const token = artifacts.require('WEMEXToken');
const liquidity = artifacts.require('Liquidity');
const factory = artifacts.require('LiquidityFactory');
const _name = 'WEMEXToken';
const _symbol = 'MEX';
const _decimals = 18;
const web3 = require('web3');
const _total_supply = web3.utils.toWei('100000', 'ether');

module.exports = async function (deployer) {
  // deployer.deploy(token, _name, _symbol, _total_supply);
  await deployer.deploy(token, _name, _symbol, _total_supply);
  await deployer.deploy(liquidity, token.address);
  await deployer.deploy(factory)

};
