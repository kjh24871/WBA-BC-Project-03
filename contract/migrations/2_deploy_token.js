const token = artifacts.require('WEMEXToken');
const exchange = artifacts.require('Exchange');
const _name = 'WEMEXToken';
const _symbol = 'MEX';
const _decimals = 18;
const web3 = require('web3');
const _total_supply = web3.utils.toWei('100000', 'ether');

module.exports = function (deployer) {
  // deployer.deploy(token, _name, _symbol, _total_supply);
  deployer.deploy(token, _name, _symbol, _total_supply).then(() => {
    return deployer.deploy(exchange, token.address);
  });
};
