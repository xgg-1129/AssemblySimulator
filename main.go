package main

import (
	"gg/computer"
)

func main() {
	   computer.InitReg(4)
		for i:=0;i<1;i++ {
			computer.Cpu.PrintReg()
			computer.PrintStack()
			computer.InstructionCycle()
			computer.Cpu.PrintReg()
			computer.PrintStack()
			}
















	/*//用来测试 mov M[0x122321],M[0x14214]
	computer.InitReg()
	computer.Cpu.Rsp=0x222
	computer.WriteMemory(1,computer.Vd2pd(0x222))

	//inst := *(*computer.Inst)(unsafe.Pointer(uintptr(computer.Cpu.Rip)))
	//fmt.Println(inst.Str)

	computer.MM[computer.Vd2pd(0x122321)]=2
	computer.MM[computer.Vd2pd(0x14214)]=1
	computer.PrintReg()
	fmt.Println("地址0x122321=",computer.MM[computer.Vd2pd(0x122321)])
	fmt.Println("地址0x14214=",computer.MM[computer.Vd2pd(0x14214)])

	computer.InstructionCycle()
	computer.PrintStack()

	computer.PrintReg()
	fmt.Println("地址0x122321=",computer.MM[computer.Vd2pd(0x122321)])
	fmt.Println("地址0x14214=",computer.MM[computer.Vd2pd(0x14214)])*/



	/*//用来测试 mov 1234 rax指令
	computer.Cpu.Rip = uint64(uintptr(unsafe.Pointer(&computer.Program[1])))
	computer.Cpu.Rax = 1111
	fmt.Println("rip=",computer.Cpu.Rip)
	fmt.Println("rax=",computer.Cpu.Rax)
	fmt.Println("寄存器rax的地址",&computer.Cpu.Rax)
	computer.InstructionCycle()
	fmt.Println("=======指令执行完毕后，寄存器和内存的状态")
	fmt.Println("rip=",computer.Cpu.Rip)
	fmt.Println("rax=",computer.Cpu.Rax)
	 */
}
