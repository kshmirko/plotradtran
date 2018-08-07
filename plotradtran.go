package main

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/kshmirko/radtran/libplot"
	"github.com/kshmirko/radtran/librt3"
)

func main() {

	f, err := os.Open("tmpout.out")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	for i := 0; i < 11; i++ {
		_ = librt3.ReadLine(f)
	}

	z := 0.0

	var phi, mu, I, Q, ang []float64
	var phi_tmp, mu_tmp, I_tmp, Q_tmp float64

	/*phi := make([]float64, NLines)
	mu := make([]float64, NLines)
	I := make([]float64, NLines)
	Q := make([]float64, NLines)
	ang := make([]float64, NLines)*/
	deg2rad := math.Pi / 180.0
	//j := 0
	for {
		_, err := fmt.Fscanf(f, "%f %f %f %f %f\n", &z, &phi_tmp, &mu_tmp, &I_tmp, &Q_tmp)
		if err != nil {
			break
		}
		if mu_tmp > 0.0 {
			phi = append(phi, phi_tmp)
			mu = append(mu, mu_tmp)
			I = append(I, I_tmp)
			Q = append(Q, Q_tmp)

			fmt.Printf("%f %f %f %f %f\n", z, phi_tmp, mu_tmp, I_tmp, Q_tmp)
			ang = append(ang, math.Cos(phi_tmp*deg2rad)*math.Acos(mu_tmp)/deg2rad)
		}
	}

	libplot.VizualizeIntensity("Intensities.pdf", 1.0, &ang, &I, &Q)
	libplot.VizualizePolarization("Polarization.pdf", 1.0, &ang, &I, &Q)
}
