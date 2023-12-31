pragma solidity >=0.4.19 <0.6.0;
contract Ownable{
    address public owner; 
    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner); 
    /*** @dev The Ownable constructor sets the original `owner` of the contract to the sender * account. */
    constructor () public { owner = msg.sender; }/*** @dev Throws if called by any account other than the owner. */ 
    modifier onlyOwner() { 
        require(msg.sender == owner); 
        _;
    }
    /*** @dev Allows the current owner to transfer control of the contract to a newOwner.
    * @param newOwner The address to transfer ownership to. */
    
}