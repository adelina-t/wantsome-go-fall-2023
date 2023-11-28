package bench

import (
	"fmt"
	"testing"
)

var blackhole int

func Test_GetFileLen(t *testing.T) {
	result, err := GetFileLen("testdata/data.txt", 1)
	if err != nil {
		t.Fatal(err)
	}

	if result != 960 {
		t.Errorf("expected 960, got %d", result)
	}
}

func Benchmark_GetFileLen(b *testing.B) {
	for _, bufSize := range []int{1, 10, 100, 1000, 100000} {
		b.Run(fmt.Sprintf("bufSize-%d", bufSize), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result, err := GetFileLen("testdata/data.txt", bufSize)
				if err != nil {
					b.Fatal(err)
				}
				blackhole = result
			}
		})
	}
}
