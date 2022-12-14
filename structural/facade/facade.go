package structural

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// The Facade design pattern is used to hide the complexity scope from the user
// and shield the code form unwanted access.
type CurrentWeatherDataRetriever interface {
	GetByCityAndCountryCode(city, countryCode string) (*Weather, error)
	GetByGeoCoordinates(lat, lon float32) (*Weather, error)
}

type CurrentWeatherData struct {
	APIkey string
}

func (p *CurrentWeatherData) responseParser(body io.Reader) (*Weather, error) {
	var w Weather
	err := json.NewDecoder(body).Decode(&w)
	if err != nil {
		return nil, err
	}
	return &w, nil
}

func (p *CurrentWeatherData) GetByCityAndCountryCode(city, countryCode string) (*Weather, error) {
	r := getMockData()

	w, err := p.responseParser(r)
	if err != nil {
		return nil, err
	}

	if w.Name == city && w.Sys.Coutry == countryCode {
		return w, nil
	}

	return nil, errors.New("not found")
}

func (p *CurrentWeatherData) GetByGeoCoordinates(lat, lon float32) (*Weather, error) {
	r := getMockData()

	w, err := p.responseParser(r)
	if err != nil {
		return nil, err
	}

	if w.Coord.Lat == lat && w.Coord.Lon == lon {
		return w, nil
	}

	return nil, errors.New("not found")

}

func getMockData() io.Reader {
	response := `{"coord":{"lon":-3.7, "lat":40.42},"weather":[{"id":803,"main":"Clouds","description":"broken clouds","icon":"04n"}],"base":"stations","main":{"temp":303.56,"pressure":1016.46,"humidity":26.8,"temp_min":300.95,"temp_max":305.93},"wind":{"speed":3.17,"deg":151.001},"rain":{"3h":0.0075},"clouds":{"all":68},"dt":1471295823,"sys":{"type":3,"id":1442829648,"message":0.0278,"country":"ES","sunrise":1471238808,"sunset":1471288232},"id":3117735,"name":"Madrid","cod":200}`
	r := bytes.NewReader([]byte(response))
	return r
}

type Weather struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Cod   int    `json:"cod"`
	Coord struct {
		Lon float32 `json:"lon"`
		Lat float32 `json:"lat"`
	} `json:"coords"`
	Weather []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp     float32 `json:"temp"`
		Pressure float32 `json:"pressure"`
		Humidity float32 `json:"humidity"`
		TempMin  float32 `json:"temp_min"`
		TempMax  float32 `json:"temp_max"`
	} `json:"main"`
	Wind struct {
		Speed float32 `json:"speed"`
		Deg   float32 `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Rain struct {
		ThreeHours float32 `json:"3h"`
	} `json:"rain"`
	Dt  uint32 `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float32 `json:"message"`
		Coutry  string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
}
