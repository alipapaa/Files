package main

import "testing"

func BenchmarkComma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ttcomma()
	}
}

func BenchmarkCommaBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ttcommabuffer()
	}
}

func ttcomma() {
	t := []string{"1111", "1111111", "111111"}
	for i := 0; i < len(t); i++ {
		comma(t[i])
	}
}

func ttcommabuffer() {
	t := []string{"1111", "1111111", "111111"}
	for i := 0; i < len(t); i++ {
		commaBuffer([]byte(t[i]))
	}
}

//练习3.10 ThomasHuke$ go test -bench=. -cpu=4
//goos: darwin
//goarch: amd64
//pkg: github.com/googege/gopl_homework/ch3/第三章练习题/练习3.10
//BenchmarkComma-4         	 5000000	       250 ns/op
//BenchmarkCommaBuffer-4   	 3000000	       456 ns/op
//PASS
//ok  	github.com/googege/gopl_homework/ch3/第三章练习题/练习3.10	3.373s
