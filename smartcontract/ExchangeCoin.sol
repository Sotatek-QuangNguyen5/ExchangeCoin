// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.9.0;

contract ExchangeCoin {
    
    address public creator;
    mapping (address => bool) listReceiver;
    event eventReceiverMoney(address addr, uint256 amount);
    event eventWithDrawMoney(address addr);

    constructor() {

        creator = msg.sender;
    }

    function getMessageHash(address receiver, string memory message, uint256 amount) public pure returns(bytes32) {

        return keccak256(abi.encodePacked(receiver, message, amount));
    }

    function getEthSignedMessageHash(bytes32 messageHash) public pure returns(bytes32) {
        
        return keccak256(abi.encodePacked("Batman vs Joker", messageHash));
    }

    function splitSignature(bytes memory sig) public pure returns (bytes32 r, bytes32 s, uint8 v) {
        
        require(sig.length == 65, "invalid signature length!!!");
        assembly {

            r := mload(add(sig, 32))
            s := mload(add(sig, 64))
            v := byte(0, mload(add(sig, 96)))
        }
    }

    function recoverSigner(bytes32 ethSignedMessageHash, bytes memory signature) public pure returns(address) {

        (bytes32 r, bytes32 s, uint8 v) = splitSignature(signature);
        return ecrecover(ethSignedMessageHash, v, r, s);
    }

    function withdrawMoney(address payable receiver, string memory message, uint256 amount, bytes memory signature) public {
        
        bytes32 messageHash = getMessageHash(receiver, message, amount);
        bytes32 ethSignedMessageHash = getEthSignedMessageHash(messageHash);
        require(recoverSigner(ethSignedMessageHash, signature) == creator, "invalid signature!!!");
        require(receiver == msg.sender, "wrong receiver!!!");
        require(amount <= address(this).balance, "not enough money!!!");
        require(listReceiver[msg.sender] == false, "you received money!!!");
        receiver.transfer(amount);
        listReceiver[msg.sender] = true;
        emit eventWithDrawMoney(receiver);
    }

    function receiveMoney() public payable {

        listReceiver[msg.sender] = false;
        emit eventReceiverMoney(msg.sender, msg.value);
    }

    function getBalance() public view returns(uint256) {

        return address(this).balance;
    }
}