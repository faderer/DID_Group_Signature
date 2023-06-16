var keythereum = require("keythereum");
var datadir = "/home/xialb/.ethereum/net2022";
var address= "0x6185d022dc9ff5cb7fcbad78f1f771e03abb4b52";
const password = "S";

var keyObject = keythereum.importFromFile(address, datadir);
var privateKey = keythereum.recover(password, keyObject);
console.log(privateKey.toString('hex'));
