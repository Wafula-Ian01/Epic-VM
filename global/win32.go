package global

import (
	"encoding/binary"
	"fmt"
)

//Contains Microsoft Windows specific code

// S1 signed data datatypes
type S1 = byte  //represents a char
type S2 = int16 //represents a short
type S4 = int64 //represents a long

// U1 unsigned data datatypes
type U1 = byte  //represents a char
type U2 = int16 //represents a short
type U4 = int64 //represents a long

// F4 float and double
type F4 = float32 //represents float
type F8 = float64 //represents double

func checkEndian() { //Checks for the endianness of the platform running the VM
	var hexValue uint32 = 0xDEED1234            //declare a hex value
	bytes := make([]byte, 4)                    //create a splice of bytes
	binary.BigEndian.PutUint32(bytes, hexValue) //insert the hex value into the splice
	var j int

	fmt.Printf("Value= %v\n", hexValue)
	if binary.BigEndian.Uint32(bytes) == hexValue { //check for big endianness
		fmt.Printf("this Platform is Little Endianness. Lower bytes come first.")
	} else if binary.LittleEndian.Uint32(bytes) == hexValue { //check for small endianness
		fmt.Printf("this platform is Big Endianness. Upper bytes come first")
	} else { //endianness unknown
		fmt.Printf("This platform's Endianness is unknown.")
	}

	//return the bytes in our splice
	fmt.Printf("\nhere are the 4 bytes\n")

	for j = 0; j < 4; j++ {
		fmt.Printf("bytes[%d]= %x", j, bytes)
	}
	return
} //End of checking Endianness

// converting big endianness to little endianness
func byteCodeToWord(bytes []byte) byte {
	var word U2
	var buffer *U1
	buffer = (*U1) & word
}

//converting little endianness to big endianness
