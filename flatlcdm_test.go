package cosmo

import (
	"math"
	"testing"
)

var zFlatLCDM = []float64{0.5, 1.0, 2.0, 3.0}

// Calculated via Python AstroPy
//   from astropy.cosmology import FlatLambdaCDM
//   z = np.asarray([0.5, 1.0, 2.0, 3.0])
var testTableFlatLCDM = map[string]struct {
	cos      FlatLCDM
	function string
	exp      []float64
}{
	"FlatLCDMDistanceModulus":    {FlatLCDM{H0: 70, Om0: 0.3}, "DistanceModulus", []float64{42.26118542, 44.10023766, 45.95719725, 47.02611193}},
	"FlatLCDMLuminosityDistance": {FlatLCDM{H0: 70, Om0: 0.3}, "LuminosityDistance", []float64{2832.9380939, 6607.65761177, 15539.58622323, 25422.74174519}},
	// Calculated via FlatLambdaCDM(70, 1.0).comoving_distance(z)
	"FlatLCDMComovingDistanceEdS":           {FlatLCDM{H0: 70, Om0: 1}, "ComovingDistance", []float64{1571.79831586, 2508.77651427, 3620.20576208, 4282.7494}},
	"FlatLCDMAngularDiameterDistance":       {FlatLCDM{H0: 70, Om0: 0.3}, "AngularDiameterDistance", []float64{1259.08359729, 1651.91440294, 1726.62069147, 1588.92135907}},
	"FlatLCDMComovingTransverseDistance":    {FlatLCDM{H0: 70, Om0: 0.3}, "ComovingTransverseDistance", []float64{1888.62539593, 3303.82880589, 5179.86207441, 6355.6854363}},
	"FlatLCDMComovingDistanceZ1Z2Integrate": {FlatLCDM{H0: 70, Om0: 0.3}, "ComovingDistanceZ1Z2", []float64{1888.62539593, 3303.82880589, 5179.86207441, 6355.6854363}},
	"FlatLCDMComovingDistanceZ1Z2Elliptic":  {FlatLCDM{H0: 70, Om0: 0.3}, "ComovingDistanceZ1Z2", []float64{1888.62539593, 3303.82880589, 5179.86207441, 6355.6854363}},
	// Calculated via FlatLambdaCDM(70, 0.3).lookback_time(z)
	"FlatLCDMLookbackTime": {FlatLCDM{H0: 70, Om0: 0.3}, "LookbackTime", []float64{5.04063793, 7.715337, 10.24035689, 11.35445676}},
	// Calculated via FlatLambdaCDM(70, 1.0).lookback_time(z)
	"FlatLCDMLookbackTimeEdS": {FlatLCDM{H0: 70, Om0: 1.0}, "LookbackTime", []float64{4.24332906, 6.0199092, 7.52015258, 8.14826851}},
	//   FlatLambdaCDM(70, 0.3).age(z)
	"FlatLCDMAge": {FlatLCDM{H0: 70, Om0: 0.3}, "Age", []float64{8.42634602, 5.75164694, 3.22662706, 2.11252719}},
	//   FlatLambdaCDM(70, 1.0).age(z)
	"FlatLCDMAgeEdS": {FlatLCDM{H0: 70, Om0: 1.0}, "Age", []float64{5.06897781, 3.29239767, 1.79215429, 1.16403836}},
}

func TestFlatLCDMCosmologyInterface(t *testing.T) {
	age_distance := func(cos FLRW) {
		z := 0.5
		age := cos.Age(z)
		dc := cos.ComovingDistance(z)
		_, _ = age, dc
	}

	cos := FlatLCDM{H0: 70, Om0: 0.27}
	age_distance(cos)
}

