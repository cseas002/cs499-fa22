// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rate.proto

/*
Package rate is a generated protocol buffer package.

It is generated from these files:
	rate.proto

It has these top-level messages:
	RateRequest
	RateResult
	RatePlan
	RoomType
	Hotel
	Address
	Image
*/
package rate

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RateRequest struct {
	HotelIds []string `protobuf:"bytes,1,rep,name=hotelIds" json:"hotelIds,omitempty"`
	InDate   string   `protobuf:"bytes,2,opt,name=InDate,json=inDate" json:"InDate,omitempty"`
	OutDate  string   `protobuf:"bytes,3,opt,name=OutDate,json=outDate" json:"OutDate,omitempty"`
}

func (m *RateRequest) Reset()                    { *m = RateRequest{} }
func (m *RateRequest) String() string            { return proto.CompactTextString(m) }
func (*RateRequest) ProtoMessage()               {}
func (*RateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RateRequest) GetHotelIds() []string {
	if m != nil {
		return m.HotelIds
	}
	return nil
}

func (m *RateRequest) GetInDate() string {
	if m != nil {
		return m.InDate
	}
	return ""
}

func (m *RateRequest) GetOutDate() string {
	if m != nil {
		return m.OutDate
	}
	return ""
}

type RateResult struct {
	RatePlans []*RatePlan `protobuf:"bytes,1,rep,name=RatePlans,json=ratePlans" json:"RatePlans,omitempty"`
}

func (m *RateResult) Reset()                    { *m = RateResult{} }
func (m *RateResult) String() string            { return proto.CompactTextString(m) }
func (*RateResult) ProtoMessage()               {}
func (*RateResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RateResult) GetRatePlans() []*RatePlan {
	if m != nil {
		return m.RatePlans
	}
	return nil
}

type RatePlan struct {
	InDate   string    `protobuf:"bytes,1,opt,name=InDate,json=inDate" json:"InDate,omitempty"`
	OutDate  string    `protobuf:"bytes,2,opt,name=OutDate,json=outDate" json:"OutDate,omitempty"`
	HotelId  string    `protobuf:"bytes,3,opt,name=HotelId,json=hotelId" json:"HotelId,omitempty"`
	RoomType *RoomType `protobuf:"bytes,4,opt,name=RoomType,json=roomType" json:"RoomType,omitempty"`
}

func (m *RatePlan) Reset()                    { *m = RatePlan{} }
func (m *RatePlan) String() string            { return proto.CompactTextString(m) }
func (*RatePlan) ProtoMessage()               {}
func (*RatePlan) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RatePlan) GetInDate() string {
	if m != nil {
		return m.InDate
	}
	return ""
}

func (m *RatePlan) GetOutDate() string {
	if m != nil {
		return m.OutDate
	}
	return ""
}

func (m *RatePlan) GetHotelId() string {
	if m != nil {
		return m.HotelId
	}
	return ""
}

func (m *RatePlan) GetRoomType() *RoomType {
	if m != nil {
		return m.RoomType
	}
	return nil
}

type RoomType struct {
	TotalRate float64 `protobuf:"fixed64,1,opt,name=TotalRate,json=totalRate" json:"TotalRate,omitempty"`
}

func (m *RoomType) Reset()                    { *m = RoomType{} }
func (m *RoomType) String() string            { return proto.CompactTextString(m) }
func (*RoomType) ProtoMessage()               {}
func (*RoomType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RoomType) GetTotalRate() float64 {
	if m != nil {
		return m.TotalRate
	}
	return 0
}

type Hotel struct {
	Id          string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	PhoneNumber string   `protobuf:"bytes,3,opt,name=phoneNumber" json:"phoneNumber,omitempty"`
	Description string   `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	Address     *Address `protobuf:"bytes,5,opt,name=address" json:"address,omitempty"`
	Images      []*Image `protobuf:"bytes,6,rep,name=images" json:"images,omitempty"`
}

func (m *Hotel) Reset()                    { *m = Hotel{} }
func (m *Hotel) String() string            { return proto.CompactTextString(m) }
func (*Hotel) ProtoMessage()               {}
func (*Hotel) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Hotel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Hotel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Hotel) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *Hotel) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Hotel) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Hotel) GetImages() []*Image {
	if m != nil {
		return m.Images
	}
	return nil
}

type Address struct {
	StreetNumber string  `protobuf:"bytes,1,opt,name=streetNumber" json:"streetNumber,omitempty"`
	StreetName   string  `protobuf:"bytes,2,opt,name=streetName" json:"streetName,omitempty"`
	City         string  `protobuf:"bytes,3,opt,name=city" json:"city,omitempty"`
	State        string  `protobuf:"bytes,4,opt,name=state" json:"state,omitempty"`
	Country      string  `protobuf:"bytes,5,opt,name=country" json:"country,omitempty"`
	PostalCode   string  `protobuf:"bytes,6,opt,name=postalCode" json:"postalCode,omitempty"`
	Lat          float32 `protobuf:"fixed32,7,opt,name=lat" json:"lat,omitempty"`
	Lon          float32 `protobuf:"fixed32,8,opt,name=lon" json:"lon,omitempty"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (m *Address) String() string            { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Address) GetStreetNumber() string {
	if m != nil {
		return m.StreetNumber
	}
	return ""
}

func (m *Address) GetStreetName() string {
	if m != nil {
		return m.StreetName
	}
	return ""
}

func (m *Address) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Address) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Address) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Address) GetPostalCode() string {
	if m != nil {
		return m.PostalCode
	}
	return ""
}

