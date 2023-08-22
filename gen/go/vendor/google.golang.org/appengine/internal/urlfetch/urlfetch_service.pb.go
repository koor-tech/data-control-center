// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: vendor/google.golang.org/appengine/internal/urlfetch/urlfetch_service.proto

package urlfetch

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

type URLFetchServiceError_ErrorCode int32

const (
	URLFetchServiceError_OK                       URLFetchServiceError_ErrorCode = 0
	URLFetchServiceError_INVALID_URL              URLFetchServiceError_ErrorCode = 1
	URLFetchServiceError_FETCH_ERROR              URLFetchServiceError_ErrorCode = 2
	URLFetchServiceError_UNSPECIFIED_ERROR        URLFetchServiceError_ErrorCode = 3
	URLFetchServiceError_RESPONSE_TOO_LARGE       URLFetchServiceError_ErrorCode = 4
	URLFetchServiceError_DEADLINE_EXCEEDED        URLFetchServiceError_ErrorCode = 5
	URLFetchServiceError_SSL_CERTIFICATE_ERROR    URLFetchServiceError_ErrorCode = 6
	URLFetchServiceError_DNS_ERROR                URLFetchServiceError_ErrorCode = 7
	URLFetchServiceError_CLOSED                   URLFetchServiceError_ErrorCode = 8
	URLFetchServiceError_INTERNAL_TRANSIENT_ERROR URLFetchServiceError_ErrorCode = 9
	URLFetchServiceError_TOO_MANY_REDIRECTS       URLFetchServiceError_ErrorCode = 10
	URLFetchServiceError_MALFORMED_REPLY          URLFetchServiceError_ErrorCode = 11
	URLFetchServiceError_CONNECTION_ERROR         URLFetchServiceError_ErrorCode = 12
)

// Enum value maps for URLFetchServiceError_ErrorCode.
var (
	URLFetchServiceError_ErrorCode_name = map[int32]string{
		0:  "OK",
		1:  "INVALID_URL",
		2:  "FETCH_ERROR",
		3:  "UNSPECIFIED_ERROR",
		4:  "RESPONSE_TOO_LARGE",
		5:  "DEADLINE_EXCEEDED",
		6:  "SSL_CERTIFICATE_ERROR",
		7:  "DNS_ERROR",
		8:  "CLOSED",
		9:  "INTERNAL_TRANSIENT_ERROR",
		10: "TOO_MANY_REDIRECTS",
		11: "MALFORMED_REPLY",
		12: "CONNECTION_ERROR",
	}
	URLFetchServiceError_ErrorCode_value = map[string]int32{
		"OK":                       0,
		"INVALID_URL":              1,
		"FETCH_ERROR":              2,
		"UNSPECIFIED_ERROR":        3,
		"RESPONSE_TOO_LARGE":       4,
		"DEADLINE_EXCEEDED":        5,
		"SSL_CERTIFICATE_ERROR":    6,
		"DNS_ERROR":                7,
		"CLOSED":                   8,
		"INTERNAL_TRANSIENT_ERROR": 9,
		"TOO_MANY_REDIRECTS":       10,
		"MALFORMED_REPLY":          11,
		"CONNECTION_ERROR":         12,
	}
)

func (x URLFetchServiceError_ErrorCode) Enum() *URLFetchServiceError_ErrorCode {
	p := new(URLFetchServiceError_ErrorCode)
	*p = x
	return p
}

func (x URLFetchServiceError_ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (URLFetchServiceError_ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_enumTypes[0].Descriptor()
}

func (URLFetchServiceError_ErrorCode) Type() protoreflect.EnumType {
	return &file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_enumTypes[0]
}

func (x URLFetchServiceError_ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *URLFetchServiceError_ErrorCode) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = URLFetchServiceError_ErrorCode(num)
	return nil
}

// Deprecated: Use URLFetchServiceError_ErrorCode.Descriptor instead.
func (URLFetchServiceError_ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_rawDescGZIP(), []int{0, 0}
}

type URLFetchRequest_RequestMethod int32

const (
	URLFetchRequest_GET    URLFetchRequest_RequestMethod = 1
	URLFetchRequest_POST   URLFetchRequest_RequestMethod = 2
	URLFetchRequest_HEAD   URLFetchRequest_RequestMethod = 3
	URLFetchRequest_PUT    URLFetchRequest_RequestMethod = 4
	URLFetchRequest_DELETE URLFetchRequest_RequestMethod = 5
	URLFetchRequest_PATCH  URLFetchRequest_RequestMethod = 6
)

