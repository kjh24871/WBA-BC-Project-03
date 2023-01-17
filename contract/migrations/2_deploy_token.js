const token = artifacts.require('WEMEXToken');
const liquidity = artifacts.require('Liquidity');
const _name = 'WEMEXToken';
const _symbol = 'MEX';
const _total_supply = web3.utils.toWei('100000', 'ether');

module.exports = async function (deployer) {
  const token1 = token
  const token2 = token
  await deployer.deploy(token1, "token1", "tk1", _total_supply)
  await deployer.deploy(token2, "token2", "tk2", _total_supply)
  await deployer.deploy(liquidity, token1.address, token2.address);
  
};
