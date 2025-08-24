// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
/*
题目描述：反转一个字符串。输入 "abcde"，输出 "edcba"
*/
contract ReverseString{
    // 定义一个函数，输入一个字符串，输出反转后的字符串
    function reverseString(string memory input) public pure returns (string memory) {
        bytes memory bytesInput = bytes(input);
        uint256 length = bytesInput.length;
        bytes memory reversedBytes = new bytes(length);
        for (uint256 i = 0; i < length; i++) {
            reversedBytes[i] = bytesInput[length - 1 - i];
        }
        return string(reversedBytes);
    }
}