// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: proto/message_service.proto

package pb

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

type FileData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type      string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	CreatedAt string `protobuf:"bytes,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Customer  string `protobuf:"bytes,4,opt,name=customer,proto3" json:"customer,omitempty"`
}

func (x *FileData) Reset() {
	*x = FileData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileData) ProtoMessage() {}

func (x *FileData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileData.ProtoReflect.Descriptor instead.
func (*FileData) Descriptor() ([]byte, []int) {
	return file_proto_message_service_proto_rawDescGZIP(), []int{0}
}

func (x *FileData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FileData) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *FileData) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *FileData) GetCustomer() string {
	if x != nil {
		return x.Customer
	}
	return ""
}

type ContractData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileData *FileData `protobuf:"bytes,1,opt,name=fileData,proto3" json:"fileData,omitempty"`
	Title    string    `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Parties  []string  `protobuf:"bytes,3,rep,name=parties,proto3" json:"parties,omitempty"`
	Object   string    `protobuf:"bytes,4,opt,name=object,proto3" json:"object,omitempty"`
	Extract  []string  `protobuf:"bytes,5,rep,name=extract,proto3" json:"extract,omitempty"`
	Invoice  []string  `protobuf:"bytes,6,rep,name=invoice,proto3" json:"invoice,omitempty"`
}

func (x *ContractData) Reset() {
	*x = ContractData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractData) ProtoMessage() {}

func (x *ContractData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractData.ProtoReflect.Descriptor instead.
func (*ContractData) Descriptor() ([]byte, []int) {
	return file_proto_message_service_proto_rawDescGZIP(), []int{1}
}

func (x *ContractData) GetFileData() *FileData {
	if x != nil {
		return x.FileData
	}
	return nil
}

func (x *ContractData) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ContractData) GetParties() []string {
	if x != nil {
		return x.Parties
	}
	return nil
}

func (x *ContractData) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *ContractData) GetExtract() []string {
	if x != nil {
		return x.Extract
	}
	return nil
}

func (x *ContractData) GetInvoice() []string {
	if x != nil {
		return x.Invoice
	}
	return nil
}

type ExtractData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileData *FileData `protobuf:"bytes,1,opt,name=fileData,proto3" json:"fileData,omitempty"`
	Account  int32     `protobuf:"varint,2,opt,name=account,proto3" json:"account,omitempty"`
	Value    float64   `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
	Category string    `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty"`
	Method   string    `protobuf:"bytes,5,opt,name=method,proto3" json:"method,omitempty"`
	Location string    `protobuf:"bytes,6,opt,name=location,proto3" json:"location,omitempty"`
	Contract string    `protobuf:"bytes,7,opt,name=contract,proto3" json:"contract,omitempty"`
}

func (x *ExtractData) Reset() {
	*x = ExtractData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtractData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtractData) ProtoMessage() {}

func (x *ExtractData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtractData.ProtoReflect.Descriptor instead.
func (*ExtractData) Descriptor() ([]byte, []int) {
	return file_proto_message_service_proto_rawDescGZIP(), []int{2}
}

func (x *ExtractData) GetFileData() *FileData {
	if x != nil {
		return x.FileData
	}
	return nil
}

func (x *ExtractData) GetAccount() int32 {
	if x != nil {
		return x.Account
	}
	return 0
}

func (x *ExtractData) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *ExtractData) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *ExtractData) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *ExtractData) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *ExtractData) GetContract() string {
	if x != nil {
		return x.Contract
	}
	return ""
}

