package room

import (
	pba "github.com/nekizz/final-project/backend/go-pbtype/amenity"
	pb "github.com/nekizz/final-project/backend/go-pbtype/room"
	amenity2 "github.com/nekizz/final-project/backend/room/amenity"
	"github.com/nekizz/final-project/backend/room/model"
	"gorm.io/gorm"
)

func prepareDataToResponse(p *model.Room) *pb.Room {
	var listAmenity []*pba.Amenity

	for _, img := range p.Amenity {
		listAmenity = append(listAmenity, amenity2.PrepareAmenityToResponse(&img))
	}
	return &pb.Room{
		Id:          int32(p.ID),
		Name:        p.Name,
		Type:        p.Type,
		Price:       p.Price,
		Description: p.Description,
		ListAmenity: listAmenity,
		Quantity:    int32(p.Quantity),
		Status:      p.Status,
		HotelId:     p.HotelID,
	}
}

func prepareDataToRequest(p *pb.Room) *model.Room {
	var listAmenity []model.Amenity

	for _, img := range p.ListAmenity {
		listAmenity = append(listAmenity, *amenity2.PrepareAmenityToRequest(img))
	}
	return &model.Room{
		Model:       gorm.Model{},
		Name:        p.Name,
		Type:        p.Type,
		Price:       p.Price,
		Status:      p.Status,
		Description: p.Description,
		Amenity:     listAmenity,
		Quantity:    uint(p.Quantity),
		HotelID:     p.HotelId,
	}
}

func prepareDataToCreateRoom(p *pb.Room) *model.Room {
	var listAmenity []model.Amenity

	for _, img := range p.ListAmenity {
		listAmenity = append(listAmenity, *amenity2.PrepareAmenityToRequest(img))
	}
	return &model.Room{
		Model:           gorm.Model{},
		Name:            p.Name,
		Type:            p.Type,
		Price:           p.Price,
		Status:          "true",
		Description:     p.Description,
		HotelID:         p.HotelId,
		Quantity:        uint(p.Quantity),
		CurrentQuantity: uint(p.Quantity),
		Amenity:         listAmenity,
	}
}
