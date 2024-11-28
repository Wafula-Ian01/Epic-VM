//go:build windows
// +build windows

//Added this because my memory access method wasn't being recognised since I'm using linux.
//And this is a windows specific file

package global

/*
#include <stdint.h>

// S1 signed data datatypes
#define S1 signed char
#define S2 signed short
#define S4 signed long
#define S8 signed __int64

// U1 unsigned data datatypes
#define U1 unsigned char
#define U2 unsigned short
#define U4 unsigned long
#define U8 unsigned __int64

// F4 float and double
#define F4 float
#define F8 double

#define PRINT_UREG(rstr, reg) printf("%-6s=%-21I64u", rstr, reg)
#define PRINT_SREG(rstr, reg) printf("%-6s=%-21I64d", rstr, reg)
#define PRINT_FREG(rstr, reg) printf("%-6s=%g", rstr, (F4)reg)
#define PRINT_DREG(rstr, reg) printf("%-6s=%g", rstr, (F8)reg)

#define pU1(arg) printf("%u", arg)
#define pU2(arg) printf("%hu", arg)
#define pU4(arg) printf("%lu", arg)
#define pU8(arg) printf("%I64lu", arg)

#define pS1(arg) printf("%d", arg)
#define pS2(arg) printf("%hd", arg)
#define pS4(arg) printf("%ld", arg)
#define pS8(arg) printf("%I64d", arg)

#define rU8(arg) scanf("%I64ld", arg)

#define fpU8(ptr, arg) fprintf(ptr, "%I64lu", arg)
#define fpS8(ptr, arg) fprintf(ptr, "%I64d", arg)

*/
import "C"

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"golang.org/x/sys/windows"
	"math"
	"os"
	"unsafe"
)

//Contains Microsoft Windows specific code

func checkEndian() { //Checks for the endianness of the platform running the VM
	var hexValue uint32 = 0xDEED1234               //declare a hex value
	bytesArr := make([]byte, 4)                    //create a splice of bytes
	binary.BigEndian.PutUint32(bytesArr, hexValue) //insert the hex value into the splice
	var j int

	fmt.Printf("Value= %v\n", hexValue)
	if binary.BigEndian.Uint32(bytesArr) == hexValue { //check for big endianness
		fmt.Printf("this Platform is Little Endianness. Lower bytes come first.")
	} else if binary.LittleEndian.Uint32(bytesArr) == hexValue { //check for small endianness
		fmt.Printf("this platform is Big Endianness. Upper bytes come first")
	} else { //endianness unknown
		fmt.Printf("This platform's Endianness is unknown.")
	}

	//return the bytes in our splice
	fmt.Printf("\nhere are the 4 bytes\n")

	for j = 0; j < 4; j++ {
		fmt.Printf("bytes[%d]= %x", j, bytesArr)
	}
	return
} //End of checking Endianness

func byteCodeToWord(bytes []C.U1) C.U2 {
	if len(bytes) != 2 {
		panic("Input slice must contain 2 bytes")
	}
	word := C.U2(bytes[1]) | C.U2(bytes[0])<<8

	return word
}

func byteCodeToDword(bytes []C.U1) C.U4 {
	if len(bytes) != 4 {
		panic("Input slice must contain 4 bytes")
	}

	dword := C.U4(bytes[3]) | C.U4(bytes[2])<<8 | C.U4(bytes[1])<<16 | C.U4(bytes[0])<<24

	return dword
}

func byteCodeToQword(bytes []C.U1) C.U8 {
	if len(bytes) != 8 {
		panic("Input slice must contain 8 bytes")
	}

	qword := C.U8(bytes[7]) | C.U8(bytes[6])<<8 | C.U8(bytes[5])<<16 | C.U8(bytes[4])<<24 | C.U8(bytes[3])<<32 |
		C.U8(bytes[2])<<40 | C.U8(bytes[1])<<48 | C.U8(bytes[0])<<56

	return qword
}

func byteCodeToFloat(bytes []C.U1) C.F4 {
	if len(bytes) != 4 {
		panic("Input slice must contain 4 bytes")
	}
	//reverse splice
	reversedBytes := []byte{bytes[3], bytes[2], bytes[1], bytes[0]}

	//convert bytes to uint32
	bits := binary.LittleEndian.Uint32(reversedBytes)
	//convert to float32
	floatValue := math.Float32frombits(bits)

	return floatValue
}

