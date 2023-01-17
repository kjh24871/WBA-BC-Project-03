const token = artifacts.require('WEMEXToken');
const liquidity = artifacts.require('Liquidity');
const factory = artifacts.require('LiquidityFactory');
const dummy1 = artifacts.require('dummyToken1');
const dummy2 = artifacts.require('dummyToken2');
const _name = 'WEMEXToken';
const _symbol = 'MEX';
const _total_supply = web3.utils.toWei('100000', 'ether');

module.exports = async function (deployer) {
  // deployer.deploy(token, _name, _symbol, _total_supply);
  await deployer.deploy(factory)
  await deployer.deploy(dummy1, "token1", "tk1", _total_supply)
  await deployer.deploy(dummy2, "token2", "tk2", _total_supply)
  await deployer.deploy(liquidity, dummy1.address, dummy2.address);
  
};
