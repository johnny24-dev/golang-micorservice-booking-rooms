package comment

import (
	pb "github.com/nekizz/final-project/backend/go-pbtype/comment"
	"github.com/nekizz/final-project/backend/hotel/model"
	"gorm.io/gorm"
)

func prepareCommentToResponse(p *model.Comment) *pb.Comment {
	return &pb.Comment{
		Id:          int32(p.ID),
		Text:        p.Text,
		Type:        p.Type,
		Rate:        p.Rate,
		CommentDate: p.CreatedAt.Format("2006-01-02 15:04:05"),
		HotelId:     int32(p.HotelID),
		CustomerId:  int32(p.CustomerID),
	}
}

func prepareCommentToRequest(p *pb.Comment) *model.Comment {
	return &model.Comment{
		Model:      gorm.Model{},
		Text:       p.Text,
		Type:       p.Type,
		Rate:       p.Rate,
		HotelID:    int(p.GetHotelId()),
		CustomerID: int(p.GetCustomerId()),
	}
}
