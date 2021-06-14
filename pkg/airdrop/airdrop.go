// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package airdrop

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

// AirdropABI is the input ABI used to generate the binding from.
const AirdropABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxAirdropAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"date\",\"type\":\"uint256\"}],\"name\":\"AirdropProcessed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"claimTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentAirdropAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxAirdropAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"processedAirdrops\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"updateAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Airdrop is an auto generated Go binding around an Ethereum contract.
type Airdrop struct {
	AirdropCaller     // Read-only binding to the contract
	AirdropTransactor // Write-only binding to the contract
	AirdropFilterer   // Log filterer for contract events
}

// AirdropCaller is an auto generated read-only Go binding around an Ethereum contract.
type AirdropCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AirdropTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AirdropTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AirdropFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AirdropFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AirdropSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AirdropSession struct {
	Contract     *Airdrop          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AirdropCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AirdropCallerSession struct {
	Contract *AirdropCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AirdropTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AirdropTransactorSession struct {
	Contract     *AirdropTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AirdropRaw is an auto generated low-level Go binding around an Ethereum contract.
type AirdropRaw struct {
	Contract *Airdrop // Generic contract binding to access the raw methods on
}

// AirdropCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AirdropCallerRaw struct {
	Contract *AirdropCaller // Generic read-only contract binding to access the raw methods on
}

// AirdropTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AirdropTransactorRaw struct {
	Contract *AirdropTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAirdrop creates a new instance of Airdrop, bound to a specific deployed contract.
func NewAirdrop(address common.Address, backend bind.ContractBackend) (*Airdrop, error) {
	contract, err := bindAirdrop(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Airdrop{AirdropCaller: AirdropCaller{contract: contract}, AirdropTransactor: AirdropTransactor{contract: contract}, AirdropFilterer: AirdropFilterer{contract: contract}}, nil
}

// NewAirdropCaller creates a new read-only instance of Airdrop, bound to a specific deployed contract.
func NewAirdropCaller(address common.Address, caller bind.ContractCaller) (*AirdropCaller, error) {
	contract, err := bindAirdrop(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AirdropCaller{contract: contract}, nil
}

// NewAirdropTransactor creates a new write-only instance of Airdrop, bound to a specific deployed contract.
func NewAirdropTransactor(address common.Address, transactor bind.ContractTransactor) (*AirdropTransactor, error) {
	contract, err := bindAirdrop(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AirdropTransactor{contract: contract}, nil
}

// NewAirdropFilterer creates a new log filterer instance of Airdrop, bound to a specific deployed contract.
func NewAirdropFilterer(address common.Address, filterer bind.ContractFilterer) (*AirdropFilterer, error) {
	contract, err := bindAirdrop(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AirdropFilterer{contract: contract}, nil
}

// bindAirdrop binds a generic wrapper to an already deployed contract.
func bindAirdrop(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AirdropABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Airdrop *AirdropRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Airdrop.Contract.AirdropCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Airdrop *AirdropRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Airdrop.Contract.AirdropTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Airdrop *AirdropRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Airdrop.Contract.AirdropTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Airdrop *AirdropCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Airdrop.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Airdrop *AirdropTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Airdrop.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Airdrop *AirdropTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Airdrop.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Airdrop *AirdropCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Airdrop *AirdropSession) Admin() (common.Address, error) {
	return _Airdrop.Contract.Admin(&_Airdrop.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Airdrop *AirdropCallerSession) Admin() (common.Address, error) {
	return _Airdrop.Contract.Admin(&_Airdrop.CallOpts)
}

// CurrentAirdropAmount is a free data retrieval call binding the contract method 0x7424bab1.
//
// Solidity: function currentAirdropAmount() view returns(uint256)
func (_Airdrop *AirdropCaller) CurrentAirdropAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "currentAirdropAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentAirdropAmount is a free data retrieval call binding the contract method 0x7424bab1.
//
// Solidity: function currentAirdropAmount() view returns(uint256)
func (_Airdrop *AirdropSession) CurrentAirdropAmount() (*big.Int, error) {
	return _Airdrop.Contract.CurrentAirdropAmount(&_Airdrop.CallOpts)
}

// CurrentAirdropAmount is a free data retrieval call binding the contract method 0x7424bab1.
//
// Solidity: function currentAirdropAmount() view returns(uint256)
func (_Airdrop *AirdropCallerSession) CurrentAirdropAmount() (*big.Int, error) {
	return _Airdrop.Contract.CurrentAirdropAmount(&_Airdrop.CallOpts)
}

// MaxAirdropAmount is a free data retrieval call binding the contract method 0x23efeb12.
//
// Solidity: function maxAirdropAmount() view returns(uint256)
func (_Airdrop *AirdropCaller) MaxAirdropAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "maxAirdropAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAirdropAmount is a free data retrieval call binding the contract method 0x23efeb12.
//
// Solidity: function maxAirdropAmount() view returns(uint256)
func (_Airdrop *AirdropSession) MaxAirdropAmount() (*big.Int, error) {
	return _Airdrop.Contract.MaxAirdropAmount(&_Airdrop.CallOpts)
}

// MaxAirdropAmount is a free data retrieval call binding the contract method 0x23efeb12.
//
// Solidity: function maxAirdropAmount() view returns(uint256)
func (_Airdrop *AirdropCallerSession) MaxAirdropAmount() (*big.Int, error) {
	return _Airdrop.Contract.MaxAirdropAmount(&_Airdrop.CallOpts)
}

// ProcessedAirdrops is a free data retrieval call binding the contract method 0x32ffa783.
//
// Solidity: function processedAirdrops(address ) view returns(bool)
func (_Airdrop *AirdropCaller) ProcessedAirdrops(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "processedAirdrops", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProcessedAirdrops is a free data retrieval call binding the contract method 0x32ffa783.
//
// Solidity: function processedAirdrops(address ) view returns(bool)
func (_Airdrop *AirdropSession) ProcessedAirdrops(arg0 common.Address) (bool, error) {
	return _Airdrop.Contract.ProcessedAirdrops(&_Airdrop.CallOpts, arg0)
}

// ProcessedAirdrops is a free data retrieval call binding the contract method 0x32ffa783.
//
// Solidity: function processedAirdrops(address ) view returns(bool)
func (_Airdrop *AirdropCallerSession) ProcessedAirdrops(arg0 common.Address) (bool, error) {
	return _Airdrop.Contract.ProcessedAirdrops(&_Airdrop.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Airdrop *AirdropCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Airdrop *AirdropSession) Token() (common.Address, error) {
	return _Airdrop.Contract.Token(&_Airdrop.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Airdrop *AirdropCallerSession) Token() (common.Address, error) {
	return _Airdrop.Contract.Token(&_Airdrop.CallOpts)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0x14d8bbf1.
//
// Solidity: function claimTokens(uint256 amount, bytes signature) returns()
func (_Airdrop *AirdropTransactor) ClaimTokens(opts *bind.TransactOpts, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "claimTokens", amount, signature)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0x14d8bbf1.
//
// Solidity: function claimTokens(uint256 amount, bytes signature) returns()
func (_Airdrop *AirdropSession) ClaimTokens(amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _Airdrop.Contract.ClaimTokens(&_Airdrop.TransactOpts, amount, signature)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0x14d8bbf1.
//
// Solidity: function claimTokens(uint256 amount, bytes signature) returns()
func (_Airdrop *AirdropTransactorSession) ClaimTokens(amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _Airdrop.Contract.ClaimTokens(&_Airdrop.TransactOpts, amount, signature)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0xe2f273bd.
//
// Solidity: function updateAdmin(address newAdmin) returns()
func (_Airdrop *AirdropTransactor) UpdateAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "updateAdmin", newAdmin)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0xe2f273bd.
//
// Solidity: function updateAdmin(address newAdmin) returns()
func (_Airdrop *AirdropSession) UpdateAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.UpdateAdmin(&_Airdrop.TransactOpts, newAdmin)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0xe2f273bd.
//
// Solidity: function updateAdmin(address newAdmin) returns()
func (_Airdrop *AirdropTransactorSession) UpdateAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Airdrop.Contract.UpdateAdmin(&_Airdrop.TransactOpts, newAdmin)
}

// AirdropAirdropProcessedIterator is returned from FilterAirdropProcessed and is used to iterate over the raw logs and unpacked data for AirdropProcessed events raised by the Airdrop contract.
type AirdropAirdropProcessedIterator struct {
	Event *AirdropAirdropProcessed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AirdropAirdropProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AirdropAirdropProcessed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AirdropAirdropProcessed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AirdropAirdropProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AirdropAirdropProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AirdropAirdropProcessed represents a AirdropProcessed event raised by the Airdrop contract.
type AirdropAirdropProcessed struct {
	Recipient common.Address
	Amount    *big.Int
	Date      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAirdropProcessed is a free log retrieval operation binding the contract event 0xc0312d3bf25e86c787fa77f4d038213e428dc722e16d38ca9719e987d20f7c7c.
//
// Solidity: event AirdropProcessed(address recipient, uint256 amount, uint256 date)
func (_Airdrop *AirdropFilterer) FilterAirdropProcessed(opts *bind.FilterOpts) (*AirdropAirdropProcessedIterator, error) {

	logs, sub, err := _Airdrop.contract.FilterLogs(opts, "AirdropProcessed")
	if err != nil {
		return nil, err
	}
	return &AirdropAirdropProcessedIterator{contract: _Airdrop.contract, event: "AirdropProcessed", logs: logs, sub: sub}, nil
}

// WatchAirdropProcessed is a free log subscription operation binding the contract event 0xc0312d3bf25e86c787fa77f4d038213e428dc722e16d38ca9719e987d20f7c7c.
//
// Solidity: event AirdropProcessed(address recipient, uint256 amount, uint256 date)
func (_Airdrop *AirdropFilterer) WatchAirdropProcessed(opts *bind.WatchOpts, sink chan<- *AirdropAirdropProcessed) (event.Subscription, error) {

	logs, sub, err := _Airdrop.contract.WatchLogs(opts, "AirdropProcessed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AirdropAirdropProcessed)
				if err := _Airdrop.contract.UnpackLog(event, "AirdropProcessed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAirdropProcessed is a log parse operation binding the contract event 0xc0312d3bf25e86c787fa77f4d038213e428dc722e16d38ca9719e987d20f7c7c.
//
// Solidity: event AirdropProcessed(address recipient, uint256 amount, uint256 date)
func (_Airdrop *AirdropFilterer) ParseAirdropProcessed(log types.Log) (*AirdropAirdropProcessed, error) {
	event := new(AirdropAirdropProcessed)
	if err := _Airdrop.contract.UnpackLog(event, "AirdropProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC20Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimedABI is the input ABI used to generate the binding from.
const TimedABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_REG_TIME\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_CLAIM_TIME\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CLAIM_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REG_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"status\",\"outputs\":[{\"internalType\":\"enumTimed.Status\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Timed is an auto generated Go binding around an Ethereum contract.
type Timed struct {
	TimedCaller     // Read-only binding to the contract
	TimedTransactor // Write-only binding to the contract
	TimedFilterer   // Log filterer for contract events
}

// TimedCaller is an auto generated read-only Go binding around an Ethereum contract.
type TimedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TimedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TimedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TimedSession struct {
	Contract     *Timed            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TimedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TimedCallerSession struct {
	Contract *TimedCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TimedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TimedTransactorSession struct {
	Contract     *TimedTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TimedRaw is an auto generated low-level Go binding around an Ethereum contract.
type TimedRaw struct {
	Contract *Timed // Generic contract binding to access the raw methods on
}

// TimedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TimedCallerRaw struct {
	Contract *TimedCaller // Generic read-only contract binding to access the raw methods on
}

// TimedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TimedTransactorRaw struct {
	Contract *TimedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTimed creates a new instance of Timed, bound to a specific deployed contract.
func NewTimed(address common.Address, backend bind.ContractBackend) (*Timed, error) {
	contract, err := bindTimed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Timed{TimedCaller: TimedCaller{contract: contract}, TimedTransactor: TimedTransactor{contract: contract}, TimedFilterer: TimedFilterer{contract: contract}}, nil
}

// NewTimedCaller creates a new read-only instance of Timed, bound to a specific deployed contract.
func NewTimedCaller(address common.Address, caller bind.ContractCaller) (*TimedCaller, error) {
	contract, err := bindTimed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TimedCaller{contract: contract}, nil
}

// NewTimedTransactor creates a new write-only instance of Timed, bound to a specific deployed contract.
func NewTimedTransactor(address common.Address, transactor bind.ContractTransactor) (*TimedTransactor, error) {
	contract, err := bindTimed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TimedTransactor{contract: contract}, nil
}

// NewTimedFilterer creates a new log filterer instance of Timed, bound to a specific deployed contract.
func NewTimedFilterer(address common.Address, filterer bind.ContractFilterer) (*TimedFilterer, error) {
	contract, err := bindTimed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TimedFilterer{contract: contract}, nil
}

// bindTimed binds a generic wrapper to an already deployed contract.
func bindTimed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TimedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Timed *TimedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Timed.Contract.TimedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Timed *TimedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Timed.Contract.TimedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Timed *TimedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Timed.Contract.TimedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Timed *TimedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Timed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Timed *TimedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Timed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Timed *TimedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Timed.Contract.contract.Transact(opts, method, params...)
}

// CLAIMTIME is a free data retrieval call binding the contract method 0x3bcaba8b.
//
// Solidity: function CLAIM_TIME() view returns(uint256)
func (_Timed *TimedCaller) CLAIMTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Timed.contract.Call(opts, &out, "CLAIM_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CLAIMTIME is a free data retrieval call binding the contract method 0x3bcaba8b.
//
// Solidity: function CLAIM_TIME() view returns(uint256)
func (_Timed *TimedSession) CLAIMTIME() (*big.Int, error) {
	return _Timed.Contract.CLAIMTIME(&_Timed.CallOpts)
}

// CLAIMTIME is a free data retrieval call binding the contract method 0x3bcaba8b.
//
// Solidity: function CLAIM_TIME() view returns(uint256)
func (_Timed *TimedCallerSession) CLAIMTIME() (*big.Int, error) {
	return _Timed.Contract.CLAIMTIME(&_Timed.CallOpts)
}

// INITTIME is a free data retrieval call binding the contract method 0xfb841c6d.
//
// Solidity: function INIT_TIME() view returns(uint256)
func (_Timed *TimedCaller) INITTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Timed.contract.Call(opts, &out, "INIT_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITTIME is a free data retrieval call binding the contract method 0xfb841c6d.
//
// Solidity: function INIT_TIME() view returns(uint256)
func (_Timed *TimedSession) INITTIME() (*big.Int, error) {
	return _Timed.Contract.INITTIME(&_Timed.CallOpts)
}

// INITTIME is a free data retrieval call binding the contract method 0xfb841c6d.
//
// Solidity: function INIT_TIME() view returns(uint256)
func (_Timed *TimedCallerSession) INITTIME() (*big.Int, error) {
	return _Timed.Contract.INITTIME(&_Timed.CallOpts)
}

// REGTIME is a free data retrieval call binding the contract method 0x47f7b722.
//
// Solidity: function REG_TIME() view returns(uint256)
func (_Timed *TimedCaller) REGTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Timed.contract.Call(opts, &out, "REG_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REGTIME is a free data retrieval call binding the contract method 0x47f7b722.
//
// Solidity: function REG_TIME() view returns(uint256)
func (_Timed *TimedSession) REGTIME() (*big.Int, error) {
	return _Timed.Contract.REGTIME(&_Timed.CallOpts)
}

// REGTIME is a free data retrieval call binding the contract method 0x47f7b722.
//
// Solidity: function REG_TIME() view returns(uint256)
func (_Timed *TimedCallerSession) REGTIME() (*big.Int, error) {
	return _Timed.Contract.REGTIME(&_Timed.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_Timed *TimedCaller) Status(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Timed.contract.Call(opts, &out, "status")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_Timed *TimedSession) Status() (uint8, error) {
	return _Timed.Contract.Status(&_Timed.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_Timed *TimedCallerSession) Status() (uint8, error) {
	return _Timed.Contract.Status(&_Timed.CallOpts)
}

// TimedAirdropABI is the input ABI used to generate the binding from.
const TimedAirdropABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxAirdropAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"date\",\"type\":\"uint256\"}],\"name\":\"AirdropProcessed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CLAIM_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INIT_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REG_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"claimTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentAirdropAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxAirdropAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"processedAirdrops\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"status\",\"outputs\":[{\"internalType\":\"enumTimed.Status\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"updateAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TimedAirdrop is an auto generated Go binding around an Ethereum contract.
type TimedAirdrop struct {
	TimedAirdropCaller     // Read-only binding to the contract
	TimedAirdropTransactor // Write-only binding to the contract
	TimedAirdropFilterer   // Log filterer for contract events
}

// TimedAirdropCaller is an auto generated read-only Go binding around an Ethereum contract.
type TimedAirdropCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimedAirdropTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TimedAirdropTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimedAirdropFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TimedAirdropFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimedAirdropSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TimedAirdropSession struct {
	Contract     *TimedAirdrop     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TimedAirdropCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TimedAirdropCallerSession struct {
	Contract *TimedAirdropCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TimedAirdropTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TimedAirdropTransactorSession struct {
	Contract     *TimedAirdropTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TimedAirdropRaw is an auto generated low-level Go binding around an Ethereum contract.
type TimedAirdropRaw struct {
	Contract *TimedAirdrop // Generic contract binding to access the raw methods on
}

// TimedAirdropCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TimedAirdropCallerRaw struct {
	Contract *TimedAirdropCaller // Generic read-only contract binding to access the raw methods on
}

// TimedAirdropTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TimedAirdropTransactorRaw struct {
	Contract *TimedAirdropTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTimedAirdrop creates a new instance of TimedAirdrop, bound to a specific deployed contract.
func NewTimedAirdrop(address common.Address, backend bind.ContractBackend) (*TimedAirdrop, error) {
	contract, err := bindTimedAirdrop(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TimedAirdrop{TimedAirdropCaller: TimedAirdropCaller{contract: contract}, TimedAirdropTransactor: TimedAirdropTransactor{contract: contract}, TimedAirdropFilterer: TimedAirdropFilterer{contract: contract}}, nil
}

// NewTimedAirdropCaller creates a new read-only instance of TimedAirdrop, bound to a specific deployed contract.
func NewTimedAirdropCaller(address common.Address, caller bind.ContractCaller) (*TimedAirdropCaller, error) {
	contract, err := bindTimedAirdrop(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TimedAirdropCaller{contract: contract}, nil
}

// NewTimedAirdropTransactor creates a new write-only instance of TimedAirdrop, bound to a specific deployed contract.
func NewTimedAirdropTransactor(address common.Address, transactor bind.ContractTransactor) (*TimedAirdropTransactor, error) {
	contract, err := bindTimedAirdrop(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TimedAirdropTransactor{contract: contract}, nil
}

// NewTimedAirdropFilterer creates a new log filterer instance of TimedAirdrop, bound to a specific deployed contract.
func NewTimedAirdropFilterer(address common.Address, filterer bind.ContractFilterer) (*TimedAirdropFilterer, error) {
	contract, err := bindTimedAirdrop(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TimedAirdropFilterer{contract: contract}, nil
}

// bindTimedAirdrop binds a generic wrapper to an already deployed contract.
func bindTimedAirdrop(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TimedAirdropABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TimedAirdrop *TimedAirdropRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TimedAirdrop.Contract.TimedAirdropCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TimedAirdrop *TimedAirdropRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimedAirdrop.Contract.TimedAirdropTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TimedAirdrop *TimedAirdropRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TimedAirdrop.Contract.TimedAirdropTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TimedAirdrop *TimedAirdropCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TimedAirdrop.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TimedAirdrop *TimedAirdropTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimedAirdrop.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TimedAirdrop *TimedAirdropTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TimedAirdrop.Contract.contract.Transact(opts, method, params...)
}

// CLAIMTIME is a free data retrieval call binding the contract method 0x3bcaba8b.
//
// Solidity: function CLAIM_TIME() view returns(uint256)
func (_TimedAirdrop *TimedAirdropCaller) CLAIMTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TimedAirdrop.contract.Call(opts, &out, "CLAIM_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CLAIMTIME is a free data retrieval call binding the contract method 0x3bcaba8b.
//
// Solidity: function CLAIM_TIME() view returns(uint256)
func (_TimedAirdrop *TimedAirdropSession) CLAIMTIME() (*big.Int, error) {
	return _TimedAirdrop.Contract.CLAIMTIME(&_TimedAirdrop.CallOpts)
}

// CLAIMTIME is a free data retrieval call binding the contract method 0x3bcaba8b.
//
// Solidity: function CLAIM_TIME() view returns(uint256)
func (_TimedAirdrop *TimedAirdropCallerSession) CLAIMTIME() (*big.Int, error) {
	return _TimedAirdrop.Contract.CLAIMTIME(&_TimedAirdrop.CallOpts)
}

// INITTIME is a free data retrieval call binding the contract method 0xfb841c6d.
//
// Solidity: function INIT_TIME() view returns(uint256)
func (_TimedAirdrop *TimedAirdropCaller) INITTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TimedAirdrop.contract.Call(opts, &out, "INIT_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITTIME is a free data retrieval call binding the contract method 0xfb841c6d.
//
// Solidity: function INIT_TIME() view returns(uint256)
func (_TimedAirdrop *TimedAirdropSession) INITTIME() (*big.Int, error) {
	return _TimedAirdrop.Contract.INITTIME(&_TimedAirdrop.CallOpts)
}

// INITTIME is a free data retrieval call binding the contract method 0xfb841c6d.
//
// Solidity: function INIT_TIME() view returns(uint256)
func (_TimedAirdrop *TimedAirdropCallerSession) INITTIME() (*big.Int, error) {
	return _TimedAirdrop.Contract.INITTIME(&_TimedAirdrop.CallOpts)
}

// REGTIME is a free data retrieval call binding the contract method 0x47f7b722.
//
// Solidity: function REG_TIME() view returns(uint256)
func (_TimedAirdrop *TimedAirdropCaller) REGTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TimedAirdrop.contract.Call(opts, &out, "REG_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REGTIME is a free data retrieval call binding the contract method 0x47f7b722.
//
// Solidity: function REG_TIME() view returns(uint256)
func (_TimedAirdrop *TimedAirdropSession) REGTIME() (*big.Int, error) {
	return _TimedAirdrop.Contract.REGTIME(&_TimedAirdrop.CallOpts)
}

// REGTIME is a free data retrieval call binding the contract method 0x47f7b722.
//
// Solidity: function REG_TIME() view returns(uint256)
func (_TimedAirdrop *TimedAirdropCallerSession) REGTIME() (*big.Int, error) {
	return _TimedAirdrop.Contract.REGTIME(&_TimedAirdrop.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_TimedAirdrop *TimedAirdropCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TimedAirdrop.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_TimedAirdrop *TimedAirdropSession) Admin() (common.Address, error) {
	return _TimedAirdrop.Contract.Admin(&_TimedAirdrop.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_TimedAirdrop *TimedAirdropCallerSession) Admin() (common.Address, error) {
	return _TimedAirdrop.Contract.Admin(&_TimedAirdrop.CallOpts)
}

// CurrentAirdropAmount is a free data retrieval call binding the contract method 0x7424bab1.
//
// Solidity: function currentAirdropAmount() view returns(uint256)
func (_TimedAirdrop *TimedAirdropCaller) CurrentAirdropAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TimedAirdrop.contract.Call(opts, &out, "currentAirdropAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentAirdropAmount is a free data retrieval call binding the contract method 0x7424bab1.
//
// Solidity: function currentAirdropAmount() view returns(uint256)
func (_TimedAirdrop *TimedAirdropSession) CurrentAirdropAmount() (*big.Int, error) {
	return _TimedAirdrop.Contract.CurrentAirdropAmount(&_TimedAirdrop.CallOpts)
}

// CurrentAirdropAmount is a free data retrieval call binding the contract method 0x7424bab1.
//
// Solidity: function currentAirdropAmount() view returns(uint256)
func (_TimedAirdrop *TimedAirdropCallerSession) CurrentAirdropAmount() (*big.Int, error) {
	return _TimedAirdrop.Contract.CurrentAirdropAmount(&_TimedAirdrop.CallOpts)
}

// MaxAirdropAmount is a free data retrieval call binding the contract method 0x23efeb12.
//
// Solidity: function maxAirdropAmount() view returns(uint256)
func (_TimedAirdrop *TimedAirdropCaller) MaxAirdropAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TimedAirdrop.contract.Call(opts, &out, "maxAirdropAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAirdropAmount is a free data retrieval call binding the contract method 0x23efeb12.
//
// Solidity: function maxAirdropAmount() view returns(uint256)
func (_TimedAirdrop *TimedAirdropSession) MaxAirdropAmount() (*big.Int, error) {
	return _TimedAirdrop.Contract.MaxAirdropAmount(&_TimedAirdrop.CallOpts)
}

// MaxAirdropAmount is a free data retrieval call binding the contract method 0x23efeb12.
//
// Solidity: function maxAirdropAmount() view returns(uint256)
func (_TimedAirdrop *TimedAirdropCallerSession) MaxAirdropAmount() (*big.Int, error) {
	return _TimedAirdrop.Contract.MaxAirdropAmount(&_TimedAirdrop.CallOpts)
}

// ProcessedAirdrops is a free data retrieval call binding the contract method 0x32ffa783.
//
// Solidity: function processedAirdrops(address ) view returns(bool)
func (_TimedAirdrop *TimedAirdropCaller) ProcessedAirdrops(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _TimedAirdrop.contract.Call(opts, &out, "processedAirdrops", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProcessedAirdrops is a free data retrieval call binding the contract method 0x32ffa783.
//
// Solidity: function processedAirdrops(address ) view returns(bool)
func (_TimedAirdrop *TimedAirdropSession) ProcessedAirdrops(arg0 common.Address) (bool, error) {
	return _TimedAirdrop.Contract.ProcessedAirdrops(&_TimedAirdrop.CallOpts, arg0)
}

// ProcessedAirdrops is a free data retrieval call binding the contract method 0x32ffa783.
//
// Solidity: function processedAirdrops(address ) view returns(bool)
func (_TimedAirdrop *TimedAirdropCallerSession) ProcessedAirdrops(arg0 common.Address) (bool, error) {
	return _TimedAirdrop.Contract.ProcessedAirdrops(&_TimedAirdrop.CallOpts, arg0)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_TimedAirdrop *TimedAirdropCaller) Status(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TimedAirdrop.contract.Call(opts, &out, "status")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_TimedAirdrop *TimedAirdropSession) Status() (uint8, error) {
	return _TimedAirdrop.Contract.Status(&_TimedAirdrop.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x200d2ed2.
//
// Solidity: function status() view returns(uint8)
func (_TimedAirdrop *TimedAirdropCallerSession) Status() (uint8, error) {
	return _TimedAirdrop.Contract.Status(&_TimedAirdrop.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TimedAirdrop *TimedAirdropCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TimedAirdrop.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TimedAirdrop *TimedAirdropSession) Token() (common.Address, error) {
	return _TimedAirdrop.Contract.Token(&_TimedAirdrop.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TimedAirdrop *TimedAirdropCallerSession) Token() (common.Address, error) {
	return _TimedAirdrop.Contract.Token(&_TimedAirdrop.CallOpts)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0x14d8bbf1.
//
// Solidity: function claimTokens(uint256 amount, bytes signature) returns()
func (_TimedAirdrop *TimedAirdropTransactor) ClaimTokens(opts *bind.TransactOpts, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _TimedAirdrop.contract.Transact(opts, "claimTokens", amount, signature)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0x14d8bbf1.
//
// Solidity: function claimTokens(uint256 amount, bytes signature) returns()
func (_TimedAirdrop *TimedAirdropSession) ClaimTokens(amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _TimedAirdrop.Contract.ClaimTokens(&_TimedAirdrop.TransactOpts, amount, signature)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0x14d8bbf1.
//
// Solidity: function claimTokens(uint256 amount, bytes signature) returns()
func (_TimedAirdrop *TimedAirdropTransactorSession) ClaimTokens(amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _TimedAirdrop.Contract.ClaimTokens(&_TimedAirdrop.TransactOpts, amount, signature)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0xe2f273bd.
//
// Solidity: function updateAdmin(address newAdmin) returns()
func (_TimedAirdrop *TimedAirdropTransactor) UpdateAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _TimedAirdrop.contract.Transact(opts, "updateAdmin", newAdmin)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0xe2f273bd.
//
// Solidity: function updateAdmin(address newAdmin) returns()
func (_TimedAirdrop *TimedAirdropSession) UpdateAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _TimedAirdrop.Contract.UpdateAdmin(&_TimedAirdrop.TransactOpts, newAdmin)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0xe2f273bd.
//
// Solidity: function updateAdmin(address newAdmin) returns()
func (_TimedAirdrop *TimedAirdropTransactorSession) UpdateAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _TimedAirdrop.Contract.UpdateAdmin(&_TimedAirdrop.TransactOpts, newAdmin)
}

// TimedAirdropAirdropProcessedIterator is returned from FilterAirdropProcessed and is used to iterate over the raw logs and unpacked data for AirdropProcessed events raised by the TimedAirdrop contract.
type TimedAirdropAirdropProcessedIterator struct {
	Event *TimedAirdropAirdropProcessed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TimedAirdropAirdropProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimedAirdropAirdropProcessed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TimedAirdropAirdropProcessed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TimedAirdropAirdropProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimedAirdropAirdropProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimedAirdropAirdropProcessed represents a AirdropProcessed event raised by the TimedAirdrop contract.
type TimedAirdropAirdropProcessed struct {
	Recipient common.Address
	Amount    *big.Int
	Date      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAirdropProcessed is a free log retrieval operation binding the contract event 0xc0312d3bf25e86c787fa77f4d038213e428dc722e16d38ca9719e987d20f7c7c.
//
// Solidity: event AirdropProcessed(address recipient, uint256 amount, uint256 date)
func (_TimedAirdrop *TimedAirdropFilterer) FilterAirdropProcessed(opts *bind.FilterOpts) (*TimedAirdropAirdropProcessedIterator, error) {

	logs, sub, err := _TimedAirdrop.contract.FilterLogs(opts, "AirdropProcessed")
	if err != nil {
		return nil, err
	}
	return &TimedAirdropAirdropProcessedIterator{contract: _TimedAirdrop.contract, event: "AirdropProcessed", logs: logs, sub: sub}, nil
}

// WatchAirdropProcessed is a free log subscription operation binding the contract event 0xc0312d3bf25e86c787fa77f4d038213e428dc722e16d38ca9719e987d20f7c7c.
//
// Solidity: event AirdropProcessed(address recipient, uint256 amount, uint256 date)
func (_TimedAirdrop *TimedAirdropFilterer) WatchAirdropProcessed(opts *bind.WatchOpts, sink chan<- *TimedAirdropAirdropProcessed) (event.Subscription, error) {

	logs, sub, err := _TimedAirdrop.contract.WatchLogs(opts, "AirdropProcessed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimedAirdropAirdropProcessed)
				if err := _TimedAirdrop.contract.UnpackLog(event, "AirdropProcessed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAirdropProcessed is a log parse operation binding the contract event 0xc0312d3bf25e86c787fa77f4d038213e428dc722e16d38ca9719e987d20f7c7c.
//
// Solidity: event AirdropProcessed(address recipient, uint256 amount, uint256 date)
func (_TimedAirdrop *TimedAirdropFilterer) ParseAirdropProcessed(log types.Log) (*TimedAirdropAirdropProcessed, error) {
	event := new(TimedAirdropAirdropProcessed)
	if err := _TimedAirdrop.contract.UnpackLog(event, "AirdropProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
