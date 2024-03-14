package wishlist

import (
	"context"
	"fmt"
	"github.com/gogo/protobuf/types"
	pbh "github.com/nekizz/final-project/backend/go-pbtype/hotel"
	pb "github.com/nekizz/final-project/backend/go-pbtype/wishlist"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"math"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s Service) Create(ctx context.Context, r *pb.Wishlist) (*pb.Wishlist, error) {
	if err := validateCreate(r); nil != err {
		return nil, err
	}

	hotel, err := NewRepository(s.db).CreatOne(prepareDataToRequest(r))
	if nil != err {
		return nil, err
	}

	return prepareDataToResponse(hotel), nil
}

func (s Service) Delete(ctx context.Context, r *pb.Wishlist) (*types.Empty, error) {
	if err := validateDelete(r); nil != err {
		return nil, err
	}

	err := NewRepository(s.db).DeleteOne(int(r.HotelId), int(r.GetCustomerId()))
	if nil != err {
		return nil, err
	}

	return &types.Empty{}, nil
}

func (s Service) List(ctx context.Context, r *pb.ListWishlistRequest) (*pb.ListWishlistResponse, error) {
	if err := validateList(r); nil != err {
		return nil, err
	}

	listWishlistID, count, err := NewRepository(s.db).ListWishlistIdOfCustomer(int(r.GetCustomerId()))
	if nil != err {
		return nil, err
	}
	fmt.Println(listWishlistID)
	grpcAddr := viper.GetString("hotel_backend")
	hotelConn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	hotelService := pbh.NewHotelServiceClient(hotelConn)
	hotelReq := &pbh.ListHotelByIDRequest{
		Page:        r.GetPage(),
		Limit:       r.GetLimit(),
		ListHotelId: convertIntToInt32(listWishlistID),
	}
	fmt.Println(2)
	listHotel, err := hotelService.GetByListID(context.Background(), hotelReq)
	if err != nil {
		return nil, err
	}
	fmt.Println(listHotel)
	fmt.Println(4)

	return &pb.ListWishlistResponse{
		Items:      listHotel.GetItems(),
		TotalCount: uint32(count),
		Page:       r.GetPage(),
		Limit:      r.GetLimit(),
		MaxPage:    uint32(math.Ceil(float64(uint32(count)) / float64(r.GetLimit()))),
	}, nil
}
