// +build ignore

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/tabwriter"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/ethereum/go-ethereum/core/vm"
)

// deployed root contract
const CryptoEuroAddress = "0xB8E7245B42529B905a8909B8FD5fC690097e71f7"

var SubConstractID = map[string]int{
	"Data":            0,
	"Accounts":        1,
	"Approving":       2,
	"Reserve":         3,
	"Enforcement":     4,
	"AccountRecovery": 5,
	"Delegation":      6,
}

var ContractOrder = []string{
	"CryptoFiat",
	"Data",
	"Accounts",
	"Approving",
	"Reserve",
	"Enforcement",
	"AccountRecovery",
	"Delegation",
}

var Assembly = map[string]*Asm{}

var (
	bindir = flag.String("bin", ".bin", "compilation output")
	abidir = flag.String("abi", "abi", "abi directory")

	solc = flag.String("solc", "solc", "compiler")
)

func main() {
	flag.Parse()

	contracts, err := compile(*solc, "CryptoFiat.sol")
	if err != nil {
		fmt.Printf("unable to compile: %v\n", err)
		os.Exit(-1)
	}

	err = WriteBind(contracts, "contract", "abi.go", bind.LangGo)
	if err != nil {
		fmt.Printf("go binding: %v\n", err)
		os.Exit(-1)
	}

	err = WriteBind(contracts, "eu.cryptoeuro.contract", "abi.java", bind.LangJava)
	if err != nil {
		fmt.Printf("write java binding: %v\n", err)
		os.Exit(-1)
	}

	err = WriteBindJS(contracts, "cryptoeuro", "abi.js")
	if err != nil {
		fmt.Printf("write js binding: %v\n", err)
		os.Exit(-1)
	}

	err = WriteABI(contracts, "abi")
	if err != nil {
		fmt.Printf("write abi: %v\n", err)
		os.Exit(-1)
	}

	err = WriteDetails(contracts, ".bin")
	if err != nil {
		fmt.Printf("write details: %v\n", err)
		os.Exit(-1)
	}

	err = GenerateWeb3J(contracts, "java")
	if err != nil {
		fmt.Printf("generate java abi: %v\n", err)
		os.Exit(-1)
	}
}

func WriteBind(contracts map[string]*compiler.Contract, pkg, file string, lang bind.Lang) error {
	var (
		abis  []string
		bins  []string
		types []string
	)

	for _, name := range ContractOrder {
		contract := contracts[name]

		abi, _ := json.Marshal(contract.Info.AbiDefinition)
		abis = append(abis, string(abi))
		bins = append(bins, contract.Code)
		types = append(types, name)
	}

	code, err := bind.Bind(types, abis, bins, pkg, lang)
	if err != nil {
		return fmt.Errorf("failed to create binding: %v", err)
	}

	err = ioutil.WriteFile(file, []byte(code), 0600)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	return nil
}

func WriteBindJS(contracts map[string]*compiler.Contract, pkg, file string) error {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "var Contract = {};\n")
	writeContract := func(name string, contract *compiler.Contract) {
		abi, _ := json.Marshal(contract.Info.AbiDefinition)
		fmt.Fprintf(buf, "Contract[%q] = web3.eth.contract(%v);\n", name, string(abi))
		if name == "CryptoFiat" {
			fmt.Fprintf(buf, "var %v = Contract.CryptoFiat.at(%q);\n", name, CryptoEuroAddress)
		} else {
			fmt.Fprintf(buf, "var %v = Contract.%v.at(CryptoFiat.contracts(%v));\n", name, name, SubConstractID[name])
		}
		fmt.Fprintf(buf, "\n")
	}

	for _, name := range ContractOrder {
		writeContract(name, contracts[name])
	}

	err := ioutil.WriteFile(file, buf.Bytes(), 0600)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	return nil
}

func WriteABI(contracts map[string]*compiler.Contract, folder string) error {
	for _, name := range ContractOrder {
		contract := contracts[name]

		abi, _ := json.Marshal(contract.Info.AbiDefinition)
		filename := filepath.Join(folder, name) + ".abi"
		err := ioutil.WriteFile(filename, abi, 0600)
		if err != nil {
			return err
		}
	}
	return nil
}

