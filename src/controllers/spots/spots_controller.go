package spots

import (
	"gitlab.com/chertokdmitry/surfproto/src/domain/spots"
	"gitlab.com/chertokdmitry/surfproto/src/services"
	"log"
)

func GetByRegionId(regionId int64) []spots.Spot {
	spots, err := services.SpotsService.GetByRegionId(regionId)
	if err != nil {
		log.Println(err)
	}

	return spots
}

func GetAll() []spots.Spot {
	spots, err := services.SpotsService.GetAll()
	if err != nil {
		log.Println(err)
	}

	return spots
}
