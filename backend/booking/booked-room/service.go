package booked_room

import (
	"context"
	"github.com/gogo/protobuf/types"
	pbBK "github.com/nekizz/final-project/backend/go-pbtype/bookedroom"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s Service) Create(ctx context.Context, r *pbBK.BookedRoom) (*pbBK.BookedRoom, error) {
	return &pbBK.BookedRoom{}, nil
}

func (s Service) Update(ctx context.Context, r *pbBK.BookedRoom) (*pbBK.BookedRoom, error) {
	return &pbBK.BookedRoom{}, nil
}

func (s Service) Get(ctx context.Context, r *pbBK.OneBookedRoomRequest) (*pbBK.BookedRoom, error) {

	return &pbBK.BookedRoom{}, nil
}

func (s Service) List(ctx context.Context, r *pbBK.ListBookedRoomRequest) (*pbBK.ListBookedRoomResponse, error) {
	return &pbBK.ListBookedRoomResponse{}, nil
}

func (s Service) Delete(ctx context.Context, r *pbBK.OneBookedRoomRequest) (*types.Empty, error) {
	return &types.Empty{}, nil
}
