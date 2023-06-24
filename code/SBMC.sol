pragma solidity >=0.4.19;
pragma experimental ABIEncoderV2;
import "./SSMC.sol";
import "./Ownable.sol";
contract SBMC is Ownable{
    SSMC ssmc;     
    uint DownloadPrice=20;
    function  ControlContract(address _SSMCAddr) onlyOwner public 
    {        
        ssmc = SSMC(_SSMCAddr);   
    }

    struct BuyerInfor
    {
        uint dataId;
        byte[4] IP;
        address Abuyer;
        uint NumOfEth;
    }
    BuyerInfor[] Buyers;
    function BuyerReg(uint _dataId,byte[4] memory _IP) public payable
    {
        require(ssmc.viewSeller(_dataId).price+ssmc.viewSeller(_dataId).PartsNum*DownloadPrice<=msg.value);
        Buyers.push(BuyerInfor(_dataId,_IP,msg.sender,msg.value));
    }
    struct BuyerMatchSP
    {
        uint buyerId;
        uint[] SPID;
        bytes32[] HashOfDoubleEncryptData;
        bool[] confirmed;
    }
    BuyerMatchSP[] BMSPS;
    function SPReg(uint _buyerId,uint[] memory _SPID) public returns(uint)
    {
        require(Buyers[_buyerId].Abuyer==msg.sender);
        BuyerMatchSP memory BMSP;
        BMSP.buyerId=_buyerId;
        BMSP.SPID=_SPID;
        BMSPS.push(BMSP);
        uint i;
        for(i=0;i<BMSP.SPID.length;i++)
        {
            BMSP.HashOfDoubleEncryptData[i]=bytes32(0x00);
            BMSP.confirmed[i]=false;
        }
        return BMSPS.length-1;
    }
    function Buyerconfirmed(uint BMID,uint SPIndex) public
    {
        require(Buyers[BMSPS[BMID].buyerId].Abuyer==msg.sender);
        require(BMSPS[BMID].HashOfDoubleEncryptData[SPIndex]!=bytes32(0x00));
        BMSPS[BMID].confirmed[SPIndex]==true;
    }
    function SetHashOfDEData(uint BMID,uint SPIndex,bytes32 _HODED) public
    {
        
        require((ssmc.viewSPbyIndex(BMSPS[BMID].SPID[SPIndex])).ASP==msg.sender);
        BMSPS[BMID].HashOfDoubleEncryptData[SPIndex]==_HODED;
        
    }
}