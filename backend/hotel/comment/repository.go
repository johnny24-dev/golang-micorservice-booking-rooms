package comment

import (
	pb "github.com/nekizz/final-project/backend/go-pbtype/comment"
	"github.com/nekizz/final-project/backend/hotel/model"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreatOne(e *model.Comment) (*model.Comment, error) {
	query := r.db.Model(&model.Comment{}).Omit("Customer, Name, Classify").Create(&e)
	if err := query.Error; nil != err {
		return nil, err
	}

	return e, nil
}

func (r *Repository) DeleteOne(id int) error {
	query := r.db.Delete(&model.Comment{}, "id = ?", id)
	if err := query.Error; nil != err {
		return err
	}

	return nil
}

func (r *Repository) ListAll(req *pb.ListCommentRequest) ([]*model.Comment, int64, error) {
	var count int64
	var list []*model.Comment

	limit := int(req.GetLimit())
	offset := limit * int(req.GetPage()-1)

	listQuery := r.db.Model(&model.Comment{}).Select("*").Limit(limit).Offset(offset).Where("hotel_id = ?", req.GetHotelId())
	countQuery := r.db.Model(&model.Comment{}).Select("id").Where("hotel_id = ?", req.GetHotelId())

	if req.GetSortAsc() == 1 {
		listQuery.Order("created_at")
	} else {
		listQuery.Order("created_at desc")
	}

	if err := countQuery.Count(&count).Error; nil != err {
		return nil, 0, err
	}
	if err := listQuery.Find(&list).Error; nil != err {
		return nil, 0, err
	}

	return list, count, nil
}

func (r *Repository) ListRateOfHotel(hotelID int) ([]float32, error) {
	var list []float32

	listQuery := r.db.Model(&model.Comment{}).Select("rate").Where("hotel_id = ?", hotelID)

	if err := listQuery.Find(&list).Error; nil != err {
		return nil, err
	}

	return list, nil
}
