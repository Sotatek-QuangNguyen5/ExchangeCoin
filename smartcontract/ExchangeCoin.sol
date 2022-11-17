// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.9.0;

contract ExchangeCoin {
    
    address public creator;
    mapping (bytes => bool) listReceiver;
    event eventReceiverMoney(address addr, uint256 amount);
    event eventWithDrawMoney(address addr, bytes signature);

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
        
        require(listReceiver[signature] == false, "signature expire time!!!");
        require(receiver == msg.sender, "wrong receiver!!!");
        require(amount <= address(this).balance, "contract is not enough money!!!");
        bytes32 messageHash = getMessageHash(receiver, message, amount);
        bytes32 ethSignedMessageHash = getEthSignedMessageHash(messageHash);
        require(recoverSigner(ethSignedMessageHash, signature) == creator, "invalid signature!!!");
        receiver.transfer(amount);
        listReceiver[signature] = true;
        emit eventWithDrawMoney(receiver, signature);
    }

    function receiveMoney() public payable {

        emit eventReceiverMoney(msg.sender, msg.value);
    }

    function getBalance() public view returns(uint256) {

        return address(this).balance;
    }
}