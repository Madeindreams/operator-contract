// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package operator

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

// SequencerLoanSequence is an auto generated low-level Go binding around an user-defined struct.
type SequencerLoanSequence struct {
	Assets  []common.Address
	Amounts []*big.Int
	Modes   []*big.Int
}

// OperatorMetaData contains all meta data concerning the Operator contract.
var OperatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPoolAddressesProvider\",\"name\":\"_addressProvider\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"Sequenced\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADDRESSES_PROVIDER\",\"outputs\":[{\"internalType\":\"contractIPoolAddressesProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"premiums\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"executeOperation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"modes\",\"type\":\"uint256[]\"}],\"internalType\":\"structSequencer.LoanSequence\",\"name\":\"loanData\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"tradeData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"useLoan\",\"type\":\"bool\"}],\"name\":\"executeSequence\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c060405234801562000010575f80fd5b506040516200208a3803806200208a8339818101604052810190620000369190620003e5565b80335f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603620000ab575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401620000a2919062000426565b60405180910390fd5b620000bc81620002ac60201b60201c565b508073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508073ffffffffffffffffffffffffffffffffffffffff1663026b1d5f6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156200013b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062000161919062000470565b73ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff1681525050508060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663026b1d5f6040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000240573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019062000266919062000470565b60025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050620004a0565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6200039c8262000371565b9050919050565b5f620003af8262000390565b9050919050565b620003c181620003a3565b8114620003cc575f80fd5b50565b5f81519050620003df81620003b6565b92915050565b5f60208284031215620003fd57620003fc6200036d565b5b5f6200040c84828501620003cf565b91505092915050565b620004208162000390565b82525050565b5f6020820190506200043b5f83018462000415565b92915050565b6200044c8162000390565b811462000457575f80fd5b50565b5f815190506200046a8162000441565b92915050565b5f602082840312156200048857620004876200036d565b5b5f62000497848285016200045a565b91505092915050565b60805160a051611bc8620004c25f395f61046301525f6101740152611bc85ff3fe608060405234801561000f575f80fd5b5060043610610086575f3560e01c80637535d246116100595780637535d246146100ea5780638da5cb5b14610108578063920f5c8414610126578063f2fde38b1461015657610086565b80630542975c1461008a57806351cff8d9146100a8578063715018a6146100c457806371c6f4cc146100ce575b5f80fd5b610092610172565b60405161009f9190610be1565b60405180910390f35b6100c260048036038101906100bd9190610c46565b610196565b005b6100cc61034d565b005b6100e860048036038101906100e39190611056565b610360565b005b6100f2610461565b6040516100ff91906110fe565b60405180910390f35b610110610485565b60405161011d9190611126565b60405180910390f35b610140600480360381019061013b9190611242565b6104ac565b60405161014d9190611347565b60405180910390f35b610170600480360381019061016b9190610c46565b61074e565b005b7f000000000000000000000000000000000000000000000000000000000000000081565b61019e6107d2565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361024f575f6101da610485565b73ffffffffffffffffffffffffffffffffffffffff16476040516101fd9061138d565b5f6040518083038185875af1925050503d805f8114610237576040519150601f19603f3d011682016040523d82523d5f602084013e61023c565b606091505b5050905080610249575f80fd5b5061034a565b8073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb610273610485565b8373ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016102ac9190611126565b602060405180830381865afa1580156102c7573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102eb91906113b5565b6040518363ffffffff1660e01b81526004016103089291906113ef565b6020604051808303815f875af1158015610324573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610348919061142a565b505b50565b6103556107d2565b61035e5f610859565b565b6103686107d2565b5f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f8490508215610419578173ffffffffffffffffffffffffffffffffffffffff1663ab9c4b5d30835f015184602001518560400151338a5f6040518863ffffffff1660e01b81526004016103e79796959493929190611683565b5f604051808303815f87803b1580156103fe575f80fd5b505af1158015610410573d5f803e3d5ffd5b5050505061045a565b6104228461091a565b7f418ec945bc0a5b00e014495f00d2406bedfc7690a64c24b6d490d63387dfd1ce336040516104519190611126565b60405180910390a15b5050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b5f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461053c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161053390611766565b60405180910390fd5b5f61058984848080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f820116905080830192505050505050506109c1565b90505f5b815f0151518110156105ea576105de825f015182815181106105b2576105b1611784565b5b6020026020010151836020015183815181106105d1576105d0611784565b5b6020026020010151610a01565b5080600101905061058d565b505f5b8b8b9050811015610704578b8b8281811061060b5761060a611784565b5b90506020020160208101906106209190610c46565b73ffffffffffffffffffffffffffffffffffffffff1663095ea7b360025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff168a8a858181106106705761066f611784565b5b905060200201358d8d8681811061068a57610689611784565b5b9050602002013561069b91906117de565b6040518363ffffffff1660e01b81526004016106b89291906113ef565b6020604051808303815f875af11580156106d4573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106f8919061142a565b508060010190506105ed565b507f418ec945bc0a5b00e014495f00d2406bedfc7690a64c24b6d490d63387dfd1ce856040516107349190611126565b60405180910390a160019150509998505050505050505050565b6107566107d2565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036107c6575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016107bd9190611126565b60405180910390fd5b6107cf81610859565b50565b6107da610b46565b73ffffffffffffffffffffffffffffffffffffffff166107f8610485565b73ffffffffffffffffffffffffffffffffffffffff16146108575761081b610b46565b6040517f118cdaa700000000000000000000000000000000000000000000000000000000815260040161084e9190611126565b60405180910390fd5b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f610924826109c1565b90505f5b815f01515181101561098557610979825f0151828151811061094d5761094c611784565b5b60200260200101518360200151838151811061096c5761096b611784565b5b6020026020010151610a01565b50806001019050610928565b507f418ec945bc0a5b00e014495f00d2406bedfc7690a64c24b6d490d63387dfd1ce336040516109b59190611126565b60405180910390a15050565b6109c9610b4d565b5f80838060200190518101906109df9190611a06565b9150915060405180604001604052808381526020018281525092505050919050565b60603073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610a71576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a6890611ac6565b60405180910390fd5b5f60608473ffffffffffffffffffffffffffffffffffffffff1684604051610a999190611b14565b5f604051808303815f865af19150503d805f8114610ad2576040519150601f19603f3d011682016040523d82523d5f602084013e610ad7565b606091505b50809250819350505081158015610aee57505f8151115b15610afb573d5f803e3d5ffd5b81610b3b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b3290611b74565b60405180910390fd5b809250505092915050565b5f33905090565b604051806040016040528060608152602001606081525090565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f610ba9610ba4610b9f84610b67565b610b86565b610b67565b9050919050565b5f610bba82610b8f565b9050919050565b5f610bcb82610bb0565b9050919050565b610bdb81610bc1565b82525050565b5f602082019050610bf45f830184610bd2565b92915050565b5f604051905090565b5f80fd5b5f80fd5b5f610c1582610b67565b9050919050565b610c2581610c0b565b8114610c2f575f80fd5b50565b5f81359050610c4081610c1c565b92915050565b5f60208284031215610c5b57610c5a610c03565b5b5f610c6884828501610c32565b91505092915050565b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610cbb82610c75565b810181811067ffffffffffffffff82111715610cda57610cd9610c85565b5b80604052505050565b5f610cec610bfa565b9050610cf88282610cb2565b919050565b5f80fd5b5f80fd5b5f67ffffffffffffffff821115610d1f57610d1e610c85565b5b602082029050602081019050919050565b5f80fd5b5f610d46610d4184610d05565b610ce3565b90508083825260208201905060208402830185811115610d6957610d68610d30565b5b835b81811015610d925780610d7e8882610c32565b845260208401935050602081019050610d6b565b5050509392505050565b5f82601f830112610db057610daf610d01565b5b8135610dc0848260208601610d34565b91505092915050565b5f67ffffffffffffffff821115610de357610de2610c85565b5b602082029050602081019050919050565b5f819050919050565b610e0681610df4565b8114610e10575f80fd5b50565b5f81359050610e2181610dfd565b92915050565b5f610e39610e3484610dc9565b610ce3565b90508083825260208201905060208402830185811115610e5c57610e5b610d30565b5b835b81811015610e855780610e718882610e13565b845260208401935050602081019050610e5e565b5050509392505050565b5f82601f830112610ea357610ea2610d01565b5b8135610eb3848260208601610e27565b91505092915050565b5f60608284031215610ed157610ed0610c71565b5b610edb6060610ce3565b90505f82013567ffffffffffffffff811115610efa57610ef9610cfd565b5b610f0684828501610d9c565b5f83015250602082013567ffffffffffffffff811115610f2957610f28610cfd565b5b610f3584828501610e8f565b602083015250604082013567ffffffffffffffff811115610f5957610f58610cfd565b5b610f6584828501610e8f565b60408301525092915050565b5f80fd5b5f67ffffffffffffffff821115610f8f57610f8e610c85565b5b610f9882610c75565b9050602081019050919050565b828183375f83830152505050565b5f610fc5610fc084610f75565b610ce3565b905082815260208101848484011115610fe157610fe0610f71565b5b610fec848285610fa5565b509392505050565b5f82601f83011261100857611007610d01565b5b8135611018848260208601610fb3565b91505092915050565b5f8115159050919050565b61103581611021565b811461103f575f80fd5b50565b5f813590506110508161102c565b92915050565b5f805f6060848603121561106d5761106c610c03565b5b5f84013567ffffffffffffffff81111561108a57611089610c07565b5b61109686828701610ebc565b935050602084013567ffffffffffffffff8111156110b7576110b6610c07565b5b6110c386828701610ff4565b92505060406110d486828701611042565b9150509250925092565b5f6110e882610bb0565b9050919050565b6110f8816110de565b82525050565b5f6020820190506111115f8301846110ef565b92915050565b61112081610c0b565b82525050565b5f6020820190506111395f830184611117565b92915050565b5f80fd5b5f8083601f84011261115857611157610d01565b5b8235905067ffffffffffffffff8111156111755761117461113f565b5b60208301915083602082028301111561119157611190610d30565b5b9250929050565b5f8083601f8401126111ad576111ac610d01565b5b8235905067ffffffffffffffff8111156111ca576111c961113f565b5b6020830191508360208202830111156111e6576111e5610d30565b5b9250929050565b5f8083601f84011261120257611201610d01565b5b8235905067ffffffffffffffff81111561121f5761121e61113f565b5b60208301915083600182028301111561123b5761123a610d30565b5b9250929050565b5f805f805f805f805f60a08a8c03121561125f5761125e610c03565b5b5f8a013567ffffffffffffffff81111561127c5761127b610c07565b5b6112888c828d01611143565b995099505060208a013567ffffffffffffffff8111156112ab576112aa610c07565b5b6112b78c828d01611198565b975097505060408a013567ffffffffffffffff8111156112da576112d9610c07565b5b6112e68c828d01611198565b955095505060606112f98c828d01610c32565b93505060808a013567ffffffffffffffff81111561131a57611319610c07565b5b6113268c828d016111ed565b92509250509295985092959850929598565b61134181611021565b82525050565b5f60208201905061135a5f830184611338565b92915050565b5f81905092915050565b50565b5f6113785f83611360565b91506113838261136a565b5f82019050919050565b5f6113978261136d565b9150819050919050565b5f815190506113af81610dfd565b92915050565b5f602082840312156113ca576113c9610c03565b5b5f6113d7848285016113a1565b91505092915050565b6113e981610df4565b82525050565b5f6040820190506114025f830185611117565b61140f60208301846113e0565b9392505050565b5f815190506114248161102c565b92915050565b5f6020828403121561143f5761143e610c03565b5b5f61144c84828501611416565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b61148781610c0b565b82525050565b5f611498838361147e565b60208301905092915050565b5f602082019050919050565b5f6114ba82611455565b6114c4818561145f565b93506114cf8361146f565b805f5b838110156114ff5781516114e6888261148d565b97506114f1836114a4565b9250506001810190506114d2565b5085935050505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b61153e81610df4565b82525050565b5f61154f8383611535565b60208301905092915050565b5f602082019050919050565b5f6115718261150c565b61157b8185611516565b935061158683611526565b805f5b838110156115b657815161159d8882611544565b97506115a88361155b565b925050600181019050611589565b5085935050505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b838110156115fa5780820151818401526020810190506115df565b5f8484015250505050565b5f61160f826115c3565b61161981856115cd565b93506116298185602086016115dd565b61163281610c75565b840191505092915050565b5f819050919050565b5f61ffff82169050919050565b5f61166d6116686116638461163d565b610b86565b611646565b9050919050565b61167d81611653565b82525050565b5f60e0820190506116965f83018a611117565b81810360208301526116a881896114b0565b905081810360408301526116bc8188611567565b905081810360608301526116d08187611567565b90506116df6080830186611117565b81810360a08301526116f18185611605565b905061170060c0830184611674565b98975050505050505050565b5f82825260208201905092915050565b7f796f7520617265206e6f742074686520706f6f6c0000000000000000000000005f82015250565b5f61175060148361170c565b915061175b8261171c565b602082019050919050565b5f6020820190508181035f83015261177d81611744565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6117e882610df4565b91506117f383610df4565b925082820190508082111561180b5761180a6117b1565b5b92915050565b5f8151905061181f81610c1c565b92915050565b5f61183761183284610d05565b610ce3565b9050808382526020820190506020840283018581111561185a57611859610d30565b5b835b81811015611883578061186f8882611811565b84526020840193505060208101905061185c565b5050509392505050565b5f82601f8301126118a1576118a0610d01565b5b81516118b1848260208601611825565b91505092915050565b5f67ffffffffffffffff8211156118d4576118d3610c85565b5b602082029050602081019050919050565b5f6118f76118f284610f75565b610ce3565b90508281526020810184848401111561191357611912610f71565b5b61191e8482856115dd565b509392505050565b5f82601f83011261193a57611939610d01565b5b815161194a8482602086016118e5565b91505092915050565b5f611965611960846118ba565b610ce3565b9050808382526020820190506020840283018581111561198857611987610d30565b5b835b818110156119cf57805167ffffffffffffffff8111156119ad576119ac610d01565b5b8086016119ba8982611926565b8552602085019450505060208101905061198a565b5050509392505050565b5f82601f8301126119ed576119ec610d01565b5b81516119fd848260208601611953565b91505092915050565b5f8060408385031215611a1c57611a1b610c03565b5b5f83015167ffffffffffffffff811115611a3957611a38610c07565b5b611a458582860161188d565b925050602083015167ffffffffffffffff811115611a6657611a65610c07565b5b611a72858286016119d9565b9150509250929050565b7f496e7465726e616c207265656e747279000000000000000000000000000000005f82015250565b5f611ab060108361170c565b9150611abb82611a7c565b602082019050919050565b5f6020820190508181035f830152611add81611aa4565b9050919050565b5f611aee826115c3565b611af88185611360565b9350611b088185602086016115dd565b80840191505092915050565b5f611b1f8284611ae4565b915081905092915050565b7f564d3a206578656375746543616c6c20726576657274656400000000000000005f82015250565b5f611b5e60188361170c565b9150611b6982611b2a565b602082019050919050565b5f6020820190508181035f830152611b8b81611b52565b905091905056fea264697066735822122000d5ca4dae961061adc38c01adf5c229acbc352249ff842b1a8a6c65ac0e4f3564736f6c63430008160033",
}

