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
	"strings"
	"text/template"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/compiler"
)

// deployed root contract
const rootAddress = "0xB8E7245B42529B905a8909B8FD5fC690097e71f7"

var (
	bindir = flag.String("bin", ".bin", "compilation output")
	abidir = flag.String("abi", "abi", "abi directory")

	solc = flag.String("solc", "solc", "compiler")
)

func main() {
	flag.Parse()

	originalContracts, err := compile(*solc, "CryptoFiat.sol")
	if err != nil {
		fmt.Printf("unable to compile: %v\n", err)
		os.Exit(-1)
	}
	contracts := make(map[string]*compiler.Contract)
	for name, contract := range originalContracts {
		parts := strings.Split(name, ":")
		typeName := parts[len(parts)-1]
		contracts[typeName] = contract
	}

	// remove abstract contracts
	delete(contracts, "Constants")
	delete(contracts, "Relay")
	delete(contracts, "InternalData")

	err = WriteBind(contracts, "contract", "abi.go", bind.LangGo)
	if err != nil {
		fmt.Printf("go binding: %v\n", err)
		os.Exit(-1)
	}

	err = WriteBind(contracts, "eu.cryptoeuro.contract", "abi.java", bind.LangJava)
	if err != nil {
		fmt.Printf("java binding: %v\n", err)
		os.Exit(-1)
	}
}

func WriteBind(contracts map[string]*compiler.Contract, pkg, file string, lang bind.Lang) error {
	var (
		abis  []string
		bins  []string
		types []string
	)
	for name, contract := range contracts {
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

// --combined-output format
type solcOutput struct {
	Contracts map[string]struct {
		Bin, Abi, Devdoc, Userdoc, Metadata string
	}
	Version string
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
		"--combined-json", "bin,abi,userdoc,devdoc,metadata",
		"--add-std",                  // include standard lib contracts
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

var SubConstractID = map[string]int{
	"Data":            0,
	"Accounts":        1,
	"Approving":       2,
	"Reserve":         3,
	"Enforcement":     4,
	"AccountRecovery": 5,
	"Delegation":      6,
	"MultiDelegation": 7,
}

/*
func WriteJS(contracts map) {
	os.MkdirAll(*abidir, 0755)

	matches, err := filepath.Glob(filepath.Join(*bindir, "*.abi"))
	checkf(err, "filepath.Glob: %v")

	type Contract struct {
		ID   int
		Name string
		ABI  string
	}

	root := Contract{}
	contracts := []Contract{}

	for _, match := range matches {
		data, err := ioutil.ReadFile(match)
		checkf(err, "ReadFile: %v")

		file := filepath.Base(match)
		ext := filepath.Ext(file)
		name := file[:len(file)-len(ext)]

		contract := Contract{}
		contract.Name = name
		contract.ABI = string(data)
		if contract.Name == "CryptoFiat" {
			contract.ID = -1
			root = contract
		} else {
			id, ok := SubConstractID[name]
			if !ok {
				continue
			}
			contract.ID = id
			contracts = append(contracts, contract)
		}

		ioutil.WriteFile(filepath.Join(*abidir, file), data, 0755)
	}

	sort.Slice(contracts, func(i, k int) bool {
		return contracts[i].ID < contracts[k].ID
	})

	f, err := os.Create(filepath.Join(*abidir, "cryptofiat.abi.js"))
	checkf(err, "create file: %v")
	defer f.Close()

	err = jsabi_template.Execute(f, map[string]interface{}{
		"rootAddress": rootAddress,
		"root":        root,
		"contracts":   contracts,
	})
	checkf(err, "template: %v")
}
*/
var jsabi_template = template.Must(template.New("").Parse(`
// GENERATED CODE
// DO NOT MODIFY

var Contract = {};

Contract['{{.root.Name}}'] = web3.eth.contract({{.root.ABI}});
{{ range .contracts }}
Contract['{{.Name}}'] = web3.eth.contract({{.ABI}});
{{- end }}

var CryptoFiat = Contract.CryptoFiat.at('{{.rootAddress}}');
{{ range .contracts }}
var {{ .Name }} = Contract.Accounts.at(CryptoFiat.contracts({{ .ID }}));
{{- end }}
`))

func solcx(args ...string) ([]byte, error) {
	return exec.Command("solc", args...).CombinedOutput()
}