type InvoiceData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileData *FileData `protobuf:"bytes,1,opt,name=fileData,proto3" json:"fileData,omitempty"`
	DueDate  string    `protobuf:"bytes,2,opt,name=dueDate,proto3" json:"dueDate,omitempty"`
	Value    float64   `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
	Method   string    `protobuf:"bytes,4,opt,name=method,proto3" json:"method,omitempty"`
	Contract string    `protobuf:"bytes,5,opt,name=contract,proto3" json:"contract,omitempty"`
}

func (x *InvoiceData) Reset() {
	*x = InvoiceData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvoiceData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvoiceData) ProtoMessage() {}

func (x *InvoiceData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvoiceData.ProtoReflect.Descriptor instead.
func (*InvoiceData) Descriptor() ([]byte, []int) {
	return file_proto_message_service_proto_rawDescGZIP(), []int{3}
}

func (x *InvoiceData) GetFileData() *FileData {
	if x != nil {
		return x.FileData
	}
	return nil
}

func (x *InvoiceData) GetDueDate() string {
	if x != nil {
		return x.DueDate
	}
	return ""
}

func (x *InvoiceData) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *InvoiceData) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *InvoiceData) GetContract() string {
	if x != nil {
		return x.Contract
	}
	return ""
}

type Contract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File     []byte        `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	Metadata *ContractData `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *Contract) Reset() {
	*x = Contract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contract) ProtoMessage() {}

func (x *Contract) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contract.ProtoReflect.Descriptor instead.
func (*Contract) Descriptor() ([]byte, []int) {
	return file_proto_message_service_proto_rawDescGZIP(), []int{4}
}

func (x *Contract) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *Contract) GetMetadata() *ContractData {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type Extract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File     []byte       `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	Metadata *ExtractData `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *Extract) Reset() {
	*x = Extract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Extract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Extract) ProtoMessage() {}

func (x *Extract) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Extract.ProtoReflect.Descriptor instead.
func (*Extract) Descriptor() ([]byte, []int) {
	return file_proto_message_service_proto_rawDescGZIP(), []int{5}
}

func (x *Extract) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *Extract) GetMetadata() *ExtractData {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type Invoice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File     []byte       `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	Metadata *InvoiceData `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *Invoice) Reset() {
	*x = Invoice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Invoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Invoice) ProtoMessage() {}

func (x *Invoice) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Invoice.ProtoReflect.Descriptor instead.
func (*Invoice) Descriptor() ([]byte, []int) {
	return file_proto_message_service_proto_rawDescGZIP(), []int{6}
}

func (x *Invoice) GetFile() []byte {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *Invoice) GetMetadata() *InvoiceData {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_message_service_proto_rawDescGZIP(), []int{7}
}

func (x *Response) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_message_service_proto protoreflect.FileDescriptor

var file_proto_message_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a,
	0x08, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x22, 0xb1, 0x01, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x65, 0x78, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x22, 0xd0, 0x01, 0x0a, 0x0b,
	0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x22, 0x98,
	0x01, 0x0a, 0x0b, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x25,
	0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x22, 0x49, 0x0a, 0x08, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x47, 0x0a, 0x07, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66,
	0x69, 0x6c, 0x65, 0x12, 0x28, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x47, 0x0a,
	0x07, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x28, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3e, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x3c, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x0f, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x09, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x32, 0x39, 0x0a, 0x0e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x08, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32,
	0x39, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x27, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x08, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x1a, 0x09, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_message_service_proto_rawDescOnce sync.Once
	file_proto_message_service_proto_rawDescData = file_proto_message_service_proto_rawDesc
)

func file_proto_message_service_proto_rawDescGZIP() []byte {
	file_proto_message_service_proto_rawDescOnce.Do(func() {
		file_proto_message_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_message_service_proto_rawDescData)
	})
	return file_proto_message_service_proto_rawDescData
}

var file_proto_message_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_message_service_proto_goTypes = []interface{}{
	(*FileData)(nil),     // 0: FileData
	(*ContractData)(nil), // 1: ContractData
	(*ExtractData)(nil),  // 2: ExtractData
	(*InvoiceData)(nil),  // 3: InvoiceData
	(*Contract)(nil),     // 4: Contract
	(*Extract)(nil),      // 5: Extract
	(*Invoice)(nil),      // 6: Invoice
	(*Response)(nil),     // 7: Response
}
var file_proto_message_service_proto_depIdxs = []int32{
	0, // 0: ContractData.fileData:type_name -> FileData
	0, // 1: ExtractData.fileData:type_name -> FileData
	0, // 2: InvoiceData.fileData:type_name -> FileData
	1, // 3: Contract.metadata:type_name -> ContractData
	2, // 4: Extract.metadata:type_name -> ExtractData
	3, // 5: Invoice.metadata:type_name -> InvoiceData
	4, // 6: ContractRequest.ContractRequest:input_type -> Contract
	5, // 7: ExtractRequest.ExtractRequest:input_type -> Extract
	6, // 8: InvoiceRequest.InvoiceRequest:input_type -> Invoice
	7, // 9: ContractRequest.ContractRequest:output_type -> Response
	7, // 10: ExtractRequest.ExtractRequest:output_type -> Response
	7, // 11: InvoiceRequest.InvoiceRequest:output_type -> Response
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_message_service_proto_init() }
func file_proto_message_service_proto_init() {
	if File_proto_message_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_message_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileData); i {
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
		file_proto_message_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractData); i {
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
		file_proto_message_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtractData); i {
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
		file_proto_message_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvoiceData); i {
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
		file_proto_message_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contract); i {
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
		file_proto_message_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Extract); i {
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
		file_proto_message_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Invoice); i {
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
		file_proto_message_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_proto_message_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_proto_message_service_proto_goTypes,
		DependencyIndexes: file_proto_message_service_proto_depIdxs,
		MessageInfos:      file_proto_message_service_proto_msgTypes,
	}.Build()
	File_proto_message_service_proto = out.File
	file_proto_message_service_proto_rawDesc = nil
	file_proto_message_service_proto_goTypes = nil
	file_proto_message_service_proto_depIdxs = nil
}