package cosmo

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/integrate/quad"
)

// WACDM provides cosmological distances, age, and look-back time
// for a w(a) cosmology:
// matter, dark energy, and curvature,
// with a
// w = w0 + wa * (1-a)
// equation-of-state parameter for dark energy.
type WACDM struct {
	H0      float64 // Hubble constant at z=0.  [km/s/Mpc]
	Om0     float64 // Matter Density at z=0
	Ol0     float64 // Dark Energy density Lambda at z=0
	W0      float64 // Dark energy equation-of-state parameter, w0 + wa*(1-a) = p/rho
	WA      float64 // Dark energy equation-of-state parameter, w0 + wa*(1-a) = p/rho
	Ogamma0 float64 // Photon density
	Onu0    float64 // Neutrino density
}

// Tcmb0   float64 // Temperature of the CMB at z=0.  [K]
// nuToPhotonDensity float64 // Neutrino density / photon density

func (cos WACDM) String() string {
	return fmt.Sprintf("WACDM{H0: %v, Om0: %v, Ol0: %v, W0: %v, WA: %v}",
		cos.H0, cos.Om0, cos.Ol0, cos.W0, cos.WA)
}

// Ok0 is the curvature density at z=0
func (cos WACDM) Ok0() (curvatureDensity float64) {
	return 1 - (cos.Om0 + cos.Ol0)
}

// DistanceModulus is the magnitude difference between 1 Mpc and
// the luminosity distance for the given z.
func (cos WACDM) DistanceModulus(z float64) (distanceModulusMag float64) {
	return 5*math.Log10(cos.LuminosityDistance(z)) + 25
}

// LuminosityDistance is the radius of effective sphere over which the light has spread out
func (cos WACDM) LuminosityDistance(z float64) (distanceMpc float64) {
	return (1 + z) * cos.ComovingTransverseDistance(z)
}

// AngularDiameterDistance is the ratio of physical transverse size to angular size
func (cos WACDM) AngularDiameterDistance(z float64) (distanceMpcRad float64) {
	return cos.ComovingTransverseDistance(z) / (1 + z)
}

// ComovingTransverseDistance is the comoving distance at z as seen from z=0
func (cos WACDM) ComovingTransverseDistance(z float64) (distanceMpcRad float64) {
	return cos.ComovingTransverseDistanceZ1Z2(0, z)
}

// ComovingTransverseDistanceZ1Z2 is the comoving distance at z2 as seen from z1
func (cos WACDM) ComovingTransverseDistanceZ1Z2(z1, z2 float64) (distanceMpcRad float64) {
	return comovingTransverseDistanceZ1Z2(cos, z1, z2)
}

// HubbleDistance is the inverse of the Hubble parameter
//   distance : [Mpc]
func (cos WACDM) HubbleDistance() float64 {
	return hubbleDistance(cos.H0)
}

// ComovingDistance is the distance that is constant with the Hubble flow
// expressed in the physical distance at z=0.
//
// As the scale factor a = 1/(1+z) increases from 0.5 to 1,
// two objects separated by a proper distance of 10 Mpc at a=0.5 (z=1)
// will be separated by a proper distance of 2*10 Mpc at a=1.0 (z=0).
// The comoving distance between these objects is 20 Mpc.
func (cos WACDM) ComovingDistance(z float64) (distanceMpc float64) {
	return cos.ComovingDistanceZ1Z2(0, z)
}

// ComovingDistanceZ1Z2Integrate is the comoving distance between two z
// in a flat lambda CDM cosmology using fixed Gaussian quadrature integration.
func (cos WACDM) comovingDistanceZ1Z2Integrate(z1, z2 float64) (distanceMpc float64) {
	n := 1000 // Integration will be n-point Gaussian quadrature
	return cos.HubbleDistance() * quad.Fixed(cos.Einv, z1, z2, n, nil, 0)
}

// ComovingDistanceZ1Z2 is the base function for calculation of comoving distances
// Here is where the choice of fundamental calculation method is made:
// Fall back to simpler cosmology, or quadature integration
func (cos WACDM) ComovingDistanceZ1Z2(z1, z2 float64) (distanceMpc float64) {
	switch {
	// Test for Ol0==0 first so that (Om0, Ol0) = (1, 0)
	// is handled by the analytic solution
	// rather than the explicit integration.
	case cos.Ol0 == 0:
		return comovingDistanceOMZ1Z2(z1, z2, cos.Om0, cos.H0)
	case cos.WA == 0:
		wcdm_cos := WCDM{H0: cos.H0, Om0: cos.Om0, Ol0: cos.Ol0, W0: cos.W0}
		return wcdm_cos.ComovingDistanceZ1Z2(z1, z2)
	default:
		return cos.comovingDistanceZ1Z2Integrate(z1, z2)
	}
}

