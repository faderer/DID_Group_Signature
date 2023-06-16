// contracts/Time.sol
// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.0;
contract Time {
    bytes public g1;
    bytes public h;
    bytes public u;
    bytes public v;
    bytes public g2;
    bytes public w;
    bytes public ehw;
    bytes public ehg2;
    bytes public minusEg1g2;

    function newGroup(bytes memory _g1, bytes memory _h, bytes memory _u, bytes memory _v, bytes memory _g2, 
    bytes memory _w, bytes memory _ehw, bytes memory _ehg2, bytes memory _minusEg1g2) public {
        g1 = _g1;
        h = _h;
        u = _u;
        v = _v;
        g2 = _g2;
        w = _w;
        ehw = _ehw;
        ehg2 = _ehg2;
        minusEg1g2 = _minusEg1g2;
    }

    function retrieve() public view returns(bytes memory){
        bytes memory output;
//        output = abi.encodePacked(g1, h, u, v, g2, w, ehw, ehg2);
        output = bytes.concat(g1, " ", h);
        output = bytes.concat(output, " ", u);
        output = bytes.concat(output, " ", v);
        output = bytes.concat(output, " ", g2);
        output = bytes.concat(output, " ", w);
        output = bytes.concat(output, " ", ehw);
        output = bytes.concat(output, " ", ehg2);
	return output;
    }

    function verify(bytes memory combine) public view returns (bytes32) {
        bytes memory input;
        bytes32 output;
        // input = bytes.concat(g1, " ", h);
        // input = bytes.concat(input, " ", u);
        // input = bytes.concat(input, " ", v);
        // input = bytes.concat(input, " ", g2);
        // input = bytes.concat(input, " ", w);
        // input = bytes.concat(input, " ", ehw);
        // input = bytes.concat(input, " ", ehg2);
        // input = bytes.concat(input, " ", combine);
        input = abi.encodePacked(g1, h, u, v, g2, w, ehw, ehg2, combine);
	uint256 len = input.length;
	for (uint i = 0; i < 100; i++){
        assembly {
            let out := mload(0x40)
            let success:=staticcall(gas(), 10, add(input,0x20), len, out, 0x20)
            switch success
            case 0 {
                revert(0,0)
            } default {
                output := mload(out)
            }
        }
	}
        return output;
    }
}
