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

// SimpleAirdropABI is the input ABI used to generate the binding from.
const SimpleAirdropABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_maxAirdropAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"date\",\"type\":\"uint256\"}],\"name\":\"AirdropProcessed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"claimTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentAirdropAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxAirdropAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"processedAirdrops\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"updateAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SimpleAirdrop is an auto generated Go binding around an Ethereum contract.
type SimpleAirdrop struct {
	SimpleAirdropCaller     // Read-only binding to the contract
	SimpleAirdropTransactor // Write-only binding to the contract
	SimpleAirdropFilterer   // Log filterer for contract events
}

// SimpleAirdropCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleAirdropCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleAirdropTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleAirdropTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleAirdropFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleAirdropFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleAirdropSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleAirdropSession struct {
	Contract     *SimpleAirdrop    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleAirdropCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleAirdropCallerSession struct {
	Contract *SimpleAirdropCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SimpleAirdropTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleAirdropTransactorSession struct {
	Contract     *SimpleAirdropTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SimpleAirdropRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleAirdropRaw struct {
	Contract *SimpleAirdrop // Generic contract binding to access the raw methods on
}

// SimpleAirdropCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleAirdropCallerRaw struct {
	Contract *SimpleAirdropCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleAirdropTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleAirdropTransactorRaw struct {
	Contract *SimpleAirdropTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleAirdrop creates a new instance of SimpleAirdrop, bound to a specific deployed contract.
func NewSimpleAirdrop(address common.Address, backend bind.ContractBackend) (*SimpleAirdrop, error) {
	contract, err := bindSimpleAirdrop(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleAirdrop{SimpleAirdropCaller: SimpleAirdropCaller{contract: contract}, SimpleAirdropTransactor: SimpleAirdropTransactor{contract: contract}, SimpleAirdropFilterer: SimpleAirdropFilterer{contract: contract}}, nil
}

// NewSimpleAirdropCaller creates a new read-only instance of SimpleAirdrop, bound to a specific deployed contract.
func NewSimpleAirdropCaller(address common.Address, caller bind.ContractCaller) (*SimpleAirdropCaller, error) {
	contract, err := bindSimpleAirdrop(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleAirdropCaller{contract: contract}, nil
}

// NewSimpleAirdropTransactor creates a new write-only instance of SimpleAirdrop, bound to a specific deployed contract.
func NewSimpleAirdropTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleAirdropTransactor, error) {
	contract, err := bindSimpleAirdrop(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleAirdropTransactor{contract: contract}, nil
}

// NewSimpleAirdropFilterer creates a new log filterer instance of SimpleAirdrop, bound to a specific deployed contract.
func NewSimpleAirdropFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleAirdropFilterer, error) {
	contract, err := bindSimpleAirdrop(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleAirdropFilterer{contract: contract}, nil
}

// bindSimpleAirdrop binds a generic wrapper to an already deployed contract.
func bindSimpleAirdrop(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimpleAirdropABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleAirdrop *SimpleAirdropRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleAirdrop.Contract.SimpleAirdropCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleAirdrop *SimpleAirdropRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleAirdrop.Contract.SimpleAirdropTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleAirdrop *SimpleAirdropRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleAirdrop.Contract.SimpleAirdropTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleAirdrop *SimpleAirdropCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleAirdrop.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleAirdrop *SimpleAirdropTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleAirdrop.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleAirdrop *SimpleAirdropTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleAirdrop.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_SimpleAirdrop *SimpleAirdropCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SimpleAirdrop.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_SimpleAirdrop *SimpleAirdropSession) Admin() (common.Address, error) {
	return _SimpleAirdrop.Contract.Admin(&_SimpleAirdrop.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_SimpleAirdrop *SimpleAirdropCallerSession) Admin() (common.Address, error) {
	return _SimpleAirdrop.Contract.Admin(&_SimpleAirdrop.CallOpts)
}

// CurrentAirdropAmount is a free data retrieval call binding the contract method 0x7424bab1.
//
// Solidity: function currentAirdropAmount() view returns(uint256)
func (_SimpleAirdrop *SimpleAirdropCaller) CurrentAirdropAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SimpleAirdrop.contract.Call(opts, &out, "currentAirdropAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentAirdropAmount is a free data retrieval call binding the contract method 0x7424bab1.
//
// Solidity: function currentAirdropAmount() view returns(uint256)
func (_SimpleAirdrop *SimpleAirdropSession) CurrentAirdropAmount() (*big.Int, error) {
	return _SimpleAirdrop.Contract.CurrentAirdropAmount(&_SimpleAirdrop.CallOpts)
}

// CurrentAirdropAmount is a free data retrieval call binding the contract method 0x7424bab1.
//
// Solidity: function currentAirdropAmount() view returns(uint256)
func (_SimpleAirdrop *SimpleAirdropCallerSession) CurrentAirdropAmount() (*big.Int, error) {
	return _SimpleAirdrop.Contract.CurrentAirdropAmount(&_SimpleAirdrop.CallOpts)
}

// MaxAirdropAmount is a free data retrieval call binding the contract method 0x23efeb12.
//
// Solidity: function maxAirdropAmount() view returns(uint256)
func (_SimpleAirdrop *SimpleAirdropCaller) MaxAirdropAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SimpleAirdrop.contract.Call(opts, &out, "maxAirdropAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAirdropAmount is a free data retrieval call binding the contract method 0x23efeb12.
//
// Solidity: function maxAirdropAmount() view returns(uint256)
func (_SimpleAirdrop *SimpleAirdropSession) MaxAirdropAmount() (*big.Int, error) {
	return _SimpleAirdrop.Contract.MaxAirdropAmount(&_SimpleAirdrop.CallOpts)
}

// MaxAirdropAmount is a free data retrieval call binding the contract method 0x23efeb12.
//
// Solidity: function maxAirdropAmount() view returns(uint256)
func (_SimpleAirdrop *SimpleAirdropCallerSession) MaxAirdropAmount() (*big.Int, error) {
	return _SimpleAirdrop.Contract.MaxAirdropAmount(&_SimpleAirdrop.CallOpts)
}

// ProcessedAirdrops is a free data retrieval call binding the contract method 0x32ffa783.
//
// Solidity: function processedAirdrops(address ) view returns(bool)
func (_SimpleAirdrop *SimpleAirdropCaller) ProcessedAirdrops(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _SimpleAirdrop.contract.Call(opts, &out, "processedAirdrops", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProcessedAirdrops is a free data retrieval call binding the contract method 0x32ffa783.
//
// Solidity: function processedAirdrops(address ) view returns(bool)
func (_SimpleAirdrop *SimpleAirdropSession) ProcessedAirdrops(arg0 common.Address) (bool, error) {
	return _SimpleAirdrop.Contract.ProcessedAirdrops(&_SimpleAirdrop.CallOpts, arg0)
}

// ProcessedAirdrops is a free data retrieval call binding the contract method 0x32ffa783.
//
// Solidity: function processedAirdrops(address ) view returns(bool)
func (_SimpleAirdrop *SimpleAirdropCallerSession) ProcessedAirdrops(arg0 common.Address) (bool, error) {
	return _SimpleAirdrop.Contract.ProcessedAirdrops(&_SimpleAirdrop.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_SimpleAirdrop *SimpleAirdropCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SimpleAirdrop.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_SimpleAirdrop *SimpleAirdropSession) Token() (common.Address, error) {
	return _SimpleAirdrop.Contract.Token(&_SimpleAirdrop.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_SimpleAirdrop *SimpleAirdropCallerSession) Token() (common.Address, error) {
	return _SimpleAirdrop.Contract.Token(&_SimpleAirdrop.CallOpts)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0x14d8bbf1.
//
// Solidity: function claimTokens(uint256 amount, bytes signature) returns()
func (_SimpleAirdrop *SimpleAirdropTransactor) ClaimTokens(opts *bind.TransactOpts, amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _SimpleAirdrop.contract.Transact(opts, "claimTokens", amount, signature)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0x14d8bbf1.
//
// Solidity: function claimTokens(uint256 amount, bytes signature) returns()
func (_SimpleAirdrop *SimpleAirdropSession) ClaimTokens(amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _SimpleAirdrop.Contract.ClaimTokens(&_SimpleAirdrop.TransactOpts, amount, signature)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0x14d8bbf1.
//
// Solidity: function claimTokens(uint256 amount, bytes signature) returns()
func (_SimpleAirdrop *SimpleAirdropTransactorSession) ClaimTokens(amount *big.Int, signature []byte) (*types.Transaction, error) {
	return _SimpleAirdrop.Contract.ClaimTokens(&_SimpleAirdrop.TransactOpts, amount, signature)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0xe2f273bd.
//
// Solidity: function updateAdmin(address newAdmin) returns()
func (_SimpleAirdrop *SimpleAirdropTransactor) UpdateAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _SimpleAirdrop.contract.Transact(opts, "updateAdmin", newAdmin)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0xe2f273bd.
//
// Solidity: function updateAdmin(address newAdmin) returns()
func (_SimpleAirdrop *SimpleAirdropSession) UpdateAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _SimpleAirdrop.Contract.UpdateAdmin(&_SimpleAirdrop.TransactOpts, newAdmin)
}

// UpdateAdmin is a paid mutator transaction binding the contract method 0xe2f273bd.
//
// Solidity: function updateAdmin(address newAdmin) returns()
func (_SimpleAirdrop *SimpleAirdropTransactorSession) UpdateAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _SimpleAirdrop.Contract.UpdateAdmin(&_SimpleAirdrop.TransactOpts, newAdmin)
}

// SimpleAirdropAirdropProcessedIterator is returned from FilterAirdropProcessed and is used to iterate over the raw logs and unpacked data for AirdropProcessed events raised by the SimpleAirdrop contract.
type SimpleAirdropAirdropProcessedIterator struct {
	Event *SimpleAirdropAirdropProcessed // Event containing the contract specifics and raw log

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
func (it *SimpleAirdropAirdropProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleAirdropAirdropProcessed)
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
		it.Event = new(SimpleAirdropAirdropProcessed)
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
func (it *SimpleAirdropAirdropProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleAirdropAirdropProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleAirdropAirdropProcessed represents a AirdropProcessed event raised by the SimpleAirdrop contract.
type SimpleAirdropAirdropProcessed struct {
	Recipient common.Address
	Amount    *big.Int
	Date      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAirdropProcessed is a free log retrieval operation binding the contract event 0xc0312d3bf25e86c787fa77f4d038213e428dc722e16d38ca9719e987d20f7c7c.
//
// Solidity: event AirdropProcessed(address recipient, uint256 amount, uint256 date)
func (_SimpleAirdrop *SimpleAirdropFilterer) FilterAirdropProcessed(opts *bind.FilterOpts) (*SimpleAirdropAirdropProcessedIterator, error) {

	logs, sub, err := _SimpleAirdrop.contract.FilterLogs(opts, "AirdropProcessed")
	if err != nil {
		return nil, err
	}
	return &SimpleAirdropAirdropProcessedIterator{contract: _SimpleAirdrop.contract, event: "AirdropProcessed", logs: logs, sub: sub}, nil
}

// WatchAirdropProcessed is a free log subscription operation binding the contract event 0xc0312d3bf25e86c787fa77f4d038213e428dc722e16d38ca9719e987d20f7c7c.
//
// Solidity: event AirdropProcessed(address recipient, uint256 amount, uint256 date)
func (_SimpleAirdrop *SimpleAirdropFilterer) WatchAirdropProcessed(opts *bind.WatchOpts, sink chan<- *SimpleAirdropAirdropProcessed) (event.Subscription, error) {

	logs, sub, err := _SimpleAirdrop.contract.WatchLogs(opts, "AirdropProcessed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleAirdropAirdropProcessed)
				if err := _SimpleAirdrop.contract.UnpackLog(event, "AirdropProcessed", log); err != nil {
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
func (_SimpleAirdrop *SimpleAirdropFilterer) ParseAirdropProcessed(log types.Log) (*SimpleAirdropAirdropProcessed, error) {
	event := new(SimpleAirdropAirdropProcessed)
	if err := _SimpleAirdrop.contract.UnpackLog(event, "AirdropProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
