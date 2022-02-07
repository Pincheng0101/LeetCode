package benchmarks

import (
	"bytes"
	"strings"
	"testing"
)

func BenchmarkNormalConcat(b *testing.B) {
	var s string = ""
	for i := 0; i < b.N; i++ {
		s += "X"
	}
	_ = s
}

func BenchmarkBytesBuffer(b *testing.B) {
	var s bytes.Buffer
	for i := 0; i < b.N; i++ {
		s.WriteString("X")
	}
	_ = s.String()
}

func BenchmarkStringsBuider(b *testing.B) {
	var s strings.Builder
	for i := 0; i < b.N; i++ {
		s.WriteString("X")
	}
	_ = s.String()
}

func BenchmarkStringsBuiderWithGrow(b *testing.B) {
	var s strings.Builder
	s.Grow(b.N)
	for i := 0; i < b.N; i++ {
		s.WriteString("X")
	}
	_ = s.String()
}

func BenchmarkCopy(b *testing.B) {
	var bytes []byte = make([]byte, b.N)
	for i := 0; i < b.N; i++ {
		copy(bytes[i:], "X")
	}
	_ = string(bytes)
}

func BenchmarkAssign(b *testing.B) {
	var bytes []byte = make([]byte, b.N)
	for i := 0; i < b.N; i++ {
		bytes[i] = byte('X')
	}
	_ = string(bytes)
}

func BenchmarkNormalConcatLongString(b *testing.B) {
	var s string = ""
	for i := 0; i < b.N; i++ {
		s += "XXXXXXXXXX"
	}
	_ = s
}

func BenchmarkBytesBufferLongString(b *testing.B) {
	var s bytes.Buffer
	for i := 0; i < b.N; i++ {
		s.WriteString("XXXXXXXXXX")
	}
	_ = s.String()
}

func BenchmarkStringsBuiderLongString(b *testing.B) {
	var s strings.Builder
	for i := 0; i < b.N; i++ {
		s.WriteString("XXXXXXXXXX")
	}
	_ = s.String()
}
