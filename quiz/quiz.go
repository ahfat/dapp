// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package quiz

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

// QuizMetaData contains all meta data concerning the Quiz contract.
var QuizMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_qn\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"_ans\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"checkBoard\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"question\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ans\",\"type\":\"bytes32\"}],\"name\":\"sendAnswer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620009c0380380620009c0833981810160405281019062000037919062000226565b8160009081620000489190620004d7565b50806001819055505050620005be565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620000c18262000076565b810181811067ffffffffffffffff82111715620000e357620000e262000087565b5b80604052505050565b6000620000f862000058565b9050620001068282620000b6565b919050565b600067ffffffffffffffff82111562000129576200012862000087565b5b620001348262000076565b9050602081019050919050565b60005b838110156200016157808201518184015260208101905062000144565b60008484015250505050565b6000620001846200017e846200010b565b620000ec565b905082815260208101848484011115620001a357620001a262000071565b5b620001b084828562000141565b509392505050565b600082601f830112620001d057620001cf6200006c565b5b8151620001e28482602086016200016d565b91505092915050565b6000819050919050565b6200020081620001eb565b81146200020c57600080fd5b50565b6000815190506200022081620001f5565b92915050565b6000806040838503121562000240576200023f62000062565b5b600083015167ffffffffffffffff81111562000261576200026062000067565b5b6200026f85828601620001b8565b925050602062000282858286016200020f565b9150509250929050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620002df57607f821691505b602082108103620002f557620002f462000297565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026200035f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000320565b6200036b868362000320565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620003b8620003b2620003ac8462000383565b6200038d565b62000383565b9050919050565b6000819050919050565b620003d48362000397565b620003ec620003e382620003bf565b8484546200032d565b825550505050565b600090565b62000403620003f4565b62000410818484620003c9565b505050565b5b8181101562000438576200042c600082620003f9565b60018101905062000416565b5050565b601f82111562000487576200045181620002fb565b6200045c8462000310565b810160208510156200046c578190505b620004846200047b8562000310565b83018262000415565b50505b505050565b600082821c905092915050565b6000620004ac600019846008026200048c565b1980831691505092915050565b6000620004c7838362000499565b9150826002028217905092915050565b620004e2826200028c565b67ffffffffffffffff811115620004fe57620004fd62000087565b5b6200050a8254620002c6565b620005178282856200043c565b600060209050601f8311600181146200054f57600084156200053a578287015190505b620005468582620004b9565b865550620005b6565b601f1984166200055f86620002fb565b60005b82811015620005895784890151825560018201915060208501945060208101905062000562565b86831015620005a95784890151620005a5601f89168262000499565b8355505b6001600288020188555050505b505050505050565b6103f280620005ce6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806317d1653c146100465780633fad9ae01461007657806377f46bff14610094575b600080fd5b610060600480360381019061005b9190610247565b6100b2565b60405161006d919061028f565b60405180910390f35b61007e6100c8565b60405161008b919061033a565b60405180910390f35b61009c610156565b6040516100a9919061028f565b60405180910390f35b60006100c182600154146101aa565b9050919050565b600080546100d59061038b565b80601f01602080910402602001604051908101604052809291908181526020018280546101019061038b565b801561014e5780601f106101235761010080835404028352916020019161014e565b820191906000526020600020905b81548152906001019060200180831161013157829003601f168201915b505050505081565b6000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905090565b600081600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555060019050919050565b600080fd5b6000819050919050565b61022481610211565b811461022f57600080fd5b50565b6000813590506102418161021b565b92915050565b60006020828403121561025d5761025c61020c565b5b600061026b84828501610232565b91505092915050565b60008115159050919050565b61028981610274565b82525050565b60006020820190506102a46000830184610280565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156102e45780820151818401526020810190506102c9565b60008484015250505050565b6000601f19601f8301169050919050565b600061030c826102aa565b61031681856102b5565b93506103268185602086016102c6565b61032f816102f0565b840191505092915050565b600060208201905081810360008301526103548184610301565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806103a357607f821691505b6020821081036103b6576103b561035c565b5b5091905056fea2646970667358221220035b7c0be72d82b9ba63f188fc6d32b194a8501aaffdcadc9f78996539f8982964736f6c63430008130033",
}

// QuizABI is the input ABI used to generate the binding from.
// Deprecated: Use QuizMetaData.ABI instead.
var QuizABI = QuizMetaData.ABI

// QuizBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use QuizMetaData.Bin instead.
var QuizBin = QuizMetaData.Bin

// DeployQuiz deploys a new Ethereum contract, binding an instance of Quiz to it.
func DeployQuiz(auth *bind.TransactOpts, backend bind.ContractBackend, _qn string, _ans [32]byte) (common.Address, *types.Transaction, *Quiz, error) {
	parsed, err := QuizMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(QuizBin), backend, _qn, _ans)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Quiz{QuizCaller: QuizCaller{contract: contract}, QuizTransactor: QuizTransactor{contract: contract}, QuizFilterer: QuizFilterer{contract: contract}}, nil
}

// Quiz is an auto generated Go binding around an Ethereum contract.
type Quiz struct {
	QuizCaller     // Read-only binding to the contract
	QuizTransactor // Write-only binding to the contract
	QuizFilterer   // Log filterer for contract events
}

