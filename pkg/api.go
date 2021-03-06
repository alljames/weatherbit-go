package weatherbit

import (
	"fmt"
	"math"
	"time"
)

// DegToRad is the constant used to convert from Degrees to Radians
// defined as 1° = π/180° = 0.005555556π = 0.01745329252 rad
const (
	DegToRad float64 = 0.01745329252
)

// GetResponse is exported and calls an internal function
func GetResponse(p Parameters) WbResponse {
	return syncHTTPGets(p)
}

// Prettyprint can be customised by the user to print the elements of
// their choosing within the WbResponse struct, but is mostly here to demonstrate
// how to access elements out of the WbResponse struct
func Prettyprint(wbr WbResponse) {

	count := len(wbr.Data)
	fmt.Printf("Number of observations returned: %d\n", count)
	fmt.Printf("City: %s\n", wbr.Data[0].CityName)
	for i := 0; i < count; i++ {

		observationtimestamp := int64(wbr.Data[i].LastObservationTimeStamp)
		observationtime := time.Unix(observationtimestamp, (observationtimestamp / 1e9))
		fmt.Printf("Observation time: %v\t (%v)\n", observationtime, (time.Until(observationtime)))

		dni := wbr.Data[i].Dni
		dhi := wbr.Data[i].Dhi
		solarElevationAngle := wbr.Data[i].SolarHourAngle // degrees
		solarElevationRadians := solarElevationAngle * DegToRad
		fmt.Printf("Degrees: %.3f Radians: %.3f\n", solarElevationAngle, solarElevationRadians)
		ghi := (dni*math.Cos(solarElevationRadians) + dhi)
		fmt.Printf("Global Horizontal Irradiance calc: %.3f\n", ghi)
		fmt.Printf("Global Horizontal Irradiance wbit: %.3f\n", wbr.Data[i].Ghi)
		fmt.Printf("Temperature: %.2f\t", wbr.Data[i].Temperature)
		fmt.Printf("CloudsHi: %.2f\n", wbr.Data[i].CloudsHi)
		fmt.Printf("UV: %.3f\t DNI: %.3f\t DHI: %.3f\t", wbr.Data[i].UV, dni, dhi)

		// Global Horizontal (GHI) = Direct Normal (DNI) X cos(θ) + Diffuse Horizontal (DHI)
		// TODO: check degree / radians. use golang geo package to convert between the two
	}

}
