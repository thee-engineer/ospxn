package ospxn_test

import (
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"testing"

	"github.com/thee-engineer/ospxn"
)

func TestDummy(t *testing.T) {
	noise := ospxn.New(0)

	noise.Eval2(0.0, 0.0)
	noise.Eval3(0.0, 0.0, 0.0)

	// norm := ospxn.NewNormalized(0)
	// norm.Eval2(0.0, 0.0)
	// norm.Eval3(0.0, 0.0, 0.0)
}

func BenchmarkEval2(b *testing.B) {
	f, err := os.Create("eval2.profile")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	noise := ospxn.New(1)
	for n := 0; n < b.N; n++ {
		noise.Eval2(float64(n), float64(n))
	}
}

func BenchmarkEval3(b *testing.B) {
	f, err := os.Create("eval3.profile")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	noise := ospxn.New(1)
	for n := 0; n < b.N; n++ {
		noise.Eval3(float64(n), float64(n), float64(n))
	}
}

func Example() {
	noise := ospxn.New(rand.Int63())

	w, h := 100, 100
	heightmap := make([]float64, w, h)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			xFloat := float64(x) / float64(w)
			yFloat := float64(y) / float64(h)
			heightmap[(y*w)+x] = noise.Eval2(xFloat, yFloat)
		}
	}
}
