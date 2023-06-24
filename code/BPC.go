// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mytoken

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
)

// SSMCSPInfor is an auto generated low-level Go binding around an user-defined struct.
type SSMCSPInfor struct {
	DataId    *big.Int
	IP        [][1]byte
	ASP       common.Address
	Confirmed bool
}

// SSMCSellerData is an auto generated low-level Go binding around an user-defined struct.
type SSMCSellerData struct {
	Aseller            common.Address
	Description        string
	PartsNum           *big.Int
	HashOfParts        string
	HashOfEncryptParts string
	RandIndex          *big.Int
	Price              *big.Int
}

// BPCMetaData contains all meta data concerning the BPC contract.
var BPCMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dataId\",\"type\":\"uint256\"},{\"internalType\":\"bytes1[]\",\"name\":\"_IP\",\"type\":\"bytes1[]\"}],\"name\":\"BuyerReg\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"BMID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"SPIndex\",\"type\":\"uint256\"}],\"name\":\"Buyerconfirmed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_SSMCAddr\",\"type\":\"address\"}],\"name\":\"ControlContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_SSMCAddr\",\"type\":\"address\"}],\"name\":\"ControlContract2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_buyerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_SPID\",\"type\":\"uint256[]\"}],\"name\":\"SPReg\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_BMID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_SPIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_BMPayID\",\"type\":\"uint256\"}],\"name\":\"SPwithdraw\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"BMID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"SPIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_HODED\",\"type\":\"bytes32\"}],\"name\":\"SetHashOfDEData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"BMID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"SPIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"SPkey\",\"type\":\"bytes32\"}],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_BMPayID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"DEdata\",\"type\":\"string\"}],\"name\":\"appeal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"94221adf": "BuyerReg(uint256,bytes1[])",
		"e88a72ad": "Buyerconfirmed(uint256,uint256)",
		"8cbeb340": "ControlContract(address)",
		"388b08c6": "ControlContract2(address)",
		"dbec6377": "SPReg(uint256,uint256[])",
		"2575cf63": "SPwithdraw(uint256,uint256,uint256)",
		"00b17f72": "SetHashOfDEData(uint256,uint256,bytes32)",
		"7901de43": "activate(uint256,uint256,bytes32)",
		"9d7f722c": "appeal(uint256,string)",
		"8da5cb5b": "owner()",
	},
	Bin: "0x6080604052601460025561271060055534801561001b57600080fd5b50600080546001600160a01b031916331790556113858061003d6000396000f3fe6080604052600436106100905760003560e01c80638da5cb5b116100595780638da5cb5b1461020f57806394221adf1461024c5780639d7f722c1461025f578063dbec63771461027f578063e88a72ad146102ad57600080fd5b8062b17f72146100955780632575cf63146100b7578063388b08c6146100ca5780637901de43146100ea5780638cbeb340146100ca575b600080fd5b3480156100a157600080fd5b506100b56100b0366004610cd2565b6102cd565b005b6100b56100c5366004610cd2565b6103dd565b3480156100d657600080fd5b506100b56100e5366004610d16565b6105bc565b3480156100f657600080fd5b506100b5610105366004610cd2565b6040805160a08101825293845260208401928352429084019081526060840191825260006080850181815260068054600181018255925294517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f60059092029182015592517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d40840155517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d41830155517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d4282015590517ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d43909101805460ff1916911515919091179055565b34801561021b57600080fd5b5060005461022f906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b6100b561025a366004610e07565b6105f5565b34801561026b57600080fd5b506100b561027a366004610eda565b6107d8565b34801561028b57600080fd5b5061029f61029a366004610f5f565b6108d2565b604051908152602001610243565b3480156102b957600080fd5b506100b56102c8366004610fed565b610a82565b6001546004805433926001600160a01b03169163da18264091879081106102f6576102f661100f565b906000526020600020906004020160010185815481106103185761031861100f565b90600052602060002001546040518263ffffffff1660e01b815260040161034191815260200190565b600060405180830381865afa15801561035e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526103869190810190611045565b604001516001600160a01b03161461039d57600080fd5b80600484815481106103b1576103b161100f565b906000526020600020906004020160020183815481106103d3576103d361100f565b5060005250505050565b6001546004805433926001600160a01b03169163da18264091879081106104065761040661100f565b906000526020600020906004020160010185815481106104285761042861100f565b90600052602060002001546040518263ffffffff1660e01b815260040161045191815260200190565b600060405180830381865afa15801561046e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526104969190810190611045565b604001516001600160a01b0316146104ad57600080fd5b82600682815481106104c1576104c161100f565b906000526020600020906005020160000154146104dd57600080fd5b81600682815481106104f1576104f161100f565b9060005260206000209060050201600101541461050d57600080fd5b600681815481106105205761052061100f565b90600052602060002090600502016002015460055461053f9190611151565b42101561054b57600080fd5b6000600682815481106105605761056061100f565b60009182526020909120600590910201600401805460ff191682151517905561058857600080fd5b600254604051339180156108fc02916000818181858888f193505050501580156105b6573d6000803e3d6000fd5b50505050565b6000546001600160a01b031633146105d357600080fd5b600180546001600160a01b0319166001600160a01b0392909216919091179055565b6002546001546040516315c96b3960e31b8152600481018590523492916001600160a01b03169063ae4b59c8906024016000604051808303816000875af1158015610644573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261066c91908101906111e2565b6040015161067a91906112c9565b6001546040516315c96b3960e31b8152600481018690526001600160a01b039091169063ae4b59c8906024016000604051808303816000875af11580156106c5573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526106ed91908101906111e2565b60c001516106fb9190611151565b111561070657600080fd5b604080516080810182528381526020808201848152339383019390935234606083015260038054600181018255600091909152825160049091027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b810191825593518051939491936107a1937fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85c909301929190910190610b76565b5060408201516002820180546001600160a01b0319166001600160a01b039092169190911790556060909101516003909101555050565b6004600683815481106107ed576107ed61100f565b9060005260206000209060050201600001548154811061080f5761080f61100f565b90600052602060002090600402016002016000815481106108325761083261100f565b906000526020600020015460028260405161084d91906112e8565b602060405180830381855afa15801561086a573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061088d9190611304565b1461089757600080fd5b6001600683815481106108ac576108ac61100f565b60009182526020909120600590910201600401805460ff19169115159190911790555050565b6000336001600160a01b0316600384815481106108f1576108f161100f565b60009182526020909120600260049092020101546001600160a01b03161461091857600080fd5b6109436040518060800160405280600081526020016060815260200160608152602001606081525090565b8381526020808201848152600480546001810182556000829052845191027f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b810191825591518051859492936109bf937f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19c909101920190610c1c565b50604082015180516109db916002840191602090910190610c1c565b50606082015180516109f7916003840191602090910190610c57565b50505060005b816020015151811015610a6a576000801b82604001518281518110610a2457610a2461100f565b602002602001018181525050600082606001518281518110610a4857610a4861100f565b9115156020928302919091019091015280610a628161131d565b9150506109fd565b600454610a7990600190611338565b95945050505050565b336001600160a01b0316600360048481548110610aa157610aa161100f565b90600052602060002090600402016000015481548110610ac357610ac361100f565b60009182526020909120600260049092020101546001600160a01b031614610aea57600080fd5b6000801b60048381548110610b0157610b0161100f565b90600052602060002090600402016002018281548110610b2357610b2361100f565b90600052602060002001541415610b3957600080fd5b60048281548110610b4c57610b4c61100f565b90600052602060002090600402016003018181548110610b6e57610b6e61100f565b506000525050565b82805482825590600052602060002090601f01602090048101928215610c0c5791602002820160005b83821115610bdd57835183826101000a81548160ff021916908360f81c02179055509260200192600101602081600001049283019260010302610b9f565b8015610c0a5782816101000a81549060ff0219169055600101602081600001049283019260010302610bdd565b505b50610c18929150610cbd565b5090565b828054828255906000526020600020908101928215610c0c579160200282015b82811115610c0c578251825591602001919060010190610c3c565b82805482825590600052602060002090601f01602090048101928215610c0c5791602002820160005b83821115610bdd57835183826101000a81548160ff0219169083151502179055509260200192600101602081600001049283019260010302610c80565b5b80821115610c185760008155600101610cbe565b600080600060608486031215610ce757600080fd5b505081359360208301359350604090920135919050565b6001600160a01b0381168114610d1357600080fd5b50565b600060208284031215610d2857600080fd5b8135610d3381610cfe565b9392505050565b634e487b7160e01b600052604160045260246000fd5b6040516080810167ffffffffffffffff81118282101715610d7357610d73610d3a565b60405290565b60405160e0810167ffffffffffffffff81118282101715610d7357610d73610d3a565b604051601f8201601f1916810167ffffffffffffffff81118282101715610dc557610dc5610d3a565b604052919050565b600067ffffffffffffffff821115610de757610de7610d3a565b5060051b60200190565b6001600160f81b031981168114610d1357600080fd5b60008060408385031215610e1a57600080fd5b8235915060208084013567ffffffffffffffff811115610e3957600080fd5b8401601f81018613610e4a57600080fd5b8035610e5d610e5882610dcd565b610d9c565b81815260059190911b82018301908381019088831115610e7c57600080fd5b928401925b82841015610ea3578335610e9481610df1565b82529284019290840190610e81565b80955050505050509250929050565b600067ffffffffffffffff821115610ecc57610ecc610d3a565b50601f01601f191660200190565b60008060408385031215610eed57600080fd5b82359150602083013567ffffffffffffffff811115610f0b57600080fd5b8301601f81018513610f1c57600080fd5b8035610f2a610e5882610eb2565b818152866020838501011115610f3f57600080fd5b816020840160208301376000602083830101528093505050509250929050565b60008060408385031215610f7257600080fd5b8235915060208084013567ffffffffffffffff811115610f9157600080fd5b8401601f81018613610fa257600080fd5b8035610fb0610e5882610dcd565b81815260059190911b82018301908381019088831115610fcf57600080fd5b928401925b82841015610ea357833582529284019290840190610fd4565b6000806040838503121561100057600080fd5b50508035926020909101359150565b634e487b7160e01b600052603260045260246000fd5b805161103081610cfe565b919050565b8051801515811461103057600080fd5b6000602080838503121561105857600080fd5b825167ffffffffffffffff8082111561107057600080fd5b908401906080828703121561108457600080fd5b61108c610d50565b8251815283830151828111156110a157600080fd5b83019150601f820187136110b457600080fd5b81516110c2610e5882610dcd565b81815260059190911b830185019085810190898311156110e157600080fd5b938601935b828510156111085784516110f981610df1565b825293860193908601906110e6565b838701525061111b905060408401611025565b604082015261112c60608401611035565b60608201529695505050505050565b634e487b7160e01b600052601160045260246000fd5b600082198211156111645761116461113b565b500190565b60005b8381101561118457818101518382015260200161116c565b838111156105b65750506000910152565b600082601f8301126111a657600080fd5b81516111b4610e5882610eb2565b8181528460208386010111156111c957600080fd5b6111da826020830160208701611169565b949350505050565b6000602082840312156111f457600080fd5b815167ffffffffffffffff8082111561120c57600080fd5b9083019060e0828603121561122057600080fd5b611228610d79565b61123183611025565b815260208301518281111561124557600080fd5b61125187828601611195565b6020830152506040830151604082015260608301518281111561127357600080fd5b61127f87828601611195565b60608301525060808301518281111561129757600080fd5b6112a387828601611195565b60808301525060a083015160a082015260c083015160c082015280935050505092915050565b60008160001904831182151516156112e3576112e361113b565b500290565b600082516112fa818460208701611169565b9190910192915050565b60006020828403121561131657600080fd5b5051919050565b60006000198214156113315761133161113b565b5060010190565b60008282101561134a5761134a61113b565b50039056fea26469706673582212204be566d685a125f9ca7fc02871e0a6445b05e67dd02351a162f853a948431db264736f6c634300080c0033",
}

