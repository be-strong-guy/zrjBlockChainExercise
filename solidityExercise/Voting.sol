// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Voting {
    /*
    一个mapping来存储候选人的得票数 candidate候选人  votes 票数
    一个vote函数，允许用户投票给某个候选人
    一个getVotes函数，返回某个候选人的得票数
    一个resetVotes函数，重置所有候选人的得票数
    */
    // 记录所有候选人的数组，方便下面清空所有候选人得票
    string[] candidateAll ;
    mapping ( string candidate => uint votes ) public votesMap;
    function vote(string memory candidateName) public {
        // 如果没有这个候选人，就加入数组
        if (!candidateExists(candidateName)){
            candidateAll.push(candidateName);
        }
        votesMap[candidateName] += 1;
    }
    function getVotes(string memory candidateName) public view returns (uint) {
        return votesMap[candidateName];
    }
    // 重置所有候选人的得票数 只能一个一个删除好像，不能重新给它new个map
    function resetVotes() public {
       for (uint i = 0;i<candidateAll.length;i++){
         delete votesMap[candidateAll[i]];
       }
    }
    // 检查候选人是否存在
    function candidateExists(string memory candidate) internal view returns (bool) {
        for (uint i = 0; i < candidateAll.length; i++) {
            if (keccak256(bytes(candidateAll[i])) == keccak256(bytes(candidate))) {
                return true;
            }
        }
        return false;
    }
}
 