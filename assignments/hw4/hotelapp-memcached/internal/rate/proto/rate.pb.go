// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/rate/proto/rate.proto

/*
Package rate is a generated protocol buffer package.

It is generated from these files:
	internal/rate/proto/rate.proto

It has these top-level messages:
	Request
	Result
	RatePlans
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

type Request struct {
	HotelIds []string `protobuf:"bytes,1,rep,name=hotelIds" bson:"hotelIds,omitempty"`
	InDate   string   `protobuf:"bytes,2,opt,name=inDate" bson:"inDate,omitempty"`
	OutDate  string   `protobuf:"bytes,3,opt,name=outDate" bson:"outDate,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetHotelIds() []string {
	if m != nil {
		return m.HotelIds
	}
	return nil
}

func (m *Request) GetInDate() string {
	if m != nil {
		return m.InDate
	}
	return ""
}

func (m *Request) GetOutDate() string {
	if m != nil {
		return m.OutDate
	}
	return ""
}

type Result struct {
	RatePlans []*RatePlans `protobuf:"bytes,1,rep,name=ratePlans" bson:"ratePlans,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Result) GetRatePlans() []*RatePlans {
	if m != nil {
		return m.RatePlans
	}
	return nil
}

type RatePlans struct {
	HotelId  string    `protobuf:"bytes,1,opt,name=hotelId" bson:"hotelId,omitempty"`
	Code     string    `protobuf:"bytes,2,opt,name=code" bson:"code,omitempty"`
	InDate   string    `protobuf:"bytes,3,opt,name=inDate" bson:"inDate,omitempty"`
	OutDate  string    `protobuf:"bytes,4,opt,name=outDate" bson:"outDate,omitempty"`
	RoomType *RoomType `protobuf:"bytes,5,opt,name=roomType" bson:"roomType,omitempty"`
}

func (m *RatePlans) Reset()                    { *m = RatePlans{} }
func (m *RatePlans) String() string            { return proto.CompactTextString(m) }
func (*RatePlans) ProtoMessage()               {}
func (*RatePlans) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RatePlans) GetHotelId() string {
	if m != nil {
		return m.HotelId
	}
	return ""
}

func (m *RatePlans) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *RatePlans) GetInDate() string {
	if m != nil {
		return m.InDate
	}
	return ""
}

func (m *RatePlans) GetOutDate() string {
	if m != nil {
		return m.OutDate
	}
	return ""
}

func (m *RatePlans) GetRoomType() *RoomType {
	if m != nil {
		return m.RoomType
	}
	return nil
}

type RoomType struct {
	BookableRate       float64 `protobuf:"fixed64,1,opt,name=bookableRate" bson:"bookableRate,omitempty"`
	TotalRate          float64 `protobuf:"fixed64,2,opt,name=totalRate" bson:"totalRate,omitempty"`
	TotalRateInclusive float64 `protobuf:"fixed64,3,opt,name=totalRateInclusive" bson:"totalRateInclusive,omitempty"`
	Code               string  `protobuf:"bytes,4,opt,name=code" bson:"code,omitempty"`
	Currency           string  `protobuf:"bytes,5,opt,name=currency" bson:"currency,omitempty"`
	RoomDescription    string  `protobuf:"bytes,6,opt,name=roomDescription" bson:"roomDescription,omitempty"`
}

func (m *RoomType) Reset()                    { *m = RoomType{} }
func (m *RoomType) String() string            { return proto.CompactTextString(m) }
func (*RoomType) ProtoMessage()               {}
func (*RoomType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RoomType) GetBookableRate() float64 {
	if m != nil {
		return m.BookableRate
	}
	return 0
}

func (m *RoomType) GetTotalRate() float64 {
	if m != nil {
		return m.TotalRate
	}
	return 0
}

func (m *RoomType) GetTotalRateInclusive() float64 {
	if m != nil {
		return m.TotalRateInclusive
	}
	return 0
}

func (m *RoomType) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *RoomType) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *RoomType) GetRoomDescription() string {
	if m != nil {
		return m.RoomDescription
	}
	return ""
}

type Hotel struct {
	Id          string   `protobuf:"bytes,1,opt,name=id" bson:"id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name" bson:"name,omitempty"`
	PhoneNumber string   `protobuf:"bytes,3,opt,name=phoneNumber" bson:"phoneNumber,omitempty"`
	Description string   `protobuf:"bytes,4,opt,name=description" bson:"description,omitempty"`
	Address     *Address `protobuf:"bytes,5,opt,name=address" bson:"address,omitempty"`
	Images      []*Image `protobuf:"bytes,6,rep,name=images" bson:"images,omitempty"`
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
	StreetNumber string  `protobuf:"bytes,1,opt,name=streetNumber" bson:"streetNumber,omitempty"`
	StreetName   string  `protobuf:"bytes,2,opt,name=streetName" bson:"streetName,omitempty"`
	City         string  `protobuf:"bytes,3,opt,name=city" bson:"city,omitempty"`
	State        string  `protobuf:"bytes,4,opt,name=state" bson:"state,omitempty"`
	Country      string  `protobuf:"bytes,5,opt,name=country" bson:"country,omitempty"`
	PostalCode   string  `protobuf:"bytes,6,opt,name=postalCode" bson:"postalCode,omitempty"`
	Lat          float32 `protobuf:"fixed32,7,opt,name=lat" bson:"lat,omitempty"`
	Lon          float32 `protobuf:"fixed32,8,opt,name=lon" bson:"lon,omitempty"`
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
	Url     string `protobuf:"bytes,1,opt,name=url" bson:"url,omitempty"`
	Default bool   `protobuf:"varint,2,opt,name=default" bson:"default,omitempty"`
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
	proto.RegisterType((*Request)(nil), "rate.Request")
	proto.RegisterType((*Result)(nil), "rate.Result")
	proto.RegisterType((*RatePlans)(nil), "rate.RatePlans")
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
	// GetRates returns rate codes for hotels for a given date range
	GetRates(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error)
}

type rateClient struct {
	cc *grpc.ClientConn
}

func NewRateClient(cc *grpc.ClientConn) RateClient {
	return &rateClient{cc}
}

func (c *rateClient) GetRates(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/rate.Rate/GetRates", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Rate service

type RateServer interface {
	// GetRates returns rate codes for hotels for a given date range
	GetRates(context.Context, *Request) (*Result, error)
}

func RegisterRateServer(s *grpc.Server, srv RateServer) {
	s.RegisterService(&_Rate_serviceDesc, srv)
}

func _Rate_GetRates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
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
		return srv.(RateServer).GetRates(ctx, req.(*Request))
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
	Metadata: "internal/rate/proto/rate.proto",
}

func init() { proto.RegisterFile("internal/rate/proto/rate.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0x66, 0xfa, 0x93, 0x26, 0xa7, 0xeb, 0xae, 0x1c, 0x44, 0xc2, 0x22, 0x4b, 0x89, 0x17, 0x5b,
	0x04, 0x5b, 0xe8, 0x5e, 0x78, 0x2d, 0x2e, 0x68, 0x6f, 0x44, 0x06, 0xc1, 0xeb, 0x69, 0x32, 0xba,
	0xc1, 0xe9, 0x4c, 0x9d, 0x99, 0x08, 0x7d, 0x12, 0x9f, 0xc6, 0xc7, 0xd0, 0xe7, 0x91, 0xf9, 0x4b,
	0x53, 0xd1, 0xbb, 0xef, 0x67, 0xda, 0xf3, 0x9d, 0x2f, 0x93, 0xc0, 0x4d, 0x2b, 0x2d, 0xd7, 0x92,
	0x89, 0xb5, 0x66, 0x96, 0xaf, 0x0f, 0x5a, 0x59, 0xe5, 0xe1, 0xca, 0x43, 0x9c, 0x38, 0x5c, 0x7d,
	0x82, 0x19, 0xe5, 0xdf, 0x3a, 0x6e, 0x2c, 0x5e, 0x43, 0xfe, 0xa0, 0x2c, 0x17, 0xdb, 0xc6, 0x94,
	0x64, 0x31, 0x5e, 0x16, 0xb4, 0xe7, 0xf8, 0x14, 0xb2, 0x56, 0xde, 0x33, 0xcb, 0xcb, 0xd1, 0x82,
	0x2c, 0x0b, 0x1a, 0x19, 0x96, 0x30, 0x53, 0x9d, 0xf5, 0xc6, 0xd8, 0x1b, 0x89, 0x56, 0xaf, 0x20,
	0xa3, 0xdc, 0x74, 0xc2, 0xe2, 0x4b, 0x28, 0xdc, 0xa8, 0x0f, 0x82, 0xc9, 0xf0, 0xc7, 0xf3, 0xcd,
	0xd5, 0xca, 0x07, 0xa1, 0x49, 0xa6, 0xa7, 0x13, 0xd5, 0x0f, 0x02, 0x45, 0x6f, 0xb8, 0x01, 0x31,
	0x44, 0x49, 0xc2, 0x80, 0x48, 0x11, 0x61, 0x52, 0xab, 0x26, 0x05, 0xf2, 0x78, 0x10, 0x73, 0xfc,
	0xbf, 0x98, 0x93, 0xb3, 0x98, 0xf8, 0x02, 0x72, 0xad, 0xd4, 0xfe, 0xe3, 0xf1, 0xc0, 0xcb, 0xe9,
	0x82, 0x2c, 0xe7, 0x9b, 0xcb, 0x98, 0x2d, 0xaa, 0xb4, 0xf7, 0xab, 0xdf, 0x04, 0xf2, 0x24, 0x63,
	0x05, 0x17, 0x3b, 0xa5, 0xbe, 0xb2, 0x9d, 0xe0, 0x2e, 0xad, 0x4f, 0x47, 0xe8, 0x99, 0x86, 0xcf,
	0xa0, 0xb0, 0xca, 0x32, 0x41, 0x53, 0x71, 0x84, 0x9e, 0x04, 0x5c, 0x01, 0xf6, 0x64, 0x2b, 0x6b,
	0xd1, 0x99, 0xf6, 0x7b, 0x08, 0x4e, 0xe8, 0x3f, 0x9c, 0x7e, 0xe1, 0xc9, 0x60, 0xe1, 0x6b, 0xc8,
	0xeb, 0x4e, 0x6b, 0x2e, 0xeb, 0xa3, 0x8f, 0x5f, 0xd0, 0x9e, 0xe3, 0x12, 0xae, 0x5c, 0xf4, 0x7b,
	0x6e, 0x6a, 0xdd, 0x1e, 0x6c, 0xab, 0x64, 0x99, 0xf9, 0x23, 0x7f, 0xcb, 0xd5, 0x4f, 0x02, 0xd3,
	0x77, 0xae, 0x56, 0xbc, 0x84, 0x51, 0x9b, 0x9a, 0x1e, 0xb5, 0xbe, 0x64, 0xc9, 0xf6, 0x7d, 0xc9,
	0x0e, 0xe3, 0x02, 0xe6, 0x87, 0x07, 0x25, 0xf9, 0xfb, 0x6e, 0xbf, 0xe3, 0x3a, 0x36, 0x3d, 0x94,
	0xdc, 0x89, 0x66, 0x30, 0x35, 0x04, 0x1e, 0x4a, 0x78, 0x0b, 0x33, 0xd6, 0x34, 0x9a, 0x1b, 0x13,
	0x5b, 0x7f, 0x14, 0x5a, 0x7f, 0x1d, 0x44, 0x9a, 0x5c, 0x7c, 0x0e, 0x59, 0xbb, 0x67, 0x5f, 0xb8,
	0x29, 0x33, 0x7f, 0x73, 0xe6, 0xe1, 0xdc, 0xd6, 0x69, 0x34, 0x5a, 0xd5, 0x2f, 0x02, 0xb3, 0xf8,
	0x4b, 0xf7, 0x5c, 0x8c, 0xd5, 0x9c, 0xdb, 0x18, 0x2f, 0xec, 0x72, 0xa6, 0xe1, 0x0d, 0x40, 0xe4,
	0xa7, 0xdd, 0x06, 0x8a, 0x6f, 0xba, 0xb5, 0xc7, 0xb8, 0x9a, 0xc7, 0xf8, 0x04, 0xa6, 0xc6, 0x9e,
	0x2e, 0x50, 0x20, 0xee, 0x62, 0xd5, 0xaa, 0x93, 0x56, 0xa7, 0xfa, 0x13, 0x75, 0x33, 0x0e, 0xca,
	0x58, 0x26, 0xde, 0xb8, 0x67, 0x16, 0x8a, 0x1f, 0x28, 0xf8, 0x18, 0xc6, 0x82, 0xd9, 0x72, 0xb6,
	0x20, 0xcb, 0x11, 0x75, 0xd0, 0x2b, 0x4a, 0x96, 0x79, 0x54, 0x94, 0xac, 0xee, 0x60, 0xea, 0x17,
	0x75, 0x56, 0xa7, 0x45, 0xdc, 0xc5, 0x41, 0x37, 0xb8, 0xe1, 0x9f, 0x59, 0x27, 0xac, 0xcf, 0x9f,
	0xd3, 0x44, 0x37, 0x6b, 0x98, 0xf8, 0xeb, 0x75, 0x0b, 0xf9, 0x5b, 0x6e, 0x1d, 0x34, 0x18, 0xdb,
	0x8d, 0x6f, 0xfa, 0xf5, 0x45, 0xa2, 0xee, 0xfd, 0xdc, 0x65, 0xfe, 0x7b, 0x70, 0xf7, 0x27, 0x00,
	0x00, 0xff, 0xff, 0xe3, 0x23, 0x73, 0x87, 0x31, 0x04, 0x00, 0x00,
}
