//function getPrivateKey(){
        var keythereum = require('keythereum');
        var fromkey = keythereum.importFromFile("0x6185d022dc9ff5cb7fcbad78f1f771e03abb4b52", "/home/xialb/.ethereum/net2022");
        //recover输出为buffer类型的私钥
        var privateKey = keythereum.recover('S', fromkey);
        console.log(privateKey.toString('hex'));
	console.log(1);



