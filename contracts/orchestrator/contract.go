// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package orchestrator

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

// OrchestratorTypesOperatorInfo is an auto generated low-level Go binding around an user-defined struct.
type OrchestratorTypesOperatorInfo struct {
	Owner    common.Address
	License  *big.Int
	Activity *big.Int
}

// OrchestratorMetaData contains all meta data concerning the Orchestrator contract.
var OrchestratorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInput\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidOperator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LicenseAlreadyDelegated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"EpochReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"licenseId\",\"type\":\"uint256\"}],\"name\":\"LicenseDelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"licenseId\",\"type\":\"uint256\"}],\"name\":\"LicenseUndelegated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"licenseId\",\"type\":\"uint256\"}],\"name\":\"Reported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_licenseId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"delegateLicense\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"epochToEpochInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalActivity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getDelegatedLicensesForOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_licenseId\",\"type\":\"uint256\"}],\"name\":\"getOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"getOperatorInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"license\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activity\",\"type\":\"uint256\"}],\"internalType\":\"structOrchestratorTypes.OperatorInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getOperatorsForOwner\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"license\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"getSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nftLicense\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"licenseId\",\"type\":\"uint256\"}],\"name\":\"licenseToLicenseInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nftLicense\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"operatorToOperatorInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"license\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"present\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"releaseEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"licenseId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"reportedInEpoch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_licenseId\",\"type\":\"uint256\"}],\"name\":\"undelegateLicense\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"signatureHash\",\"type\":\"bytes32\"}],\"name\":\"usedSignatures\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"used\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// OrchestratorABI is the input ABI used to generate the binding from.
// Deprecated: Use OrchestratorMetaData.ABI instead.
var OrchestratorABI = OrchestratorMetaData.ABI

// Orchestrator is an auto generated Go binding around an Ethereum contract.
type Orchestrator struct {
	OrchestratorCaller     // Read-only binding to the contract
	OrchestratorTransactor // Write-only binding to the contract
	OrchestratorFilterer   // Log filterer for contract events
}

// OrchestratorCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrchestratorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrchestratorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrchestratorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrchestratorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrchestratorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrchestratorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrchestratorSession struct {
	Contract     *Orchestrator     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrchestratorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrchestratorCallerSession struct {
	Contract *OrchestratorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// OrchestratorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrchestratorTransactorSession struct {
	Contract     *OrchestratorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// OrchestratorRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrchestratorRaw struct {
	Contract *Orchestrator // Generic contract binding to access the raw methods on
}

// OrchestratorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrchestratorCallerRaw struct {
	Contract *OrchestratorCaller // Generic read-only contract binding to access the raw methods on
}

// OrchestratorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrchestratorTransactorRaw struct {
	Contract *OrchestratorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrchestrator creates a new instance of Orchestrator, bound to a specific deployed contract.
func NewOrchestrator(address common.Address, backend bind.ContractBackend) (*Orchestrator, error) {
	contract, err := bindOrchestrator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Orchestrator{OrchestratorCaller: OrchestratorCaller{contract: contract}, OrchestratorTransactor: OrchestratorTransactor{contract: contract}, OrchestratorFilterer: OrchestratorFilterer{contract: contract}}, nil
}

// NewOrchestratorCaller creates a new read-only instance of Orchestrator, bound to a specific deployed contract.
func NewOrchestratorCaller(address common.Address, caller bind.ContractCaller) (*OrchestratorCaller, error) {
	contract, err := bindOrchestrator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrchestratorCaller{contract: contract}, nil
}

// NewOrchestratorTransactor creates a new write-only instance of Orchestrator, bound to a specific deployed contract.
func NewOrchestratorTransactor(address common.Address, transactor bind.ContractTransactor) (*OrchestratorTransactor, error) {
	contract, err := bindOrchestrator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrchestratorTransactor{contract: contract}, nil
}

// NewOrchestratorFilterer creates a new log filterer instance of Orchestrator, bound to a specific deployed contract.
func NewOrchestratorFilterer(address common.Address, filterer bind.ContractFilterer) (*OrchestratorFilterer, error) {
	contract, err := bindOrchestrator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrchestratorFilterer{contract: contract}, nil
}

