package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lcgodemo
#include "cgodemo.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	data := make([]DemoData, 9, 9)
	C.demo((*C.int)(unsafe.Pointer(&data[0])))
	fmt.Println(data)

	dataLong := make([]DemoData, 9, 9)
	C.demoLong((*C.long)(unsafe.Pointer(&dataLong[0])))
	fmt.Println(dataLong)

	fmt.Println("----------------------------------------")
	fmt.Printf("int32 size: %d\n", unsafe.Sizeof(int32(0)))
	fmt.Printf("uint64 size: %d\n", unsafe.Sizeof(uint64(0)))
	fmt.Printf("data sizeof: %d aligonf: %d\n", unsafe.Sizeof(data[0]), unsafe.Alignof(data[0]))
	fmt.Printf("dataLong sizeof: %d aligonf: %d\n", unsafe.Sizeof(dataLong[0]), unsafe.Alignof(dataLong[0]))
}

type DemoData struct {
	nIndex  int32
	nId0    uint64
	nLevel0 int32
	nId1    uint64
	nLevel1 int32
	nId2    uint64
	nLevel2 int32
}