func WriteDetails(contracts map[string]*compiler.Contract, folder string) error {
	os.MkdirAll(folder, 0755)

	write := func(name, ext string, data string) error {
		filename := filepath.Join(folder, name) + ext
		return ioutil.WriteFile(filename, []byte(data), 0600)
	}

	for _, name := range ContractOrder {
		contract := contracts[name]

		abi, _ := json.MarshalIndent(contract.Info.AbiDefinition, "", "\t")
		if err := write(name, ".abi", string(abi)); err != nil {
			return err
		}

		if err := write(name, ".bin", contract.Code[2:]); err != nil {
			return err
		}

		if err := write(name, ".evm", Assembly[name].String()); err != nil {
			return err
		}
	}
	return nil
}

func GenerateWeb3J(contracts map[string]*compiler.Contract, folder string) error {
	os.MkdirAll(folder, 0755)

	for _, name := range ContractOrder {
		cmd := exec.Command("web3j",
			"solidity", "generate",
			"--javaTypes",
			filepath.Join(".bin", name+".bin"),
			filepath.Join(".bin", name+".abi"),
			"-o", folder,
			"-p", "eu.cryptoeuro.contract",
		)
		if err := cmd.Run(); err != nil {
			fmt.Println("uanble to write abi for", name, ":", err)
			return err
		}
	}
	return nil
}

// --combined-output format
type solcOutput struct {
	Contracts map[string]struct {
		Bin, Abi, Devdoc, Userdoc, Metadata string

		Asm *Asm
	}
	Version string
}

type Asm struct {
	Source string
	Code   []Instr `json:".code"`
}

type Instr struct {
	Name     string `json:"name"`
	Begin    int    `json:"begin"`
	End      int    `json:"end"`
	Value    string `json:"value"`
	JumpType string `json:"jumpType"`
}

func (asm *Asm) String() string {
	buf := new(bytes.Buffer)
	w := new(tabwriter.Writer)
	w.Init(buf, 4, 8, 4, ' ', 0)

	for _, instr := range asm.Code {
		opcode := vm.StringToOp(instr.Name)

		if instr.Begin == -1 && instr.End == -1 {
			fmt.Fprint(w, "\t")
		} else {
			fmt.Fprintf(w, "%v:%v\t", instr.Begin, instr.End)
		}
		if cost := gasCost(opcode); cost != UNKNOWN {
			fmt.Fprintf(w, "%v\t", cost)
		} else {
			fmt.Fprint(w, "\t")
		}
		fmt.Fprint(w, instr.Name+"\t")
		fmt.Fprint(w, instr.Value+"\t")
		fmt.Fprint(w, instr.JumpType+"\t")
		fmt.Fprint(w, "\n")
	}
	w.Flush()

	lines := strings.Split(buf.String(), "\n")
	for i, line := range lines {
		lines[i] = strings.TrimRight(line, " \n")
	}
	return strings.Join(lines, "\n") + "\n"
}

func compile(solc string, sourcefiles ...string) (map[string]*compiler.Contract, error) {
	if len(sourcefiles) == 0 {
		return nil, errors.New("solc: no source files")
	}

	source, err := slurpFiles(sourcefiles)
	if err != nil {
		return nil, err
	}

	s, err := compiler.SolidityVersion(solc)
	if err != nil {
		return nil, err
	}

	args := []string{
		"--combined-json", "asm,bin,abi,userdoc,devdoc,metadata",
		"--optimize",                 // code optimizer switched on
		"--optimize-runs", "1000000", // optimize as much as possible
		"--",
	}
	cmd := exec.Command(s.Path, append(args, sourcefiles...)...)
	return runCompiler(s, args, cmd, source)
}

