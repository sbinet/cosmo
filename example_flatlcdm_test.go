package cosmo

import (
	"fmt"
)

// Calculated via
//   from astropy.cosmology import FlatLCDM
//   z = np.asarray([0.5, 1.0, 2.0, 3.0])
//   FlatLCDM(70, 0.3, 0.7).distmod(z)
//   FlatLCDM(70, 0.3, 0.7).luminosity_distance(z)
//   FlatLCDM(70, 0.3, 0.7).angular_diameter_distance(z)

func ExampleFlatLCDM() {
	cos := FlatLCDM{Om0: 0.3, H0: 70, Tcmb0: 0.}

	z_vec := []float64{0.5, 1.0, 2.0, 3.0}
	distmod_vec := make([]float64, len(z_vec))
	lumdist_vec := make([]float64, len(z_vec))
	angdist_vec := make([]float64, len(z_vec))
	for i, z := range z_vec {
		distmod_vec[i] = cos.DistanceModulus(z)
		lumdist_vec[i] = cos.LuminosityDistance(z)
		angdist_vec[i] = cos.AngularDiameterDistance(z)
	}

	fmt.Println("Distance Modulus [mag]")
	fmt.Println(distmod_vec)
	fmt.Println("Luminosity Distance [Mpc]")
	fmt.Println(lumdist_vec)
	fmt.Println("Angular Diameter Distance [Mpc]")
	fmt.Println(angdist_vec)
	// Output:
	// Distance Modulus [mag]
	// [42.26118542154089 44.10023765554372 45.95719725271018 47.026111928689645]
	// Luminosity Distance [Mpc]
	// [2832.938093900105 6607.65761177494 15539.58622322811 25422.74174518986]
	// Angular Diameter Distance [Mpc]
	// [1259.0835972889354 1651.914402943735 1726.6206914697902 1588.9213590743661]
}