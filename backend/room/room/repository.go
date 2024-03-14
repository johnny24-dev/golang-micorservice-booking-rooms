package room

import (
	"fmt"
	pb "github.com/nekizz/final-project/backend/go-pbtype/room"
	"github.com/nekizz/final-project/backend/room/model"
	"gorm.io/gorm"
	"time"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreatOne(e *model.Room) (*model.Room, error) {
	fmt.Println(e.HotelID)
	query := r.db.Model(&model.Room{}).Create(e)
	if err := query.Error; nil != err {
		return nil, err
	}

	return e, nil
}

func (r *Repository) UpdateOne(id int, c *model.Room) (*model.Room, error) {
	query := r.db.Model(&model.Room{}).Where("id = ?", id).UpdateColumns(c)
	if err := query.Error; nil != err {
		return nil, err
	}

	return c, nil
}

func (r *Repository) FindOne(id int) (*model.Room, error) {
	var room model.Room

	query := r.db.Model(&model.Room{}).Where("id = ? AND is_deleted != 1", id).Preload("Amenity").Find(&room)
	if err := query.Error; nil != err {
		return nil, err
	}

	return &room, nil
}

func (r *Repository) DeleteOne(id int) error {
	query := r.db.Model(&model.Room{}).Where("id = ?", id).Updates(map[string]interface{}{"is_deleted": 1})
	if err := query.Error; nil != err {
		return err
	}

	return nil
}

func (r *Repository) ListAll(req *pb.ListRoomRequest) ([]*model.Room, int64, error) {
	var count int64
	var list []*model.Room

	limit := int(req.GetLimit())
	offset := limit * int(req.GetPage()-1)

	listQuery := r.db.Model(&model.Room{}).Select("*").Limit(limit).Offset(offset).Where("hotel_id = ? AND current_quantity > 0 AND is_deleted != 1", req.HotelId).Preload("Amenity")
	countQuery := r.db.Model(&model.Room{}).Select("id").Where("hotel_id = ? AND current_quantity > 0", req.HotelId)

	if err := countQuery.Count(&count).Error; nil != err {
		return nil, 0, err
	}

	if err := listQuery.Find(&list).Limit(limit).Offset(offset).Error; nil != err {
		return nil, 0, err
	}

	return list, count, nil
}

func (r *Repository) CheckingRoomQuantity(req *pb.ListRoomRequest) ([]*model.Room, int64, error) {
	var count int64
	var list []*model.Room

	limit := int(req.GetLimit())
	offset := limit * int(req.GetPage()-1)

	listQuery := r.db.Model(&model.Room{}).Select("*").Limit(limit).Offset(offset).Where("hotel_id = ? AND current_quantity > 0 AND is_deleted != 1", req.HotelId).Preload("Amenity")
	countQuery := r.db.Model(&model.Room{}).Select("id").Where("hotel_id = ? AND current_quantity > 0", req.HotelId)

	if err := countQuery.Count(&count).Error; nil != err {
		return nil, 0, err
	}

	if err := listQuery.Find(&list).Limit(limit).Offset(offset).Error; nil != err {
		return nil, 0, err
	}

	return list, count, nil
}

func (r *Repository) UpdateRoomQuantity(listRoomID []int) error {
	query := r.db.Model(&model.Room{}).Where("id IN ?", listRoomID).Updates("CurrentQuantity = CurrentQuantity - 1")
	if err := query.Error; nil != err {
		return err
	}

	return nil
}

func (r *Repository) UpdateRoomCurrentQuantity() error {
	var listHotelId []int

	//start transaction

	checkingTime := time.Now()

	queryBR := r.db.Table("booked_rooms").Select("room_id").Where("check_in <= ? AND check_out >= ?", checkingTime, checkingTime).Scan(&listHotelId)
	if err := queryBR.Error; err != nil {
		return err
	}

	if len(listHotelId) > 0 {
		mapping := map[int]int{}

		for _, val := range listHotelId {
			mapping[val] += 1
		}

		tx := r.db.Begin()

		for idx, val := range mapping {
			updateStr := fmt.Sprintf("current_quantity = quantity - %s", val)
			query := tx.Model(&model.Room{}).Where("id = ?", idx).Updates(updateStr)
			if err := query.Error; nil != err {
				tx.Rollback()
				return err
			}
		}

		if err := tx.Commit().Error; err != nil {
			return err
		}
	}

	return nil
}

func (r *Repository) GetRoomByBooked(listRoomID []uint32) ([]*model.Room, error) {
	var listRoom []*model.Room

	queryBR := r.db.Model(&model.Room{}).Where("id IN ?", listRoomID).Find(&listRoom)
	if err := queryBR.Error; err != nil {
		return nil, err
	}

	return listRoom, nil
}