// Enum value maps for URLFetchRequest_RequestMethod.
var (
	URLFetchRequest_RequestMethod_name = map[int32]string{
		1: "GET",
		2: "POST",
		3: "HEAD",
		4: "PUT",
		5: "DELETE",
		6: "PATCH",
	}
	URLFetchRequest_RequestMethod_value = map[string]int32{
		"GET":    1,
		"POST":   2,
		"HEAD":   3,
		"PUT":    4,
		"DELETE": 5,
		"PATCH":  6,
	}
)

func (x URLFetchRequest_RequestMethod) Enum() *URLFetchRequest_RequestMethod {
	p := new(URLFetchRequest_RequestMethod)
	*p = x
	return p
}

func (x URLFetchRequest_RequestMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (URLFetchRequest_RequestMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_enumTypes[1].Descriptor()
}

func (URLFetchRequest_RequestMethod) Type() protoreflect.EnumType {
	return &file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_enumTypes[1]
}

func (x URLFetchRequest_RequestMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *URLFetchRequest_RequestMethod) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = URLFetchRequest_RequestMethod(num)
	return nil
}

// Deprecated: Use URLFetchRequest_RequestMethod.Descriptor instead.
func (URLFetchRequest_RequestMethod) EnumDescriptor() ([]byte, []int) {
	return file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_rawDescGZIP(), []int{1, 0}
}

type URLFetchServiceError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *URLFetchServiceError) Reset() {
	*x = URLFetchServiceError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *URLFetchServiceError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URLFetchServiceError) ProtoMessage() {}

func (x *URLFetchServiceError) ProtoReflect() protoreflect.Message {
	mi := &file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URLFetchServiceError.ProtoReflect.Descriptor instead.
func (*URLFetchServiceError) Descriptor() ([]byte, []int) {
	return file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_rawDescGZIP(), []int{0}
}

type URLFetchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method                        *URLFetchRequest_RequestMethod `protobuf:"varint,1,req,name=Method,enum=appengine.URLFetchRequest_RequestMethod" json:"Method,omitempty"`
	Url                           *string                        `protobuf:"bytes,2,req,name=Url" json:"Url,omitempty"`
	Header                        []*URLFetchRequest_Header      `protobuf:"group,3,rep,name=Header,json=header" json:"header,omitempty"`
	Payload                       []byte                         `protobuf:"bytes,6,opt,name=Payload" json:"Payload,omitempty"`
	FollowRedirects               *bool                          `protobuf:"varint,7,opt,name=FollowRedirects,def=1" json:"FollowRedirects,omitempty"`
	Deadline                      *float64                       `protobuf:"fixed64,8,opt,name=Deadline" json:"Deadline,omitempty"`
	MustValidateServerCertificate *bool                          `protobuf:"varint,9,opt,name=MustValidateServerCertificate,def=1" json:"MustValidateServerCertificate,omitempty"`
}

// Default values for URLFetchRequest fields.
const (
	Default_URLFetchRequest_FollowRedirects               = bool(true)
	Default_URLFetchRequest_MustValidateServerCertificate = bool(true)
)

func (x *URLFetchRequest) Reset() {
	*x = URLFetchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *URLFetchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URLFetchRequest) ProtoMessage() {}

func (x *URLFetchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URLFetchRequest.ProtoReflect.Descriptor instead.
func (*URLFetchRequest) Descriptor() ([]byte, []int) {
	return file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_rawDescGZIP(), []int{1}
}

func (x *URLFetchRequest) GetMethod() URLFetchRequest_RequestMethod {
	if x != nil && x.Method != nil {
		return *x.Method
	}
	return URLFetchRequest_GET
}

func (x *URLFetchRequest) GetUrl() string {
	if x != nil && x.Url != nil {
		return *x.Url
	}
	return ""
}

func (x *URLFetchRequest) GetHeader() []*URLFetchRequest_Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *URLFetchRequest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *URLFetchRequest) GetFollowRedirects() bool {
	if x != nil && x.FollowRedirects != nil {
		return *x.FollowRedirects
	}
	return Default_URLFetchRequest_FollowRedirects
}

func (x *URLFetchRequest) GetDeadline() float64 {
	if x != nil && x.Deadline != nil {
		return *x.Deadline
	}
	return 0
}

func (x *URLFetchRequest) GetMustValidateServerCertificate() bool {
	if x != nil && x.MustValidateServerCertificate != nil {
		return *x.MustValidateServerCertificate
	}
	return Default_URLFetchRequest_MustValidateServerCertificate
}

type URLFetchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content               []byte                     `protobuf:"bytes,1,opt,name=Content" json:"Content,omitempty"`
	StatusCode            *int32                     `protobuf:"varint,2,req,name=StatusCode" json:"StatusCode,omitempty"`
	Header                []*URLFetchResponse_Header `protobuf:"group,3,rep,name=Header,json=header" json:"header,omitempty"`
	ContentWasTruncated   *bool                      `protobuf:"varint,6,opt,name=ContentWasTruncated,def=0" json:"ContentWasTruncated,omitempty"`
	ExternalBytesSent     *int64                     `protobuf:"varint,7,opt,name=ExternalBytesSent" json:"ExternalBytesSent,omitempty"`
	ExternalBytesReceived *int64                     `protobuf:"varint,8,opt,name=ExternalBytesReceived" json:"ExternalBytesReceived,omitempty"`
	FinalUrl              *string                    `protobuf:"bytes,9,opt,name=FinalUrl" json:"FinalUrl,omitempty"`
	ApiCpuMilliseconds    *int64                     `protobuf:"varint,10,opt,name=ApiCpuMilliseconds,def=0" json:"ApiCpuMilliseconds,omitempty"`
	ApiBytesSent          *int64                     `protobuf:"varint,11,opt,name=ApiBytesSent,def=0" json:"ApiBytesSent,omitempty"`
	ApiBytesReceived      *int64                     `protobuf:"varint,12,opt,name=ApiBytesReceived,def=0" json:"ApiBytesReceived,omitempty"`
}

// Default values for URLFetchResponse fields.
const (
	Default_URLFetchResponse_ContentWasTruncated = bool(false)
	Default_URLFetchResponse_ApiCpuMilliseconds  = int64(0)
	Default_URLFetchResponse_ApiBytesSent        = int64(0)
	Default_URLFetchResponse_ApiBytesReceived    = int64(0)
)

func (x *URLFetchResponse) Reset() {
	*x = URLFetchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *URLFetchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URLFetchResponse) ProtoMessage() {}

func (x *URLFetchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URLFetchResponse.ProtoReflect.Descriptor instead.
func (*URLFetchResponse) Descriptor() ([]byte, []int) {
	return file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_rawDescGZIP(), []int{2}
}

func (x *URLFetchResponse) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *URLFetchResponse) GetStatusCode() int32 {
	if x != nil && x.StatusCode != nil {
		return *x.StatusCode
	}
	return 0
}

func (x *URLFetchResponse) GetHeader() []*URLFetchResponse_Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *URLFetchResponse) GetContentWasTruncated() bool {
	if x != nil && x.ContentWasTruncated != nil {
		return *x.ContentWasTruncated
	}
	return Default_URLFetchResponse_ContentWasTruncated
}

func (x *URLFetchResponse) GetExternalBytesSent() int64 {
	if x != nil && x.ExternalBytesSent != nil {
		return *x.ExternalBytesSent
	}
	return 0
}

func (x *URLFetchResponse) GetExternalBytesReceived() int64 {
	if x != nil && x.ExternalBytesReceived != nil {
		return *x.ExternalBytesReceived
	}
	return 0
}

func (x *URLFetchResponse) GetFinalUrl() string {
	if x != nil && x.FinalUrl != nil {
		return *x.FinalUrl
	}
	return ""
}

func (x *URLFetchResponse) GetApiCpuMilliseconds() int64 {
	if x != nil && x.ApiCpuMilliseconds != nil {
		return *x.ApiCpuMilliseconds
	}
	return Default_URLFetchResponse_ApiCpuMilliseconds
}

func (x *URLFetchResponse) GetApiBytesSent() int64 {
	if x != nil && x.ApiBytesSent != nil {
		return *x.ApiBytesSent
	}
	return Default_URLFetchResponse_ApiBytesSent
}

func (x *URLFetchResponse) GetApiBytesReceived() int64 {
	if x != nil && x.ApiBytesReceived != nil {
		return *x.ApiBytesReceived
	}
	return Default_URLFetchResponse_ApiBytesReceived
}

type URLFetchRequest_Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   *string `protobuf:"bytes,4,req,name=Key" json:"Key,omitempty"`
	Value *string `protobuf:"bytes,5,req,name=Value" json:"Value,omitempty"`
}

