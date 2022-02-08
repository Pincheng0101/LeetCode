package benchmarks

import (
	"testing"
	"unsafe"
)

func BenchmarkConvertBytesToString(b *testing.B) {
	var bytes []byte = []byte("XXXXXXXXXX")
	for i := 0; i < b.N; i++ {
		_ = string(bytes)
	}
}

func BenchmarkConvertBytesToStringWithUnsafePointer(b *testing.B) {
	var bytes []byte = []byte("XXXXXXXXXX")
	for i := 0; i < b.N; i++ {
		p := unsafe.Pointer(&bytes)
		_ = *(*string)(p)
	}
}

func BenchmarkConvertStringToBytes(b *testing.B) {
	var s string = "XXXXXXXXXX"
	for i := 0; i < b.N; i++ {
		_ = []byte(s)
	}
}

func BenchmarkConvertStringToBytesWithUnsafePointer(b *testing.B) {
	var s string = "XXXXXXXXXX"
	for i := 0; i < b.N; i++ {
		p := unsafe.Pointer(&s)
		_ = *(*[]byte)(unsafe.Pointer(&p)) // dangerous!
	}
}

func TestConvertStringToBytes(t *testing.T) {
	s := "Test"
	bytes := []byte(s)
	bytes[0] = 'X'
}

func TestConvertStringToBytesWithUnsafePointer(t *testing.T) {
	s := "Test"
	bytes := *(*[]byte)(unsafe.Pointer(&s))
	bytes[0] = 'X'
}
