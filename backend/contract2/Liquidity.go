// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract2

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

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SwapTokenToToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxToken\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"getLiquidityAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"inputAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"getSwapRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"liquidityProvider\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lpTokenAmount\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minToken\",\"type\":\"uint256\"}],\"name\":\"swapCoinToToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minCoin\",\"type\":\"uint256\"}],\"name\":\"swapTokenToCoin\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162002ada38038062002ada83398181016040528101906200003791906200017b565b6040518060400160405280600781526020017f6c70546f6b656e000000000000000000000000000000000000000000000000008152506040518060400160405280600281526020017f4c500000000000000000000000000000000000000000000000000000000000008152508160039081620000b4919062000427565b508060049081620000c6919062000427565b50505080600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200050e565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620001438262000116565b9050919050565b620001558162000136565b81146200016157600080fd5b50565b60008151905062000175816200014a565b92915050565b60006020828403121562000194576200019362000111565b5b6000620001a48482850162000164565b91505092915050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200022f57607f821691505b602082108103620002455762000244620001e7565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620002af7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000270565b620002bb868362000270565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006200030862000302620002fc84620002d3565b620002dd565b620002d3565b9050919050565b6000819050919050565b6200032483620002e7565b6200033c62000333826200030f565b8484546200027d565b825550505050565b600090565b6200035362000344565b6200036081848462000319565b505050565b5b8181101562000388576200037c60008262000349565b60018101905062000366565b5050565b601f821115620003d757620003a1816200024b565b620003ac8462000260565b81016020851015620003bc578190505b620003d4620003cb8562000260565b83018262000365565b50505b505050565b600082821c905092915050565b6000620003fc60001984600802620003dc565b1980831691505092915050565b6000620004178383620003e9565b9150826002028217905092915050565b6200043282620001ad565b67ffffffffffffffff8111156200044e576200044d620001b8565b5b6200045a825462000216565b620004678282856200038c565b600060209050601f8311600181146200049f57600084156200048a578287015190505b62000496858262000409565b86555062000506565b601f198416620004af866200024b565b60005b82811015620004d957848901518255600182019150602085019450602081019050620004b2565b86831015620004f95784890151620004f5601f891682620003e9565b8355505b6001600288020188555050505b505050505050565b6125bc806200051e6000396000f3fe6080604052600436106101145760003560e01c80636f7e9362116100a0578063a2687c6311610064578063a2687c63146103bf578063a457c2d7146103fc578063a5cb0b8414610439578063a9059cbb14610455578063dd62ed3e1461049257610114565b80636f7e9362146102b457806370a08231146102f157806395d89b411461032e5780639c8f9f2314610359578063a1f408b41461038257610114565b806323b872dd116100e757806323b872dd146101d7578063313ce56714610214578063395093511461023f578063411df6f31461027c57806351c6590a1461029857610114565b806306fdde0314610119578063095ea7b31461014457806312065fe01461018157806318160ddd146101ac575b600080fd5b34801561012557600080fd5b5061012e6104cf565b60405161013b9190611976565b60405180910390f35b34801561015057600080fd5b5061016b60048036038101906101669190611a31565b610561565b6040516101789190611a8c565b60405180910390f35b34801561018d57600080fd5b50610196610584565b6040516101a39190611ab6565b60405180910390f35b3480156101b857600080fd5b506101c161058c565b6040516101ce9190611ab6565b60405180910390f35b3480156101e357600080fd5b506101fe60048036038101906101f99190611ad1565b610596565b60405161020b9190611a8c565b60405180910390f35b34801561022057600080fd5b506102296105c5565b6040516102369190611b40565b60405180910390f35b34801561024b57600080fd5b5061026660048036038101906102619190611a31565b6105ce565b6040516102739190611a8c565b60405180910390f35b61029660048036038101906102919190611b5b565b610605565b005b6102b260048036038101906102ad9190611b9b565b61081c565b005b3480156102c057600080fd5b506102db60048036038101906102d69190611bc8565b610ace565b6040516102e89190611ab6565b60405180910390f35b3480156102fd57600080fd5b5061031860048036038101906103139190611c1b565b610b06565b6040516103259190611ab6565b60405180910390f35b34801561033a57600080fd5b50610343610b4e565b6040516103509190611976565b60405180910390f35b34801561036557600080fd5b50610380600480360381019061037b9190611b9b565b610be0565b005b34801561038e57600080fd5b506103a960048036038101906103a49190611c1b565b610db5565b6040516103b69190611ab6565b60405180910390f35b3480156103cb57600080fd5b506103e660048036038101906103e19190611c1b565b610dfe565b6040516103f39190611ab6565b60405180910390f35b34801561040857600080fd5b50610423600480360381019061041e9190611a31565b610e16565b6040516104309190611a8c565b60405180910390f35b610453600480360381019061044e9190611b9b565b610e8d565b005b34801561046157600080fd5b5061047c60048036038101906104779190611a31565b61103c565b6040516104899190611a8c565b60405180910390f35b34801561049e57600080fd5b506104b960048036038101906104b49190611c48565b61105f565b6040516104c69190611ab6565b60405180910390f35b6060600380546104de90611cb7565b80601f016020809104026020016040519081016040528092919081815260200182805461050a90611cb7565b80156105575780601f1061052c57610100808354040283529160200191610557565b820191906000526020600020905b81548152906001019060200180831161053a57829003601f168201915b5050505050905090565b60008061056c6110e6565b90506105798185856110ee565b600191505092915050565b600047905090565b6000600254905090565b6000806105a16110e6565b90506105ae8582856112b7565b6105b9858585611343565b60019150509392505050565b60006012905090565b6000806105d96110e6565b90506105fa8185856105eb858961105f565b6105f59190611d17565b6110ee565b600191505092915050565b60006106ad83600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016106669190611d5a565b602060405180830381865afa158015610683573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106a79190611d8a565b47610ace565b9050818110156106f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e990611e03565b60405180910390fd5b7f66c9b9a0940362117f399802fc8dee32d47e0f132432c3ccebe00af0206144d533308560405161072593929190611e23565b60405180910390a1600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330866040518463ffffffff1660e01b815260040161078c93929190611e23565b6020604051808303816000875af11580156107ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107cf9190611e86565b503373ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015610816573d6000803e3d6000fd5b50505050565b600061082661058c565b90506000811115610a10576000344761083f9190611eb3565b90506000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b815260040161089e9190611d5a565b602060405180830381865afa1580156108bb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108df9190611d8a565b905060008282346108f09190611ee7565b6108fa9190611f58565b90508085101561093f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161093690611fd5565b60405180910390fd5b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b815260040161099e93929190611e23565b6020604051808303816000875af11580156109bd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109e19190611e86565b5060008334866109f19190611ee7565b6109fb9190611f58565b9050610a0733826115b9565b50505050610aca565b60008290506000479050610a2433826115b9565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330856040518463ffffffff1660e01b8152600401610a8393929190611e23565b6020604051808303816000875af1158015610aa2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ac69190611e86565b5050505b5050565b6000808483610add9190611ee7565b905060008585610aed9190611d17565b90508082610afb9190611f58565b925050509392505050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b606060048054610b5d90611cb7565b80601f0160208091040260200160405190810160405280929190818152602001828054610b8990611cb7565b8015610bd65780601f10610bab57610100808354040283529160200191610bd6565b820191906000526020600020905b815481529060010190602001808311610bb957829003601f168201915b5050505050905090565b6000610bea61058c565b90506000814784610bfb9190611ee7565b610c059190611f58565b9050600082600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401610c659190611d5a565b602060405180830381865afa158015610c82573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ca69190611d8a565b85610cb19190611ee7565b610cbb9190611f58565b9050610cc7338561170f565b3373ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f19350505050158015610d0d573d6000803e3d6000fd5b50600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b8152600401610d6b929190611ff5565b6020604051808303816000875af1158015610d8a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dae9190611e86565b5050505050565b6000600660008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60066020528060005260406000206000915090505481565b600080610e216110e6565b90506000610e2f828661105f565b905083811015610e74576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e6b90612090565b60405180910390fd5b610e8182868684036110ee565b60019250505092915050565b600034905060008147610ea09190611eb3565b90506000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401610eff9190611d5a565b602060405180830381865afa158015610f1c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f409190611d8a565b90506000610f4f348484610ace565b905084811015610f94576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f8b90611e03565b60405180910390fd5b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b8152600401610ff1929190611ff5565b6020604051808303816000875af1158015611010573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110349190611e86565b505050505050565b6000806110476110e6565b9050611054818585611343565b600191505092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361115d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161115490612122565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036111cc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111c3906121b4565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040516112aa9190611ab6565b60405180910390a3505050565b60006112c3848461105f565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff811461133d578181101561132f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161132690612220565b60405180910390fd5b61133c84848484036110ee565b5b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036113b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113a9906122b2565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611421576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161141890612344565b60405180910390fd5b61142c8383836118dc565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050818110156114b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114a9906123d6565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516115a09190611ab6565b60405180910390a36115b38484846118e1565b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611628576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161161f90612442565b60405180910390fd5b611634600083836118dc565b80600260008282546116469190611d17565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516116f79190611ab6565b60405180910390a361170b600083836118e1565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361177e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611775906124d4565b60405180910390fd5b61178a826000836118dc565b60008060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015611810576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161180790612566565b60405180910390fd5b8181036000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600260008282540392505081905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516118c39190611ab6565b60405180910390a36118d7836000846118e1565b505050565b505050565b505050565b600081519050919050565b600082825260208201905092915050565b60005b83811015611920578082015181840152602081019050611905565b60008484015250505050565b6000601f19601f8301169050919050565b6000611948826118e6565b61195281856118f1565b9350611962818560208601611902565b61196b8161192c565b840191505092915050565b60006020820190508181036000830152611990818461193d565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006119c88261199d565b9050919050565b6119d8816119bd565b81146119e357600080fd5b50565b6000813590506119f5816119cf565b92915050565b6000819050919050565b611a0e816119fb565b8114611a1957600080fd5b50565b600081359050611a2b81611a05565b92915050565b60008060408385031215611a4857611a47611998565b5b6000611a56858286016119e6565b9250506020611a6785828601611a1c565b9150509250929050565b60008115159050919050565b611a8681611a71565b82525050565b6000602082019050611aa16000830184611a7d565b92915050565b611ab0816119fb565b82525050565b6000602082019050611acb6000830184611aa7565b92915050565b600080600060608486031215611aea57611ae9611998565b5b6000611af8868287016119e6565b9350506020611b09868287016119e6565b9250506040611b1a86828701611a1c565b9150509250925092565b600060ff82169050919050565b611b3a81611b24565b82525050565b6000602082019050611b556000830184611b31565b92915050565b60008060408385031215611b7257611b71611998565b5b6000611b8085828601611a1c565b9250506020611b9185828601611a1c565b9150509250929050565b600060208284031215611bb157611bb0611998565b5b6000611bbf84828501611a1c565b91505092915050565b600080600060608486031215611be157611be0611998565b5b6000611bef86828701611a1c565b9350506020611c0086828701611a1c565b9250506040611c1186828701611a1c565b9150509250925092565b600060208284031215611c3157611c30611998565b5b6000611c3f848285016119e6565b91505092915050565b60008060408385031215611c5f57611c5e611998565b5b6000611c6d858286016119e6565b9250506020611c7e858286016119e6565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680611ccf57607f821691505b602082108103611ce257611ce1611c88565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611d22826119fb565b9150611d2d836119fb565b9250828201905080821115611d4557611d44611ce8565b5b92915050565b611d54816119bd565b82525050565b6000602082019050611d6f6000830184611d4b565b92915050565b600081519050611d8481611a05565b92915050565b600060208284031215611da057611d9f611998565b5b6000611dae84828501611d75565b91505092915050565b7f6c61636b206f6620616d6f756e74000000000000000000000000000000000000600082015250565b6000611ded600e836118f1565b9150611df882611db7565b602082019050919050565b60006020820190508181036000830152611e1c81611de0565b9050919050565b6000606082019050611e386000830186611d4b565b611e456020830185611d4b565b611e526040830184611aa7565b949350505050565b611e6381611a71565b8114611e6e57600080fd5b50565b600081519050611e8081611e5a565b92915050565b600060208284031215611e9c57611e9b611998565b5b6000611eaa84828501611e71565b91505092915050565b6000611ebe826119fb565b9150611ec9836119fb565b9250828203905081811115611ee157611ee0611ce8565b5b92915050565b6000611ef2826119fb565b9150611efd836119fb565b9250828202611f0b816119fb565b91508282048414831517611f2257611f21611ce8565b5b5092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000611f63826119fb565b9150611f6e836119fb565b925082611f7e57611f7d611f29565b5b828204905092915050565b7f2d2d2d2d2d3f0000000000000000000000000000000000000000000000000000600082015250565b6000611fbf6006836118f1565b9150611fca82611f89565b602082019050919050565b60006020820190508181036000830152611fee81611fb2565b9050919050565b600060408201905061200a6000830185611d4b565b6120176020830184611aa7565b9392505050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b600061207a6025836118f1565b91506120858261201e565b604082019050919050565b600060208201905081810360008301526120a98161206d565b9050919050565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b600061210c6024836118f1565b9150612117826120b0565b604082019050919050565b6000602082019050818103600083015261213b816120ff565b9050919050565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b600061219e6022836118f1565b91506121a982612142565b604082019050919050565b600060208201905081810360008301526121cd81612191565b9050919050565b7f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000600082015250565b600061220a601d836118f1565b9150612215826121d4565b602082019050919050565b60006020820190508181036000830152612239816121fd565b9050919050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b600061229c6025836118f1565b91506122a782612240565b604082019050919050565b600060208201905081810360008301526122cb8161228f565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b600061232e6023836118f1565b9150612339826122d2565b604082019050919050565b6000602082019050818103600083015261235d81612321565b9050919050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b60006123c06026836118f1565b91506123cb82612364565b604082019050919050565b600060208201905081810360008301526123ef816123b3565b9050919050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b600061242c601f836118f1565b9150612437826123f6565b602082019050919050565b6000602082019050818103600083015261245b8161241f565b9050919050565b7f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360008201527f7300000000000000000000000000000000000000000000000000000000000000602082015250565b60006124be6021836118f1565b91506124c982612462565b604082019050919050565b600060208201905081810360008301526124ed816124b1565b9050919050565b7f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60008201527f6365000000000000000000000000000000000000000000000000000000000000602082015250565b60006125506022836118f1565b915061255b826124f4565b604082019050919050565b6000602082019050818103600083015261257f81612543565b905091905056fea264697066735822122072c4163cd83cdd5d4012513fdf50eb785fc3d4421e9f3bf583c114d6366ea11764736f6c63430008110033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Contracts *ContractsCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Contracts *ContractsSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Contracts.Contract.Allowance(&_Contracts.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Contracts *ContractsCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Contracts.Contract.Allowance(&_Contracts.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Contracts *ContractsCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Contracts *ContractsSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Contracts.Contract.BalanceOf(&_Contracts.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Contracts *ContractsCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Contracts.Contract.BalanceOf(&_Contracts.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Contracts *ContractsCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Contracts *ContractsSession) Decimals() (uint8, error) {
	return _Contracts.Contract.Decimals(&_Contracts.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Contracts *ContractsCallerSession) Decimals() (uint8, error) {
	return _Contracts.Contract.Decimals(&_Contracts.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Contracts *ContractsCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Contracts *ContractsSession) GetBalance() (*big.Int, error) {
	return _Contracts.Contract.GetBalance(&_Contracts.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Contracts *ContractsCallerSession) GetBalance() (*big.Int, error) {
	return _Contracts.Contract.GetBalance(&_Contracts.CallOpts)
}

// GetLiquidityAmount is a free data retrieval call binding the contract method 0xa1f408b4.
//
// Solidity: function getLiquidityAmount(address provider) view returns(uint256)
func (_Contracts *ContractsCaller) GetLiquidityAmount(opts *bind.CallOpts, provider common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getLiquidityAmount", provider)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLiquidityAmount is a free data retrieval call binding the contract method 0xa1f408b4.
//
// Solidity: function getLiquidityAmount(address provider) view returns(uint256)
func (_Contracts *ContractsSession) GetLiquidityAmount(provider common.Address) (*big.Int, error) {
	return _Contracts.Contract.GetLiquidityAmount(&_Contracts.CallOpts, provider)
}

// GetLiquidityAmount is a free data retrieval call binding the contract method 0xa1f408b4.
//
// Solidity: function getLiquidityAmount(address provider) view returns(uint256)
func (_Contracts *ContractsCallerSession) GetLiquidityAmount(provider common.Address) (*big.Int, error) {
	return _Contracts.Contract.GetLiquidityAmount(&_Contracts.CallOpts, provider)
}

// GetSwapRatio is a free data retrieval call binding the contract method 0x6f7e9362.
//
// Solidity: function getSwapRatio(uint256 inputAmount, uint256 x, uint256 y) pure returns(uint256)
func (_Contracts *ContractsCaller) GetSwapRatio(opts *bind.CallOpts, inputAmount *big.Int, x *big.Int, y *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getSwapRatio", inputAmount, x, y)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSwapRatio is a free data retrieval call binding the contract method 0x6f7e9362.
//
// Solidity: function getSwapRatio(uint256 inputAmount, uint256 x, uint256 y) pure returns(uint256)
func (_Contracts *ContractsSession) GetSwapRatio(inputAmount *big.Int, x *big.Int, y *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetSwapRatio(&_Contracts.CallOpts, inputAmount, x, y)
}

// GetSwapRatio is a free data retrieval call binding the contract method 0x6f7e9362.
//
// Solidity: function getSwapRatio(uint256 inputAmount, uint256 x, uint256 y) pure returns(uint256)
func (_Contracts *ContractsCallerSession) GetSwapRatio(inputAmount *big.Int, x *big.Int, y *big.Int) (*big.Int, error) {
	return _Contracts.Contract.GetSwapRatio(&_Contracts.CallOpts, inputAmount, x, y)
}

// LiquidityProvider is a free data retrieval call binding the contract method 0xa2687c63.
//
// Solidity: function liquidityProvider(address ) view returns(uint256)
func (_Contracts *ContractsCaller) LiquidityProvider(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "liquidityProvider", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidityProvider is a free data retrieval call binding the contract method 0xa2687c63.
//
// Solidity: function liquidityProvider(address ) view returns(uint256)
func (_Contracts *ContractsSession) LiquidityProvider(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.LiquidityProvider(&_Contracts.CallOpts, arg0)
}

// LiquidityProvider is a free data retrieval call binding the contract method 0xa2687c63.
//
// Solidity: function liquidityProvider(address ) view returns(uint256)
func (_Contracts *ContractsCallerSession) LiquidityProvider(arg0 common.Address) (*big.Int, error) {
	return _Contracts.Contract.LiquidityProvider(&_Contracts.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contracts *ContractsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contracts *ContractsSession) Name() (string, error) {
	return _Contracts.Contract.Name(&_Contracts.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contracts *ContractsCallerSession) Name() (string, error) {
	return _Contracts.Contract.Name(&_Contracts.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contracts *ContractsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contracts *ContractsSession) Symbol() (string, error) {
	return _Contracts.Contract.Symbol(&_Contracts.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contracts *ContractsCallerSession) Symbol() (string, error) {
	return _Contracts.Contract.Symbol(&_Contracts.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contracts *ContractsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contracts *ContractsSession) TotalSupply() (*big.Int, error) {
	return _Contracts.Contract.TotalSupply(&_Contracts.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contracts *ContractsCallerSession) TotalSupply() (*big.Int, error) {
	return _Contracts.Contract.TotalSupply(&_Contracts.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x51c6590a.
//
// Solidity: function addLiquidity(uint256 _maxToken) payable returns()
func (_Contracts *ContractsTransactor) AddLiquidity(opts *bind.TransactOpts, _maxToken *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "addLiquidity", _maxToken)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x51c6590a.
//
// Solidity: function addLiquidity(uint256 _maxToken) payable returns()
func (_Contracts *ContractsSession) AddLiquidity(_maxToken *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AddLiquidity(&_Contracts.TransactOpts, _maxToken)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x51c6590a.
//
// Solidity: function addLiquidity(uint256 _maxToken) payable returns()
func (_Contracts *ContractsTransactorSession) AddLiquidity(_maxToken *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.AddLiquidity(&_Contracts.TransactOpts, _maxToken)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Contracts *ContractsTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Contracts *ContractsSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Approve(&_Contracts.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Contracts *ContractsTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Approve(&_Contracts.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Contracts *ContractsTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Contracts *ContractsSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.DecreaseAllowance(&_Contracts.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Contracts *ContractsTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.DecreaseAllowance(&_Contracts.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Contracts *ContractsTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Contracts *ContractsSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.IncreaseAllowance(&_Contracts.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Contracts *ContractsTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.IncreaseAllowance(&_Contracts.TransactOpts, spender, addedValue)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 _lpTokenAmount) returns()
func (_Contracts *ContractsTransactor) RemoveLiquidity(opts *bind.TransactOpts, _lpTokenAmount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "removeLiquidity", _lpTokenAmount)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 _lpTokenAmount) returns()
func (_Contracts *ContractsSession) RemoveLiquidity(_lpTokenAmount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.RemoveLiquidity(&_Contracts.TransactOpts, _lpTokenAmount)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x9c8f9f23.
//
// Solidity: function removeLiquidity(uint256 _lpTokenAmount) returns()
func (_Contracts *ContractsTransactorSession) RemoveLiquidity(_lpTokenAmount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.RemoveLiquidity(&_Contracts.TransactOpts, _lpTokenAmount)
}

// SwapCoinToToken is a paid mutator transaction binding the contract method 0xa5cb0b84.
//
// Solidity: function swapCoinToToken(uint256 _minToken) payable returns()
func (_Contracts *ContractsTransactor) SwapCoinToToken(opts *bind.TransactOpts, _minToken *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "swapCoinToToken", _minToken)
}

// SwapCoinToToken is a paid mutator transaction binding the contract method 0xa5cb0b84.
//
// Solidity: function swapCoinToToken(uint256 _minToken) payable returns()
func (_Contracts *ContractsSession) SwapCoinToToken(_minToken *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SwapCoinToToken(&_Contracts.TransactOpts, _minToken)
}

// SwapCoinToToken is a paid mutator transaction binding the contract method 0xa5cb0b84.
//
// Solidity: function swapCoinToToken(uint256 _minToken) payable returns()
func (_Contracts *ContractsTransactorSession) SwapCoinToToken(_minToken *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SwapCoinToToken(&_Contracts.TransactOpts, _minToken)
}

// SwapTokenToCoin is a paid mutator transaction binding the contract method 0x411df6f3.
//
// Solidity: function swapTokenToCoin(uint256 _tokenAmount, uint256 _minCoin) payable returns()
func (_Contracts *ContractsTransactor) SwapTokenToCoin(opts *bind.TransactOpts, _tokenAmount *big.Int, _minCoin *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "swapTokenToCoin", _tokenAmount, _minCoin)
}

// SwapTokenToCoin is a paid mutator transaction binding the contract method 0x411df6f3.
//
// Solidity: function swapTokenToCoin(uint256 _tokenAmount, uint256 _minCoin) payable returns()
func (_Contracts *ContractsSession) SwapTokenToCoin(_tokenAmount *big.Int, _minCoin *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SwapTokenToCoin(&_Contracts.TransactOpts, _tokenAmount, _minCoin)
}

// SwapTokenToCoin is a paid mutator transaction binding the contract method 0x411df6f3.
//
// Solidity: function swapTokenToCoin(uint256 _tokenAmount, uint256 _minCoin) payable returns()
func (_Contracts *ContractsTransactorSession) SwapTokenToCoin(_tokenAmount *big.Int, _minCoin *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SwapTokenToCoin(&_Contracts.TransactOpts, _tokenAmount, _minCoin)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Contracts *ContractsTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Contracts *ContractsSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Transfer(&_Contracts.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Contracts *ContractsTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Transfer(&_Contracts.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Contracts *ContractsTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Contracts *ContractsSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.TransferFrom(&_Contracts.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Contracts *ContractsTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.TransferFrom(&_Contracts.TransactOpts, from, to, amount)
}

// ContractsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Contracts contract.
type ContractsApprovalIterator struct {
	Event *ContractsApproval // Event containing the contract specifics and raw log

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
func (it *ContractsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsApproval)
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
		it.Event = new(ContractsApproval)
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
func (it *ContractsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsApproval represents a Approval event raised by the Contracts contract.
type ContractsApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Contracts *ContractsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ContractsApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ContractsApprovalIterator{contract: _Contracts.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Contracts *ContractsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ContractsApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsApproval)
				if err := _Contracts.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseApproval(log types.Log) (*ContractsApproval, error) {
	event := new(ContractsApproval)
	if err := _Contracts.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsSwapTokenToTokenIterator is returned from FilterSwapTokenToToken and is used to iterate over the raw logs and unpacked data for SwapTokenToToken events raised by the Contracts contract.
type ContractsSwapTokenToTokenIterator struct {
	Event *ContractsSwapTokenToToken // Event containing the contract specifics and raw log

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
func (it *ContractsSwapTokenToTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsSwapTokenToToken)
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
		it.Event = new(ContractsSwapTokenToToken)
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
func (it *ContractsSwapTokenToTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsSwapTokenToTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsSwapTokenToToken represents a SwapTokenToToken event raised by the Contracts contract.
type ContractsSwapTokenToToken struct {
	Caller common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSwapTokenToToken is a free log retrieval operation binding the contract event 0x66c9b9a0940362117f399802fc8dee32d47e0f132432c3ccebe00af0206144d5.
//
// Solidity: event SwapTokenToToken(address caller, address to, uint256 amount)
func (_Contracts *ContractsFilterer) FilterSwapTokenToToken(opts *bind.FilterOpts) (*ContractsSwapTokenToTokenIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "SwapTokenToToken")
	if err != nil {
		return nil, err
	}
	return &ContractsSwapTokenToTokenIterator{contract: _Contracts.contract, event: "SwapTokenToToken", logs: logs, sub: sub}, nil
}

// WatchSwapTokenToToken is a free log subscription operation binding the contract event 0x66c9b9a0940362117f399802fc8dee32d47e0f132432c3ccebe00af0206144d5.
//
// Solidity: event SwapTokenToToken(address caller, address to, uint256 amount)
func (_Contracts *ContractsFilterer) WatchSwapTokenToToken(opts *bind.WatchOpts, sink chan<- *ContractsSwapTokenToToken) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "SwapTokenToToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsSwapTokenToToken)
				if err := _Contracts.contract.UnpackLog(event, "SwapTokenToToken", log); err != nil {
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

// ParseSwapTokenToToken is a log parse operation binding the contract event 0x66c9b9a0940362117f399802fc8dee32d47e0f132432c3ccebe00af0206144d5.
//
// Solidity: event SwapTokenToToken(address caller, address to, uint256 amount)
func (_Contracts *ContractsFilterer) ParseSwapTokenToToken(log types.Log) (*ContractsSwapTokenToToken, error) {
	event := new(ContractsSwapTokenToToken)
	if err := _Contracts.contract.UnpackLog(event, "SwapTokenToToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Contracts contract.
type ContractsTransferIterator struct {
	Event *ContractsTransfer // Event containing the contract specifics and raw log

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
func (it *ContractsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsTransfer)
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
		it.Event = new(ContractsTransfer)
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
func (it *ContractsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsTransfer represents a Transfer event raised by the Contracts contract.
type ContractsTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Contracts *ContractsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ContractsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractsTransferIterator{contract: _Contracts.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Contracts *ContractsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ContractsTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsTransfer)
				if err := _Contracts.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Contracts *ContractsFilterer) ParseTransfer(log types.Log) (*ContractsTransfer, error) {
	event := new(ContractsTransfer)
	if err := _Contracts.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}