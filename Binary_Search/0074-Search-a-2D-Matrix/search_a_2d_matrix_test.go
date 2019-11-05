package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// funcName for easier name representation
func funcName(fnc func([][]int,int) bool) string {
	var ptr = reflect.ValueOf(fnc).Pointer()
	var path = runtime.FuncForPC(ptr).Name()
	var names1 = strings.Split(path, "/")
	var names2 = strings.Split(names1[len(names1)-1], ".")
	return names2[1]
}

func test(t *testing.T, fnc func([][]int,int) bool) {

	var name = funcName(fnc)

	t.Parallel()
	for i, test := range TestCases {
		t.Run(fmt.Sprintf("%s(%d)", name, i), func(t *testing.T) {
			assert.Equal(t, test.expected, fnc(test.matrix, test.target))
		})
	}
}

// prevention of inline optimization.
var result bool

func bench(b *testing.B, fnc func([][]int,int) bool) {
	b.StopTimer()
	b.ReportAllocs()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var r bool
			for _, test := range TestCases {
				r = fnc(test.matrix, test.target)
			}
			result = r
		}
	})
}

// Actual Benchmarks & Tests -------------------------------------

func TestSearchMatrix(t *testing.T) {
	test(t, searchMatrix)
}

func BenchmarkSearchMatrix(b *testing.B) {
	bench(b, searchMatrix)
}
