// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dpos_interface

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

// DposInterfaceDelegationData is an auto generated low-level Go binding around an user-defined struct.
type DposInterfaceDelegationData struct {
	Account    common.Address
	Delegation DposInterfaceDelegatorInfo
}

// DposInterfaceDelegatorInfo is an auto generated low-level Go binding around an user-defined struct.
type DposInterfaceDelegatorInfo struct {
	Stake   *big.Int
	Rewards *big.Int
}

// DposInterfaceUndelegationData is an auto generated low-level Go binding around an user-defined struct.
type DposInterfaceUndelegationData struct {
	Stake     *big.Int
	Block     uint64
	Validator common.Address
}

// DposInterfaceValidatorBasicInfo is an auto generated low-level Go binding around an user-defined struct.
type DposInterfaceValidatorBasicInfo struct {
	TotalStake           *big.Int
	CommissionReward     *big.Int
	Commission           uint16
	LastCommissionChange uint64
	Owner                common.Address
	Description          string
	Endpoint             string
}

// DposInterfaceValidatorData is an auto generated low-level Go binding around an user-defined struct.
type DposInterfaceValidatorData struct {
	Account common.Address
	Info    DposInterfaceValidatorBasicInfo
}

// TaraxaDposInterfaceMetaData contains all meta data concerning the TaraxaDposInterface contract.
var TaraxaDposInterfaceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"CommissionRewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"comission\",\"type\":\"uint16\"}],\"name\":\"CommissionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Delegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Redelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"RewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UndelegateCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UndelegateConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Undelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorInfoSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"cancelUndelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"claimCommissionRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"confirmUndelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"batch\",\"type\":\"uint32\"}],\"name\":\"getDelegations\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"internalType\":\"structDposInterface.DelegatorInfo\",\"name\":\"delegation\",\"type\":\"tuple\"}],\"internalType\":\"structDposInterface.DelegationData[]\",\"name\":\"delegations\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"end\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalEligibleVotesCount\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"batch\",\"type\":\"uint32\"}],\"name\":\"getUndelegations\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"block\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"internalType\":\"structDposInterface.UndelegationData[]\",\"name\":\"undelegations\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"end\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"getValidator\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"total_stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commission_reward\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"commission\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"last_commission_change\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"internalType\":\"structDposInterface.ValidatorBasicInfo\",\"name\":\"validator_info\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"getValidatorDelegators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"delegators\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"getValidatorEligibleVotesCount\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"batch\",\"type\":\"uint32\"}],\"name\":\"getValidators\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"total_stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commission_reward\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"commission\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"last_commission_change\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"internalType\":\"structDposInterface.ValidatorBasicInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"internalType\":\"structDposInterface.ValidatorData[]\",\"name\":\"validators\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"end\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"batch\",\"type\":\"uint32\"}],\"name\":\"getValidatorsFor\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"total_stake\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commission_reward\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"commission\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"last_commission_change\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"internalType\":\"structDposInterface.ValidatorBasicInfo\",\"name\":\"info\",\"type\":\"tuple\"}],\"internalType\":\"structDposInterface.ValidatorData[]\",\"name\":\"validators\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"end\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"isValidatorEligible\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validator_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"reDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"vrf_key\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"commission\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"name\":\"registerValidator\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"commission\",\"type\":\"uint16\"}],\"name\":\"setCommission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"name\":\"setValidatorInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"undelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TaraxaDposInterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use TaraxaDposInterfaceMetaData.ABI instead.
var TaraxaDposInterfaceABI = TaraxaDposInterfaceMetaData.ABI

// TaraxaDposInterface is an auto generated Go binding around an Ethereum contract.
type TaraxaDposInterface struct {
	TaraxaDposInterfaceCaller     // Read-only binding to the contract
	TaraxaDposInterfaceTransactor // Write-only binding to the contract
	TaraxaDposInterfaceFilterer   // Log filterer for contract events
}

// TaraxaDposInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaraxaDposInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaraxaDposInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaraxaDposInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaraxaDposInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaraxaDposInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaraxaDposInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaraxaDposInterfaceSession struct {
	Contract     *TaraxaDposInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TaraxaDposInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaraxaDposInterfaceCallerSession struct {
	Contract *TaraxaDposInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// TaraxaDposInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaraxaDposInterfaceTransactorSession struct {
	Contract     *TaraxaDposInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// TaraxaDposInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaraxaDposInterfaceRaw struct {
	Contract *TaraxaDposInterface // Generic contract binding to access the raw methods on
}

// TaraxaDposInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaraxaDposInterfaceCallerRaw struct {
	Contract *TaraxaDposInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// TaraxaDposInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaraxaDposInterfaceTransactorRaw struct {
	Contract *TaraxaDposInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTaraxaDposInterface creates a new instance of TaraxaDposInterface, bound to a specific deployed contract.
func NewTaraxaDposInterface(address common.Address, backend bind.ContractBackend) (*TaraxaDposInterface, error) {
	contract, err := bindTaraxaDposInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TaraxaDposInterface{TaraxaDposInterfaceCaller: TaraxaDposInterfaceCaller{contract: contract}, TaraxaDposInterfaceTransactor: TaraxaDposInterfaceTransactor{contract: contract}, TaraxaDposInterfaceFilterer: TaraxaDposInterfaceFilterer{contract: contract}}, nil
}

// NewTaraxaDposInterfaceCaller creates a new read-only instance of TaraxaDposInterface, bound to a specific deployed contract.
func NewTaraxaDposInterfaceCaller(address common.Address, caller bind.ContractCaller) (*TaraxaDposInterfaceCaller, error) {
	contract, err := bindTaraxaDposInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaraxaDposInterfaceCaller{contract: contract}, nil
}

// NewTaraxaDposInterfaceTransactor creates a new write-only instance of TaraxaDposInterface, bound to a specific deployed contract.
func NewTaraxaDposInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*TaraxaDposInterfaceTransactor, error) {
	contract, err := bindTaraxaDposInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaraxaDposInterfaceTransactor{contract: contract}, nil
}

// NewTaraxaDposInterfaceFilterer creates a new log filterer instance of TaraxaDposInterface, bound to a specific deployed contract.
func NewTaraxaDposInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*TaraxaDposInterfaceFilterer, error) {
	contract, err := bindTaraxaDposInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaraxaDposInterfaceFilterer{contract: contract}, nil
}

// bindTaraxaDposInterface binds a generic wrapper to an already deployed contract.
func bindTaraxaDposInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TaraxaDposInterfaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaraxaDposInterface *TaraxaDposInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaraxaDposInterface.Contract.TaraxaDposInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaraxaDposInterface *TaraxaDposInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaraxaDposInterface.Contract.TaraxaDposInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaraxaDposInterface *TaraxaDposInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaraxaDposInterface.Contract.TaraxaDposInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaraxaDposInterface *TaraxaDposInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaraxaDposInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaraxaDposInterface *TaraxaDposInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaraxaDposInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaraxaDposInterface *TaraxaDposInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaraxaDposInterface.Contract.contract.Transact(opts, method, params...)
}

