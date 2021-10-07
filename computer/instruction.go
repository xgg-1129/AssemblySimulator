package computer
//指令集
import (
	"fmt"
	"unsafe"
)

const InstType = 11

var (
	Functable =[InstType]func(src, dst uint64){
		movIMM2Reg,
		movIMM2MM,
		movMM2MM,
		movMM2Reg,
		movReg2MM,
		Push_IMM,
		Im_Call,
		Id_Call,
		Pop_Reg,
		Push_Reg,
		RetFunc,
	}
)

type code uint64
const (
	Mov_IMM2Reg code =0
	Mov_IMM2MM code =1
	Mov_MM2MM code =2
	Mov_MM2Reg code =3
	Mov_Reg2MM code=4
	PushImm  code=5
	IMCall code=6
	IDCall code=7
	PopReg code=8
	PushReg code=9
	Ret code=10
	)

type DecodeType uint64
const (
	IMM            DecodeType =0
	REG            DecodeType =1
	MM_IMM         DecodeType =2
	MM_RA          DecodeType =3
	MM_IMM_RA      DecodeType =4
	MM_RA_RB       DecodeType =5
	MM_IMM_RA_RB   DecodeType =6
	MM_RA_S        DecodeType =7
	MM_IMM_RA_S    DecodeType =8
	MM_RA_RB_S     DecodeType =9
	MM_IMM_RA_RB_S DecodeType =10
)
type InscObject struct {
	IMM int64
	S uint64
	//ri和rb寄存器 指向一个uint64的地址
	RA *uint64
	RB *uint64

	AddressType DecodeType
}

type Inst struct {
	Str  string
	Name code
	Src  InscObject
	Dst  InscObject
}

func (src InscObject)DecodeAddress() uint64{
	if src.AddressType ==0 {
		return uint64(src.IMM)
	} else if src.AddressType == 1 {
		return *(*uint64)(unsafe.Pointer(&src.RA))
	} else{
		fmt.Print("else")
		index := uint64(src.IMM) + (*src.RB*src.S) + *src.RB
		return index
	}
}

func InstructionCycle(){
	//根据rip指令获取当前指向的指令
	 inst := *(*Inst)(unsafe.Pointer(uintptr(Cpu.Rip)))
	 fmt.Println(inst.Str)
	 var src uint64 = inst.Src.DecodeAddress()
	 var dst uint64 = inst.Dst.DecodeAddress()
	 fun:= Functable[inst.Name]
	 fun(src,dst)
	 //执行完一条指令后，跟新rip寄存器的值
	 Cpu.Rip= Cpu.Rip+uint64(unsafe.Sizeof(inst))
}

func movIMM2Reg(src,dst uint64){
 	//此函数实现mov imm reg
	*(*uint64)(unsafe.Pointer(uintptr(dst)))=src
	return
}
func movIMM2MM(src,dst uint64){
//	//此函数实现mov imm memory
//  直接使用MM[索引]的方法访问，是不经过cache的，后续可以加一个cache，采用函数的方法对内存进行读写
//	传来的是两个地址，因为我这是64的cpu，数据总线为64位，每次都是对内存的64位进行改写
//	64=8*8
	WriteMemory(src,dst)
	return
}

func movMM2MM(src,dst uint64){
	//获取src地址的64位数据，然后传给write函数
	data:= ReadMemory(Vd2pd(src))
	WriteMemory(data,Vd2pd(dst))
	return
}
func movMM2Reg(src,dst uint64){
	data:= ReadMemory(src)
	*(*uint64)(unsafe.Pointer(uintptr(dst)))=data
	return
}
func movReg2MM(src,dst uint64){
	data:=*(*uint64)(unsafe.Pointer(uintptr(src)))
	WriteMemory(data,dst)

	return
}
func Push_IMM(src,dst uint64)  {
	//因为我们的cpu的字是8个字节，所以-8
	Cpu.Rsp=Cpu.Rsp-8
	WriteMemory(src,Vd2pd(Cpu.Rsp))
}
func Push_Reg(src,dst uint64)  {
	//因为我们的cpu的字是8个字节，所以-8
	Cpu.Rsp=Cpu.Rsp-8
	fmt.Println("src11111111111111111111111111111111111111111111",src)
	WriteMemory(*(*uint64)(unsafe.Pointer(uintptr(src))),Vd2pd(Cpu.Rsp))
}
func Pop_Reg(src,dst uint64){
	//弹栈至寄存器
	*((*uint64)(unsafe.Pointer(uintptr(src))))=ReadMemory(Vd2pd(Cpu.Rsp))
	Cpu.Rsp=Cpu.Rsp+8
}

//直接调用
func Im_Call(src,dst uint64) {
	//将返回地址写入新的栈顶
	Cpu.Rsp=Cpu.Rsp-8
	WriteMemory(Cpu.Rip+uint64(unsafe.Sizeof(Inst{})),Vd2pd(Cpu.Rsp))
	//更新R
	Cpu.Rip=src
}
func  Id_Call(src,dst uint64)  {
	//将返回地址写入新的栈顶
	Cpu.Rsp=Cpu.Rsp-8
	WriteMemory(Cpu.Rip+uint64(unsafe.Sizeof(Inst{})),Vd2pd(Cpu.Rsp))

	Cpu.Rip=ReadMemory(Vd2pd(src))
}
func RetFunc(src,dst uint64){
	address:=ReadMemory(Vd2pd(Cpu.Rsp))
	Cpu.Rsp=Cpu.Rsp+8
	//这里的实现和理论符，原因是我再cycle里设置了每次执行指令自动加rip的值
	Cpu.Rip=address-uint64(unsafe.Sizeof(Inst{}))
	Cpu.Rip=address-uint64(unsafe.Sizeof(Inst{}))
}




