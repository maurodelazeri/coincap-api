package coincap

import (
	"encoding/json"
	"fmt"
)

type Ticks []*Tick

type Tick struct {
	Short       string  `json:"short"`
	Long        float64 `json:"long"`
	Mktcap      float64 `json:"mktcap"`
	Perc        float64 `json:"perc"`
	Price       float64 `json:"price"`
	ShapeShift  float64 `json:"shapeshift"`
	Supply      int64   `json:"supply"`
	UsdVolume   float64 `json:"usdVolume"`
	Volume      float64 `json:"volume"`
	VwapData    float64 `json:"vwapData"`
	VwapDataBTC float64 `json:"vwapDataBTC"`
}

// coincap API implementation of Ticker endpoint.
//
// Endpoint: /front/
// Method: GET
//
/*
[
    {
        "cap24hrChange": -6.05,
        "long": "Bitcoin",
        "mktcap": 65173805891.25,
        "perc": -6.05,
        "price": 3934.85,
        "shapeshift": true,
        "short": "BTC",
        "supply": 16563225,
        "usdVolume": 2337600000,
        "volume": 2337600000,
        "vwapData": 3997.5639538606733,
        "vwapDataBTC": 3997.5639538606733
    },
    {
        "cap24hrChange": -6.59,
        "long": "Ethereum",
        "mktcap": 26016428866.32,
        "perc": -6.59,
        "price": 275.02,
        "shapeshift": true,
        "short": "ETH",
        "supply": 94598316,
        "usdVolume": 945732000,
        "volume": 945732000,
        "vwapData": 278.03921067242516,
        "vwapDataBTC": 278.03921067242516
    }
]
*/

func (client *Client) GetTickers() (Ticks, error) {

	resp, err := client.do("front", nil)
	if err != nil {
		return nil, fmt.Errorf("Client.do: %v", err)
	}

	res := make(Ticks, 0)

	if err := json.Unmarshal(resp, &res); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %v", err)
	}

	return res, nil
}
