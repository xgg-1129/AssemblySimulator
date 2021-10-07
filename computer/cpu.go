package computer

import (
	"fmt"
	"unsafe"
)

//64位的cpu

/*最好使用联合体实现rax和rah、ral，但是golang不支持union，[也许结构体嵌套可以实现？]
所以我这里暂时仅仅实现了64位的寄存器
 */

var Cpu Reg
type Reg struct {
	//通用寄存器
	Rax uint64
	Rbx uint64
	Rcx uint64
	Rdx uint64
	Rsi uint64
	Rbp uint64
	//栈指针
	Rsp uint64

	//指令指针
	Rip uint64

	Mmu uint64
	PageTable uint64
}

func InitReg(index int){
	Cpu.Rax=111
	Cpu.Rbx=112
	Cpu.Rcx=113
	Cpu.Rdx=114
	Cpu.Rsi=115
	Cpu.Rbp=116
	//栈顶
	Cpu.Rsp=200
	Cpu.Rip=uint64(uintptr(unsafe.Pointer(&Program[index])))
}
func (r Reg) PrintReg()  {
	fmt.Println("===========当前cpu寄存器的状态============")
	fmt.Printf("Rax=%#016x   Rbx=%#016x  Rcx=%#016x \nRdx=%#016x   Rsi=%#016x  Rbp=%#016x\n",
		r.Rax,r.Rbx,r.Rcx,r.Rdx,r.Rsi,r.Rbp)
	fmt.Printf("栈顶寄存器        Rsp=%#016x\n",r.Rsp)
	fmt.Printf("指令ip寄存器      Rip=%#016x\n",r.Rip)
	fmt.Println("=======================================")
}

func PrintStack(){
	fmt.Println("===========当前的程序的栈帧============")
	//打印栈附件的20个字
	//获取栈顶
	n:=10
	var top uint64= uint64(uintptr(unsafe.Pointer(&MM[Vd2pd(Cpu.Rsp)])))


	top=top+uint64(n)*uint64(unsafe.Sizeof(top))
	vd:=Cpu.Rsp
	vd=vd+uint64(n)*uint64(unsafe.Sizeof(vd))
	for i:=0;i<2*n;i++{
		//获取每个address和data
		address := (*uint64)(unsafe.Pointer(uintptr(top)))
		data := *address
		fmt.Printf("%#016x: %#016x",vd,data)
		//sizeof(uint64)=8，所以可以把函数直接改为8
		top=top-uint64(unsafe.Sizeof(top))
		vd=vd-8
		if i==n{
			fmt.Printf("<==rsp 当前栈顶")
		}
		fmt.Printf("\n")
	}
	fmt.Println("=======================================")
}

func cacheExit() bool{
	return false
}


//
