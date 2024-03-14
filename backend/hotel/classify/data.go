package classify

import (
	pb "github.com/nekizz/final-project/backend/go-pbtype/classify"
	"github.com/nekizz/final-project/backend/hotel/model"
	"strconv"
)

func prepareClassifyToResponse(p *model.Classify) *pb.Classify {
	return &pb.Classify{
		Id:                strconv.Itoa(p.ID),
		Emotion:           strconv.Itoa(p.Emotion),
		Service:           strconv.Itoa(p.Service),
		AmenityAndLeisure: strconv.Itoa(p.AmenityAndLeisure),
		Room:              strconv.Itoa(p.Room),
		Location:          strconv.Itoa(p.Location),
		Cuisine:           strconv.Itoa(p.Cuisine),
		Cost:              strconv.Itoa(p.Cost),
	}
}
