const token = artifacts.require('WEMEXToken');
const dummy1 = artifacts.require('dummyToken1');
const dummy2 = artifacts.require('dummyToken2');
const liquidity = artifacts.require('Liquidity');
const _name = 'WEMEXToken';
const _symbol = 'MEX';
const _total_supply = web3.utils.toWei('100000', 'ether');

module.exports = async function (deployer) {

  await deployer.deploy(dummy1, "token1", "tk1", _total_supply)
  await deployer.deploy(dummy2, "token2", "tk2", _total_supply)
  await deployer.deploy(liquidity, dummy1.address, dummy2.address);
  
};