// BPCABI is the input ABI used to generate the binding from.
// Deprecated: Use BPCMetaData.ABI instead.
var BPCABI = BPCMetaData.ABI

// Deprecated: Use BPCMetaData.Sigs instead.
// BPCFuncSigs maps the 4-byte function signature to its string representation.
var BPCFuncSigs = BPCMetaData.Sigs

// BPCBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BPCMetaData.Bin instead.
var BPCBin = BPCMetaData.Bin

// DeployBPC deploys a new Ethereum contract, binding an instance of BPC to it.
func DeployBPC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BPC, error) {
	parsed, err := BPCMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BPCBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BPC{BPCCaller: BPCCaller{contract: contract}, BPCTransactor: BPCTransactor{contract: contract}, BPCFilterer: BPCFilterer{contract: contract}}, nil
}

// BPC is an auto generated Go binding around an Ethereum contract.
type BPC struct {
	BPCCaller     // Read-only binding to the contract
	BPCTransactor // Write-only binding to the contract
	BPCFilterer   // Log filterer for contract events
}

// BPCCaller is an auto generated read-only Go binding around an Ethereum contract.
type BPCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BPCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BPCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BPCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BPCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BPCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BPCSession struct {
	Contract     *BPC              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BPCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BPCCallerSession struct {
	Contract *BPCCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BPCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BPCTransactorSession struct {
	Contract     *BPCTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BPCRaw is an auto generated low-level Go binding around an Ethereum contract.
type BPCRaw struct {
	Contract *BPC // Generic contract binding to access the raw methods on
}

// BPCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BPCCallerRaw struct {
	Contract *BPCCaller // Generic read-only contract binding to access the raw methods on
}

// BPCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BPCTransactorRaw struct {
	Contract *BPCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBPC creates a new instance of BPC, bound to a specific deployed contract.
func NewBPC(address common.Address, backend bind.ContractBackend) (*BPC, error) {
	contract, err := bindBPC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BPC{BPCCaller: BPCCaller{contract: contract}, BPCTransactor: BPCTransactor{contract: contract}, BPCFilterer: BPCFilterer{contract: contract}}, nil
}

// NewBPCCaller creates a new read-only instance of BPC, bound to a specific deployed contract.
func NewBPCCaller(address common.Address, caller bind.ContractCaller) (*BPCCaller, error) {
	contract, err := bindBPC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BPCCaller{contract: contract}, nil
}

// NewBPCTransactor creates a new write-only instance of BPC, bound to a specific deployed contract.
func NewBPCTransactor(address common.Address, transactor bind.ContractTransactor) (*BPCTransactor, error) {
	contract, err := bindBPC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BPCTransactor{contract: contract}, nil
}

// NewBPCFilterer creates a new log filterer instance of BPC, bound to a specific deployed contract.
func NewBPCFilterer(address common.Address, filterer bind.ContractFilterer) (*BPCFilterer, error) {
	contract, err := bindBPC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BPCFilterer{contract: contract}, nil
}

// bindBPC binds a generic wrapper to an already deployed contract.
func bindBPC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BPCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BPC *BPCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BPC.Contract.BPCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BPC *BPCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BPC.Contract.BPCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BPC *BPCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BPC.Contract.BPCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BPC *BPCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BPC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BPC *BPCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BPC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BPC *BPCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BPC.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BPC *BPCCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BPC.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BPC *BPCSession) Owner() (common.Address, error) {
	return _BPC.Contract.Owner(&_BPC.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BPC *BPCCallerSession) Owner() (common.Address, error) {
	return _BPC.Contract.Owner(&_BPC.CallOpts)
}

// BuyerReg is a paid mutator transaction binding the contract method 0x94221adf.
//
// Solidity: function BuyerReg(uint256 _dataId, bytes1[] _IP) payable returns()
func (_BPC *BPCTransactor) BuyerReg(opts *bind.TransactOpts, _dataId *big.Int, _IP [][1]byte) (*types.Transaction, error) {
	return _BPC.contract.Transact(opts, "BuyerReg", _dataId, _IP)
}

// BuyerReg is a paid mutator transaction binding the contract method 0x94221adf.
//
// Solidity: function BuyerReg(uint256 _dataId, bytes1[] _IP) payable returns()
func (_BPC *BPCSession) BuyerReg(_dataId *big.Int, _IP [][1]byte) (*types.Transaction, error) {
	return _BPC.Contract.BuyerReg(&_BPC.TransactOpts, _dataId, _IP)
}

// BuyerReg is a paid mutator transaction binding the contract method 0x94221adf.
//
// Solidity: function BuyerReg(uint256 _dataId, bytes1[] _IP) payable returns()
func (_BPC *BPCTransactorSession) BuyerReg(_dataId *big.Int, _IP [][1]byte) (*types.Transaction, error) {
	return _BPC.Contract.BuyerReg(&_BPC.TransactOpts, _dataId, _IP)
}

// Buyerconfirmed is a paid mutator transaction binding the contract method 0xe88a72ad.
//
// Solidity: function Buyerconfirmed(uint256 BMID, uint256 SPIndex) returns()
func (_BPC *BPCTransactor) Buyerconfirmed(opts *bind.TransactOpts, BMID *big.Int, SPIndex *big.Int) (*types.Transaction, error) {
	return _BPC.contract.Transact(opts, "Buyerconfirmed", BMID, SPIndex)
}

// Buyerconfirmed is a paid mutator transaction binding the contract method 0xe88a72ad.
//
// Solidity: function Buyerconfirmed(uint256 BMID, uint256 SPIndex) returns()
func (_BPC *BPCSession) Buyerconfirmed(BMID *big.Int, SPIndex *big.Int) (*types.Transaction, error) {
	return _BPC.Contract.Buyerconfirmed(&_BPC.TransactOpts, BMID, SPIndex)
}

// Buyerconfirmed is a paid mutator transaction binding the contract method 0xe88a72ad.
//
// Solidity: function Buyerconfirmed(uint256 BMID, uint256 SPIndex) returns()
func (_BPC *BPCTransactorSession) Buyerconfirmed(BMID *big.Int, SPIndex *big.Int) (*types.Transaction, error) {
	return _BPC.Contract.Buyerconfirmed(&_BPC.TransactOpts, BMID, SPIndex)
}

// ControlContract is a paid mutator transaction binding the contract method 0x8cbeb340.
//
// Solidity: function ControlContract(address _SSMCAddr) returns()
func (_BPC *BPCTransactor) ControlContract(opts *bind.TransactOpts, _SSMCAddr common.Address) (*types.Transaction, error) {
	return _BPC.contract.Transact(opts, "ControlContract", _SSMCAddr)
}

// ControlContract is a paid mutator transaction binding the contract method 0x8cbeb340.
//
// Solidity: function ControlContract(address _SSMCAddr) returns()
func (_BPC *BPCSession) ControlContract(_SSMCAddr common.Address) (*types.Transaction, error) {
	return _BPC.Contract.ControlContract(&_BPC.TransactOpts, _SSMCAddr)
}

// ControlContract is a paid mutator transaction binding the contract method 0x8cbeb340.
//
// Solidity: function ControlContract(address _SSMCAddr) returns()
func (_BPC *BPCTransactorSession) ControlContract(_SSMCAddr common.Address) (*types.Transaction, error) {
	return _BPC.Contract.ControlContract(&_BPC.TransactOpts, _SSMCAddr)
}

// ControlContract2 is a paid mutator transaction binding the contract method 0x388b08c6.
//
// Solidity: function ControlContract2(address _SSMCAddr) returns()
func (_BPC *BPCTransactor) ControlContract2(opts *bind.TransactOpts, _SSMCAddr common.Address) (*types.Transaction, error) {
	return _BPC.contract.Transact(opts, "ControlContract2", _SSMCAddr)
}

// ControlContract2 is a paid mutator transaction binding the contract method 0x388b08c6.
//
// Solidity: function ControlContract2(address _SSMCAddr) returns()
func (_BPC *BPCSession) ControlContract2(_SSMCAddr common.Address) (*types.Transaction, error) {
	return _BPC.Contract.ControlContract2(&_BPC.TransactOpts, _SSMCAddr)
}

// ControlContract2 is a paid mutator transaction binding the contract method 0x388b08c6.
//
// Solidity: function ControlContract2(address _SSMCAddr) returns()
func (_BPC *BPCTransactorSession) ControlContract2(_SSMCAddr common.Address) (*types.Transaction, error) {
	return _BPC.Contract.ControlContract2(&_BPC.TransactOpts, _SSMCAddr)
}

// SPReg is a paid mutator transaction binding the contract method 0xdbec6377.
//
// Solidity: function SPReg(uint256 _buyerId, uint256[] _SPID) returns(uint256)
func (_BPC *BPCTransactor) SPReg(opts *bind.TransactOpts, _buyerId *big.Int, _SPID []*big.Int) (*types.Transaction, error) {
	return _BPC.contract.Transact(opts, "SPReg", _buyerId, _SPID)
}

// SPReg is a paid mutator transaction binding the contract method 0xdbec6377.
//
// Solidity: function SPReg(uint256 _buyerId, uint256[] _SPID) returns(uint256)
func (_BPC *BPCSession) SPReg(_buyerId *big.Int, _SPID []*big.Int) (*types.Transaction, error) {
	return _BPC.Contract.SPReg(&_BPC.TransactOpts, _buyerId, _SPID)
}

// SPReg is a paid mutator transaction binding the contract method 0xdbec6377.
//
// Solidity: function SPReg(uint256 _buyerId, uint256[] _SPID) returns(uint256)
func (_BPC *BPCTransactorSession) SPReg(_buyerId *big.Int, _SPID []*big.Int) (*types.Transaction, error) {
	return _BPC.Contract.SPReg(&_BPC.TransactOpts, _buyerId, _SPID)
}

// SPwithdraw is a paid mutator transaction binding the contract method 0x2575cf63.
//
// Solidity: function SPwithdraw(uint256 _BMID, uint256 _SPIndex, uint256 _BMPayID) payable returns()
func (_BPC *BPCTransactor) SPwithdraw(opts *bind.TransactOpts, _BMID *big.Int, _SPIndex *big.Int, _BMPayID *big.Int) (*types.Transaction, error) {
	return _BPC.contract.Transact(opts, "SPwithdraw", _BMID, _SPIndex, _BMPayID)
}

// SPwithdraw is a paid mutator transaction binding the contract method 0x2575cf63.
//
// Solidity: function SPwithdraw(uint256 _BMID, uint256 _SPIndex, uint256 _BMPayID) payable returns()
func (_BPC *BPCSession) SPwithdraw(_BMID *big.Int, _SPIndex *big.Int, _BMPayID *big.Int) (*types.Transaction, error) {
	return _BPC.Contract.SPwithdraw(&_BPC.TransactOpts, _BMID, _SPIndex, _BMPayID)
}

// SPwithdraw is a paid mutator transaction binding the contract method 0x2575cf63.
//
// Solidity: function SPwithdraw(uint256 _BMID, uint256 _SPIndex, uint256 _BMPayID) payable returns()
func (_BPC *BPCTransactorSession) SPwithdraw(_BMID *big.Int, _SPIndex *big.Int, _BMPayID *big.Int) (*types.Transaction, error) {
	return _BPC.Contract.SPwithdraw(&_BPC.TransactOpts, _BMID, _SPIndex, _BMPayID)
}

// SetHashOfDEData is a paid mutator transaction binding the contract method 0x00b17f72.
//
// Solidity: function SetHashOfDEData(uint256 BMID, uint256 SPIndex, bytes32 _HODED) returns()
func (_BPC *BPCTransactor) SetHashOfDEData(opts *bind.TransactOpts, BMID *big.Int, SPIndex *big.Int, _HODED [32]byte) (*types.Transaction, error) {
	return _BPC.contract.Transact(opts, "SetHashOfDEData", BMID, SPIndex, _HODED)
}

// SetHashOfDEData is a paid mutator transaction binding the contract method 0x00b17f72.
//
// Solidity: function SetHashOfDEData(uint256 BMID, uint256 SPIndex, bytes32 _HODED) returns()
func (_BPC *BPCSession) SetHashOfDEData(BMID *big.Int, SPIndex *big.Int, _HODED [32]byte) (*types.Transaction, error) {
	return _BPC.Contract.SetHashOfDEData(&_BPC.TransactOpts, BMID, SPIndex, _HODED)
}

// SetHashOfDEData is a paid mutator transaction binding the contract method 0x00b17f72.
//
// Solidity: function SetHashOfDEData(uint256 BMID, uint256 SPIndex, bytes32 _HODED) returns()
func (_BPC *BPCTransactorSession) SetHashOfDEData(BMID *big.Int, SPIndex *big.Int, _HODED [32]byte) (*types.Transaction, error) {
	return _BPC.Contract.SetHashOfDEData(&_BPC.TransactOpts, BMID, SPIndex, _HODED)
}

// Activate is a paid mutator transaction binding the contract method 0x7901de43.
//
// Solidity: function activate(uint256 BMID, uint256 SPIndex, bytes32 SPkey) returns()
func (_BPC *BPCTransactor) Activate(opts *bind.TransactOpts, BMID *big.Int, SPIndex *big.Int, SPkey [32]byte) (*types.Transaction, error) {
	return _BPC.contract.Transact(opts, "activate", BMID, SPIndex, SPkey)
}

// Activate is a paid mutator transaction binding the contract method 0x7901de43.
//
// Solidity: function activate(uint256 BMID, uint256 SPIndex, bytes32 SPkey) returns()
func (_BPC *BPCSession) Activate(BMID *big.Int, SPIndex *big.Int, SPkey [32]byte) (*types.Transaction, error) {
	return _BPC.Contract.Activate(&_BPC.TransactOpts, BMID, SPIndex, SPkey)
}

// Activate is a paid mutator transaction binding the contract method 0x7901de43.
//
// Solidity: function activate(uint256 BMID, uint256 SPIndex, bytes32 SPkey) returns()
func (_BPC *BPCTransactorSession) Activate(BMID *big.Int, SPIndex *big.Int, SPkey [32]byte) (*types.Transaction, error) {
	return _BPC.Contract.Activate(&_BPC.TransactOpts, BMID, SPIndex, SPkey)
}

// Appeal is a paid mutator transaction binding the contract method 0x9d7f722c.
//
// Solidity: function appeal(uint256 _BMPayID, string DEdata) returns()
func (_BPC *BPCTransactor) Appeal(opts *bind.TransactOpts, _BMPayID *big.Int, DEdata string) (*types.Transaction, error) {
	return _BPC.contract.Transact(opts, "appeal", _BMPayID, DEdata)
}

// Appeal is a paid mutator transaction binding the contract method 0x9d7f722c.
//
// Solidity: function appeal(uint256 _BMPayID, string DEdata) returns()
func (_BPC *BPCSession) Appeal(_BMPayID *big.Int, DEdata string) (*types.Transaction, error) {
	return _BPC.Contract.Appeal(&_BPC.TransactOpts, _BMPayID, DEdata)
}

// Appeal is a paid mutator transaction binding the contract method 0x9d7f722c.
//
// Solidity: function appeal(uint256 _BMPayID, string DEdata) returns()
func (_BPC *BPCTransactorSession) Appeal(_BMPayID *big.Int, DEdata string) (*types.Transaction, error) {
	return _BPC.Contract.Appeal(&_BPC.TransactOpts, _BMPayID, DEdata)
}

// BPCOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BPC contract.
type BPCOwnershipTransferredIterator struct {
	Event *BPCOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BPCOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BPCOwnershipTransferred)
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
		it.Event = new(BPCOwnershipTransferred)
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
func (it *BPCOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BPCOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BPCOwnershipTransferred represents a OwnershipTransferred event raised by the BPC contract.
type BPCOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BPC *BPCFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BPCOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BPC.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BPCOwnershipTransferredIterator{contract: _BPC.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BPC *BPCFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BPCOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BPC.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BPCOwnershipTransferred)
				if err := _BPC.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BPC *BPCFilterer) ParseOwnershipTransferred(log types.Log) (*BPCOwnershipTransferred, error) {
	event := new(BPCOwnershipTransferred)
	if err := _BPC.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnableMetaData contains all meta data concerning the Ownable contract.
var OwnableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8da5cb5b": "owner()",
	},
	Bin: "0x6080604052348015600f57600080fd5b50600080546001600160a01b031916331790556091806100306000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80638da5cb5b14602d575b600080fd5b600054603f906001600160a01b031681565b6040516001600160a01b03909116815260200160405180910390f3fea2646970667358221220ee23b2f5024c799d2d6f2588057ae28d9a0e7ad60200b9b9a88e64b9365b097f64736f6c634300080c0033",
}

// OwnableABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableMetaData.ABI instead.
var OwnableABI = OwnableMetaData.ABI

// Deprecated: Use OwnableMetaData.Sigs instead.
// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = OwnableMetaData.Sigs

// OwnableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OwnableMetaData.Bin instead.
var OwnableBin = OwnableMetaData.Bin

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := OwnableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SBMCMetaData contains all meta data concerning the SBMC contract.
var SBMCMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dataId\",\"type\":\"uint256\"},{\"internalType\":\"bytes1[]\",\"name\":\"_IP\",\"type\":\"bytes1[]\"}],\"name\":\"BuyerReg\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"BMID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"SPIndex\",\"type\":\"uint256\"}],\"name\":\"Buyerconfirmed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_SSMCAddr\",\"type\":\"address\"}],\"name\":\"ControlContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_buyerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_SPID\",\"type\":\"uint256[]\"}],\"name\":\"SPReg\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"BMID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"SPIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_HODED\",\"type\":\"bytes32\"}],\"name\":\"SetHashOfDEData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"94221adf": "BuyerReg(uint256,bytes1[])",
		"e88a72ad": "Buyerconfirmed(uint256,uint256)",
		"8cbeb340": "ControlContract(address)",
		"dbec6377": "SPReg(uint256,uint256[])",
		"00b17f72": "SetHashOfDEData(uint256,uint256,bytes32)",
		"8da5cb5b": "owner()",
	},
	Bin: "0x6080604052601460025534801561001557600080fd5b50600080546001600160a01b03191633179055610e49806100376000396000f3fe6080604052600436106100545760003560e01c8062b17f72146100595780638cbeb3401461007b5780638da5cb5b1461009b57806394221adf146100d8578063dbec6377146100eb578063e88a72ad14610119575b600080fd5b34801561006557600080fd5b50610079610074366004610865565b610139565b005b34801561008757600080fd5b506100796100963660046108a9565b610249565b3480156100a757600080fd5b506000546100bb906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b6100796100e636600461099a565b610282565b3480156100f757600080fd5b5061010b610106366004610a45565b610465565b6040519081526020016100cf565b34801561012557600080fd5b50610079610134366004610ad3565b610615565b6001546004805433926001600160a01b03169163da182640918790811061016257610162610af5565b9060005260206000209060040201600101858154811061018457610184610af5565b90600052602060002001546040518263ffffffff1660e01b81526004016101ad91815260200190565b600060405180830381865afa1580156101ca573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526101f29190810190610b2b565b604001516001600160a01b03161461020957600080fd5b806004848154811061021d5761021d610af5565b9060005260206000209060040201600201838154811061023f5761023f610af5565b5060005250505050565b6000546001600160a01b0316331461026057600080fd5b600180546001600160a01b0319166001600160a01b0392909216919091179055565b6002546001546040516315c96b3960e31b8152600481018590523492916001600160a01b03169063ae4b59c8906024016000604051808303816000875af11580156102d1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526102f99190810190610cad565b604001516103079190610daa565b6001546040516315c96b3960e31b8152600481018690526001600160a01b039091169063ae4b59c8906024016000604051808303816000875af1158015610352573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261037a9190810190610cad565b60c001516103889190610dc9565b111561039357600080fd5b604080516080810182528381526020808201848152339383019390935234606083015260038054600181018255600091909152825160049091027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b8101918255935180519394919361042e937fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85c909301929190910190610709565b5060408201516002820180546001600160a01b0319166001600160a01b039092169190911790556060909101516003909101555050565b6000336001600160a01b03166003848154811061048457610484610af5565b60009182526020909120600260049092020101546001600160a01b0316146104ab57600080fd5b6104d66040518060800160405280600081526020016060815260200160608152602001606081525090565b8381526020808201848152600480546001810182556000829052845191027f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b81019182559151805185949293610552937f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19c9091019201906107af565b506040820151805161056e9160028401916020909101906107af565b506060820151805161058a9160038401916020909101906107ea565b50505060005b8160200151518110156105fd576000801b826040015182815181106105b7576105b7610af5565b6020026020010181815250506000826060015182815181106105db576105db610af5565b91151560209283029190910190910152806105f581610de1565b915050610590565b60045461060c90600190610dfc565b95945050505050565b336001600160a01b031660036004848154811061063457610634610af5565b9060005260206000209060040201600001548154811061065657610656610af5565b60009182526020909120600260049092020101546001600160a01b03161461067d57600080fd5b6000801b6004838154811061069457610694610af5565b906000526020600020906004020160020182815481106106b6576106b6610af5565b906000526020600020015414156106cc57600080fd5b600482815481106106df576106df610af5565b9060005260206000209060040201600301818154811061070157610701610af5565b506000525050565b82805482825590600052602060002090601f0160209004810192821561079f5791602002820160005b8382111561077057835183826101000a81548160ff021916908360f81c02179055509260200192600101602081600001049283019260010302610732565b801561079d5782816101000a81549060ff0219169055600101602081600001049283019260010302610770565b505b506107ab929150610850565b5090565b82805482825590600052602060002090810192821561079f579160200282015b8281111561079f5782518255916020019190600101906107cf565b82805482825590600052602060002090601f0160209004810192821561079f5791602002820160005b8382111561077057835183826101000a81548160ff0219169083151502179055509260200192600101602081600001049283019260010302610813565b5b808211156107ab5760008155600101610851565b60008060006060848603121561087a57600080fd5b505081359360208301359350604090920135919050565b6001600160a01b03811681146108a657600080fd5b50565b6000602082840312156108bb57600080fd5b81356108c681610891565b9392505050565b634e487b7160e01b600052604160045260246000fd5b6040516080810167ffffffffffffffff81118282101715610906576109066108cd565b60405290565b60405160e0810167ffffffffffffffff81118282101715610906576109066108cd565b604051601f8201601f1916810167ffffffffffffffff81118282101715610958576109586108cd565b604052919050565b600067ffffffffffffffff82111561097a5761097a6108cd565b5060051b60200190565b6001600160f81b0319811681146108a657600080fd5b600080604083850312156109ad57600080fd5b8235915060208084013567ffffffffffffffff8111156109cc57600080fd5b8401601f810186136109dd57600080fd5b80356109f06109eb82610960565b61092f565b81815260059190911b82018301908381019088831115610a0f57600080fd5b928401925b82841015610a36578335610a2781610984565b82529284019290840190610a14565b80955050505050509250929050565b60008060408385031215610a5857600080fd5b8235915060208084013567ffffffffffffffff811115610a7757600080fd5b8401601f81018613610a8857600080fd5b8035610a966109eb82610960565b81815260059190911b82018301908381019088831115610ab557600080fd5b928401925b82841015610a3657833582529284019290840190610aba565b60008060408385031215610ae657600080fd5b50508035926020909101359150565b634e487b7160e01b600052603260045260246000fd5b8051610b1681610891565b919050565b80518015158114610b1657600080fd5b60006020808385031215610b3e57600080fd5b825167ffffffffffffffff80821115610b5657600080fd5b9084019060808287031215610b6a57600080fd5b610b726108e3565b825181528383015182811115610b8757600080fd5b83019150601f82018713610b9a57600080fd5b8151610ba86109eb82610960565b81815260059190911b83018501908581019089831115610bc757600080fd5b938601935b82851015610bee578451610bdf81610984565b82529386019390860190610bcc565b8387015250610c01905060408401610b0b565b6040820152610c1260608401610b1b565b60608201529695505050505050565b600082601f830112610c3257600080fd5b815167ffffffffffffffff811115610c4c57610c4c6108cd565b6020610c60601f8301601f1916820161092f565b8281528582848701011115610c7457600080fd5b60005b83811015610c92578581018301518282018401528201610c77565b83811115610ca35760008385840101525b5095945050505050565b600060208284031215610cbf57600080fd5b815167ffffffffffffffff80821115610cd757600080fd5b9083019060e08286031215610ceb57600080fd5b610cf361090c565b610cfc83610b0b565b8152602083015182811115610d1057600080fd5b610d1c87828601610c21565b60208301525060408301516040820152606083015182811115610d3e57600080fd5b610d4a87828601610c21565b606083015250608083015182811115610d6257600080fd5b610d6e87828601610c21565b60808301525060a083015160a082015260c083015160c082015280935050505092915050565b634e487b7160e01b600052601160045260246000fd5b6000816000190483118215151615610dc457610dc4610d94565b500290565b60008219821115610ddc57610ddc610d94565b500190565b6000600019821415610df557610df5610d94565b5060010190565b600082821015610e0e57610e0e610d94565b50039056fea264697066735822122024d067f3a5b14848f77c30f5e535c6334295a009e49e197891c1f72be58ef3f164736f6c634300080c0033",
}

