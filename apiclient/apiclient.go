package apiclient

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"octo-pull/models"

	"gopkg.in/h2non/gentleman.v2"
)

func getClient(apiKey string) *gentleman.Client {
	cli := gentleman.New()
	cli.URL("https://api.octopus.energy")
	cli.AddHeader("Authorization", authorisationValue(apiKey))

	return cli
}

func GetRecent(apiKey string, mpan string, serialnumber string) models.ConsumptionResponse {
	cli := getClient(apiKey)
	req := cli.Request()
	req.Method("GET")
	req.Path(fmt.Sprintf("/v1/electricity-meter-points/%s/meters/%s/consumption/", mpan, serialnumber))
	req.AddQuery("group_by", "day")
	req.AddQuery("page_size", "25000")
	res, err := req.Send()
	if err != nil {
		panic(err)
	}

	var consumptionResponse models.ConsumptionResponse
	err = json.Unmarshal(res.Bytes(), &consumptionResponse)

	if err != nil {
		panic(err)
	}

	return consumptionResponse
}

func authorisationValue(apiKey string) string {
	data := apiKey + ":"
	encoded := base64.StdEncoding.EncodeToString([]byte(data))
	return "Basic " + encoded
}
