package cosmo

import (
	"testing"
)

func benchmarkFlatLCDMEN(n int, b *testing.B) {
	cos := FlatLCDM{H0: 70, Om0: 0.27}
	var z float64
	zMax := 1.0
	step := zMax / float64(n)
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			z = 0.001 + step*float64(j)
			cos.E(z)
		}
	}
}

func BenchmarkFlatLCDMEN(b *testing.B) {
	benchmarkFlatLCDMEN(10000, b)
}

func BenchmarkFlatLCDMENdistance(b *testing.B) {
	cos := FlatLCDM{H0: 70, Om0: 0.27}
	benchmarkFlatLCDMNdistance(10000, cos.E, b)
}

func BenchmarkFlatLCDME(b *testing.B) {
	cos := FlatLCDM{H0: 70, Om0: 0.27}
	z := 1.0
	for i := 0; i < b.N; i++ {
		cos.E(z)
	}
}

func BenchmarkFlatLCDMEinv(b *testing.B) {
	cos := FlatLCDM{H0: 70, Om0: 0.27}
	z := 1.0
	for i := 0; i < b.N; i++ {
		cos.Einv(z)
	}
}

// benchmarkFlatLCDMDistance is a helper function to be called by specific benchmarkFlatLCDMs
func benchmarkFlatLCDMDistance(f func(float64) float64, b *testing.B) {
	z := 1.0
	for i := 0; i < b.N; i++ {
		f(z)
	}
}

// benchmarkFlatLCDMNdistance is a helper function to be called by specific benchmarkFlatLCDMs
func benchmarkFlatLCDMNdistance(n int, f func(float64) float64, b *testing.B) {
	var z float64
	zMax := 1.0
	step := zMax / float64(n)
	for i := 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			z = 0.001 + step*float64(j)
			f(z)
		}
	}
}

func BenchmarkFlatLCDMComovingDistance(b *testing.B) {
	cos := FlatLCDM{H0: 70, Om0: 0.27}
	benchmarkFlatLCDMDistance(cos.ComovingDistance, b)
}

func BenchmarkFlatLCDMComovingTransverseDistance(b *testing.B) {
	cos := FlatLCDM{H0: 70, Om0: 0.27}
	benchmarkFlatLCDMDistance(cos.ComovingTransverseDistance, b)
}

func BenchmarkFlatLCDMLuminosityDistance(b *testing.B) {
	cos := FlatLCDM{H0: 70, Om0: 0.27}
	benchmarkFlatLCDMDistance(cos.LuminosityDistance, b)
}

func BenchmarkFlatLCDMLookbackTime(b *testing.B) {
	cos := FlatLCDM{H0: 70, Om0: 0.27}
	benchmarkFlatLCDMDistance(cos.LookbackTime, b)
}

func BenchmarkFlatLCDMNComovingDistance(b *testing.B) {
	cos := FlatLCDM{H0: 70, Om0: 0.27}
	benchmarkFlatLCDMNdistance(10000, cos.ComovingDistance, b)
}

func BenchmarkFlatLCDMNLuminosityDistance(b *testing.B) {
	cos := FlatLCDM{H0: 70, Om0: 0.27}
	benchmarkFlatLCDMNdistance(10000, cos.LuminosityDistance, b)
}

func BenchmarkFlatLCDMNE(b *testing.B) {
	cos := FlatLCDM{H0: 70, Om0: 0.27}
	benchmarkFlatLCDMNdistance(10000, cos.E, b)
}
