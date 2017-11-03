package utils

import (
	"testing"
	. "github.com/franela/goblin"
	"math/rand"
)

func TestEncodeDecode(t *testing.T) {
	g := Goblin(t)
	g.Describe("Utils: Encode & Decode", func() {
		g.It("Should encode an int ID", func() {
			g.Assert(Encode(12345)).Equal("3wwsx")
		})
		g.It("Should decode a string", func() {
			g.Assert(Decode("3wwsx")).Equal(12345)
		})
	})
}

func TestGetURL(t *testing.T) {
	g := Goblin(t)
	g.Describe("Utils: GetURL", func() {
		g.It("Should get the correct URL if protocol identifier is absent", func() {
			g.Assert(GetURL("example.com", "SomeCode")).Equal("https://example.com/SomeCode")
		})
		g.It("Should get the correct URL if protocol identifier is present", func() {
			g.Assert(GetURL("https://example.com", "SomeCode")).Equal("https://example.com/SomeCode")
		})
	})
}

func BenchmarkDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Decode("3wwsx")
	}
}

func BenchmarkEncode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Encode(rand.Intn(100000) + 1)
	}
}
