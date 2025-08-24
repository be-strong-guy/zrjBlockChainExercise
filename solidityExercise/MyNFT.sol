// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
/*
任务目标
使用 Solidity 编写一个符合 ERC721 标准的 NFT 合约。
将图文数据上传到 IPFS，生成元数据链接。
将合约部署到以太坊测试网（如 Goerli 或 Sepolia）。
铸造 NFT 并在测试网环境中查看。
任务步骤
编写 NFT 合约
使用 OpenZeppelin 的 ERC721 库编写一个 NFT 合约。
合约应包含以下功能：
构造函数：设置 NFT 的名称和符号。
mintNFT 函数：允许用户铸造 NFT，并关联元数据链接（tokenURI）。
在 Remix IDE 中编译合约。
准备图文数据
准备一张图片，并将其上传到 IPFS（可以使用 Pinata 或其他工具）。
创建一个 JSON 文件，描述 NFT 的属性（如名称、描述、图片链接等）。
将 JSON 文件上传到 IPFS，获取元数据链接。
JSON文件参考 https://docs.opensea.io/docs/metadata-standards
部署合约到测试网
在 Remix IDE 中连接 MetaMask，并确保 MetaMask 连接到 Goerli 或 Sepolia 测试网。
部署 NFT 合约到测试网，并记录合约地址。
铸造 NFT
使用 mintNFT 函数铸造 NFT：
在 recipient 字段中输入你的钱包地址。
在 tokenURI 字段中输入元数据的 IPFS 链接。
在 MetaMask 中确认交易。
查看 NFT
打开 OpenSea 测试网 或 Etherscan 测试网。
连接你的钱包，查看你铸造的 NFT。
*/

/*
图片上传成功在：https://ipfs.io/ipfs/bafkreihlvxzs2fqvvkpbnmjpssmap47w6aumoscpi55pahfks2erg6cmwy
sepolia测试也成功地址在：
https://sepolia.etherscan.io/address/0xFDDd75fc06042E05276D5D09C29866ad86C03CA7
个人钱包中，也出现收藏品一张图片
*/
// 导入 OpenZeppelin ERC721 库
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract MyNFT is ERC721, Ownable {
    uint256 public nextTokenId;

    // tokenId -> tokenURI 映射
    mapping(uint256 => string) private _tokenURIs;


    constructor(string memory _name, string memory _symbol)
     ERC721(_name, _symbol)
      Ownable(msg.sender)  {
        nextTokenId = 1; // tokenId 从1开始
    }

    // 铸造 NFT
    function mintNFT(address recipient, string memory _tokenURI) external returns (uint256) {
        uint256 tokenId = nextTokenId;
        _safeMint(recipient, tokenId);
        _setTokenURI(tokenId, _tokenURI);
        nextTokenId++;
        return tokenId;
    }

    // 内部函数，设置 tokenURI
    function _setTokenURI(uint256 _tokenId, string memory _tokenURI) internal tokenIsexist(_tokenId) {
        _tokenURIs[_tokenId] = _tokenURI;
    }

    // 覆盖 ERC721 的 tokenURI 函数
    function tokenURI(uint256 _tokenId) public view override tokenIsexist(_tokenId) returns (string memory) {
        return _tokenURIs[_tokenId];
    }
    modifier tokenIsexist(uint256 _tokenId){
        require(ownerOf(_tokenId) != address(0), "ERC721Metadata: URI query for nonexistent token");
        _;
    }
}