// OperatorABI is the input ABI used to generate the binding from.
// Deprecated: Use OperatorMetaData.ABI instead.
var OperatorABI = OperatorMetaData.ABI

// OperatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OperatorMetaData.Bin instead.
var OperatorBin = OperatorMetaData.Bin

// DeployOperator deploys a new Ethereum contract, binding an instance of Operator to it.
func DeployOperator(auth *bind.TransactOpts, backend bind.ContractBackend, _addressProvider common.Address) (common.Address, *types.Transaction, *Operator, error) {
	parsed, err := OperatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OperatorBin), backend, _addressProvider)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Operator{OperatorCaller: OperatorCaller{contract: contract}, OperatorTransactor: OperatorTransactor{contract: contract}, OperatorFilterer: OperatorFilterer{contract: contract}}, nil
}

// Operator is an auto generated Go binding around an Ethereum contract.
type Operator struct {
	OperatorCaller     // Read-only binding to the contract
	OperatorTransactor // Write-only binding to the contract
	OperatorFilterer   // Log filterer for contract events
}

// OperatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type OperatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OperatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OperatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OperatorSession struct {
	Contract     *Operator         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OperatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OperatorCallerSession struct {
	Contract *OperatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// OperatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OperatorTransactorSession struct {
	Contract     *OperatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// OperatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type OperatorRaw struct {
	Contract *Operator // Generic contract binding to access the raw methods on
}

// OperatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OperatorCallerRaw struct {
	Contract *OperatorCaller // Generic read-only contract binding to access the raw methods on
}

// OperatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OperatorTransactorRaw struct {
	Contract *OperatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOperator creates a new instance of Operator, bound to a specific deployed contract.
func NewOperator(address common.Address, backend bind.ContractBackend) (*Operator, error) {
	contract, err := bindOperator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Operator{OperatorCaller: OperatorCaller{contract: contract}, OperatorTransactor: OperatorTransactor{contract: contract}, OperatorFilterer: OperatorFilterer{contract: contract}}, nil
}

// NewOperatorCaller creates a new read-only instance of Operator, bound to a specific deployed contract.
func NewOperatorCaller(address common.Address, caller bind.ContractCaller) (*OperatorCaller, error) {
	contract, err := bindOperator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OperatorCaller{contract: contract}, nil
}

// NewOperatorTransactor creates a new write-only instance of Operator, bound to a specific deployed contract.
func NewOperatorTransactor(address common.Address, transactor bind.ContractTransactor) (*OperatorTransactor, error) {
	contract, err := bindOperator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OperatorTransactor{contract: contract}, nil
}

// NewOperatorFilterer creates a new log filterer instance of Operator, bound to a specific deployed contract.
func NewOperatorFilterer(address common.Address, filterer bind.ContractFilterer) (*OperatorFilterer, error) {
	contract, err := bindOperator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OperatorFilterer{contract: contract}, nil
}

