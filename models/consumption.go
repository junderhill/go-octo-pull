package models

type Consumption struct {
	Consumption    float64
	Interval_start string
	Interval_end   string
}

type ConsumptionResponse struct {
	Count   int
	Results []Consumption
}
