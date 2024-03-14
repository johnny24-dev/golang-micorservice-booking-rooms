package hotel

import (
	"fmt"
	pb "github.com/nekizz/final-project/backend/go-pbtype/hotel"
	"github.com/nekizz/final-project/backend/hotel/model"
	"github.com/nekizz/final-project/backend/hotel/util"
	"gorm.io/gorm"
	"strings"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreatOne(e *model.Hotel) (*model.Hotel, error) {
	query := r.db.Model(&model.Hotel{}).Create(e)
	if err := query.Error; nil != err {
		return nil, err
	}

	return e, nil
}

func (r *Repository) UpdateOne(id int, c *model.Hotel) (*model.Hotel, error) {
	query := r.db.Model(&model.Hotel{}).Omit("rate", "is_available, is_deleted, customer_id, address_id").Where("id = ?", id).UpdateColumns(getModel(uint(id), c))

	if err := query.Error; nil != err {
		return nil, err
	}

	return c, nil
}

func (r *Repository) FindOne(id int) (*model.Hotel, error) {
	var hotel model.Hotel

	query := r.db.Model(&model.Hotel{}).Preload("Image").Preload("Address").Preload("Comment").Where("id = ?", id).Find(&hotel)
	if err := query.Error; nil != err {
		return nil, err
	}

	return &hotel, nil
}

func (r *Repository) ListAll(req *pb.ListHotelRequest) ([]*model.Hotel, int64, error) {
	var count int64
	var list []*model.Hotel

	limit := int(req.GetLimit())
	offset := limit * int(req.GetPage()-1)

	listQuery := r.db.Model(&model.Hotel{}).Preload("Image").Preload("Address").Select("*").Limit(limit).Offset(offset).Where("customer_id = ? AND is_deleted = 1", req.CustomerId)
	countQuery := r.db.Model(&model.Hotel{}).Select("id").Where("customer_id = ? AND is_deleted = 1", req.CustomerId)

	if err := countQuery.Count(&count).Error; nil != err {
		return nil, 0, err
	}

	if err := listQuery.Find(&list).Limit(limit).Offset(offset).Error; nil != err {
		return nil, 0, err
	}

	return list, count, nil
}

func (r *Repository) ListAllTwo(req *pb.ListHotelCustomerRequest, mappingSearch map[string]string) ([]*model.Hotel, int64, error) {
	var count int64
	var list []*model.Hotel

	limit := int(req.GetLimit())
	offset := limit * int(req.GetPage()-1)

	listQuery := r.db.Model(&model.Hotel{}).Preload("Image").Preload("Address").Select("*").Where("is_deleted = 1").Limit(limit).Offset(offset)
	countQuery := r.db.Model(&model.Hotel{}).Select("id")

	for field, value := range mappingSearch {
		sql := util.BuildQuerySearch(field, value)
		listQuery = listQuery.Where(sql)
		countQuery = countQuery.Where(sql)
	}

	if err := countQuery.Count(&count).Error; nil != err {
		return nil, 0, err
	}

	if err := listQuery.Find(&list).Error; nil != err {
		return nil, 0, err
	}

	sql1 := r.db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		var list2 []*model.Hotel

		testQuery := tx.Model(&model.Hotel{}).Select("*").Limit(limit).Offset(offset)

		for field, value := range mappingSearch {
			sql := util.BuildQuerySearch(field, value)
			testQuery = testQuery.Where(sql)
		}

		return testQuery.Find(&list2)
	})
	fmt.Println("test query: ", sql1)

	return list, count, nil
}

func (r *Repository) DeleteOne(id int) error {
	//	query := r.db.Delete(&model.Hotel{}, "id = ?", id)
	query := r.db.Table("hotels").Where("id = ?", id).Updates(map[string]interface{}{"is_deleted": 0})
	if err := query.Error; nil != err {
		return err
	}

	return nil
}

func (r *Repository) listAllAddressByObject(id int) (map[uint]model.Address, error) {
	var listIdHotelByID []uint32
	var listAddressByObject []model.Address
	mapping := map[uint]model.Address{}

	query := r.db.Model(&model.Hotel{}).Select("address_id").Where("customer_id = ? AND is_deleted = 1", id).Find(&listIdHotelByID)
	if err := query.Error; nil != err {
		return nil, err
	}

	subQuery := r.db.Model(&model.Address{}).Select("*").Where("id IN ? ", listIdHotelByID).Find(&listAddressByObject)
	if err := subQuery.Error; nil != err {
		return nil, err
	}

	if len(listAddressByObject) > 0 {
		for _, val := range listAddressByObject {
			mapping[val.ID] = val
		}
	}

	return mapping, nil
}

func (r *Repository) ListHotelByIDs(req *pb.ListHotelByIDRequest) ([]*model.Hotel, int64, error) {
	var count int64
	var list []*model.Hotel

	limit := int(req.GetLimit())
	offset := limit * int(req.GetPage()-1)

	listQuery := r.db.Model(&model.Hotel{}).Preload("Image").Preload("Address").Select("*").Limit(limit).Offset(offset).Where("id IN ? AND is_deleted = 1", req.GetListHotelId())
	countQuery := r.db.Model(&model.Hotel{}).Select("id").Where("id IN ? AND is_deleted = 1", req.GetListHotelId())

	if err := countQuery.Count(&count).Error; nil != err {
		return nil, 0, err
	}

	if err := listQuery.Find(&list).Limit(limit).Offset(offset).Error; nil != err {
		return nil, 0, err
	}

	return list, count, nil
}

func (r *Repository) SearchHotelWithFilter(req *pb.ListHotelRequest) ([]*model.Hotel, int64, error) {
	var count int64
	var list []*model.Hotel

	sql := ""
	limit := int(req.GetLimit())
	offset := limit * int(req.GetPage()-1)

	if req.GetSearchField() != "" && req.GetSearchValue() != "" {
		searchFields := strings.Split(req.GetSearchField(), ",")
		searchValue := fmt.Sprintf("'%%%s%%'", req.GetSearchValue())

		for idx, field := range searchFields {
			if idx == 0 {
				sql += fmt.Sprintf("%s LIKE %s", field, searchValue)
				continue
			}
			sql += fmt.Sprintf(" OR %s LIKE %s", field, searchValue)
		}
	}

	listQuery := r.db.Model(&model.Hotel{}).Preload("Image").Select("*").Limit(limit).Offset(offset)
	countQuery := r.db.Model(&model.Hotel{}).Select("id")

	//if sql != "" {
	//	countQuery = countQuery.Where(sql)
	//	listQuery = listQuery.Where(sql)
	//}

	if err := countQuery.Count(&count).Error; nil != err {
		return nil, 0, err
	}

	if err := listQuery.Find(&list).Limit(limit).Offset(offset).Error; nil != err {
		return nil, 0, err
	}

	return list, count, nil
}