func runCompiler(s *compiler.Solidity, args []string, cmd *exec.Cmd, source string) (map[string]*compiler.Contract, error) {
	var stderr, stdout bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("solc: %v\n%s", err, stderr.Bytes())
	}
	var output solcOutput
	if err := json.Unmarshal(stdout.Bytes(), &output); err != nil {
		return nil, err
	}

	// Compilation succeeded, assemble and return the contracts.
	contracts := make(map[string]*compiler.Contract)
	for name, info := range output.Contracts {
		parts := strings.Split(name, ":")
		name = parts[len(parts)-1]

		// Parse the individual compilation results.
		var abi interface{}
		if err := json.Unmarshal([]byte(info.Abi), &abi); err != nil {
			return nil, fmt.Errorf("solc: error reading abi definition (%v)", err)
		}
		var userdoc interface{}
		if err := json.Unmarshal([]byte(info.Userdoc), &userdoc); err != nil {
			return nil, fmt.Errorf("solc: error reading user doc: %v", err)
		}
		var devdoc interface{}
		if err := json.Unmarshal([]byte(info.Devdoc), &devdoc); err != nil {
			return nil, fmt.Errorf("solc: error reading dev doc: %v", err)
		}

		info.Asm.Source = source
		Assembly[name] = info.Asm
		contracts[name] = &compiler.Contract{
			Code: "0x" + info.Bin,
			Info: compiler.ContractInfo{
				Source:          source,
				Language:        "Solidity",
				LanguageVersion: s.Version,
				CompilerVersion: s.Version,
				CompilerOptions: strings.Join(args, " "),
				AbiDefinition:   abi,
				UserDoc:         userdoc,
				DeveloperDoc:    devdoc,
				Metadata:        info.Metadata,
			},
		}
	}
	return contracts, nil
}

func slurpFiles(files []string) (string, error) {
	var concat bytes.Buffer
	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			return "", err
		}
		concat.Write(content)
	}
	return concat.String(), nil
}

func checkf(err error, format string) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, format+"\n", err)
	os.Exit(1)
}

const UNKNOWN = 0xFFFFFFFFFFFFFFFF