// GetDelegations is a free data retrieval call binding the contract method 0x8b49d394.
//
// Solidity: function getDelegations(address delegator, uint32 batch) view returns((address,(uint256,uint256))[] delegations, bool end)
func (_TaraxaDposInterface *TaraxaDposInterfaceCaller) GetDelegations(opts *bind.CallOpts, delegator common.Address, batch uint32) (struct {
	Delegations []DposInterfaceDelegationData
	End         bool
}, error) {
	var out []interface{}
	err := _TaraxaDposInterface.contract.Call(opts, &out, "getDelegations", delegator, batch)

	outstruct := new(struct {
		Delegations []DposInterfaceDelegationData
		End         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Delegations = *abi.ConvertType(out[0], new([]DposInterfaceDelegationData)).(*[]DposInterfaceDelegationData)
	outstruct.End = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetDelegations is a free data retrieval call binding the contract method 0x8b49d394.
//
// Solidity: function getDelegations(address delegator, uint32 batch) view returns((address,(uint256,uint256))[] delegations, bool end)
func (_TaraxaDposInterface *TaraxaDposInterfaceSession) GetDelegations(delegator common.Address, batch uint32) (struct {
	Delegations []DposInterfaceDelegationData
	End         bool
}, error) {
	return _TaraxaDposInterface.Contract.GetDelegations(&_TaraxaDposInterface.CallOpts, delegator, batch)
}

// GetDelegations is a free data retrieval call binding the contract method 0x8b49d394.
//
// Solidity: function getDelegations(address delegator, uint32 batch) view returns((address,(uint256,uint256))[] delegations, bool end)
func (_TaraxaDposInterface *TaraxaDposInterfaceCallerSession) GetDelegations(delegator common.Address, batch uint32) (struct {
	Delegations []DposInterfaceDelegationData
	End         bool
}, error) {
	return _TaraxaDposInterface.Contract.GetDelegations(&_TaraxaDposInterface.CallOpts, delegator, batch)
}

// GetTotalEligibleVotesCount is a free data retrieval call binding the contract method 0xde8e4b50.
//
// Solidity: function getTotalEligibleVotesCount() view returns(uint64)
func (_TaraxaDposInterface *TaraxaDposInterfaceCaller) GetTotalEligibleVotesCount(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _TaraxaDposInterface.contract.Call(opts, &out, "getTotalEligibleVotesCount")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetTotalEligibleVotesCount is a free data retrieval call binding the contract method 0xde8e4b50.
//
// Solidity: function getTotalEligibleVotesCount() view returns(uint64)
func (_TaraxaDposInterface *TaraxaDposInterfaceSession) GetTotalEligibleVotesCount() (uint64, error) {
	return _TaraxaDposInterface.Contract.GetTotalEligibleVotesCount(&_TaraxaDposInterface.CallOpts)
}

// GetTotalEligibleVotesCount is a free data retrieval call binding the contract method 0xde8e4b50.
//
// Solidity: function getTotalEligibleVotesCount() view returns(uint64)
func (_TaraxaDposInterface *TaraxaDposInterfaceCallerSession) GetTotalEligibleVotesCount() (uint64, error) {
	return _TaraxaDposInterface.Contract.GetTotalEligibleVotesCount(&_TaraxaDposInterface.CallOpts)
}

// GetUndelegations is a free data retrieval call binding the contract method 0x4edd9943.
//
// Solidity: function getUndelegations(address delegator, uint32 batch) view returns((uint256,uint64,address)[] undelegations, bool end)
func (_TaraxaDposInterface *TaraxaDposInterfaceCaller) GetUndelegations(opts *bind.CallOpts, delegator common.Address, batch uint32) (struct {
	Undelegations []DposInterfaceUndelegationData
	End           bool
}, error) {
	var out []interface{}
	err := _TaraxaDposInterface.contract.Call(opts, &out, "getUndelegations", delegator, batch)

	outstruct := new(struct {
		Undelegations []DposInterfaceUndelegationData
		End           bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Undelegations = *abi.ConvertType(out[0], new([]DposInterfaceUndelegationData)).(*[]DposInterfaceUndelegationData)
	outstruct.End = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetUndelegations is a free data retrieval call binding the contract method 0x4edd9943.
//
// Solidity: function getUndelegations(address delegator, uint32 batch) view returns((uint256,uint64,address)[] undelegations, bool end)
func (_TaraxaDposInterface *TaraxaDposInterfaceSession) GetUndelegations(delegator common.Address, batch uint32) (struct {
	Undelegations []DposInterfaceUndelegationData
	End           bool
}, error) {
	return _TaraxaDposInterface.Contract.GetUndelegations(&_TaraxaDposInterface.CallOpts, delegator, batch)
}

// GetUndelegations is a free data retrieval call binding the contract method 0x4edd9943.
//
// Solidity: function getUndelegations(address delegator, uint32 batch) view returns((uint256,uint64,address)[] undelegations, bool end)
func (_TaraxaDposInterface *TaraxaDposInterfaceCallerSession) GetUndelegations(delegator common.Address, batch uint32) (struct {
	Undelegations []DposInterfaceUndelegationData
	End           bool
}, error) {
	return _TaraxaDposInterface.Contract.GetUndelegations(&_TaraxaDposInterface.CallOpts, delegator, batch)
}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address validator) view returns((uint256,uint256,uint16,uint64,address,string,string) validator_info)
func (_TaraxaDposInterface *TaraxaDposInterfaceCaller) GetValidator(opts *bind.CallOpts, validator common.Address) (DposInterfaceValidatorBasicInfo, error) {
	var out []interface{}
	err := _TaraxaDposInterface.contract.Call(opts, &out, "getValidator", validator)

	if err != nil {
		return *new(DposInterfaceValidatorBasicInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(DposInterfaceValidatorBasicInfo)).(*DposInterfaceValidatorBasicInfo)

	return out0, err

}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address validator) view returns((uint256,uint256,uint16,uint64,address,string,string) validator_info)
func (_TaraxaDposInterface *TaraxaDposInterfaceSession) GetValidator(validator common.Address) (DposInterfaceValidatorBasicInfo, error) {
	return _TaraxaDposInterface.Contract.GetValidator(&_TaraxaDposInterface.CallOpts, validator)
}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address validator) view returns((uint256,uint256,uint16,uint64,address,string,string) validator_info)
func (_TaraxaDposInterface *TaraxaDposInterfaceCallerSession) GetValidator(validator common.Address) (DposInterfaceValidatorBasicInfo, error) {
	return _TaraxaDposInterface.Contract.GetValidator(&_TaraxaDposInterface.CallOpts, validator)
}

// GetValidatorDelegators is a free data retrieval call binding the contract method 0xc1deb4f1.
//
// Solidity: function getValidatorDelegators(address validator) view returns(address[] delegators)
func (_TaraxaDposInterface *TaraxaDposInterfaceCaller) GetValidatorDelegators(opts *bind.CallOpts, validator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _TaraxaDposInterface.contract.Call(opts, &out, "getValidatorDelegators", validator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidatorDelegators is a free data retrieval call binding the contract method 0xc1deb4f1.
//
// Solidity: function getValidatorDelegators(address validator) view returns(address[] delegators)
func (_TaraxaDposInterface *TaraxaDposInterfaceSession) GetValidatorDelegators(validator common.Address) ([]common.Address, error) {
	return _TaraxaDposInterface.Contract.GetValidatorDelegators(&_TaraxaDposInterface.CallOpts, validator)
}

// GetValidatorDelegators is a free data retrieval call binding the contract method 0xc1deb4f1.
//
// Solidity: function getValidatorDelegators(address validator) view returns(address[] delegators)
func (_TaraxaDposInterface *TaraxaDposInterfaceCallerSession) GetValidatorDelegators(validator common.Address) ([]common.Address, error) {
	return _TaraxaDposInterface.Contract.GetValidatorDelegators(&_TaraxaDposInterface.CallOpts, validator)
}

// GetValidatorEligibleVotesCount is a free data retrieval call binding the contract method 0x618e3862.
//
// Solidity: function getValidatorEligibleVotesCount(address validator) view returns(uint64)
func (_TaraxaDposInterface *TaraxaDposInterfaceCaller) GetValidatorEligibleVotesCount(opts *bind.CallOpts, validator common.Address) (uint64, error) {
	var out []interface{}
	err := _TaraxaDposInterface.contract.Call(opts, &out, "getValidatorEligibleVotesCount", validator)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetValidatorEligibleVotesCount is a free data retrieval call binding the contract method 0x618e3862.
//
// Solidity: function getValidatorEligibleVotesCount(address validator) view returns(uint64)
func (_TaraxaDposInterface *TaraxaDposInterfaceSession) GetValidatorEligibleVotesCount(validator common.Address) (uint64, error) {
	return _TaraxaDposInterface.Contract.GetValidatorEligibleVotesCount(&_TaraxaDposInterface.CallOpts, validator)
}

// GetValidatorEligibleVotesCount is a free data retrieval call binding the contract method 0x618e3862.
//
// Solidity: function getValidatorEligibleVotesCount(address validator) view returns(uint64)
func (_TaraxaDposInterface *TaraxaDposInterfaceCallerSession) GetValidatorEligibleVotesCount(validator common.Address) (uint64, error) {
	return _TaraxaDposInterface.Contract.GetValidatorEligibleVotesCount(&_TaraxaDposInterface.CallOpts, validator)
}

// GetValidators is a free data retrieval call binding the contract method 0x19d8024f.
//
// Solidity: function getValidators(uint32 batch) view returns((address,(uint256,uint256,uint16,uint64,address,string,string))[] validators, bool end)
func (_TaraxaDposInterface *TaraxaDposInterfaceCaller) GetValidators(opts *bind.CallOpts, batch uint32) (struct {
	Validators []DposInterfaceValidatorData
	End        bool
}, error) {
	var out []interface{}
	err := _TaraxaDposInterface.contract.Call(opts, &out, "getValidators", batch)

	outstruct := new(struct {
		Validators []DposInterfaceValidatorData
		End        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Validators = *abi.ConvertType(out[0], new([]DposInterfaceValidatorData)).(*[]DposInterfaceValidatorData)
	outstruct.End = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetValidators is a free data retrieval call binding the contract method 0x19d8024f.
//
// Solidity: function getValidators(uint32 batch) view returns((address,(uint256,uint256,uint16,uint64,address,string,string))[] validators, bool end)
func (_TaraxaDposInterface *TaraxaDposInterfaceSession) GetValidators(batch uint32) (struct {
	Validators []DposInterfaceValidatorData
	End        bool
}, error) {
	return _TaraxaDposInterface.Contract.GetValidators(&_TaraxaDposInterface.CallOpts, batch)
}

// GetValidators is a free data retrieval call binding the contract method 0x19d8024f.
//
// Solidity: function getValidators(uint32 batch) view returns((address,(uint256,uint256,uint16,uint64,address,string,string))[] validators, bool end)
func (_TaraxaDposInterface *TaraxaDposInterfaceCallerSession) GetValidators(batch uint32) (struct {
	Validators []DposInterfaceValidatorData
	End        bool
}, error) {
	return _TaraxaDposInterface.Contract.GetValidators(&_TaraxaDposInterface.CallOpts, batch)
}

// GetValidatorsFor is a free data retrieval call binding the contract method 0x724ac6b0.
//
// Solidity: function getValidatorsFor(address owner, uint32 batch) view returns((address,(uint256,uint256,uint16,uint64,address,string,string))[] validators, bool end)
func (_TaraxaDposInterface *TaraxaDposInterfaceCaller) GetValidatorsFor(opts *bind.CallOpts, owner common.Address, batch uint32) (struct {
	Validators []DposInterfaceValidatorData
	End        bool
}, error) {
	var out []interface{}
	err := _TaraxaDposInterface.contract.Call(opts, &out, "getValidatorsFor", owner, batch)

	outstruct := new(struct {
		Validators []DposInterfaceValidatorData
		End        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Validators = *abi.ConvertType(out[0], new([]DposInterfaceValidatorData)).(*[]DposInterfaceValidatorData)
	outstruct.End = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetValidatorsFor is a free data retrieval call binding the contract method 0x724ac6b0.
//
// Solidity: function getValidatorsFor(address owner, uint32 batch) view returns((address,(uint256,uint256,uint16,uint64,address,string,string))[] validators, bool end)
func (_TaraxaDposInterface *TaraxaDposInterfaceSession) GetValidatorsFor(owner common.Address, batch uint32) (struct {
	Validators []DposInterfaceValidatorData
	End        bool
}, error) {
	return _TaraxaDposInterface.Contract.GetValidatorsFor(&_TaraxaDposInterface.CallOpts, owner, batch)
}

// GetValidatorsFor is a free data retrieval call binding the contract method 0x724ac6b0.
//
// Solidity: function getValidatorsFor(address owner, uint32 batch) view returns((address,(uint256,uint256,uint16,uint64,address,string,string))[] validators, bool end)
func (_TaraxaDposInterface *TaraxaDposInterfaceCallerSession) GetValidatorsFor(owner common.Address, batch uint32) (struct {
	Validators []DposInterfaceValidatorData
	End        bool
}, error) {
	return _TaraxaDposInterface.Contract.GetValidatorsFor(&_TaraxaDposInterface.CallOpts, owner, batch)
}

// IsValidatorEligible is a free data retrieval call binding the contract method 0xf3094e90.
//
// Solidity: function isValidatorEligible(address validator) view returns(bool)
func (_TaraxaDposInterface *TaraxaDposInterfaceCaller) IsValidatorEligible(opts *bind.CallOpts, validator common.Address) (bool, error) {
	var out []interface{}
	err := _TaraxaDposInterface.contract.Call(opts, &out, "isValidatorEligible", validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidatorEligible is a free data retrieval call binding the contract method 0xf3094e90.
//
// Solidity: function isValidatorEligible(address validator) view returns(bool)
func (_TaraxaDposInterface *TaraxaDposInterfaceSession) IsValidatorEligible(validator common.Address) (bool, error) {
	return _TaraxaDposInterface.Contract.IsValidatorEligible(&_TaraxaDposInterface.CallOpts, validator)
}

// IsValidatorEligible is a free data retrieval call binding the contract method 0xf3094e90.
//
// Solidity: function isValidatorEligible(address validator) view returns(bool)
func (_TaraxaDposInterface *TaraxaDposInterfaceCallerSession) IsValidatorEligible(validator common.Address) (bool, error) {
	return _TaraxaDposInterface.Contract.IsValidatorEligible(&_TaraxaDposInterface.CallOpts, validator)
}

// CancelUndelegate is a paid mutator transaction binding the contract method 0x399ff554.
//
// Solidity: function cancelUndelegate(address validator) returns()
func (_TaraxaDposInterface *TaraxaDposInterfaceTransactor) CancelUndelegate(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _TaraxaDposInterface.contract.Transact(opts, "cancelUndelegate", validator)
}

// CancelUndelegate is a paid mutator transaction binding the contract method 0x399ff554.
//
// Solidity: function cancelUndelegate(address validator) returns()
func (_TaraxaDposInterface *TaraxaDposInterfaceSession) CancelUndelegate(validator common.Address) (*types.Transaction, error) {
	return _TaraxaDposInterface.Contract.CancelUndelegate(&_TaraxaDposInterface.TransactOpts, validator)
}

// CancelUndelegate is a paid mutator transaction binding the contract method 0x399ff554.
//
// Solidity: function cancelUndelegate(address validator) returns()
func (_TaraxaDposInterface *TaraxaDposInterfaceTransactorSession) CancelUndelegate(validator common.Address) (*types.Transaction, error) {
	return _TaraxaDposInterface.Contract.CancelUndelegate(&_TaraxaDposInterface.TransactOpts, validator)
}

// ClaimCommissionRewards is a paid mutator transaction binding the contract method 0xd0eebfe2.
//
// Solidity: function claimCommissionRewards(address validator) returns()
func (_TaraxaDposInterface *TaraxaDposInterfaceTransactor) ClaimCommissionRewards(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _TaraxaDposInterface.contract.Transact(opts, "claimCommissionRewards", validator)
}

// ClaimCommissionRewards is a paid mutator transaction binding the contract method 0xd0eebfe2.
//
// Solidity: function claimCommissionRewards(address validator) returns()
func (_TaraxaDposInterface *TaraxaDposInterfaceSession) ClaimCommissionRewards(validator common.Address) (*types.Transaction, error) {
	return _TaraxaDposInterface.Contract.ClaimCommissionRewards(&_TaraxaDposInterface.TransactOpts, validator)
}

// ClaimCommissionRewards is a paid mutator transaction binding the contract method 0xd0eebfe2.
//
// Solidity: function claimCommissionRewards(address validator) returns()
func (_TaraxaDposInterface *TaraxaDposInterfaceTransactorSession) ClaimCommissionRewards(validator common.Address) (*types.Transaction, error) {
	return _TaraxaDposInterface.Contract.ClaimCommissionRewards(&_TaraxaDposInterface.TransactOpts, validator)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address validator) returns()
func (_TaraxaDposInterface *TaraxaDposInterfaceTransactor) ClaimRewards(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _TaraxaDposInterface.contract.Transact(opts, "claimRewards", validator)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address validator) returns()
func (_TaraxaDposInterface *TaraxaDposInterfaceSession) ClaimRewards(validator common.Address) (*types.Transaction, error) {
	return _TaraxaDposInterface.Contract.ClaimRewards(&_TaraxaDposInterface.TransactOpts, validator)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address validator) returns()
func (_TaraxaDposInterface *TaraxaDposInterfaceTransactorSession) ClaimRewards(validator common.Address) (*types.Transaction, error) {
	return _TaraxaDposInterface.Contract.ClaimRewards(&_TaraxaDposInterface.TransactOpts, validator)
}

// ConfirmUndelegate is a paid mutator transaction binding the contract method 0x45a02561.
//
// Solidity: function confirmUndelegate(address validator) returns()
func (_TaraxaDposInterface *TaraxaDposInterfaceTransactor) ConfirmUndelegate(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _TaraxaDposInterface.contract.Transact(opts, "confirmUndelegate", validator)
}

// ConfirmUndelegate is a paid mutator transaction binding the contract method 0x45a02561.
//
// Solidity: function confirmUndelegate(address validator) returns()
func (_TaraxaDposInterface *TaraxaDposInterfaceSession) ConfirmUndelegate(validator common.Address) (*types.Transaction, error) {
	return _TaraxaDposInterface.Contract.ConfirmUndelegate(&_TaraxaDposInterface.TransactOpts, validator)
}

// ConfirmUndelegate is a paid mutator transaction binding the contract method 0x45a02561.
//
// Solidity: function confirmUndelegate(address validator) returns()
func (_TaraxaDposInterface *TaraxaDposInterfaceTransactorSession) ConfirmUndelegate(validator common.Address) (*types.Transaction, error) {
	return _TaraxaDposInterface.Contract.ConfirmUndelegate(&_TaraxaDposInterface.TransactOpts, validator)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address validator) payable returns()
func (_TaraxaDposInterface *TaraxaDposInterfaceTransactor) Delegate(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _TaraxaDposInterface.contract.Transact(opts, "delegate", validator)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address validator) payable returns()
func (_TaraxaDposInterface *TaraxaDposInterfaceSession) Delegate(validator common.Address) (*types.Transaction, error) {
	return _TaraxaDposInterface.Contract.Delegate(&_TaraxaDposInterface.TransactOpts, validator)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address validator) payable returns()
func (_TaraxaDposInterface *TaraxaDposInterfaceTransactorSession) Delegate(validator common.Address) (*types.Transaction, error) {
	return _TaraxaDposInterface.Contract.Delegate(&_TaraxaDposInterface.TransactOpts, validator)
}

// ReDelegate is a paid mutator transaction binding the contract method 0x703812cc.
//
// Solidity: function reDelegate(address validator_from, address validator_to, uint256 amount) returns()
func (_TaraxaDposInterface *TaraxaDposInterfaceTransactor) ReDelegate(opts *bind.TransactOpts, validator_from common.Address, validator_to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaraxaDposInterface.contract.Transact(opts, "reDelegate", validator_from, validator_to, amount)
}

// ReDelegate is a paid mutator transaction binding the contract method 0x703812cc.
//
// Solidity: function reDelegate(address validator_from, address validator_to, uint256 amount) returns()
func (_TaraxaDposInterface *TaraxaDposInterfaceSession) ReDelegate(validator_from common.Address, validator_to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaraxaDposInterface.Contract.ReDelegate(&_TaraxaDposInterface.TransactOpts, validator_from, validator_to, amount)
}

// ReDelegate is a paid mutator transaction binding the contract method 0x703812cc.
//
// Solidity: function reDelegate(address validator_from, address validator_to, uint256 amount) returns()
func (_TaraxaDposInterface *TaraxaDposInterfaceTransactorSession) ReDelegate(validator_from common.Address, validator_to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaraxaDposInterface.Contract.ReDelegate(&_TaraxaDposInterface.TransactOpts, validator_from, validator_to, amount)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xd6fdc127.
//
// Solidity: function registerValidator(address validator, bytes proof, bytes vrf_key, uint16 commission, string description, string endpoint) payable returns()
func (_TaraxaDposInterface *TaraxaDposInterfaceTransactor) RegisterValidator(opts *bind.TransactOpts, validator common.Address, proof []byte, vrf_key []byte, commission uint16, description string, endpoint string) (*types.Transaction, error) {
	return _TaraxaDposInterface.contract.Transact(opts, "registerValidator", validator, proof, vrf_key, commission, description, endpoint)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xd6fdc127.
//
// Solidity: function registerValidator(address validator, bytes proof, bytes vrf_key, uint16 commission, string description, string endpoint) payable returns()
func (_TaraxaDposInterface *TaraxaDposInterfaceSession) RegisterValidator(validator common.Address, proof []byte, vrf_key []byte, commission uint16, description string, endpoint string) (*types.Transaction, error) {
	return _TaraxaDposInterface.Contract.RegisterValidator(&_TaraxaDposInterface.TransactOpts, validator, proof, vrf_key, commission, description, endpoint)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xd6fdc127.
//
// Solidity: function registerValidator(address validator, bytes proof, bytes vrf_key, uint16 commission, string description, string endpoint) payable returns()
func (_TaraxaDposInterface *TaraxaDposInterfaceTransactorSession) RegisterValidator(validator common.Address, proof []byte, vrf_key []byte, commission uint16, description string, endpoint string) (*types.Transaction, error) {
	return _TaraxaDposInterface.Contract.RegisterValidator(&_TaraxaDposInterface.TransactOpts, validator, proof, vrf_key, commission, description, endpoint)
}

// SetCommission is a paid mutator transaction binding the contract method 0xf000322c.
//
// Solidity: function setCommission(address validator, uint16 commission) returns()
func (_TaraxaDposInterface *TaraxaDposInterfaceTransactor) SetCommission(opts *bind.TransactOpts, validator common.Address, commission uint16) (*types.Transaction, error) {
	return _TaraxaDposInterface.contract.Transact(opts, "setCommission", validator, commission)
}

// SetCommission is a paid mutator transaction binding the contract method 0xf000322c.
//
// Solidity: function setCommission(address validator, uint16 commission) returns()
func (_TaraxaDposInterface *TaraxaDposInterfaceSession) SetCommission(validator common.Address, commission uint16) (*types.Transaction, error) {
	return _TaraxaDposInterface.Contract.SetCommission(&_TaraxaDposInterface.TransactOpts, validator, commission)
}

// SetCommission is a paid mutator transaction binding the contract method 0xf000322c.
//
// Solidity: function setCommission(address validator, uint16 commission) returns()
func (_TaraxaDposInterface *TaraxaDposInterfaceTransactorSession) SetCommission(validator common.Address, commission uint16) (*types.Transaction, error) {
	return _TaraxaDposInterface.Contract.SetCommission(&_TaraxaDposInterface.TransactOpts, validator, commission)
}

// SetValidatorInfo is a paid mutator transaction binding the contract method 0x0babea4c.
//
// Solidity: function setValidatorInfo(address validator, string description, string endpoint) returns()
func (_TaraxaDposInterface *TaraxaDposInterfaceTransactor) SetValidatorInfo(opts *bind.TransactOpts, validator common.Address, description string, endpoint string) (*types.Transaction, error) {
	return _TaraxaDposInterface.contract.Transact(opts, "setValidatorInfo", validator, description, endpoint)
}

// SetValidatorInfo is a paid mutator transaction binding the contract method 0x0babea4c.
//
// Solidity: function setValidatorInfo(address validator, string description, string endpoint) returns()
func (_TaraxaDposInterface *TaraxaDposInterfaceSession) SetValidatorInfo(validator common.Address, description string, endpoint string) (*types.Transaction, error) {
	return _TaraxaDposInterface.Contract.SetValidatorInfo(&_TaraxaDposInterface.TransactOpts, validator, description, endpoint)
}

// SetValidatorInfo is a paid mutator transaction binding the contract method 0x0babea4c.
//
// Solidity: function setValidatorInfo(address validator, string description, string endpoint) returns()
func (_TaraxaDposInterface *TaraxaDposInterfaceTransactorSession) SetValidatorInfo(validator common.Address, description string, endpoint string) (*types.Transaction, error) {
	return _TaraxaDposInterface.Contract.SetValidatorInfo(&_TaraxaDposInterface.TransactOpts, validator, description, endpoint)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validator, uint256 amount) returns()
func (_TaraxaDposInterface *TaraxaDposInterfaceTransactor) Undelegate(opts *bind.TransactOpts, validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaraxaDposInterface.contract.Transact(opts, "undelegate", validator, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validator, uint256 amount) returns()
func (_TaraxaDposInterface *TaraxaDposInterfaceSession) Undelegate(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaraxaDposInterface.Contract.Undelegate(&_TaraxaDposInterface.TransactOpts, validator, amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address validator, uint256 amount) returns()
func (_TaraxaDposInterface *TaraxaDposInterfaceTransactorSession) Undelegate(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TaraxaDposInterface.Contract.Undelegate(&_TaraxaDposInterface.TransactOpts, validator, amount)
}

// TaraxaDposInterfaceCommissionRewardsClaimedIterator is returned from FilterCommissionRewardsClaimed and is used to iterate over the raw logs and unpacked data for CommissionRewardsClaimed events raised by the TaraxaDposInterface contract.
type TaraxaDposInterfaceCommissionRewardsClaimedIterator struct {
	Event *TaraxaDposInterfaceCommissionRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *TaraxaDposInterfaceCommissionRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraxaDposInterfaceCommissionRewardsClaimed)
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
		it.Event = new(TaraxaDposInterfaceCommissionRewardsClaimed)
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
func (it *TaraxaDposInterfaceCommissionRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraxaDposInterfaceCommissionRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraxaDposInterfaceCommissionRewardsClaimed represents a CommissionRewardsClaimed event raised by the TaraxaDposInterface contract.
type TaraxaDposInterfaceCommissionRewardsClaimed struct {
	Account   common.Address
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCommissionRewardsClaimed is a free log retrieval operation binding the contract event 0x9fe6fb5c5703216904280d8ed59c57daa2bedb91e204b4f23e347ecaaf313c3b.
//
// Solidity: event CommissionRewardsClaimed(address indexed account, address indexed validator)
func (_TaraxaDposInterface *TaraxaDposInterfaceFilterer) FilterCommissionRewardsClaimed(opts *bind.FilterOpts, account []common.Address, validator []common.Address) (*TaraxaDposInterfaceCommissionRewardsClaimedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposInterface.contract.FilterLogs(opts, "CommissionRewardsClaimed", accountRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &TaraxaDposInterfaceCommissionRewardsClaimedIterator{contract: _TaraxaDposInterface.contract, event: "CommissionRewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchCommissionRewardsClaimed is a free log subscription operation binding the contract event 0x9fe6fb5c5703216904280d8ed59c57daa2bedb91e204b4f23e347ecaaf313c3b.
//
// Solidity: event CommissionRewardsClaimed(address indexed account, address indexed validator)
func (_TaraxaDposInterface *TaraxaDposInterfaceFilterer) WatchCommissionRewardsClaimed(opts *bind.WatchOpts, sink chan<- *TaraxaDposInterfaceCommissionRewardsClaimed, account []common.Address, validator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposInterface.contract.WatchLogs(opts, "CommissionRewardsClaimed", accountRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraxaDposInterfaceCommissionRewardsClaimed)
				if err := _TaraxaDposInterface.contract.UnpackLog(event, "CommissionRewardsClaimed", log); err != nil {
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

// ParseCommissionRewardsClaimed is a log parse operation binding the contract event 0x9fe6fb5c5703216904280d8ed59c57daa2bedb91e204b4f23e347ecaaf313c3b.
//
// Solidity: event CommissionRewardsClaimed(address indexed account, address indexed validator)
func (_TaraxaDposInterface *TaraxaDposInterfaceFilterer) ParseCommissionRewardsClaimed(log types.Log) (*TaraxaDposInterfaceCommissionRewardsClaimed, error) {
	event := new(TaraxaDposInterfaceCommissionRewardsClaimed)
	if err := _TaraxaDposInterface.contract.UnpackLog(event, "CommissionRewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaraxaDposInterfaceCommissionSetIterator is returned from FilterCommissionSet and is used to iterate over the raw logs and unpacked data for CommissionSet events raised by the TaraxaDposInterface contract.
type TaraxaDposInterfaceCommissionSetIterator struct {
	Event *TaraxaDposInterfaceCommissionSet // Event containing the contract specifics and raw log

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
func (it *TaraxaDposInterfaceCommissionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraxaDposInterfaceCommissionSet)
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
		it.Event = new(TaraxaDposInterfaceCommissionSet)
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
func (it *TaraxaDposInterfaceCommissionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraxaDposInterfaceCommissionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraxaDposInterfaceCommissionSet represents a CommissionSet event raised by the TaraxaDposInterface contract.
type TaraxaDposInterfaceCommissionSet struct {
	Validator common.Address
	Comission uint16
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCommissionSet is a free log retrieval operation binding the contract event 0xc909daf778d180f43dac53b55d0de934d2f1e0b70412ca274982e4e6e894eb1a.
//
// Solidity: event CommissionSet(address indexed validator, uint16 comission)
func (_TaraxaDposInterface *TaraxaDposInterfaceFilterer) FilterCommissionSet(opts *bind.FilterOpts, validator []common.Address) (*TaraxaDposInterfaceCommissionSetIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposInterface.contract.FilterLogs(opts, "CommissionSet", validatorRule)
	if err != nil {
		return nil, err
	}
	return &TaraxaDposInterfaceCommissionSetIterator{contract: _TaraxaDposInterface.contract, event: "CommissionSet", logs: logs, sub: sub}, nil
}

// WatchCommissionSet is a free log subscription operation binding the contract event 0xc909daf778d180f43dac53b55d0de934d2f1e0b70412ca274982e4e6e894eb1a.
//
// Solidity: event CommissionSet(address indexed validator, uint16 comission)
func (_TaraxaDposInterface *TaraxaDposInterfaceFilterer) WatchCommissionSet(opts *bind.WatchOpts, sink chan<- *TaraxaDposInterfaceCommissionSet, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposInterface.contract.WatchLogs(opts, "CommissionSet", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraxaDposInterfaceCommissionSet)
				if err := _TaraxaDposInterface.contract.UnpackLog(event, "CommissionSet", log); err != nil {
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

// ParseCommissionSet is a log parse operation binding the contract event 0xc909daf778d180f43dac53b55d0de934d2f1e0b70412ca274982e4e6e894eb1a.
//
// Solidity: event CommissionSet(address indexed validator, uint16 comission)
func (_TaraxaDposInterface *TaraxaDposInterfaceFilterer) ParseCommissionSet(log types.Log) (*TaraxaDposInterfaceCommissionSet, error) {
	event := new(TaraxaDposInterfaceCommissionSet)
	if err := _TaraxaDposInterface.contract.UnpackLog(event, "CommissionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaraxaDposInterfaceDelegatedIterator is returned from FilterDelegated and is used to iterate over the raw logs and unpacked data for Delegated events raised by the TaraxaDposInterface contract.
type TaraxaDposInterfaceDelegatedIterator struct {
	Event *TaraxaDposInterfaceDelegated // Event containing the contract specifics and raw log

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
func (it *TaraxaDposInterfaceDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraxaDposInterfaceDelegated)
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
		it.Event = new(TaraxaDposInterfaceDelegated)
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
func (it *TaraxaDposInterfaceDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraxaDposInterfaceDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraxaDposInterfaceDelegated represents a Delegated event raised by the TaraxaDposInterface contract.
type TaraxaDposInterfaceDelegated struct {
	Delegator common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegated is a free log retrieval operation binding the contract event 0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b.
//
// Solidity: event Delegated(address indexed delegator, address indexed validator, uint256 amount)
func (_TaraxaDposInterface *TaraxaDposInterfaceFilterer) FilterDelegated(opts *bind.FilterOpts, delegator []common.Address, validator []common.Address) (*TaraxaDposInterfaceDelegatedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposInterface.contract.FilterLogs(opts, "Delegated", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &TaraxaDposInterfaceDelegatedIterator{contract: _TaraxaDposInterface.contract, event: "Delegated", logs: logs, sub: sub}, nil
}

// WatchDelegated is a free log subscription operation binding the contract event 0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b.
//
// Solidity: event Delegated(address indexed delegator, address indexed validator, uint256 amount)
func (_TaraxaDposInterface *TaraxaDposInterfaceFilterer) WatchDelegated(opts *bind.WatchOpts, sink chan<- *TaraxaDposInterfaceDelegated, delegator []common.Address, validator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposInterface.contract.WatchLogs(opts, "Delegated", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraxaDposInterfaceDelegated)
				if err := _TaraxaDposInterface.contract.UnpackLog(event, "Delegated", log); err != nil {
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

// ParseDelegated is a log parse operation binding the contract event 0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b.
//
// Solidity: event Delegated(address indexed delegator, address indexed validator, uint256 amount)
func (_TaraxaDposInterface *TaraxaDposInterfaceFilterer) ParseDelegated(log types.Log) (*TaraxaDposInterfaceDelegated, error) {
	event := new(TaraxaDposInterfaceDelegated)
	if err := _TaraxaDposInterface.contract.UnpackLog(event, "Delegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaraxaDposInterfaceRedelegatedIterator is returned from FilterRedelegated and is used to iterate over the raw logs and unpacked data for Redelegated events raised by the TaraxaDposInterface contract.
type TaraxaDposInterfaceRedelegatedIterator struct {
	Event *TaraxaDposInterfaceRedelegated // Event containing the contract specifics and raw log

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
func (it *TaraxaDposInterfaceRedelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraxaDposInterfaceRedelegated)
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
		it.Event = new(TaraxaDposInterfaceRedelegated)
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
func (it *TaraxaDposInterfaceRedelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraxaDposInterfaceRedelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraxaDposInterfaceRedelegated represents a Redelegated event raised by the TaraxaDposInterface contract.
type TaraxaDposInterfaceRedelegated struct {
	Delegator common.Address
	From      common.Address
	To        common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRedelegated is a free log retrieval operation binding the contract event 0x12e144c27d0bad08abc77c66a640b5cf15a03a93f6582f40de6932b033a5fa5e.
//
// Solidity: event Redelegated(address indexed delegator, address indexed from, address indexed to, uint256 amount)
func (_TaraxaDposInterface *TaraxaDposInterfaceFilterer) FilterRedelegated(opts *bind.FilterOpts, delegator []common.Address, from []common.Address, to []common.Address) (*TaraxaDposInterfaceRedelegatedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TaraxaDposInterface.contract.FilterLogs(opts, "Redelegated", delegatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TaraxaDposInterfaceRedelegatedIterator{contract: _TaraxaDposInterface.contract, event: "Redelegated", logs: logs, sub: sub}, nil
}

// WatchRedelegated is a free log subscription operation binding the contract event 0x12e144c27d0bad08abc77c66a640b5cf15a03a93f6582f40de6932b033a5fa5e.
//
// Solidity: event Redelegated(address indexed delegator, address indexed from, address indexed to, uint256 amount)
func (_TaraxaDposInterface *TaraxaDposInterfaceFilterer) WatchRedelegated(opts *bind.WatchOpts, sink chan<- *TaraxaDposInterfaceRedelegated, delegator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TaraxaDposInterface.contract.WatchLogs(opts, "Redelegated", delegatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraxaDposInterfaceRedelegated)
				if err := _TaraxaDposInterface.contract.UnpackLog(event, "Redelegated", log); err != nil {
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

// ParseRedelegated is a log parse operation binding the contract event 0x12e144c27d0bad08abc77c66a640b5cf15a03a93f6582f40de6932b033a5fa5e.
//
// Solidity: event Redelegated(address indexed delegator, address indexed from, address indexed to, uint256 amount)
func (_TaraxaDposInterface *TaraxaDposInterfaceFilterer) ParseRedelegated(log types.Log) (*TaraxaDposInterfaceRedelegated, error) {
	event := new(TaraxaDposInterfaceRedelegated)
	if err := _TaraxaDposInterface.contract.UnpackLog(event, "Redelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaraxaDposInterfaceRewardsClaimedIterator is returned from FilterRewardsClaimed and is used to iterate over the raw logs and unpacked data for RewardsClaimed events raised by the TaraxaDposInterface contract.
type TaraxaDposInterfaceRewardsClaimedIterator struct {
	Event *TaraxaDposInterfaceRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *TaraxaDposInterfaceRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraxaDposInterfaceRewardsClaimed)
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
		it.Event = new(TaraxaDposInterfaceRewardsClaimed)
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
func (it *TaraxaDposInterfaceRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraxaDposInterfaceRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraxaDposInterfaceRewardsClaimed represents a RewardsClaimed event raised by the TaraxaDposInterface contract.
type TaraxaDposInterfaceRewardsClaimed struct {
	Account   common.Address
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardsClaimed is a free log retrieval operation binding the contract event 0x76fefaa578826cd28a6d72d47460f755534b9a481c773721143222296cc4cc5a.
//
// Solidity: event RewardsClaimed(address indexed account, address indexed validator)
func (_TaraxaDposInterface *TaraxaDposInterfaceFilterer) FilterRewardsClaimed(opts *bind.FilterOpts, account []common.Address, validator []common.Address) (*TaraxaDposInterfaceRewardsClaimedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposInterface.contract.FilterLogs(opts, "RewardsClaimed", accountRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &TaraxaDposInterfaceRewardsClaimedIterator{contract: _TaraxaDposInterface.contract, event: "RewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardsClaimed is a free log subscription operation binding the contract event 0x76fefaa578826cd28a6d72d47460f755534b9a481c773721143222296cc4cc5a.
//
// Solidity: event RewardsClaimed(address indexed account, address indexed validator)
func (_TaraxaDposInterface *TaraxaDposInterfaceFilterer) WatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *TaraxaDposInterfaceRewardsClaimed, account []common.Address, validator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposInterface.contract.WatchLogs(opts, "RewardsClaimed", accountRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraxaDposInterfaceRewardsClaimed)
				if err := _TaraxaDposInterface.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
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

// ParseRewardsClaimed is a log parse operation binding the contract event 0x76fefaa578826cd28a6d72d47460f755534b9a481c773721143222296cc4cc5a.
//
// Solidity: event RewardsClaimed(address indexed account, address indexed validator)
func (_TaraxaDposInterface *TaraxaDposInterfaceFilterer) ParseRewardsClaimed(log types.Log) (*TaraxaDposInterfaceRewardsClaimed, error) {
	event := new(TaraxaDposInterfaceRewardsClaimed)
	if err := _TaraxaDposInterface.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaraxaDposInterfaceUndelegateCanceledIterator is returned from FilterUndelegateCanceled and is used to iterate over the raw logs and unpacked data for UndelegateCanceled events raised by the TaraxaDposInterface contract.
type TaraxaDposInterfaceUndelegateCanceledIterator struct {
	Event *TaraxaDposInterfaceUndelegateCanceled // Event containing the contract specifics and raw log

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
func (it *TaraxaDposInterfaceUndelegateCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraxaDposInterfaceUndelegateCanceled)
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
		it.Event = new(TaraxaDposInterfaceUndelegateCanceled)
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
func (it *TaraxaDposInterfaceUndelegateCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraxaDposInterfaceUndelegateCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraxaDposInterfaceUndelegateCanceled represents a UndelegateCanceled event raised by the TaraxaDposInterface contract.
type TaraxaDposInterfaceUndelegateCanceled struct {
	Delegator common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUndelegateCanceled is a free log retrieval operation binding the contract event 0xfc25f8a919d19f2c2dfce21115718abc9ef2b1e0c9218a488f614c75be4184b7.
//
// Solidity: event UndelegateCanceled(address indexed delegator, address indexed validator, uint256 amount)
func (_TaraxaDposInterface *TaraxaDposInterfaceFilterer) FilterUndelegateCanceled(opts *bind.FilterOpts, delegator []common.Address, validator []common.Address) (*TaraxaDposInterfaceUndelegateCanceledIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposInterface.contract.FilterLogs(opts, "UndelegateCanceled", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &TaraxaDposInterfaceUndelegateCanceledIterator{contract: _TaraxaDposInterface.contract, event: "UndelegateCanceled", logs: logs, sub: sub}, nil
}

// WatchUndelegateCanceled is a free log subscription operation binding the contract event 0xfc25f8a919d19f2c2dfce21115718abc9ef2b1e0c9218a488f614c75be4184b7.
//
// Solidity: event UndelegateCanceled(address indexed delegator, address indexed validator, uint256 amount)
func (_TaraxaDposInterface *TaraxaDposInterfaceFilterer) WatchUndelegateCanceled(opts *bind.WatchOpts, sink chan<- *TaraxaDposInterfaceUndelegateCanceled, delegator []common.Address, validator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposInterface.contract.WatchLogs(opts, "UndelegateCanceled", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraxaDposInterfaceUndelegateCanceled)
				if err := _TaraxaDposInterface.contract.UnpackLog(event, "UndelegateCanceled", log); err != nil {
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

// ParseUndelegateCanceled is a log parse operation binding the contract event 0xfc25f8a919d19f2c2dfce21115718abc9ef2b1e0c9218a488f614c75be4184b7.
//
// Solidity: event UndelegateCanceled(address indexed delegator, address indexed validator, uint256 amount)
func (_TaraxaDposInterface *TaraxaDposInterfaceFilterer) ParseUndelegateCanceled(log types.Log) (*TaraxaDposInterfaceUndelegateCanceled, error) {
	event := new(TaraxaDposInterfaceUndelegateCanceled)
	if err := _TaraxaDposInterface.contract.UnpackLog(event, "UndelegateCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaraxaDposInterfaceUndelegateConfirmedIterator is returned from FilterUndelegateConfirmed and is used to iterate over the raw logs and unpacked data for UndelegateConfirmed events raised by the TaraxaDposInterface contract.
type TaraxaDposInterfaceUndelegateConfirmedIterator struct {
	Event *TaraxaDposInterfaceUndelegateConfirmed // Event containing the contract specifics and raw log

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
func (it *TaraxaDposInterfaceUndelegateConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraxaDposInterfaceUndelegateConfirmed)
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
		it.Event = new(TaraxaDposInterfaceUndelegateConfirmed)
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
func (it *TaraxaDposInterfaceUndelegateConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraxaDposInterfaceUndelegateConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraxaDposInterfaceUndelegateConfirmed represents a UndelegateConfirmed event raised by the TaraxaDposInterface contract.
type TaraxaDposInterfaceUndelegateConfirmed struct {
	Delegator common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUndelegateConfirmed is a free log retrieval operation binding the contract event 0xf8bef3a6fe3b4c932b5b51c6472a89f171d039f4bacf18cff632208938bf0426.
//
// Solidity: event UndelegateConfirmed(address indexed delegator, address indexed validator, uint256 amount)
func (_TaraxaDposInterface *TaraxaDposInterfaceFilterer) FilterUndelegateConfirmed(opts *bind.FilterOpts, delegator []common.Address, validator []common.Address) (*TaraxaDposInterfaceUndelegateConfirmedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposInterface.contract.FilterLogs(opts, "UndelegateConfirmed", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &TaraxaDposInterfaceUndelegateConfirmedIterator{contract: _TaraxaDposInterface.contract, event: "UndelegateConfirmed", logs: logs, sub: sub}, nil
}

// WatchUndelegateConfirmed is a free log subscription operation binding the contract event 0xf8bef3a6fe3b4c932b5b51c6472a89f171d039f4bacf18cff632208938bf0426.
//
// Solidity: event UndelegateConfirmed(address indexed delegator, address indexed validator, uint256 amount)
func (_TaraxaDposInterface *TaraxaDposInterfaceFilterer) WatchUndelegateConfirmed(opts *bind.WatchOpts, sink chan<- *TaraxaDposInterfaceUndelegateConfirmed, delegator []common.Address, validator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposInterface.contract.WatchLogs(opts, "UndelegateConfirmed", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraxaDposInterfaceUndelegateConfirmed)
				if err := _TaraxaDposInterface.contract.UnpackLog(event, "UndelegateConfirmed", log); err != nil {
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

// ParseUndelegateConfirmed is a log parse operation binding the contract event 0xf8bef3a6fe3b4c932b5b51c6472a89f171d039f4bacf18cff632208938bf0426.
//
// Solidity: event UndelegateConfirmed(address indexed delegator, address indexed validator, uint256 amount)
func (_TaraxaDposInterface *TaraxaDposInterfaceFilterer) ParseUndelegateConfirmed(log types.Log) (*TaraxaDposInterfaceUndelegateConfirmed, error) {
	event := new(TaraxaDposInterfaceUndelegateConfirmed)
	if err := _TaraxaDposInterface.contract.UnpackLog(event, "UndelegateConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaraxaDposInterfaceUndelegatedIterator is returned from FilterUndelegated and is used to iterate over the raw logs and unpacked data for Undelegated events raised by the TaraxaDposInterface contract.
type TaraxaDposInterfaceUndelegatedIterator struct {
	Event *TaraxaDposInterfaceUndelegated // Event containing the contract specifics and raw log

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
func (it *TaraxaDposInterfaceUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraxaDposInterfaceUndelegated)
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
		it.Event = new(TaraxaDposInterfaceUndelegated)
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
func (it *TaraxaDposInterfaceUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraxaDposInterfaceUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraxaDposInterfaceUndelegated represents a Undelegated event raised by the TaraxaDposInterface contract.
type TaraxaDposInterfaceUndelegated struct {
	Delegator common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUndelegated is a free log retrieval operation binding the contract event 0x4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c.
//
// Solidity: event Undelegated(address indexed delegator, address indexed validator, uint256 amount)
func (_TaraxaDposInterface *TaraxaDposInterfaceFilterer) FilterUndelegated(opts *bind.FilterOpts, delegator []common.Address, validator []common.Address) (*TaraxaDposInterfaceUndelegatedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposInterface.contract.FilterLogs(opts, "Undelegated", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &TaraxaDposInterfaceUndelegatedIterator{contract: _TaraxaDposInterface.contract, event: "Undelegated", logs: logs, sub: sub}, nil
}

// WatchUndelegated is a free log subscription operation binding the contract event 0x4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c.
//
// Solidity: event Undelegated(address indexed delegator, address indexed validator, uint256 amount)
func (_TaraxaDposInterface *TaraxaDposInterfaceFilterer) WatchUndelegated(opts *bind.WatchOpts, sink chan<- *TaraxaDposInterfaceUndelegated, delegator []common.Address, validator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposInterface.contract.WatchLogs(opts, "Undelegated", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraxaDposInterfaceUndelegated)
				if err := _TaraxaDposInterface.contract.UnpackLog(event, "Undelegated", log); err != nil {
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

// ParseUndelegated is a log parse operation binding the contract event 0x4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c.
//
// Solidity: event Undelegated(address indexed delegator, address indexed validator, uint256 amount)
func (_TaraxaDposInterface *TaraxaDposInterfaceFilterer) ParseUndelegated(log types.Log) (*TaraxaDposInterfaceUndelegated, error) {
	event := new(TaraxaDposInterfaceUndelegated)
	if err := _TaraxaDposInterface.contract.UnpackLog(event, "Undelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaraxaDposInterfaceValidatorInfoSetIterator is returned from FilterValidatorInfoSet and is used to iterate over the raw logs and unpacked data for ValidatorInfoSet events raised by the TaraxaDposInterface contract.
type TaraxaDposInterfaceValidatorInfoSetIterator struct {
	Event *TaraxaDposInterfaceValidatorInfoSet // Event containing the contract specifics and raw log

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
func (it *TaraxaDposInterfaceValidatorInfoSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraxaDposInterfaceValidatorInfoSet)
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
		it.Event = new(TaraxaDposInterfaceValidatorInfoSet)
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
func (it *TaraxaDposInterfaceValidatorInfoSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraxaDposInterfaceValidatorInfoSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraxaDposInterfaceValidatorInfoSet represents a ValidatorInfoSet event raised by the TaraxaDposInterface contract.
type TaraxaDposInterfaceValidatorInfoSet struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorInfoSet is a free log retrieval operation binding the contract event 0x7aa20e1f59764c9066578febd688a51375adbd654aff86cef56593a17a99071d.
//
// Solidity: event ValidatorInfoSet(address indexed validator)
func (_TaraxaDposInterface *TaraxaDposInterfaceFilterer) FilterValidatorInfoSet(opts *bind.FilterOpts, validator []common.Address) (*TaraxaDposInterfaceValidatorInfoSetIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposInterface.contract.FilterLogs(opts, "ValidatorInfoSet", validatorRule)
	if err != nil {
		return nil, err
	}
	return &TaraxaDposInterfaceValidatorInfoSetIterator{contract: _TaraxaDposInterface.contract, event: "ValidatorInfoSet", logs: logs, sub: sub}, nil
}

// WatchValidatorInfoSet is a free log subscription operation binding the contract event 0x7aa20e1f59764c9066578febd688a51375adbd654aff86cef56593a17a99071d.
//
// Solidity: event ValidatorInfoSet(address indexed validator)
func (_TaraxaDposInterface *TaraxaDposInterfaceFilterer) WatchValidatorInfoSet(opts *bind.WatchOpts, sink chan<- *TaraxaDposInterfaceValidatorInfoSet, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposInterface.contract.WatchLogs(opts, "ValidatorInfoSet", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraxaDposInterfaceValidatorInfoSet)
				if err := _TaraxaDposInterface.contract.UnpackLog(event, "ValidatorInfoSet", log); err != nil {
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

// ParseValidatorInfoSet is a log parse operation binding the contract event 0x7aa20e1f59764c9066578febd688a51375adbd654aff86cef56593a17a99071d.
//
// Solidity: event ValidatorInfoSet(address indexed validator)
func (_TaraxaDposInterface *TaraxaDposInterfaceFilterer) ParseValidatorInfoSet(log types.Log) (*TaraxaDposInterfaceValidatorInfoSet, error) {
	event := new(TaraxaDposInterfaceValidatorInfoSet)
	if err := _TaraxaDposInterface.contract.UnpackLog(event, "ValidatorInfoSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaraxaDposInterfaceValidatorRegisteredIterator is returned from FilterValidatorRegistered and is used to iterate over the raw logs and unpacked data for ValidatorRegistered events raised by the TaraxaDposInterface contract.
type TaraxaDposInterfaceValidatorRegisteredIterator struct {
	Event *TaraxaDposInterfaceValidatorRegistered // Event containing the contract specifics and raw log

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
func (it *TaraxaDposInterfaceValidatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaraxaDposInterfaceValidatorRegistered)
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
		it.Event = new(TaraxaDposInterfaceValidatorRegistered)
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
func (it *TaraxaDposInterfaceValidatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaraxaDposInterfaceValidatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaraxaDposInterfaceValidatorRegistered represents a ValidatorRegistered event raised by the TaraxaDposInterface contract.
type TaraxaDposInterfaceValidatorRegistered struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorRegistered is a free log retrieval operation binding the contract event 0xd09501348473474a20c772c79c653e1fd7e8b437e418fe235d277d2c88853251.
//
// Solidity: event ValidatorRegistered(address indexed validator)
func (_TaraxaDposInterface *TaraxaDposInterfaceFilterer) FilterValidatorRegistered(opts *bind.FilterOpts, validator []common.Address) (*TaraxaDposInterfaceValidatorRegisteredIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposInterface.contract.FilterLogs(opts, "ValidatorRegistered", validatorRule)
	if err != nil {
		return nil, err
	}
	return &TaraxaDposInterfaceValidatorRegisteredIterator{contract: _TaraxaDposInterface.contract, event: "ValidatorRegistered", logs: logs, sub: sub}, nil
}

// WatchValidatorRegistered is a free log subscription operation binding the contract event 0xd09501348473474a20c772c79c653e1fd7e8b437e418fe235d277d2c88853251.
//
// Solidity: event ValidatorRegistered(address indexed validator)
func (_TaraxaDposInterface *TaraxaDposInterfaceFilterer) WatchValidatorRegistered(opts *bind.WatchOpts, sink chan<- *TaraxaDposInterfaceValidatorRegistered, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _TaraxaDposInterface.contract.WatchLogs(opts, "ValidatorRegistered", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaraxaDposInterfaceValidatorRegistered)
				if err := _TaraxaDposInterface.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
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

// ParseValidatorRegistered is a log parse operation binding the contract event 0xd09501348473474a20c772c79c653e1fd7e8b437e418fe235d277d2c88853251.
//
// Solidity: event ValidatorRegistered(address indexed validator)
func (_TaraxaDposInterface *TaraxaDposInterfaceFilterer) ParseValidatorRegistered(log types.Log) (*TaraxaDposInterfaceValidatorRegistered, error) {
	event := new(TaraxaDposInterfaceValidatorRegistered)
	if err := _TaraxaDposInterface.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
