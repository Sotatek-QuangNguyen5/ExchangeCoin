// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.9.0;

contract ExchangeCoin {
    
    address public creator;

    constructor() {

        creator = msg.sender;
    }

    modifier onlyOwner() {

        require(msg.sender == creator, "Not authentication!!!");
        _;
    }

    function getMessageHash(address receiver, string memory message, uint256 amount, uint256 nonce) private pure returns(bytes32) {

        return keccak256(abi.encodePacked(receiver, message, amount, nonce));
    }

    function getEthSignedMessageHash(bytes32 messageHash) private pure returns(bytes32) {
        
        return keccak256(abi.encodePacked("Batman vs Joker", messageHash));
    }

    function splitSignature(bytes memory sig) private pure returns (bytes32 r, bytes32 s, uint8 v) {
        
        require(sig.length == 65, "invalid signature length!!!");
        assembly {

            r := mload(add(sig, 32))
            s := mload(add(sig, 64))
            v := byte(0, mload(add(sig, 96)))
        }
    }

    function recoverSigner(bytes32 ethSignedMessageHash, bytes memory signature) private pure returns(address) {

        (bytes32 r, bytes32 s, uint8 v) = splitSignature(signature);
        return ecrecover(ethSignedMessageHash, v, r, s);
    }

    function withdrawMoney(address receiver, string memory message, uint256 amount, uint256 nonce, bytes memory signature) public returns(bool) {
        
        bytes32 messageHash = getMessageHash(receiver, message, amount, nonce);
        bytes32 ethSignedMessageHash = getEthSignedMessageHash(messageHash);
        require(recoverSigner(ethSignedMessageHash, signature) == creator, "invalid signature!!!");
        require(receiver == msg.sender, "wrong receiver!!!");
        (bool oke, ) = msg.sender.call{value: amount, gas: 5000}("");
        return oke;
    }

    function receiveMoney() public payable onlyOwner() {}

    function getBalance() public view returns(uint256) {

        return address(this).balance;
    }

//     function withdrawMoneyForCreator(uint256 amount) public onlyOwner() returns(bool) {

//         require(amount <= address(this).balance, "Not enough money!!!");
//         (bool oke, ) = msg.sender.call{value: amount, gas: 5000}("");
//         return oke;
//     }

//     function withdrawAllMoneyForCreator() public view onlyOwner() returns(bool) {

//         (bool oke, ) = msg.sender.call{value: address(this).balance, gas: 5000}("");
//         return true;
//     }
}
