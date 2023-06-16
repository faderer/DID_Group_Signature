// contracts/Ver.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Ver{
    function callBn256Add() public view returns (bytes32 result) {
        bytes32 input = hex"0000000000000000000000000000000000000000000000000000000000000002";
	//result = hex"0000000000000000000000000000000000000000000000000000000000000001";
	//bytes32[26] memory input;
	//bytes32 temp = 0x111122223333444455556666777788889999AAAABBBBCCCCDDDDEEEEFFFFCCCC;	
	//uint32 i = 0;
        //for (i = 0; i < 26; i++){
	//	input[i] = temp;
	//}
	bool success;
        assembly {
            success := staticcall(gas(), 0x02, input, 0x20, result, 0x20)
            switch success
                case 0 {
                    revert(0,0)
                }
        }
	require(success, "pairing-add-failed");
    }
function callDatacopy(bytes memory data) public view returns (bytes memory) {
    bytes memory ret = new bytes(data.length);
    assembly {
        let len := mload(data)
        if iszero(staticcall(gas(), 10, add(data, 0x20), len, add(ret,0x20), len)) {
            invalid()
        }
    }

    return ret;
}
function callDatacopy() public returns (bytes memory) {
        bytes memory input = hex"2222";
        bytes memory ret = new bytes(input.length);
        assembly {
                let len := mload(input)
                if iszero(call(gas(), 10, 0, add(input, 0x20), len, add(ret,0x20), len)) {
                invalid()
                }
        }
        return ret;
        }    
        function callBlake2F() public view returns (bytes32[2] memory) {
                bytes32[2] memory output;
		
		uint32 rounds = 12;

  bytes32[2] memory h;
  h[0] = hex"48c9bdf267e6096a3ba7ca8485ae67bb2bf894fe72f36e3cf1361d5f3af54fa5";
  h[1] = hex"d182e6ad7f520e511f6c3e2b8c68059b6bbd41fbabd9831f79217e1319cde05b";

  bytes32[4] memory m;
  m[0] = hex"6162630000000000000000000000000000000000000000000000000000000000";
  m[1] = hex"0000000000000000000000000000000000000000000000000000000000000000";
  m[2] = hex"0000000000000000000000000000000000000000000000000000000000000000";
  m[3] = hex"0000000000000000000000000000000000000000000000000000000000000000";

  bytes8[2] memory t;
  t[0] = hex"03000000";
  t[1] = hex"00000000";

  bool f = true;

                bytes memory args = abi.encodePacked(rounds, h[0], h[1], m[0], m[1], m[2], m[3], t[0], t[1], f);

                assembly {
                        if iszero(staticcall(not(0), 0x09, add(args, 32), 0xd5, output, 0x40)) {
                        revert(0, 0)
                        }
                }

                return output;
        }
}

