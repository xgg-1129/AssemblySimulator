package computer

const MemoryLength = 8*1000

var (
	MM [MemoryLength]uint8
)
//这里实际传的dst是物理地址，无需进行地址转换
func WriteMemory(data ,dst uint64){
	if cacheExit()   {

	}else{
		MM[dst+0] = uint8(data & 0xff)
		MM[dst+1] = uint8(data>>8 & 0xff)
		MM[dst+2] = uint8(data>>16 & 0xff)
		MM[dst+3] = uint8(data>>24 & 0xff)
		MM[dst+4] = uint8(data>>32 & 0xff)
		MM[dst+5] = uint8(data>>40 & 0xff)
		MM[dst+6] = uint8(data>>48 & 0xff)
		MM[dst+7] = uint8(data>>56 & 0xff)
	}

	return
}
//返回dst的64位置数据
func ReadMemory(dst uint64)uint64{
	if cacheExit() {
		return 0
	}else{
		var data uint64 = 0
		data=data | uint64(MM[dst+0])
		data=data | (uint64(MM[dst+1])<<8)
		data=data | (uint64(MM[dst+2])<<16)
		data=data | (uint64(MM[dst+3])<<24)
		data=data | (uint64(MM[dst+4])<<32)
		data=data | (uint64(MM[dst+5])<<40)
		data=data | (uint64(MM[dst+6])<<48)
		data=data | (uint64(MM[dst+7])<<56)
		return data
	}

}




