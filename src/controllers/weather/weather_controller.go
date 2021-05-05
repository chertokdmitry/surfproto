package weather

import (
	"gitlab.com/chertokdmitry/surfproto/src/domain/weather"
	"gitlab.com/chertokdmitry/surfproto/src/services"
	"log"
)

func Get(spotId int64) *weather.Weather {

	weather, err := services.WeatherService.GetWeather(spotId)
	if err != nil {
		log.Println(err)
	}

	return weather
}
