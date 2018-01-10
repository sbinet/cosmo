package cosmo

import (
	"testing"
)

func benchmarkWACDMEN(n int, b *testing.B) {
	cos := WACDM{H0: 70, Om0: 0.2, Ol0: 0.7, W0: -1.2, WA: 2}
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

func BenchmarkWACDMEN(b *testing.B) {
	benchmarkWACDMEN(10000, b)
}

func BenchmarkWACDMENdistance(b *testing.B) {
	cos := WACDM{H0: 70, Om0: 0.2, Ol0: 0.7, W0: -1.2}
	benchmarkWACDMNdistance(10000, cos.E, b)
}

func BenchmarkWACDME(b *testing.B) {
	cos := WACDM{H0: 70, Om0: 0.2, Ol0: 0.7, W0: -1.2, WA: 2}
	z := 1.0
	for i := 0; i < b.N; i++ {
		cos.E(z)
	}
}

func BenchmarkWACDMEinv(b *testing.B) {
	cos := WACDM{H0: 70, Om0: 0.2, Ol0: 0.7, W0: -1.2, WA: 2}
	z := 1.0
	for i := 0; i < b.N; i++ {
		cos.Einv(z)
	}
}

// benchmarkWACDMDistance is a helper function to be called by specific benchmarkWACDMs
func benchmarkWACDMDistance(f func(float64) float64, b *testing.B) {
	z := 1.0
	for i := 0; i < b.N; i++ {
		f(z)
	}
}

// benchmarkWACDMNdistance is a helper function to be called by specific benchmarkWACDMs
func benchmarkWACDMNdistance(n int, f func(float64) float64, b *testing.B) {
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

func BenchmarkWACDMComovingDistance(b *testing.B) {
	cos := WACDM{H0: 70, Om0: 0.2, Ol0: 0.7, W0: -1.2}
	benchmarkWACDMDistance(cos.ComovingDistance, b)
}

func BenchmarkWACDMComovingTransverseDistance(b *testing.B) {
	cos := WACDM{H0: 70, Om0: 0.2, Ol0: 0.7, W0: -1.2}
	benchmarkWACDMDistance(cos.ComovingTransverseDistance, b)
}

func BenchmarkWACDMLuminosityDistance(b *testing.B) {
	cos := WACDM{H0: 70, Om0: 0.2, Ol0: 0.7, W0: -1.2}
	benchmarkWACDMDistance(cos.LuminosityDistance, b)
}

func BenchmarkWACDMLookbackTime(b *testing.B) {
	cos := WACDM{H0: 70, Om0: 0.2, Ol0: 0.7, W0: -1.2}
	benchmarkWACDMDistance(cos.LookbackTime, b)
}

func BenchmarkWACDMNComovingDistance(b *testing.B) {
	cos := WACDM{H0: 70, Om0: 0.2, Ol0: 0.7, W0: -1.2}
	benchmarkWACDMNdistance(10000, cos.ComovingDistance, b)
}

func BenchmarkWACDMNLuminosityDistance(b *testing.B) {
	cos := WACDM{H0: 70, Om0: 0.2, Ol0: 0.7, W0: -1.2}
	benchmarkWACDMNdistance(10000, cos.LuminosityDistance, b)
}

func BenchmarkWACDMNE(b *testing.B) {
	cos := WACDM{H0: 70, Om0: 0.2, Ol0: 0.7, W0: -1.2}
	benchmarkWACDMNdistance(10000, cos.E, b)
}

func BenchmarkWACDMComovingDistanceOM(b *testing.B) {
	cos := WACDM{H0: 70, Om0: 0.3, Ol0: 0, W0: -1.2, WA: 2}
	benchmarkWACDMDistance(cos.ComovingDistance, b)
}

func BenchmarkWACDMLookbackTimeOM(b *testing.B) {
	cos := WACDM{H0: 70, Om0: 0.3, Ol0: 0, W0: -1.2, WA: 2}
	benchmarkWACDMDistance(cos.LookbackTime, b)
}

func BenchmarkWACDMComovingDistanceFlat(b *testing.B) {
	cos := WACDM{H0: 70, Om0: 0.3, Ol0: 0.7, W0: -1, WA: 0}
	benchmarkWACDMDistance(cos.ComovingDistance, b)
}

func BenchmarkWACDMComovingDistancePositiveOk0(b *testing.B) {
	cos := WACDM{H0: 70, Om0: 0.3, Ol0: 0.9, W0: -1, WA: 2}
	benchmarkWACDMDistance(cos.ComovingDistance, b)
}
