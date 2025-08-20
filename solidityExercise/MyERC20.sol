// SPDX-License-Identifier: MIT
pragma solidity ^0.8;
/*
任务：参考 openzeppelin-contracts/contracts/token/ERC20/IERC20.sol实现一个简单的 ERC20 代币合约。要求：
合约包含以下标准 ERC20 功能：
balanceOf：查询账户余额。
transfer：转账。
approve 和 transferFrom：授权和代扣转账。
使用 event 记录转账和授权操作。
提供 mint 函数，允许合约所有者增发代币。
提示：
使用 mapping 存储账户余额和授权信息。
使用 event 定义 Transfer 和 Approval 事件。
部署到sepolia 测试网，导入到自己的钱包
 
 任务完成已经部署到sepolia测试网，不过需要flattened展开，这是测试网部署的结果：
 https://sepolia.etherscan.io/address/0xdBd0Dc72CBA5eC79717eBce16c63172B4A22748E#code
*/
import {IERC20} from "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/v4.9.0/contracts/token/ERC20/IERC20.sol";
contract MyERC20 is IERC20{
    //1、定义发行代币（通证）的相关信息，比如name，简称，额度，精度等，当然这些ERC2.0里面默认的就有的
    string public name = "ZRJToken";
    string public symbol = "ZRJT";
    uint8 public decimals = 18;
    uint256 private _totalSupply;
    address public owner;//定义合约拥有者
    mapping(address => uint256) private balances; //记录某个人账号余额
    mapping(address => mapping(address => uint256)) private allowances;//记录某个人授权信息
    constructor(uint256 initialSupply) {
        //在构造函数中初始化owner
        owner = msg.sender;
        // 在部署时直接铸造 initialSupply 个代币给合约部署者
        _mint(owner, initialSupply);
    }
    // ===== 自定义功能：mint =====
    function mint(address account, uint256 amount) external {
        require(msg.sender == owner, "Only owner can mint");
        _mint(account, amount);
    }

    // 内部增发函数
    function _mint(address account, uint256 amount) internal {
        balances[account] += amount;
        _totalSupply += amount;
        emit Transfer(address(0), account, amount);
    }

     // ===== IERC20 接口默认实现 =====
    function totalSupply() external view override returns (uint256) {
        return _totalSupply;
    }

    function balanceOf(address account) external view override returns (uint256) {
        return balances[account];
    }

    function transfer(address recipient, uint256 amount) external override returns (bool) {
        require(balances[msg.sender] >= amount, "ERC20: transfer amount exceeds balance");
        balances[msg.sender] -= amount;
        balances[recipient] += amount;
        emit Transfer(msg.sender, recipient, amount);
        return true;
    }

    function approve(address spender, uint256 amount) external override returns (bool) {
        allowances[msg.sender][spender] = amount;
        emit Approval(msg.sender, spender, amount);
        return true;
    }

    function allowance(address owner_, address spender) external view override returns (uint256) {
        return allowances[owner_][spender];
    }

    function transferFrom(address sender, address recipient, uint256 amount) external override returns (bool) {
        require(balances[sender] >= amount, "ERC20: transfer amount exceeds balance");
        require(allowances[sender][msg.sender] >= amount, "ERC20: transfer amount exceeds allowance");

        balances[sender] -= amount;
        balances[recipient] += amount;
        allowances[sender][msg.sender] -= amount;

        emit Transfer(sender, recipient, amount);
        return true;
    }
}