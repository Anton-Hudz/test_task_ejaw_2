package main

import (
	"crypto/rand"
	"testing"
)

const (
	numString    = 100
	stringLength = 10
)

func BenchmarkConcatination(b *testing.B) {
	data := make([]string, 0, numString)

	for i := 0; i < numString; i++ {
		randomBytes := make([]byte, stringLength)
		rand.Read(randomBytes)
		data = append(data, string(randomBytes))
	}

	b.Run("Concat", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			concat(data)
		}
	})

	b.Run("ConcatStrinBuilder", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			concatStringBuilder(data)
		}
	})

	b.Run("ConcatJoin", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			concatWithJoin(data)
		}
	})

	b.Run("ConcatBuffer", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			concatWithBuffer(data)
		}
	})
}
