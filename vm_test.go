package vm

func ExampleVM_halt() {
	hello := []int{
		HALT}
	vm := NewVM(hello, 0, 0)
	vm.SetTrace(true)
	vm.cpu()
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
	vm.cpu()
	// Output:
	// 0000: iconst 99 []
	// 0002: print [99]
	// 99
	// 0003: halt []
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
	vm.cpu()
	// Output:
	// 0000: iconst 99 []
	// 0002: gstore 0 [99]
	// 0004: gload 0 []
	// 0006: print [99]
	// 99
	// 0007: halt []
	// Memory:
	// 0000: 99
	// 0001: 0
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
	vm.cpu()
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
