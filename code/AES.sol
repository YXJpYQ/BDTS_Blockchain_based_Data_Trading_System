pragma solidity >=0.4.19;
pragma experimental ABIEncoderV2;
contract AES
{
	
	uint8[4][4] mix_martix;
	function AES_mul(byte a,uint8 b) private returns(byte)
	{
		byte flag=0x80;byte const=0x1b;byte x=a;
		if(b>=2)
		{
			if(a&flag!=0)
			{a<<=1;a^=const;}
			else
			{a<<=1;}
			
			
			
			if(b==3)
			{
				a^=x;
			}
		}
		
		return a;
	}
	
	
	function mix_col(byte[4][4] memory s) public returns(byte[4][4] memory)
	{
		byte[4][4] memory martix;
		uint8 i; uint8 j;uint8 r;byte tem;
		for(i=0;i<4;i++)
		{
			for(j=0;j<4;j++)
			{
				tem=0;
				for(r=0;r<4;r++)
				{
					tem^=AES_mul(s[r][j],mix_martix[i][r]);
				}
				martix[i][j]=tem;
			}
		}
		return martix;
	}
	constructor() public
	{
	mix_martix=[[2,3,1,1],[1,2,3,1],[1,1,2,3],[3,1,1,2]];
	rcon=[[byte(0x01),0x00,0x00,0x00],[byte(0x02),0x00,0x00,0x00],[byte(0x04),0x00,0x00,0x00],[byte(0x08),0x00,0x00,0x00],[byte(0x10),0x00,0x00,0x00],
    [byte(0x20),0x00,0x00,0x00],[byte(0x40),0x00,0x00,0x00],[byte(0x80),0x00,0x00,0x00],[byte(0x1B),0x00,0x00,0x00],[byte(0x36),0x00,0x00,0x00]];
	S=[byte(0x63), 0x7C, 0x77, 0x7B, 0xF2, 0x6B, 0x6F, 0xC5, 0x30, 0x01, 0x67, 0x2B, 0xFE, 0xD7, 0xAB, 0x76,
        0xCA, 0x82, 0xC9, 0x7D, 0xFA, 0x59, 0x47, 0xF0, 0xAD, 0xD4, 0xA2, 0xAF, 0x9C, 0xA4, 0x72, 0xC0,
        0xB7, 0xFD, 0x93, 0x26, 0x36, 0x3F, 0xF7, 0xCC, 0x34, 0xA5, 0xE5, 0xF1, 0x71, 0xD8, 0x31, 0x15,
        0x04, 0xC7, 0x23, 0xC3, 0x18, 0x96, 0x05, 0x9A, 0x07, 0x12, 0x80, 0xE2, 0xEB, 0x27, 0xB2, 0x75,
        0x09, 0x83, 0x2C, 0x1A, 0x1B, 0x6E, 0x5A, 0xA0, 0x52, 0x3B, 0xD6, 0xB3, 0x29, 0xE3, 0x2F, 0x84,
        0x53, 0xD1, 0x00, 0xED, 0x20, 0xFC, 0xB1, 0x5B, 0x6A, 0xCB, 0xBE, 0x39, 0x4A, 0x4C, 0x58, 0xCF,
        0xD0, 0xEF, 0xAA, 0xFB, 0x43, 0x4D, 0x33, 0x85, 0x45, 0xF9, 0x02, 0x7F, 0x50, 0x3C, 0x9F, 0xA8,
        0x51, 0xA3, 0x40, 0x8F, 0x92, 0x9D, 0x38, 0xF5, 0xBC, 0xB6, 0xDA, 0x21, 0x10, 0xFF, 0xF3, 0xD2,
        0xCD, 0x0C, 0x13, 0xEC, 0x5F, 0x97, 0x44, 0x17, 0xC4, 0xA7, 0x7E, 0x3D, 0x64, 0x5D, 0x19, 0x73,
        0x60, 0x81, 0x4F, 0xDC, 0x22, 0x2A, 0x90, 0x88, 0x46, 0xEE, 0xB8, 0x14, 0xDE, 0x5E, 0x0B, 0xDB,
        0xE0, 0x32, 0x3A, 0x0A, 0x49, 0x06, 0x24, 0x5C, 0xC2, 0xD3, 0xAC, 0x62, 0x91, 0x95, 0xE4, 0x79,
        0xE7, 0xC8, 0x37, 0x6D, 0x8D, 0xD5, 0x4E, 0xA9, 0x6C, 0x56, 0xF4, 0xEA, 0x65, 0x7A, 0xAE, 0x08,
        0xBA, 0x78, 0x25, 0x2E, 0x1C, 0xA6, 0xB4, 0xC6, 0xE8, 0xDD, 0x74, 0x1F, 0x4B, 0xBD, 0x8B, 0x8A,
        0x70, 0x3E, 0xB5, 0x66, 0x48, 0x03, 0xF6, 0x0E, 0x61, 0x35, 0x57, 0xB9, 0x86, 0xC1, 0x1D, 0x9E,
        0xE1, 0xF8, 0x98, 0x11, 0x69, 0xD9, 0x8E, 0x94, 0x9B, 0x1E, 0x87, 0xE9, 0xCE, 0x55, 0x28, 0xDF,
        0x8C, 0xA1, 0x89, 0x0D, 0xBF, 0xE6, 0x42, 0x68, 0x41, 0x99, 0x2D, 0x0F, 0xB0, 0x54, 0xBB, 0x16];
		
	}
    function _array_remove(uint[] memory arr,uint HeadIndex,uint TailIndex) private
    //remove elements from arr[HeadIndex] to arr[HeadIndex](contain tail)
    {
        uint i;
        uint len=TailIndex-HeadIndex+1;
        for(i=HeadIndex;i<=arr.length-len+1;i++)
        {
            arr[i]=arr[i+len];
        }
    }
    function _OneVer2TwoVer(bytes memory OneVer) private returns(byte[4][4] memory)
    //one ver 2 TwoVer [4][4]
    {
	
        uint i;
        uint j;
        byte[4][4] memory TwoVer;
        for(i=0;i<4;i++)
        {
            for(j=0;j<4;j++)
            {
                TwoVer[i][j]=OneVer[i*4+j];
            }
        }
        return TwoVer;
    }
	function _Inv_OneVer2TwoVer(bytes memory OneVer) private returns(byte[4][4] memory)
    //one ver 2 TwoVer [4][4]
    {
	
        uint i;
        uint j;
        byte[4][4] memory TwoVer;
        for(i=0;i<4;i++)
        {
            for(j=0;j<4;j++)
            {
                TwoVer[j][i]=OneVer[i*4+j];
            }
        }
        return TwoVer;
    }
	
	
    byte[4][10] rcon;
	
    byte[256] S;
    function _subWords(byte Word) private returns(byte)
    {
		return S[uint8(Word)];
    }
    function bytesToUint(bytes memory b) public view returns (uint256){
        
        uint256 number=0;
        for(uint i= 0; i<b.length; i++){
            number = number + uint8(b[i])*(2**(8*(b.length-(i+1))));
        }
        return (number);
    }
    function _ConstArr2Arr(byte[4] memory ConstA) private returns (bytes memory)
    {
        bytes memory Arr=new bytes(4);
        Arr[0]=ConstA[0];Arr[1]=ConstA[1];Arr[2]=ConstA[2];Arr[3]=ConstA[3];
        return Arr;
    }
    function _RotWord(byte[4] memory w) private returns(byte[4] memory)
    {
        byte temp=w[0];
        w[0]=w[1];w[1]=w[2];w[2]=w[3];w[3]=temp;
        return w;
    }
    function _XOR(byte[4] memory Arr_a,byte[4] memory Arr_b) private returns(byte[4] memory)
    {
        Arr_a[0]^=Arr_b[0]; Arr_a[1]^=Arr_b[1]; Arr_a[2]^=Arr_b[2]; Arr_a[3]^=Arr_b[3];
        return Arr_a;
    }
    function abc(string memory a) public view returns(byte)
	{
		return S[uint8(bytes(a)[0])];
	}
	byte[4][44] w;
	function viewW() public view returns(byte[4][44] memory)
	{
	return w;
	}
    function _KeyExpansions(byte[4][4] memory key) public returns (byte[4][44] memory)
    {
        uint i;
        uint j;
        
        for(i=0;i<4;i++)
        {
            for(j=0;j<4;j++)
            {
                w[i][j]=key[i][j];
            }
        }
		
        for(i=4;i<44;i++)
        {
            if(i%4==0)
            {
				
                w[i]=_RotWord(w[i-1]);
				
                for(j=0;j<4;j++)
                {
				
					w[i][j]=_subWords(w[i][j]);
      
                }
                w[i]=_XOR(w[i],rcon[i/4-1]);
                w[i]=_XOR(w[i-4],w[i]);
            }
            if(i%4!=0)
            {
                w[i]=_XOR(w[i-4],w[i-1]);
            }
        }
		return w;
    }
	function _Row_Shift(byte[4] memory w,uint8 i) private returns(byte[4] memory)
	{
		byte[4] memory x;uint j;
		for(j=0;j<4;j++)
		{
			x[(j-i)%4]=w[j];
		}
		return x;
	}
	function _Rows_Shift(byte[4][4] memory w) private returns(byte[4][4] memory)
	{
		uint8 i;
		for(i=0;i<4;i++)
		{
			w[i]=_Row_Shift(w[i],i);
		}
		return w;
	}
    function  AES_Encrypt(string memory plaintext,string memory key) public returns(byte[4][4] memory)
    {
        require(bytes(plaintext).length==16);
        require(bytes(key).length==16);
        return _AES_Encrypt(bytes(plaintext),bytes(key));
        
    }
	function add_roundkey(byte[4][4] memory a,byte[4][44] memory b,uint dev) private returns(byte[4][4] memory)
	{
		uint i;uint j;
		for(i=0;i<4;i++)
			for(j=0;j<4;j++)
				a[j][i]^=b[i+dev][j];
		return a;
	}
	function SubWords(byte[4][4] memory s) private returns(byte[4][4] memory)
	{
		uint8 i;uint8 j;
		for(i=0;i<4;i++)
			for(j=0;j<4;j++)
				s[i][j]=_subWords(s[i][j]);
		return s;
	}
	byte[4][4] state;
	function viewS() public view returns(byte[4][4] memory)
	{return state;}
	
    function _AES_Encrypt(bytes memory plaintext,bytes memory _key) public returns(byte[4][4] memory)
    {
        state=_Inv_OneVer2TwoVer(plaintext);
        byte[4][4] memory key=_OneVer2TwoVer(_key);
		
        w = _KeyExpansions(key);
		
		state=add_roundkey(state,w,0);
		
		uint i;
		for(i=1;i<10;i++)
		{
			state=SubWords(state);
			state=_Rows_Shift(state);
			state=mix_col(state);
			state=add_roundkey(state,w,4*i);
		}
		state=SubWords(state);
		state=_Rows_Shift(state);
		state=add_roundkey(state, w,40);
		return state;
    }
	
	
}