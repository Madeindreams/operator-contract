// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iquoter

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IquoterMetaData contains all meta data concerning the Iquoter contract.
var IquoterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"quoteExactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"name\":\"quoteExactInputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"name\":\"quoteExactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"name\":\"quoteExactOutputSingle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IquoterABI is the input ABI used to generate the binding from.
// Deprecated: Use IquoterMetaData.ABI instead.
var IquoterABI = IquoterMetaData.ABI

// Iquoter is an auto generated Go binding around an Ethereum contract.
type Iquoter struct {
	IquoterCaller     // Read-only binding to the contract
	IquoterTransactor // Write-only binding to the contract
	IquoterFilterer   // Log filterer for contract events
}

// IquoterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IquoterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IquoterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IquoterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IquoterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IquoterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IquoterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IquoterSession struct {
	Contract     *Iquoter          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IquoterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IquoterCallerSession struct {
	Contract *IquoterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IquoterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IquoterTransactorSession struct {
	Contract     *IquoterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IquoterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IquoterRaw struct {
	Contract *Iquoter // Generic contract binding to access the raw methods on
}

// IquoterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IquoterCallerRaw struct {
	Contract *IquoterCaller // Generic read-only contract binding to access the raw methods on
}

// IquoterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IquoterTransactorRaw struct {
	Contract *IquoterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIquoter creates a new instance of Iquoter, bound to a specific deployed contract.
func NewIquoter(address common.Address, backend bind.ContractBackend) (*Iquoter, error) {
	contract, err := bindIquoter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Iquoter{IquoterCaller: IquoterCaller{contract: contract}, IquoterTransactor: IquoterTransactor{contract: contract}, IquoterFilterer: IquoterFilterer{contract: contract}}, nil
}

// NewIquoterCaller creates a new read-only instance of Iquoter, bound to a specific deployed contract.
func NewIquoterCaller(address common.Address, caller bind.ContractCaller) (*IquoterCaller, error) {
	contract, err := bindIquoter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IquoterCaller{contract: contract}, nil
}

// NewIquoterTransactor creates a new write-only instance of Iquoter, bound to a specific deployed contract.
func NewIquoterTransactor(address common.Address, transactor bind.ContractTransactor) (*IquoterTransactor, error) {
	contract, err := bindIquoter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IquoterTransactor{contract: contract}, nil
}

// NewIquoterFilterer creates a new log filterer instance of Iquoter, bound to a specific deployed contract.
func NewIquoterFilterer(address common.Address, filterer bind.ContractFilterer) (*IquoterFilterer, error) {
	contract, err := bindIquoter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IquoterFilterer{contract: contract}, nil
}

// bindIquoter binds a generic wrapper to an already deployed contract.
func bindIquoter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IquoterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iquoter *IquoterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iquoter.Contract.IquoterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iquoter *IquoterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iquoter.Contract.IquoterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iquoter *IquoterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iquoter.Contract.IquoterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Iquoter *IquoterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Iquoter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Iquoter *IquoterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Iquoter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Iquoter *IquoterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Iquoter.Contract.contract.Transact(opts, method, params...)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut)
func (_Iquoter *IquoterTransactor) QuoteExactInput(opts *bind.TransactOpts, path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _Iquoter.contract.Transact(opts, "quoteExactInput", path, amountIn)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut)
func (_Iquoter *IquoterSession) QuoteExactInput(path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _Iquoter.Contract.QuoteExactInput(&_Iquoter.TransactOpts, path, amountIn)
}

