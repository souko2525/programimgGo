package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	longString = "longlonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglonglong"
	sepapator  = "\n"
)

var byteString []string

func init() {
	byteString = strings.Split(longString, "")
}

func measureTime(f func() string) int64 {
	start := time.Now()
	f()
	return time.Since(start).Nanoseconds()
}

func echo2() string {
	var joined string
	for _, c := range byteString {
		joined += c
		joined += sepapator
	}
	return joined
}
func echo3() string {
	return strings.Join(byteString, sepapator)
}

func main() {
	slow := measureTime(echo2)
	fast := measureTime(echo3)
	fmt.Printf("slow version %v nsec\n", slow)
	fmt.Printf("fast version %v nsec\n", fast)
}