// TestE* tests that basic calculation of E
//   https://github.com/astropy/astropy/blob/master/astropy/cosmology/tests/test_cosmology.py
func TestFlatLCDME(t *testing.T) {
	var z, exp float64
	cos := FlatLCDM{H0: 70, Om0: 0.27}

	// Check value of E(z=1.0)
	//   OM, OL, OK, z = 0.27, 0.73, 0.0, 1.0
	//   sqrt(OM*(1+z)**3 + OK * (1+z)**2 + OL)
	//   sqrt(0.27*(1+1.0)**3 + 0.0 * (1+1.0)**2 + 0.73)
	//   sqrt(0.27*8 + 0 + 0.73)
	//   sqrt(2.89)
	z = 1.0
	exp = 1.7
	runTest(cos.E, z, exp, eTol, t, 0)

	exp = 1 / 1.7
	runTest(cos.Einv, z, exp, eTol, t, 0)
}

func TestFlatLCDMDistanceModulusReflect(t *testing.T) {
	test := testTableFlatLCDM["FlatLCDMDistanceModulus"]
	runTestsByName(test.cos, test.function, zFlatLCDM, test.exp, distmodTol, t)
}

func TestFlatLCDMDistanceModulus(t *testing.T) {
	test := testTableFlatLCDM["FlatLCDMDistanceModulus"]
	runTestsByName(test.cos, test.function, zFlatLCDM, test.exp, distmodTol, t)
}

func TestFlatLCDMLuminosityDistance(t *testing.T) {
	test := testTableFlatLCDM["FlatLCDMLuminosityDistance"]
	runTestsByName(test.cos, test.function, zFlatLCDM, test.exp, distmodTol, t)
}

func TestFlatLCDMAngularDiameterDistance(t *testing.T) {
	test := testTableFlatLCDM["FlatLCDMAngularDiameterDistance"]
	runTestsByName(test.cos, test.function, zFlatLCDM, test.exp, distTol, t)
}

func TestFlatLCDMComovingTransverseDistance(t *testing.T) {
	test := testTableFlatLCDM["FlatLCDMComovingTransverseDistance"]
	runTestsByName(test.cos, test.function, zFlatLCDM, test.exp, distTol, t)
}

func TestFlatLCDMComovingDistanceZ1Z2Integrate(t *testing.T) {
	test := testTableFlatLCDM["FlatLCDMComovingDistanceZ1Z2Integrate"]
	runTestsZ0Z2ByName(test.cos, test.function, zFlatLCDM, test.exp, distTol, t)
}

func TestFlatLCDMComovingDistanceZ1Z2Elliptic(t *testing.T) {
	test := testTableFlatLCDM["FlatLCDMComovingDistanceZ1Z2Elliptic"]
	runTestsZ0Z2ByName(test.cos, test.function, zFlatLCDM, test.exp, distTol, t)
}

func TestFlatLCDMLookbackTime(t *testing.T) {
	test := testTableFlatLCDM["FlatLCDMLookbackTime"]
	runTestsByName(test.cos, test.function, zFlatLCDM, test.exp, ageTol, t)
}

func TestFlatLCDMLookbackTimeEdS(t *testing.T) {
	test := testTableFlatLCDM["FlatLCDMLookbackTimeEdS"]
	runTestsByName(test.cos, test.function, zFlatLCDM, test.exp, ageTol, t)
}

func TestFlatLCDMAge(t *testing.T) {
	test := testTableFlatLCDM["FlatLCDMAge"]
	runTestsByName(test.cos, test.function, zFlatLCDM, test.exp, ageTol, t)
}

func TestFlatLCDMAgeEdS(t *testing.T) {
	test := testTableFlatLCDM["FlatLCDMAgeEdS"]
	runTestsByName(test.cos, test.function, zFlatLCDM, test.exp, ageTol, t)
}

// Analytic case of Omega_Lambda = 0
func TestFlatLCDMEOm(t *testing.T) {
	cos := FlatLCDM{H0: 70, Om0: 1.0}
	z_vec := []float64{1.0, 10.0, 500.0, 1000.0}
	hubbleDistance := SpeedOfLightKmS / cos.H0
	exp_vec := make([]float64, len(z_vec))
	for i, z := range z_vec {
		exp_vec[i] = 2.0 * hubbleDistance * (1 - math.Sqrt(1/(1+z)))
	}
	runTests(cos.ComovingDistance, z_vec, exp_vec, distTol, t)
}
