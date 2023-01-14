//SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "./Liquidity.sol";

contract LiquidityFactory{
	//관리자만 컨트랙트에 접근하게 하기 위해 owner를 추가했는데, 오픈제플린에 ownable이란게 있어서 나중에 확인해보겠습니다.
	address public owner;
	mapping (address => Liquidity) public contracts;
	mapping (string => address) poolAddress;

	constructor(){
		owner = msg.sender;
	}

	modifier onlyBy(address _account) {
		require(msg.sender == _account, "Sender not authorized.");
		_;
	}

	function changeOwner(address _newOwner) public onlyBy(owner){
		owner = _newOwner;
	}

	//새로운 풀 추가
	function newLiquidity(address _tokenAddress) public onlyBy(owner) returns (Liquidity){
		Liquidity l = new Liquidity(_tokenAddress);
		// "토큰심볼-ETH"(나중에 WEMIX로 바꾸면 좋겠습니다)이라는 풀이름을 만들어줍니다. (필요할지 모르겠는데 주소 찾을때 이름으로 찾는게 편할것 같아 넣어봤습니다)
		string memory tokenName = ERC20(_tokenAddress).symbol();
		string memory name = string(bytes.concat(bytes(tokenName), "-ETH"));
		contracts[address(l)] = l;
		poolAddress[name] = address(l);
		return l;
	}

	//pool 목록 반환
	// function getLiquidityList() public view returns (mapping(string => address) memory){
	// 	return poolAddress;
	// }
}