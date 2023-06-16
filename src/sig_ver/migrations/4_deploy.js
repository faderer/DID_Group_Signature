// migrations/4_deploy.js
const Ver = artifacts.require('Ver');

module.exports = async function (deployer) {
  await deployer.deploy(Ver);
};
