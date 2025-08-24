// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;
/*
✅  用 solidity 实现整数转罗马数字
题目描述在 https://leetcode.cn/problems/roman-to-integer/description/3.
✅  用 solidity 实现罗马数字转数整数
题目描述在 https://leetcode.cn/problems/integer-to-roman/description/
*/

contract RomeNumber{
    // 先把单个字符转成整数
    // 将单个罗马字符转成对应的整数
    function romanCharToInt(bytes1 char) internal pure returns (uint) {
        if (char == "I") return 1;
        if (char == "V") return 5;
        if (char == "X") return 10;
        if (char == "L") return 50;
        if (char == "C") return 100;
        if (char == "D") return 500;
        if (char == "M") return 1000;
        revert("Invalid Roman numeral character");
    }

    function romeToInt(string memory romeNumber)external pure returns(uint){
        bytes memory s = bytes(romeNumber);
        uint total = 0;
        uint prev = 0;

        // 从右往左遍历
        for (int i = int(s.length) - 1; i >= 0; i--) {
            uint curr = romanCharToInt(s[uint(i)]);
            if (curr < prev) {
                total -= curr; // 特殊情况：小数字在大数字左边，做减法
            } else {
                total += curr; // 否则做加法
            }
            prev = curr;
        }
        return total;
    }

    // 再把数字转罗马数字：
    function intToRome(uint num) external pure returns (string memory) {
        require(num > 0 && num < 4000, "Number out of range"); 
        // 定义罗马数字符号及对应的整数值
        uint256[13] memory values = [
            uint256(1000), uint256(900), uint256(500), uint256(400), uint256(100), uint256(90), uint256(50), 
            uint256(40), uint256(10), uint256(9), uint256(5), uint256(4), uint256(1)
        ];

        string[13] memory symbols = [
            "M","CM","D","CD","C","XC","L","XL","X","IX","V","IV","I"
        ];
        string memory result = "";
        
        uint i = 0;
        while (num > 0) {
            while (num >= values[i]) {
                result = string(abi.encodePacked(result, symbols[i]));
                num -= values[i];
            }
            i++;
        }
        return result;
    }
}