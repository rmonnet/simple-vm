# simple-vm

A simple Virtual Machine (in go) based on Terence Parr's talk.

## License

The project is published under the MIT license.

## Reference

Terence Parr's talk about building a Virtual Machine can be found
[here](https://www.youtube.com/watch?v=OjaAToVkoTw)
and the associated slides can be found [here](https://www.slideshare.net/parrt/how-to-build-a-virtual-machine).

## Summary

The project demonstrates a subset of a Virtual Machine (integer operations)
as well as test and branch instructions and subroutine call.

Example of simple programs and their output can be found in `vm_test.go`.

## The Virtual Machine

The Virtual Machine memory consists of separate code, stack and global data spaces.
The Virtual Machine, like the JVM, is a stack based interpreter. Arguments are first
pushed on the stack, the bytecode executed and the result left on the stack.

The bytecode documentation below shows the stack, before and after, the bytecode
is executed.

The bytecodes are as follow:

- IADD (l r -- l+r)
    - adds the two integers on top of the stack (leaves result on the stack)
- ISUB (l r -- l-r)
	- substract the integer on top of the stack from the 2nd on the stack (leaves result on the stack)
- IMUL (l -r -- l*r)
    - multiplies the 2 integers on top of the stack (leaves result on the stack)
- ILT (l r -- l<r)
    - checks if the 2nd integer on the stack is less than the one on top of the stack (leaves 0=false or 1=true on top of the stack)
- IEQ (l r -- l==r)
    - checks if the two integers on the top of the stack are equal (leaves 0=false or 1=true on top of the stack)
- BR addr ( -- )
    - jumps to the given address
- BRT addr (b -- )
    - jumps to the given address if the condition on top of the stack is true
- BRF addr (b -- )
    - jumps to the given address if the condition on top of the stack is false
- ICONST val ( -- v)
    - pushes an integer on the stack
- LOAD i ( -- v)
    - loads the ith value from the local context (indicated by the frame pointer) on top of the stack. The ith value is a positive or negative offset from the frame pointer (fp).
    - the local context is setup during a subroutine call. It is stored on the stack as follow:
        - fp+0: previous stack pointer
        - fp-1: previous frame pointer
        - fp-2: subroutine number of arguments
        - fp-3: nth argument of the subroutine (before call)
        - fp-4: (n-1)th argument of the subroutine (before call)
        - fp-2-n: 1st argument of the subroutine (before call)
- STORE i (v -- )
    - stores the top of the stack to the ith value from the local context (see LOAD for local context details)
- GLOAD i ( -- v)
    - push the ith global variable to the top of the stack
- STORE i (v -- )
    - stores the top of the stack as the ith global variable
- PRINT (v -- )
    - prints the top of the stack
- POP (v -- )
    - pops and discards the top of the stack
- CALL addr ( -- )
    - calls the subroutine at the addr
- RET ( -- )
    - returns from a function call
- HALT ( -- )
    - stops the current program

## To do

- missing implementation for the CALL and RET bytecode
- implement an assembler



