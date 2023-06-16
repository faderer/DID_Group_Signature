// scripts/index3.js
module.exports = async function main (callback) {
  try {
    const Ver = artifacts.require('Ver');
    const ver = await Ver.deployed();
    const value = await ver.callBn256Add();
    console.log('Verify result is', value);
    const value2 = await ver.callDatacopy(0x1111);
    console.log('Box value is', value2);
    const value3 = await ver.callDatacopy();
    console.log('Box value is', value3);
    callback(0);
  } catch (error) {
    console.error(error);
    callback(1);
  }
};