// LookbackTime is the time from redshift 0 to z in Gyr.
func (cos WACDM) LookbackTime(z float64) (timeGyr float64) {
	switch {
	case (cos.Ol0 == 0) && (0 < cos.Om0) && (cos.Om0 != 1):
		return lookbackTimeOM(z, cos.Om0, cos.H0)
	case cos.WA == 0:
		wcdm_cos := WCDM{H0: cos.H0, Om0: cos.Om0, Ol0: cos.Ol0, W0: cos.W0}
		return wcdm_cos.LookbackTime(z)
	default:
		return cos.lookbackTimeIntegrate(z)
	}
}

// lookbackTimeIntegrate is the lookback time using explicit integration
func (cos WACDM) lookbackTimeIntegrate(z float64) (timeGyr float64) {
	n := 1000 // Integration will be n-point Gaussian quadrature
	integrand := func(z float64) float64 { return cos.Einv(z) / (1 + z) }
	return hubbleTime(cos.H0) * quad.Fixed(integrand, 0, z, n, nil, 0)
}

// Age is the time from redshift ∞ to z in Gyr.
func (cos WACDM) Age(z float64) (timeGyr float64) {
	switch {
	case (cos.Ol0 == 0) && (0 < cos.Om0) && (cos.Om0 != 1):
		return ageOM(z, cos.Om0, cos.H0)
	case cos.WA == 0:
		wcdm_cos := WCDM{H0: cos.H0, Om0: cos.Om0, Ol0: cos.Ol0, W0: cos.W0}
		return wcdm_cos.Age(z)
	default:
		return cos.ageIntegrate(z)
	}
}

// ageIntegrate is the time from redshift ∞ to z
// using explicit integration.
//
// Basic integrand can be found in many texts
// I happened to copy this from
// Thomas and Kantowski, 2000, PRD, 62, 103507.  Eq. 1.
// Current implementation is fixed quadrature using mathext.integrate.quad.Fixed
func (cos WACDM) ageIntegrate(z float64) (timeGyr float64) {
	n := 1000 // Integration will be n-point Gaussian quadrature
	integrand := func(z float64) float64 {
		denom := (1 + z) * cos.E(z)
		return 1 / denom
	}
	// When given math.Inf(), quad.Fixed automatically redefines variables
	// to successfully do the numerical integration.
	return hubbleTime(cos.H0) * quad.Fixed(integrand, z, math.Inf(1), n, nil, 0)
}

// E is the Hubble parameter as a fraction of its present value.
// E.g., Hogg arXiv:9905116  Eq. 14
// Linder, 2003, PhRvL, 90, 130, Eq. 5, 7
func (cos WACDM) E(z float64) (fractionalHubbleParameter float64) {
	oR := cos.Ogamma0 + cos.Onu0
	var deScale float64
	switch {
	case (cos.W0 == -1) && (cos.WA == 0):
		deScale = 1
	case cos.WA == 0:
		deScale = math.Pow(1+z, 3*(1+cos.W0))
	default:
		deScale = math.Pow(1+z, 3*(1+cos.W0+cos.WA)) * math.Exp(-3*cos.WA*z/(1+z))
	}
	Ok0 := 1 - (cos.Om0 + cos.Ol0)
	return math.Sqrt((1+z)*(1+z)*(1+z)*(1+z)*oR + (1+z)*(1+z)*(1+z)*cos.Om0 +
		(1+z)*(1+z)*Ok0 + cos.Ol0*deScale)
}

// Einv is the inverse Hubble parameter
// Implementation is just to return E(z)
func (cos WACDM) Einv(z float64) (invFractionalHubbleParameter float64) {
	// 1/Sqrt() is not notably slower than Pow(-0.5)
	//
	// Pow(-0.5) is in fact implemented as 1/Sqrt() in math.pow.go
	// func pow(x, y float64) float64 {
	//    [...]
	// case y == -0.5:
	//    return 1 / Sqrt(x)
	//
	// Thus we just return the inverse of E(z) instead of rewriting out here.
	return 1 / cos.E(z)
}