func (x *URLFetchRequest_Header) Reset() {
	*x = URLFetchRequest_Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *URLFetchRequest_Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URLFetchRequest_Header) ProtoMessage() {}

func (x *URLFetchRequest_Header) ProtoReflect() protoreflect.Message {
	mi := &file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URLFetchRequest_Header.ProtoReflect.Descriptor instead.
func (*URLFetchRequest_Header) Descriptor() ([]byte, []int) {
	return file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *URLFetchRequest_Header) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *URLFetchRequest_Header) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

type URLFetchResponse_Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   *string `protobuf:"bytes,4,req,name=Key" json:"Key,omitempty"`
	Value *string `protobuf:"bytes,5,req,name=Value" json:"Value,omitempty"`
}

func (x *URLFetchResponse_Header) Reset() {
	*x = URLFetchResponse_Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *URLFetchResponse_Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URLFetchResponse_Header) ProtoMessage() {}

func (x *URLFetchResponse_Header) ProtoReflect() protoreflect.Message {
	mi := &file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use URLFetchResponse_Header.ProtoReflect.Descriptor instead.
func (*URLFetchResponse_Header) Descriptor() ([]byte, []int) {
	return file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_rawDescGZIP(), []int{2, 0}
}

func (x *URLFetchResponse_Header) GetKey() string {
	if x != nil && x.Key != nil {
		return *x.Key
	}
	return ""
}

func (x *URLFetchResponse_Header) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

var File_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto protoreflect.FileDescriptor

var file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_rawDesc = []byte{
	0x0a, 0x4b, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x75, 0x72,
	0x6c, 0x66, 0x65, 0x74, 0x63, 0x68, 0x2f, 0x75, 0x72, 0x6c, 0x66, 0x65, 0x74, 0x63, 0x68, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61,
	0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x22, 0xab, 0x02, 0x0a, 0x14, 0x55, 0x52, 0x4c,
	0x46, 0x65, 0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x22, 0x92, 0x02, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x5f, 0x55, 0x52, 0x4c, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x45, 0x54, 0x43,
	0x48, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03,
	0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x54, 0x4f, 0x4f,
	0x5f, 0x4c, 0x41, 0x52, 0x47, 0x45, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x45, 0x41, 0x44,
	0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x45, 0x58, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x05, 0x12,
	0x19, 0x0a, 0x15, 0x53, 0x53, 0x4c, 0x5f, 0x43, 0x45, 0x52, 0x54, 0x49, 0x46, 0x49, 0x43, 0x41,
	0x54, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x4e,
	0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x07, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4c, 0x4f,
	0x53, 0x45, 0x44, 0x10, 0x08, 0x12, 0x1c, 0x0a, 0x18, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41,
	0x4c, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x09, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x59, 0x5f,
	0x52, 0x45, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x53, 0x10, 0x0a, 0x12, 0x13, 0x0a, 0x0f, 0x4d,
	0x41, 0x4c, 0x46, 0x4f, 0x52, 0x4d, 0x45, 0x44, 0x5f, 0x52, 0x45, 0x50, 0x4c, 0x59, 0x10, 0x0b,
	0x12, 0x14, 0x0a, 0x10, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0x0c, 0x22, 0xd6, 0x03, 0x0a, 0x0f, 0x55, 0x52, 0x4c, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x06, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x61, 0x70, 0x70,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x55, 0x52, 0x4c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x52, 0x06, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09, 0x52, 0x03, 0x55, 0x72, 0x6c, 0x12, 0x39,
	0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0a, 0x32, 0x21,
	0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x55, 0x52, 0x4c, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x07, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x02, 0x08, 0x01, 0x52, 0x07,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x2e, 0x0a, 0x0f, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x3a, 0x04, 0x74, 0x72, 0x75, 0x65, 0x52, 0x0f, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x65, 0x61, 0x64, 0x6c,
	0x69, 0x6e, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x44, 0x65, 0x61, 0x64, 0x6c,
	0x69, 0x6e, 0x65, 0x12, 0x4a, 0x0a, 0x1d, 0x4d, 0x75, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x04, 0x74, 0x72, 0x75, 0x65,
	0x52, 0x1d, 0x4d, 0x75, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x1a,
	0x30, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79,
	0x18, 0x04, 0x20, 0x02, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x4c, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x45, 0x54, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x50,
	0x4f, 0x53, 0x54, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x45, 0x41, 0x44, 0x10, 0x03, 0x12,
	0x07, 0x0a, 0x03, 0x50, 0x55, 0x54, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45,
	0x54, 0x45, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x41, 0x54, 0x43, 0x48, 0x10, 0x06, 0x22,
	0xfc, 0x03, 0x0a, 0x10, 0x55, 0x52, 0x4c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x02,
	0x28, 0x05, 0x52, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x3a,
	0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0a, 0x32, 0x22,
	0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x55, 0x52, 0x4c, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x13, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x57, 0x61, 0x73, 0x54, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x65,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x13,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x57, 0x61, 0x73, 0x54, 0x72, 0x75, 0x6e, 0x63, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x53, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x42, 0x79, 0x74, 0x65, 0x73, 0x53, 0x65, 0x6e,
	0x74, 0x12, 0x34, 0x0a, 0x15, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x15, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x42, 0x79, 0x74, 0x65, 0x73, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x61, 0x6c,
	0x55, 0x72, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6e, 0x61, 0x6c,
	0x55, 0x72, 0x6c, 0x12, 0x31, 0x0a, 0x12, 0x41, 0x70, 0x69, 0x43, 0x70, 0x75, 0x4d, 0x69, 0x6c,
	0x6c, 0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x3a,
	0x01, 0x30, 0x52, 0x12, 0x41, 0x70, 0x69, 0x43, 0x70, 0x75, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x25, 0x0a, 0x0c, 0x41, 0x70, 0x69, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x53, 0x65, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x3a, 0x01, 0x30, 0x52,
	0x0c, 0x41, 0x70, 0x69, 0x42, 0x79, 0x74, 0x65, 0x73, 0x53, 0x65, 0x6e, 0x74, 0x12, 0x2d, 0x0a,
	0x10, 0x41, 0x70, 0x69, 0x42, 0x79, 0x74, 0x65, 0x73, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65,
	0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x3a, 0x01, 0x30, 0x52, 0x10, 0x41, 0x70, 0x69, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x64, 0x1a, 0x30, 0x0a, 0x06,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x04, 0x20,
	0x02, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x05, 0x20, 0x02, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0xcf,
	0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x42, 0x14, 0x55, 0x72, 0x6c, 0x66, 0x65, 0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x64, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x6f, 0x72, 0x2d, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x63, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x75, 0x72, 0x6c, 0x66, 0x65, 0x74, 0x63, 0x68, 0xa2, 0x02,
	0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0xca, 0x02, 0x09, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0xe2, 0x02, 0x15, 0x41,
	0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
}

var (
	file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_rawDescOnce sync.Once
	file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_rawDescData = file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_rawDesc
)

func file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_rawDescGZIP() []byte {
	file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_rawDescOnce.Do(func() {
		file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_rawDescData)
	})
	return file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_rawDescData
}

