pragma solidity >=0.4.19;
pragma experimental ABIEncoderV2;
contract SSMC
{
    struct SellerData
    {
        address Aseller;
        string Description;
        uint PartsNum;
        string HashOfParts;
        string HashOfEncryptParts;
        uint RandIndex;
        uint price;
    }
    struct ExposedPart
    {
        string[] plaintexts;
        string[] EncryptParts;
    }
    SellerData[] Sellers;
    ExposedPart[] ExposedParts; 
    function SellerDataset(string memory _des,uint _PartsNum,string memory _HOP,string memory _HOEP,uint _price) public
    {
        require(_PartsNum<=10000);
        require(_PartsNum>=10);
        require(bytes(_HOP).length==bytes(_HOEP).length);
        require(bytes(_HOP).length==_PartsNum*256);
        Sellers.push(SellerData(msg.sender,_des,_PartsNum,_HOP,_HOEP,_rand(_PartsNum),_price));
        
    }
    function _rand(uint256 _length) public view returns(uint256) {
        uint256 random = uint256(keccak256(abi.encodePacked(block.difficulty, now)));
        return random%_length;
    }
    function _BytesTransfer(string memory str,uint index) private returns(bytes32)
    {
        bytes32 result=0x0;
        uint i;
        for(i=0;i<32;i++)
        {
            result|=bytes(str)[index+i];
            result<<=8;
        }
    }
    
    function ExposeParts(uint _DataId,string[] memory Plaintexts,string[] memory EncryptParts) public
    {
        require(Sellers[_DataId].Aseller==msg.sender);
        uint i;
        for(i=0;i<10;i++)
        {
            assert(sha256(bytes(Plaintexts[i]))==_BytesTransfer(Sellers[_DataId].HashOfParts,Sellers[_DataId].RandIndex+i*Sellers[_DataId].PartsNum/10));
            assert(sha256(bytes(EncryptParts[i]))==_BytesTransfer(Sellers[_DataId].HashOfEncryptParts,Sellers[_DataId].RandIndex+i*Sellers[_DataId].PartsNum/10));
            ExposedParts[_DataId]=(ExposedPart(Plaintexts,EncryptParts));
        }
    }
    struct SPInfor
    {
        uint dataId;
        byte[4] IP;
        address ASP;
        bool confirmed;
    }
    SPInfor[] SPInfors;
    function SPReg(uint _DataId,byte[4] memory IP) public
    {
        SPInfors.push(SPInfor(_DataId,IP,msg.sender,false));
    }
    function Confirm(uint _SPId) public
    {
        uint _dataid=SPInfors[_SPId].dataId;
        require(Sellers[_dataid].Aseller==msg.sender);
        SPInfors[_SPId].confirmed=true;
    }
    function viewSeller(uint _ID)  public returns(SellerData memory)
    {
        return Sellers[_ID];
    }
    function viewSPbydataId(uint _dataId) view public returns(SPInfor[] memory)
    {
        uint i;uint j;
        SPInfor[] memory SPS;
        j=0;
        for(i=0;i<SPInfors.length;i++)
        {
            if(SPInfors[i].dataId==_dataId)
            {
                SPS[j]=SPInfors[i];
                j++;
            }
        }
        return SPS;
    }
    function viewSPbyIndex(uint _index) view public returns(SPInfor memory)
    {
        return SPInfors[_index];
    }
}