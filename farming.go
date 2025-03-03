// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package farming

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

// FarmingMetaData contains all meta data concerning the Farming contract.
var FarmingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"code\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"AssetAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"AssetUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BuyNFTRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"CancelBuyNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"CancelSellNFT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"code\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"DebtAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"DebtUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newEditor\",\"type\":\"address\"}],\"name\":\"EditorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNAV\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"NAVUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"NFTMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refund\",\"type\":\"uint256\"}],\"name\":\"PendingBuyNFTProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sellId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PendingSellCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sellId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PendingSellNFTProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CEIL_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FLOOR_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NATIVE_TOKEN_ADDRESS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"addAllowContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"buyNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"buyNFTCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"buyNFTHistory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refund\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"processed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_totalAsset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_winRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_risk\",\"type\":\"uint256\"}],\"name\":\"calculateNFTPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_transactionId\",\"type\":\"uint256\"}],\"name\":\"cancelBuyNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_sellId\",\"type\":\"uint256\"}],\"name\":\"cancelSellNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ceil\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_transactionId\",\"type\":\"uint256\"}],\"name\":\"finalizeBuyNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_sellId\",\"type\":\"uint256\"}],\"name\":\"finalizeSellNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"floor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllPendingBuyInSession\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllPendingSellInSession\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lockNFT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nftToken\",\"outputs\":[{\"internalType\":\"contractIBEQNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingBuyNFTInSession\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingSellNFTInSession\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveFee\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"removeAllowContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"sellNFT\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sellNFTCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sellNFTHistory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"processed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sessionCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sessionHistory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nftPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAsset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalBuyNFT\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalSellNFT\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalNFT\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"winRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"risk\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"sweepToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBuyInSession\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSellInSession\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdcDecimal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdcToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FarmingABI is the input ABI used to generate the binding from.
// Deprecated: Use FarmingMetaData.ABI instead.
var FarmingABI = FarmingMetaData.ABI

// Farming is an auto generated Go binding around an Ethereum contract.
type Farming struct {
	FarmingCaller     // Read-only binding to the contract
	FarmingTransactor // Write-only binding to the contract
	FarmingFilterer   // Log filterer for contract events
}

