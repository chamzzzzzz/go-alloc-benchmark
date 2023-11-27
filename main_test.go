package main

import (
	"testing"
)

func BenchmarkGetTestDataValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetTestDataValue()
	}
}

func BenchmarkGetTestDataPtr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetTestDataPtr()
	}
}

func BenchmarkGetTestDataValueNoInline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetTestDataValueNoInline()
	}
}

func BenchmarkGetTestDataPtrNoInline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetTestDataPtrNoInline()
	}
}
