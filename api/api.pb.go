// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: api/api.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FetchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *FetchRequest) Reset() {
	*x = FetchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchRequest) ProtoMessage() {}

func (x *FetchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchRequest.ProtoReflect.Descriptor instead.
func (*FetchRequest) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{0}
}

func (x *FetchRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type FetchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *FetchResponse) Reset() {
	*x = FetchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchResponse) ProtoMessage() {}

func (x *FetchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchResponse.ProtoReflect.Descriptor instead.
func (*FetchResponse) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{1}
}

func (x *FetchResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Column         string `protobuf:"bytes,1,opt,name=column,proto3" json:"column,omitempty"`
	Order          int32  `protobuf:"varint,2,opt,name=order,proto3" json:"order,omitempty"`
	PageNumber     int64  `protobuf:"varint,3,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	ResultsPerPage int64  `protobuf:"varint,4,opt,name=results_per_page,json=resultsPerPage,proto3" json:"results_per_page,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{2}
}

func (x *ListRequest) GetColumn() string {
	if x != nil {
		return x.Column
	}
	return ""
}

func (x *ListRequest) GetOrder() int32 {
	if x != nil {
		return x.Order
	}
	return 0
}

func (x *ListRequest) GetPageNumber() int64 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *ListRequest) GetResultsPerPage() int64 {
	if x != nil {
		return x.ResultsPerPage
	}
	return 0
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*Result `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{3}
}

func (x *ListResponse) GetResults() []*Result {
	if x != nil {
		return x.Results
	}
	return nil
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product           string  `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	Price             float64 `protobuf:"fixed64,2,opt,name=price,proto3" json:"price,omitempty"`
	Timespricechanged int64   `protobuf:"varint,3,opt,name=timespricechanged,proto3" json:"timespricechanged,omitempty"`
	Requesttime       int64   `protobuf:"varint,4,opt,name=requesttime,proto3" json:"requesttime,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_api_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_api_api_proto_rawDescGZIP(), []int{4}
}

func (x *Result) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *Result) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Result) GetTimespricechanged() int64 {
	if x != nil {
		return x.Timespricechanged
	}
	return 0
}

func (x *Result) GetRequesttime() int64 {
	if x != nil {
		return x.Requesttime
	}
	return 0
}

var File_api_api_proto protoreflect.FileDescriptor

var file_api_api_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x70, 0x69, 0x22, 0x20, 0x0a, 0x0c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x25, 0x0a, 0x0d, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x86, 0x01,
	0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x10,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x50,
	0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x22, 0x35, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0x88, 0x01,
	0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x70, 0x72, 0x69, 0x63, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x11, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x70, 0x72, 0x69, 0x63, 0x65, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x74, 0x69, 0x6d, 0x65, 0x32, 0x66, 0x0a, 0x03, 0x41, 0x70, 0x69, 0x12,
	0x30, 0x0a, 0x05, 0x46, 0x65, 0x74, 0x63, 0x68, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x2d, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b,
	0x73, 0x75, 0x6b, 0x68, 0x6f, 0x72, 0x75, 0x6b, 0x6f, 0x76, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x6e,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_api_proto_rawDescOnce sync.Once
	file_api_api_proto_rawDescData = file_api_api_proto_rawDesc
)

func file_api_api_proto_rawDescGZIP() []byte {
	file_api_api_proto_rawDescOnce.Do(func() {
		file_api_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_api_proto_rawDescData)
	})
	return file_api_api_proto_rawDescData
}

var file_api_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_api_proto_goTypes = []interface{}{
	(*FetchRequest)(nil),  // 0: api.FetchRequest
	(*FetchResponse)(nil), // 1: api.FetchResponse
	(*ListRequest)(nil),   // 2: api.ListRequest
	(*ListResponse)(nil),  // 3: api.ListResponse
	(*Result)(nil),        // 4: api.Result
}
var file_api_api_proto_depIdxs = []int32{
	4, // 0: api.ListResponse.results:type_name -> api.Result
	0, // 1: api.Api.Fetch:input_type -> api.FetchRequest
	2, // 2: api.Api.List:input_type -> api.ListRequest
	1, // 3: api.Api.Fetch:output_type -> api.FetchResponse
	3, // 4: api.Api.List:output_type -> api.ListResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_api_proto_init() }
func file_api_api_proto_init() {
	if File_api_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_api_proto_goTypes,
		DependencyIndexes: file_api_api_proto_depIdxs,
		MessageInfos:      file_api_api_proto_msgTypes,
	}.Build()
	File_api_api_proto = out.File
	file_api_api_proto_rawDesc = nil
	file_api_api_proto_goTypes = nil
	file_api_api_proto_depIdxs = nil
}
