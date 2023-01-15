const LiquidityFactory = artifacts.require('LiquidityFactory');
const { expect } = require('chai');
const { ethers, providers } = require('ethers');
const web3 = require('web3');
const { toWei, toBN, fromWei } = web3.utils;
const Token = artifacts.require('WEMEXToken');

contract('LiquidityFactory', (accounts) => {
  let token;
  let liquidityFactory;
  // contract() 실행 후 & 내부의 테스트 함수들이 실행되기 전에 실행되는 부분
  before(async () => {
	console.log(accounts);
    token = await Token.deployed();
    console.log('Token Contract Address:', token.address);
    liquidityFactory = await LiquidityFactory.deployed();
    console.log('factory Address', liquidityFactory.address);
  });

  describe('newLiquidity', async function() {
    it('newLiquidity는 새로운 토큰주소를 입력 받는다.', async function() {
      let tokenAddress = await token.address;
      // token은 처음에 100를 가지고 있다.
      liquidityFactory.newLiquidity(tokenAddress)
    });
  });
  describe('getLiquidityList', async function() {
    it('getLiquidityList는 풀 리스트를 반환한다', async function(){
      liquidityFactory.getLiquidityList();
    })
  })

  
});
