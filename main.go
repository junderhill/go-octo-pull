package main

import (
	"fmt"
	"octo-pull/apiclient"
	"os"
)

func main() {
	apikey := os.Getenv("OCTOPUS_API_KEY")
	mpan := os.Getenv("OCTOPUS_MPAN")
	serialnumber := os.Getenv("OCTOPUS_SERIAL_NUMBER")
	consumption := apiclient.GetRecent(apikey, mpan, serialnumber)

	sum := 0.0
	for _, c := range consumption {
		sum += c.Consumption
		fmt.Printf("Start %s, End %s: Consumption (kwh) %f\n", c.Interval_start, c.Interval_end, c.Consumption)
	}

	fmt.Printf("Average Consumption (kwh) %f\n", sum/float64(len(consumption)))
	fmt.Printf("Total consumption (kwh) %f\n", sum)

}