func gasCost(opcode vm.OpCode) uint64 {
	switch opcode {
	case vm.RETURNDATASIZE:
		return vm.GasQuickStep
	case vm.STOP:
		return 0
	case vm.ADD:
		return vm.GasFastestStep
	case vm.MUL:
		return vm.GasFastStep
	case vm.SUB:
		return vm.GasFastestStep
	case vm.DIV:
		return vm.GasFastStep
	case vm.SDIV:
		return vm.GasFastStep
	case vm.MOD:
		return vm.GasFastStep
	case vm.SMOD:
		return vm.GasFastStep
	case vm.ADDMOD:
		return vm.GasMidStep
	case vm.MULMOD:
		return vm.GasMidStep
	case vm.SIGNEXTEND:
		return vm.GasFastStep
	case vm.LT:
		return vm.GasFastestStep
	case vm.GT:
		return vm.GasFastestStep
	case vm.SLT:
		return vm.GasFastestStep
	case vm.SGT:
		return vm.GasFastestStep
	case vm.EQ:
		return vm.GasFastestStep
	case vm.ISZERO:
		return vm.GasFastestStep
	case vm.AND:
		return vm.GasFastestStep
	case vm.XOR:
		return vm.GasFastestStep
	case vm.OR:
		return vm.GasFastestStep
	case vm.NOT:
		return vm.GasFastestStep
	case vm.BYTE:
		return vm.GasFastestStep
	case vm.ADDRESS:
		return vm.GasQuickStep
	case vm.ORIGIN:
		return vm.GasQuickStep
	case vm.CALLER:
		return vm.GasQuickStep
	case vm.CALLVALUE:
		return vm.GasQuickStep
	case vm.CALLDATALOAD:
		return vm.GasFastestStep
	case vm.CALLDATASIZE:
		return vm.GasQuickStep
	case vm.CODESIZE:
		return vm.GasQuickStep
	case vm.GASPRICE:
		return vm.GasQuickStep
	case vm.BLOCKHASH:
		return vm.GasExtStep
	case vm.COINBASE:
		return vm.GasQuickStep
	case vm.TIMESTAMP:
		return vm.GasQuickStep
	case vm.NUMBER:
		return vm.GasQuickStep
	case vm.DIFFICULTY:
		return vm.GasQuickStep
	case vm.GASLIMIT:
		return vm.GasQuickStep
	case vm.POP:
		return vm.GasQuickStep
	case vm.JUMP:
		return vm.GasMidStep
	case vm.JUMPI:
		return vm.GasSlowStep
	case vm.PC:
		return vm.GasQuickStep
	case vm.MSIZE:
		return vm.GasQuickStep
	case vm.GAS:
		return vm.GasQuickStep
	case vm.JUMPDEST:
		return 1
	case vm.PUSH1:
		return vm.GasFastestStep
	case vm.PUSH2:
		return vm.GasFastestStep
	case vm.PUSH3:
		return vm.GasFastestStep
	case vm.PUSH4:
		return vm.GasFastestStep
	case vm.PUSH5:
		return vm.GasFastestStep
	case vm.PUSH6:
		return vm.GasFastestStep
	case vm.PUSH7:
		return vm.GasFastestStep
	case vm.PUSH8:
		return vm.GasFastestStep
	case vm.PUSH9:
		return vm.GasFastestStep
	case vm.PUSH10:
		return vm.GasFastestStep
	case vm.PUSH11:
		return vm.GasFastestStep
	case vm.PUSH12:
		return vm.GasFastestStep
	case vm.PUSH13:
		return vm.GasFastestStep
	case vm.PUSH14:
		return vm.GasFastestStep
	case vm.PUSH15:
		return vm.GasFastestStep
	case vm.PUSH16:
		return vm.GasFastestStep
	case vm.PUSH17:
		return vm.GasFastestStep
	case vm.PUSH18:
		return vm.GasFastestStep
	case vm.PUSH19:
		return vm.GasFastestStep
	case vm.PUSH20:
		return vm.GasFastestStep
	case vm.PUSH21:
		return vm.GasFastestStep
	case vm.PUSH22:
		return vm.GasFastestStep
	case vm.PUSH23:
		return vm.GasFastestStep
	case vm.PUSH24:
		return vm.GasFastestStep
	case vm.PUSH25:
		return vm.GasFastestStep
	case vm.PUSH26:
		return vm.GasFastestStep
	case vm.PUSH27:
		return vm.GasFastestStep
	case vm.PUSH28:
		return vm.GasFastestStep
	case vm.PUSH29:
		return vm.GasFastestStep
	case vm.PUSH30:
		return vm.GasFastestStep
	case vm.PUSH31:
		return vm.GasFastestStep
	case vm.PUSH32:
		return vm.GasFastestStep
	case vm.DUP1:
		return vm.GasFastestStep
	case vm.DUP2:
		return vm.GasFastestStep
	case vm.DUP3:
		return vm.GasFastestStep
	case vm.DUP4:
		return vm.GasFastestStep
	case vm.DUP5:
		return vm.GasFastestStep
	case vm.DUP6:
		return vm.GasFastestStep
	case vm.DUP7:
		return vm.GasFastestStep
	case vm.DUP8:
		return vm.GasFastestStep
	case vm.DUP9:
		return vm.GasFastestStep
	case vm.DUP10:
		return vm.GasFastestStep
	case vm.DUP11:
		return vm.GasFastestStep
	case vm.DUP12:
		return vm.GasFastestStep
	case vm.DUP13:
		return vm.GasFastestStep
	case vm.DUP14:
		return vm.GasFastestStep
	case vm.DUP15:
		return vm.GasFastestStep
	case vm.DUP16:
		return vm.GasFastestStep
	case vm.SWAP1:
		return vm.GasFastestStep
	case vm.SWAP2:
		return vm.GasFastestStep
	case vm.SWAP3:
		return vm.GasFastestStep
	case vm.SWAP4:
		return vm.GasFastestStep
	case vm.SWAP5:
		return vm.GasFastestStep
	case vm.SWAP6:
		return vm.GasFastestStep
	case vm.SWAP7:
		return vm.GasFastestStep
	case vm.SWAP8:
		return vm.GasFastestStep
	case vm.SWAP9:
		return vm.GasFastestStep
	case vm.SWAP10:
		return vm.GasFastestStep
	case vm.SWAP11:
		return vm.GasFastestStep
	case vm.SWAP12:
		return vm.GasFastestStep
	case vm.SWAP13:
		return vm.GasFastestStep
	case vm.SWAP14:
		return vm.GasFastestStep
	case vm.SWAP15:
		return vm.GasFastestStep
	case vm.SWAP16:
		return vm.GasFastestStep
	}
	return UNKNOWN
}
