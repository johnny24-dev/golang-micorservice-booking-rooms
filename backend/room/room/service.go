package room

import (
	"context"
	"errors"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/gogo/protobuf/types"
	pb "github.com/nekizz/final-project/backend/go-pbtype/room"
	"gorm.io/gorm"
	"math"
	"strconv"
)

type Service struct {
	publisher *kafka.Publisher
	db        *gorm.DB
}

func NewService(db *gorm.DB, publisher *kafka.Publisher) *Service {
	return &Service{
		db:        db,
		publisher: publisher,
	}
}

func (s *Service) Create(ctx context.Context, r *pb.Room) (*pb.Room, error) {
	if err := validateCreate(r); nil != err {
		return nil, err
	}

	hotel, err := NewRepository(s.db).CreatOne(prepareDataToCreateRoom(r))
	if nil != err {
		return nil, err
	}

	return prepareDataToResponse(hotel), nil
}

func (s *Service) Update(ctx context.Context, r *pb.Room) (*pb.Room, error) {
	if err := validateUpdate(r); nil != err {
		return nil, err
	}

	hotel, err := NewRepository(s.db).UpdateOne(int(r.GetId()), prepareDataToRequest(r))

	if nil != err {
		return nil, err
	}

	return prepareDataToResponse(hotel), nil
}

func (s *Service) Get(ctx context.Context, r *pb.OneRoomRequest) (*pb.Room, error) {
	if err := validateOne(r); nil != err {
		return nil, err
	}

	id, _ := strconv.Atoi(r.GetId())
	customer, err := NewRepository(s.db).FindOne(id)
	if err != nil {
		return nil, err
	}

	return prepareDataToResponse(customer), nil
}

func (s *Service) Delete(ctx context.Context, r *pb.OneRoomRequest) (*types.Empty, error) {
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

func (s *Service) List(ctx context.Context, r *pb.ListRoomRequest) (*pb.ListRoomResponse, error) {
	if err := validateList(r); nil != err {
		return nil, err
	}

	var list []*pb.Room

	rooms, count, err := NewRepository(s.db).ListAll(r)
	if nil != err {
		return nil, err
	}

	for _, room := range rooms {
		list = append(list, prepareDataToResponse(room))
	}

	return &pb.ListRoomResponse{
		Items:      list,
		TotalCount: uint32(count),
		Page:       r.GetPage(),
		Limit:      r.GetLimit(),
		MaxPage:    uint32(math.Ceil(float64(uint32(count)) / float64(r.GetLimit()))),
	}, nil
}

func (s *Service) CheckRoomQuantity(ctx context.Context, r *pb.CheckingRoomRequest) (*types.Empty, error) {
	err := NewRepository(s.db).UpdateRoomCurrentQuantity()
	if nil != err {
		return nil, err
	}

	return &types.Empty{}, nil
}

func (s *Service) ListRoomByBooked(ctx context.Context, req *pb.ListRoomByBookedRoomRequest) (*pb.ListRoomByBookedRoomResponse, error) {
	if len(req.GetListRoomId()) == 0 {
		return nil, errors.New("Invalid list room id")
	}

	var listRoomPb []*pb.Room

	rooms, err := NewRepository(s.db).GetRoomByBooked(req.GetListRoomId())
	if nil != err {
		return nil, err
	}

	for _, room := range rooms {
		listRoomPb = append(listRoomPb, prepareDataToResponse(room))
	}

	return &pb.ListRoomByBookedRoomResponse{
		Items: listRoomPb,
	}, nil
}