func byteCodeToDouble(bytes []C.U1) C.F8 {
	if len(bytes) != 8 {
		panic("Input slice must contain 8 bytes")
	}

	reverseBytes := []byte{bytes[7], bytes[6], bytes[5], bytes[4],
		bytes[3], bytes[2], bytes[1], bytes[0]}

	bits := binary.LittleEndian.Uint64(reverseBytes)

	doubleValue := math.Float64frombits(bits)

	return doubleValue
}
func wordToByteCode(word C.U2) []C.U1 {
	arr := make([]C.U1, 2)

	arr[0] = C.U1(word >> 8)
	arr[1] = C.U1(word & 0xFF)

	return arr
}

func dwordToByteCode(dword C.U4) []C.U1 {
	arr := make([]C.U1, 4)

	arr[0] = C.U1(dword & 0xFF)
	arr[1] = C.U1((dword >> 8) & 0xFF)
	arr[2] = C.U1((dword >> 16) & 0xFF)
	arr[3] = C.U1((dword >> 24) & 0xFF)

	return arr
}

func qwordToByteCode(qword C.U8) []C.U1 {
	arr := make([]C.U1, 8)

	arr[0] = C.U1(qword & 0xFF)
	arr[1] = C.U1((qword >> 8) & 0xFF)
	arr[2] = C.U1((qword >> 16) & 0xFF)
	arr[3] = C.U1((qword >> 24) & 0xFF)
	arr[4] = C.U1((qword >> 32) & 0xFF)
	arr[5] = C.U1((qword >> 40) & 0xFF)
	arr[6] = C.U1((qword >> 48) & 0xFF)
	arr[7] = C.U1((qword >> 56) & 0xFF)

	return arr
}

func floatToByteCode(floatValue C.F4) []C.U1 {
	bits := math.Float32bits(floatValue)

	buff := new(bytes.Buffer)

	err := binary.Write(buff, binary.LittleEndian, bits)

	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	return buff.Bytes()
}

func doubleToByteCode(doubleValue C.F8) []C.U1 {
	bits := math.Float64bits(doubleValue)

	buff := new(bytes.Buffer) // holds the byte sequence

	err := binary.Write(buff, binary.LittleEndian, bits)

	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	return buff.Bytes()
}

// byte swapping for endian conversions
func formatWord(arr []C.U1, start int) {
	fb0, fb1 := arr[start+1], arr[start]
	arr[start], arr[start+1] = fb0, fb1
}

func formatDword(arr []C.U1, start int) {
	fb0, fb1, fb2, fb3 := arr[start+3], arr[start+2], arr[start+1], arr[start]
	arr[start], arr[start+1], arr[start+2], arr[start+3] = fb0, fb1, fb2, fb3
}

func formatQword(arr []C.U1, start int) {
	fb0, fb1, fb2, fb3, fb4, fb5, fb6, fb7 :=
		arr[start+7], arr[start+6], arr[start+5], arr[start+4],
		arr[start+3], arr[start+2], arr[start+1], arr[start]
	arr[start], arr[start+1], arr[start+2], arr[start+3],
		arr[start+4], arr[start+5], arr[start+6], arr[start+7] =
		fb0, fb1, fb2, fb3, fb4, fb5, fb6, fb7
}

// MEMORYSTATUSEX Obtaining memory stats and file details
type MEMORYSTATUSEX struct {
	Length               uint32
	MemoryLoad           uint32
	TotalPhys            uint64
	AvailPhys            uint64
	TotalPageFile        uint64
	AvailPageFile        uint64
	TotalVirtual         uint64
	AvailVirtual         uint64
	AvailExtendedVirtual uint64
}

func getAvailableMemory() (uint64, uint64) {
	msx := MEMORYSTATUSEX{
		Length: uint32(unsafe.Sizeof(MEMORYSTATUSEX{})),
	}

	if err := windows.GlobalMemoryStatusEx(msx); err != nil {
		fmt.Printf("Failed to get memory status: %v", err)
		return 0, 0
	}
	return msx.TotalPhys, msx.AvailPhys
} /*end getAvailableMemory*/

// ensures that bytecode executable loaded to memory ain't bigger than the memory its being loaded to
func getFileSize(filePath *string, err error) U4 {
	if filePath == nil {
		return 0, fmt.Errorf("filePath is nil.")
	}

	fileInfo, err := os.Open(*filePath)
	if err != nil {
		return 0, err
	}

	return fileInfo.Size(), nil
}