// QuoteExactInput is a paid mutator transaction binding the contract method 0xcdca1753.
//
// Solidity: function quoteExactInput(bytes path, uint256 amountIn) returns(uint256 amountOut)
func (_Iquoter *IquoterTransactorSession) QuoteExactInput(path []byte, amountIn *big.Int) (*types.Transaction, error) {
	return _Iquoter.Contract.QuoteExactInput(&_Iquoter.TransactOpts, path, amountIn)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0xf7729d43.
//
// Solidity: function quoteExactInputSingle(address tokenIn, address tokenOut, uint24 fee, uint256 amountIn, uint160 sqrtPriceLimitX96) returns(uint256 amountOut)
func (_Iquoter *IquoterTransactor) QuoteExactInputSingle(opts *bind.TransactOpts, tokenIn common.Address, tokenOut common.Address, fee *big.Int, amountIn *big.Int, sqrtPriceLimitX96 *big.Int) (*types.Transaction, error) {
	return _Iquoter.contract.Transact(opts, "quoteExactInputSingle", tokenIn, tokenOut, fee, amountIn, sqrtPriceLimitX96)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0xf7729d43.
//
// Solidity: function quoteExactInputSingle(address tokenIn, address tokenOut, uint24 fee, uint256 amountIn, uint160 sqrtPriceLimitX96) returns(uint256 amountOut)
func (_Iquoter *IquoterSession) QuoteExactInputSingle(tokenIn common.Address, tokenOut common.Address, fee *big.Int, amountIn *big.Int, sqrtPriceLimitX96 *big.Int) (*types.Transaction, error) {
	return _Iquoter.Contract.QuoteExactInputSingle(&_Iquoter.TransactOpts, tokenIn, tokenOut, fee, amountIn, sqrtPriceLimitX96)
}

// QuoteExactInputSingle is a paid mutator transaction binding the contract method 0xf7729d43.
//
// Solidity: function quoteExactInputSingle(address tokenIn, address tokenOut, uint24 fee, uint256 amountIn, uint160 sqrtPriceLimitX96) returns(uint256 amountOut)
func (_Iquoter *IquoterTransactorSession) QuoteExactInputSingle(tokenIn common.Address, tokenOut common.Address, fee *big.Int, amountIn *big.Int, sqrtPriceLimitX96 *big.Int) (*types.Transaction, error) {
	return _Iquoter.Contract.QuoteExactInputSingle(&_Iquoter.TransactOpts, tokenIn, tokenOut, fee, amountIn, sqrtPriceLimitX96)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn)
func (_Iquoter *IquoterTransactor) QuoteExactOutput(opts *bind.TransactOpts, path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _Iquoter.contract.Transact(opts, "quoteExactOutput", path, amountOut)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn)
func (_Iquoter *IquoterSession) QuoteExactOutput(path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _Iquoter.Contract.QuoteExactOutput(&_Iquoter.TransactOpts, path, amountOut)
}

// QuoteExactOutput is a paid mutator transaction binding the contract method 0x2f80bb1d.
//
// Solidity: function quoteExactOutput(bytes path, uint256 amountOut) returns(uint256 amountIn)
func (_Iquoter *IquoterTransactorSession) QuoteExactOutput(path []byte, amountOut *big.Int) (*types.Transaction, error) {
	return _Iquoter.Contract.QuoteExactOutput(&_Iquoter.TransactOpts, path, amountOut)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0x30d07f21.
//
// Solidity: function quoteExactOutputSingle(address tokenIn, address tokenOut, uint24 fee, uint256 amountOut, uint160 sqrtPriceLimitX96) returns(uint256 amountIn)
func (_Iquoter *IquoterTransactor) QuoteExactOutputSingle(opts *bind.TransactOpts, tokenIn common.Address, tokenOut common.Address, fee *big.Int, amountOut *big.Int, sqrtPriceLimitX96 *big.Int) (*types.Transaction, error) {
	return _Iquoter.contract.Transact(opts, "quoteExactOutputSingle", tokenIn, tokenOut, fee, amountOut, sqrtPriceLimitX96)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0x30d07f21.
//
// Solidity: function quoteExactOutputSingle(address tokenIn, address tokenOut, uint24 fee, uint256 amountOut, uint160 sqrtPriceLimitX96) returns(uint256 amountIn)
func (_Iquoter *IquoterSession) QuoteExactOutputSingle(tokenIn common.Address, tokenOut common.Address, fee *big.Int, amountOut *big.Int, sqrtPriceLimitX96 *big.Int) (*types.Transaction, error) {
	return _Iquoter.Contract.QuoteExactOutputSingle(&_Iquoter.TransactOpts, tokenIn, tokenOut, fee, amountOut, sqrtPriceLimitX96)
}

// QuoteExactOutputSingle is a paid mutator transaction binding the contract method 0x30d07f21.
//
// Solidity: function quoteExactOutputSingle(address tokenIn, address tokenOut, uint24 fee, uint256 amountOut, uint160 sqrtPriceLimitX96) returns(uint256 amountIn)
func (_Iquoter *IquoterTransactorSession) QuoteExactOutputSingle(tokenIn common.Address, tokenOut common.Address, fee *big.Int, amountOut *big.Int, sqrtPriceLimitX96 *big.Int) (*types.Transaction, error) {
	return _Iquoter.Contract.QuoteExactOutputSingle(&_Iquoter.TransactOpts, tokenIn, tokenOut, fee, amountOut, sqrtPriceLimitX96)
}
