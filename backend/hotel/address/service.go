package address

import (
	"context"
	pb "github.com/nekizz/final-project/backend/go-pbtype/address"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s Service) Update(ctx context.Context, r *pb.Address) (*pb.Address, error) {
	if err := validateUpdate(r); nil != err {
		return nil, err
	}

	address, err := NewRepository(s.db).UpdateOne(int(r.GetId()), PrepareAddressToRequest(r))
	if nil != err {
		return nil, err
	}

	return PrepareDataToResponse(address), nil
}

func (s Service) Create(ctx context.Context, r *pb.Address) (*pb.Address, error) {
	if err := validateUpdate(r); nil != err {
		return nil, err
	}

	address, err := NewRepository(s.db).CreatOne(PrepareAddressToRequest(r))
	if nil != err {
		return nil, err
	}

	return PrepareDataToResponse(address), nil
}

func (s Service) Get(ctx context.Context, r *pb.OneAddressRequest) (*pb.Address, error) {
	if err := validateOne(r); nil != err {
		return nil, err
	}

	address, err := NewRepository(s.db).FindOne(int(r.GetId()))
	if nil != err {
		return nil, err
	}

	return PrepareDataToResponse(address), nil
}
