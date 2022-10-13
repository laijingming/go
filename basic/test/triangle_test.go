package test

import (
	"testing"
)

func TestTriangle(t *testing.T) {
	testes := []struct{ a, b, c int }{
		{3, 4, 5},
		{8, 12, 14},
		{3000, 4000, 5000},
	}
	for _, tt := range testes {
		if actual := calcTriangle(tt.a, tt.b); actual != tt.c {
			t.Errorf("calcTriangle(%d,%d) got %d,expected %d", tt.a, tt.b, actual, tt.c)
		}
	}
}

//func BenchmarkTriangle(b *testing.B) {
//	a := 3
//	bb := 4
//	c := 5
//	for i := 0; i < b.N; i++ {
//		if actual := calcTriangle(a, bb); actual != c {
//			b.Errorf("calcTriangle(%d,%d) got %d,expected %d", a, bb, actual, c)
//		}
//	}
//}
//
//func BenchmarkLengthOfNonRepeatingSubStr(b *testing.B) {
//	a := "sssa测试数"
//	c := 5
//	for i := 0; i < b.N; i++ {
//		if actual := lengthOfNonRepeatingSubStr(a); actual != c {
//			b.Errorf("lengthOfNonRepeatingSubStr(%s) got %d,expected %d", a, actual, c)
//		}
//	}
//}
func BenchmarkLengthOfNonRepeatingSubStr2(b *testing.B) {
	a := "sssa测试数"
	c := 5
	for i := 0; i < b.N; i++ {
		if actual := lengthOfNonRepeatingSubStr2(a); actual != c {
			b.Errorf("lengthOfNonRepeatingSubStr(%s) got %d,expected %d", a, actual, c)
		}
	}
}