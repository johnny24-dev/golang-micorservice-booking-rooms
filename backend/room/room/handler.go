package room

//func (s Service) UpdateRoomQuantity(msg *message.Message) ([]*message.Message, error) {
//	fmt.Println("Update quantity room service")
//	var listBookedRoomID []int
//	var replyMsgs []*message.Message
//
//	booking, pbBooking, err := util.DecodeCreateBookingCmd(msg.Payload)
//	if err != nil {
//		return nil, err
//	}
//	for _, bookedRoom := range pbBooking.GetBookedroom() {
//		listBookedRoomID = append(listBookedRoomID, int(bookedRoom.Id))
//	}
//	reply := pbB.CreateBookingResponse{
//		BookingId: uint32(booking.ID),
//		Booking:   pbBooking,
//	}
//	err = NewRepository(s.db).UpdateRoomQuantity(listBookedRoomID)
//	if err != nil {
//		reply.Success = false
//		reply.Error = err.Error()
//	} else {
//		reply.Success = true
//		reply.Error = ""
//	}
//
//	//reply.Timestamp = util.Time2pbTimestamp(time.Now())
//	payload, err := json.Marshal(&reply)
//	if err != nil {
//		return nil, err
//	}
//	replyMsg := message.NewMessage(watermill.NewUUID(), payload)
//	replyMsgs = append(replyMsgs, replyMsg)
//
//	return replyMsgs, nil
//}
//
//func (s Service) RollbackRoomQuantity(msg *message.Message) ([]*message.Message, error) {
//	var cmd pbB.RollbackCmd
//	var replyMsgs []*message.Message
//
//	if err := json.Unmarshal(msg.Payload, &cmd); err != nil {
//		return nil, err
//	}
//
//	reply := pbB.RollbackResponse{
//		CustomerId: cmd.CustomerId,
//		BookingId:  cmd.BookingId,
//	}
//	//err := NewRepository(s.db).RollbackRoomQuantity(int(cmd.BookingId))
//	//if err != nil {
//	//	reply.Success = false
//	//	reply.Error = err.Error()
//	//} else {
//	//	reply.Success = true
//	//	reply.Error = ""
//	//}
//	//reply.Timestamp = pkg.Time2pbTimestamp(time.Now())
//
//	payload, err := json.Marshal(&reply)
//	if err != nil {
//		return nil, err
//	}
//
//	replyMsg := message.NewMessage(watermill.NewUUID(), payload)
//	replyMsgs = append(replyMsgs, replyMsg)
//
//	return replyMsgs, nil
//}
//
//func (s Service) CheckRoom() {
//
//}
