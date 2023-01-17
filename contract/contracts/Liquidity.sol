//SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./ILiquidity.sol";




contract Liquidity is ILiquidity{

  IERC20 token;

  //TODO 사용자 address에 liquidity mapping한 값 가지고 있기

  // 풀을 만들고자하는 토큰의 주소를 초기값으로 받습니다.
  constructor (address _token){
    token = IERC20(_token);
  }
  // 유동성 공급을 위한 함수
  function addLiquidity(uint256 _amount) public payable{
    // 해당 함수 호출 시 ERC20토큰의 수량과 이더리움을 함께 받습니다.
    //TODO 사용자 address에 liquidity 기록하기
    token.transferFrom(msg.sender, address(this), _amount);
  }
  // Exchange 컨트랙트가 현재 가지고 있는 이더리움의 개수를 보여주기 위함입니다.
  function getBalance() public view returns(uint256){
    return address(this).balance;
  }

  // CSMM
  // 이더리움을 받고 토큰을 내주는 스왑 함수입니다.
  function ethToERC20CSMMSwap() public payable{
    // msg.value: 사용자가 보낸 이더리움의 개수입니다.
    // CSMM에 의해 1대1 비율로 나(이더를 주고 토큰을 받으려는 사람:해당 함수를 호출한 사람)에게 전송합니다. 
    // inputAmount와 ouputAmount가 같은 이유는 1:1비율을 의미합니다.
    // 사용자가 이더리움 2개를 보내면 그에 맞게 2개의 WEMEX 토큰을 보내주는 의미입니다.
    // 굳이 이렇게 안써도되지만 이해를 돕기위해 작성하였습니다.
    uint256 inputAmount = msg.value;
    uint256 outputAmount = inputAmount;
    token.transfer(msg.sender, outputAmount);
  }
  // x : 넣는 ETH, y : 받은 토큰양
  function getPrice(uint256 x, uint256 y) public pure returns (uint256){
    uint256 numerator = x;
    uint256 denominator = y;
    return numerator / denominator;
  }

////////////////////////////////////////////////////////////////
// CPMM
// _minToken : 슬리피지가 포함된 값, 최소 사용자에게 줘야하는 토큰
  function ethToERC20CPMMSwap(uint256 _minToken) public payable{
    // 사용자게에 받은 이더의 양
    uint256 inputAmount = msg.value;
    // Exchange가 가지고 있는 이더리움의 양
    uint256 x = address(this).balance - inputAmount;
    // Exchange가 가지고 있는 토큰의 양
    uint256 y = token.balanceOf(address(this));
    uint256 ouputAmount = getOuputAmount(msg.value, x, y);
    require(ouputAmount >= _minToken , "lack of amount");
    token.transfer(msg.sender, ouputAmount);
  }

  // CPMM
  // y를 구하기 위함, x는 Exchange가 가지고 있는 이더리움의 양, y는 토큰의 양
  function getOuputAmount(uint256 inputAmount, uint256 x, uint256 y) public pure returns (uint256){
    uint256 numerator = y * inputAmount;
    uint256 denominator = x + inputAmount;
    return numerator/denominator;
  }
}