var file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_goTypes = []interface{}{
	(URLFetchServiceError_ErrorCode)(0), // 0: appengine.URLFetchServiceError.ErrorCode
	(URLFetchRequest_RequestMethod)(0),  // 1: appengine.URLFetchRequest.RequestMethod
	(*URLFetchServiceError)(nil),        // 2: appengine.URLFetchServiceError
	(*URLFetchRequest)(nil),             // 3: appengine.URLFetchRequest
	(*URLFetchResponse)(nil),            // 4: appengine.URLFetchResponse
	(*URLFetchRequest_Header)(nil),      // 5: appengine.URLFetchRequest.Header
	(*URLFetchResponse_Header)(nil),     // 6: appengine.URLFetchResponse.Header
}
var file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_depIdxs = []int32{
	1, // 0: appengine.URLFetchRequest.Method:type_name -> appengine.URLFetchRequest.RequestMethod
	5, // 1: appengine.URLFetchRequest.header:type_name -> appengine.URLFetchRequest.Header
	6, // 2: appengine.URLFetchResponse.header:type_name -> appengine.URLFetchResponse.Header
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_init() }
func file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_init() {
	if File_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*URLFetchServiceError); i {
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
		file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*URLFetchRequest); i {
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
		file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*URLFetchResponse); i {
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
		file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*URLFetchRequest_Header); i {
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
		file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*URLFetchResponse_Header); i {
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
			RawDescriptor: file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_goTypes,
		DependencyIndexes: file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_depIdxs,
		EnumInfos:         file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_enumTypes,
		MessageInfos:      file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_msgTypes,
	}.Build()
	File_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto = out.File
	file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_rawDesc = nil
	file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_goTypes = nil
	file_vendor_google_golang_org_appengine_internal_urlfetch_urlfetch_service_proto_depIdxs = nil
}
