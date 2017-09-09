// analyse program prints different statistics about the cost of operations with CryptoFiat contracts
package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/sha3"
)

//go:generate abigen --sol ../CryptoFiat.sol --pkg main --out abi.go

// Accounts
var (
	// Master Roles
	master          = newAccount("master", `3c3d916536295378289df64893d7cb075a223ed425018809ee73b6cb13b649b7`)
	reserveBank     = newAccount("reserve", `327dbe38fd1842424c2dac0c8c833c49e45d4f827efdb29c1f8af1877b8c7ebc`)
	accountApprover = newAccount("account approver", `10463c1b057ac60c09451a3c935ed2cc10d0fa93362c680d893a99ccc807666f`)

	lawEnforcer                 = newAccount("law enforcer", `760902010f6bc423120fe197d487a32edb7f72f104298325e4584d28d9335bde`)
	enforcmentAccountDesignator = newAccount("enforcement account designator", `bd4132d9b7d1e93553bd7bc831a85b6005de4d131b701c5d9ea8ab78a0952538`)
	enforcmentAccount           = newAccount("enforcement account", `c6a3afcca53233a72c68cd05efd350011bac46b75754b3d1a60f58c2efd70732`)

	delegationServer = newAccount("delegation server", `c6a3afcca53233a72c68cd05efd352011bac46b75754b3d1a60f58c2efd70732`)

	// Users
	alice = newAccount("alice", `9ba72a7d290bc324b9bc0999ad46b37c8e8e31c6dc850ec8d6008896a3318a6c`)
	bob   = newAccount("bob", `bc5531b4c4b2b36906e7743d872684d3c8c38f56a644f303b11605344a62f6d6`)
	carol = newAccount("carol", `c6a3aecca53233a72c68cd05efd350011bac46b75754b3d1a60f58c2efd70732`)
	erin  = newAccount("erin", `d31daef2b844db4118dfa9d1cfe6436d647efc963cca16432366e7c589e8f4ba`)
)

// Backend
var (
	initialBalance = big.NewInt(1e9)

	backend = newBackend(initialBalance,
		master,
		reserveBank,
		accountApprover,

		lawEnforcer,
		enforcmentAccountDesignator,
		enforcmentAccount,

		delegationServer,

		alice,
		bob,
		carol,
		erin,
	)

	approveAddresses = []common.Address{
		reserveBank.Address,
		enforcmentAccount.Address,

		delegationServer.Address,

		alice.Address,
		bob.Address,
		carol.Address,
		erin.Address,
	}

	userAccounts = []*account{
		alice,
		bob,
		carol,
		erin,
	}
)

var (
	DATA             = big.NewInt(1)
	ACCOUNTS         = big.NewInt(2)
	APPROVING        = big.NewInt(3)
	RESERVE          = big.NewInt(4)
	ENFORCEMENT      = big.NewInt(5)
	ACCOUNT_RECOVERY = big.NewInt(6)
	DELEGATION       = big.NewInt(7)
)

