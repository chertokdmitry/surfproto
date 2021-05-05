package cameras

import (
	"gitlab.com/chertokdmitry/surfproto/src/services"
	"log"
)

func GetBySpot(spotId int64) []int64 {
	cameras, err := services.CamerasService.GetBySpot(spotId)
	if err != nil {
		log.Println(err)
	}

	return cameras
}
