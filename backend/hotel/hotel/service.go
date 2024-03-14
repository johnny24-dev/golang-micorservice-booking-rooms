package hotel

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogo/protobuf/types"
	pb "github.com/nekizz/final-project/backend/go-pbtype/hotel"
	pbi "github.com/nekizz/final-project/backend/go-pbtype/image"
	"github.com/nekizz/final-project/backend/hotel/address"
	"github.com/nekizz/final-project/backend/hotel/comment"
	"github.com/nekizz/final-project/backend/hotel/image"
	"github.com/nekizz/final-project/backend/hotel/util"
	"gorm.io/gorm"
	"log"
	"math"
	"strconv"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) Create(ctx context.Context, r *pb.Hotel) (*pb.Hotel, error) {
	if err := validateCreate(r); nil != err {
		return nil, err
	}

	if len(r.ListImage) < 5 {
		linkImageDefault := "https://chatbizfly.mediacdn.vn/2023/03/28backend_chat/_linkimagedefaultpng1679991208.png"
		for len(r.ListImage) < 5 {
			image := &pbi.Image{
				Url:     linkImageDefault,
				Type:    "hotel",
				HotelId: r.GetId(),
			}
			r.ListImage = append(r.ListImage, image)
		}
	}
	hotel, err := NewRepository(s.db).CreatOne(prepareDataToRequest(r))
	if nil != err {
		return nil, err
	}

	return prepareDataToResponse(hotel), nil
}

func (s *Service) Update(ctx context.Context, r *pb.Hotel) (*pb.Hotel, error) {
	if err := validateUpdate(r); nil != err {
		return nil, err
	}

	tx := s.db.Begin()

	hotel, err := NewRepository(tx).UpdateOne(int(r.GetId()), prepareDataToRequest(r))
	if nil != err {
		tx.Rollback()
		return nil, err
	}

	if r.GetAddress() != nil {
		_, err := address.UpdateAddressesWorker(r.GetAddress(), tx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if r.GetListImage() != nil && len(r.GetListImage()) == 5 {
		listImageChange := r.GetListImage()
		listImage, err := image.GetImageWorker(int(r.GetId()), s.db)
		if err != nil || len(listImage) == 0 {
			return nil, err
		}
		for idx, img := range listImageChange {
			if listImage[idx].Url != listImageChange[idx].Url {
				_, err := image.UpdateImageWorker(img, tx)
				if err != nil {
					tx.Rollback()
					return nil, err
				}

			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return prepareDataToResponse(hotel), nil
}

func (s *Service) Get(ctx context.Context, r *pb.OneHotelRequest) (*pb.Hotel, error) {
	if err := validateOne(r); nil != err {
		return nil, err
	}

	id, _ := strconv.Atoi(r.GetId())
	hotel, err := NewRepository(s.db).FindOne(id)
	if err != nil || hotel == nil {
		return nil, err
	}

	if hotel.IsDeleted == 0 {
		return nil, errors.New("Hotel doesn't exist")
	}

	address, err := address.NewRepository(s.db).FindOne(hotel.AddressID)
	if nil != err {
		return nil, err
	}
	hotel.Address = *address

	listRate, err := comment.CalculateRateOfHotel(int(hotel.ID), s.db)
	if err != nil {
		return nil, err
	}
	hotelPb := prepareDataToResponse(hotel)
	rate, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", util.CaulculateAverange(listRate)), 32)
	hotelPb.Rate = float32(rate)

	return hotelPb, nil
}

func (s *Service) ListHotelByHostID(ctx context.Context, r *pb.ListHotelRequest) (*pb.ListHotelResponse, error) {
	if err := validateList(r.GetPage(), r.GetLimit()); nil != err {
		return nil, err
	}

	var list []*pb.Hotel

	hotels, count, err := NewRepository(s.db).ListAll(r)
	if nil != err {
		return nil, err
	}

	mapping, err := NewRepository(s.db).listAllAddressByObject(int(r.GetCustomerId()))
	if err != nil {
		return nil, err
	}

	for _, hotel := range hotels {
		hotel.Address = mapping[uint(hotel.AddressID)]
		list = append(list, prepareDataToResponse(hotel))
	}

	return &pb.ListHotelResponse{
		Items:      list,
		TotalCount: uint32(count),
		Page:       r.GetPage(),
		Limit:      r.GetLimit(),
		MaxPage:    uint32(math.Ceil(float64(uint32(count)) / float64(r.GetLimit()))),
	}, nil
}

func (s *Service) ListAll(ctx context.Context, r *pb.ListHotelCustomerRequest) (*pb.ListHotelResponse, error) {
	if err := validateList(r.GetPage(), r.GetLimit()); nil != err {
		return nil, err
	}

	var list []*pb.Hotel
	mappingSearch := map[string]string{}

	if r.GetName() != "" {
		mappingSearch["name"] = r.GetName()
	}
	if r.GetLocation() != "" {
		mappingSearch["location"] = r.GetLocation()
	}
	if r.GetRoomType() != "" {
		mappingSearch["room_type"] = r.GetRoomType()
	}
	if r.GetCheckIn() != "" {
		mappingSearch["check_in"] = r.GetCheckIn()
	}
	if r.GetCheckout() != "" {
		mappingSearch["check_out"] = r.GetCheckout()
	}
	if r.GetStarLevel() != "" {
		mappingSearch["star_level"] = r.GetStarLevel()
	}

	hotels, count, err := NewRepository(s.db).ListAllTwo(r, mappingSearch)
	if nil != err {
		return nil, err
	}

	for _, hotel := range hotels {
		listRate, err := comment.CalculateRateOfHotel(int(hotel.ID), s.db)
		if err != nil {
			log.Println(err)
			continue
		}
		hotelPb := prepareDataToResponse(hotel)
		rate, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", util.CaulculateAverange(listRate)), 32)
		hotelPb.Rate = float32(rate)
		list = append(list, hotelPb)
	}

	return &pb.ListHotelResponse{
		Items:      list,
		TotalCount: uint32(count),
		Page:       r.GetPage(),
		Limit:      r.GetLimit(),
		MaxPage:    uint32(math.Ceil(float64(uint32(count)) / float64(r.GetLimit()))),
	}, nil
}

func (s *Service) GetByListID(ctx context.Context, r *pb.ListHotelByIDRequest) (*pb.ListHotelResponse, error) {
	var list []*pb.Hotel

	hotels, count, err := NewRepository(s.db).ListHotelByIDs(r)
	if nil != err {
		return nil, err
	}

	for _, hotel := range hotels {
		list = append(list, prepareDataToResponse(hotel))
	}

	return &pb.ListHotelResponse{
		Items:      list,
		TotalCount: uint32(count),
		Page:       r.GetPage(),
		Limit:      r.GetLimit(),
		MaxPage:    uint32(math.Ceil(float64(uint32(count)) / float64(r.GetLimit()))),
	}, nil
}

func (s *Service) Delete(ctx context.Context, r *pb.OneHotelRequest) (*types.Empty, error) {
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