func main() {
	var stats transactions

	cryptofiatAddress, tx, cryptofiat, err := DeployCryptoFiat(master.tx(0), backend)
	stats.add("deploy", "cryptofiat", tx, err)
	backend.Commit()

	// deploy contracts
	dataAddress, tx, data, err := DeployData(master.tx(0), backend,
		cryptofiatAddress)
	stats.add("deploy", "data", tx, err)
	backend.Commit()
	_ = data

	accountsAddress, tx, accounts, err := DeployAccounts(master.tx(0), backend,
		cryptofiatAddress)
	stats.add("deploy", "accounts", tx, err)
	backend.Commit()

	approvingAddress, tx, approving, err := DeployApproving(master.tx(0), backend,
		cryptofiatAddress, accountApprover.Address)
	stats.add("deploy", "approving", tx, err)
	backend.Commit()

	reserveAddress, tx, reserve, err := DeployReserve(master.tx(0), backend,
		cryptofiatAddress, reserveBank.Address)
	stats.add("deploy", "reserve", tx, err)
	backend.Commit()

	enforcementAddress, tx, enforcement, err := DeployEnforcement(master.tx(0), backend,
		cryptofiatAddress,
		lawEnforcer.Address, enforcmentAccountDesignator.Address, enforcmentAccount.Address)
	stats.add("deploy", "enforcement", tx, err)
	backend.Commit()

	accountRecoveryAddress, tx, accountRecovery, err := DeployAccountRecovery(master.tx(0), backend,
		cryptofiatAddress)
	stats.add("deploy", "account recovery", tx, err)
	backend.Commit()

	delegationAddress, tx, delegation, err := DeployDelegation(master.tx(0), backend,
		cryptofiatAddress)
	stats.add("deploy", "delegation", tx, err)
	backend.Commit()

	{
		// upgrade contracts
		tx, err = cryptofiat.Upgrade(master.tx(0), DATA, dataAddress)
		stats.add("upgrade", "data", tx, err)
		backend.Commit()

		tx, err = cryptofiat.Upgrade(master.tx(0), ACCOUNTS, accountsAddress)
		stats.add("upgrade", "accounts", tx, err)
		backend.Commit()

		tx, err = cryptofiat.Upgrade(master.tx(0), APPROVING, approvingAddress)
		stats.add("upgrade", "approving", tx, err)
		backend.Commit()

		tx, err = cryptofiat.Upgrade(master.tx(0), RESERVE, reserveAddress)
		stats.add("upgrade", "reserve", tx, err)
		backend.Commit()

		tx, err = cryptofiat.Upgrade(master.tx(0), ENFORCEMENT, enforcementAddress)
		stats.add("upgrade", "enforcement", tx, err)
		backend.Commit()

		tx, err = cryptofiat.Upgrade(master.tx(0), ACCOUNT_RECOVERY, accountRecoveryAddress)
		stats.add("upgrade", "account recovery", tx, err)
		backend.Commit()

		tx, err = cryptofiat.Upgrade(master.tx(0), DELEGATION, delegationAddress)
		stats.add("upgrade", "delegation", tx, err)
		backend.Commit()
	}

	{ // crypto fiat functions
		tx, err = cryptofiat.AppointMasterAccount(master.tx(0), master.Address)
		stats.add("cryptofiat", "appointMasterAccount", tx, err)
		backend.Commit()
	}

	{ // approving
		tx, err = approving.ApproveAccount(accountApprover.tx(0), master.Address)
		stats.add("approving", "approve account", tx, err)
		backend.Commit()

		tx, err = approving.CloseAccount(accountApprover.tx(0), master.Address)
		stats.add("approving", "close account", tx, err)
		backend.Commit()

		tx, err = approving.ApproveAccounts(accountApprover.tx(0), approveAddresses)
		stats.add("approving", "approve accounts "+strconv.Itoa(len(approveAddresses)), tx, err)
		backend.Commit()
	}

	{ // reserve
		tx, err = reserve.AppointReserveBank(reserveBank.tx(0), reserveBank.Address)
		stats.add("reserve", "appointReserveBank", tx, err)
		backend.Commit()

		tx, err = reserve.IncreaseSupply(reserveBank.tx(0), big.NewInt(10000000))
		stats.add("reserve", "increaseSupply", tx, err)
		backend.Commit()

		tx, err = reserve.DecreaseSupply(reserveBank.tx(0), big.NewInt(100))
		stats.add("reserve", "decreaseSupply", tx, err)
		backend.Commit()
	}

	{ // distribute money as a starting point
		for _, acc := range userAccounts {
			tx, err = accounts.Transfer(reserveBank.tx(0), acc.Address, big.NewInt(100000))
			stats.add("reserve deposit", acc.Name, tx, err)
			backend.Commit()
		}

		for _, acc := range userAccounts {
			tx, err = accounts.Transfer(reserveBank.tx(0), acc.Address, big.NewInt(100000))
			stats.add("reserve send", acc.Name, tx, err)
			backend.Commit()
		}
	}

	{ // accounts
		tx, err = accounts.Transfer(alice.tx(0), bob.Address, big.NewInt(1000))
		stats.add("accounts", "alice -> bob", tx, err)
		backend.Commit()

		tx, err = accounts.Transfer(bob.tx(0), alice.Address, big.NewInt(1000))
		stats.add("accounts", "bob -> alice", tx, err)
		backend.Commit()
	}

	{ // enforcement
		tx, err = enforcement.FreezeAccount(lawEnforcer.tx(0), carol.Address)
		stats.add("enforcement", "freeze carol", tx, err)
		backend.Commit()

		tx, err = enforcement.Withdraw(lawEnforcer.tx(0), carol.Address, big.NewInt(1000))
		stats.add("enforcement", "withdraw carol", tx, err)
		backend.Commit()

		tx, err = enforcement.UnfreezeAccount(lawEnforcer.tx(0), carol.Address)
		stats.add("enforcement", "unfreeze carol", tx, err)
		backend.Commit()
	}

	{ // account recovery
		tx, err = accountRecovery.DesignateRecoveryAccount(erin.tx(0), carol.Address)
		stats.add("account recovery", "designate account erin", tx, err)
		backend.Commit()

		tx, err = accountRecovery.RecoverBalance(carol.tx(0), erin.Address, bob.Address)
		stats.add("account recovery", "recover balance", tx, err)
		backend.Commit()
	}

	{ // delegated transfer
		xfer := newDelegatedTransfer(alice, big.NewInt(1), bob.Address, big.NewInt(5000), big.NewInt(20))
		tx, err = delegation.Transfer(delegationServer.tx(0),
			xfer.nonce, xfer.destination, xfer.amount, xfer.fee, xfer.signature,
			delegationServer.Address)
		stats.add("delegation", "alice -> bob", tx, err)
		backend.Commit()

		xfer = newDelegatedTransfer(alice, big.NewInt(2), bob.Address, big.NewInt(5000), big.NewInt(20))
		tx, err = delegation.Transfer(delegationServer.tx(0),
			xfer.nonce, xfer.destination, xfer.amount, xfer.fee, xfer.signature,
			delegationServer.Address)
		stats.add("delegation", "alice -> bob", tx, err)
		backend.Commit()

		var xfers []*delegatedTransfer
		xfer = newDelegatedTransfer(alice, big.NewInt(3), bob.Address, big.NewInt(100), big.NewInt(10))
		xfers = append(xfers, xfer)
		tx, err = delegation.Multitransfer(delegationServer.tx(0),
			big.NewInt(int64(len(xfers))),
			delegatedTransfersToBytes(xfers),
			delegationServer.Address)
		stats.add("multi delegation", "1x [alice  -> bob]", tx, err)
		backend.Commit()

		xfers = []*delegatedTransfer{}
		xfer = newDelegatedTransfer(alice, big.NewInt(4), bob.Address, big.NewInt(100), big.NewInt(10))
		xfers = append(xfers, xfer)
		xfer = newDelegatedTransfer(bob, big.NewInt(1), alice.Address, big.NewInt(100), big.NewInt(10))
		xfers = append(xfers, xfer)
		xfer = newDelegatedTransfer(alice, big.NewInt(5), bob.Address, big.NewInt(100), big.NewInt(10))
		xfers = append(xfers, xfer)
		xfer = newDelegatedTransfer(bob, big.NewInt(2), alice.Address, big.NewInt(100), big.NewInt(10))
		xfers = append(xfers, xfer)

		tx, err = delegation.Multitransfer(delegationServer.tx(0),
			big.NewInt(int64(len(xfers))),
			delegatedTransfersToBytes(xfers),
			delegationServer.Address)
		stats.add("multi delegation", "4x [alice <-> bob]", tx, err)
		backend.Commit()
	}

	{ // print final stats
		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 4, 8, 4, ' ', 0)
		defer w.Flush()

		fmt.Fprintf(w, "Category\tTransaction\tGasUsed\t1gwei\t5gwei\t20gwei\t40gwei\tError\n")
		prevcat := ""
		for _, transaction := range stats {
			category := transaction.category
			if prevcat == transaction.category {
				category = ""
			}
			prevcat = transaction.category

			if transaction.err != nil {
				fmt.Fprintf(w, "%v\t%v\t-\t-\t-\t-\t-\t%v\n", category, transaction.name, transaction.err)
				continue
			}

			rcpt, err := backend.TransactionReceipt(context.TODO(), transaction.tx.Hash())
			check(err)

			used := rcpt.GasUsed
			usedUpGas := ""
			if used.Cmp(big.NewInt(4000000)) > 0 {
				// when something fails it will use up all the gas which is ~4e6
				usedUpGas = "gas exhausted"
			}
			if rcpt.Failed {
				usedUpGas = "failed"
			}

			fmt.Fprintf(w, "%v\t%v\t%v\t%0.2f€\t%0.2f€\t%0.2f€\t%0.2f€\t%v\n",
				category, transaction.name,
				used, price(used, 1, EthEur), price(used, 5, EthEur), price(used, 20, EthEur), price(used, 40, EthEur),
				usedUpGas)
		}
	}
}

