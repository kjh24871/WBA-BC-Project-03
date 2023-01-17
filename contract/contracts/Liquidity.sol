//SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../node_modules/@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract Liquidity is ERC20{
  IERC20 tokenA;
  IERC20 tokenB;

  /**
    실제 가격 비율: 1:20
    풀 기준 가격 : amountB:amountA
  */
  event Swap(address caller, address input, address output, uint256 inputAmount, uint256 outputAmount);
  event AddLiquidity(address caller, uint256 amount);
  constructor (address _tokenA, address _tokenB) ERC20("lpToken", "LP"){
    tokenA = IERC20(_tokenA);
    tokenB = IERC20(_tokenB);
  }

  function addLiquidity(uint256 _desiredAmountA, uint256 _desiredAmountB) public
  returns(uint256 actualAmountA, uint256 actualAmountB){
    if (totalSupply() == 0){
      actualAmountA = _desiredAmountA;
      actualAmountB = _desiredAmountB;
    } else{
      uint256 amountA = tokenA.balanceOf(address(this));
      uint256 amountB = tokenB.balanceOf(address(this));
      
      actualAmountA = _desiredAmountB * amountA / amountB;
      if (actualAmountA <= _desiredAmountA){
        actualAmountB = _desiredAmountB;
      } else{
        actualAmountB = _desiredAmountA * amountB / amountA;
        actualAmountA = _desiredAmountA;
      }
      
    }
    tokenA.transferFrom(msg.sender, address(this), actualAmountA);
    tokenB.transferFrom(msg.sender, address(this), actualAmountB);
    _mint(msg.sender, actualAmountB);
    emit AddLiquidity(msg.sender, actualAmountB);
    return (actualAmountA, actualAmountB);
  }

  function removeLiquidity(uint256 _amount) public returns(uint256 outputA, uint256 outputB){
    uint256 balance = balanceOf(msg.sender);
    require(_amount >= balance, "balance is not enough");
    uint256 totalSupply = totalSupply();
    outputA = tokenA.balanceOf(address(this)) * _amount / totalSupply;
    outputA = tokenA.balanceOf(address(this)) * _amount / totalSupply;
    
    tokenA.transferFrom(address(this), msg.sender, outputA);
    tokenB.transferFrom(address(this), msg.sender, outputB);
    _burn(msg.sender, _amount);
    return (outputA, outputB);
  }
  
  function swap(address _inputToken, uint256 _inputAmount) public returns(uint256 outputAmount){
    IERC20 inputToken = IERC20(_inputToken);
    IERC20 outputToken = (tokenA == inputToken) ? tokenB : tokenA;
    outputAmount = getSwapRatio(_inputAmount, inputToken.balanceOf(address(this)), outputToken.balanceOf(address(this)));

    inputToken.transferFrom(msg.sender, address(this), _inputAmount);
    outputToken.transferFrom(address(this), msg.sender, _inputAmount);
    emit Swap(msg.sender, address(inputToken), address(outputToken), _inputAmount, outputAmount);
    return outputAmount;
}

  function getSwapRatio(uint256 _inputAmount, uint256 _liquidityInput, uint256 _liquidityOutput)
   public pure returns(uint256){
    uint256 numerator = _liquidityOutput * _inputAmount;
    uint256 denominator = _liquidityInput + _inputAmount;
    return numerator/denominator;
  }
}