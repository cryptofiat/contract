package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"text/template"
)

// deployed root contract
const rootAddress = "0xB8E7245B42529B905a8909B8FD5fC690097e71f7"

var (
	bindir = flag.String("bin", ".bin", "compilation output")
	abidir = flag.String("abi", "abi", "abi directory")
)

func main() {
	flag.Parse()
	compile()
	jsabi()
}

func checkf(err error, format string) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, format+"\n", err)
	os.Exit(1)
}

func compile() {
	os.MkdirAll(*bindir, 0755)
	data, err := solc(
		"-o", *bindir,
		"--bin", "--abi", "--asm", "--gas",
		"--optimize", "--optimize-runs", "100000",
		"--overwrite",

		"CryptoFiat.sol",
	)

	checkf(err, "solc failed: %v \n"+string(data))
	ioutil.WriteFile(filepath.Join(*bindir, "_gas.log"), data, 0755)
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

func jsabi() {
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

func solc(args ...string) ([]byte, error) {
	return exec.Command("solc", args...).CombinedOutput()
}
