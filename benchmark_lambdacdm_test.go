package cosmo

import (
	"testing"
)

func benchmarkLambdaCDMEN(n int, b *testing.B) {
	cos := LambdaCDM{H0: 70, Om0: 0.3, Ol0: 0.9}

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

func BenchmarkLambdaCDMEN(b *testing.B) {
	benchmarkLambdaCDMEN(10000, b)
}

func BenchmarkLambdaCDMENdistance(b *testing.B) {
	cos := LambdaCDM{H0: 70, Om0: 0.3, Ol0: 0.9}
	benchmarkLambdaCDMNdistance(10000, cos.E, b)
}

func BenchmarkLambdaCDME(b *testing.B) {
	cos := LambdaCDM{H0: 70, Om0: 0.3, Ol0: 0.9}
	z := 1.0
	for i := 0; i < b.N; i++ {
		cos.E(z)
	}
}

func BenchmarkLambdaCDMEinv(b *testing.B) {
	cos := LambdaCDM{H0: 70, Om0: 0.3, Ol0: 0.9}
	z := 1.0
	for i := 0; i < b.N; i++ {
		cos.Einv(z)
	}
}

// benchmarkLambdaCDMDistance is a helper function to be called by specific benchmarkLambdaCDMs
func benchmarkLambdaCDMDistance(f func(float64) float64, b *testing.B) {
	z := 1.0
	for i := 0; i < b.N; i++ {
		f(z)
	}
}

// benchmarkLambdaCDMNdistance is a helper function to be called by specific benchmarkLambdaCDMs
func benchmarkLambdaCDMNdistance(n int, f func(float64) float64, b *testing.B) {
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

func BenchmarkLambdaCDMComovingDistance(b *testing.B) {
	cos := LambdaCDM{H0: 70, Om0: 0.3, Ol0: 0.9}
	benchmarkLambdaCDMDistance(cos.ComovingDistance, b)
}

func BenchmarkLambdaCDMComovingTransverseDistance(b *testing.B) {
	cos := LambdaCDM{H0: 70, Om0: 0.3, Ol0: 0.9}
	benchmarkLambdaCDMDistance(cos.ComovingTransverseDistance, b)
}

func BenchmarkLambdaCDMLuminosityDistance(b *testing.B) {
	cos := LambdaCDM{H0: 70, Om0: 0.3, Ol0: 0.9}
	benchmarkLambdaCDMDistance(cos.LuminosityDistance, b)
}

func BenchmarkLambdaCDMLookbackTime(b *testing.B) {
	cos := LambdaCDM{H0: 70, Om0: 0.3, Ol0: 0.9}
	benchmarkLambdaCDMDistance(cos.LookbackTime, b)
}

func BenchmarkLambdaCDMNComovingDistance(b *testing.B) {
	cos := LambdaCDM{H0: 70, Om0: 0.3, Ol0: 0.9}
	benchmarkLambdaCDMNdistance(10000, cos.ComovingDistance, b)
}

func BenchmarkLambdaCDMNLuminosityDistance(b *testing.B) {
	cos := LambdaCDM{H0: 70, Om0: 0.3, Ol0: 0.9}
	benchmarkLambdaCDMNdistance(10000, cos.LuminosityDistance, b)
}

func BenchmarkLambdaCDMNE(b *testing.B) {
	cos := LambdaCDM{H0: 70, Om0: 0.3, Ol0: 0.9}
	benchmarkLambdaCDMNdistance(10000, cos.E, b)
}

func BenchmarkLambdaCDMComovingDistanceOM(b *testing.B) {
	cos := LambdaCDM{H0: 70, Om0: 0.27, Ol0: 0.}
	benchmarkLambdaCDMDistance(cos.ComovingDistance, b)
}

func BenchmarkLambdaCDMLookbackTimeOM(b *testing.B) {
	cos := LambdaCDM{H0: 70, Om0: 0.27, Ol0: 0.}
	benchmarkLambdaCDMDistance(cos.LookbackTime, b)
}

func BenchmarkLambdaCDMComovingDistanceFlat(b *testing.B) {
	cos := LambdaCDM{H0: 70, Om0: 0.3, Ol0: 0.7}
	benchmarkLambdaCDMDistance(cos.ComovingDistance, b)
}

func BenchmarkLambdaCDMComovingDistancePositiveOk0(b *testing.B) {
	cos := LambdaCDM{H0: 70, Om0: 0.3, Ol0: 0.9}
	benchmarkLambdaCDMDistance(cos.ComovingDistance, b)
}
