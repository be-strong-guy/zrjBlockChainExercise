// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/*
✅  合并两个有序数组 (Merge Sorted Array)
题目描述：将两个有序数组合并为一个有序数组。
✅  二分查找 (Binary Search)
题目描述：在一个有序数组中查找目标值。
*/
contract MergeAndBinarySearch{

    function mergeSortedArray(uint256[] memory array1,uint256[] memory array2)external pure returns  (uint256[] memory){
        require(array1.length > 0 && array2.length > 0, "Array must not be empty"); 
        uint len1 = array1.length;
        uint len2 = array2.length;
        uint256[] memory result = new uint256[](len1 + len2);

        uint i = 0; // array1 索引
        uint j = 0; // array2 索引
        uint k = 0; // result 索引

        // 双指针合并
        while (i < len1 && j < len2) {
            if (array1[i] <= array2[j]) {
                result[k] = array1[i];
                i++;
            } else {
                result[k] = array2[j];
                j++;
            }
            k++;
        }

        // 复制剩余元素
        while (i < len1) {
            result[k] = array1[i];
            i++;
            k++;
        }

        while (j < len2) {
            result[k] = array2[j];
            j++;
            k++;
        }
        return result ;
    }
    // 二分查找 (Binary Search)题目描述：在一个有序数组中查找目标值。
    function binarySearch(uint[] memory array, uint target)external pure returns (int ){
        require(array.length > 0, "Array must not be empty"); 
        int left = 0;
        int right = int(array.length - 1);
        while (left <= right) {
            //防止mid数组越界，就这么写
           int mid = left + (right - left) / 2;
            if (array[uint(mid)] == target) {
                return mid; // 返回索引
            } else if (array[uint(mid)] < target) {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }
        // 没有找到
        return -1;
    }

}