// QuizCaller is an auto generated read-only Go binding around an Ethereum contract.
type QuizCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuizTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QuizTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuizFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QuizFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuizSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QuizSession struct {
	Contract     *Quiz             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuizCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QuizCallerSession struct {
	Contract *QuizCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// QuizTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QuizTransactorSession struct {
	Contract     *QuizTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuizRaw is an auto generated low-level Go binding around an Ethereum contract.
type QuizRaw struct {
	Contract *Quiz // Generic contract binding to access the raw methods on
}

// QuizCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QuizCallerRaw struct {
	Contract *QuizCaller // Generic read-only contract binding to access the raw methods on
}

// QuizTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QuizTransactorRaw struct {
	Contract *QuizTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQuiz creates a new instance of Quiz, bound to a specific deployed contract.
func NewQuiz(address common.Address, backend bind.ContractBackend) (*Quiz, error) {
	contract, err := bindQuiz(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Quiz{QuizCaller: QuizCaller{contract: contract}, QuizTransactor: QuizTransactor{contract: contract}, QuizFilterer: QuizFilterer{contract: contract}}, nil
}

// NewQuizCaller creates a new read-only instance of Quiz, bound to a specific deployed contract.
func NewQuizCaller(address common.Address, caller bind.ContractCaller) (*QuizCaller, error) {
	contract, err := bindQuiz(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QuizCaller{contract: contract}, nil
}

// NewQuizTransactor creates a new write-only instance of Quiz, bound to a specific deployed contract.
func NewQuizTransactor(address common.Address, transactor bind.ContractTransactor) (*QuizTransactor, error) {
	contract, err := bindQuiz(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QuizTransactor{contract: contract}, nil
}

// NewQuizFilterer creates a new log filterer instance of Quiz, bound to a specific deployed contract.
func NewQuizFilterer(address common.Address, filterer bind.ContractFilterer) (*QuizFilterer, error) {
	contract, err := bindQuiz(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QuizFilterer{contract: contract}, nil
}

// bindQuiz binds a generic wrapper to an already deployed contract.
func bindQuiz(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := QuizMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Quiz *QuizRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Quiz.Contract.QuizCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Quiz *QuizRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quiz.Contract.QuizTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Quiz *QuizRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Quiz.Contract.QuizTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Quiz *QuizCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Quiz.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Quiz *QuizTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quiz.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Quiz *QuizTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Quiz.Contract.contract.Transact(opts, method, params...)
}

// CheckBoard is a free data retrieval call binding the contract method 0x77f46bff.
//
// Solidity: function checkBoard() view returns(bool)
func (_Quiz *QuizCaller) CheckBoard(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Quiz.contract.Call(opts, &out, "checkBoard")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckBoard is a free data retrieval call binding the contract method 0x77f46bff.
//
// Solidity: function checkBoard() view returns(bool)
func (_Quiz *QuizSession) CheckBoard() (bool, error) {
	return _Quiz.Contract.CheckBoard(&_Quiz.CallOpts)
}

// CheckBoard is a free data retrieval call binding the contract method 0x77f46bff.
//
// Solidity: function checkBoard() view returns(bool)
func (_Quiz *QuizCallerSession) CheckBoard() (bool, error) {
	return _Quiz.Contract.CheckBoard(&_Quiz.CallOpts)
}

// Question is a free data retrieval call binding the contract method 0x3fad9ae0.
//
// Solidity: function question() view returns(string)
func (_Quiz *QuizCaller) Question(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Quiz.contract.Call(opts, &out, "question")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Question is a free data retrieval call binding the contract method 0x3fad9ae0.
//
// Solidity: function question() view returns(string)
func (_Quiz *QuizSession) Question() (string, error) {
	return _Quiz.Contract.Question(&_Quiz.CallOpts)
}

// Question is a free data retrieval call binding the contract method 0x3fad9ae0.
//
// Solidity: function question() view returns(string)
func (_Quiz *QuizCallerSession) Question() (string, error) {
	return _Quiz.Contract.Question(&_Quiz.CallOpts)
}

// SendAnswer is a paid mutator transaction binding the contract method 0x17d1653c.
//
// Solidity: function sendAnswer(bytes32 _ans) returns(bool)
func (_Quiz *QuizTransactor) SendAnswer(opts *bind.TransactOpts, _ans [32]byte) (*types.Transaction, error) {
	return _Quiz.contract.Transact(opts, "sendAnswer", _ans)
}

// SendAnswer is a paid mutator transaction binding the contract method 0x17d1653c.
//
// Solidity: function sendAnswer(bytes32 _ans) returns(bool)
func (_Quiz *QuizSession) SendAnswer(_ans [32]byte) (*types.Transaction, error) {
	return _Quiz.Contract.SendAnswer(&_Quiz.TransactOpts, _ans)
}

// SendAnswer is a paid mutator transaction binding the contract method 0x17d1653c.
//
// Solidity: function sendAnswer(bytes32 _ans) returns(bool)
func (_Quiz *QuizTransactorSession) SendAnswer(_ans [32]byte) (*types.Transaction, error) {
	return _Quiz.Contract.SendAnswer(&_Quiz.TransactOpts, _ans)
}