// bindOperator binds a generic wrapper to an already deployed contract.
func bindOperator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OperatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Operator *OperatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Operator.Contract.OperatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Operator *OperatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Operator.Contract.OperatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Operator *OperatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Operator.Contract.OperatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Operator *OperatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Operator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Operator *OperatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Operator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Operator *OperatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Operator.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_Operator *OperatorCaller) ADDRESSESPROVIDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "ADDRESSES_PROVIDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_Operator *OperatorSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _Operator.Contract.ADDRESSESPROVIDER(&_Operator.CallOpts)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_Operator *OperatorCallerSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _Operator.Contract.ADDRESSESPROVIDER(&_Operator.CallOpts)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_Operator *OperatorCaller) POOL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "POOL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_Operator *OperatorSession) POOL() (common.Address, error) {
	return _Operator.Contract.POOL(&_Operator.CallOpts)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_Operator *OperatorCallerSession) POOL() (common.Address, error) {
	return _Operator.Contract.POOL(&_Operator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Operator *OperatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Operator *OperatorSession) Owner() (common.Address, error) {
	return _Operator.Contract.Owner(&_Operator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Operator *OperatorCallerSession) Owner() (common.Address, error) {
	return _Operator.Contract.Owner(&_Operator.CallOpts)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x920f5c84.
//
// Solidity: function executeOperation(address[] assets, uint256[] amounts, uint256[] premiums, address initiator, bytes params) returns(bool)
func (_Operator *OperatorTransactor) ExecuteOperation(opts *bind.TransactOpts, assets []common.Address, amounts []*big.Int, premiums []*big.Int, initiator common.Address, params []byte) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "executeOperation", assets, amounts, premiums, initiator, params)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x920f5c84.
//
// Solidity: function executeOperation(address[] assets, uint256[] amounts, uint256[] premiums, address initiator, bytes params) returns(bool)
func (_Operator *OperatorSession) ExecuteOperation(assets []common.Address, amounts []*big.Int, premiums []*big.Int, initiator common.Address, params []byte) (*types.Transaction, error) {
	return _Operator.Contract.ExecuteOperation(&_Operator.TransactOpts, assets, amounts, premiums, initiator, params)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x920f5c84.
//
// Solidity: function executeOperation(address[] assets, uint256[] amounts, uint256[] premiums, address initiator, bytes params) returns(bool)
func (_Operator *OperatorTransactorSession) ExecuteOperation(assets []common.Address, amounts []*big.Int, premiums []*big.Int, initiator common.Address, params []byte) (*types.Transaction, error) {
	return _Operator.Contract.ExecuteOperation(&_Operator.TransactOpts, assets, amounts, premiums, initiator, params)
}

// ExecuteSequence is a paid mutator transaction binding the contract method 0x71c6f4cc.
//
// Solidity: function executeSequence((address[],uint256[],uint256[]) loanData, bytes tradeData, bool useLoan) returns()
func (_Operator *OperatorTransactor) ExecuteSequence(opts *bind.TransactOpts, loanData SequencerLoanSequence, tradeData []byte, useLoan bool) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "executeSequence", loanData, tradeData, useLoan)
}

// ExecuteSequence is a paid mutator transaction binding the contract method 0x71c6f4cc.
//
// Solidity: function executeSequence((address[],uint256[],uint256[]) loanData, bytes tradeData, bool useLoan) returns()
func (_Operator *OperatorSession) ExecuteSequence(loanData SequencerLoanSequence, tradeData []byte, useLoan bool) (*types.Transaction, error) {
	return _Operator.Contract.ExecuteSequence(&_Operator.TransactOpts, loanData, tradeData, useLoan)
}

// ExecuteSequence is a paid mutator transaction binding the contract method 0x71c6f4cc.
//
// Solidity: function executeSequence((address[],uint256[],uint256[]) loanData, bytes tradeData, bool useLoan) returns()
func (_Operator *OperatorTransactorSession) ExecuteSequence(loanData SequencerLoanSequence, tradeData []byte, useLoan bool) (*types.Transaction, error) {
	return _Operator.Contract.ExecuteSequence(&_Operator.TransactOpts, loanData, tradeData, useLoan)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Operator *OperatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Operator *OperatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Operator.Contract.RenounceOwnership(&_Operator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Operator *OperatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Operator.Contract.RenounceOwnership(&_Operator.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Operator *OperatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Operator *OperatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Operator.Contract.TransferOwnership(&_Operator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Operator *OperatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Operator.Contract.TransferOwnership(&_Operator.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _asset) returns()
func (_Operator *OperatorTransactor) Withdraw(opts *bind.TransactOpts, _asset common.Address) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "withdraw", _asset)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _asset) returns()
func (_Operator *OperatorSession) Withdraw(_asset common.Address) (*types.Transaction, error) {
	return _Operator.Contract.Withdraw(&_Operator.TransactOpts, _asset)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _asset) returns()
func (_Operator *OperatorTransactorSession) Withdraw(_asset common.Address) (*types.Transaction, error) {
	return _Operator.Contract.Withdraw(&_Operator.TransactOpts, _asset)
}

// OperatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Operator contract.
type OperatorOwnershipTransferredIterator struct {
	Event *OperatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OperatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorOwnershipTransferred)
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
		it.Event = new(OperatorOwnershipTransferred)
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
func (it *OperatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorOwnershipTransferred represents a OwnershipTransferred event raised by the Operator contract.
type OperatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Operator *OperatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OperatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Operator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OperatorOwnershipTransferredIterator{contract: _Operator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Operator *OperatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OperatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Operator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorOwnershipTransferred)
				if err := _Operator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Operator *OperatorFilterer) ParseOwnershipTransferred(log types.Log) (*OperatorOwnershipTransferred, error) {
	event := new(OperatorOwnershipTransferred)
	if err := _Operator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorSequencedIterator is returned from FilterSequenced and is used to iterate over the raw logs and unpacked data for Sequenced events raised by the Operator contract.
type OperatorSequencedIterator struct {
	Event *OperatorSequenced // Event containing the contract specifics and raw log

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
func (it *OperatorSequencedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorSequenced)
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
		it.Event = new(OperatorSequenced)
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
func (it *OperatorSequencedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorSequencedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorSequenced represents a Sequenced event raised by the Operator contract.
type OperatorSequenced struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSequenced is a free log retrieval operation binding the contract event 0x418ec945bc0a5b00e014495f00d2406bedfc7690a64c24b6d490d63387dfd1ce.
//
// Solidity: event Sequenced(address user)
func (_Operator *OperatorFilterer) FilterSequenced(opts *bind.FilterOpts) (*OperatorSequencedIterator, error) {

	logs, sub, err := _Operator.contract.FilterLogs(opts, "Sequenced")
	if err != nil {
		return nil, err
	}
	return &OperatorSequencedIterator{contract: _Operator.contract, event: "Sequenced", logs: logs, sub: sub}, nil
}

// WatchSequenced is a free log subscription operation binding the contract event 0x418ec945bc0a5b00e014495f00d2406bedfc7690a64c24b6d490d63387dfd1ce.
//
// Solidity: event Sequenced(address user)
func (_Operator *OperatorFilterer) WatchSequenced(opts *bind.WatchOpts, sink chan<- *OperatorSequenced) (event.Subscription, error) {

	logs, sub, err := _Operator.contract.WatchLogs(opts, "Sequenced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorSequenced)
				if err := _Operator.contract.UnpackLog(event, "Sequenced", log); err != nil {
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

// ParseSequenced is a log parse operation binding the contract event 0x418ec945bc0a5b00e014495f00d2406bedfc7690a64c24b6d490d63387dfd1ce.
//
// Solidity: event Sequenced(address user)
func (_Operator *OperatorFilterer) ParseSequenced(log types.Log) (*OperatorSequenced, error) {
	event := new(OperatorSequenced)
	if err := _Operator.contract.UnpackLog(event, "Sequenced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
