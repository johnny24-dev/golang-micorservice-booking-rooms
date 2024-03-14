package wishlist

import (
	pb "github.com/nekizz/final-project/backend/go-pbtype/wishlist"
	"github.com/nekizz/final-project/backend/user/model"
)

func prepareDataToResponse(p *model.Wishlist) *pb.Wishlist {
	return &pb.Wishlist{
		HotelId:    int32(p.HotelID),
		CustomerId: int32(p.CustomerID),
	}
}

func prepareDataToRequest(p *pb.Wishlist) *model.Wishlist {
	return &model.Wishlist{
		HotelID:    uint(p.HotelId),
		CustomerID: uint(p.CustomerId),
	}
}

func convertIntToInt32(listInt []int) []int32 {
	var listInt32 []int32

	for _, value := range listInt {
		listInt32 = append(listInt32, int32(value))
	}

	return listInt32
}
