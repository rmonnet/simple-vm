package vm

import (
	"testing"
)

func Test_badopcode(t *testing.T) {
	defer func() {
		recover()
	}()
	pgm := []int{66}
	vm := NewVM(pgm, 0, 0)
	vm.Exec()
	t.Errorf("bad opcode didn't cause a panic")
}
func ExampleVM_halt() {
	pgm := []int{
		HALT}
	vm := NewVM(pgm, 0, 0)
	vm.SetTrace(true)
	vm.Exec()
	// Output:
	// 0000: halt []
}

func ExampleVM_print() {
	pgm := []int{
		ICONST, 99,
		PRINT,
		HALT}
	vm := NewVM(pgm, 0, 0)
	vm.SetTrace(true)
	vm.Exec()
	// Output:
	// 0000: iconst 99 []
	// 0002: print [99]
	// 99
	// 0003: halt []
}

func ExampleVM_eq() {
	pgm := []int{
		ICONST, 5,
		ICONST, 5,
		IEQ,
		PRINT,
		ICONST, 5,
		ICONST, 4,
		IEQ,
		PRINT,
	}
	vm := NewVM(pgm, 0, 0)
	vm.Exec()
	// Output:
	// 1
	// 0
}

func ExampleVM_brt() {
	pgm := []int{
		ICONST, 0, // 0
		BRT, 8, // 2
		ICONST, 5, // 4
		BR, 10, // 6
		ICONST, 10, // 8
		PRINT, // 10
	}
	vm := NewVM(pgm, 0, 0)
	vm.Exec()
	// Output: 5
}

func ExampleVM_brt2() {
	pgm := []int{
		ICONST, 1, // 0
		BRT, 8, // 2
		ICONST, 5, // 4
		BR, 10, // 6
		ICONST, 10, // 8
		PRINT, // 10
	}
	vm := NewVM(pgm, 0, 0)
	vm.Exec()
	// Output: 10
}

func ExampleVM_brf() {
	pgm := []int{
		ICONST, 0, // 0
		BRF, 8, // 2
		ICONST, 5, // 4
		BR, 10, // 6
		ICONST, 10, // 8
		PRINT, // 10
	}
	vm := NewVM(pgm, 0, 0)
	vm.Exec()
	// Output: 10
}

func ExampleVM_pop() {
	pgm := []int{
		ICONST, 10,
		ICONST, 5,
		POP,
		PRINT,
	}
	vm := NewVM(pgm, 0, 0)
	vm.Exec()
	// Output: 10
}

func ExampleVM_loadstore() {
	pgm := []int{
		// reserve space for 2 locals
		ICONST, 0, // local1, bottom of the stack (fp+1)
		ICONST, 1, // local2, 2nd on stack (fp + 2)
		LOAD, 1, // retrieve local1
		PRINT,
		LOAD, 2, // retrieve local2
		PRINT,
		ICONST, 10,
		STORE, 1, // local1 = 10
		ICONST, 20,
		STORE, 2, // local2 = 20
		LOAD, 1, // retrieve local1
		PRINT,
		LOAD, 2, // retrieve local2
		PRINT,
	}
	vm := NewVM(pgm, 0, 0)
	vm.Exec()
	// Output:
	// 0
	// 1
	// 10
	// 20
}

func ExampleVM_global() {
	pgm := []int{
		ICONST, 99,
		GSTORE, 0,
		GLOAD, 0,
		PRINT,
		HALT}
	vm := NewVM(pgm, 0, 1)
	vm.SetTrace(true)
	vm.Exec()
	// Output:
	// 0000: iconst 99 []
	// 0002: gstore 0 [99]
	// 0004: gload 0 []
	// 0006: print [99]
	// 99
	// 0007: halt []
	// Memory:
	// 0000: 99
}

func ExampleVM_isub() {
	pgm := []int{
		ICONST, 10,
		ICONST, 3,
		ISUB,
		PRINT,
		HALT,
	}
	vm := NewVM(pgm, 0, 0)
	vm.Exec()
	// Output: 7
}

func ExampleVM_imul() {
	pgm := []int{
		ICONST, 10,
		ICONST, 3,
		IMUL,
		PRINT,
		HALT,
	}
	vm := NewVM(pgm, 0, 0)
	vm.Exec()
	// Output: 30
}

func ExampleVM_call() {
	iN := -3    // second arg for power(x, n) at fp-3
	iX := -4    // first arg for power(x, n) at fp-4
	iRes := 1   // local 1 for power(x, n) at fp+1
	iIdx := 2   // local 2 for power(x, n) at fp+2
	lPow := 0   // entry point for power(x, n)
	lInc := 13  // label for x increment
	lTest := 4  // label for power test
	lMain := 29 // label for the main entry point
	pgm := []int{
		// lPow - pow(x, n)
		ICONST, 1, // 0 - iRes (local 1)
		ICONST, 0, // 2 - iIdx (local 2)
		// lTest
		LOAD, iIdx, // 4
		LOAD, iN, // 6
		ILT,       // 8
		BRT, lInc, // 9
		POP, // 11 -remove local 2 from stack
		RET, // 12
		// lInc
		LOAD, iX, // 13
		LOAD, iRes, // 15
		IMUL,        // 17
		STORE, iRes, // 18
		LOAD, iIdx, // 20
		ICONST, 1, // 22
		IADD,        // 24
		STORE, iIdx, // 25
		BR, lTest, // 27
		// lMain - main entry point
		ICONST, 2, // 29 - x=2
		ICONST, 3, // 31 - n=3
		ICONST, 2, // 33 - nargs = 2
		CALL, lPow, // 35 - power(2,3)
		PRINT,     // 37
		ICONST, 3, // 39 - x=3
		ICONST, 4, // 41 - n=4
		ICONST, 2, // 43 - nargs = 2
		CALL, lPow, // 45 - power(3,4)
		PRINT, // 81
	}
	vm := NewVM(pgm, lMain, 0)
	vm.Exec()
	// Output:
	// 8
	// 81
}

func ExampleVM_loop() {
	// global variables
	gN := 0
	gI := 1
	// labels used by the program (by address)
	lStart := 8
	lDone := 27
	pgm := []int{
		// N = 10
		ICONST, 10, // 0
		GSTORE, gN, // 2
		// I = 0
		ICONST, 0, // 4
		GSTORE, gI, // 6
		// WHILE I < N:
		// START (8):
		GLOAD, gI, // 8
		GLOAD, gN, // 10
		ILT,        // 12
		BRF, lDone, // 13
		// I = I + 1
		GLOAD, gI, // 15
		ICONST, 1, // 17
		IADD,       // 19
		GSTORE, gI, // 20
		GLOAD, gI, // 22
		PRINT,      // 24
		BR, lStart, // 25
		// DONE (27):
		HALT, // 27
	}
	vm := NewVM(pgm, 0, 2)
	vm.Exec()
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
	// 10
}
