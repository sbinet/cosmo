package cosmo

import (
	"math"
	"strings"
	"testing"
)

var zLambdaCDM = []float64{0.5, 1.0, 2.0, 3.0}

// Calculated via Python AstroPy
//   from astropy.cosmology import LambdaCDM
//   z = np.asarray([0.5, 1.0, 2.0, 3.0])
var testTableLambdaCDM = map[string]struct {
	cos      LambdaCDM
	function string
	exp      []float64
}{
	"LambdaCDMDistanceModulus": {LambdaCDM{H0: 70, Om0: 0.3, Ol0: 0.7}, "DistanceModulus", []float64{42.26118542, 44.10023766, 45.95719725, 47.02611193}},
	//   LambdaCDM(70, 0.3, 0.7).luminosity_distance(z)
	"LambdaCDMLuminosityDistanceFlat": {LambdaCDM{H0: 70, Om0: 0.3, Ol0: 0.7}, "LuminosityDistance", []float64{2832.9380939, 6607.65761177, 15539.58622323, 25422.74174519}},
	// LambdaCDM(70, 0.3, 0.6).luminosity_distance(z)
	"LambdaCDMLuminosityDistancePositiveOkLCDM": {LambdaCDM{H0: 70, Om0: 0.3, Ol0: 0.6}, "LuminosityDistance", []float64{2787.51504671, 6479.83450953, 15347.21516211, 25369.7240234}},
	// LambdaCDM(70, 0.3, 0.9).luminosity_distance(z)
	"LambdaCDMLuminosityDistanceNegativeOkLCDM": {LambdaCDM{H0: 70, Om0: 0.3, Ol0: 0.9}, "LuminosityDistance", []float64{2933.96568944, 6896.93040403, 15899.60122012, 25287.53295915}},
	"LambdaCDMAngularDiameterDistance":          {LambdaCDM{H0: 70, Om0: 0.3, Ol0: 0.7}, "AngularDiameterDistance", []float64{1259.08359729, 1651.91440294, 1726.62069147, 1588.92135907}},
	"LambdaCDMComovingTransverseDistance":       {LambdaCDM{H0: 70, Om0: 0.3, Ol0: 0.7}, "ComovingTransverseDistance", []float64{1888.62539593, 3303.82880589, 5179.86207441, 6355.6854363}},
	// LambdaCDM(70, 0.3, 0.6, -1).comoving_transverse_distance(z)
	"LambdaCDMComovingTransverseDistancePositiveOkLCDM": {LambdaCDM{H0: 70, Om0: 0.3, Ol0: 0.6}, "ComovingTransverseDistance", []float64{1858.34336447, 3239.91725476, 5115.73838737, 6342.43100585}},
	// LambdaCDM(70, 0.3, 0.9, -1).comoving_transverse_distance(z)
	"LambdaCDMComovingTransverseDistanceNegativeOkLCDM": {LambdaCDM{H0: 70, Om0: 0.3, Ol0: 0.9}, "ComovingTransverseDistance", []float64{1955.97712629, 3448.46520202, 5299.86707337, 6321.88323979}},
	"LambdaCDMComovingDistanceZ1Z2Integrate":            {LambdaCDM{H0: 70, Om0: 0.3, Ol0: 0.7}, "ComovingDistanceZ1Z2", []float64{1888.62539593, 3303.82880589, 5179.86207441, 6355.6854363}},
	"LambdaCDMComovingDistanceZ1Z2Elliptic":             {LambdaCDM{H0: 70, Om0: 0.3, Ol0: 0.7}, "ComovingDistanceZ1Z2", []float64{1888.62539593, 3303.82880589, 5179.86207441, 6355.6854363}},
	// LambdaCDM(70, 0.3, 0).comoving_distance(z)
	"LambdaCDMComovingDistanceNonflatOM": {LambdaCDM{H0: 70, Om0: 0.3, Ol0: 0.}, "ComovingDistance", []float64{1679.81156606, 2795.15602075, 4244.25192263, 5178.38877021}},
	// LambdaCDM(70, 0.3, 0).comoving_transverse_distance(z)
	"LambdaCDMComovingTransverseDistanceNonflatOM": {LambdaCDM{H0: 70, Om0: 0.3, Ol0: 0.}, "ComovingTransverseDistance", []float64{1710.1240353, 2936.1472205, 4747.54480615, 6107.95517311}},
	// FlatLambdaCDM(70, 1.0).comoving_distance(z)
	"LambdaCDMComovingDistanceEdS": {LambdaCDM{H0: 70, Om0: 1.0, Ol0: 0.}, "ComovingDistance", []float64{1571.79831586, 2508.77651427, 3620.20576208, 4282.7494}},
	// LambdaCDM(70, 0.3, 0).lookback_time(z)
	"LambdaCDMLookbackTime": {LambdaCDM{H0: 70, Om0: 0.3, Ol0: 0.7}, "LookbackTime", []float64{5.04063793, 7.715337, 10.24035689, 11.35445676}},
	// LambdaCDM(70, 0.3, 0).lookback_time(z)
	"LambdaCDMLookbackTimeOM": {LambdaCDM{H0: 70, Om0: 0.3, Ol0: 0.}, "LookbackTime", []float64{4.51471693, 6.62532254, 8.57486509, 9.45923582}},
	// LambdaCDM(70, 0.3, 0.7).lookback_time(z)
	"LambdaCDMLookbackTimeOL": {LambdaCDM{H0: 70, Om0: 0., Ol0: 0.5}, "LookbackTime", []float64{5.0616361, 7.90494991, 10.94241739, 12.52244605}},
	"LambdaCDMAge":            {LambdaCDM{H0: 70, Om0: 0.3, Ol0: 0.6}, "Age", []float64{8.11137578, 5.54558439, 3.13456008, 2.06445301}},
	"LambdaCDMAgeFlatLCDM":    {LambdaCDM{H0: 70, Om0: 0.3, Ol0: 0.7}, "Age", []float64{8.42634602, 5.75164694, 3.22662706, 2.11252719}},
	// FlatLambdaCDM(70, 1.0).age(z)
	"LambdaCDMAgeEdS": {LambdaCDM{H0: 70, Om0: 1.0, Ol0: 0.}, "Age", []float64{5.06897781, 3.29239767, 1.79215429, 1.16403836}},
	// LambdaCDM(70, 0.3, 0.).age(z)
	"LambdaCDMAgeOM": {LambdaCDM{H0: 70, Om0: 0.3, Ol0: 0.}, "Age", []float64{6.78287955, 4.67227393, 2.72273139, 1.83836065}},
	// FlatLambdaCDM(70, 0, 0.5).lookback_time
	"LambdaCDMAgeOL": {LambdaCDM{H0: 70, Om0: 0., Ol0: 0.5}, "Age", []float64{12.34935796, 9.50604415, 6.46857667, 4.88854801}},
}