// FarmingCaller is an auto generated read-only Go binding around an Ethereum contract.
type FarmingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FarmingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FarmingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FarmingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FarmingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FarmingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FarmingSession struct {
	Contract     *Farming          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FarmingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FarmingCallerSession struct {
	Contract *FarmingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FarmingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FarmingTransactorSession struct {
	Contract     *FarmingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FarmingRaw is an auto generated low-level Go binding around an Ethereum contract.
type FarmingRaw struct {
	Contract *Farming // Generic contract binding to access the raw methods on
}

// FarmingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FarmingCallerRaw struct {
	Contract *FarmingCaller // Generic read-only contract binding to access the raw methods on
}

// FarmingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FarmingTransactorRaw struct {
	Contract *FarmingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFarming creates a new instance of Farming, bound to a specific deployed contract.
func NewFarming(address common.Address, backend bind.ContractBackend) (*Farming, error) {
	contract, err := bindFarming(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Farming{FarmingCaller: FarmingCaller{contract: contract}, FarmingTransactor: FarmingTransactor{contract: contract}, FarmingFilterer: FarmingFilterer{contract: contract}}, nil
}

// NewFarmingCaller creates a new read-only instance of Farming, bound to a specific deployed contract.
func NewFarmingCaller(address common.Address, caller bind.ContractCaller) (*FarmingCaller, error) {
	contract, err := bindFarming(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FarmingCaller{contract: contract}, nil
}

// NewFarmingTransactor creates a new write-only instance of Farming, bound to a specific deployed contract.
func NewFarmingTransactor(address common.Address, transactor bind.ContractTransactor) (*FarmingTransactor, error) {
	contract, err := bindFarming(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FarmingTransactor{contract: contract}, nil
}

// NewFarmingFilterer creates a new log filterer instance of Farming, bound to a specific deployed contract.
func NewFarmingFilterer(address common.Address, filterer bind.ContractFilterer) (*FarmingFilterer, error) {
	contract, err := bindFarming(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FarmingFilterer{contract: contract}, nil
}

// bindFarming binds a generic wrapper to an already deployed contract.
func bindFarming(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FarmingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Farming *FarmingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Farming.Contract.FarmingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Farming *FarmingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Farming.Contract.FarmingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Farming *FarmingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Farming.Contract.FarmingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Farming *FarmingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Farming.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Farming *FarmingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Farming.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Farming *FarmingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Farming.Contract.contract.Transact(opts, method, params...)
}

// CEILFACTOR is a free data retrieval call binding the contract method 0xb8f08248.
//
// Solidity: function CEIL_FACTOR() view returns(uint256)
func (_Farming *FarmingCaller) CEILFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Farming.contract.Call(opts, &out, "CEIL_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CEILFACTOR is a free data retrieval call binding the contract method 0xb8f08248.
//
// Solidity: function CEIL_FACTOR() view returns(uint256)
func (_Farming *FarmingSession) CEILFACTOR() (*big.Int, error) {
	return _Farming.Contract.CEILFACTOR(&_Farming.CallOpts)
}

// CEILFACTOR is a free data retrieval call binding the contract method 0xb8f08248.
//
// Solidity: function CEIL_FACTOR() view returns(uint256)
func (_Farming *FarmingCallerSession) CEILFACTOR() (*big.Int, error) {
	return _Farming.Contract.CEILFACTOR(&_Farming.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Farming *FarmingCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Farming.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Farming *FarmingSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Farming.Contract.DEFAULTADMINROLE(&_Farming.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Farming *FarmingCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Farming.Contract.DEFAULTADMINROLE(&_Farming.CallOpts)
}

// FEEFACTOR is a free data retrieval call binding the contract method 0x5afbc4a8.
//
// Solidity: function FEE_FACTOR() view returns(uint256)
func (_Farming *FarmingCaller) FEEFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Farming.contract.Call(opts, &out, "FEE_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEEFACTOR is a free data retrieval call binding the contract method 0x5afbc4a8.
//
// Solidity: function FEE_FACTOR() view returns(uint256)
func (_Farming *FarmingSession) FEEFACTOR() (*big.Int, error) {
	return _Farming.Contract.FEEFACTOR(&_Farming.CallOpts)
}

// FEEFACTOR is a free data retrieval call binding the contract method 0x5afbc4a8.
//
// Solidity: function FEE_FACTOR() view returns(uint256)
func (_Farming *FarmingCallerSession) FEEFACTOR() (*big.Int, error) {
	return _Farming.Contract.FEEFACTOR(&_Farming.CallOpts)
}

// FLOORFACTOR is a free data retrieval call binding the contract method 0x78908bd9.
//
// Solidity: function FLOOR_FACTOR() view returns(uint256)
func (_Farming *FarmingCaller) FLOORFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Farming.contract.Call(opts, &out, "FLOOR_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FLOORFACTOR is a free data retrieval call binding the contract method 0x78908bd9.
//
// Solidity: function FLOOR_FACTOR() view returns(uint256)
func (_Farming *FarmingSession) FLOORFACTOR() (*big.Int, error) {
	return _Farming.Contract.FLOORFACTOR(&_Farming.CallOpts)
}

// FLOORFACTOR is a free data retrieval call binding the contract method 0x78908bd9.
//
// Solidity: function FLOOR_FACTOR() view returns(uint256)
func (_Farming *FarmingCallerSession) FLOORFACTOR() (*big.Int, error) {
	return _Farming.Contract.FLOORFACTOR(&_Farming.CallOpts)
}

// NATIVETOKENADDRESS is a free data retrieval call binding the contract method 0xdf2ebdbb.
//
// Solidity: function NATIVE_TOKEN_ADDRESS() view returns(address)
func (_Farming *FarmingCaller) NATIVETOKENADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Farming.contract.Call(opts, &out, "NATIVE_TOKEN_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVETOKENADDRESS is a free data retrieval call binding the contract method 0xdf2ebdbb.
//
// Solidity: function NATIVE_TOKEN_ADDRESS() view returns(address)
func (_Farming *FarmingSession) NATIVETOKENADDRESS() (common.Address, error) {
	return _Farming.Contract.NATIVETOKENADDRESS(&_Farming.CallOpts)
}

// NATIVETOKENADDRESS is a free data retrieval call binding the contract method 0xdf2ebdbb.
//
// Solidity: function NATIVE_TOKEN_ADDRESS() view returns(address)
func (_Farming *FarmingCallerSession) NATIVETOKENADDRESS() (common.Address, error) {
	return _Farming.Contract.NATIVETOKENADDRESS(&_Farming.CallOpts)
}

// BaseFee is a free data retrieval call binding the contract method 0x6ef25c3a.
//
// Solidity: function baseFee() view returns(uint256)
func (_Farming *FarmingCaller) BaseFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Farming.contract.Call(opts, &out, "baseFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseFee is a free data retrieval call binding the contract method 0x6ef25c3a.
//
// Solidity: function baseFee() view returns(uint256)
func (_Farming *FarmingSession) BaseFee() (*big.Int, error) {
	return _Farming.Contract.BaseFee(&_Farming.CallOpts)
}

// BaseFee is a free data retrieval call binding the contract method 0x6ef25c3a.
//
// Solidity: function baseFee() view returns(uint256)
func (_Farming *FarmingCallerSession) BaseFee() (*big.Int, error) {
	return _Farming.Contract.BaseFee(&_Farming.CallOpts)
}

// BuyNFTCounter is a free data retrieval call binding the contract method 0xdb409b44.
//
// Solidity: function buyNFTCounter() view returns(uint256)
func (_Farming *FarmingCaller) BuyNFTCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Farming.contract.Call(opts, &out, "buyNFTCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BuyNFTCounter is a free data retrieval call binding the contract method 0xdb409b44.
//
// Solidity: function buyNFTCounter() view returns(uint256)
func (_Farming *FarmingSession) BuyNFTCounter() (*big.Int, error) {
	return _Farming.Contract.BuyNFTCounter(&_Farming.CallOpts)
}

// BuyNFTCounter is a free data retrieval call binding the contract method 0xdb409b44.
//
// Solidity: function buyNFTCounter() view returns(uint256)
func (_Farming *FarmingCallerSession) BuyNFTCounter() (*big.Int, error) {
	return _Farming.Contract.BuyNFTCounter(&_Farming.CallOpts)
}

// BuyNFTHistory is a free data retrieval call binding the contract method 0x76628cb2.
//
// Solidity: function buyNFTHistory(uint256 ) view returns(uint256 id, uint256 sessionId, address buyer, uint256 amount, uint256 refund, uint256 timestamp, bool processed)
func (_Farming *FarmingCaller) BuyNFTHistory(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id        *big.Int
	SessionId *big.Int
	Buyer     common.Address
	Amount    *big.Int
	Refund    *big.Int
	Timestamp *big.Int
	Processed bool
}, error) {
	var out []interface{}
	err := _Farming.contract.Call(opts, &out, "buyNFTHistory", arg0)

	outstruct := new(struct {
		Id        *big.Int
		SessionId *big.Int
		Buyer     common.Address
		Amount    *big.Int
		Refund    *big.Int
		Timestamp *big.Int
		Processed bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SessionId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Buyer = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Refund = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Processed = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// BuyNFTHistory is a free data retrieval call binding the contract method 0x76628cb2.
//
// Solidity: function buyNFTHistory(uint256 ) view returns(uint256 id, uint256 sessionId, address buyer, uint256 amount, uint256 refund, uint256 timestamp, bool processed)
func (_Farming *FarmingSession) BuyNFTHistory(arg0 *big.Int) (struct {
	Id        *big.Int
	SessionId *big.Int
	Buyer     common.Address
	Amount    *big.Int
	Refund    *big.Int
	Timestamp *big.Int
	Processed bool
}, error) {
	return _Farming.Contract.BuyNFTHistory(&_Farming.CallOpts, arg0)
}

// BuyNFTHistory is a free data retrieval call binding the contract method 0x76628cb2.
//
// Solidity: function buyNFTHistory(uint256 ) view returns(uint256 id, uint256 sessionId, address buyer, uint256 amount, uint256 refund, uint256 timestamp, bool processed)
func (_Farming *FarmingCallerSession) BuyNFTHistory(arg0 *big.Int) (struct {
	Id        *big.Int
	SessionId *big.Int
	Buyer     common.Address
	Amount    *big.Int
	Refund    *big.Int
	Timestamp *big.Int
	Processed bool
}, error) {
	return _Farming.Contract.BuyNFTHistory(&_Farming.CallOpts, arg0)
}

// Ceil is a free data retrieval call binding the contract method 0x5ed03e57.
//
// Solidity: function ceil() view returns(uint256)
func (_Farming *FarmingCaller) Ceil(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Farming.contract.Call(opts, &out, "ceil")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Ceil is a free data retrieval call binding the contract method 0x5ed03e57.
//
// Solidity: function ceil() view returns(uint256)
func (_Farming *FarmingSession) Ceil() (*big.Int, error) {
	return _Farming.Contract.Ceil(&_Farming.CallOpts)
}

// Ceil is a free data retrieval call binding the contract method 0x5ed03e57.
//
// Solidity: function ceil() view returns(uint256)
func (_Farming *FarmingCallerSession) Ceil() (*big.Int, error) {
	return _Farming.Contract.Ceil(&_Farming.CallOpts)
}

// Floor is a free data retrieval call binding the contract method 0x40695363.
//
// Solidity: function floor() view returns(uint256)
func (_Farming *FarmingCaller) Floor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Farming.contract.Call(opts, &out, "floor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Floor is a free data retrieval call binding the contract method 0x40695363.
//
// Solidity: function floor() view returns(uint256)
func (_Farming *FarmingSession) Floor() (*big.Int, error) {
	return _Farming.Contract.Floor(&_Farming.CallOpts)
}

// Floor is a free data retrieval call binding the contract method 0x40695363.
//
// Solidity: function floor() view returns(uint256)
func (_Farming *FarmingCallerSession) Floor() (*big.Int, error) {
	return _Farming.Contract.Floor(&_Farming.CallOpts)
}

// GetAllPendingBuyInSession is a free data retrieval call binding the contract method 0x65f6f2ad.
//
// Solidity: function getAllPendingBuyInSession() view returns(uint256[])
func (_Farming *FarmingCaller) GetAllPendingBuyInSession(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Farming.contract.Call(opts, &out, "getAllPendingBuyInSession")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAllPendingBuyInSession is a free data retrieval call binding the contract method 0x65f6f2ad.
//
// Solidity: function getAllPendingBuyInSession() view returns(uint256[])
func (_Farming *FarmingSession) GetAllPendingBuyInSession() ([]*big.Int, error) {
	return _Farming.Contract.GetAllPendingBuyInSession(&_Farming.CallOpts)
}

// GetAllPendingBuyInSession is a free data retrieval call binding the contract method 0x65f6f2ad.
//
// Solidity: function getAllPendingBuyInSession() view returns(uint256[])
func (_Farming *FarmingCallerSession) GetAllPendingBuyInSession() ([]*big.Int, error) {
	return _Farming.Contract.GetAllPendingBuyInSession(&_Farming.CallOpts)
}

// GetAllPendingSellInSession is a free data retrieval call binding the contract method 0x2e50f48b.
//
// Solidity: function getAllPendingSellInSession() view returns(uint256[])
func (_Farming *FarmingCaller) GetAllPendingSellInSession(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Farming.contract.Call(opts, &out, "getAllPendingSellInSession")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAllPendingSellInSession is a free data retrieval call binding the contract method 0x2e50f48b.
//
// Solidity: function getAllPendingSellInSession() view returns(uint256[])
func (_Farming *FarmingSession) GetAllPendingSellInSession() ([]*big.Int, error) {
	return _Farming.Contract.GetAllPendingSellInSession(&_Farming.CallOpts)
}

// GetAllPendingSellInSession is a free data retrieval call binding the contract method 0x2e50f48b.
//
// Solidity: function getAllPendingSellInSession() view returns(uint256[])
func (_Farming *FarmingCallerSession) GetAllPendingSellInSession() ([]*big.Int, error) {
	return _Farming.Contract.GetAllPendingSellInSession(&_Farming.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Farming *FarmingCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Farming.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Farming *FarmingSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Farming.Contract.GetRoleAdmin(&_Farming.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Farming *FarmingCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Farming.Contract.GetRoleAdmin(&_Farming.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Farming *FarmingCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Farming.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Farming *FarmingSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Farming.Contract.HasRole(&_Farming.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Farming *FarmingCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Farming.Contract.HasRole(&_Farming.CallOpts, role, account)
}

// LockNFT is a free data retrieval call binding the contract method 0x2b7170d0.
//
// Solidity: function lockNFT(uint256 ) view returns(uint256)
func (_Farming *FarmingCaller) LockNFT(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Farming.contract.Call(opts, &out, "lockNFT", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockNFT is a free data retrieval call binding the contract method 0x2b7170d0.
//
// Solidity: function lockNFT(uint256 ) view returns(uint256)
func (_Farming *FarmingSession) LockNFT(arg0 *big.Int) (*big.Int, error) {
	return _Farming.Contract.LockNFT(&_Farming.CallOpts, arg0)
}

// LockNFT is a free data retrieval call binding the contract method 0x2b7170d0.
//
// Solidity: function lockNFT(uint256 ) view returns(uint256)
func (_Farming *FarmingCallerSession) LockNFT(arg0 *big.Int) (*big.Int, error) {
	return _Farming.Contract.LockNFT(&_Farming.CallOpts, arg0)
}

// NftToken is a free data retrieval call binding the contract method 0xd06fcba8.
//
// Solidity: function nftToken() view returns(address)
func (_Farming *FarmingCaller) NftToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Farming.contract.Call(opts, &out, "nftToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftToken is a free data retrieval call binding the contract method 0xd06fcba8.
//
// Solidity: function nftToken() view returns(address)
func (_Farming *FarmingSession) NftToken() (common.Address, error) {
	return _Farming.Contract.NftToken(&_Farming.CallOpts)
}

// NftToken is a free data retrieval call binding the contract method 0xd06fcba8.
//
// Solidity: function nftToken() view returns(address)
func (_Farming *FarmingCallerSession) NftToken() (common.Address, error) {
	return _Farming.Contract.NftToken(&_Farming.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Farming *FarmingCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Farming.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Farming *FarmingSession) Paused() (bool, error) {
	return _Farming.Contract.Paused(&_Farming.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Farming *FarmingCallerSession) Paused() (bool, error) {
	return _Farming.Contract.Paused(&_Farming.CallOpts)
}

// PendingBuyNFTInSession is a free data retrieval call binding the contract method 0xfb1a5551.
//
// Solidity: function pendingBuyNFTInSession(uint256 ) view returns(uint256)
func (_Farming *FarmingCaller) PendingBuyNFTInSession(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Farming.contract.Call(opts, &out, "pendingBuyNFTInSession", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingBuyNFTInSession is a free data retrieval call binding the contract method 0xfb1a5551.
//
// Solidity: function pendingBuyNFTInSession(uint256 ) view returns(uint256)
func (_Farming *FarmingSession) PendingBuyNFTInSession(arg0 *big.Int) (*big.Int, error) {
	return _Farming.Contract.PendingBuyNFTInSession(&_Farming.CallOpts, arg0)
}

// PendingBuyNFTInSession is a free data retrieval call binding the contract method 0xfb1a5551.
//
// Solidity: function pendingBuyNFTInSession(uint256 ) view returns(uint256)
func (_Farming *FarmingCallerSession) PendingBuyNFTInSession(arg0 *big.Int) (*big.Int, error) {
	return _Farming.Contract.PendingBuyNFTInSession(&_Farming.CallOpts, arg0)
}

// PendingSellNFTInSession is a free data retrieval call binding the contract method 0x1c5df308.
//
// Solidity: function pendingSellNFTInSession(uint256 ) view returns(uint256)
func (_Farming *FarmingCaller) PendingSellNFTInSession(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Farming.contract.Call(opts, &out, "pendingSellNFTInSession", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingSellNFTInSession is a free data retrieval call binding the contract method 0x1c5df308.
//
// Solidity: function pendingSellNFTInSession(uint256 ) view returns(uint256)
func (_Farming *FarmingSession) PendingSellNFTInSession(arg0 *big.Int) (*big.Int, error) {
	return _Farming.Contract.PendingSellNFTInSession(&_Farming.CallOpts, arg0)
}

// PendingSellNFTInSession is a free data retrieval call binding the contract method 0x1c5df308.
//
// Solidity: function pendingSellNFTInSession(uint256 ) view returns(uint256)
func (_Farming *FarmingCallerSession) PendingSellNFTInSession(arg0 *big.Int) (*big.Int, error) {
	return _Farming.Contract.PendingSellNFTInSession(&_Farming.CallOpts, arg0)
}

// ReceiveFee is a free data retrieval call binding the contract method 0xe18b7fcf.
//
// Solidity: function receiveFee() view returns(address)
func (_Farming *FarmingCaller) ReceiveFee(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Farming.contract.Call(opts, &out, "receiveFee")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReceiveFee is a free data retrieval call binding the contract method 0xe18b7fcf.
//
// Solidity: function receiveFee() view returns(address)
func (_Farming *FarmingSession) ReceiveFee() (common.Address, error) {
	return _Farming.Contract.ReceiveFee(&_Farming.CallOpts)
}

// ReceiveFee is a free data retrieval call binding the contract method 0xe18b7fcf.
//
// Solidity: function receiveFee() view returns(address)
func (_Farming *FarmingCallerSession) ReceiveFee() (common.Address, error) {
	return _Farming.Contract.ReceiveFee(&_Farming.CallOpts)
}

// SellNFTCounter is a free data retrieval call binding the contract method 0x783b1b01.
//
// Solidity: function sellNFTCounter() view returns(uint256)
func (_Farming *FarmingCaller) SellNFTCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Farming.contract.Call(opts, &out, "sellNFTCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SellNFTCounter is a free data retrieval call binding the contract method 0x783b1b01.
//
// Solidity: function sellNFTCounter() view returns(uint256)
func (_Farming *FarmingSession) SellNFTCounter() (*big.Int, error) {
	return _Farming.Contract.SellNFTCounter(&_Farming.CallOpts)
}

// SellNFTCounter is a free data retrieval call binding the contract method 0x783b1b01.
//
// Solidity: function sellNFTCounter() view returns(uint256)
func (_Farming *FarmingCallerSession) SellNFTCounter() (*big.Int, error) {
	return _Farming.Contract.SellNFTCounter(&_Farming.CallOpts)
}

// SellNFTHistory is a free data retrieval call binding the contract method 0xe565bce4.
//
// Solidity: function sellNFTHistory(uint256 ) view returns(uint256 id, uint256 sessionId, address seller, address buyer, uint256 tokenId, uint256 price, uint256 timestamp, bool processed)
func (_Farming *FarmingCaller) SellNFTHistory(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id        *big.Int
	SessionId *big.Int
	Seller    common.Address
	Buyer     common.Address
	TokenId   *big.Int
	Price     *big.Int
	Timestamp *big.Int
	Processed bool
}, error) {
	var out []interface{}
	err := _Farming.contract.Call(opts, &out, "sellNFTHistory", arg0)

	outstruct := new(struct {
		Id        *big.Int
		SessionId *big.Int
		Seller    common.Address
		Buyer     common.Address
		TokenId   *big.Int
		Price     *big.Int
		Timestamp *big.Int
		Processed bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SessionId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Seller = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Buyer = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Price = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Processed = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// SellNFTHistory is a free data retrieval call binding the contract method 0xe565bce4.
//
// Solidity: function sellNFTHistory(uint256 ) view returns(uint256 id, uint256 sessionId, address seller, address buyer, uint256 tokenId, uint256 price, uint256 timestamp, bool processed)
func (_Farming *FarmingSession) SellNFTHistory(arg0 *big.Int) (struct {
	Id        *big.Int
	SessionId *big.Int
	Seller    common.Address
	Buyer     common.Address
	TokenId   *big.Int
	Price     *big.Int
	Timestamp *big.Int
	Processed bool
}, error) {
	return _Farming.Contract.SellNFTHistory(&_Farming.CallOpts, arg0)
}

// SellNFTHistory is a free data retrieval call binding the contract method 0xe565bce4.
//
// Solidity: function sellNFTHistory(uint256 ) view returns(uint256 id, uint256 sessionId, address seller, address buyer, uint256 tokenId, uint256 price, uint256 timestamp, bool processed)
func (_Farming *FarmingCallerSession) SellNFTHistory(arg0 *big.Int) (struct {
	Id        *big.Int
	SessionId *big.Int
	Seller    common.Address
	Buyer     common.Address
	TokenId   *big.Int
	Price     *big.Int
	Timestamp *big.Int
	Processed bool
}, error) {
	return _Farming.Contract.SellNFTHistory(&_Farming.CallOpts, arg0)
}

// SessionCounter is a free data retrieval call binding the contract method 0xcc64e2af.
//
// Solidity: function sessionCounter() view returns(uint256)
func (_Farming *FarmingCaller) SessionCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Farming.contract.Call(opts, &out, "sessionCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SessionCounter is a free data retrieval call binding the contract method 0xcc64e2af.
//
// Solidity: function sessionCounter() view returns(uint256)
func (_Farming *FarmingSession) SessionCounter() (*big.Int, error) {
	return _Farming.Contract.SessionCounter(&_Farming.CallOpts)
}

// SessionCounter is a free data retrieval call binding the contract method 0xcc64e2af.
//
// Solidity: function sessionCounter() view returns(uint256)
func (_Farming *FarmingCallerSession) SessionCounter() (*big.Int, error) {
	return _Farming.Contract.SessionCounter(&_Farming.CallOpts)
}

// SessionHistory is a free data retrieval call binding the contract method 0x5119d25e.
//
// Solidity: function sessionHistory(uint256 ) view returns(uint256 id, uint256 nftPrice, uint256 totalAsset, uint256 totalBuyNFT, uint256 totalSellNFT, uint256 totalNFT, uint256 winRate, uint256 risk, uint256 timestamp)
func (_Farming *FarmingCaller) SessionHistory(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id           *big.Int
	NftPrice     *big.Int
	TotalAsset   *big.Int
	TotalBuyNFT  *big.Int
	TotalSellNFT *big.Int
	TotalNFT     *big.Int
	WinRate      *big.Int
	Risk         *big.Int
	Timestamp    *big.Int
}, error) {
	var out []interface{}
	err := _Farming.contract.Call(opts, &out, "sessionHistory", arg0)

	outstruct := new(struct {
		Id           *big.Int
		NftPrice     *big.Int
		TotalAsset   *big.Int
		TotalBuyNFT  *big.Int
		TotalSellNFT *big.Int
		TotalNFT     *big.Int
		WinRate      *big.Int
		Risk         *big.Int
		Timestamp    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NftPrice = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalAsset = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalBuyNFT = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TotalSellNFT = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TotalNFT = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.WinRate = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Risk = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SessionHistory is a free data retrieval call binding the contract method 0x5119d25e.
//
// Solidity: function sessionHistory(uint256 ) view returns(uint256 id, uint256 nftPrice, uint256 totalAsset, uint256 totalBuyNFT, uint256 totalSellNFT, uint256 totalNFT, uint256 winRate, uint256 risk, uint256 timestamp)
func (_Farming *FarmingSession) SessionHistory(arg0 *big.Int) (struct {
	Id           *big.Int
	NftPrice     *big.Int
	TotalAsset   *big.Int
	TotalBuyNFT  *big.Int
	TotalSellNFT *big.Int
	TotalNFT     *big.Int
	WinRate      *big.Int
	Risk         *big.Int
	Timestamp    *big.Int
}, error) {
	return _Farming.Contract.SessionHistory(&_Farming.CallOpts, arg0)
}

// SessionHistory is a free data retrieval call binding the contract method 0x5119d25e.
//
// Solidity: function sessionHistory(uint256 ) view returns(uint256 id, uint256 nftPrice, uint256 totalAsset, uint256 totalBuyNFT, uint256 totalSellNFT, uint256 totalNFT, uint256 winRate, uint256 risk, uint256 timestamp)
func (_Farming *FarmingCallerSession) SessionHistory(arg0 *big.Int) (struct {
	Id           *big.Int
	NftPrice     *big.Int
	TotalAsset   *big.Int
	TotalBuyNFT  *big.Int
	TotalSellNFT *big.Int
	TotalNFT     *big.Int
	WinRate      *big.Int
	Risk         *big.Int
	Timestamp    *big.Int
}, error) {
	return _Farming.Contract.SessionHistory(&_Farming.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Farming *FarmingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Farming.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Farming *FarmingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Farming.Contract.SupportsInterface(&_Farming.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Farming *FarmingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Farming.Contract.SupportsInterface(&_Farming.CallOpts, interfaceId)
}

// TotalBuyInSession is a free data retrieval call binding the contract method 0xf6ae4a74.
//
// Solidity: function totalBuyInSession() view returns(uint256)
func (_Farming *FarmingCaller) TotalBuyInSession(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Farming.contract.Call(opts, &out, "totalBuyInSession")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBuyInSession is a free data retrieval call binding the contract method 0xf6ae4a74.
//
// Solidity: function totalBuyInSession() view returns(uint256)
func (_Farming *FarmingSession) TotalBuyInSession() (*big.Int, error) {
	return _Farming.Contract.TotalBuyInSession(&_Farming.CallOpts)
}

// TotalBuyInSession is a free data retrieval call binding the contract method 0xf6ae4a74.
//
// Solidity: function totalBuyInSession() view returns(uint256)
func (_Farming *FarmingCallerSession) TotalBuyInSession() (*big.Int, error) {
	return _Farming.Contract.TotalBuyInSession(&_Farming.CallOpts)
}

// TotalSellInSession is a free data retrieval call binding the contract method 0x6152c341.
//
// Solidity: function totalSellInSession() view returns(uint256)
func (_Farming *FarmingCaller) TotalSellInSession(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Farming.contract.Call(opts, &out, "totalSellInSession")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSellInSession is a free data retrieval call binding the contract method 0x6152c341.
//
// Solidity: function totalSellInSession() view returns(uint256)
func (_Farming *FarmingSession) TotalSellInSession() (*big.Int, error) {
	return _Farming.Contract.TotalSellInSession(&_Farming.CallOpts)
}

// TotalSellInSession is a free data retrieval call binding the contract method 0x6152c341.
//
// Solidity: function totalSellInSession() view returns(uint256)
func (_Farming *FarmingCallerSession) TotalSellInSession() (*big.Int, error) {
	return _Farming.Contract.TotalSellInSession(&_Farming.CallOpts)
}

// UsdcDecimal is a free data retrieval call binding the contract method 0x4abe3ce5.
//
// Solidity: function usdcDecimal() view returns(uint256)
func (_Farming *FarmingCaller) UsdcDecimal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Farming.contract.Call(opts, &out, "usdcDecimal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdcDecimal is a free data retrieval call binding the contract method 0x4abe3ce5.
//
// Solidity: function usdcDecimal() view returns(uint256)
func (_Farming *FarmingSession) UsdcDecimal() (*big.Int, error) {
	return _Farming.Contract.UsdcDecimal(&_Farming.CallOpts)
}

// UsdcDecimal is a free data retrieval call binding the contract method 0x4abe3ce5.
//
// Solidity: function usdcDecimal() view returns(uint256)
func (_Farming *FarmingCallerSession) UsdcDecimal() (*big.Int, error) {
	return _Farming.Contract.UsdcDecimal(&_Farming.CallOpts)
}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_Farming *FarmingCaller) UsdcToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Farming.contract.Call(opts, &out, "usdcToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_Farming *FarmingSession) UsdcToken() (common.Address, error) {
	return _Farming.Contract.UsdcToken(&_Farming.CallOpts)
}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_Farming *FarmingCallerSession) UsdcToken() (common.Address, error) {
	return _Farming.Contract.UsdcToken(&_Farming.CallOpts)
}

// AddAllowContract is a paid mutator transaction binding the contract method 0x0c7f6d67.
//
// Solidity: function addAllowContract(address _contractAddress) returns()
func (_Farming *FarmingTransactor) AddAllowContract(opts *bind.TransactOpts, _contractAddress common.Address) (*types.Transaction, error) {
	return _Farming.contract.Transact(opts, "addAllowContract", _contractAddress)
}

// AddAllowContract is a paid mutator transaction binding the contract method 0x0c7f6d67.
//
// Solidity: function addAllowContract(address _contractAddress) returns()
func (_Farming *FarmingSession) AddAllowContract(_contractAddress common.Address) (*types.Transaction, error) {
	return _Farming.Contract.AddAllowContract(&_Farming.TransactOpts, _contractAddress)
}

// AddAllowContract is a paid mutator transaction binding the contract method 0x0c7f6d67.
//
// Solidity: function addAllowContract(address _contractAddress) returns()
func (_Farming *FarmingTransactorSession) AddAllowContract(_contractAddress common.Address) (*types.Transaction, error) {
	return _Farming.Contract.AddAllowContract(&_Farming.TransactOpts, _contractAddress)
}

// BuyNFT is a paid mutator transaction binding the contract method 0x51ed8288.
//
// Solidity: function buyNFT(uint256 amount) returns()
func (_Farming *FarmingTransactor) BuyNFT(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Farming.contract.Transact(opts, "buyNFT", amount)
}

// BuyNFT is a paid mutator transaction binding the contract method 0x51ed8288.
//
// Solidity: function buyNFT(uint256 amount) returns()
func (_Farming *FarmingSession) BuyNFT(amount *big.Int) (*types.Transaction, error) {
	return _Farming.Contract.BuyNFT(&_Farming.TransactOpts, amount)
}

// BuyNFT is a paid mutator transaction binding the contract method 0x51ed8288.
//
// Solidity: function buyNFT(uint256 amount) returns()
func (_Farming *FarmingTransactorSession) BuyNFT(amount *big.Int) (*types.Transaction, error) {
	return _Farming.Contract.BuyNFT(&_Farming.TransactOpts, amount)
}

// CalculateNFTPrice is a paid mutator transaction binding the contract method 0x25af38d9.
//
// Solidity: function calculateNFTPrice(uint256 _totalAsset, uint256 _winRate, uint256 _risk) returns()
func (_Farming *FarmingTransactor) CalculateNFTPrice(opts *bind.TransactOpts, _totalAsset *big.Int, _winRate *big.Int, _risk *big.Int) (*types.Transaction, error) {
	return _Farming.contract.Transact(opts, "calculateNFTPrice", _totalAsset, _winRate, _risk)
}

// CalculateNFTPrice is a paid mutator transaction binding the contract method 0x25af38d9.
//
// Solidity: function calculateNFTPrice(uint256 _totalAsset, uint256 _winRate, uint256 _risk) returns()
func (_Farming *FarmingSession) CalculateNFTPrice(_totalAsset *big.Int, _winRate *big.Int, _risk *big.Int) (*types.Transaction, error) {
	return _Farming.Contract.CalculateNFTPrice(&_Farming.TransactOpts, _totalAsset, _winRate, _risk)
}

// CalculateNFTPrice is a paid mutator transaction binding the contract method 0x25af38d9.
//
// Solidity: function calculateNFTPrice(uint256 _totalAsset, uint256 _winRate, uint256 _risk) returns()
func (_Farming *FarmingTransactorSession) CalculateNFTPrice(_totalAsset *big.Int, _winRate *big.Int, _risk *big.Int) (*types.Transaction, error) {
	return _Farming.Contract.CalculateNFTPrice(&_Farming.TransactOpts, _totalAsset, _winRate, _risk)
}

// CancelBuyNFT is a paid mutator transaction binding the contract method 0xb9902bc1.
//
// Solidity: function cancelBuyNFT(uint256 _transactionId) returns()
func (_Farming *FarmingTransactor) CancelBuyNFT(opts *bind.TransactOpts, _transactionId *big.Int) (*types.Transaction, error) {
	return _Farming.contract.Transact(opts, "cancelBuyNFT", _transactionId)
}

// CancelBuyNFT is a paid mutator transaction binding the contract method 0xb9902bc1.
//
// Solidity: function cancelBuyNFT(uint256 _transactionId) returns()
func (_Farming *FarmingSession) CancelBuyNFT(_transactionId *big.Int) (*types.Transaction, error) {
	return _Farming.Contract.CancelBuyNFT(&_Farming.TransactOpts, _transactionId)
}

// CancelBuyNFT is a paid mutator transaction binding the contract method 0xb9902bc1.
//
// Solidity: function cancelBuyNFT(uint256 _transactionId) returns()
func (_Farming *FarmingTransactorSession) CancelBuyNFT(_transactionId *big.Int) (*types.Transaction, error) {
	return _Farming.Contract.CancelBuyNFT(&_Farming.TransactOpts, _transactionId)
}

// CancelSellNFT is a paid mutator transaction binding the contract method 0xacd3db01.
//
// Solidity: function cancelSellNFT(uint256 _sellId) returns()
func (_Farming *FarmingTransactor) CancelSellNFT(opts *bind.TransactOpts, _sellId *big.Int) (*types.Transaction, error) {
	return _Farming.contract.Transact(opts, "cancelSellNFT", _sellId)
}

// CancelSellNFT is a paid mutator transaction binding the contract method 0xacd3db01.
//
// Solidity: function cancelSellNFT(uint256 _sellId) returns()
func (_Farming *FarmingSession) CancelSellNFT(_sellId *big.Int) (*types.Transaction, error) {
	return _Farming.Contract.CancelSellNFT(&_Farming.TransactOpts, _sellId)
}

// CancelSellNFT is a paid mutator transaction binding the contract method 0xacd3db01.
//
// Solidity: function cancelSellNFT(uint256 _sellId) returns()
func (_Farming *FarmingTransactorSession) CancelSellNFT(_sellId *big.Int) (*types.Transaction, error) {
	return _Farming.Contract.CancelSellNFT(&_Farming.TransactOpts, _sellId)
}

// FinalizeBuyNFT is a paid mutator transaction binding the contract method 0x0fc6e4bf.
//
// Solidity: function finalizeBuyNFT(uint256 _transactionId) returns()
func (_Farming *FarmingTransactor) FinalizeBuyNFT(opts *bind.TransactOpts, _transactionId *big.Int) (*types.Transaction, error) {
	return _Farming.contract.Transact(opts, "finalizeBuyNFT", _transactionId)
}

// FinalizeBuyNFT is a paid mutator transaction binding the contract method 0x0fc6e4bf.
//
// Solidity: function finalizeBuyNFT(uint256 _transactionId) returns()
func (_Farming *FarmingSession) FinalizeBuyNFT(_transactionId *big.Int) (*types.Transaction, error) {
	return _Farming.Contract.FinalizeBuyNFT(&_Farming.TransactOpts, _transactionId)
}

// FinalizeBuyNFT is a paid mutator transaction binding the contract method 0x0fc6e4bf.
//
// Solidity: function finalizeBuyNFT(uint256 _transactionId) returns()
func (_Farming *FarmingTransactorSession) FinalizeBuyNFT(_transactionId *big.Int) (*types.Transaction, error) {
	return _Farming.Contract.FinalizeBuyNFT(&_Farming.TransactOpts, _transactionId)
}

// FinalizeSellNFT is a paid mutator transaction binding the contract method 0x2719cbd4.
//
// Solidity: function finalizeSellNFT(uint256 _sellId) returns()
func (_Farming *FarmingTransactor) FinalizeSellNFT(opts *bind.TransactOpts, _sellId *big.Int) (*types.Transaction, error) {
	return _Farming.contract.Transact(opts, "finalizeSellNFT", _sellId)
}

// FinalizeSellNFT is a paid mutator transaction binding the contract method 0x2719cbd4.
//
// Solidity: function finalizeSellNFT(uint256 _sellId) returns()
func (_Farming *FarmingSession) FinalizeSellNFT(_sellId *big.Int) (*types.Transaction, error) {
	return _Farming.Contract.FinalizeSellNFT(&_Farming.TransactOpts, _sellId)
}

// FinalizeSellNFT is a paid mutator transaction binding the contract method 0x2719cbd4.
//
// Solidity: function finalizeSellNFT(uint256 _sellId) returns()
func (_Farming *FarmingTransactorSession) FinalizeSellNFT(_sellId *big.Int) (*types.Transaction, error) {
	return _Farming.Contract.FinalizeSellNFT(&_Farming.TransactOpts, _sellId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Farming *FarmingTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Farming.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Farming *FarmingSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Farming.Contract.GrantRole(&_Farming.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Farming *FarmingTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Farming.Contract.GrantRole(&_Farming.TransactOpts, role, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Farming *FarmingTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Farming.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Farming *FarmingSession) Pause() (*types.Transaction, error) {
	return _Farming.Contract.Pause(&_Farming.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Farming *FarmingTransactorSession) Pause() (*types.Transaction, error) {
	return _Farming.Contract.Pause(&_Farming.TransactOpts)
}

// RemoveAllowContract is a paid mutator transaction binding the contract method 0x89382392.
//
// Solidity: function removeAllowContract(address _contractAddress) returns()
func (_Farming *FarmingTransactor) RemoveAllowContract(opts *bind.TransactOpts, _contractAddress common.Address) (*types.Transaction, error) {
	return _Farming.contract.Transact(opts, "removeAllowContract", _contractAddress)
}

// RemoveAllowContract is a paid mutator transaction binding the contract method 0x89382392.
//
// Solidity: function removeAllowContract(address _contractAddress) returns()
func (_Farming *FarmingSession) RemoveAllowContract(_contractAddress common.Address) (*types.Transaction, error) {
	return _Farming.Contract.RemoveAllowContract(&_Farming.TransactOpts, _contractAddress)
}

// RemoveAllowContract is a paid mutator transaction binding the contract method 0x89382392.
//
// Solidity: function removeAllowContract(address _contractAddress) returns()
func (_Farming *FarmingTransactorSession) RemoveAllowContract(_contractAddress common.Address) (*types.Transaction, error) {
	return _Farming.Contract.RemoveAllowContract(&_Farming.TransactOpts, _contractAddress)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Farming *FarmingTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Farming.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Farming *FarmingSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Farming.Contract.RenounceRole(&_Farming.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Farming *FarmingTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Farming.Contract.RenounceRole(&_Farming.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Farming *FarmingTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Farming.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Farming *FarmingSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Farming.Contract.RevokeRole(&_Farming.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Farming *FarmingTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Farming.Contract.RevokeRole(&_Farming.TransactOpts, role, account)
}

// SellNFT is a paid mutator transaction binding the contract method 0xee9cdfb3.
//
// Solidity: function sellNFT(uint256 _tokenId) returns()
func (_Farming *FarmingTransactor) SellNFT(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _Farming.contract.Transact(opts, "sellNFT", _tokenId)
}

// SellNFT is a paid mutator transaction binding the contract method 0xee9cdfb3.
//
// Solidity: function sellNFT(uint256 _tokenId) returns()
func (_Farming *FarmingSession) SellNFT(_tokenId *big.Int) (*types.Transaction, error) {
	return _Farming.Contract.SellNFT(&_Farming.TransactOpts, _tokenId)
}

// SellNFT is a paid mutator transaction binding the contract method 0xee9cdfb3.
//
// Solidity: function sellNFT(uint256 _tokenId) returns()
func (_Farming *FarmingTransactorSession) SellNFT(_tokenId *big.Int) (*types.Transaction, error) {
	return _Farming.Contract.SellNFT(&_Farming.TransactOpts, _tokenId)
}

// SweepToken is a paid mutator transaction binding the contract method 0x1be19560.
//
// Solidity: function sweepToken(address _token) returns()
func (_Farming *FarmingTransactor) SweepToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Farming.contract.Transact(opts, "sweepToken", _token)
}

// SweepToken is a paid mutator transaction binding the contract method 0x1be19560.
//
// Solidity: function sweepToken(address _token) returns()
func (_Farming *FarmingSession) SweepToken(_token common.Address) (*types.Transaction, error) {
	return _Farming.Contract.SweepToken(&_Farming.TransactOpts, _token)
}

// SweepToken is a paid mutator transaction binding the contract method 0x1be19560.
//
// Solidity: function sweepToken(address _token) returns()
func (_Farming *FarmingTransactorSession) SweepToken(_token common.Address) (*types.Transaction, error) {
	return _Farming.Contract.SweepToken(&_Farming.TransactOpts, _token)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Farming *FarmingTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Farming.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Farming *FarmingSession) Unpause() (*types.Transaction, error) {
	return _Farming.Contract.Unpause(&_Farming.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Farming *FarmingTransactorSession) Unpause() (*types.Transaction, error) {
	return _Farming.Contract.Unpause(&_Farming.TransactOpts)
}

// FarmingAssetAddedIterator is returned from FilterAssetAdded and is used to iterate over the raw logs and unpacked data for AssetAdded events raised by the Farming contract.
type FarmingAssetAddedIterator struct {
	Event *FarmingAssetAdded // Event containing the contract specifics and raw log

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
func (it *FarmingAssetAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmingAssetAdded)
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
		it.Event = new(FarmingAssetAdded)
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
func (it *FarmingAssetAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmingAssetAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmingAssetAdded represents a AssetAdded event raised by the Farming contract.
type FarmingAssetAdded struct {
	Name  string
	Code  string
	Value *big.Int
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAssetAdded is a free log retrieval operation binding the contract event 0x70d8d7024d92c8032daf01f544cb61321ce9e0607e32ab331e4fd547c53f9b1e.
//
// Solidity: event AssetAdded(string name, string code, uint256 value, uint256 index)
func (_Farming *FarmingFilterer) FilterAssetAdded(opts *bind.FilterOpts) (*FarmingAssetAddedIterator, error) {

	logs, sub, err := _Farming.contract.FilterLogs(opts, "AssetAdded")
	if err != nil {
		return nil, err
	}
	return &FarmingAssetAddedIterator{contract: _Farming.contract, event: "AssetAdded", logs: logs, sub: sub}, nil
}

// WatchAssetAdded is a free log subscription operation binding the contract event 0x70d8d7024d92c8032daf01f544cb61321ce9e0607e32ab331e4fd547c53f9b1e.
//
// Solidity: event AssetAdded(string name, string code, uint256 value, uint256 index)
func (_Farming *FarmingFilterer) WatchAssetAdded(opts *bind.WatchOpts, sink chan<- *FarmingAssetAdded) (event.Subscription, error) {

	logs, sub, err := _Farming.contract.WatchLogs(opts, "AssetAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmingAssetAdded)
				if err := _Farming.contract.UnpackLog(event, "AssetAdded", log); err != nil {
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

// ParseAssetAdded is a log parse operation binding the contract event 0x70d8d7024d92c8032daf01f544cb61321ce9e0607e32ab331e4fd547c53f9b1e.
//
// Solidity: event AssetAdded(string name, string code, uint256 value, uint256 index)
func (_Farming *FarmingFilterer) ParseAssetAdded(log types.Log) (*FarmingAssetAdded, error) {
	event := new(FarmingAssetAdded)
	if err := _Farming.contract.UnpackLog(event, "AssetAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmingAssetUpdatedIterator is returned from FilterAssetUpdated and is used to iterate over the raw logs and unpacked data for AssetUpdated events raised by the Farming contract.
type FarmingAssetUpdatedIterator struct {
	Event *FarmingAssetUpdated // Event containing the contract specifics and raw log

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
func (it *FarmingAssetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmingAssetUpdated)
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
		it.Event = new(FarmingAssetUpdated)
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
func (it *FarmingAssetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmingAssetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmingAssetUpdated represents a AssetUpdated event raised by the Farming contract.
type FarmingAssetUpdated struct {
	Index *big.Int
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAssetUpdated is a free log retrieval operation binding the contract event 0x314fa82f40c9e1ea2150ffdfe35d573ea4663cf019f4945ea8130a2df4454cdf.
//
// Solidity: event AssetUpdated(uint256 index, uint256 value)
func (_Farming *FarmingFilterer) FilterAssetUpdated(opts *bind.FilterOpts) (*FarmingAssetUpdatedIterator, error) {

	logs, sub, err := _Farming.contract.FilterLogs(opts, "AssetUpdated")
	if err != nil {
		return nil, err
	}
	return &FarmingAssetUpdatedIterator{contract: _Farming.contract, event: "AssetUpdated", logs: logs, sub: sub}, nil
}

// WatchAssetUpdated is a free log subscription operation binding the contract event 0x314fa82f40c9e1ea2150ffdfe35d573ea4663cf019f4945ea8130a2df4454cdf.
//
// Solidity: event AssetUpdated(uint256 index, uint256 value)
func (_Farming *FarmingFilterer) WatchAssetUpdated(opts *bind.WatchOpts, sink chan<- *FarmingAssetUpdated) (event.Subscription, error) {

	logs, sub, err := _Farming.contract.WatchLogs(opts, "AssetUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmingAssetUpdated)
				if err := _Farming.contract.UnpackLog(event, "AssetUpdated", log); err != nil {
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

// ParseAssetUpdated is a log parse operation binding the contract event 0x314fa82f40c9e1ea2150ffdfe35d573ea4663cf019f4945ea8130a2df4454cdf.
//
// Solidity: event AssetUpdated(uint256 index, uint256 value)
func (_Farming *FarmingFilterer) ParseAssetUpdated(log types.Log) (*FarmingAssetUpdated, error) {
	event := new(FarmingAssetUpdated)
	if err := _Farming.contract.UnpackLog(event, "AssetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmingBuyNFTRequestIterator is returned from FilterBuyNFTRequest and is used to iterate over the raw logs and unpacked data for BuyNFTRequest events raised by the Farming contract.
type FarmingBuyNFTRequestIterator struct {
	Event *FarmingBuyNFTRequest // Event containing the contract specifics and raw log

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
func (it *FarmingBuyNFTRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmingBuyNFTRequest)
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
		it.Event = new(FarmingBuyNFTRequest)
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
func (it *FarmingBuyNFTRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmingBuyNFTRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmingBuyNFTRequest represents a BuyNFTRequest event raised by the Farming contract.
type FarmingBuyNFTRequest struct {
	Buyer         common.Address
	Amount        *big.Int
	TransactionId *big.Int
	Timestamp     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBuyNFTRequest is a free log retrieval operation binding the contract event 0x9c4a57a2d23e491271812302c1880d053c5930a97fdb06e917684d82435d6e9d.
//
// Solidity: event BuyNFTRequest(address buyer, uint256 amount, uint256 transactionId, uint256 timestamp)
func (_Farming *FarmingFilterer) FilterBuyNFTRequest(opts *bind.FilterOpts) (*FarmingBuyNFTRequestIterator, error) {

	logs, sub, err := _Farming.contract.FilterLogs(opts, "BuyNFTRequest")
	if err != nil {
		return nil, err
	}
	return &FarmingBuyNFTRequestIterator{contract: _Farming.contract, event: "BuyNFTRequest", logs: logs, sub: sub}, nil
}

// WatchBuyNFTRequest is a free log subscription operation binding the contract event 0x9c4a57a2d23e491271812302c1880d053c5930a97fdb06e917684d82435d6e9d.
//
// Solidity: event BuyNFTRequest(address buyer, uint256 amount, uint256 transactionId, uint256 timestamp)
func (_Farming *FarmingFilterer) WatchBuyNFTRequest(opts *bind.WatchOpts, sink chan<- *FarmingBuyNFTRequest) (event.Subscription, error) {

	logs, sub, err := _Farming.contract.WatchLogs(opts, "BuyNFTRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmingBuyNFTRequest)
				if err := _Farming.contract.UnpackLog(event, "BuyNFTRequest", log); err != nil {
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

// ParseBuyNFTRequest is a log parse operation binding the contract event 0x9c4a57a2d23e491271812302c1880d053c5930a97fdb06e917684d82435d6e9d.
//
// Solidity: event BuyNFTRequest(address buyer, uint256 amount, uint256 transactionId, uint256 timestamp)
func (_Farming *FarmingFilterer) ParseBuyNFTRequest(log types.Log) (*FarmingBuyNFTRequest, error) {
	event := new(FarmingBuyNFTRequest)
	if err := _Farming.contract.UnpackLog(event, "BuyNFTRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmingCancelBuyNFTIterator is returned from FilterCancelBuyNFT and is used to iterate over the raw logs and unpacked data for CancelBuyNFT events raised by the Farming contract.
type FarmingCancelBuyNFTIterator struct {
	Event *FarmingCancelBuyNFT // Event containing the contract specifics and raw log

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
func (it *FarmingCancelBuyNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmingCancelBuyNFT)
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
		it.Event = new(FarmingCancelBuyNFT)
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
func (it *FarmingCancelBuyNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmingCancelBuyNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmingCancelBuyNFT represents a CancelBuyNFT event raised by the Farming contract.
type FarmingCancelBuyNFT struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCancelBuyNFT is a free log retrieval operation binding the contract event 0x79b35cecb42a71843efa68708196bd251e8642787a9bf446fb1545cb93eb3c8a.
//
// Solidity: event CancelBuyNFT(uint256 transactionId)
func (_Farming *FarmingFilterer) FilterCancelBuyNFT(opts *bind.FilterOpts) (*FarmingCancelBuyNFTIterator, error) {

	logs, sub, err := _Farming.contract.FilterLogs(opts, "CancelBuyNFT")
	if err != nil {
		return nil, err
	}
	return &FarmingCancelBuyNFTIterator{contract: _Farming.contract, event: "CancelBuyNFT", logs: logs, sub: sub}, nil
}

// WatchCancelBuyNFT is a free log subscription operation binding the contract event 0x79b35cecb42a71843efa68708196bd251e8642787a9bf446fb1545cb93eb3c8a.
//
// Solidity: event CancelBuyNFT(uint256 transactionId)
func (_Farming *FarmingFilterer) WatchCancelBuyNFT(opts *bind.WatchOpts, sink chan<- *FarmingCancelBuyNFT) (event.Subscription, error) {

	logs, sub, err := _Farming.contract.WatchLogs(opts, "CancelBuyNFT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmingCancelBuyNFT)
				if err := _Farming.contract.UnpackLog(event, "CancelBuyNFT", log); err != nil {
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

// ParseCancelBuyNFT is a log parse operation binding the contract event 0x79b35cecb42a71843efa68708196bd251e8642787a9bf446fb1545cb93eb3c8a.
//
// Solidity: event CancelBuyNFT(uint256 transactionId)
func (_Farming *FarmingFilterer) ParseCancelBuyNFT(log types.Log) (*FarmingCancelBuyNFT, error) {
	event := new(FarmingCancelBuyNFT)
	if err := _Farming.contract.UnpackLog(event, "CancelBuyNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmingCancelSellNFTIterator is returned from FilterCancelSellNFT and is used to iterate over the raw logs and unpacked data for CancelSellNFT events raised by the Farming contract.
type FarmingCancelSellNFTIterator struct {
	Event *FarmingCancelSellNFT // Event containing the contract specifics and raw log

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
func (it *FarmingCancelSellNFTIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmingCancelSellNFT)
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
		it.Event = new(FarmingCancelSellNFT)
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
func (it *FarmingCancelSellNFTIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmingCancelSellNFTIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmingCancelSellNFT represents a CancelSellNFT event raised by the Farming contract.
type FarmingCancelSellNFT struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCancelSellNFT is a free log retrieval operation binding the contract event 0xc2c3f52c8ed415c00c97e91f5621f1506edff3108b0749037c71312ac3dec64a.
//
// Solidity: event CancelSellNFT(uint256 transactionId)
func (_Farming *FarmingFilterer) FilterCancelSellNFT(opts *bind.FilterOpts) (*FarmingCancelSellNFTIterator, error) {

	logs, sub, err := _Farming.contract.FilterLogs(opts, "CancelSellNFT")
	if err != nil {
		return nil, err
	}
	return &FarmingCancelSellNFTIterator{contract: _Farming.contract, event: "CancelSellNFT", logs: logs, sub: sub}, nil
}

// WatchCancelSellNFT is a free log subscription operation binding the contract event 0xc2c3f52c8ed415c00c97e91f5621f1506edff3108b0749037c71312ac3dec64a.
//
// Solidity: event CancelSellNFT(uint256 transactionId)
func (_Farming *FarmingFilterer) WatchCancelSellNFT(opts *bind.WatchOpts, sink chan<- *FarmingCancelSellNFT) (event.Subscription, error) {

	logs, sub, err := _Farming.contract.WatchLogs(opts, "CancelSellNFT")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmingCancelSellNFT)
				if err := _Farming.contract.UnpackLog(event, "CancelSellNFT", log); err != nil {
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

// ParseCancelSellNFT is a log parse operation binding the contract event 0xc2c3f52c8ed415c00c97e91f5621f1506edff3108b0749037c71312ac3dec64a.
//
// Solidity: event CancelSellNFT(uint256 transactionId)
func (_Farming *FarmingFilterer) ParseCancelSellNFT(log types.Log) (*FarmingCancelSellNFT, error) {
	event := new(FarmingCancelSellNFT)
	if err := _Farming.contract.UnpackLog(event, "CancelSellNFT", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmingDebtAddedIterator is returned from FilterDebtAdded and is used to iterate over the raw logs and unpacked data for DebtAdded events raised by the Farming contract.
type FarmingDebtAddedIterator struct {
	Event *FarmingDebtAdded // Event containing the contract specifics and raw log

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
func (it *FarmingDebtAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmingDebtAdded)
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
		it.Event = new(FarmingDebtAdded)
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
func (it *FarmingDebtAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmingDebtAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmingDebtAdded represents a DebtAdded event raised by the Farming contract.
type FarmingDebtAdded struct {
	Name  string
	Code  string
	Value *big.Int
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDebtAdded is a free log retrieval operation binding the contract event 0x222bc523bb6981e25032bfc7daf8772951a49e677b4b5281b9c9667eaa7f3950.
//
// Solidity: event DebtAdded(string name, string code, uint256 value, uint256 index)
func (_Farming *FarmingFilterer) FilterDebtAdded(opts *bind.FilterOpts) (*FarmingDebtAddedIterator, error) {

	logs, sub, err := _Farming.contract.FilterLogs(opts, "DebtAdded")
	if err != nil {
		return nil, err
	}
	return &FarmingDebtAddedIterator{contract: _Farming.contract, event: "DebtAdded", logs: logs, sub: sub}, nil
}

// WatchDebtAdded is a free log subscription operation binding the contract event 0x222bc523bb6981e25032bfc7daf8772951a49e677b4b5281b9c9667eaa7f3950.
//
// Solidity: event DebtAdded(string name, string code, uint256 value, uint256 index)
func (_Farming *FarmingFilterer) WatchDebtAdded(opts *bind.WatchOpts, sink chan<- *FarmingDebtAdded) (event.Subscription, error) {

	logs, sub, err := _Farming.contract.WatchLogs(opts, "DebtAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmingDebtAdded)
				if err := _Farming.contract.UnpackLog(event, "DebtAdded", log); err != nil {
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

// ParseDebtAdded is a log parse operation binding the contract event 0x222bc523bb6981e25032bfc7daf8772951a49e677b4b5281b9c9667eaa7f3950.
//
// Solidity: event DebtAdded(string name, string code, uint256 value, uint256 index)
func (_Farming *FarmingFilterer) ParseDebtAdded(log types.Log) (*FarmingDebtAdded, error) {
	event := new(FarmingDebtAdded)
	if err := _Farming.contract.UnpackLog(event, "DebtAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmingDebtUpdatedIterator is returned from FilterDebtUpdated and is used to iterate over the raw logs and unpacked data for DebtUpdated events raised by the Farming contract.
type FarmingDebtUpdatedIterator struct {
	Event *FarmingDebtUpdated // Event containing the contract specifics and raw log

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
func (it *FarmingDebtUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmingDebtUpdated)
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
		it.Event = new(FarmingDebtUpdated)
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
func (it *FarmingDebtUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmingDebtUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmingDebtUpdated represents a DebtUpdated event raised by the Farming contract.
type FarmingDebtUpdated struct {
	Index *big.Int
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDebtUpdated is a free log retrieval operation binding the contract event 0x37b69423fc1031ce118998719845c41735dd96f6be8cd19cc670f3416764f873.
//
// Solidity: event DebtUpdated(uint256 index, uint256 value)
func (_Farming *FarmingFilterer) FilterDebtUpdated(opts *bind.FilterOpts) (*FarmingDebtUpdatedIterator, error) {

	logs, sub, err := _Farming.contract.FilterLogs(opts, "DebtUpdated")
	if err != nil {
		return nil, err
	}
	return &FarmingDebtUpdatedIterator{contract: _Farming.contract, event: "DebtUpdated", logs: logs, sub: sub}, nil
}

// WatchDebtUpdated is a free log subscription operation binding the contract event 0x37b69423fc1031ce118998719845c41735dd96f6be8cd19cc670f3416764f873.
//
// Solidity: event DebtUpdated(uint256 index, uint256 value)
func (_Farming *FarmingFilterer) WatchDebtUpdated(opts *bind.WatchOpts, sink chan<- *FarmingDebtUpdated) (event.Subscription, error) {

	logs, sub, err := _Farming.contract.WatchLogs(opts, "DebtUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmingDebtUpdated)
				if err := _Farming.contract.UnpackLog(event, "DebtUpdated", log); err != nil {
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

// ParseDebtUpdated is a log parse operation binding the contract event 0x37b69423fc1031ce118998719845c41735dd96f6be8cd19cc670f3416764f873.
//
// Solidity: event DebtUpdated(uint256 index, uint256 value)
func (_Farming *FarmingFilterer) ParseDebtUpdated(log types.Log) (*FarmingDebtUpdated, error) {
	event := new(FarmingDebtUpdated)
	if err := _Farming.contract.UnpackLog(event, "DebtUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmingEditorUpdatedIterator is returned from FilterEditorUpdated and is used to iterate over the raw logs and unpacked data for EditorUpdated events raised by the Farming contract.
type FarmingEditorUpdatedIterator struct {
	Event *FarmingEditorUpdated // Event containing the contract specifics and raw log

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
func (it *FarmingEditorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmingEditorUpdated)
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
		it.Event = new(FarmingEditorUpdated)
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
func (it *FarmingEditorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmingEditorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmingEditorUpdated represents a EditorUpdated event raised by the Farming contract.
type FarmingEditorUpdated struct {
	NewEditor common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEditorUpdated is a free log retrieval operation binding the contract event 0x72ee7959ecf19b763a256ecb0108ff6d89e422718b073f85591f40bb4da32456.
//
// Solidity: event EditorUpdated(address newEditor)
func (_Farming *FarmingFilterer) FilterEditorUpdated(opts *bind.FilterOpts) (*FarmingEditorUpdatedIterator, error) {

	logs, sub, err := _Farming.contract.FilterLogs(opts, "EditorUpdated")
	if err != nil {
		return nil, err
	}
	return &FarmingEditorUpdatedIterator{contract: _Farming.contract, event: "EditorUpdated", logs: logs, sub: sub}, nil
}

// WatchEditorUpdated is a free log subscription operation binding the contract event 0x72ee7959ecf19b763a256ecb0108ff6d89e422718b073f85591f40bb4da32456.
//
// Solidity: event EditorUpdated(address newEditor)
func (_Farming *FarmingFilterer) WatchEditorUpdated(opts *bind.WatchOpts, sink chan<- *FarmingEditorUpdated) (event.Subscription, error) {

	logs, sub, err := _Farming.contract.WatchLogs(opts, "EditorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmingEditorUpdated)
				if err := _Farming.contract.UnpackLog(event, "EditorUpdated", log); err != nil {
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

// ParseEditorUpdated is a log parse operation binding the contract event 0x72ee7959ecf19b763a256ecb0108ff6d89e422718b073f85591f40bb4da32456.
//
// Solidity: event EditorUpdated(address newEditor)
func (_Farming *FarmingFilterer) ParseEditorUpdated(log types.Log) (*FarmingEditorUpdated, error) {
	event := new(FarmingEditorUpdated)
	if err := _Farming.contract.UnpackLog(event, "EditorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmingNAVUpdatedIterator is returned from FilterNAVUpdated and is used to iterate over the raw logs and unpacked data for NAVUpdated events raised by the Farming contract.
type FarmingNAVUpdatedIterator struct {
	Event *FarmingNAVUpdated // Event containing the contract specifics and raw log

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
func (it *FarmingNAVUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmingNAVUpdated)
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
		it.Event = new(FarmingNAVUpdated)
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
func (it *FarmingNAVUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmingNAVUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmingNAVUpdated represents a NAVUpdated event raised by the Farming contract.
type FarmingNAVUpdated struct {
	NewNAV    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNAVUpdated is a free log retrieval operation binding the contract event 0x224bbf4ba043354bee6487d05aa90892db4426afe8e3aa79007d0eae1b985ebc.
//
// Solidity: event NAVUpdated(uint256 newNAV, uint256 timestamp)
func (_Farming *FarmingFilterer) FilterNAVUpdated(opts *bind.FilterOpts) (*FarmingNAVUpdatedIterator, error) {

	logs, sub, err := _Farming.contract.FilterLogs(opts, "NAVUpdated")
	if err != nil {
		return nil, err
	}
	return &FarmingNAVUpdatedIterator{contract: _Farming.contract, event: "NAVUpdated", logs: logs, sub: sub}, nil
}

// WatchNAVUpdated is a free log subscription operation binding the contract event 0x224bbf4ba043354bee6487d05aa90892db4426afe8e3aa79007d0eae1b985ebc.
//
// Solidity: event NAVUpdated(uint256 newNAV, uint256 timestamp)
func (_Farming *FarmingFilterer) WatchNAVUpdated(opts *bind.WatchOpts, sink chan<- *FarmingNAVUpdated) (event.Subscription, error) {

	logs, sub, err := _Farming.contract.WatchLogs(opts, "NAVUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmingNAVUpdated)
				if err := _Farming.contract.UnpackLog(event, "NAVUpdated", log); err != nil {
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

// ParseNAVUpdated is a log parse operation binding the contract event 0x224bbf4ba043354bee6487d05aa90892db4426afe8e3aa79007d0eae1b985ebc.
//
// Solidity: event NAVUpdated(uint256 newNAV, uint256 timestamp)
func (_Farming *FarmingFilterer) ParseNAVUpdated(log types.Log) (*FarmingNAVUpdated, error) {
	event := new(FarmingNAVUpdated)
	if err := _Farming.contract.UnpackLog(event, "NAVUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmingNFTMintedIterator is returned from FilterNFTMinted and is used to iterate over the raw logs and unpacked data for NFTMinted events raised by the Farming contract.
type FarmingNFTMintedIterator struct {
	Event *FarmingNFTMinted // Event containing the contract specifics and raw log

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
func (it *FarmingNFTMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmingNFTMinted)
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
		it.Event = new(FarmingNFTMinted)
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
func (it *FarmingNFTMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmingNFTMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmingNFTMinted represents a NFTMinted event raised by the Farming contract.
type FarmingNFTMinted struct {
	Buyer     common.Address
	Price     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNFTMinted is a free log retrieval operation binding the contract event 0x3a8a89b59a31c39a36febecb987e0657ab7b7c73b60ebacb44dcb9886c2d5c8a.
//
// Solidity: event NFTMinted(address indexed buyer, uint256 price, uint256 timestamp)
func (_Farming *FarmingFilterer) FilterNFTMinted(opts *bind.FilterOpts, buyer []common.Address) (*FarmingNFTMintedIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Farming.contract.FilterLogs(opts, "NFTMinted", buyerRule)
	if err != nil {
		return nil, err
	}
	return &FarmingNFTMintedIterator{contract: _Farming.contract, event: "NFTMinted", logs: logs, sub: sub}, nil
}

// WatchNFTMinted is a free log subscription operation binding the contract event 0x3a8a89b59a31c39a36febecb987e0657ab7b7c73b60ebacb44dcb9886c2d5c8a.
//
// Solidity: event NFTMinted(address indexed buyer, uint256 price, uint256 timestamp)
func (_Farming *FarmingFilterer) WatchNFTMinted(opts *bind.WatchOpts, sink chan<- *FarmingNFTMinted, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Farming.contract.WatchLogs(opts, "NFTMinted", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmingNFTMinted)
				if err := _Farming.contract.UnpackLog(event, "NFTMinted", log); err != nil {
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

// ParseNFTMinted is a log parse operation binding the contract event 0x3a8a89b59a31c39a36febecb987e0657ab7b7c73b60ebacb44dcb9886c2d5c8a.
//
// Solidity: event NFTMinted(address indexed buyer, uint256 price, uint256 timestamp)
func (_Farming *FarmingFilterer) ParseNFTMinted(log types.Log) (*FarmingNFTMinted, error) {
	event := new(FarmingNFTMinted)
	if err := _Farming.contract.UnpackLog(event, "NFTMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmingPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Farming contract.
type FarmingPausedIterator struct {
	Event *FarmingPaused // Event containing the contract specifics and raw log

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
func (it *FarmingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmingPaused)
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
		it.Event = new(FarmingPaused)
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
func (it *FarmingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmingPaused represents a Paused event raised by the Farming contract.
type FarmingPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Farming *FarmingFilterer) FilterPaused(opts *bind.FilterOpts) (*FarmingPausedIterator, error) {

	logs, sub, err := _Farming.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &FarmingPausedIterator{contract: _Farming.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Farming *FarmingFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *FarmingPaused) (event.Subscription, error) {

	logs, sub, err := _Farming.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmingPaused)
				if err := _Farming.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Farming *FarmingFilterer) ParsePaused(log types.Log) (*FarmingPaused, error) {
	event := new(FarmingPaused)
	if err := _Farming.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmingPendingBuyNFTProcessedIterator is returned from FilterPendingBuyNFTProcessed and is used to iterate over the raw logs and unpacked data for PendingBuyNFTProcessed events raised by the Farming contract.
type FarmingPendingBuyNFTProcessedIterator struct {
	Event *FarmingPendingBuyNFTProcessed // Event containing the contract specifics and raw log

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
func (it *FarmingPendingBuyNFTProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmingPendingBuyNFTProcessed)
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
		it.Event = new(FarmingPendingBuyNFTProcessed)
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
func (it *FarmingPendingBuyNFTProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmingPendingBuyNFTProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmingPendingBuyNFTProcessed represents a PendingBuyNFTProcessed event raised by the Farming contract.
type FarmingPendingBuyNFTProcessed struct {
	TransactionId *big.Int
	Timestamp     *big.Int
	Refund        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPendingBuyNFTProcessed is a free log retrieval operation binding the contract event 0xca7c2d9fc29e2c238661a7b15ca26db54f140cdef084ebb8ca6269fe882575ba.
//
// Solidity: event PendingBuyNFTProcessed(uint256 transactionId, uint256 timestamp, uint256 refund)
func (_Farming *FarmingFilterer) FilterPendingBuyNFTProcessed(opts *bind.FilterOpts) (*FarmingPendingBuyNFTProcessedIterator, error) {

	logs, sub, err := _Farming.contract.FilterLogs(opts, "PendingBuyNFTProcessed")
	if err != nil {
		return nil, err
	}
	return &FarmingPendingBuyNFTProcessedIterator{contract: _Farming.contract, event: "PendingBuyNFTProcessed", logs: logs, sub: sub}, nil
}

// WatchPendingBuyNFTProcessed is a free log subscription operation binding the contract event 0xca7c2d9fc29e2c238661a7b15ca26db54f140cdef084ebb8ca6269fe882575ba.
//
// Solidity: event PendingBuyNFTProcessed(uint256 transactionId, uint256 timestamp, uint256 refund)
func (_Farming *FarmingFilterer) WatchPendingBuyNFTProcessed(opts *bind.WatchOpts, sink chan<- *FarmingPendingBuyNFTProcessed) (event.Subscription, error) {

	logs, sub, err := _Farming.contract.WatchLogs(opts, "PendingBuyNFTProcessed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmingPendingBuyNFTProcessed)
				if err := _Farming.contract.UnpackLog(event, "PendingBuyNFTProcessed", log); err != nil {
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

// ParsePendingBuyNFTProcessed is a log parse operation binding the contract event 0xca7c2d9fc29e2c238661a7b15ca26db54f140cdef084ebb8ca6269fe882575ba.
//
// Solidity: event PendingBuyNFTProcessed(uint256 transactionId, uint256 timestamp, uint256 refund)
func (_Farming *FarmingFilterer) ParsePendingBuyNFTProcessed(log types.Log) (*FarmingPendingBuyNFTProcessed, error) {
	event := new(FarmingPendingBuyNFTProcessed)
	if err := _Farming.contract.UnpackLog(event, "PendingBuyNFTProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmingPendingSellCreatedIterator is returned from FilterPendingSellCreated and is used to iterate over the raw logs and unpacked data for PendingSellCreated events raised by the Farming contract.
type FarmingPendingSellCreatedIterator struct {
	Event *FarmingPendingSellCreated // Event containing the contract specifics and raw log

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
func (it *FarmingPendingSellCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmingPendingSellCreated)
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
		it.Event = new(FarmingPendingSellCreated)
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
func (it *FarmingPendingSellCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmingPendingSellCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmingPendingSellCreated represents a PendingSellCreated event raised by the Farming contract.
type FarmingPendingSellCreated struct {
	SellId    *big.Int
	TokenId   *big.Int
	Seller    common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPendingSellCreated is a free log retrieval operation binding the contract event 0xdbec9f79d3cb02939ad2f4b733a4321701b706969ac6c32a6a538a172628ae40.
//
// Solidity: event PendingSellCreated(uint256 sellId, uint256 tokenId, address seller, uint256 timestamp)
func (_Farming *FarmingFilterer) FilterPendingSellCreated(opts *bind.FilterOpts) (*FarmingPendingSellCreatedIterator, error) {

	logs, sub, err := _Farming.contract.FilterLogs(opts, "PendingSellCreated")
	if err != nil {
		return nil, err
	}
	return &FarmingPendingSellCreatedIterator{contract: _Farming.contract, event: "PendingSellCreated", logs: logs, sub: sub}, nil
}

// WatchPendingSellCreated is a free log subscription operation binding the contract event 0xdbec9f79d3cb02939ad2f4b733a4321701b706969ac6c32a6a538a172628ae40.
//
// Solidity: event PendingSellCreated(uint256 sellId, uint256 tokenId, address seller, uint256 timestamp)
func (_Farming *FarmingFilterer) WatchPendingSellCreated(opts *bind.WatchOpts, sink chan<- *FarmingPendingSellCreated) (event.Subscription, error) {

	logs, sub, err := _Farming.contract.WatchLogs(opts, "PendingSellCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmingPendingSellCreated)
				if err := _Farming.contract.UnpackLog(event, "PendingSellCreated", log); err != nil {
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

// ParsePendingSellCreated is a log parse operation binding the contract event 0xdbec9f79d3cb02939ad2f4b733a4321701b706969ac6c32a6a538a172628ae40.
//
// Solidity: event PendingSellCreated(uint256 sellId, uint256 tokenId, address seller, uint256 timestamp)
func (_Farming *FarmingFilterer) ParsePendingSellCreated(log types.Log) (*FarmingPendingSellCreated, error) {
	event := new(FarmingPendingSellCreated)
	if err := _Farming.contract.UnpackLog(event, "PendingSellCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmingPendingSellNFTProcessedIterator is returned from FilterPendingSellNFTProcessed and is used to iterate over the raw logs and unpacked data for PendingSellNFTProcessed events raised by the Farming contract.
type FarmingPendingSellNFTProcessedIterator struct {
	Event *FarmingPendingSellNFTProcessed // Event containing the contract specifics and raw log

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
func (it *FarmingPendingSellNFTProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmingPendingSellNFTProcessed)
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
		it.Event = new(FarmingPendingSellNFTProcessed)
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
func (it *FarmingPendingSellNFTProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmingPendingSellNFTProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmingPendingSellNFTProcessed represents a PendingSellNFTProcessed event raised by the Farming contract.
type FarmingPendingSellNFTProcessed struct {
	SellId    *big.Int
	TokenId   *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPendingSellNFTProcessed is a free log retrieval operation binding the contract event 0x2eee9ff196680fe6873dd18bae2874f088da335081f4cbe911b8b0482823940b.
//
// Solidity: event PendingSellNFTProcessed(uint256 sellId, uint256 tokenId, uint256 timestamp)
func (_Farming *FarmingFilterer) FilterPendingSellNFTProcessed(opts *bind.FilterOpts) (*FarmingPendingSellNFTProcessedIterator, error) {

	logs, sub, err := _Farming.contract.FilterLogs(opts, "PendingSellNFTProcessed")
	if err != nil {
		return nil, err
	}
	return &FarmingPendingSellNFTProcessedIterator{contract: _Farming.contract, event: "PendingSellNFTProcessed", logs: logs, sub: sub}, nil
}

// WatchPendingSellNFTProcessed is a free log subscription operation binding the contract event 0x2eee9ff196680fe6873dd18bae2874f088da335081f4cbe911b8b0482823940b.
//
// Solidity: event PendingSellNFTProcessed(uint256 sellId, uint256 tokenId, uint256 timestamp)
func (_Farming *FarmingFilterer) WatchPendingSellNFTProcessed(opts *bind.WatchOpts, sink chan<- *FarmingPendingSellNFTProcessed) (event.Subscription, error) {

	logs, sub, err := _Farming.contract.WatchLogs(opts, "PendingSellNFTProcessed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmingPendingSellNFTProcessed)
				if err := _Farming.contract.UnpackLog(event, "PendingSellNFTProcessed", log); err != nil {
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

// ParsePendingSellNFTProcessed is a log parse operation binding the contract event 0x2eee9ff196680fe6873dd18bae2874f088da335081f4cbe911b8b0482823940b.
//
// Solidity: event PendingSellNFTProcessed(uint256 sellId, uint256 tokenId, uint256 timestamp)
func (_Farming *FarmingFilterer) ParsePendingSellNFTProcessed(log types.Log) (*FarmingPendingSellNFTProcessed, error) {
	event := new(FarmingPendingSellNFTProcessed)
	if err := _Farming.contract.UnpackLog(event, "PendingSellNFTProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmingRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Farming contract.
type FarmingRoleAdminChangedIterator struct {
	Event *FarmingRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *FarmingRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmingRoleAdminChanged)
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
		it.Event = new(FarmingRoleAdminChanged)
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
func (it *FarmingRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmingRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmingRoleAdminChanged represents a RoleAdminChanged event raised by the Farming contract.
type FarmingRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Farming *FarmingFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*FarmingRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Farming.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &FarmingRoleAdminChangedIterator{contract: _Farming.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Farming *FarmingFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *FarmingRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Farming.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmingRoleAdminChanged)
				if err := _Farming.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Farming *FarmingFilterer) ParseRoleAdminChanged(log types.Log) (*FarmingRoleAdminChanged, error) {
	event := new(FarmingRoleAdminChanged)
	if err := _Farming.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmingRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Farming contract.
type FarmingRoleGrantedIterator struct {
	Event *FarmingRoleGranted // Event containing the contract specifics and raw log

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
func (it *FarmingRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmingRoleGranted)
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
		it.Event = new(FarmingRoleGranted)
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
func (it *FarmingRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmingRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmingRoleGranted represents a RoleGranted event raised by the Farming contract.
type FarmingRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Farming *FarmingFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FarmingRoleGrantedIterator, error) {

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

	logs, sub, err := _Farming.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FarmingRoleGrantedIterator{contract: _Farming.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Farming *FarmingFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *FarmingRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Farming.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmingRoleGranted)
				if err := _Farming.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Farming *FarmingFilterer) ParseRoleGranted(log types.Log) (*FarmingRoleGranted, error) {
	event := new(FarmingRoleGranted)
	if err := _Farming.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmingRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Farming contract.
type FarmingRoleRevokedIterator struct {
	Event *FarmingRoleRevoked // Event containing the contract specifics and raw log

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
func (it *FarmingRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmingRoleRevoked)
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
		it.Event = new(FarmingRoleRevoked)
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
func (it *FarmingRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmingRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmingRoleRevoked represents a RoleRevoked event raised by the Farming contract.
type FarmingRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Farming *FarmingFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FarmingRoleRevokedIterator, error) {

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

	logs, sub, err := _Farming.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FarmingRoleRevokedIterator{contract: _Farming.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Farming *FarmingFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *FarmingRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Farming.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmingRoleRevoked)
				if err := _Farming.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Farming *FarmingFilterer) ParseRoleRevoked(log types.Log) (*FarmingRoleRevoked, error) {
	event := new(FarmingRoleRevoked)
	if err := _Farming.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmingUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Farming contract.
type FarmingUnpausedIterator struct {
	Event *FarmingUnpaused // Event containing the contract specifics and raw log

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
func (it *FarmingUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmingUnpaused)
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
		it.Event = new(FarmingUnpaused)
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
func (it *FarmingUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmingUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmingUnpaused represents a Unpaused event raised by the Farming contract.
type FarmingUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Farming *FarmingFilterer) FilterUnpaused(opts *bind.FilterOpts) (*FarmingUnpausedIterator, error) {

	logs, sub, err := _Farming.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &FarmingUnpausedIterator{contract: _Farming.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Farming *FarmingFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *FarmingUnpaused) (event.Subscription, error) {

	logs, sub, err := _Farming.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmingUnpaused)
				if err := _Farming.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Farming *FarmingFilterer) ParseUnpaused(log types.Log) (*FarmingUnpaused, error) {
	event := new(FarmingUnpaused)
	if err := _Farming.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
