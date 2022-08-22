package cmd

import (
	"encoding/csv"
	"fmt"
	"octo-pull/apiclient"
	"octo-pull/models"
	"os"

	"github.com/spf13/cobra"
)

type MeterReadOptions struct {
	ApiKey         string
	Mpan           string
	SerialNumber   string
	ExportFilename string
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
	electricityCmd.Flags().StringVar(&opts.ExportFilename, "export", "", "Export to CSV")

	electricityCmd.MarkFlagRequired("apikey")
	electricityCmd.MarkFlagRequired("mpan")
	electricityCmd.MarkFlagRequired("serialnumber")

	return electricityCmd
}

func GetMeterReadings(opts *MeterReadOptions) {
	fmt.Println("Obtaining electricity meter readings with API Key: ", opts.ApiKey)
	consumption := apiclient.GetRecent(opts.ApiKey, opts.Mpan, opts.SerialNumber)

	if opts.ExportFilename != "" {
		CreateCSVFromResult(opts, consumption)
	}

	sum := 0.0
	for _, c := range consumption.Results {
		sum += c.Consumption
		fmt.Printf("Start %s, End %s: Consumption (kwh) %f\n", c.Interval_start, c.Interval_end, c.Consumption)
	}

	fmt.Printf("Average Consumption (kwh) %f\n", sum/float64(len(consumption.Results)))
	fmt.Printf("Total consumption (kwh) %f\n", sum)
}

func CreateCSVFromResult(opts *MeterReadOptions, consumption models.ConsumptionResponse) {
	file, err := os.Create(opts.ExportFilename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	write := csv.NewWriter(file)
	defer write.Flush()

	write.Write([]string{"Start", "End", "Consumptions (Kwh)"})

	for _, v := range consumption.Results {
		write.Write([]string{v.Interval_start, v.Interval_end, fmt.Sprint(v.Consumption)})
	}
}
