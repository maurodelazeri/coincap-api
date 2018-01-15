package coincap

import (
	"encoding/json"
	"fmt"
)

type GlobalData struct {
	AltCap        float64 `json:"altCap"`
	BitnodesCount int     `json:"bitnodesCount"`
	BtcCap        float64 `json:"btcCap"`
	BtcPrice      float64 `json:"btcPrice"`
	Dom           float64 `json:"dom"`
	TotalCap      float64 `json:"totalCap"`
	VolumeAlt     float64 `json:"volumeAlt"`
	VolumeBtc     float64 `json:"volumeBtc"`
	VolumeTotal   float64 `json:"volumeTotal"`
}

// coincap API implementation of Gobal Data endpoint
//
// Endpoint: /global/
// Method: GET
//
// Example: http://coincap.io/global
//
/*
{
    "altCap": 70056946653.0021,
    "bitnodesCount": 9350,
    "btcCap": 65003614189.66167,
    "btcPrice": 3924.57472440673,
    "dom": 69.48,
    "totalCap": 135060560842.66382,
    "volumeAlt": 578834063.4608318,
    "volumeBtc": 1317521389.0430577,
    "volumeTotal": 1896355452.5038888
}
*/

func (client *Client) GetGlobalData() (*GlobalData, error) {

	resp, err := client.do("global", nil)
	if err != nil {
		return nil, fmt.Errorf("Client.do: %v", err)
	}

	res := GlobalData{}

	if err := json.Unmarshal(resp, &res); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}

	return &res, nil
}
