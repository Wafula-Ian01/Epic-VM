package global

// Contains instruction set definitions
//The opcodes here only serve as
//identifiers for some low-level operations because they
//are abstracted by GO

/*
//data transfer
#define LBI 0 //load byte immediately
#define LWI 1 //load word immediately
#define LDI 2 //load double word immediately
#define LQI 3 //load quad word immediately
#define LF1I 4 //load single-precision float immediately
#define LF2I 5 //load double-precision float immediately

#define LAD 6 //load address direct
#define LAI 7 //load address indirect

#define LB 8 //load byte
#define LW 9 //load word
#define LD 10 //load double word
#define LQ 11 //load quad word
#define LF1 12 //load single-precision float
#define LF2 13 //load double-precision float

#define SB 14 //store byte
#define SW 15 //store word
#define SD 16 //store double word
#define SQ 17 //store quad word
#define SF1 18 //store single-precision float
#define SF2 19 //store double-precision float

#define PUSHB 20 //push byte onto the stack
#define PUSHW 21 //push word onto the stack
#define PUSHD 22 //push double word onto the stack
#define PUSHQ 23 //push quad word onto the stack
#define PUSHF1 24 //push single-precision float onto the stack
#define PUSHF2 25 //push double-precision float onto the stack

#define POPB 26 //pop byte from the stack
#define POPW 27 //pop word from the stack
#define POPD 28 //pop double word from the stack
#define POPQ 29 //pop quad word from the stack
#define POPF1 30 //pop single-precision float from the stack
#define POPF2 31 //pop double-precision float from the stack

#define MOV 32 //move an integer
#define MOVF 33 //move a single-precision value
#define MOVD 34 //move a double-precision value

//program flow control
#define JMP 35 //Unconditional jump
#define JE 36 //jump if equal
#define JNE 37 //jump if not equal
#define SLT 38 //set less than
#define INT 39 //perform interrupt
#define DI 40 //disable interrupt
#define EI 41 //enable interrupt
#define HALT 42 //stop virtual machine
#define NOP 43 //No Operation

//Bitwise operators
#define AND 44 //bitwise AND
#define OR 45 //bitwise OR
#define XOR 46 //bitwise XOR
#define NOT 47 //bitwise NOT
#define BT 48 //bitwise Test
#define BS 49 //bitwise Set

//Shift
#define SRA 50 //shift arithmetic rigt
#define SRL 51 //shift logic right
#define SL 52 //shift left

//Integer Aritmetic
#define ADD 53 //integer addition
#define SUB 54 //integer subtraction
#define MULT 55 //integer multiplication
#define DIV 56 //integer division

//
#define CAST_IF 57 //convert a single-precision float to an integer
#define CAST_ID 58 //convert a double-precision float to an integer
#define CAST_FI 59 //convert an integer to a single-precision float
#define CAST_FD 60 //convert a double-precision float to a single precision float
#define CAST_DI 61 //convert an integer to a double-precision float
#define CAST_DF 62 //convert a single-precision float to a double-precision float
*/
import "C"
