package service

import (
	"context"
	"github.com/gogo/protobuf/types"
	pb "github.com/nekizz/final-project/backend/go-pbtype/service"
	"gorm.io/gorm"
	"math"
	"strconv"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s Service) Create(ctx context.Context, r *pb.Service) (*pb.Service, error) {
	if err := validateCreate(r); nil != err {
		return nil, err
	}

	hotel, err := NewRepository(s.db).CreatOne(prepareDataToRequest(r))

	if nil != err {
		return nil, err
	}

	return prepareDataToResponse(hotel), nil
}

func (s Service) Update(ctx context.Context, r *pb.Service) (*pb.Service, error) {
	if err := validateUpdate(r); nil != err {
		return nil, err
	}

	hotel, err := NewRepository(s.db).UpdateOne(int(r.GetId()), prepareDataToRequest(r))

	if nil != err {
		return nil, err
	}

	return prepareDataToResponse(hotel), nil
}

func (s Service) Get(ctx context.Context, r *pb.OneServiceRequest) (*pb.Service, error) {
	if err := validateOne(r); nil != err {
		return nil, err
	}

	id, _ := strconv.Atoi(r.GetId())
	employee, err := NewRepository(s.db).FindOne(id)

	if nil != err {
		return nil, err
	}

	return prepareDataToResponse(employee), nil
}

func (s Service) List(ctx context.Context, r *pb.ListServiceRequest) (*pb.ListServiceResponse, error) {
	var list []*pb.Service
	if err := validateList(r); nil != err {
		return nil, err
	}

	employee, count, err := NewRepository(s.db).ListAll(r)

	if nil != err {
		return nil, err
	}
	for i := range employee {
		list = append(list, prepareDataToResponse(employee[i]))
	}

	return &pb.ListServiceResponse{
		Items:      list,
		TotalCount: uint32(count),
		Page:       r.GetPage(),
		Limit:      r.GetLimit(),
		MaxPage:    uint32(math.Ceil(float64(uint32(count)) / float64(r.GetLimit()))),
	}, nil
}

func (s Service) Delete(ctx context.Context, r *pb.OneServiceRequest) (*types.Empty, error) {
	if err := validateOne(r); nil != err {
		return nil, err
	}

	id, _ := strconv.Atoi(r.GetId())
	err := NewRepository(s.db).DeleteOne(id)
	if nil != err {
		return nil, err
	}

	return &types.Empty{}, nil
}