func (m *Address) GetLat() float32 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *Address) GetLon() float32 {
	if m != nil {
		return m.Lon
	}
	return 0
}

type Image struct {
	Url     string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Default bool   `protobuf:"varint,2,opt,name=default" json:"default,omitempty"`
}

func (m *Image) Reset()                    { *m = Image{} }
func (m *Image) String() string            { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()               {}
func (*Image) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Image) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Image) GetDefault() bool {
	if m != nil {
		return m.Default
	}
	return false
}

func init() {
	proto.RegisterType((*RateRequest)(nil), "rate.RateRequest")
	proto.RegisterType((*RateResult)(nil), "rate.RateResult")
	proto.RegisterType((*RatePlan)(nil), "rate.RatePlan")
	proto.RegisterType((*RoomType)(nil), "rate.RoomType")
	proto.RegisterType((*Hotel)(nil), "rate.Hotel")
	proto.RegisterType((*Address)(nil), "rate.Address")
	proto.RegisterType((*Image)(nil), "rate.Image")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Rate service

type RateClient interface {
	GetRates(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResult, error)
}

type rateClient struct {
	cc *grpc.ClientConn
}

func NewRateClient(cc *grpc.ClientConn) RateClient {
	return &rateClient{cc}
}

func (c *rateClient) GetRates(ctx context.Context, in *RateRequest, opts ...grpc.CallOption) (*RateResult, error) {
	out := new(RateResult)
	err := grpc.Invoke(ctx, "/rate.Rate/GetRates", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Rate service

type RateServer interface {
	GetRates(context.Context, *RateRequest) (*RateResult, error)
}

func RegisterRateServer(s *grpc.Server, srv RateServer) {
	s.RegisterService(&_Rate_serviceDesc, srv)
}

func _Rate_GetRates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateServer).GetRates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rate.Rate/GetRates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateServer).GetRates(ctx, req.(*RateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Rate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rate.Rate",
	HandlerType: (*RateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRates",
			Handler:    _Rate_GetRates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rate.proto",
}

func init() { proto.RegisterFile("rate.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 458 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xcf, 0x8a, 0xd4, 0x40,
	0x10, 0xc6, 0xe9, 0xcc, 0x4c, 0xfe, 0xd4, 0xe8, 0xb2, 0x16, 0x22, 0xcd, 0x22, 0x12, 0xe2, 0xc1,
	0x41, 0x64, 0x85, 0xdd, 0x83, 0xe0, 0x4d, 0x14, 0x74, 0x2e, 0x2a, 0xcd, 0xde, 0x3c, 0xf5, 0x4e,
	0x5a, 0x37, 0x90, 0x49, 0xc7, 0xee, 0xca, 0x61, 0xae, 0xbe, 0x97, 0x8f, 0xe1, 0xfb, 0x48, 0xff,
	0x33, 0xd9, 0x83, 0xb7, 0xfa, 0x7e, 0x55, 0x49, 0x7d, 0x5f, 0xd2, 0x0d, 0x60, 0x24, 0xa9, 0xcb,
	0xd1, 0x68, 0xd2, 0xb8, 0x76, 0x75, 0xf3, 0x0d, 0xb6, 0x42, 0x92, 0x12, 0xea, 0xe7, 0xa4, 0x2c,
	0xe1, 0x05, 0x94, 0x77, 0x9a, 0x54, 0xbf, 0x6f, 0x2d, 0x67, 0xf5, 0x6a, 0x57, 0x89, 0x7f, 0x1a,
	0x9f, 0x40, 0xbe, 0x1f, 0x3e, 0x48, 0x52, 0x3c, 0xab, 0xd9, 0xae, 0x12, 0x79, 0xe7, 0x15, 0x72,
	0x28, 0xbe, 0x4c, 0xe4, 0x1b, 0x2b, 0xdf, 0x28, 0x74, 0x90, 0xcd, 0x5b, 0x80, 0xf0, 0x72, 0x3b,
	0xf5, 0x84, 0xaf, 0xa0, 0x72, 0xea, 0x6b, 0x2f, 0x87, 0xf0, 0xf2, 0xed, 0xd5, 0xd9, 0xa5, 0x37,
	0x94, 0xb0, 0xa8, 0x4c, 0x1a, 0x68, 0x7e, 0x31, 0x28, 0x13, 0x5f, 0xac, 0x66, 0xff, 0x5b, 0x9d,
	0xdd, 0x5b, 0xed, 0x3a, 0x9f, 0x82, 0xf1, 0x64, 0x2a, 0xe6, 0xc0, 0x97, 0x50, 0x0a, 0xad, 0x8f,
	0x37, 0xa7, 0x51, 0xf1, 0x75, 0xcd, 0x16, 0x2e, 0x22, 0x15, 0xa5, 0x89, 0x55, 0xb3, 0x9b, 0x67,
	0xf1, 0x29, 0x54, 0x37, 0x9a, 0x64, 0x2f, 0x92, 0x0d, 0x26, 0x2a, 0x4a, 0xa0, 0xf9, 0xcd, 0x60,
	0xe3, 0x17, 0xe2, 0x19, 0x64, 0x5d, 0x1b, 0x7d, 0x66, 0x5d, 0x8b, 0x08, 0xeb, 0x41, 0x1e, 0x93,
	0x41, 0x5f, 0x63, 0x0d, 0xdb, 0xf1, 0x4e, 0x0f, 0xea, 0xf3, 0x74, 0xbc, 0x55, 0x26, 0x3a, 0x5c,
	0x22, 0x37, 0xd1, 0x2a, 0x7b, 0x30, 0xdd, 0x48, 0x9d, 0x1e, 0xbc, 0xd1, 0x4a, 0x2c, 0x11, 0xbe,
	0x80, 0x42, 0xb6, 0xad, 0x51, 0xd6, 0xf2, 0x8d, 0x8f, 0xf1, 0x30, 0xc4, 0x78, 0x17, 0xa0, 0x48,
	0x5d, 0x7c, 0x0e, 0x79, 0x77, 0x94, 0x3f, 0x94, 0xe5, 0xb9, 0xff, 0xe8, 0xdb, 0x30, 0xb7, 0x77,
	0x4c, 0xc4, 0x56, 0xf3, 0x87, 0x41, 0x11, 0x9f, 0xc4, 0x06, 0x1e, 0x58, 0x32, 0x4a, 0x51, 0xb4,
	0x17, 0xb2, 0xdc, 0x63, 0xf8, 0x0c, 0x20, 0xea, 0x39, 0xdb, 0x82, 0xb8, 0xd4, 0x87, 0x8e, 0x4e,
	0x31, 0x9a, 0xaf, 0xf1, 0x31, 0x6c, 0x2c, 0xb9, 0xaf, 0x17, 0xd2, 0x04, 0xe1, 0xfe, 0xd4, 0x41,
	0x4f, 0x03, 0x99, 0x93, 0xcf, 0x51, 0x89, 0x24, 0xdd, 0x8e, 0x51, 0x5b, 0x92, 0xfd, 0x7b, 0xdd,
	0x2a, 0x9e, 0x87, 0x1d, 0x33, 0xc1, 0x73, 0x58, 0xf5, 0x92, 0x78, 0x51, 0xb3, 0x5d, 0x26, 0x5c,
	0xe9, 0x89, 0x1e, 0x78, 0x19, 0x89, 0x1e, 0x9a, 0x6b, 0xd8, 0xf8, 0xa0, 0xae, 0x35, 0x99, 0x3e,
	0x66, 0x71, 0xa5, 0x5b, 0xdc, 0xaa, 0xef, 0x72, 0xea, 0xc9, 0xfb, 0x2f, 0x45, 0x92, 0x57, 0x6f,
	0x60, 0xed, 0x7e, 0x2a, 0xbe, 0x86, 0xf2, 0xa3, 0x22, 0x57, 0x5a, 0x7c, 0x34, 0x1f, 0xd5, 0x78,
	0x59, 0x2e, 0xce, 0x97, 0xc8, 0x1d, 0xf1, 0xdb, 0xdc, 0x5f, 0xad, 0xeb, 0xbf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x1a, 0x92, 0x62, 0x87, 0x68, 0x03, 0x00, 0x00,
}
