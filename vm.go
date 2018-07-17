package vm

import "fmt"

const (
	stackSize = 1000
)

// VM represents a Virtual Machine.
type VM struct {
	data  []int
	code  []int
	stack []int
	sp    int
	fp    int
	ip    int
	trace bool
}

// NewVM creates a new Virtual Machine.
func NewVM(code []int, main int, datasize int) *VM {
	vm := new(VM)
	vm.code = make([]int, len(code))
	copy(vm.code, code)
	vm.ip = main
	vm.data = make([]int, datasize)
	vm.stack = make([]int, stackSize)
	vm.sp = -1
	// frame pointer only really makes sense inside a CALL
	// starting at -1 is like pretenting the main program
	// is within a call (but no argument or return values are
	// available).
	vm.fp = -1
	return vm
}

// SetTrace allows to switch tracing on and off
func (vm *VM) SetTrace(on bool) {
	vm.trace = on
}

func (vm *VM) disassemble(opcode int) {
	inst := Instructions[opcode]
	fmt.Printf("%04d: %s", vm.ip-1, inst.Name())
	for i := 0; i < inst.NumArgs(); i++ {
		fmt.Printf(" %d", vm.code[vm.ip+i])
	}
	fmt.Printf(" %v", vm.stack[0:vm.sp+1])
	fmt.Println("")
}

func (vm *VM) dumpMemory() {
	if len(vm.data) <= 0 {
		return
	}
	fmt.Println("Memory:")
	for i := 0; i < len(vm.data); i++ {
		fmt.Printf("%04d: %d\n", i, vm.data[i])
	}
}

// Exec executes the program loaded in the Virtual Machine
func (vm *VM) Exec() {
	for vm.ip < len(vm.code) {
		// fetch
		opcode := vm.fetch()
		if vm.trace {
			vm.disassemble(opcode)
		}
		// execute
		switch opcode {
		case HALT:
			break
		case ICONST:
			v := vm.fetch()
			vm.push(v)
		case PRINT:
			v := vm.pop()
			fmt.Println(v)
		case GSTORE:
			v := vm.pop()
			addr := vm.fetch()
			vm.data[addr] = v
		case GLOAD:
			addr := vm.fetch()
			vm.push(vm.data[addr])
		case ILT:
			r := vm.pop()
			l := vm.pop()
			cond := 0
			if l < r {
				cond = 1
			}
			vm.push(cond)
		case IEQ:
			r := vm.pop()
			l := vm.pop()
			cond := 0
			if l == r {
				cond = 1
			}
			vm.push(cond)
		case BRF:
			cond := vm.pop()
			addr := vm.fetch()
			if cond == 0 {
				vm.ip = addr
			}
		case BRT:
			cond := vm.pop()
			addr := vm.fetch()
			if cond != 0 {
				vm.ip = addr
			}
		case BR:
			addr := vm.fetch()
			vm.ip = addr
		case IADD:
			r := vm.pop()
			l := vm.pop()
			vm.push(l + r)
		case ISUB:
			r := vm.pop()
			l := vm.pop()
			vm.push(l - r)
		case IMUL:
			r := vm.pop()
			l := vm.pop()
			vm.push(l * r)
		case POP:
			vm.pop()
		case LOAD:
			i := vm.fetch()
			vm.push(vm.stack[vm.fp+i])
		case STORE:
			i := vm.fetch()
			vm.stack[vm.fp+i] = vm.pop()
		default:
			panic(fmt.Sprintf("unrecognized opcode: %d", opcode))
		}
	}
	if vm.trace {
		vm.dumpMemory()
	}

}

func (vm *VM) fetch() int {
	opcode := vm.code[vm.ip]
	vm.ip++
	return opcode
}

func (vm *VM) push(v int) {
	vm.sp++
	vm.stack[vm.sp] = v
}

func (vm *VM) pop() int {
	v := vm.stack[vm.sp]
	vm.sp--
	return v
}
