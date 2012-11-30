package main
import (
	"testing"
	"fmt"
)
func BenchmarkT1(b *testing.B) {
	a:=0
	for i:=0;i<b.N;i++{
		a=1
		fmt.Print(a)
	}
	
}