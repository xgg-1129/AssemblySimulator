package computer


var (

	Program = []Inst{
		{ Str:  " mov M[0x122321],M[0x14214]",
			Name: Mov_MM2MM,
			Src:  InscObject{
				IMM:         0x122321,
				S:           0,
				RA:          nil,
				RB:          nil,
				AddressType: IMM,
			},
			Dst:  InscObject{
				IMM:         0x14214,
				S:           0,
				RA:          nil,
				RB:          nil,
				AddressType: IMM,
			},
		},
		{ Str:  "mov 0x1234,reg",
			Name: Mov_IMM2Reg,
			Src:  InscObject{
				IMM:         1234,
				S:           0,
				RA:          nil,
				RB:          nil,
				AddressType: IMM,
			},
			Dst:  InscObject{
				IMM:         0x14214,
				S:           0,
				RA:          &Cpu.Rax,
				RB:          nil,
				AddressType: REG,
			},
		},
		{ Str:  "Push 0xffff",
			Name: PushImm,
			Src:  InscObject{
				IMM:         0xffff,
				S:           0,
				RA:          nil,
				RB:          nil,
				AddressType: IMM,
			},
			Dst:  InscObject{
				IMM:         0x14214,
				S:           0,
				RA:          &Cpu.Rax,
				RB:          nil,
				AddressType: REG,
			},
		},
		{ Str:  "CALL 0X0fff",
			Name: IMCall,
			//因为这里go语言无法直接对IMM赋值，所以我这里用了src和dst共同实现
			Src:  InscObject{
				IMM:       	 0x0fff,
				S:           0,
				RA:          &Cpu.Rip,
				RB:          nil,
				AddressType: IMM,
			},
			Dst:  InscObject{
				IMM:         1,
				S:           0,
				RA:          &Cpu.Rax,
				RB:          nil,
				AddressType: IMM,
			},
		},
		{ Str:  "Push %rax",
			Name: PushReg,
			Src:  InscObject{
				IMM:         0xffff,
				S:           0,
				RA:          &Cpu.Rax,
				RB:          nil,
				AddressType: REG,
			},
			Dst:  InscObject{
				IMM:         0x14214,
				S:           0,
				RA:          &Cpu.Rax,
				RB:          nil,
				AddressType: REG,
			},
		},
	}
)



