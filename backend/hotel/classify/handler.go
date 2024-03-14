package classify

import (
	pb "github.com/nekizz/final-project/backend/go-pbtype/classify"
	"github.com/nekizz/final-project/backend/hotel/model"
	"gorm.io/gorm"
)

func GetClassifyOfComment(listCmt []*model.Comment, db *gorm.DB) (map[int]*pb.Classify, error) {
	var listIDCmt []int32
	mapping := map[int]*pb.Classify{}

	for _, cmt := range listCmt {
		listIDCmt = append(listIDCmt, int32(cmt.ID))
	}

	listClassify, err := NewRepository(db).getAllClassify(listIDCmt)
	if err != nil {
		return nil, err
	}

	for _, classify := range listClassify {
		mapping[classify.ID] = prepareClassifyToResponse(classify)
	}

	return mapping, nil
}
