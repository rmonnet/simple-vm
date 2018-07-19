package vm

const (
	// IADD  adds the 2 integers on top of the stack
	IADD = 1
	// ISUB subtracts the 2 integers on top of the stack
	ISUB = 2
	// IMUL multipies the 2 integers on top of the stack
	IMUL = 3
	// ILT compares if the integer on top of the stack is less than ...
	ILT = 4
	// IEQ checks if the integer at the top of the stack if equal to ...
	IEQ = 5
	// BR jumps to the given address
	BR = 6
	// BRT jumps to the given address if the condition is true
	BRT = 7
	// BRF jumps to the given address if the condition is false
	BRF = 8
	// ICONST pushes an integer on the stack
	ICONST = 9
	// LOAD loads from local context
	LOAD = 10
	// GLOAD loads from global memory
	GLOAD = 11
	// STORE stores in local context
	STORE = 12
	// GSTORE stores in global memory
	GSTORE = 13
	// PRINT prints the top of the stack
	PRINT = 14
	// POP throws away the top of the stack
	POP = 15
	// CALL calls a function
	CALL = 16
	// RET returns from a function call
	RET = 17
	// HALT stops the current program
	HALT = 18
)

// Instruction represents one of the VM instructions
type Instruction struct {
	name  string
	nargs int
}

// Name returns the name of the instruction
func (inst *Instruction) Name() string {
	return inst.name
}

// NumArgs returns the number of arguments for this instruction
func (inst *Instruction) NumArgs() int {
	return inst.nargs
}

// Instructions maps opcodes to their associated instructions
var Instructions = [...]Instruction{
	{"na", 0},
	{"iadd", 0},
	{"isub", 0},
	{"imul", 0},
	{"ilt", 0},
	{"ieq", 0},
	{"br", 1},
	{"brt", 1},
	{"brf", 1},
	{"iconst", 1},
	{"load", 1},
	{"gload", 1},
	{"store", 1},
	{"gstore", 1},
	{"print", 0},
	{"pop", 0},
	{"call", 1},
	{"ret", 0},
	{"halt", 0}}
