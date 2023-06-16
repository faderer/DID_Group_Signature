// migrations/2_deploy.js
const Test = artifacts.require('Test');
const Verify = artifacts.require('Verify');
const Greeting = artifacts.require("Greeting");
const Time = artifacts.require('Time');

module.exports = async function (deployer) {
  await deployer.deploy(Test);
  await deployer.deploy(Verify);
  await deployer.deploy(Greeting);
  await deployer.deploy(Time);
};
