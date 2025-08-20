// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;
/*
编写一个讨饭合约
任务目标
使用 Solidity 编写一个合约，允许用户向合约地址发送以太币。
记录每个捐赠者的地址和捐赠金额。
允许合约所有者提取所有捐赠的资金。

任务步骤
编写合约
创建一个名为 BeggingContract 的合约。
合约应包含以下功能：
一个 mapping 来记录每个捐赠者的捐赠金额。
一个 donate 函数，允许用户向合约发送以太币，并记录捐赠信息。
一个 withdraw 函数，允许合约所有者提取所有资金。
一个 getDonation 函数，允许查询某个地址的捐赠金额。
使用 payable 修饰符和 address.transfer 实现支付和提款。
部署合约
在 Remix IDE 中编译合约。
部署合约到 Goerli 或 Sepolia 测试网。
测试合约
使用 MetaMask 向合约发送以太币，测试 donate 功能。
调用 withdraw 函数，测试合约所有者是否可以提取资金。
调用 getDonation 函数，查询某个地址的捐赠金额。

任务要求
合约代码：
使用 mapping 记录捐赠者的地址和金额。
使用 payable 修饰符实现 donate 和 withdraw 函数。
使用 onlyOwner 修饰符限制 withdraw 函数只能由合约所有者调用。
测试网部署：
合约必须部署到 Goerli 或 Sepolia 测试网。
功能测试：
确保 donate、withdraw 和 getDonation 函数正常工作。

提交内容
合约代码：提交 Solidity 合约文件（如 BeggingContract.sol）。
合约地址：提交部署到测试网的合约地址。
测试截图：提交在 Remix 或 Etherscan 上测试合约的截图。
*/
contract BeggingContract {
    // 合约所有者（部署者）
    address payable public owner;

    // 记录每位捐赠者的累计捐款额
    mapping(address => uint256) private donations;

    // 事件：收到捐赠 & 提款
    event DonationReceived(address indexed from, uint256 amount);
    event Withdrawn(address indexed to, uint256 amount);

    // 仅所有者可调用
    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner");
        _;
    }

    constructor() {
        owner = payable(msg.sender);
    }

    /// 主动捐赠函数
    function donate() external payable {
        require(msg.value > 0, "No ETH sent");
        donations[msg.sender] += msg.value;
        emit DonationReceived(msg.sender, msg.value);
    }

    /// 允许直接向合约地址转账（例如在钱包里直接“发送”）
    /// 也会自动记录到 donations
    receive() external payable {
        require(msg.value > 0, "No ETH sent");
        donations[msg.sender] += msg.value;
        emit DonationReceived(msg.sender, msg.value);
    }

    /// 兜底函数（如误带数据调用），同样接收并记录
    fallback() external payable {
        require(msg.value > 0, "No ETH sent");
        donations[msg.sender] += msg.value;
        emit DonationReceived(msg.sender, msg.value);
    }

    ///  查询某地址累计捐赠金额
    function getDonation(address donor) external view returns (uint256) {
        return donations[donor];
    }

    /// 所有者提走合约中全部余额
    function withdraw() external onlyOwner {
        uint256 amount = address(this).balance;
        require(amount > 0, "Nothing to withdraw");
        owner.transfer(amount);
        emit Withdrawn(owner, amount);
    }
}
