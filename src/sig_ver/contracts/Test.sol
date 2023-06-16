// contracts/Ver.sol
// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.0;
contract Test {
        function callBn256Add(uint256 ax, uint256 ay, uint256 bx, uint256 by) public view returns (bytes32[2] memory result) {
                uint256[4] memory input;
                input[0] = ax;
                input[1] = ay;
                input[2] = bx;
                input[3] = by;
                bool success;
                // solium-disable-next-line security/no-inline-assembly
                assembly {
                        success := staticcall(sub(gas(), 2000), 6, input, 0x80, result, 0x40)
                        // Use "invalid" to make gas estimation work
                        switch success
                        case 0 {
                                revert(0,0)
                        }
                }
                require(success, "pairing-add-failed");
        }
}
