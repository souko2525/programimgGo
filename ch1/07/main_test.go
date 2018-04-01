//go test -bench=. -benchmem

package main

import "testing"

//実行しちゃだめよ
func BenchmarkBuf1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf1()
	}
}

func BenchmarkBuf2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf2()
	}
}
