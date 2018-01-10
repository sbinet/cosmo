package cosmo

import (
	"testing"
)

func benchmarkWCDMEN(n int, b *testing.B) {
	cos := WCDM{H0: 70, Om0: 0.2, Ol0: 0.7, W0: -1.2}

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

func BenchmarkWCDMEN(b *testing.B) {
	benchmarkWCDMEN(10000, b)
}

func BenchmarkWCDMENdistance(b *testing.B) {
	cos := WCDM{H0: 70, Om0: 0.2, Ol0: 0.7, W0: -1.2}
	benchmarkWCDMNdistance(10000, cos.E, b)
}

func BenchmarkWCDME(b *testing.B) {
	cos := WCDM{H0: 70, Om0: 0.2, Ol0: 0.7, W0: -1.2}
	z := 1.0
	for i := 0; i < b.N; i++ {
		cos.E(z)
	}
}

func BenchmarkWCDMEinv(b *testing.B) {
	cos := WCDM{H0: 70, Om0: 0.2, Ol0: 0.7, W0: -1.2}
	z := 1.0
	for i := 0; i < b.N; i++ {
		cos.Einv(z)
	}
}

// benchmarkWCDMDistance is a helper function to be called by specific benchmarkWCDMs
func benchmarkWCDMDistance(f func(float64) float64, b *testing.B) {
	z := 1.0
	for i := 0; i < b.N; i++ {
		f(z)
	}
}

// benchmarkWCDMNdistance is a helper function to be called by specific benchmarkWCDMs
func benchmarkWCDMNdistance(n int, f func(float64) float64, b *testing.B) {
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

func BenchmarkWCDMComovingDistance(b *testing.B) {
	cos := WCDM{H0: 70, Om0: 0.2, Ol0: 0.7, W0: -1.2}
	benchmarkWCDMDistance(cos.ComovingDistance, b)
}

func BenchmarkWCDMComovingTransverseDistance(b *testing.B) {
	cos := WCDM{H0: 70, Om0: 0.2, Ol0: 0.7, W0: -1.2}
	benchmarkWCDMDistance(cos.ComovingTransverseDistance, b)
}

func BenchmarkWCDMLuminosityDistance(b *testing.B) {
	cos := WCDM{H0: 70, Om0: 0.2, Ol0: 0.7, W0: -1.2}
	benchmarkWCDMDistance(cos.LuminosityDistance, b)
}

func BenchmarkWCDMLookbackTime(b *testing.B) {
	cos := WCDM{H0: 70, Om0: 0.2, Ol0: 0.7, W0: -1.2}
	benchmarkWCDMDistance(cos.LookbackTime, b)
}

func BenchmarkWCDMNComovingDistance(b *testing.B) {
	cos := WCDM{H0: 70, Om0: 0.2, Ol0: 0.7, W0: -1.2}
	benchmarkWCDMNdistance(10000, cos.ComovingDistance, b)
}

func BenchmarkWCDMNLuminosityDistance(b *testing.B) {
	cos := WCDM{H0: 70, Om0: 0.2, Ol0: 0.7, W0: -1.2}
	benchmarkWCDMNdistance(10000, cos.LuminosityDistance, b)
}

func BenchmarkWCDMNE(b *testing.B) {
	cos := WCDM{H0: 70, Om0: 0.2, Ol0: 0.7, W0: -1.2}
	benchmarkWCDMNdistance(10000, cos.E, b)
}

func BenchmarkWCDMComovingDistanceOM(b *testing.B) {
	cos := WCDM{H0: 70, Om0: 0.3, Ol0: 0, W0: -1.2}
	benchmarkWCDMDistance(cos.ComovingDistance, b)
}

func BenchmarkWCDMLookbackTimeOM(b *testing.B) {
	cos := WCDM{H0: 70, Om0: 0.3, Ol0: 0, W0: -1.2}
	benchmarkWCDMDistance(cos.LookbackTime, b)
}

func BenchmarkWCDMComovingDistanceFlat(b *testing.B) {
	cos := WCDM{H0: 70, Om0: 0.3, Ol0: 0.7, W0: -1}
	benchmarkWCDMDistance(cos.ComovingDistance, b)
}

func BenchmarkWCDMComovingDistancePositiveOk0(b *testing.B) {
	cos := WCDM{H0: 70, Om0: 0.3, Ol0: 0.9, W0: -1}
	benchmarkWCDMDistance(cos.ComovingDistance, b)
}
