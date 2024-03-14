package booked_room

import (
	"context"
	"errors"
	pbBR "github.com/nekizz/final-project/backend/go-pbtype/bookedroom"
	pbh "github.com/nekizz/final-project/backend/go-pbtype/hotel"
	pbr "github.com/nekizz/final-project/backend/go-pbtype/room"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"strconv"
)

func GetIDRoomByBookingID(bookingID int, db *gorm.DB) ([]*pbBR.BookedRoom, float32, *pbh.Hotel, error) {
	var listRoomID []uint32
	var listPbBookedRoom []*pbBR.BookedRoom
	var total, totalNotVat float32

	listBookedRoom, err := NewRepository(db).GetListBookedRoomOfBooking(bookingID)
	if err != nil {
		return nil, 0, nil, err
	}

	if len(listBookedRoom) == 0 {
		return nil, 0, nil, errors.New("No BookedRoom")
	}

	for _, bookedRoom := range listBookedRoom {
		bookedDay := float32((bookedRoom.CheckOut.Sub(*bookedRoom.CheckIn)).Hours() / 24)
		totalNotVat += (bookedRoom.Price * float32(bookedRoom.Quantity)) * bookedDay
		total = totalNotVat + totalNotVat*0.1
		listRoomID = append(listRoomID, uint32(bookedRoom.RoomID))
	}

	roomAddr := viper.GetString("room_backend")
	roomConn, err := grpc.Dial(roomAddr, grpc.WithInsecure())
	if err != nil {
		return nil, 0, nil, err
	}

	roomReq := &pbr.ListRoomByBookedRoomRequest{
		ListRoomId: listRoomID,
	}
	roomService := pbr.NewRoomServiceClient(roomConn)
	listRoom, err := roomService.ListRoomByBooked(context.Background(), roomReq)
	if err != nil {
		return nil, 0, nil, err
	}

	for _, bookedRoom := range listBookedRoom {
		for _, room := range listRoom.GetItems() {
			if bookedRoom.RoomID == uint(room.Id) {
				bookedRoomPb := prepareBookingToResponse(bookedRoom)
				bookedRoomPb.Room = room
				listPbBookedRoom = append(listPbBookedRoom, bookedRoomPb)
			}
		}
	}

	if listPbBookedRoom[0].GetRoom().GetHotelId() == 0 {
		return nil, 0, nil, err
	}

	hotelAddr := viper.GetString("hotel_backend")
	hotelConn, err := grpc.Dial(hotelAddr, grpc.WithInsecure())
	if err != nil {
		return nil, 0, nil, err
	}

	hotelReq := &pbh.OneHotelRequest{
		Id: strconv.Itoa(int(listPbBookedRoom[0].GetRoom().GetHotelId())),
	}

	hotelService := pbh.NewHotelServiceClient(hotelConn)
	hotel, err := hotelService.Get(context.Background(), hotelReq)
	if err != nil {
		return nil, 0, nil, err
	}

	return listPbBookedRoom, total, hotel, nil
}
