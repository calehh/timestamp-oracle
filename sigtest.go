// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SigtestABI is the input ABI used to generate the binding from.
const SigtestABI = "[{\"inputs\":[],\"name\":\"auth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lasttimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Sigtest is an auto generated Go binding around an Ethereum contract.
type Sigtest struct {
	SigtestCaller     // Read-only binding to the contract
	SigtestTransactor // Write-only binding to the contract
	SigtestFilterer   // Log filterer for contract events
}

// SigtestCaller is an auto generated read-only Go binding around an Ethereum contract.
type SigtestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigtestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SigtestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigtestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SigtestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SigtestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SigtestSession struct {
	Contract     *Sigtest          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SigtestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SigtestCallerSession struct {
	Contract *SigtestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SigtestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SigtestTransactorSession struct {
	Contract     *SigtestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SigtestRaw is an auto generated low-level Go binding around an Ethereum contract.
type SigtestRaw struct {
	Contract *Sigtest // Generic contract binding to access the raw methods on
}

// SigtestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SigtestCallerRaw struct {
	Contract *SigtestCaller // Generic read-only contract binding to access the raw methods on
}

// SigtestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SigtestTransactorRaw struct {
	Contract *SigtestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSigtest creates a new instance of Sigtest, bound to a specific deployed contract.
func NewSigtest(address common.Address, backend bind.ContractBackend) (*Sigtest, error) {
	contract, err := bindSigtest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sigtest{SigtestCaller: SigtestCaller{contract: contract}, SigtestTransactor: SigtestTransactor{contract: contract}, SigtestFilterer: SigtestFilterer{contract: contract}}, nil
}

// NewSigtestCaller creates a new read-only instance of Sigtest, bound to a specific deployed contract.
func NewSigtestCaller(address common.Address, caller bind.ContractCaller) (*SigtestCaller, error) {
	contract, err := bindSigtest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SigtestCaller{contract: contract}, nil
}

// NewSigtestTransactor creates a new write-only instance of Sigtest, bound to a specific deployed contract.
func NewSigtestTransactor(address common.Address, transactor bind.ContractTransactor) (*SigtestTransactor, error) {
	contract, err := bindSigtest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SigtestTransactor{contract: contract}, nil
}

// NewSigtestFilterer creates a new log filterer instance of Sigtest, bound to a specific deployed contract.
func NewSigtestFilterer(address common.Address, filterer bind.ContractFilterer) (*SigtestFilterer, error) {
	contract, err := bindSigtest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SigtestFilterer{contract: contract}, nil
}

// bindSigtest binds a generic wrapper to an already deployed contract.
func bindSigtest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SigtestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sigtest *SigtestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sigtest.Contract.SigtestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sigtest *SigtestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sigtest.Contract.SigtestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sigtest *SigtestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sigtest.Contract.SigtestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sigtest *SigtestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sigtest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sigtest *SigtestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sigtest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sigtest *SigtestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sigtest.Contract.contract.Transact(opts, method, params...)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Sigtest *SigtestCaller) Auth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Sigtest.contract.Call(opts, &out, "auth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Sigtest *SigtestSession) Auth() (common.Address, error) {
	return _Sigtest.Contract.Auth(&_Sigtest.CallOpts)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() view returns(address)
func (_Sigtest *SigtestCallerSession) Auth() (common.Address, error) {
	return _Sigtest.Contract.Auth(&_Sigtest.CallOpts)
}

// Lasttimestamp is a free data retrieval call binding the contract method 0x6747401c.
//
// Solidity: function lasttimestamp() view returns(uint256)
func (_Sigtest *SigtestCaller) Lasttimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Sigtest.contract.Call(opts, &out, "lasttimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Lasttimestamp is a free data retrieval call binding the contract method 0x6747401c.
//
// Solidity: function lasttimestamp() view returns(uint256)
func (_Sigtest *SigtestSession) Lasttimestamp() (*big.Int, error) {
	return _Sigtest.Contract.Lasttimestamp(&_Sigtest.CallOpts)
}

// Lasttimestamp is a free data retrieval call binding the contract method 0x6747401c.
//
// Solidity: function lasttimestamp() view returns(uint256)
func (_Sigtest *SigtestCallerSession) Lasttimestamp() (*big.Int, error) {
	return _Sigtest.Contract.Lasttimestamp(&_Sigtest.CallOpts)
}

// Verify is a paid mutator transaction binding the contract method 0x85852ce4.
//
// Solidity: function verify(uint256 timestamp, bytes signature) returns()
func (_Sigtest *SigtestTransactor) Verify(opts *bind.TransactOpts, timestamp *big.Int, signature []byte) (*types.Transaction, error) {
	return _Sigtest.contract.Transact(opts, "verify", timestamp, signature)
}

// Verify is a paid mutator transaction binding the contract method 0x85852ce4.
//
// Solidity: function verify(uint256 timestamp, bytes signature) returns()
func (_Sigtest *SigtestSession) Verify(timestamp *big.Int, signature []byte) (*types.Transaction, error) {
	return _Sigtest.Contract.Verify(&_Sigtest.TransactOpts, timestamp, signature)
}

// Verify is a paid mutator transaction binding the contract method 0x85852ce4.
//
// Solidity: function verify(uint256 timestamp, bytes signature) returns()
func (_Sigtest *SigtestTransactorSession) Verify(timestamp *big.Int, signature []byte) (*types.Transaction, error) {
	return _Sigtest.Contract.Verify(&_Sigtest.TransactOpts, timestamp, signature)
}
