package cmd

import (
	"fmt"
	"octo-pull/apiclient"

	"github.com/spf13/cobra"
)

type MeterReadOptions struct {
	ApiKey       string
	Mpan         string
	SerialNumber string
}

func NewCmdElectricity() *cobra.Command {
	opts := MeterReadOptions{}

	var electricityCmd = &cobra.Command{
		Use:   "electricity",
		Short: "Obtains electricity meter readings",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			GetMeterReadings(&opts)
		},
	}

	electricityCmd.Flags().StringVar(&opts.ApiKey, "apikey", "", "Octopus API Key (Required)")
	electricityCmd.Flags().StringVar(&opts.Mpan, "mpan", "", "Octopus Meter MPAN (Required)")
	electricityCmd.Flags().StringVar(&opts.SerialNumber, "serialnumber", "", "Octopus Meter Serial Number(Required)")

	return electricityCmd
}

func GetMeterReadings(opts *MeterReadOptions) {
	fmt.Println("Obtaining electricity meter readings with API Key: ", opts.ApiKey)
	consumption := apiclient.GetRecent(opts.ApiKey, opts.Mpan, opts.SerialNumber)

	sum := 0.0

	for _, c := range consumption {
		sum += c.Consumption
		fmt.Printf("Start %s, End %s: Consumption (kwh) %f\n", c.Interval_start, c.Interval_end, c.Consumption)
	}

	fmt.Printf("Average Consumption (kwh) %f\n", sum/float64(len(consumption)))
	fmt.Printf("Total consumption (kwh) %f\n", sum)
}
