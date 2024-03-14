package hotel

import (
	pbc "github.com/nekizz/final-project/backend/go-pbtype/comment"
	pb "github.com/nekizz/final-project/backend/go-pbtype/hotel"
	pb2 "github.com/nekizz/final-project/backend/go-pbtype/image"
	address "github.com/nekizz/final-project/backend/hotel/address"
	image "github.com/nekizz/final-project/backend/hotel/image"
	"github.com/nekizz/final-project/backend/hotel/model"
	"gorm.io/gorm"
	"time"
)

func prepareDataToResponse(p *model.Hotel) *pb.Hotel {
	var listImage []*pb2.Image
	var listComment []*pbc.Comment

	for _, img := range p.Image {
		listImage = append(listImage, image.PrepareImageToResponse(&img))
	}
	if p.Address.ID == 0 {
		return &pb.Hotel{
			Id:          int32(p.ID),
			Name:        p.Name,
			StarLevel:   int32(p.StarLevel),
			Rate:        p.Rate,
			Description: p.Description,
			IsAvailable: p.IsAvailable,
			ListImage:   listImage,
			ListComment: listComment,
			Address:     nil,
		}
	}
	address := address.PrepareDataToResponse(&p.Address)
	return &pb.Hotel{
		Id:          int32(p.ID),
		Name:        p.Name,
		StarLevel:   int32(p.StarLevel),
		Rate:        p.Rate,
		Description: p.Description,
		IsAvailable: p.IsAvailable,
		ListImage:   listImage,
		ListComment: listComment,
		Address:     address,
	}
}

func prepareDataToRequest(p *pb.Hotel) *model.Hotel {
	var listImage []model.Image

	if len(p.ListImage) > 0 {
		for _, img := range p.ListImage {
			listImage = append(listImage, *image.PrepareImageToRequest(img))
		}
	}
	address := address.PrepareAddressToRequest(p.Address)

	return &model.Hotel{
		Model:       gorm.Model{},
		Name:        p.Name,
		StarLevel:   int(p.StarLevel),
		Description: p.Description,
		IsAvailable: p.IsAvailable,
		Address:     *address,
		Image:       listImage,
		AddressID:   int(p.AddressId),
		CustomerID:  int(p.GetCustomerId()),
		IsDeleted:   1,
	}
}

func getModel(id uint, c *model.Hotel) *model.Hotel {
	c.ID = id
	return &model.Hotel{
		Model: gorm.Model{
			ID:        id,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		Name:        c.Name,
		StarLevel:   c.StarLevel,
		Rate:        c.Rate,
		IsAvailable: c.IsAvailable,
		Description: c.Description,
		AddressID:   c.AddressID,
		CustomerID:  c.CustomerID,
		Address:     c.Address,
		Image:       c.Image,
		Comment:     c.Comment,
	}
}
