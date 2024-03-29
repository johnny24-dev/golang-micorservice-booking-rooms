// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comment/comment.proto

package comment

import (
	fmt "fmt"
	math "math"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/nekizz/final-project/backend/go-pbtype/classify"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Comment) Validate() error {
	if this.Classify != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Classify); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Classify", err)
		}
	}
	return nil
}
func (this *ListCommentRequest) Validate() error {
	return nil
}
func (this *ListCommentResponse) Validate() error {
	for _, item := range this.Items {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Items", err)
			}
		}
	}
	return nil
}
func (this *OneCommentRequest) Validate() error {
	return nil
}
