const Liquidity = artifacts.require('Liquidity');
const { expect } = require('chai');
const { ethers, providers } = require('ethers');

const web3 = require('web3');
const { toWei, toBN, fromWei } = web3.utils;
const dummy1 = artifacts.require('dummyToken1');
const dummy2 = artifacts.require('dummyToken2');

contract('Liquidity', (accounts) => {
  let dum1;
  let dum2;
  let liquidity;
  // contract() 실행 후 & 내부의 테스트 함수들이 실행되기 전에 실행되는 부분
  before(async () => {
    console.log(accounts);
    dum1 = await dummy1.deployed();
    dum2 = await dummy2.deployed();
    // console.log('Token Contract Address:', dummy1.address);
    liquidity = await Liquidity.deployed();
    console.log('exchangeInstance Address', liquidity.address);
  });

  describe.skip('Token 생성자', async () => {
    it('Token은 TotalSupply로 10000WEMEX토큰을 가지고 있어야한다.', async () => {
      let totalSupply = await token.totalSupply();
      let temp = web3.utils.toBN(totalSupply).toString();
      // console.log('totalSupply', web3.utils.toBN(totalSupply).toString());
      // assert.equal(web3.utils.toBN(totalSupply).toString()), web3.utils.toWei('100', 'ether'));
      // console.log('totalSupply: ', web3.utils.toBN(totalSupply).toString());

      // token은 처음에 100를 가지고 있다.
      expect(web3.utils.toBN(totalSupply).toString()).to.equal(
        web3.utils.toWei('100000', 'ether')
      );
    });
  });
  describe('유동성 공급 (유동성 0일때)', async () => {
    it('유동성 공급한다.', async () => {
      // liquidity 컨트랙트에게 10개의 WEMEX토큰을 사용할 수 있도록 승인
      await dum1.approve(liquidity.address, web3.utils.toWei('10', 'ether'));
      await dum2.approve(liquidity.address, web3.utils.toWei('20', 'ether'));
      // token.transferFrom(msg.sender, address(this), _amount);

      // 유동성 공급자는 10개의 dummy1 토큰과 20개의 dummy2 토큰을 Liquidity 컨트랙트에 유동성 공급한다.
      await liquidity.addLiquidity(web3.utils.toWei('10', 'ether'), web3.utils.toWei('20', 'ether'));
      // Liquidity 컨트랙트의 잔고는 2이더가 있어야한다.
      let totalSupply = await liquidity.totalSupply();
      console.log('total supply', web3.utils.toBN(totalSupply).toString())
      expect(web3.utils.toBN(totalSupply).toString()).to.equal(
        web3.utils.toWei('20', 'ether')
      );

      //  Liquidity 컨트랙트에 들어간 토큰 개수는 10개여야한다.
      var balance = await dum1.balanceOf(liquidity.address);
      console.log('balance', web3.utils.toBN(balance).toString());
      expect(web3.utils.toBN(balance).toString()).to.equal(
        web3.utils.toWei('10', 'ether')
      );
      //  Liquidity 컨트랙트에 들어간 토큰 개수는 20개여야한다.
      var balance = await dum2.balanceOf(liquidity.address);
      console.log('balance', web3.utils.toBN(balance).toString());
      expect(web3.utils.toBN(balance).toString()).to.equal(
        web3.utils.toWei('20', 'ether')
      );    
    });
  });

  describe('유동성 공급 (유동성 있을 때)', async () => {
    it('유동성 공급한다.', async () => {
      // liquidity 컨트랙트에게 10개의 WEMEX토큰을 사용할 수 있도록 승인
      await dum1.approve(liquidity.address, web3.utils.toWei('20', 'ether'));
      await dum2.approve(liquidity.address, web3.utils.toWei('20', 'ether'));
      // token.transferFrom(msg.sender, address(this), _amount);

      // 유동성 공급자는 20개의 dummy1 토큰과 20개의 dummy2 토큰을 Liquidity 컨트랙트에 유동성 공급사도한다.
      await liquidity.addLiquidity(web3.utils.toWei('20', 'ether'), web3.utils.toWei('20', 'ether'));
      
      // 현재 풀에 1:2의 비유로 토큰이 들어있으므로 20 / 20의 토큰을 입력 시도하면 10 / 20만 입력되고 유동성은 20이 공급되어야 한다.


      let totalSupply = await liquidity.totalSupply();
      console.log('total supply', web3.utils.toBN(totalSupply).toString())
      expect(web3.utils.toBN(totalSupply).toString()).to.equal(
        web3.utils.toWei('40', 'ether')
      );

      //  Liquidity 컨트랙트에 들어간 토큰 개수는 10개여야한다.
      var balance = await dum1.balanceOf(liquidity.address);
      console.log('balance', web3.utils.toBN(balance).toString());
      expect(web3.utils.toBN(balance).toString()).to.equal(
        web3.utils.toWei('20', 'ether')
      );
      //  Liquidity 컨트랙트에 들어간 토큰 개수는 20개여야한다.
      var balance = await dum2.balanceOf(liquidity.address);
      console.log('balance', web3.utils.toBN(balance).toString());
      expect(web3.utils.toBN(balance).toString()).to.equal(
        web3.utils.toWei('40', 'ether')
      );    
    });
  });

  describe('스왑 ', async () => {
    it('사용자가 토근 주소와 입력 토큰 양을 입력하면 교환해준다.', async () => {
      // liquidity 컨트랙트에게 10개의 dummy2토큰을 사용할 수 있도록 승인
      await dum2.approve(liquidity.address, web3.utils.toWei('10', 'ether'));
      await liquidity.swap(dum2.address, web3.utils.toWei('10', 'ether'));

      //  uint256 numerator = _liquidityOutput * _inputAmount;
      //  uint256 denominator = _liquidityInput + _inputAmount;
      // 20 * 10
      // 40 + 10
      // ratio = 4

      console.log(liquidity.address)
      let balance = await dum1.balanceOf(liquidity.adddress);
      expect(web3.utils.toBN(balance).toString()).to.equal(
        web3.utils.toWei('16', 'ether')
      );
     
      balance = await dum2.balanceOf(liquidity.address);
      console.log('balance', web3.utils.toBN(balance).toString());
      expect(web3.utils.toBN(balance).toString()).to.equal(
        web3.utils.toWei('50', 'ether')
      );
    });
  });
  describe.skip('CSMM', async () => {
    it('가격은 일정해야한다.', async () => {
      // liquidity 컨트랙트에게 100개의 WEMEX토큰을 사용할 수 있도록 승인
      // await token.approve(liquidity.address, web3.utils.toWei('100', 'ether'));
      // 유동성 공급자는 70개의 WEMEX 토큰과 10개의 이더리움을 Exchange 컨트랙트에 유동성 공급한다.
      let exchangeBalance = await liquidity.getBalance();
      var balance = await token.balanceOf(liquidity.address);
      let price = await liquidity.getPrice(exchangeBalance, balance);
      console.log(web3.utils.toBN(exchangeBalance).toString());
      console.log(web3.utils.toBN(balance).toString());
      // expect(web3.utils.toBN(totalSupply).toString()).to.equal(
      //   web3.utils.toWei('100', 'ether')
      // );
    });
  });
  describe.skip('CPMM', async () => {
    it('단순 슬리피지 계산', async () => {
      await token.approve(liquidity.address, web3.utils.toWei('50', 'ether'));
      // 유동성 공급자가 40개의 토큰과 10개의 이더리움을 추가로 공급했다. (총 27개 이더 보유/115개 토큰 보유)
      await liquidity.addLiquidity(web3.utils.toWei('40', 'ether'), {
        value: web3.utils.toWei('10', 'ether'),
      });
      // 사용자는 1이더를 넣었을 때 몇개의 토큰을 받게될까
      let exchangeBalance = await liquidity.getBalance();

      // 115000000000000000000 토큰
      // 27000000000000000000 이더리움
      // 사용자는 1이더 넣으면 4.25925925926 기대함
      // y 115 토큰 / x 27 이더 = 4.25....토큰을 기대한다.

      // 하지만 실제로는 4.107142857142857142 토큰을 얻게된다.
      // x 27이더 + 1이더 = 28
      // y 115토큰
      // 받게될 토큰 = 27(기존 이더) + 1(사용자에게 받을 이더) / 115(기존 토큰)
      // 받게될 토큰 = 4.107142857142857142
      console.log(
        '몇개의 이더리움? =>',
        web3.utils.toBN(exchangeBalance).toString()
      );
      // 21개 이더리움

      let exchangeTokenBalance = await token.balanceOf(liquidity.address);
      console.log(
        '거래소에는 몇개의 토큰? ->',
        web3.utils.toBN(exchangeTokenBalance).toString()
      );

      let temp = await liquidity.getOuputAmount(
        web3.utils.toWei('1', 'ether'),
        web3.utils.toBN(exchangeBalance).toString(),
        web3.utils.toBN(exchangeTokenBalance).toString()
      );
      // 1이더를 넣었을 때 받게되는 토큰의 양
      console.log(web3.utils.fromWei(web3.utils.toBN(temp).toString()));
      expect(web3.utils.fromWei(web3.utils.toBN(temp).toString())).to.equal(
        '4.107142857142857142'
      );

      let temp2 = await liquidity.getOuputAmount(
        web3.utils.toWei('115', 'ether'),
        web3.utils.toBN(exchangeBalance).toString(),
        web3.utils.toBN(exchangeTokenBalance).toString()
      );
      // 1이더를 넣었을 때 받게되는 토큰의 양
      console.log(web3.utils.fromWei(web3.utils.toBN(temp2).toString()));
    });
  });
  describe.skip('swap coin to token', async () => {
    it('사용자는 이더리움을 넣고 CPMM 알고리즘에 의해 토큰을 받을 수 있다.', async () => {
      await token.approve(liquidity.address, toWei('1000', 'ether'));

      // 유동성 공급 => 토큰 1000 : 이더 200
      await liquidity.addLiquidity(toWei('1000', 'ether'), {
        from: accounts[0],
        value: toWei('200', 'ether'),
      });

      // 현재 유동성: 토큰 1000개, 이더 200개
      // 사용자가 1개의 토큰을 넣는다. => 사용자는 1개를 넣으면 토큰 5개를 받을 것으로 예상
      // 실제 받은 토큰은 497.xx개의 토큰을 받음

      // 받을 토큰 y는 = 1000MEX * 1 / 200ETH + 1

      let liquidityTokenBalance = await token.balanceOf(liquidity.address);
      let liquidityBalance = await liquidity.getBalance();
      console.log(
        'liquidityTokenBalance : ',
        toBN(liquidityTokenBalance).toString()
      );
      console.log('liquidityBalance : ', toBN(liquidityBalance).toString());
      let tempRatio = await liquidity.getSwapRatio(
        web3.utils.toWei('1', 'ether'),
        liquidityBalance,
        liquidityTokenBalance
      );

      console.log('슬리피지 적용 시 : ', toBN(tempRatio).toString());

      await liquidity.swapCoinToToken(toWei('1', 'ether'), {
        from: accounts[1],
        value: web3.utils.toWei('1', 'ether'),
      });

      let accountTokenBalance = await token.balanceOf(accounts[1]);
      console.log(
        'address가 받은 Token : ',
        toBN(accountTokenBalance).toString()
      );
    });
  });
  describe.skip('swap token to coin', async () => {
    it('사용자는 토큰을 넣고 CPMM 알고리즘에 의해 위믹스 코인을 받을 수 있다.', async () => {
      await token.transfer(accounts[1], toWei('5000', 'ether'));
      let sender = await token.balanceOf(accounts[1]);
      console.log('accounts[1]: ', toBN(sender).toString());

      await token.approve(liquidity.address, toWei('1000', 'ether'));
      // 유동성 공급 => 토큰 1000 : 이더 200
      await liquidity.addLiquidity(toWei('200', 'ether'), {
        from: accounts[0],
        value: toWei('200', 'ether'),
      });

      // 현재 유동성 : 토큰 1000개, 이더 200개
      // 사용자가 1개의 토큰을 넣는다. => 사용자는 토큰 1개를 넣으면 위믹스 코인 0.2개를 받을 것으로 예상
      // 실제 받은 토큰은 497.xx개의 토큰을 받음

      // 받을 토큰 y는 = 200WEMIX * 1 / 200Token + 1

      let liquidityTokenBalance = await token.balanceOf(liquidity.address);
      let liquidityBalance = await liquidity.getBalance();
      console.log(
        'liquidityTokenBalance : ',
        toBN(liquidityTokenBalance).toString()
      );
      console.log('liquidityBalance : ', toBN(liquidityBalance).toString());
      let tempRatio = await liquidity.getSwapRatio(
        toWei('1', 'ether'),
        liquidityTokenBalance,
        liquidityBalance
      );

      console.log('슬리피지 적용 시 : ', toBN(tempRatio).toString());
      console.log(accounts[1]);
      sender = await token.balanceOf(accounts[1]);

      await token.approve(liquidity.address, toWei('1000', 'ether'), {
        from: accounts[1],
      });
      await liquidity.swapTokenToCoin(
        toWei('2', 'ether'),
        toWei('1', 'ether'),
        { from: accounts[1] }
      );

      // let accountTokenBalance = await token.balanceOf(accounts[1]);
      // console.log(
      //   'address가 받은 Token : ',
      //   toBN(accountTokenBalance).toString()
      // );
    });
  });
});
