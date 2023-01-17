const token = artifacts.require('WEMEXToken');
const liquidity = artifacts.require('Liquidity');
const _name = 'WEMEXToken';
const _symbol = 'MEX';
const _total_supply = web3.utils.toWei('100000', 'ether');

module.exports = function (deployer) {
  // deployer.deploy(token, _name, _symbol, _total_supply);
  deployer.deploy(token, _name, _symbol, _total_supply).then(() => {
    return deployer.deploy(liquidity, token.address);
  });
};