// bindOrchestrator binds a generic wrapper to an already deployed contract.
func bindOrchestrator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OrchestratorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Orchestrator *OrchestratorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Orchestrator.Contract.OrchestratorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Orchestrator *OrchestratorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Orchestrator.Contract.OrchestratorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Orchestrator *OrchestratorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Orchestrator.Contract.OrchestratorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Orchestrator *OrchestratorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Orchestrator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Orchestrator *OrchestratorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Orchestrator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Orchestrator *OrchestratorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Orchestrator.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Orchestrator *OrchestratorCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Orchestrator.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Orchestrator *OrchestratorSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Orchestrator.Contract.DEFAULTADMINROLE(&_Orchestrator.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Orchestrator *OrchestratorCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Orchestrator.Contract.DEFAULTADMINROLE(&_Orchestrator.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Orchestrator *OrchestratorCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Orchestrator.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Orchestrator *OrchestratorSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Orchestrator.Contract.UPGRADEINTERFACEVERSION(&_Orchestrator.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Orchestrator *OrchestratorCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Orchestrator.Contract.UPGRADEINTERFACEVERSION(&_Orchestrator.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Orchestrator *OrchestratorCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Orchestrator.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Orchestrator *OrchestratorSession) CurrentEpoch() (*big.Int, error) {
	return _Orchestrator.Contract.CurrentEpoch(&_Orchestrator.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_Orchestrator *OrchestratorCallerSession) CurrentEpoch() (*big.Int, error) {
	return _Orchestrator.Contract.CurrentEpoch(&_Orchestrator.CallOpts)
}

// EpochToEpochInfo is a free data retrieval call binding the contract method 0x17a8463d.
//
// Solidity: function epochToEpochInfo(uint256 epochId) view returns(uint256 epochId, uint256 timestamp, uint256 totalActivity)
func (_Orchestrator *OrchestratorCaller) EpochToEpochInfo(opts *bind.CallOpts, epochId *big.Int) (struct {
	EpochId       *big.Int
	Timestamp     *big.Int
	TotalActivity *big.Int
}, error) {
	var out []interface{}
	err := _Orchestrator.contract.Call(opts, &out, "epochToEpochInfo", epochId)

	outstruct := new(struct {
		EpochId       *big.Int
		Timestamp     *big.Int
		TotalActivity *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EpochId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalActivity = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EpochToEpochInfo is a free data retrieval call binding the contract method 0x17a8463d.
//
// Solidity: function epochToEpochInfo(uint256 epochId) view returns(uint256 epochId, uint256 timestamp, uint256 totalActivity)
func (_Orchestrator *OrchestratorSession) EpochToEpochInfo(epochId *big.Int) (struct {
	EpochId       *big.Int
	Timestamp     *big.Int
	TotalActivity *big.Int
}, error) {
	return _Orchestrator.Contract.EpochToEpochInfo(&_Orchestrator.CallOpts, epochId)
}

// EpochToEpochInfo is a free data retrieval call binding the contract method 0x17a8463d.
//
// Solidity: function epochToEpochInfo(uint256 epochId) view returns(uint256 epochId, uint256 timestamp, uint256 totalActivity)
func (_Orchestrator *OrchestratorCallerSession) EpochToEpochInfo(epochId *big.Int) (struct {
	EpochId       *big.Int
	Timestamp     *big.Int
	TotalActivity *big.Int
}, error) {
	return _Orchestrator.Contract.EpochToEpochInfo(&_Orchestrator.CallOpts, epochId)
}

// GetDelegatedLicensesForOwner is a free data retrieval call binding the contract method 0x4d6fc9e7.
//
// Solidity: function getDelegatedLicensesForOwner(address _owner) view returns(uint256[])
func (_Orchestrator *OrchestratorCaller) GetDelegatedLicensesForOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Orchestrator.contract.Call(opts, &out, "getDelegatedLicensesForOwner", _owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetDelegatedLicensesForOwner is a free data retrieval call binding the contract method 0x4d6fc9e7.
//
// Solidity: function getDelegatedLicensesForOwner(address _owner) view returns(uint256[])
func (_Orchestrator *OrchestratorSession) GetDelegatedLicensesForOwner(_owner common.Address) ([]*big.Int, error) {
	return _Orchestrator.Contract.GetDelegatedLicensesForOwner(&_Orchestrator.CallOpts, _owner)
}

// GetDelegatedLicensesForOwner is a free data retrieval call binding the contract method 0x4d6fc9e7.
//
// Solidity: function getDelegatedLicensesForOwner(address _owner) view returns(uint256[])
func (_Orchestrator *OrchestratorCallerSession) GetDelegatedLicensesForOwner(_owner common.Address) ([]*big.Int, error) {
	return _Orchestrator.Contract.GetDelegatedLicensesForOwner(&_Orchestrator.CallOpts, _owner)
}

// GetOperator is a free data retrieval call binding the contract method 0x05f63c8a.
//
// Solidity: function getOperator(uint256 _licenseId) view returns(address)
func (_Orchestrator *OrchestratorCaller) GetOperator(opts *bind.CallOpts, _licenseId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Orchestrator.contract.Call(opts, &out, "getOperator", _licenseId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOperator is a free data retrieval call binding the contract method 0x05f63c8a.
//
// Solidity: function getOperator(uint256 _licenseId) view returns(address)
func (_Orchestrator *OrchestratorSession) GetOperator(_licenseId *big.Int) (common.Address, error) {
	return _Orchestrator.Contract.GetOperator(&_Orchestrator.CallOpts, _licenseId)
}

// GetOperator is a free data retrieval call binding the contract method 0x05f63c8a.
//
// Solidity: function getOperator(uint256 _licenseId) view returns(address)
func (_Orchestrator *OrchestratorCallerSession) GetOperator(_licenseId *big.Int) (common.Address, error) {
	return _Orchestrator.Contract.GetOperator(&_Orchestrator.CallOpts, _licenseId)
}

// GetOperatorInfo is a free data retrieval call binding the contract method 0x27d9ab5d.
//
// Solidity: function getOperatorInfo(address _operator) view returns((address,uint256,uint256))
func (_Orchestrator *OrchestratorCaller) GetOperatorInfo(opts *bind.CallOpts, _operator common.Address) (OrchestratorTypesOperatorInfo, error) {
	var out []interface{}
	err := _Orchestrator.contract.Call(opts, &out, "getOperatorInfo", _operator)

	if err != nil {
		return *new(OrchestratorTypesOperatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(OrchestratorTypesOperatorInfo)).(*OrchestratorTypesOperatorInfo)

	return out0, err

}

// GetOperatorInfo is a free data retrieval call binding the contract method 0x27d9ab5d.
//
// Solidity: function getOperatorInfo(address _operator) view returns((address,uint256,uint256))
func (_Orchestrator *OrchestratorSession) GetOperatorInfo(_operator common.Address) (OrchestratorTypesOperatorInfo, error) {
	return _Orchestrator.Contract.GetOperatorInfo(&_Orchestrator.CallOpts, _operator)
}

// GetOperatorInfo is a free data retrieval call binding the contract method 0x27d9ab5d.
//
// Solidity: function getOperatorInfo(address _operator) view returns((address,uint256,uint256))
func (_Orchestrator *OrchestratorCallerSession) GetOperatorInfo(_operator common.Address) (OrchestratorTypesOperatorInfo, error) {
	return _Orchestrator.Contract.GetOperatorInfo(&_Orchestrator.CallOpts, _operator)
}

// GetOperatorsForOwner is a free data retrieval call binding the contract method 0x4720dae2.
//
// Solidity: function getOperatorsForOwner(address _owner) view returns(address[])
func (_Orchestrator *OrchestratorCaller) GetOperatorsForOwner(opts *bind.CallOpts, _owner common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Orchestrator.contract.Call(opts, &out, "getOperatorsForOwner", _owner)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOperatorsForOwner is a free data retrieval call binding the contract method 0x4720dae2.
//
// Solidity: function getOperatorsForOwner(address _owner) view returns(address[])
func (_Orchestrator *OrchestratorSession) GetOperatorsForOwner(_owner common.Address) ([]common.Address, error) {
	return _Orchestrator.Contract.GetOperatorsForOwner(&_Orchestrator.CallOpts, _owner)
}

// GetOperatorsForOwner is a free data retrieval call binding the contract method 0x4720dae2.
//
// Solidity: function getOperatorsForOwner(address _owner) view returns(address[])
func (_Orchestrator *OrchestratorCallerSession) GetOperatorsForOwner(_owner common.Address) ([]common.Address, error) {
	return _Orchestrator.Contract.GetOperatorsForOwner(&_Orchestrator.CallOpts, _owner)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Orchestrator *OrchestratorCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Orchestrator.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Orchestrator *OrchestratorSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Orchestrator.Contract.GetRoleAdmin(&_Orchestrator.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Orchestrator *OrchestratorCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Orchestrator.Contract.GetRoleAdmin(&_Orchestrator.CallOpts, role)
}

// GetSigner is a free data retrieval call binding the contract method 0x5858aa26.
//
// Solidity: function getSigner(address user, uint256 license, bytes signature) pure returns(address)
func (_Orchestrator *OrchestratorCaller) GetSigner(opts *bind.CallOpts, user common.Address, license *big.Int, signature []byte) (common.Address, error) {
	var out []interface{}
	err := _Orchestrator.contract.Call(opts, &out, "getSigner", user, license, signature)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSigner is a free data retrieval call binding the contract method 0x5858aa26.
//
// Solidity: function getSigner(address user, uint256 license, bytes signature) pure returns(address)
func (_Orchestrator *OrchestratorSession) GetSigner(user common.Address, license *big.Int, signature []byte) (common.Address, error) {
	return _Orchestrator.Contract.GetSigner(&_Orchestrator.CallOpts, user, license, signature)
}

// GetSigner is a free data retrieval call binding the contract method 0x5858aa26.
//
// Solidity: function getSigner(address user, uint256 license, bytes signature) pure returns(address)
func (_Orchestrator *OrchestratorCallerSession) GetSigner(user common.Address, license *big.Int, signature []byte) (common.Address, error) {
	return _Orchestrator.Contract.GetSigner(&_Orchestrator.CallOpts, user, license, signature)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Orchestrator *OrchestratorCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Orchestrator.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Orchestrator *OrchestratorSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Orchestrator.Contract.HasRole(&_Orchestrator.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Orchestrator *OrchestratorCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Orchestrator.Contract.HasRole(&_Orchestrator.CallOpts, role, account)
}

// LicenseToLicenseInfo is a free data retrieval call binding the contract method 0xf9dceec4.
//
// Solidity: function licenseToLicenseInfo(uint256 licenseId) view returns(address owner, address operator)
func (_Orchestrator *OrchestratorCaller) LicenseToLicenseInfo(opts *bind.CallOpts, licenseId *big.Int) (struct {
	Owner    common.Address
	Operator common.Address
}, error) {
	var out []interface{}
	err := _Orchestrator.contract.Call(opts, &out, "licenseToLicenseInfo", licenseId)

	outstruct := new(struct {
		Owner    common.Address
		Operator common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Operator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// LicenseToLicenseInfo is a free data retrieval call binding the contract method 0xf9dceec4.
//
// Solidity: function licenseToLicenseInfo(uint256 licenseId) view returns(address owner, address operator)
func (_Orchestrator *OrchestratorSession) LicenseToLicenseInfo(licenseId *big.Int) (struct {
	Owner    common.Address
	Operator common.Address
}, error) {
	return _Orchestrator.Contract.LicenseToLicenseInfo(&_Orchestrator.CallOpts, licenseId)
}

// LicenseToLicenseInfo is a free data retrieval call binding the contract method 0xf9dceec4.
//
// Solidity: function licenseToLicenseInfo(uint256 licenseId) view returns(address owner, address operator)
func (_Orchestrator *OrchestratorCallerSession) LicenseToLicenseInfo(licenseId *big.Int) (struct {
	Owner    common.Address
	Operator common.Address
}, error) {
	return _Orchestrator.Contract.LicenseToLicenseInfo(&_Orchestrator.CallOpts, licenseId)
}

// NftLicense is a free data retrieval call binding the contract method 0xc3f07cbf.
//
// Solidity: function nftLicense() view returns(address)
func (_Orchestrator *OrchestratorCaller) NftLicense(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Orchestrator.contract.Call(opts, &out, "nftLicense")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftLicense is a free data retrieval call binding the contract method 0xc3f07cbf.
//
// Solidity: function nftLicense() view returns(address)
func (_Orchestrator *OrchestratorSession) NftLicense() (common.Address, error) {
	return _Orchestrator.Contract.NftLicense(&_Orchestrator.CallOpts)
}

// NftLicense is a free data retrieval call binding the contract method 0xc3f07cbf.
//
// Solidity: function nftLicense() view returns(address)
func (_Orchestrator *OrchestratorCallerSession) NftLicense() (common.Address, error) {
	return _Orchestrator.Contract.NftLicense(&_Orchestrator.CallOpts)
}

// OperatorToOperatorInfo is a free data retrieval call binding the contract method 0x6f39f71c.
//
// Solidity: function operatorToOperatorInfo(address operator) view returns(address owner, uint256 license, uint256 activity)
func (_Orchestrator *OrchestratorCaller) OperatorToOperatorInfo(opts *bind.CallOpts, operator common.Address) (struct {
	Owner    common.Address
	License  *big.Int
	Activity *big.Int
}, error) {
	var out []interface{}
	err := _Orchestrator.contract.Call(opts, &out, "operatorToOperatorInfo", operator)

	outstruct := new(struct {
		Owner    common.Address
		License  *big.Int
		Activity *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.License = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Activity = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OperatorToOperatorInfo is a free data retrieval call binding the contract method 0x6f39f71c.
//
// Solidity: function operatorToOperatorInfo(address operator) view returns(address owner, uint256 license, uint256 activity)
func (_Orchestrator *OrchestratorSession) OperatorToOperatorInfo(operator common.Address) (struct {
	Owner    common.Address
	License  *big.Int
	Activity *big.Int
}, error) {
	return _Orchestrator.Contract.OperatorToOperatorInfo(&_Orchestrator.CallOpts, operator)
}

// OperatorToOperatorInfo is a free data retrieval call binding the contract method 0x6f39f71c.
//
// Solidity: function operatorToOperatorInfo(address operator) view returns(address owner, uint256 license, uint256 activity)
func (_Orchestrator *OrchestratorCallerSession) OperatorToOperatorInfo(operator common.Address) (struct {
	Owner    common.Address
	License  *big.Int
	Activity *big.Int
}, error) {
	return _Orchestrator.Contract.OperatorToOperatorInfo(&_Orchestrator.CallOpts, operator)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Orchestrator *OrchestratorCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Orchestrator.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Orchestrator *OrchestratorSession) ProxiableUUID() ([32]byte, error) {
	return _Orchestrator.Contract.ProxiableUUID(&_Orchestrator.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Orchestrator *OrchestratorCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Orchestrator.Contract.ProxiableUUID(&_Orchestrator.CallOpts)
}

// ReportedInEpoch is a free data retrieval call binding the contract method 0xd0f6f594.
//
// Solidity: function reportedInEpoch(uint256 licenseId, uint256 epochId) view returns(bool)
func (_Orchestrator *OrchestratorCaller) ReportedInEpoch(opts *bind.CallOpts, licenseId *big.Int, epochId *big.Int) (bool, error) {
	var out []interface{}
	err := _Orchestrator.contract.Call(opts, &out, "reportedInEpoch", licenseId, epochId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ReportedInEpoch is a free data retrieval call binding the contract method 0xd0f6f594.
//
// Solidity: function reportedInEpoch(uint256 licenseId, uint256 epochId) view returns(bool)
func (_Orchestrator *OrchestratorSession) ReportedInEpoch(licenseId *big.Int, epochId *big.Int) (bool, error) {
	return _Orchestrator.Contract.ReportedInEpoch(&_Orchestrator.CallOpts, licenseId, epochId)
}

// ReportedInEpoch is a free data retrieval call binding the contract method 0xd0f6f594.
//
// Solidity: function reportedInEpoch(uint256 licenseId, uint256 epochId) view returns(bool)
func (_Orchestrator *OrchestratorCallerSession) ReportedInEpoch(licenseId *big.Int, epochId *big.Int) (bool, error) {
	return _Orchestrator.Contract.ReportedInEpoch(&_Orchestrator.CallOpts, licenseId, epochId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Orchestrator *OrchestratorCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Orchestrator.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Orchestrator *OrchestratorSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Orchestrator.Contract.SupportsInterface(&_Orchestrator.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Orchestrator *OrchestratorCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Orchestrator.Contract.SupportsInterface(&_Orchestrator.CallOpts, interfaceId)
}

// UsedSignatures is a free data retrieval call binding the contract method 0xf978fd61.
//
// Solidity: function usedSignatures(bytes32 signatureHash) view returns(bool used)
func (_Orchestrator *OrchestratorCaller) UsedSignatures(opts *bind.CallOpts, signatureHash [32]byte) (bool, error) {
	var out []interface{}
	err := _Orchestrator.contract.Call(opts, &out, "usedSignatures", signatureHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedSignatures is a free data retrieval call binding the contract method 0xf978fd61.
//
// Solidity: function usedSignatures(bytes32 signatureHash) view returns(bool used)
func (_Orchestrator *OrchestratorSession) UsedSignatures(signatureHash [32]byte) (bool, error) {
	return _Orchestrator.Contract.UsedSignatures(&_Orchestrator.CallOpts, signatureHash)
}

// UsedSignatures is a free data retrieval call binding the contract method 0xf978fd61.
//
// Solidity: function usedSignatures(bytes32 signatureHash) view returns(bool used)
func (_Orchestrator *OrchestratorCallerSession) UsedSignatures(signatureHash [32]byte) (bool, error) {
	return _Orchestrator.Contract.UsedSignatures(&_Orchestrator.CallOpts, signatureHash)
}

// DelegateLicense is a paid mutator transaction binding the contract method 0x1f49e345.
//
// Solidity: function delegateLicense(address _operator, uint256 _licenseId, bytes signature) returns()
func (_Orchestrator *OrchestratorTransactor) DelegateLicense(opts *bind.TransactOpts, _operator common.Address, _licenseId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Orchestrator.contract.Transact(opts, "delegateLicense", _operator, _licenseId, signature)
}

// DelegateLicense is a paid mutator transaction binding the contract method 0x1f49e345.
//
// Solidity: function delegateLicense(address _operator, uint256 _licenseId, bytes signature) returns()
func (_Orchestrator *OrchestratorSession) DelegateLicense(_operator common.Address, _licenseId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Orchestrator.Contract.DelegateLicense(&_Orchestrator.TransactOpts, _operator, _licenseId, signature)
}

// DelegateLicense is a paid mutator transaction binding the contract method 0x1f49e345.
//
// Solidity: function delegateLicense(address _operator, uint256 _licenseId, bytes signature) returns()
func (_Orchestrator *OrchestratorTransactorSession) DelegateLicense(_operator common.Address, _licenseId *big.Int, signature []byte) (*types.Transaction, error) {
	return _Orchestrator.Contract.DelegateLicense(&_Orchestrator.TransactOpts, _operator, _licenseId, signature)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Orchestrator *OrchestratorTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Orchestrator.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Orchestrator *OrchestratorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Orchestrator.Contract.GrantRole(&_Orchestrator.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Orchestrator *OrchestratorTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Orchestrator.Contract.GrantRole(&_Orchestrator.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address admin, address _nftLicense) returns()
func (_Orchestrator *OrchestratorTransactor) Initialize(opts *bind.TransactOpts, admin common.Address, _nftLicense common.Address) (*types.Transaction, error) {
	return _Orchestrator.contract.Transact(opts, "initialize", admin, _nftLicense)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address admin, address _nftLicense) returns()
func (_Orchestrator *OrchestratorSession) Initialize(admin common.Address, _nftLicense common.Address) (*types.Transaction, error) {
	return _Orchestrator.Contract.Initialize(&_Orchestrator.TransactOpts, admin, _nftLicense)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address admin, address _nftLicense) returns()
func (_Orchestrator *OrchestratorTransactorSession) Initialize(admin common.Address, _nftLicense common.Address) (*types.Transaction, error) {
	return _Orchestrator.Contract.Initialize(&_Orchestrator.TransactOpts, admin, _nftLicense)
}

// Present is a paid mutator transaction binding the contract method 0x9d7b745e.
//
// Solidity: function present(uint256 epochId) returns()
func (_Orchestrator *OrchestratorTransactor) Present(opts *bind.TransactOpts, epochId *big.Int) (*types.Transaction, error) {
	return _Orchestrator.contract.Transact(opts, "present", epochId)
}

// Present is a paid mutator transaction binding the contract method 0x9d7b745e.
//
// Solidity: function present(uint256 epochId) returns()
func (_Orchestrator *OrchestratorSession) Present(epochId *big.Int) (*types.Transaction, error) {
	return _Orchestrator.Contract.Present(&_Orchestrator.TransactOpts, epochId)
}

// Present is a paid mutator transaction binding the contract method 0x9d7b745e.
//
// Solidity: function present(uint256 epochId) returns()
func (_Orchestrator *OrchestratorTransactorSession) Present(epochId *big.Int) (*types.Transaction, error) {
	return _Orchestrator.Contract.Present(&_Orchestrator.TransactOpts, epochId)
}

// ReleaseEpoch is a paid mutator transaction binding the contract method 0xc6503de1.
//
// Solidity: function releaseEpoch(uint256 epochId) returns()
func (_Orchestrator *OrchestratorTransactor) ReleaseEpoch(opts *bind.TransactOpts, epochId *big.Int) (*types.Transaction, error) {
	return _Orchestrator.contract.Transact(opts, "releaseEpoch", epochId)
}

// ReleaseEpoch is a paid mutator transaction binding the contract method 0xc6503de1.
//
// Solidity: function releaseEpoch(uint256 epochId) returns()
func (_Orchestrator *OrchestratorSession) ReleaseEpoch(epochId *big.Int) (*types.Transaction, error) {
	return _Orchestrator.Contract.ReleaseEpoch(&_Orchestrator.TransactOpts, epochId)
}

// ReleaseEpoch is a paid mutator transaction binding the contract method 0xc6503de1.
//
// Solidity: function releaseEpoch(uint256 epochId) returns()
func (_Orchestrator *OrchestratorTransactorSession) ReleaseEpoch(epochId *big.Int) (*types.Transaction, error) {
	return _Orchestrator.Contract.ReleaseEpoch(&_Orchestrator.TransactOpts, epochId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Orchestrator *OrchestratorTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Orchestrator.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Orchestrator *OrchestratorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Orchestrator.Contract.RenounceRole(&_Orchestrator.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Orchestrator *OrchestratorTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Orchestrator.Contract.RenounceRole(&_Orchestrator.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Orchestrator *OrchestratorTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Orchestrator.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Orchestrator *OrchestratorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Orchestrator.Contract.RevokeRole(&_Orchestrator.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Orchestrator *OrchestratorTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Orchestrator.Contract.RevokeRole(&_Orchestrator.TransactOpts, role, account)
}

// UndelegateLicense is a paid mutator transaction binding the contract method 0xc350ef6e.
//
// Solidity: function undelegateLicense(uint256 _licenseId) returns()
func (_Orchestrator *OrchestratorTransactor) UndelegateLicense(opts *bind.TransactOpts, _licenseId *big.Int) (*types.Transaction, error) {
	return _Orchestrator.contract.Transact(opts, "undelegateLicense", _licenseId)
}

// UndelegateLicense is a paid mutator transaction binding the contract method 0xc350ef6e.
//
// Solidity: function undelegateLicense(uint256 _licenseId) returns()
func (_Orchestrator *OrchestratorSession) UndelegateLicense(_licenseId *big.Int) (*types.Transaction, error) {
	return _Orchestrator.Contract.UndelegateLicense(&_Orchestrator.TransactOpts, _licenseId)
}

// UndelegateLicense is a paid mutator transaction binding the contract method 0xc350ef6e.
//
// Solidity: function undelegateLicense(uint256 _licenseId) returns()
func (_Orchestrator *OrchestratorTransactorSession) UndelegateLicense(_licenseId *big.Int) (*types.Transaction, error) {
	return _Orchestrator.Contract.UndelegateLicense(&_Orchestrator.TransactOpts, _licenseId)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Orchestrator *OrchestratorTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Orchestrator.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Orchestrator *OrchestratorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Orchestrator.Contract.UpgradeToAndCall(&_Orchestrator.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Orchestrator *OrchestratorTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Orchestrator.Contract.UpgradeToAndCall(&_Orchestrator.TransactOpts, newImplementation, data)
}

// OrchestratorEpochReleasedIterator is returned from FilterEpochReleased and is used to iterate over the raw logs and unpacked data for EpochReleased events raised by the Orchestrator contract.
type OrchestratorEpochReleasedIterator struct {
	Event *OrchestratorEpochReleased // Event containing the contract specifics and raw log

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
func (it *OrchestratorEpochReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrchestratorEpochReleased)
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
		it.Event = new(OrchestratorEpochReleased)
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
func (it *OrchestratorEpochReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrchestratorEpochReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrchestratorEpochReleased represents a EpochReleased event raised by the Orchestrator contract.
type OrchestratorEpochReleased struct {
	Epoch     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEpochReleased is a free log retrieval operation binding the contract event 0xc8b3799b2c521d4290e6122fa5a50c584e7e42e4b84547efaf7e9d86903c0d22.
//
// Solidity: event EpochReleased(uint256 epoch, uint256 timestamp)
func (_Orchestrator *OrchestratorFilterer) FilterEpochReleased(opts *bind.FilterOpts) (*OrchestratorEpochReleasedIterator, error) {

	logs, sub, err := _Orchestrator.contract.FilterLogs(opts, "EpochReleased")
	if err != nil {
		return nil, err
	}
	return &OrchestratorEpochReleasedIterator{contract: _Orchestrator.contract, event: "EpochReleased", logs: logs, sub: sub}, nil
}

// WatchEpochReleased is a free log subscription operation binding the contract event 0xc8b3799b2c521d4290e6122fa5a50c584e7e42e4b84547efaf7e9d86903c0d22.
//
// Solidity: event EpochReleased(uint256 epoch, uint256 timestamp)
func (_Orchestrator *OrchestratorFilterer) WatchEpochReleased(opts *bind.WatchOpts, sink chan<- *OrchestratorEpochReleased) (event.Subscription, error) {

	logs, sub, err := _Orchestrator.contract.WatchLogs(opts, "EpochReleased")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrchestratorEpochReleased)
				if err := _Orchestrator.contract.UnpackLog(event, "EpochReleased", log); err != nil {
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

// ParseEpochReleased is a log parse operation binding the contract event 0xc8b3799b2c521d4290e6122fa5a50c584e7e42e4b84547efaf7e9d86903c0d22.
//
// Solidity: event EpochReleased(uint256 epoch, uint256 timestamp)
func (_Orchestrator *OrchestratorFilterer) ParseEpochReleased(log types.Log) (*OrchestratorEpochReleased, error) {
	event := new(OrchestratorEpochReleased)
	if err := _Orchestrator.contract.UnpackLog(event, "EpochReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrchestratorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Orchestrator contract.
type OrchestratorInitializedIterator struct {
	Event *OrchestratorInitialized // Event containing the contract specifics and raw log

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
func (it *OrchestratorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrchestratorInitialized)
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
		it.Event = new(OrchestratorInitialized)
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
func (it *OrchestratorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrchestratorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrchestratorInitialized represents a Initialized event raised by the Orchestrator contract.
type OrchestratorInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Orchestrator *OrchestratorFilterer) FilterInitialized(opts *bind.FilterOpts) (*OrchestratorInitializedIterator, error) {

	logs, sub, err := _Orchestrator.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OrchestratorInitializedIterator{contract: _Orchestrator.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Orchestrator *OrchestratorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OrchestratorInitialized) (event.Subscription, error) {

	logs, sub, err := _Orchestrator.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrchestratorInitialized)
				if err := _Orchestrator.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Orchestrator *OrchestratorFilterer) ParseInitialized(log types.Log) (*OrchestratorInitialized, error) {
	event := new(OrchestratorInitialized)
	if err := _Orchestrator.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrchestratorLicenseDelegatedIterator is returned from FilterLicenseDelegated and is used to iterate over the raw logs and unpacked data for LicenseDelegated events raised by the Orchestrator contract.
type OrchestratorLicenseDelegatedIterator struct {
	Event *OrchestratorLicenseDelegated // Event containing the contract specifics and raw log

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
func (it *OrchestratorLicenseDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrchestratorLicenseDelegated)
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
		it.Event = new(OrchestratorLicenseDelegated)
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
func (it *OrchestratorLicenseDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrchestratorLicenseDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrchestratorLicenseDelegated represents a LicenseDelegated event raised by the Orchestrator contract.
type OrchestratorLicenseDelegated struct {
	Operator  common.Address
	LicenseId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLicenseDelegated is a free log retrieval operation binding the contract event 0x8d2ebf625a050eca3f30de31e3f58e552a78b87d3dda324be0abcdcfff9c8280.
//
// Solidity: event LicenseDelegated(address operator, uint256 licenseId)
func (_Orchestrator *OrchestratorFilterer) FilterLicenseDelegated(opts *bind.FilterOpts) (*OrchestratorLicenseDelegatedIterator, error) {

	logs, sub, err := _Orchestrator.contract.FilterLogs(opts, "LicenseDelegated")
	if err != nil {
		return nil, err
	}
	return &OrchestratorLicenseDelegatedIterator{contract: _Orchestrator.contract, event: "LicenseDelegated", logs: logs, sub: sub}, nil
}

// WatchLicenseDelegated is a free log subscription operation binding the contract event 0x8d2ebf625a050eca3f30de31e3f58e552a78b87d3dda324be0abcdcfff9c8280.
//
// Solidity: event LicenseDelegated(address operator, uint256 licenseId)
func (_Orchestrator *OrchestratorFilterer) WatchLicenseDelegated(opts *bind.WatchOpts, sink chan<- *OrchestratorLicenseDelegated) (event.Subscription, error) {

	logs, sub, err := _Orchestrator.contract.WatchLogs(opts, "LicenseDelegated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrchestratorLicenseDelegated)
				if err := _Orchestrator.contract.UnpackLog(event, "LicenseDelegated", log); err != nil {
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

// ParseLicenseDelegated is a log parse operation binding the contract event 0x8d2ebf625a050eca3f30de31e3f58e552a78b87d3dda324be0abcdcfff9c8280.
//
// Solidity: event LicenseDelegated(address operator, uint256 licenseId)
func (_Orchestrator *OrchestratorFilterer) ParseLicenseDelegated(log types.Log) (*OrchestratorLicenseDelegated, error) {
	event := new(OrchestratorLicenseDelegated)
	if err := _Orchestrator.contract.UnpackLog(event, "LicenseDelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrchestratorLicenseUndelegatedIterator is returned from FilterLicenseUndelegated and is used to iterate over the raw logs and unpacked data for LicenseUndelegated events raised by the Orchestrator contract.
type OrchestratorLicenseUndelegatedIterator struct {
	Event *OrchestratorLicenseUndelegated // Event containing the contract specifics and raw log

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
func (it *OrchestratorLicenseUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrchestratorLicenseUndelegated)
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
		it.Event = new(OrchestratorLicenseUndelegated)
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
func (it *OrchestratorLicenseUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrchestratorLicenseUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrchestratorLicenseUndelegated represents a LicenseUndelegated event raised by the Orchestrator contract.
type OrchestratorLicenseUndelegated struct {
	LicenseId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLicenseUndelegated is a free log retrieval operation binding the contract event 0xc7f53507e2ed0afbf5239d59671fbef580350decb5b9d0a319cee4743e06e5ad.
//
// Solidity: event LicenseUndelegated(uint256 licenseId)
func (_Orchestrator *OrchestratorFilterer) FilterLicenseUndelegated(opts *bind.FilterOpts) (*OrchestratorLicenseUndelegatedIterator, error) {

	logs, sub, err := _Orchestrator.contract.FilterLogs(opts, "LicenseUndelegated")
	if err != nil {
		return nil, err
	}
	return &OrchestratorLicenseUndelegatedIterator{contract: _Orchestrator.contract, event: "LicenseUndelegated", logs: logs, sub: sub}, nil
}

// WatchLicenseUndelegated is a free log subscription operation binding the contract event 0xc7f53507e2ed0afbf5239d59671fbef580350decb5b9d0a319cee4743e06e5ad.
//
// Solidity: event LicenseUndelegated(uint256 licenseId)
func (_Orchestrator *OrchestratorFilterer) WatchLicenseUndelegated(opts *bind.WatchOpts, sink chan<- *OrchestratorLicenseUndelegated) (event.Subscription, error) {

	logs, sub, err := _Orchestrator.contract.WatchLogs(opts, "LicenseUndelegated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrchestratorLicenseUndelegated)
				if err := _Orchestrator.contract.UnpackLog(event, "LicenseUndelegated", log); err != nil {
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

// ParseLicenseUndelegated is a log parse operation binding the contract event 0xc7f53507e2ed0afbf5239d59671fbef580350decb5b9d0a319cee4743e06e5ad.
//
// Solidity: event LicenseUndelegated(uint256 licenseId)
func (_Orchestrator *OrchestratorFilterer) ParseLicenseUndelegated(log types.Log) (*OrchestratorLicenseUndelegated, error) {
	event := new(OrchestratorLicenseUndelegated)
	if err := _Orchestrator.contract.UnpackLog(event, "LicenseUndelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrchestratorReportedIterator is returned from FilterReported and is used to iterate over the raw logs and unpacked data for Reported events raised by the Orchestrator contract.
type OrchestratorReportedIterator struct {
	Event *OrchestratorReported // Event containing the contract specifics and raw log

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
func (it *OrchestratorReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrchestratorReported)
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
		it.Event = new(OrchestratorReported)
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
func (it *OrchestratorReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrchestratorReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrchestratorReported represents a Reported event raised by the Orchestrator contract.
type OrchestratorReported struct {
	Epoch     *big.Int
	LicenseId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReported is a free log retrieval operation binding the contract event 0x6fea36bfb4b7c912613756951fb5bf8ddaedb74de1f256d0bef81533158cda54.
//
// Solidity: event Reported(uint256 epoch, uint256 licenseId)
func (_Orchestrator *OrchestratorFilterer) FilterReported(opts *bind.FilterOpts) (*OrchestratorReportedIterator, error) {

	logs, sub, err := _Orchestrator.contract.FilterLogs(opts, "Reported")
	if err != nil {
		return nil, err
	}
	return &OrchestratorReportedIterator{contract: _Orchestrator.contract, event: "Reported", logs: logs, sub: sub}, nil
}

// WatchReported is a free log subscription operation binding the contract event 0x6fea36bfb4b7c912613756951fb5bf8ddaedb74de1f256d0bef81533158cda54.
//
// Solidity: event Reported(uint256 epoch, uint256 licenseId)
func (_Orchestrator *OrchestratorFilterer) WatchReported(opts *bind.WatchOpts, sink chan<- *OrchestratorReported) (event.Subscription, error) {

	logs, sub, err := _Orchestrator.contract.WatchLogs(opts, "Reported")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrchestratorReported)
				if err := _Orchestrator.contract.UnpackLog(event, "Reported", log); err != nil {
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

// ParseReported is a log parse operation binding the contract event 0x6fea36bfb4b7c912613756951fb5bf8ddaedb74de1f256d0bef81533158cda54.
//
// Solidity: event Reported(uint256 epoch, uint256 licenseId)
func (_Orchestrator *OrchestratorFilterer) ParseReported(log types.Log) (*OrchestratorReported, error) {
	event := new(OrchestratorReported)
	if err := _Orchestrator.contract.UnpackLog(event, "Reported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrchestratorRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Orchestrator contract.
type OrchestratorRoleAdminChangedIterator struct {
	Event *OrchestratorRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *OrchestratorRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrchestratorRoleAdminChanged)
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
		it.Event = new(OrchestratorRoleAdminChanged)
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
func (it *OrchestratorRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrchestratorRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrchestratorRoleAdminChanged represents a RoleAdminChanged event raised by the Orchestrator contract.
type OrchestratorRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Orchestrator *OrchestratorFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*OrchestratorRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Orchestrator.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &OrchestratorRoleAdminChangedIterator{contract: _Orchestrator.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Orchestrator *OrchestratorFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *OrchestratorRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Orchestrator.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrchestratorRoleAdminChanged)
				if err := _Orchestrator.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Orchestrator *OrchestratorFilterer) ParseRoleAdminChanged(log types.Log) (*OrchestratorRoleAdminChanged, error) {
	event := new(OrchestratorRoleAdminChanged)
	if err := _Orchestrator.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrchestratorRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Orchestrator contract.
type OrchestratorRoleGrantedIterator struct {
	Event *OrchestratorRoleGranted // Event containing the contract specifics and raw log

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
func (it *OrchestratorRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrchestratorRoleGranted)
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
		it.Event = new(OrchestratorRoleGranted)
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
func (it *OrchestratorRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrchestratorRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrchestratorRoleGranted represents a RoleGranted event raised by the Orchestrator contract.
type OrchestratorRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Orchestrator *OrchestratorFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*OrchestratorRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Orchestrator.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &OrchestratorRoleGrantedIterator{contract: _Orchestrator.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Orchestrator *OrchestratorFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *OrchestratorRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Orchestrator.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrchestratorRoleGranted)
				if err := _Orchestrator.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Orchestrator *OrchestratorFilterer) ParseRoleGranted(log types.Log) (*OrchestratorRoleGranted, error) {
	event := new(OrchestratorRoleGranted)
	if err := _Orchestrator.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrchestratorRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Orchestrator contract.
type OrchestratorRoleRevokedIterator struct {
	Event *OrchestratorRoleRevoked // Event containing the contract specifics and raw log

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
func (it *OrchestratorRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrchestratorRoleRevoked)
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
		it.Event = new(OrchestratorRoleRevoked)
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
func (it *OrchestratorRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrchestratorRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrchestratorRoleRevoked represents a RoleRevoked event raised by the Orchestrator contract.
type OrchestratorRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Orchestrator *OrchestratorFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*OrchestratorRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Orchestrator.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &OrchestratorRoleRevokedIterator{contract: _Orchestrator.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Orchestrator *OrchestratorFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *OrchestratorRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Orchestrator.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrchestratorRoleRevoked)
				if err := _Orchestrator.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Orchestrator *OrchestratorFilterer) ParseRoleRevoked(log types.Log) (*OrchestratorRoleRevoked, error) {
	event := new(OrchestratorRoleRevoked)
	if err := _Orchestrator.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrchestratorUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Orchestrator contract.
type OrchestratorUpgradedIterator struct {
	Event *OrchestratorUpgraded // Event containing the contract specifics and raw log

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
func (it *OrchestratorUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrchestratorUpgraded)
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
		it.Event = new(OrchestratorUpgraded)
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
func (it *OrchestratorUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrchestratorUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrchestratorUpgraded represents a Upgraded event raised by the Orchestrator contract.
type OrchestratorUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Orchestrator *OrchestratorFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*OrchestratorUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Orchestrator.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &OrchestratorUpgradedIterator{contract: _Orchestrator.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Orchestrator *OrchestratorFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *OrchestratorUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Orchestrator.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrchestratorUpgraded)
				if err := _Orchestrator.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Orchestrator *OrchestratorFilterer) ParseUpgraded(log types.Log) (*OrchestratorUpgraded, error) {
	event := new(OrchestratorUpgraded)
	if err := _Orchestrator.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