// SBMCABI is the input ABI used to generate the binding from.
// Deprecated: Use SBMCMetaData.ABI instead.
var SBMCABI = SBMCMetaData.ABI

// Deprecated: Use SBMCMetaData.Sigs instead.
// SBMCFuncSigs maps the 4-byte function signature to its string representation.
var SBMCFuncSigs = SBMCMetaData.Sigs

// SBMCBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SBMCMetaData.Bin instead.
var SBMCBin = SBMCMetaData.Bin

// DeploySBMC deploys a new Ethereum contract, binding an instance of SBMC to it.
func DeploySBMC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SBMC, error) {
	parsed, err := SBMCMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SBMCBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SBMC{SBMCCaller: SBMCCaller{contract: contract}, SBMCTransactor: SBMCTransactor{contract: contract}, SBMCFilterer: SBMCFilterer{contract: contract}}, nil
}

// SBMC is an auto generated Go binding around an Ethereum contract.
type SBMC struct {
	SBMCCaller     // Read-only binding to the contract
	SBMCTransactor // Write-only binding to the contract
	SBMCFilterer   // Log filterer for contract events
}

// SBMCCaller is an auto generated read-only Go binding around an Ethereum contract.
type SBMCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SBMCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SBMCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SBMCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SBMCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SBMCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SBMCSession struct {
	Contract     *SBMC             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SBMCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SBMCCallerSession struct {
	Contract *SBMCCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SBMCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SBMCTransactorSession struct {
	Contract     *SBMCTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SBMCRaw is an auto generated low-level Go binding around an Ethereum contract.
type SBMCRaw struct {
	Contract *SBMC // Generic contract binding to access the raw methods on
}

// SBMCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SBMCCallerRaw struct {
	Contract *SBMCCaller // Generic read-only contract binding to access the raw methods on
}

// SBMCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SBMCTransactorRaw struct {
	Contract *SBMCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSBMC creates a new instance of SBMC, bound to a specific deployed contract.
func NewSBMC(address common.Address, backend bind.ContractBackend) (*SBMC, error) {
	contract, err := bindSBMC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SBMC{SBMCCaller: SBMCCaller{contract: contract}, SBMCTransactor: SBMCTransactor{contract: contract}, SBMCFilterer: SBMCFilterer{contract: contract}}, nil
}

// NewSBMCCaller creates a new read-only instance of SBMC, bound to a specific deployed contract.
func NewSBMCCaller(address common.Address, caller bind.ContractCaller) (*SBMCCaller, error) {
	contract, err := bindSBMC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SBMCCaller{contract: contract}, nil
}

// NewSBMCTransactor creates a new write-only instance of SBMC, bound to a specific deployed contract.
func NewSBMCTransactor(address common.Address, transactor bind.ContractTransactor) (*SBMCTransactor, error) {
	contract, err := bindSBMC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SBMCTransactor{contract: contract}, nil
}

// NewSBMCFilterer creates a new log filterer instance of SBMC, bound to a specific deployed contract.
func NewSBMCFilterer(address common.Address, filterer bind.ContractFilterer) (*SBMCFilterer, error) {
	contract, err := bindSBMC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SBMCFilterer{contract: contract}, nil
}

// bindSBMC binds a generic wrapper to an already deployed contract.
func bindSBMC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SBMCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SBMC *SBMCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SBMC.Contract.SBMCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SBMC *SBMCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SBMC.Contract.SBMCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SBMC *SBMCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SBMC.Contract.SBMCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SBMC *SBMCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SBMC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SBMC *SBMCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SBMC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SBMC *SBMCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SBMC.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SBMC *SBMCCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SBMC.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SBMC *SBMCSession) Owner() (common.Address, error) {
	return _SBMC.Contract.Owner(&_SBMC.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SBMC *SBMCCallerSession) Owner() (common.Address, error) {
	return _SBMC.Contract.Owner(&_SBMC.CallOpts)
}

// BuyerReg is a paid mutator transaction binding the contract method 0x94221adf.
//
// Solidity: function BuyerReg(uint256 _dataId, bytes1[] _IP) payable returns()
func (_SBMC *SBMCTransactor) BuyerReg(opts *bind.TransactOpts, _dataId *big.Int, _IP [][1]byte) (*types.Transaction, error) {
	return _SBMC.contract.Transact(opts, "BuyerReg", _dataId, _IP)
}

// BuyerReg is a paid mutator transaction binding the contract method 0x94221adf.
//
// Solidity: function BuyerReg(uint256 _dataId, bytes1[] _IP) payable returns()
func (_SBMC *SBMCSession) BuyerReg(_dataId *big.Int, _IP [][1]byte) (*types.Transaction, error) {
	return _SBMC.Contract.BuyerReg(&_SBMC.TransactOpts, _dataId, _IP)
}

// BuyerReg is a paid mutator transaction binding the contract method 0x94221adf.
//
// Solidity: function BuyerReg(uint256 _dataId, bytes1[] _IP) payable returns()
func (_SBMC *SBMCTransactorSession) BuyerReg(_dataId *big.Int, _IP [][1]byte) (*types.Transaction, error) {
	return _SBMC.Contract.BuyerReg(&_SBMC.TransactOpts, _dataId, _IP)
}

// Buyerconfirmed is a paid mutator transaction binding the contract method 0xe88a72ad.
//
// Solidity: function Buyerconfirmed(uint256 BMID, uint256 SPIndex) returns()
func (_SBMC *SBMCTransactor) Buyerconfirmed(opts *bind.TransactOpts, BMID *big.Int, SPIndex *big.Int) (*types.Transaction, error) {
	return _SBMC.contract.Transact(opts, "Buyerconfirmed", BMID, SPIndex)
}

// Buyerconfirmed is a paid mutator transaction binding the contract method 0xe88a72ad.
//
// Solidity: function Buyerconfirmed(uint256 BMID, uint256 SPIndex) returns()
func (_SBMC *SBMCSession) Buyerconfirmed(BMID *big.Int, SPIndex *big.Int) (*types.Transaction, error) {
	return _SBMC.Contract.Buyerconfirmed(&_SBMC.TransactOpts, BMID, SPIndex)
}

// Buyerconfirmed is a paid mutator transaction binding the contract method 0xe88a72ad.
//
// Solidity: function Buyerconfirmed(uint256 BMID, uint256 SPIndex) returns()
func (_SBMC *SBMCTransactorSession) Buyerconfirmed(BMID *big.Int, SPIndex *big.Int) (*types.Transaction, error) {
	return _SBMC.Contract.Buyerconfirmed(&_SBMC.TransactOpts, BMID, SPIndex)
}

// ControlContract is a paid mutator transaction binding the contract method 0x8cbeb340.
//
// Solidity: function ControlContract(address _SSMCAddr) returns()
func (_SBMC *SBMCTransactor) ControlContract(opts *bind.TransactOpts, _SSMCAddr common.Address) (*types.Transaction, error) {
	return _SBMC.contract.Transact(opts, "ControlContract", _SSMCAddr)
}

// ControlContract is a paid mutator transaction binding the contract method 0x8cbeb340.
//
// Solidity: function ControlContract(address _SSMCAddr) returns()
func (_SBMC *SBMCSession) ControlContract(_SSMCAddr common.Address) (*types.Transaction, error) {
	return _SBMC.Contract.ControlContract(&_SBMC.TransactOpts, _SSMCAddr)
}

// ControlContract is a paid mutator transaction binding the contract method 0x8cbeb340.
//
// Solidity: function ControlContract(address _SSMCAddr) returns()
func (_SBMC *SBMCTransactorSession) ControlContract(_SSMCAddr common.Address) (*types.Transaction, error) {
	return _SBMC.Contract.ControlContract(&_SBMC.TransactOpts, _SSMCAddr)
}

// SPReg is a paid mutator transaction binding the contract method 0xdbec6377.
//
// Solidity: function SPReg(uint256 _buyerId, uint256[] _SPID) returns(uint256)
func (_SBMC *SBMCTransactor) SPReg(opts *bind.TransactOpts, _buyerId *big.Int, _SPID []*big.Int) (*types.Transaction, error) {
	return _SBMC.contract.Transact(opts, "SPReg", _buyerId, _SPID)
}

// SPReg is a paid mutator transaction binding the contract method 0xdbec6377.
//
// Solidity: function SPReg(uint256 _buyerId, uint256[] _SPID) returns(uint256)
func (_SBMC *SBMCSession) SPReg(_buyerId *big.Int, _SPID []*big.Int) (*types.Transaction, error) {
	return _SBMC.Contract.SPReg(&_SBMC.TransactOpts, _buyerId, _SPID)
}

// SPReg is a paid mutator transaction binding the contract method 0xdbec6377.
//
// Solidity: function SPReg(uint256 _buyerId, uint256[] _SPID) returns(uint256)
func (_SBMC *SBMCTransactorSession) SPReg(_buyerId *big.Int, _SPID []*big.Int) (*types.Transaction, error) {
	return _SBMC.Contract.SPReg(&_SBMC.TransactOpts, _buyerId, _SPID)
}

// SetHashOfDEData is a paid mutator transaction binding the contract method 0x00b17f72.
//
// Solidity: function SetHashOfDEData(uint256 BMID, uint256 SPIndex, bytes32 _HODED) returns()
func (_SBMC *SBMCTransactor) SetHashOfDEData(opts *bind.TransactOpts, BMID *big.Int, SPIndex *big.Int, _HODED [32]byte) (*types.Transaction, error) {
	return _SBMC.contract.Transact(opts, "SetHashOfDEData", BMID, SPIndex, _HODED)
}

// SetHashOfDEData is a paid mutator transaction binding the contract method 0x00b17f72.
//
// Solidity: function SetHashOfDEData(uint256 BMID, uint256 SPIndex, bytes32 _HODED) returns()
func (_SBMC *SBMCSession) SetHashOfDEData(BMID *big.Int, SPIndex *big.Int, _HODED [32]byte) (*types.Transaction, error) {
	return _SBMC.Contract.SetHashOfDEData(&_SBMC.TransactOpts, BMID, SPIndex, _HODED)
}

// SetHashOfDEData is a paid mutator transaction binding the contract method 0x00b17f72.
//
// Solidity: function SetHashOfDEData(uint256 BMID, uint256 SPIndex, bytes32 _HODED) returns()
func (_SBMC *SBMCTransactorSession) SetHashOfDEData(BMID *big.Int, SPIndex *big.Int, _HODED [32]byte) (*types.Transaction, error) {
	return _SBMC.Contract.SetHashOfDEData(&_SBMC.TransactOpts, BMID, SPIndex, _HODED)
}

// SBMCOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SBMC contract.
type SBMCOwnershipTransferredIterator struct {
	Event *SBMCOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SBMCOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SBMCOwnershipTransferred)
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
		it.Event = new(SBMCOwnershipTransferred)
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
func (it *SBMCOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SBMCOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SBMCOwnershipTransferred represents a OwnershipTransferred event raised by the SBMC contract.
type SBMCOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SBMC *SBMCFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SBMCOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SBMC.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SBMCOwnershipTransferredIterator{contract: _SBMC.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SBMC *SBMCFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SBMCOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SBMC.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SBMCOwnershipTransferred)
				if err := _SBMC.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SBMC *SBMCFilterer) ParseOwnershipTransferred(log types.Log) (*SBMCOwnershipTransferred, error) {
	event := new(SBMCOwnershipTransferred)
	if err := _SBMC.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SSMCMetaData contains all meta data concerning the SSMC contract.
var SSMCMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_SPId\",\"type\":\"uint256\"}],\"name\":\"Confirm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_DataId\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"Plaintexts\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"EncryptParts\",\"type\":\"string[]\"}],\"name\":\"ExposeParts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_DataId\",\"type\":\"uint256\"},{\"internalType\":\"bytes1[]\",\"name\":\"IP\",\"type\":\"bytes1[]\"}],\"name\":\"SPReg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_des\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_PartsNum\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_HOP\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_HOEP\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"SellerDataset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"}],\"name\":\"_rand\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"viewSPbyIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dataId\",\"type\":\"uint256\"},{\"internalType\":\"bytes1[]\",\"name\":\"IP\",\"type\":\"bytes1[]\"},{\"internalType\":\"address\",\"name\":\"ASP\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"confirmed\",\"type\":\"bool\"}],\"internalType\":\"structSSMC.SPInfor\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dataId\",\"type\":\"uint256\"}],\"name\":\"viewSPbydataId\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dataId\",\"type\":\"uint256\"},{\"internalType\":\"bytes1[]\",\"name\":\"IP\",\"type\":\"bytes1[]\"},{\"internalType\":\"address\",\"name\":\"ASP\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"confirmed\",\"type\":\"bool\"}],\"internalType\":\"structSSMC.SPInfor[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ID\",\"type\":\"uint256\"}],\"name\":\"viewSeller\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"Aseller\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"Description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"PartsNum\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"HashOfParts\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"HashOfEncryptParts\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"RandIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"internalType\":\"structSSMC.SellerData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"0bab98c8": "Confirm(uint256)",
		"60d3a57e": "ExposeParts(uint256,string[],string[])",
		"769dab04": "SPReg(uint256,bytes1[])",
		"959d38dd": "SellerDataset(string,uint256,string,string,uint256)",
		"e1eddc6d": "_rand(uint256)",
		"da182640": "viewSPbyIndex(uint256)",
		"038f4c61": "viewSPbydataId(uint256)",
		"ae4b59c8": "viewSeller(uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b5061153f806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063959d38dd1161005b578063959d38dd146100f1578063ae4b59c814610104578063da18264014610124578063e1eddc6d1461014457600080fd5b8063038f4c611461008d5780630bab98c8146100b657806360d3a57e146100cb578063769dab04146100de575b600080fd5b6100a061009b366004610ea2565b610165565b6040516100ad9190610f3e565b60405180910390f35b6100c96100c4366004610ea2565b6102c7565b005b6100c96100d936600461110b565b610372565b6100c96100ec366004611178565b610685565b6100c96100ff36600461122c565b61075f565b610117610112366004610ea2565b61089c565b6040516100ad919061131e565b610137610132366004610ea2565b610b03565b6040516100ad91906113b4565b610157610152366004610ea2565b610c1c565b6040519081526020016100ad565b6060600080825b6002548310156102bf57846002848154811061018a5761018a6113c7565b90600052602060002090600302016000015414156102ad57600283815481106101b5576101b56113c7565b9060005260206000209060030201604051806080016040529081600082015481526020016001820180548060200260200160405190810160405280929190818152602001828054801561024f57602002820191906000526020600020906000905b825461010083900a900460f81b6001600160f81b0319168152602060019283018181049485019490930390920291018084116102165790505b5050509183525050600291909101546001600160a01b0381166020830152600160a01b900460ff1615156040909101528151829084908110610293576102936113c7565b602002602001018190525081806102a9906113f3565b9250505b826102b7816113f3565b93505061016c565b949350505050565b6000600282815481106102dc576102dc6113c7565b9060005260206000209060030201600001549050336001600160a01b03166000828154811061030d5761030d6113c7565b60009182526020909120600790910201546001600160a01b03161461033157600080fd5b600160028381548110610346576103466113c7565b906000526020600020906003020160020160146101000a81548160ff0219169083151502179055505050565b336001600160a01b03166000848154811061038f5761038f6113c7565b60009182526020909120600790910201546001600160a01b0316146103b357600080fd5b60005b600a81101561067f576104df600085815481106103d5576103d56113c7565b906000526020600020906007020160030180546103f19061140e565b80601f016020809104026020016040519081016040528092919081815260200182805461041d9061140e565b801561046a5780601f1061043f5761010080835404028352916020019161046a565b820191906000526020600020905b81548152906001019060200180831161044d57829003601f168201915b5050505050600a60008781548110610484576104846113c7565b906000526020600020906007020160020154846104a19190611449565b6104ab919061147e565b600087815481106104be576104be6113c7565b9060005260206000209060070201600501546104da9190611492565b610c66565b60028483815181106104f3576104f36113c7565b602002602001015160405161050891906114aa565b602060405180830381855afa158015610525573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061054891906114c6565b14610555576105556114df565b6105876000858154811061056b5761056b6113c7565b906000526020600020906007020160040180546103f19061140e565b600283838151811061059b5761059b6113c7565b60200260200101516040516105b091906114aa565b602060405180830381855afa1580156105cd573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906105f091906114c6565b146105fd576105fd6114df565b60405180604001604052808481526020018381525060018581548110610625576106256113c7565b9060005260206000209060020201600082015181600001908051906020019061064f929190610cc0565b5060208281015180516106689260018501920190610cc0565b509050508080610677906113f3565b9150506103b6565b50505050565b6040805160808101825283815260208082018481523393830193909352600060608301819052600280546001810182559152825160039091027f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace8101918255935180519394919361071f937f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5acf909301929190910190610d1d565b506040820151600290910180546060909301511515600160a01b026001600160a81b03199093166001600160a01b03909216919091179190911790555050565b61271084111561076e57600080fd5b600a84101561077c57600080fd5b815183511461078a57600080fd5b61079684610100611449565b8351146107a257600080fd5b60006040518060e00160405280336001600160a01b031681526020018781526020018681526020018581526020018481526020016107df87610c1c565b81526020908101849052825460018082018555600094855293829020835160079092020180546001600160a01b0319166001600160a01b039092169190911781558282015180519394919361083c93928501929190910190610dbf565b506040820151600282015560608201518051610862916003840191602090910190610dbf565b506080820151805161087e916004840191602090910190610dbf565b5060a0820151816005015560c0820151816006015550505050505050565b6108e56040518060e0016040528060006001600160a01b031681526020016060815260200160008152602001606081526020016060815260200160008152602001600081525090565b600082815481106108f8576108f86113c7565b60009182526020918290206040805160e0810190915260079092020180546001600160a01b0316825260018101805492939192918401916109389061140e565b80601f01602080910402602001604051908101604052809291908181526020018280546109649061140e565b80156109b15780601f10610986576101008083540402835291602001916109b1565b820191906000526020600020905b81548152906001019060200180831161099457829003601f168201915b50505050508152602001600282015481526020016003820180546109d49061140e565b80601f0160208091040260200160405190810160405280929190818152602001828054610a009061140e565b8015610a4d5780601f10610a2257610100808354040283529160200191610a4d565b820191906000526020600020905b815481529060010190602001808311610a3057829003601f168201915b50505050508152602001600482018054610a669061140e565b80601f0160208091040260200160405190810160405280929190818152602001828054610a929061140e565b8015610adf5780601f10610ab457610100808354040283529160200191610adf565b820191906000526020600020905b815481529060010190602001808311610ac257829003601f168201915b50505050508152602001600582015481526020016006820154815250509050919050565b610b396040518060800160405280600081526020016060815260200160006001600160a01b031681526020016000151581525090565b60028281548110610b4c57610b4c6113c7565b90600052602060002090600302016040518060800160405290816000820154815260200160018201805480602002602001604051908101604052809291908181526020018280548015610be657602002820191906000526020600020906000905b825461010083900a900460f81b6001600160f81b031916815260206001928301818104948501949093039092029101808411610bad5790505b5050509183525050600291909101546001600160a01b0381166020830152600160a01b900460ff16151560409091015292915050565b6000804442604051602001610c3b929190918252602082015260400190565b60408051601f1981840301815291905280516020909101209050610c5f83826114f5565b9392505050565b600080805b6020811015610cb85784610c7f8286611492565b81518110610c8f57610c8f6113c7565b01602001516001600160f81b0319169190911760081b9080610cb0816113f3565b915050610c6b565b505092915050565b828054828255906000526020600020908101928215610d0d579160200282015b82811115610d0d5782518051610cfd918491602090910190610dbf565b5091602001919060010190610ce0565b50610d19929150610e33565b5090565b82805482825590600052602060002090601f01602090048101928215610db35791602002820160005b83821115610d8457835183826101000a81548160ff021916908360f81c02179055509260200192600101602081600001049283019260010302610d46565b8015610db15782816101000a81549060ff0219169055600101602081600001049283019260010302610d84565b505b50610d19929150610e50565b828054610dcb9061140e565b90600052602060002090601f016020900481019282610ded5760008555610db3565b82601f10610e0657805160ff1916838001178555610db3565b82800160010185558215610db3579182015b82811115610db3578251825591602001919060010190610e18565b80821115610d19576000610e478282610e65565b50600101610e33565b5b80821115610d195760008155600101610e51565b508054610e719061140e565b6000825580601f10610e81575050565b601f016020900490600052602060002090810190610e9f9190610e50565b50565b600060208284031215610eb457600080fd5b5035919050565b6000608083018251845260208084015160808287015282815180855260a0880191508383019450600092505b80831015610f115784516001600160f81b0319168252938301936001929092019190830190610ee7565b506040868101516001600160a01b0316908801526060958601511515959096019490945250929392505050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b82811015610f9357603f19888603018452610f81858351610ebb565b94509285019290850190600101610f65565b5092979650505050505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715610fdf57610fdf610fa0565b604052919050565b600067ffffffffffffffff82111561100157611001610fa0565b5060051b60200190565b600082601f83011261101c57600080fd5b813567ffffffffffffffff81111561103657611036610fa0565b611049601f8201601f1916602001610fb6565b81815284602083860101111561105e57600080fd5b816020850160208301376000918101602001919091529392505050565b600082601f83011261108c57600080fd5b813560206110a161109c83610fe7565b610fb6565b82815260059290921b840181019181810190868411156110c057600080fd5b8286015b8481101561110057803567ffffffffffffffff8111156110e45760008081fd5b6110f28986838b010161100b565b8452509183019183016110c4565b509695505050505050565b60008060006060848603121561112057600080fd5b83359250602084013567ffffffffffffffff8082111561113f57600080fd5b61114b8783880161107b565b9350604086013591508082111561116157600080fd5b5061116e8682870161107b565b9150509250925092565b6000806040838503121561118b57600080fd5b8235915060208084013567ffffffffffffffff8111156111aa57600080fd5b8401601f810186136111bb57600080fd5b80356111c961109c82610fe7565b81815260059190911b820183019083810190888311156111e857600080fd5b928401925b8284101561121d5783356001600160f81b03198116811461120e5760008081fd5b825292840192908401906111ed565b80955050505050509250929050565b600080600080600060a0868803121561124457600080fd5b853567ffffffffffffffff8082111561125c57600080fd5b61126889838a0161100b565b965060208801359550604088013591508082111561128557600080fd5b61129189838a0161100b565b945060608801359150808211156112a757600080fd5b506112b48882890161100b565b95989497509295608001359392505050565b60005b838110156112e15781810151838201526020016112c9565b8381111561067f5750506000910152565b6000815180845261130a8160208601602086016112c6565b601f01601f19169290920160200192915050565b602080825282516001600160a01b03168282015282015160e0604083015260009061134d6101008401826112f2565b9050604084015160608401526060840151601f198085840301608086015261137583836112f2565b925060808601519150808584030160a08601525061139382826112f2565b91505060a084015160c084015260c084015160e08401528091505092915050565b602081526000610c5f6020830184610ebb565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415611407576114076113dd565b5060010190565b600181811c9082168061142257607f821691505b6020821081141561144357634e487b7160e01b600052602260045260246000fd5b50919050565b6000816000190483118215151615611463576114636113dd565b500290565b634e487b7160e01b600052601260045260246000fd5b60008261148d5761148d611468565b500490565b600082198211156114a5576114a56113dd565b500190565b600082516114bc8184602087016112c6565b9190910192915050565b6000602082840312156114d857600080fd5b5051919050565b634e487b7160e01b600052600160045260246000fd5b60008261150457611504611468565b50069056fea264697066735822122086668fecae9fcb83e3b098606e0562f869145467be57875e8d7d976667485fb064736f6c634300080c0033",
}

// SSMCABI is the input ABI used to generate the binding from.
// Deprecated: Use SSMCMetaData.ABI instead.
var SSMCABI = SSMCMetaData.ABI

// Deprecated: Use SSMCMetaData.Sigs instead.
// SSMCFuncSigs maps the 4-byte function signature to its string representation.
var SSMCFuncSigs = SSMCMetaData.Sigs

// SSMCBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SSMCMetaData.Bin instead.
var SSMCBin = SSMCMetaData.Bin

// DeploySSMC deploys a new Ethereum contract, binding an instance of SSMC to it.
func DeploySSMC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SSMC, error) {
	parsed, err := SSMCMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SSMCBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SSMC{SSMCCaller: SSMCCaller{contract: contract}, SSMCTransactor: SSMCTransactor{contract: contract}, SSMCFilterer: SSMCFilterer{contract: contract}}, nil
}

// SSMC is an auto generated Go binding around an Ethereum contract.
type SSMC struct {
	SSMCCaller     // Read-only binding to the contract
	SSMCTransactor // Write-only binding to the contract
	SSMCFilterer   // Log filterer for contract events
}

// SSMCCaller is an auto generated read-only Go binding around an Ethereum contract.
type SSMCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SSMCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SSMCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SSMCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SSMCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SSMCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SSMCSession struct {
	Contract     *SSMC             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SSMCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SSMCCallerSession struct {
	Contract *SSMCCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SSMCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SSMCTransactorSession struct {
	Contract     *SSMCTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SSMCRaw is an auto generated low-level Go binding around an Ethereum contract.
type SSMCRaw struct {
	Contract *SSMC // Generic contract binding to access the raw methods on
}

// SSMCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SSMCCallerRaw struct {
	Contract *SSMCCaller // Generic read-only contract binding to access the raw methods on
}

// SSMCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SSMCTransactorRaw struct {
	Contract *SSMCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSSMC creates a new instance of SSMC, bound to a specific deployed contract.
func NewSSMC(address common.Address, backend bind.ContractBackend) (*SSMC, error) {
	contract, err := bindSSMC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SSMC{SSMCCaller: SSMCCaller{contract: contract}, SSMCTransactor: SSMCTransactor{contract: contract}, SSMCFilterer: SSMCFilterer{contract: contract}}, nil
}

// NewSSMCCaller creates a new read-only instance of SSMC, bound to a specific deployed contract.
func NewSSMCCaller(address common.Address, caller bind.ContractCaller) (*SSMCCaller, error) {
	contract, err := bindSSMC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SSMCCaller{contract: contract}, nil
}

// NewSSMCTransactor creates a new write-only instance of SSMC, bound to a specific deployed contract.
func NewSSMCTransactor(address common.Address, transactor bind.ContractTransactor) (*SSMCTransactor, error) {
	contract, err := bindSSMC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SSMCTransactor{contract: contract}, nil
}

// NewSSMCFilterer creates a new log filterer instance of SSMC, bound to a specific deployed contract.
func NewSSMCFilterer(address common.Address, filterer bind.ContractFilterer) (*SSMCFilterer, error) {
	contract, err := bindSSMC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SSMCFilterer{contract: contract}, nil
}

// bindSSMC binds a generic wrapper to an already deployed contract.
func bindSSMC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SSMCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SSMC *SSMCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SSMC.Contract.SSMCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SSMC *SSMCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SSMC.Contract.SSMCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SSMC *SSMCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SSMC.Contract.SSMCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SSMC *SSMCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SSMC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SSMC *SSMCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SSMC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SSMC *SSMCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SSMC.Contract.contract.Transact(opts, method, params...)
}

// Rand is a free data retrieval call binding the contract method 0xe1eddc6d.
//
// Solidity: function _rand(uint256 _length) view returns(uint256)
func (_SSMC *SSMCCaller) Rand(opts *bind.CallOpts, _length *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SSMC.contract.Call(opts, &out, "_rand", _length)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rand is a free data retrieval call binding the contract method 0xe1eddc6d.
//
// Solidity: function _rand(uint256 _length) view returns(uint256)
func (_SSMC *SSMCSession) Rand(_length *big.Int) (*big.Int, error) {
	return _SSMC.Contract.Rand(&_SSMC.CallOpts, _length)
}

// Rand is a free data retrieval call binding the contract method 0xe1eddc6d.
//
// Solidity: function _rand(uint256 _length) view returns(uint256)
func (_SSMC *SSMCCallerSession) Rand(_length *big.Int) (*big.Int, error) {
	return _SSMC.Contract.Rand(&_SSMC.CallOpts, _length)
}

// ViewSPbyIndex is a free data retrieval call binding the contract method 0xda182640.
//
// Solidity: function viewSPbyIndex(uint256 _index) view returns((uint256,bytes1[],address,bool))
func (_SSMC *SSMCCaller) ViewSPbyIndex(opts *bind.CallOpts, _index *big.Int) (SSMCSPInfor, error) {
	var out []interface{}
	err := _SSMC.contract.Call(opts, &out, "viewSPbyIndex", _index)

	if err != nil {
		return *new(SSMCSPInfor), err
	}

	out0 := *abi.ConvertType(out[0], new(SSMCSPInfor)).(*SSMCSPInfor)

	return out0, err

}

// ViewSPbyIndex is a free data retrieval call binding the contract method 0xda182640.
//
// Solidity: function viewSPbyIndex(uint256 _index) view returns((uint256,bytes1[],address,bool))
func (_SSMC *SSMCSession) ViewSPbyIndex(_index *big.Int) (SSMCSPInfor, error) {
	return _SSMC.Contract.ViewSPbyIndex(&_SSMC.CallOpts, _index)
}

// ViewSPbyIndex is a free data retrieval call binding the contract method 0xda182640.
//
// Solidity: function viewSPbyIndex(uint256 _index) view returns((uint256,bytes1[],address,bool))
func (_SSMC *SSMCCallerSession) ViewSPbyIndex(_index *big.Int) (SSMCSPInfor, error) {
	return _SSMC.Contract.ViewSPbyIndex(&_SSMC.CallOpts, _index)
}

// ViewSPbydataId is a free data retrieval call binding the contract method 0x038f4c61.
//
// Solidity: function viewSPbydataId(uint256 _dataId) view returns((uint256,bytes1[],address,bool)[])
func (_SSMC *SSMCCaller) ViewSPbydataId(opts *bind.CallOpts, _dataId *big.Int) ([]SSMCSPInfor, error) {
	var out []interface{}
	err := _SSMC.contract.Call(opts, &out, "viewSPbydataId", _dataId)

	if err != nil {
		return *new([]SSMCSPInfor), err
	}

	out0 := *abi.ConvertType(out[0], new([]SSMCSPInfor)).(*[]SSMCSPInfor)

	return out0, err

}

// ViewSPbydataId is a free data retrieval call binding the contract method 0x038f4c61.
//
// Solidity: function viewSPbydataId(uint256 _dataId) view returns((uint256,bytes1[],address,bool)[])
func (_SSMC *SSMCSession) ViewSPbydataId(_dataId *big.Int) ([]SSMCSPInfor, error) {
	return _SSMC.Contract.ViewSPbydataId(&_SSMC.CallOpts, _dataId)
}

// ViewSPbydataId is a free data retrieval call binding the contract method 0x038f4c61.
//
// Solidity: function viewSPbydataId(uint256 _dataId) view returns((uint256,bytes1[],address,bool)[])
func (_SSMC *SSMCCallerSession) ViewSPbydataId(_dataId *big.Int) ([]SSMCSPInfor, error) {
	return _SSMC.Contract.ViewSPbydataId(&_SSMC.CallOpts, _dataId)
}

// Confirm is a paid mutator transaction binding the contract method 0x0bab98c8.
//
// Solidity: function Confirm(uint256 _SPId) returns()
func (_SSMC *SSMCTransactor) Confirm(opts *bind.TransactOpts, _SPId *big.Int) (*types.Transaction, error) {
	return _SSMC.contract.Transact(opts, "Confirm", _SPId)
}

// Confirm is a paid mutator transaction binding the contract method 0x0bab98c8.
//
// Solidity: function Confirm(uint256 _SPId) returns()
func (_SSMC *SSMCSession) Confirm(_SPId *big.Int) (*types.Transaction, error) {
	return _SSMC.Contract.Confirm(&_SSMC.TransactOpts, _SPId)
}

// Confirm is a paid mutator transaction binding the contract method 0x0bab98c8.
//
// Solidity: function Confirm(uint256 _SPId) returns()
func (_SSMC *SSMCTransactorSession) Confirm(_SPId *big.Int) (*types.Transaction, error) {
	return _SSMC.Contract.Confirm(&_SSMC.TransactOpts, _SPId)
}

// ExposeParts is a paid mutator transaction binding the contract method 0x60d3a57e.
//
// Solidity: function ExposeParts(uint256 _DataId, string[] Plaintexts, string[] EncryptParts) returns()
func (_SSMC *SSMCTransactor) ExposeParts(opts *bind.TransactOpts, _DataId *big.Int, Plaintexts []string, EncryptParts []string) (*types.Transaction, error) {
	return _SSMC.contract.Transact(opts, "ExposeParts", _DataId, Plaintexts, EncryptParts)
}

// ExposeParts is a paid mutator transaction binding the contract method 0x60d3a57e.
//
// Solidity: function ExposeParts(uint256 _DataId, string[] Plaintexts, string[] EncryptParts) returns()
func (_SSMC *SSMCSession) ExposeParts(_DataId *big.Int, Plaintexts []string, EncryptParts []string) (*types.Transaction, error) {
	return _SSMC.Contract.ExposeParts(&_SSMC.TransactOpts, _DataId, Plaintexts, EncryptParts)
}

// ExposeParts is a paid mutator transaction binding the contract method 0x60d3a57e.
//
// Solidity: function ExposeParts(uint256 _DataId, string[] Plaintexts, string[] EncryptParts) returns()
func (_SSMC *SSMCTransactorSession) ExposeParts(_DataId *big.Int, Plaintexts []string, EncryptParts []string) (*types.Transaction, error) {
	return _SSMC.Contract.ExposeParts(&_SSMC.TransactOpts, _DataId, Plaintexts, EncryptParts)
}

// SPReg is a paid mutator transaction binding the contract method 0x769dab04.
//
// Solidity: function SPReg(uint256 _DataId, bytes1[] IP) returns()
func (_SSMC *SSMCTransactor) SPReg(opts *bind.TransactOpts, _DataId *big.Int, IP [][1]byte) (*types.Transaction, error) {
	return _SSMC.contract.Transact(opts, "SPReg", _DataId, IP)
}

// SPReg is a paid mutator transaction binding the contract method 0x769dab04.
//
// Solidity: function SPReg(uint256 _DataId, bytes1[] IP) returns()
func (_SSMC *SSMCSession) SPReg(_DataId *big.Int, IP [][1]byte) (*types.Transaction, error) {
	return _SSMC.Contract.SPReg(&_SSMC.TransactOpts, _DataId, IP)
}

// SPReg is a paid mutator transaction binding the contract method 0x769dab04.
//
// Solidity: function SPReg(uint256 _DataId, bytes1[] IP) returns()
func (_SSMC *SSMCTransactorSession) SPReg(_DataId *big.Int, IP [][1]byte) (*types.Transaction, error) {
	return _SSMC.Contract.SPReg(&_SSMC.TransactOpts, _DataId, IP)
}

// SellerDataset is a paid mutator transaction binding the contract method 0x959d38dd.
//
// Solidity: function SellerDataset(string _des, uint256 _PartsNum, string _HOP, string _HOEP, uint256 _price) returns()
func (_SSMC *SSMCTransactor) SellerDataset(opts *bind.TransactOpts, _des string, _PartsNum *big.Int, _HOP string, _HOEP string, _price *big.Int) (*types.Transaction, error) {
	return _SSMC.contract.Transact(opts, "SellerDataset", _des, _PartsNum, _HOP, _HOEP, _price)
}

// SellerDataset is a paid mutator transaction binding the contract method 0x959d38dd.
//
// Solidity: function SellerDataset(string _des, uint256 _PartsNum, string _HOP, string _HOEP, uint256 _price) returns()
func (_SSMC *SSMCSession) SellerDataset(_des string, _PartsNum *big.Int, _HOP string, _HOEP string, _price *big.Int) (*types.Transaction, error) {
	return _SSMC.Contract.SellerDataset(&_SSMC.TransactOpts, _des, _PartsNum, _HOP, _HOEP, _price)
}

// SellerDataset is a paid mutator transaction binding the contract method 0x959d38dd.
//
// Solidity: function SellerDataset(string _des, uint256 _PartsNum, string _HOP, string _HOEP, uint256 _price) returns()
func (_SSMC *SSMCTransactorSession) SellerDataset(_des string, _PartsNum *big.Int, _HOP string, _HOEP string, _price *big.Int) (*types.Transaction, error) {
	return _SSMC.Contract.SellerDataset(&_SSMC.TransactOpts, _des, _PartsNum, _HOP, _HOEP, _price)
}

// ViewSeller is a paid mutator transaction binding the contract method 0xae4b59c8.
//
// Solidity: function viewSeller(uint256 _ID) returns((address,string,uint256,string,string,uint256,uint256))
func (_SSMC *SSMCTransactor) ViewSeller(opts *bind.TransactOpts, _ID *big.Int) (*types.Transaction, error) {
	return _SSMC.contract.Transact(opts, "viewSeller", _ID)
}

// ViewSeller is a paid mutator transaction binding the contract method 0xae4b59c8.
//
// Solidity: function viewSeller(uint256 _ID) returns((address,string,uint256,string,string,uint256,uint256))
func (_SSMC *SSMCSession) ViewSeller(_ID *big.Int) (*types.Transaction, error) {
	return _SSMC.Contract.ViewSeller(&_SSMC.TransactOpts, _ID)
}

// ViewSeller is a paid mutator transaction binding the contract method 0xae4b59c8.
//
// Solidity: function viewSeller(uint256 _ID) returns((address,string,uint256,string,string,uint256,uint256))
func (_SSMC *SSMCTransactorSession) ViewSeller(_ID *big.Int) (*types.Transaction, error) {
	return _SSMC.Contract.ViewSeller(&_SSMC.TransactOpts, _ID)
}
