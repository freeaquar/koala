package ch

import (
	"reflect"
	"fmt"
	"unsafe"
	"encoding/binary"
	"net"
)

func Dump(typ reflect.Type) {
	fmt.Println("=== " + typ.String() + " ===")
	for i := 0; i < typ.NumField(); i++ {
		fmt.Println(typ.Field(i).Name, typ.Field(i).Type.String(), typ.Field(i).Type.Kind())
	}
}

func FieldOf(typ reflect.Type, name string) *reflect.StructField {
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.Name == name {
			return &field
		}
	}
	panic(name + " not found in " + typ.String())
}

func GetUint16(ptr unsafe.Pointer, field *reflect.StructField) uint16 {
	if field.Type.Kind() != reflect.Uint16 {
		panic("kind mismatch")
	}
	fieldPtr := unsafe.Pointer(uintptr(ptr) + field.Offset)
	return *(*uint16)(fieldPtr)
}

func GetUint32(ptr unsafe.Pointer, field *reflect.StructField) uint32 {
	if field.Type.Kind() != reflect.Uint32 {
		panic("kind mismatch")
	}
	fieldPtr := unsafe.Pointer(uintptr(ptr) + field.Offset)
	return *(*uint32)(fieldPtr)
}

func GetPtr(ptr unsafe.Pointer, field *reflect.StructField) unsafe.Pointer {
	return unsafe.Pointer(uintptr(ptr) + field.Offset)
}

func Ntohl(i uint32) uint32 {
	return binary.BigEndian.Uint32((*(*[4]byte)(unsafe.Pointer(&i)))[:])
}
func Htonl(i uint32) uint32 {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, i)
	return *(*uint32)(unsafe.Pointer(&b[0]))
}

func Ntohs(i uint16) uint16 {
	return binary.BigEndian.Uint16((*(*[2]byte)(unsafe.Pointer(&i)))[:])
}
func Htons(i uint16) uint16 {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, i)
	return *(*uint16)(unsafe.Pointer(&b[0]))
}

func Int2ip(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.LittleEndian.PutUint32(ip, nn)
	return ip
}
