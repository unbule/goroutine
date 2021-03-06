package main

import (
	"runtime"
	"testing"
)

func init() {
	runtime.GOMAXPROCS(8)
}
func BenchmarkFindfileBygo(b *testing.B) {
	files := directoryList()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findfileBygo(files)
	}
}
func BenchmarkFindfileBygoGroup(b *testing.B) {
	files := directoryList()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findfileBygoGroup(files)
	}
}

func BenchmarkFindfileBynomal(b *testing.B) {
	files := directoryList()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findfileBynomal(files)
	}
}
func BenchmarkFindfileBygotask(b *testing.B) {
	files := directoryList()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findfileBygotask(files)
	}
}