type delegatedTransfer struct {
	nonce       *big.Int
	destination common.Address
	amount      *big.Int
	fee         *big.Int
	signature   []byte
}

func newDelegatedTransfer(
	source *account, nonce *big.Int,
	destination common.Address, amount *big.Int,
	fee *big.Int) *delegatedTransfer {

	xfer := &delegatedTransfer{}
	xfer.nonce = nonce
	xfer.destination = destination
	xfer.amount = amount
	xfer.fee = fee

	hw := sha3.NewKeccak256()
	hw.Write(big256(xfer.nonce))
	hw.Write(xfer.destination[:])
	hw.Write(big256(xfer.amount))
	hw.Write(big256(xfer.fee))
	hash := hw.Sum(nil)

	sig, err := crypto.Sign(hash[:], source.Key)
	check(err)
	xfer.signature = sig

	return xfer
}

func (xfer *delegatedTransfer) Bytes() []byte {
	var data []byte
	data = append(data, big256(xfer.nonce)...)
	data = append(data, xfer.destination[:]...)
	data = append(data, big256(xfer.amount)...)
	data = append(data, big256(xfer.fee)...)
	data = append(data, xfer.signature...)
	return data
}

func delegatedTransfersToBytes(xfers []*delegatedTransfer) []byte {
	var data []byte
	for _, xfer := range xfers {
		data = append(data, xfer.Bytes()...)
	}
	return data
}