func TestTableLambdaCDM(t *testing.T) {
	for _, test := range testTableLambdaCDM {
		switch {
		case strings.HasSuffix(test.function, "Z1Z2"):
			runTestsZ0Z2ByName(test.cos, test.function, zLambdaCDM, test.exp, distTol, t)
		default:
			runTestsByName(test.cos, test.function, zLambdaCDM, test.exp, distTol, t)
		}
	}
}

func TestLambdaCDMCosmologyInterface(t *testing.T) {
	ageDistance := func(cos FLRW) {
		z := 0.5
		age := cos.Age(z)
		dc := cos.ComovingDistance(z)
		_, _ = age, dc
	}

	cos := LambdaCDM{H0: 70, Om0: 0.27, Ol0: 0.73}
	ageDistance(cos)
}

// TestE* tests that basic calculation of E
//   https://github.com/astropy/astropy/blob/master/astropy/cosmology/tests/test_cosmology.py
func TestLambdaCDMELcdm(t *testing.T) {
	var z, exp, tol float64
	cos := LambdaCDM{H0: 70, Om0: 0.27, Ol0: 0.73}

	// Check value of E(z=1.0)
	//   OM, OL, OK, z = 0.27, 0.73, 0.0, 1.0
	//   sqrt(OM*(1+z)**3 + OK * (1+z)**2 + OL)
	//   sqrt(0.27*(1+1.0)**3 + 0.0 * (1+1.0)**2 + 0.73)
	//   sqrt(0.27*8 + 0 + 0.73)
	//   sqrt(2.89)
	z = 1.0
	exp = 1.7
	tol = 1e-9
	runTest(cos.E, z, exp, tol, t, 0)

	exp = 1 / 1.7
	runTest(cos.Einv, z, exp, tol, t, 0)
}

// Analytic case of Omega_Lambda = 0
func TestLambdaCDMEOm(t *testing.T) {
	zLambdaCDM := []float64{1.0, 10.0, 500.0, 1000.0}
	cos := LambdaCDM{H0: 70, Om0: 1.0, Ol0: 0.}
	hubbleDistance := SpeedOfLightKmS / cos.H0
	expVec := make([]float64, len(zLambdaCDM))
	for i, z := range zLambdaCDM {
		expVec[i] = 2.0 * hubbleDistance * (1 - math.Sqrt(1/(1+z)))
	}
	runTests(cos.ComovingDistance, zLambdaCDM, expVec, distTol, t)
}
