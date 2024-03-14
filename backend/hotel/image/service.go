package image

import (
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

//func (s Service) Delete(ctx context.Context, r *pb.OneHotelRequest) (*types.Empty, error) {
//	if err := validateOne(r); nil != err {
//		return nil, err
//	}
//
//	id, _ := strconv.Atoi(r.GetId())
//	err := NewRepository(s.db).DeleteOne(id)
//	if nil != err {
//		return nil, err
//	}
//
//	return &types.Empty{}, nil
//}
