//go:build windows
// +build windows

//Added this because my memory access method wasn't being recognised since I'm using linux.

package global

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

// S1 signed data datatypes
type S1 = byte  //represents a char
type S2 = int16 //represents a short
type S4 = int32 //represents a long
type S8 = int64 //represents long long __int64

// U1 unsigned data datatypes
type U1 = byte   //represents a char
type U2 = uint16 //represents a short
type U4 = uint32 //represents a long
type U8 = uint64 //represents long long or __uint64

// F4 float and double
type F4 = float32 //represents float
type F8 = float64 //represents double

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

func byteCodeToWord(bytes []U1) U2 {
	if len(bytes) != 2 {
		panic("Input slice must contain 2 bytes")
	}
	word := U2(bytes[1]) | U2(bytes[0])<<8

	return word

}

func byteCodeToDword(bytes []U1) U4 {
	if len(bytes) != 4 {
		panic("Input slice must contain 4 bytes")
	}

	dword := U4(bytes[3]) | U4(bytes[2])<<8 | U4(bytes[1])<<16 | U4(bytes[0])<<24

	return dword
}

func byteCodeToQword(bytes []U1) U8 {
	if len(bytes) != 8 {
		panic("Input slice must contain 8 bytes")
	}

	qword := U8(bytes[7]) | U8(bytes[6])<<8 | U8(bytes[5])<<16 | U8(bytes[4])<<24 | U8(bytes[3])<<32 |
		U8(bytes[2])<<40 | U8(bytes[1])<<48 | U8(bytes[0])<<56

	return qword
}

func byteCodeToFloat(bytes []U1) F4 {
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

func byteCodeToDouble(bytes []U1) F8 {
	if len(bytes) != 8 {
		panic("Input slice must contain 8 bytes")
	}

	reverseBytes := []byte{bytes[7], bytes[6], bytes[5], bytes[4],
		bytes[3], bytes[2], bytes[1], bytes[0]}

	bits := binary.LittleEndian.Uint64(reverseBytes)

	doubleValue := math.Float64frombits(bits)

	return doubleValue
}
func wordToByteCode(word U2) []U1 {
	arr := make([]U1, 2)

	arr[0] = U1(word >> 8)
	arr[1] = U1(word & 0xFF)

	return arr
}

func dwordToByteCode(dword U4) []U1 {
	arr := make([]U1, 4)

	arr[0] = U1(dword & 0xFF)
	arr[1] = U1((dword >> 8) & 0xFF)
	arr[2] = U1((dword >> 16) & 0xFF)
	arr[3] = U1((dword >> 24) & 0xFF)

	return arr
}

func qwordToByteCode(qword U8) []U1 {
	arr := make([]U1, 8)

	arr[0] = U1(qword & 0xFF)
	arr[1] = U1((qword >> 8) & 0xFF)
	arr[2] = U1((qword >> 16) & 0xFF)
	arr[3] = U1((qword >> 24) & 0xFF)
	arr[4] = U1((qword >> 32) & 0xFF)
	arr[5] = U1((qword >> 40) & 0xFF)
	arr[6] = U1((qword >> 48) & 0xFF)
	arr[7] = U1((qword >> 56) & 0xFF)

	return arr
}

func floatToByteCode(floatValue F4) []U1 {
	bits := math.Float32bits(floatValue)

	buff := new(bytes.Buffer)

	err := binary.Write(buff, binary.LittleEndian, bits)

	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	return buff.Bytes()
}

func doubleToByteCode(doubleValue F8) []U1 {
	bits := math.Float64bits(doubleValue)

	buff := new(bytes.Buffer) // holds the byte sequence

	err := binary.Write(buff, binary.LittleEndian, bits)

	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}

	return buff.Bytes()
}

// byte swapping for endian conversions
func formatWord(arr []U1, start int) {
	fb0, fb1 := arr[start+1], arr[start]
	arr[start], arr[start+1] = fb0, fb1
}

func formatDword(arr []U1, start int) {
	fb0, fb1, fb2, fb3 := arr[start+3], arr[start+2], arr[start+1], arr[start]
	arr[start], arr[start+1], arr[start+2], arr[start+3] = fb0, fb1, fb2, fb3
}

func formatQword(arr []U1, start int) {
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
