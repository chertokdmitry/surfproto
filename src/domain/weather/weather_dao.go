package weather

import (
	"gitlab.com/chertokdmitry/surfproto/src/db_resource/weather_db"
	"gitlab.com/chertokdmitry/surfproto/src/logger"
	"gitlab.com/chertokdmitry/surfproto/src/utils/errors"
)

const querySelectWeatherBySpot = "SELECT spots.title, weather.hourly FROM spots LEFT JOIN  weather ON spots.id =  weather.spot_id WHERE spots.id = $1 ORDER BY weather.created_at DESC limit 1"

func (weather *Weather) Get() *errors.RestErr {
	db := db_resource.GetDB()
	defer db.Close()

	result := db.QueryRow(querySelectWeatherBySpot, int(weather.SpotId))
	if err := result.Scan(&weather.Title, &weather.Hourly); err != nil {
		logger.Error("error when trying to get weather", err)
		return errors.NewInternalServerError("error when trying to get weather")
	}

	return nil
}
