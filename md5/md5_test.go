package md5

import (
	"testing"
	"math/rand"
)

type randomDataMaker struct {
	src rand.Source
}

func (r *randomDataMaker) Read(p []byte) (n int, err error) {
    todo := len(p)
    offset := 0
    for {
        val := int64(r.src.Int63())
        for i := 0; i < 8; i++ {
            p[offset] = byte(val & 0xff)
            todo--
            if todo == 0 {
                return len(p), nil
            }
            offset++
            val >>= 8
        }
    }
    panic("unreachable")
}

/*
func BenchmarkRandomDataMaker(b *testing.B) {
    randomSrc := randomDataMaker{rand.NewSource(1028890720402726901)}
    for i := 0; i < b.N; i++ {
        b.SetBytes(int64(i))
        _, err := io.CopyN(ioutil.Discard, &randomSrc, int64(i))
        if err != nil {
            b.Fatalf("Error copying at %v: %v", i, err)
        }
    }
}
*/

func BenchmarkMd5(b *testing.B) {
	dig := new(digest)
	dig.Reset()
    randomSrc := randomDataMaker{rand.NewSource(1028890720402726901)}
	buffer := make([]byte, chunk)
	for i := 0; i < b.N; i++ {
		// Copy n bytes from random source.
		randomSrc.Read(buffer)
		// Call block.
		block(dig, buffer)
	}
}
