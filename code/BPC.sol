pragma solidity >=0.4.19;
pragma experimental ABIEncoderV2;
import "./SSMC.sol";
import "./SBMC.sol";
contract BPC is SBMC{
    
    SSMC ssmc;
    uint waittime=10000;
    struct BMPay
    {
        uint BMID;
        uint SPIndex;
        uint recordtime;
        bytes32 key;
        bool appeal;
    }
    BMPay[] BMPays;
    function  ControlContract2(address _SSMCAddr) onlyOwner public 
    {        
        ssmc = SSMC(_SSMCAddr);   
    }
    function activate(uint BMID,uint SPIndex,bytes32 SPkey) public
    {
        BMPays.push(BMPay(BMID,SPIndex,now,SPkey,false));
    }
    function SPwithdraw(uint _BMID,uint _SPIndex,uint _BMPayID) public payable
    {
        require(ssmc.viewSPbyIndex(BMSPS[_BMID].SPID[_SPIndex]).ASP==msg.sender);
        require(BMPays[_BMPayID].BMID==_BMID);
        require(BMPays[_BMPayID].SPIndex==_SPIndex);
        require(now>=waittime+BMPays[_BMPayID].recordtime);
        require(BMPays[_BMPayID].appeal=false);
        
        msg.sender.transfer(DownloadPrice);
    }
    function appeal(uint _BMPayID,string memory DEdata) public
    {
        require(sha256(bytes(DEdata))==BMSPS[BMPays[_BMPayID].BMID].HashOfDoubleEncryptData[0]);
        BMPays[_BMPayID].appeal=true;
    }
}