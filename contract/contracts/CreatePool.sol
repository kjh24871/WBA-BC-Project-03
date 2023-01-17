//SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "./Liquidity.sol";

contract LiquidityFactory{
	//관리자만 컨트랙트에 접근하게 하기 위해 owner를 추가했는데, 오픈제플린에 ownable이란게 있어서 나중에 확인해보겠습니다.
	address public owner;
	mapping (address => Liquidity) public contracts;
	mapping (string => address) poolAddress;
	string[] poolName;

	//정보 반환횽 구조체
	struct info{
		address poolAddress;
		string poolName;
		uint256 totalLiquidity;
		uint256 amountTokenA;
		uint256 amountTokenB;
	}
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
	function newLiquidity(address _token1Address, address _token2Address) public onlyBy(owner) returns (Liquidity){
		string memory token1Name = ERC20(_token1Address).symbol();
		string memory token2Name = ERC20(_token2Address).symbol();
		string memory name = string(bytes.concat(bytes(token1Name), "-", bytes(token2Name)));
		string memory sym = string(bytes.concat("WEMEX-" , bytes(Strings.toString(poolName.length))));
		Liquidity l = new Liquidity(_token1Address, _token2Address, name, sym);
		poolName.push(name);
		contracts[address(l)] = l;
		poolAddress[name] = address(l);
		return l;
	}

	// pool 목록 반환
	function getLiquidityList() public view returns (info[] memory){
		info[] memory list;
		for (uint i = 0 ; i < poolName.length; i++){
			(uint256 liquidity, uint256 amountTokenA, uint256 amountTokenB) = contracts[poolAddress[poolName[i]]].getLiquidity();
			list[i] = info(
				poolAddress[poolName[i]],
				poolName[i],
				liquidity,
				amountTokenA,
				amountTokenB
				);
		}
		return list;
	}
}