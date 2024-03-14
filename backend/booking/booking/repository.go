package booking

import (
	model "github.com/nekizz/final-project/backend/booking/model"
	"github.com/nekizz/final-project/backend/booking/util"
	pb "github.com/nekizz/final-project/backend/go-pbtype/booking"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreatOne(e *model.BookingV2) (*model.BookingV2, error) {
	query := r.db.Model(&model.Booking{}).Create(e)
	if err := query.Error; nil != err {
		return nil, err
	}

	return e, nil
}

func (r *Repository) ListCustomerBooking(req *pb.ListCustomerBookingRequest) ([]*model.Booking, int64, error) {
	var count int64
	var list []*model.Booking

	limit := int(req.GetLimit())
	offset := limit * int(req.GetPage()-1)

	listQuery := r.db.Model(&model.Booking{}).Select("*").Where("customer_id = ? ", req.GetCustomerId()).Limit(limit).Offset(offset)
	countQuery := r.db.Model(&model.Booking{}).Select("id").Where("customer_id = ? ", req.GetCustomerId())

	if err := countQuery.Count(&count).Error; nil != err {
		return nil, 0, err
	}

	if err := listQuery.Find(&list).Limit(limit).Offset(offset).Error; nil != err {
		return nil, 0, err
	}

	return list, count, nil
}

func (r *Repository) ListHostBooking(req *pb.ListHostBookingRequest) ([]*model.Booking, int64, error) {
	var count int64
	var list []*model.Booking

	sql := util.BuildQueryHostBooking(int(req.GetHostId()))
	limit := int(req.GetLimit())
	offset := limit * int(req.GetPage()-1)

	listQuery := r.db.Model(&model.Booking{}).Select("*").Where(sql).Limit(limit).Offset(offset)
	countQuery := r.db.Model(&model.Booking{}).Select("id").Where(sql)

	if err := countQuery.Count(&count).Error; nil != err {
		return nil, 0, err
	}
	if err := listQuery.Find(&list).Limit(limit).Offset(offset).Error; nil != err {
		return nil, 0, err
	}

	return list, count, nil
}

func (r *Repository) UpdateStatusHotel(bookingId int, status string) error {
	query := r.db.Model(model.Booking{}).Where("id = ?", bookingId).Updates(map[string]interface{}{"status": status})
	if err := query.Error; nil != err {
		return err
	}

	return nil
}
