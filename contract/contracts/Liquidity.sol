//SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../node_modules/@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "../node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract Liquidity is ERC20{
  IERC20 tokenA;
  IERC20 tokenB;
  uint256 fee = 0;
  address payable[] delegators;

  event Swap(address caller, address input, address output, uint256 inputAmount, uint256 outputAmount);
  event AddLiquidity(address caller, uint256 amount);
  event RemoveLiquidity(address caller, uint256 amount);
  constructor (
    address _tokenA, 
    address _tokenB, 
    string memory name, 
    string memory sym
    ) ERC20(name, sym){
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
    if (!isExists(msg.sender)){
      delegators.push(payable(msg.sender));
    }
    emit AddLiquidity(msg.sender, actualAmountB);
    return (actualAmountA, actualAmountB);
  }

  function removeLiquidity(uint256 _amount) public returns(uint256 outputA, uint256 outputB){
    uint256 balance = balanceOf(msg.sender);
    require(_amount <= balance, "balance is not enough");
    uint256 totalSupply = totalSupply();
    outputA = tokenA.balanceOf(address(this)) * _amount / totalSupply;
    outputB = tokenB.balanceOf(address(this)) * _amount / totalSupply;
  
    tokenA.transfer(msg.sender, outputA);
    tokenB.transfer(msg.sender, outputB);
    _burn(msg.sender, _amount);
    emit RemoveLiquidity(msg.sender, _amount);
    return (outputA, outputB);
  }
  
  function swap(address _inputToken, uint256 _inputAmount) public payable returns(uint256 outputAmount){
    require(msg.value >= fee, "not enough fee");
    IERC20 inputToken = IERC20(_inputToken);
    IERC20 outputToken = (tokenA == inputToken) ? tokenB : tokenA;
    outputAmount = getSwapRatio(_inputAmount, inputToken.balanceOf(address(this)), outputToken.balanceOf(address(this)));

    inputToken.transferFrom(msg.sender, address(this), _inputAmount);
    outputToken.transfer(msg.sender, outputAmount);
    emit Swap(msg.sender, address(inputToken), address(outputToken), _inputAmount, outputAmount);
    return outputAmount;
  }

  function pullRewardWithServer() public{
    uint256 amount;
    uint256 totalSupply = totalSupply();
    for(uint256 i = 0 ; i < delegators.length ; i++){
      amount = address(this).balance * balanceOf(delegators[i]) / totalSupply;
      delegators[i].transfer(amount);
    }
  }

  function pullRewardWithClient() public{
    uint256 myRatio;
    uint256 amount;
    uint256 totalSupply = totalSupply();
    myRatio = msg.sender.balance / totalSupply;
    amount = address(this).balance * myRatio;
    
    // Todo : PullReward 방식 수정
    msg.sender.transfer(amount);
  }

  function getSwapRatio(uint256 _inputAmount, uint256 _liquidityInput, uint256 _liquidityOutput)
   public pure returns(uint256){
    uint256 numerator = _liquidityOutput * _inputAmount;
    uint256 denominator = _liquidityInput + _inputAmount;
    return numerator/denominator;
  }

  function getTotalFee() public view returns(uint256){
    return address(this).balance;
  }

  function isExists(address user) private view returns (bool) {
    for (uint i = 0; i < delegators.length; i++) {
        if (delegators[i] == user) {
            return true;
        }
    }
    return false;
  }
  function getLiquidity() public view returns (uint256 liquidity, uint256 amountA, uint256 amountB){
    liquidity = totalSupply();
    amountA = tokenA.balanceOf(address(this));
    amountB = tokenB.balanceOf(address(this));
    return (liquidity, amountA, amountB);
  }

  function getTokenAddress() public view returns (address, address){
    return (address(tokenA), address(tokenB));
  }
}