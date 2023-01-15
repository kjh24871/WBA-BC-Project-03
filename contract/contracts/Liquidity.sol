//SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "./ILiquidity.sol";

contract Liquidity is ILiquidity, ERC20{

  IERC20 token;

  //TODO 사용자 address에 liquidity mapping한 값 가지고 있기
  event SwapTokenToToken(address caller, address to, uint256 amount);

  // 풀을 만들고자하는 토큰의 주소를 초기값으로 받습니다.
  constructor (address _token) ERC20("lpToken","LP"){
    token = IERC20(_token);
  }
  // [CPMM을 적용한 유동성 공급]
  // 공급 시 유동성 공급자는 Lp토큰을 받는다.(임시)
  // 조건 : 공급되는 코인과 토큰의 양은 1대1 비율로 공급되어야한다. 
  function addLiquidity(uint256 _amount) public payable{
    //  해당 풀이 가지고 있는 전체 공급량
    uint256 totalLiquidity = totalSupply();
    // 기존에 유동성이 존재할 떄
    if (totalLiquidity > 0){
      uint256 ethAmount = address(this).balance - msg.value;
      uint256 tokenAmount = token.balanceOf(address(this));
      // 유동성 공급하려는 토큰의 개수
      uint256 inputTokenAmount = msg.value * tokenAmount / ethAmount;
      require(_amount >= inputTokenAmount);
      token.transferFrom(msg.sender, address(this), inputTokenAmount);
      uint256 liquidityToken = totalLiquidity * msg.value / ethAmount;
      // _mint(msg.sender,liquidityToken); (임시)
    }else{
      // 유동성이 없었을 때
      // 입력받은 토큰의 개수만큼 
      uint tokenAmount = _amount;
      // 가지고 있는 이더리움의 개수
      uint initLiquidity = address(this).balance;
      // 가지고 있던 이더리움의 개수 만큼 토큰을 발행한다.(임시)
      // _mint(msg.sender, initLiquidity);
      token.transferFrom(msg.sender, address(this), tokenAmount);
    }

    // 해당 함수 호출 시 ERC20토큰의 수량과 이더리움을 함께 받습니다.
    //TODO 사용자 address에 liquidity 기록하기

  }
  // Exchange 컨트랙트가 현재 가지고 있는 이더리움의 개수를 보여주기 위함입니다.
  function getBalance() public view returns(uint256){
    return address(this).balance;
  }

  // _minToken : 슬리피지가 포함된 값, 최소 사용자에게 줘야하는 토큰, 사용자가 입력한 슬리피지 값
  function swapCoinToToken(uint256 _minToken) public payable{
    // 사용자게에 받은 이더의 양 
    uint256 inputAmount = msg.value;
    // Exchange가 가지고 있는 이더리움의 양
    uint256 x = address(this).balance - inputAmount;
    // Exchange가 가지고 있는 토큰의 양
    uint256 y = token.balanceOf(address(this));
    uint256 outputAmount = getSwapRatio(msg.value, x, y);
    // 사용자가 입력한 최소의 슬리피지값 이상은 되야 교환가능하다.
    require(outputAmount >= _minToken , "lack of amount");
    // 사용자에게 보낸다.
    IERC20(token).transfer(msg.sender, outputAmount);
  }

  // _tokenAmount : 몇개의 토큰을 바꿀껀지
  // _minCoin : 슬리피지가 입력된 값
  function swapTokenToCoin(uint256 _tokenAmount, uint256 _minCoin) public payable{
    // Liquidity 컨트랙트가 가지고 있는 토큰의 양, 코인의 양 
    uint256 outputAmount = getSwapRatio(_tokenAmount, token.balanceOf(address(this)), address(this).balance);
    // 입력받은 슬리피지보다 많이 받아야한다.
    require(outputAmount >= _minCoin , "lack of amount");
    // 사용자가 토큰을 Liquidity(컨트랙트)에게 보내게한다.
    emit SwapTokenToToken(msg.sender, address(this), _tokenAmount);
    IERC20(token).transferFrom(msg.sender, address(this),_tokenAmount);
    // 이더를 보낸다. 함수를 호출한 사용자에게
    // payable(msg.sender) : https://ethereum.stackexchange.com/questions/113243/payablemsg-sender
    payable(msg.sender).transfer(outputAmount);
  }

  // CPMM
  // inputAmount : 사용자에게 받은 값,  1
  // x : liquidity의 input,  500 + 1 = 501
  // y : liquidity output, 1000
  function getSwapRatio(uint256 inputAmount, uint256 x, uint256 y) public pure returns (uint256){
    uint256 numerator = y * inputAmount; // 1000 x 1
    uint256 denominator = x + inputAmount; // 501 + 1
    return numerator/denominator;
  }
}