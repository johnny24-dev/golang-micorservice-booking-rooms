package comment

import (
	"context"
	"fmt"
	"github.com/gogo/protobuf/types"
	pb "github.com/nekizz/final-project/backend/go-pbtype/comment"
	pbc "github.com/nekizz/final-project/backend/go-pbtype/customer"
	"github.com/nekizz/final-project/backend/hotel/classify"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"math"
	"strconv"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s Service) CreateComment(ctx context.Context, r *pb.Comment) (*pb.Comment, error) {
	if err := validateCreate(r); nil != err {
		return nil, err
	}

	comment, err := NewRepository(s.db).CreatOne(prepareCommentToRequest(r))
	if nil != err {
		return nil, err
	}

	return prepareCommentToResponse(comment), nil
}

func (s Service) DeleteComment(ctx context.Context, r *pb.OneCommentRequest) (*types.Empty, error) {
	if err := validateDeleteC(r); nil != err {
		return nil, err
	}

	err := NewRepository(s.db).DeleteOne(int(r.GetId()))
	if nil != err {
		return nil, err
	}

	return &types.Empty{}, nil
}

func (s Service) ListCommentOfHotel(ctx context.Context, r *pb.ListCommentRequest) (*pb.ListCommentResponse, error) {
	var list []*pb.Comment
	if err := validateList(r); nil != err {
		return nil, err
	}

	comments, count, err := NewRepository(s.db).ListAll(r)
	if nil != err {
		return nil, err
	}

	mapping, err := classify.GetClassifyOfComment(comments, s.db)
	if nil != err {
		return nil, err
	}

	for _, comment := range comments {
		commentPb := prepareCommentToResponse(comment)

		userAddr := viper.GetString("user_backend")
		userConn, err := grpc.Dial(userAddr, grpc.WithInsecure())
		if err != nil {
			fmt.Println("Error comment: ", comment.ID)
			continue
		}
		addressService := pbc.NewCustomerServiceClient(userConn)
		cusReq := &pbc.OneCustomerRequest{
			Id: strconv.Itoa(comment.CustomerID),
		}
		customer, err := addressService.Get(context.Background(), cusReq)
		if err != nil {
			fmt.Println("Error comment: ", comment.ID)
			continue
		}

		commentPb.CustomerName = customer.User.Name

		classify, ok := mapping[int(comment.ID)]
		if !ok {
			classify = nil
		}
		commentPb.Classify = classify

		list = append(list, commentPb)
	}

	return &pb.ListCommentResponse{
		Items:      list,
		TotalCount: uint32(count),
		Page:       r.GetPage(),
		Limit:      r.GetLimit(),
		MaxPage:    uint32(math.Ceil(float64(uint32(count)) / float64(r.GetLimit()))),
	}, nil

}