func big256(v *big.Int) []byte {
	var r [32]byte
	data := v.Bytes()
	off := 32 - len(data)
	copy(r[off:], data)
	return r[:]
}

const EthEur = 250

func price(gas *big.Int, gwei, ethfiat float64) float64 {
	return float64(gas.Int64()) * gwei * ethfiat * 1e9 / 1e18
}

type transactions []transaction

func (txs *transactions) add(category, name string, tx *types.Transaction, err error) {
	*txs = append(*txs, transaction{category, name, tx, err})
}

type transaction struct {
	category string
	name     string
	tx       *types.Transaction
	err      error
}

type account struct {
	Name    string
	Address common.Address
	Key     *ecdsa.PrivateKey
}

func newAccount(name string, hex string) *account {
	key, err := crypto.HexToECDSA(hex)
	check(err)
	addr := crypto.PubkeyToAddress(key.PublicKey)
	return &account{
		Name:    name,
		Address: addr,
		Key:     key,
	}
}

func (acc *account) tx(amount int64) *bind.TransactOpts {
	tx := bind.NewKeyedTransactor(acc.Key)
	tx.Value = big.NewInt(amount)
	return tx
}

func newBackend(initialBalance *big.Int, accounts ...*account) *backends.SimulatedBackend {
	alloc := core.GenesisAlloc{}
	for _, acc := range accounts {
		alloc[acc.Address] = core.GenesisAccount{
			Balance: initialBalance,
		}
	}
	return backends.NewSimulatedBackend(alloc)
}
