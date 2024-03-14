package amenity

import (
	pb "github.com/nekizz/final-project/backend/go-pbtype/amenity"
	"github.com/nekizz/final-project/backend/room/model"
	"gorm.io/gorm"
)

func PrepareAmenityToResponse(p *model.Amenity) *pb.Amenity {
	return &pb.Amenity{
		Id:   int32(p.ID),
		Name: p.Name,
	}
}

func PrepareAmenityToRequest(p *pb.Amenity) *model.Amenity {
	return &model.Amenity{
		Model: gorm.Model{},
		Name:  p.Name,
	}
}
