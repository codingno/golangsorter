package golangsorter

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkBubbleSort(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	arr := ArrTest
	rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	for i := 0; i < b.N; i++ {
		BubbleSort(arr)
	}
}
