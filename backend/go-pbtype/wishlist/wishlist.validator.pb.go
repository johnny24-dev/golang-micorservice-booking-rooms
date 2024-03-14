// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: wishlist/wishlist.proto

package wishlist

import (
	fmt "fmt"
	math "math"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	_ "github.com/gogo/protobuf/gogoproto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/nekizz/final-project/backend/go-pbtype/hotel"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Wishlist) Validate() error {
	return nil
}
func (this *OneWishlistRequest) Validate() error {
	return nil
}
func (this *ListWishlistRequest) Validate() error {
	return nil
}
func (this *ListWishlistResponse) Validate() error {
	for _, item := range this.Items {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Items", err)
			}
		}
	}
	return nil
}