// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: svr/chart/v1/chart_db.proto

package chartv1

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Properties 基本属性
type Chart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repid      string `protobuf:"bytes,1,opt,name=repid,proto3" json:"repid,omitempty"`            // 表信息主键，node ×, java ×
	Techind01  string `protobuf:"bytes,2,opt,name=techind01,proto3" json:"techind01,omitempty"`    // 静态动态标识 [node] √, java √ [FX,RD,CD,MX]
	Techind02  string `protobuf:"bytes,3,opt,name=techind02,proto3" json:"techind02,omitempty"`    // 填表展示标识 node √, java √ [WB,RO,OT]
	Repdesc    string `protobuf:"bytes,4,opt,name=repdesc,proto3" json:"repdesc,omitempty"`        // 报表描述 [node] √, [java] √
	Repremark  string `protobuf:"bytes,5,opt,name=repremark,proto3" json:"repremark,omitempty"`    // 报表说明 [node] √, [java] √
	Repclass01 string `protobuf:"bytes,6,opt,name=repclass01,proto3" json:"repclass01,omitempty"`  // 报表类型 [node] √, java √ [01,02,03,04,05]
	Repclass02 string `protobuf:"bytes,7,opt,name=repclass02,proto3" json:"repclass02,omitempty"`  // 业务类型 [node] √, java √ [FL,FM,AL]
	Customerid string `protobuf:"bytes,8,opt,name=customerid,proto3" json:"customerid,omitempty"`  // 客户方表编码 [node] √, java √ [业务类型-YYYY-nnnn]
	Jqind      string `protobuf:"bytes,9,opt,name=jqind,proto3" json:"jqind,omitempty"`            // 久其标识[node] √, java √  [N,Y]
	Repjqid    string `protobuf:"bytes,10,opt,name=repjqid,proto3" json:"repjqid,omitempty"`       // 久其报表编码[node jqnId = Y] √, java √
	Multimodel string `protobuf:"bytes,11,opt,name=multimodel,proto3" json:"multimodel,omitempty"` // 多模型标识[node] √, java √
	Apprange   string `protobuf:"bytes,12,opt,name=apprange,proto3" json:"apprange,omitempty"`     // 适用范围[node] √, java √
	Repuser    string `protobuf:"bytes,13,opt,name=repuser,proto3" json:"repuser,omitempty"`       // 总部负责人 node √, java √
	Mdate      int32  `protobuf:"varint,14,opt,name=mdate,proto3" json:"mdate,omitempty"`          // 制表日期node ×, java √
	Mtime      int32  `protobuf:"varint,15,opt,name=mtime,proto3" json:"mtime,omitempty"`          // 制表时间node ×, java √
	Muser      string `protobuf:"bytes,16,opt,name=muser,proto3" json:"muser,omitempty"`           // 制表人姓名，node ×, java √ userInfo读取
}

func (x *Chart) Reset() {
	*x = Chart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svr_chart_v1_chart_db_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chart) ProtoMessage() {}

func (x *Chart) ProtoReflect() protoreflect.Message {
	mi := &file_svr_chart_v1_chart_db_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chart.ProtoReflect.Descriptor instead.
func (*Chart) Descriptor() ([]byte, []int) {
	return file_svr_chart_v1_chart_db_proto_rawDescGZIP(), []int{0}
}

func (x *Chart) GetRepid() string {
	if x != nil {
		return x.Repid
	}
	return ""
}

func (x *Chart) GetTechind01() string {
	if x != nil {
		return x.Techind01
	}
	return ""
}

func (x *Chart) GetTechind02() string {
	if x != nil {
		return x.Techind02
	}
	return ""
}

func (x *Chart) GetRepdesc() string {
	if x != nil {
		return x.Repdesc
	}
	return ""
}

func (x *Chart) GetRepremark() string {
	if x != nil {
		return x.Repremark
	}
	return ""
}

func (x *Chart) GetRepclass01() string {
	if x != nil {
		return x.Repclass01
	}
	return ""
}

func (x *Chart) GetRepclass02() string {
	if x != nil {
		return x.Repclass02
	}
	return ""
}

func (x *Chart) GetCustomerid() string {
	if x != nil {
		return x.Customerid
	}
	return ""
}

func (x *Chart) GetJqind() string {
	if x != nil {
		return x.Jqind
	}
	return ""
}

func (x *Chart) GetRepjqid() string {
	if x != nil {
		return x.Repjqid
	}
	return ""
}

func (x *Chart) GetMultimodel() string {
	if x != nil {
		return x.Multimodel
	}
	return ""
}

func (x *Chart) GetApprange() string {
	if x != nil {
		return x.Apprange
	}
	return ""
}

func (x *Chart) GetRepuser() string {
	if x != nil {
		return x.Repuser
	}
	return ""
}

func (x *Chart) GetMdate() int32 {
	if x != nil {
		return x.Mdate
	}
	return 0
}

func (x *Chart) GetMtime() int32 {
	if x != nil {
		return x.Mtime
	}
	return 0
}

func (x *Chart) GetMuser() string {
	if x != nil {
		return x.Muser
	}
	return ""
}

// Status 状态信息
type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repid     string `protobuf:"bytes,1,opt,name=repid,proto3" json:"repid,omitempty"`         // 表信息主键，node ×, java ×
	Repstatus string `protobuf:"bytes,2,opt,name=repstatus,proto3" json:"repstatus,omitempty"` // 表状态 node √, java √  [DF,RL,AB]
	Validfrom string `protobuf:"bytes,3,opt,name=validfrom,proto3" json:"validfrom,omitempty"` // 生效日期[node] √, java √
	Validto   string `protobuf:"bytes,4,opt,name=validto,proto3" json:"validto,omitempty"`     // 截至日期 [node] √, java √
	Mdate     int32  `protobuf:"varint,5,opt,name=mdate,proto3" json:"mdate,omitempty"`        // 制表日期node ×, java √
	Mtime     int32  `protobuf:"varint,6,opt,name=mtime,proto3" json:"mtime,omitempty"`        // 制表时间node ×, java √
	Muser     string `protobuf:"bytes,7,opt,name=muser,proto3" json:"muser,omitempty"`         // 制表人姓名，node ×, java √ userInfo读取
	Rdate     int32  `protobuf:"varint,8,opt,name=rdate,proto3" json:"rdate,omitempty"`        // 发布日期 node ×, java √
	Rtime     int32  `protobuf:"varint,9,opt,name=rtime,proto3" json:"rtime,omitempty"`        // 发布时间 node ×, java √
	Ruser     string `protobuf:"bytes,10,opt,name=ruser,proto3" json:"ruser,omitempty"`        // 发布人 node ×, java √ userInfo读取
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svr_chart_v1_chart_db_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_svr_chart_v1_chart_db_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_svr_chart_v1_chart_db_proto_rawDescGZIP(), []int{1}
}

func (x *Status) GetRepid() string {
	if x != nil {
		return x.Repid
	}
	return ""
}

func (x *Status) GetRepstatus() string {
	if x != nil {
		return x.Repstatus
	}
	return ""
}

func (x *Status) GetValidfrom() string {
	if x != nil {
		return x.Validfrom
	}
	return ""
}

func (x *Status) GetValidto() string {
	if x != nil {
		return x.Validto
	}
	return ""
}

func (x *Status) GetMdate() int32 {
	if x != nil {
		return x.Mdate
	}
	return 0
}

func (x *Status) GetMtime() int32 {
	if x != nil {
		return x.Mtime
	}
	return 0
}

func (x *Status) GetMuser() string {
	if x != nil {
		return x.Muser
	}
	return ""
}

func (x *Status) GetRdate() int32 {
	if x != nil {
		return x.Rdate
	}
	return 0
}

func (x *Status) GetRtime() int32 {
	if x != nil {
		return x.Rtime
	}
	return 0
}

func (x *Status) GetRuser() string {
	if x != nil {
		return x.Ruser
	}
	return ""
}

// Entity 适用单位
type Entity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repid    string `protobuf:"bytes,1,opt,name=repid,proto3" json:"repid,omitempty"`       // 表信息主键，node ×, java ×
	Hiername string `protobuf:"bytes,2,opt,name=hiername,proto3" json:"hiername,omitempty"` // 不同的架构 node √, java √  [H1,H2]
	Entity   string `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity,omitempty"`     // 适用单位 node √, java √
	Rdate    int32  `protobuf:"varint,8,opt,name=rdate,proto3" json:"rdate,omitempty"`      // 发布日期 node ×, java √
	Rtime    int32  `protobuf:"varint,9,opt,name=rtime,proto3" json:"rtime,omitempty"`      // 发布时间 node ×, java √
	Ruser    string `protobuf:"bytes,10,opt,name=ruser,proto3" json:"ruser,omitempty"`      // 发布人 node ×, java √ userInfo读取
}

func (x *Entity) Reset() {
	*x = Entity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svr_chart_v1_chart_db_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entity) ProtoMessage() {}

func (x *Entity) ProtoReflect() protoreflect.Message {
	mi := &file_svr_chart_v1_chart_db_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entity.ProtoReflect.Descriptor instead.
func (*Entity) Descriptor() ([]byte, []int) {
	return file_svr_chart_v1_chart_db_proto_rawDescGZIP(), []int{2}
}

func (x *Entity) GetRepid() string {
	if x != nil {
		return x.Repid
	}
	return ""
}

func (x *Entity) GetHiername() string {
	if x != nil {
		return x.Hiername
	}
	return ""
}

func (x *Entity) GetEntity() string {
	if x != nil {
		return x.Entity
	}
	return ""
}

func (x *Entity) GetRdate() int32 {
	if x != nil {
		return x.Rdate
	}
	return 0
}

func (x *Entity) GetRtime() int32 {
	if x != nil {
		return x.Rtime
	}
	return 0
}

func (x *Entity) GetRuser() string {
	if x != nil {
		return x.Ruser
	}
	return ""
}

// Style  表样式
type Style struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repid     string `protobuf:"bytes,1,opt,name=repid,proto3" json:"repid,omitempty"`         // 表信息主键，node ×, java ×
	Styleid   string `protobuf:"bytes,2,opt,name=styleid,proto3" json:"styleid,omitempty"`     // 表样式主键，node ×, java ×
	Styletext string `protobuf:"bytes,3,opt,name=styletext,proto3" json:"styletext,omitempty"` // 报表样式style node √，java √
	Mdate     int32  `protobuf:"varint,4,opt,name=mdate,proto3" json:"mdate,omitempty"`        // 制表日期node ×, java √
	Mtime     int32  `protobuf:"varint,5,opt,name=mtime,proto3" json:"mtime,omitempty"`        // 制表时间node ×, java √
	Muser     string `protobuf:"bytes,6,opt,name=muser,proto3" json:"muser,omitempty"`         // 制表人姓名，node ×, java √ userInfo读取
	Sheetid   string `protobuf:"bytes,7,opt,name=sheetid,proto3" json:"sheetid,omitempty"`
	Sheetname string `protobuf:"bytes,8,opt,name=sheetname,proto3" json:"sheetname,omitempty"`
}

func (x *Style) Reset() {
	*x = Style{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svr_chart_v1_chart_db_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Style) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Style) ProtoMessage() {}

func (x *Style) ProtoReflect() protoreflect.Message {
	mi := &file_svr_chart_v1_chart_db_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Style.ProtoReflect.Descriptor instead.
func (*Style) Descriptor() ([]byte, []int) {
	return file_svr_chart_v1_chart_db_proto_rawDescGZIP(), []int{3}
}

func (x *Style) GetRepid() string {
	if x != nil {
		return x.Repid
	}
	return ""
}

func (x *Style) GetStyleid() string {
	if x != nil {
		return x.Styleid
	}
	return ""
}

func (x *Style) GetStyletext() string {
	if x != nil {
		return x.Styletext
	}
	return ""
}

func (x *Style) GetMdate() int32 {
	if x != nil {
		return x.Mdate
	}
	return 0
}

func (x *Style) GetMtime() int32 {
	if x != nil {
		return x.Mtime
	}
	return 0
}

func (x *Style) GetMuser() string {
	if x != nil {
		return x.Muser
	}
	return ""
}

func (x *Style) GetSheetid() string {
	if x != nil {
		return x.Sheetid
	}
	return ""
}

func (x *Style) GetSheetname() string {
	if x != nil {
		return x.Sheetname
	}
	return ""
}

// Dimension 表配置
type Dimension struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repid         string `protobuf:"bytes,1,opt,name=repid,proto3" json:"repid,omitempty"`
	Dimension     string `protobuf:"bytes,2,opt,name=dimension,proto3" json:"dimension,omitempty"`
	Dimensiontext string `protobuf:"bytes,3,opt,name=dimensiontext,proto3" json:"dimensiontext,omitempty"`
	Muser         string `protobuf:"bytes,4,opt,name=muser,proto3" json:"muser,omitempty"`
	Spaceallow    int32  `protobuf:"varint,5,opt,name=spaceallow,proto3" json:"spaceallow,omitempty"`
}

func (x *Dimension) Reset() {
	*x = Dimension{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svr_chart_v1_chart_db_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dimension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dimension) ProtoMessage() {}

func (x *Dimension) ProtoReflect() protoreflect.Message {
	mi := &file_svr_chart_v1_chart_db_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dimension.ProtoReflect.Descriptor instead.
func (*Dimension) Descriptor() ([]byte, []int) {
	return file_svr_chart_v1_chart_db_proto_rawDescGZIP(), []int{4}
}

func (x *Dimension) GetRepid() string {
	if x != nil {
		return x.Repid
	}
	return ""
}

func (x *Dimension) GetDimension() string {
	if x != nil {
		return x.Dimension
	}
	return ""
}

func (x *Dimension) GetDimensiontext() string {
	if x != nil {
		return x.Dimensiontext
	}
	return ""
}

func (x *Dimension) GetMuser() string {
	if x != nil {
		return x.Muser
	}
	return ""
}

func (x *Dimension) GetSpaceallow() int32 {
	if x != nil {
		return x.Spaceallow
	}
	return 0
}

// ConInfo 落地信息
type ConInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repldid    string `protobuf:"bytes,1,opt,name=repldid,proto3" json:"repldid,omitempty"`
	Repid      string `protobuf:"bytes,2,opt,name=repid,proto3" json:"repid,omitempty"`
	Repdesc    string `protobuf:"bytes,3,opt,name=repdesc,proto3" json:"repdesc,omitempty"`
	Hbhir      string `protobuf:"bytes,4,opt,name=hbhir,proto3" json:"hbhir,omitempty"`
	Scope      string `protobuf:"bytes,5,opt,name=scope,proto3" json:"scope,omitempty"`
	Entity     string `protobuf:"bytes,6,opt,name=entity,proto3" json:"entity,omitempty"`
	Time       string `protobuf:"bytes,7,opt,name=time,proto3" json:"time,omitempty"`
	Jqind      string `protobuf:"bytes,8,opt,name=jqind,proto3" json:"jqind,omitempty"`
	Repjqid    string `protobuf:"bytes,9,opt,name=repjqid,proto3" json:"repjqid,omitempty"`
	Pushcnt    int32  `protobuf:"varint,10,opt,name=pushcnt,proto3" json:"pushcnt,omitempty"`
	Ptimestamp int64  `protobuf:"varint,11,opt,name=ptimestamp,proto3" json:"ptimestamp,omitempty"`
	Mdate      int64  `protobuf:"varint,12,opt,name=mdate,proto3" json:"mdate,omitempty"`
	Mtime      int64  `protobuf:"varint,13,opt,name=mtime,proto3" json:"mtime,omitempty"`
	Muser      string `protobuf:"bytes,14,opt,name=muser,proto3" json:"muser,omitempty"`
	Techind01  string `protobuf:"bytes,15,opt,name=techind01,proto3" json:"techind01,omitempty"`
	Auditid    string `protobuf:"bytes,16,opt,name=auditid,proto3" json:"auditid,omitempty"`
}

func (x *ConInfo) Reset() {
	*x = ConInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_svr_chart_v1_chart_db_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConInfo) ProtoMessage() {}

func (x *ConInfo) ProtoReflect() protoreflect.Message {
	mi := &file_svr_chart_v1_chart_db_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConInfo.ProtoReflect.Descriptor instead.
func (*ConInfo) Descriptor() ([]byte, []int) {
	return file_svr_chart_v1_chart_db_proto_rawDescGZIP(), []int{5}
}

func (x *ConInfo) GetRepldid() string {
	if x != nil {
		return x.Repldid
	}
	return ""
}

func (x *ConInfo) GetRepid() string {
	if x != nil {
		return x.Repid
	}
	return ""
}

func (x *ConInfo) GetRepdesc() string {
	if x != nil {
		return x.Repdesc
	}
	return ""
}

func (x *ConInfo) GetHbhir() string {
	if x != nil {
		return x.Hbhir
	}
	return ""
}

func (x *ConInfo) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *ConInfo) GetEntity() string {
	if x != nil {
		return x.Entity
	}
	return ""
}

func (x *ConInfo) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *ConInfo) GetJqind() string {
	if x != nil {
		return x.Jqind
	}
	return ""
}

func (x *ConInfo) GetRepjqid() string {
	if x != nil {
		return x.Repjqid
	}
	return ""
}

func (x *ConInfo) GetPushcnt() int32 {
	if x != nil {
		return x.Pushcnt
	}
	return 0
}

func (x *ConInfo) GetPtimestamp() int64 {
	if x != nil {
		return x.Ptimestamp
	}
	return 0
}

func (x *ConInfo) GetMdate() int64 {
	if x != nil {
		return x.Mdate
	}
	return 0
}

func (x *ConInfo) GetMtime() int64 {
	if x != nil {
		return x.Mtime
	}
	return 0
}

func (x *ConInfo) GetMuser() string {
	if x != nil {
		return x.Muser
	}
	return ""
}

func (x *ConInfo) GetTechind01() string {
	if x != nil {
		return x.Techind01
	}
	return ""
}

func (x *ConInfo) GetAuditid() string {
	if x != nil {
		return x.Auditid
	}
	return ""
}

var File_svr_chart_v1_chart_db_proto protoreflect.FileDescriptor

var file_svr_chart_v1_chart_db_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x73, 0x76, 0x72, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x68, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73,
	0x76, 0x72, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x22, 0xb9, 0x03, 0x0a, 0x05,
	0x43, 0x68, 0x61, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x65, 0x63, 0x68, 0x69, 0x6e, 0x64, 0x30, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x74, 0x65, 0x63, 0x68, 0x69, 0x6e, 0x64, 0x30, 0x31, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x63,
	0x68, 0x69, 0x6e, 0x64, 0x30, 0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65,
	0x63, 0x68, 0x69, 0x6e, 0x64, 0x30, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x64, 0x65,
	0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x70, 0x64, 0x65, 0x73,
	0x63, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x70, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12,
	0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x30, 0x31, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x30, 0x31, 0x12,
	0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x30, 0x32, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x30, 0x32, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x6a, 0x71, 0x69, 0x6e, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6a, 0x71, 0x69, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6a, 0x71, 0x69, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6a, 0x71, 0x69, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x61, 0x70, 0x70, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x70, 0x75, 0x73, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65,
	0x70, 0x75, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x75, 0x73, 0x65, 0x72, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6d, 0x75, 0x73, 0x65, 0x72, 0x22, 0xf8, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x70,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x66,
	0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x66, 0x72, 0x6f, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x74, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x74, 0x6f, 0x12, 0x14,
	0x0a, 0x05, 0x6d, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x72, 0x64, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x72, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x72, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x72, 0x75, 0x73, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x75, 0x73,
	0x65, 0x72, 0x22, 0x94, 0x01, 0x0a, 0x06, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x72, 0x65, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65,
	0x70, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x69, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x69, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x72, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x72, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x72, 0x74,
	0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x75, 0x73, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x72, 0x75, 0x73, 0x65, 0x72, 0x22, 0xcf, 0x01, 0x0a, 0x05, 0x53, 0x74,
	0x79, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x79,
	0x6c, 0x65, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x79, 0x6c,
	0x65, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6d, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x75, 0x73, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x68, 0x65, 0x65, 0x74, 0x69, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x68, 0x65, 0x65, 0x74, 0x69, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x68, 0x65, 0x65, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x68, 0x65, 0x65, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x9b, 0x01, 0x0a, 0x09,
	0x44, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x70,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x70, 0x69, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a,
	0x0d, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6d, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x22, 0x8f, 0x03, 0x0a, 0x07, 0x43, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x64, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x64, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x72, 0x65, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x72, 0x65, 0x70, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x64, 0x65, 0x73, 0x63,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x70, 0x64, 0x65, 0x73, 0x63, 0x12,
	0x14, 0x0a, 0x05, 0x68, 0x62, 0x68, 0x69, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x68, 0x62, 0x68, 0x69, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6a, 0x71, 0x69, 0x6e, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x71, 0x69, 0x6e, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x70, 0x6a, 0x71, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x72, 0x65, 0x70, 0x6a, 0x71, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x75, 0x73, 0x68, 0x63,
	0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x75, 0x73, 0x68, 0x63, 0x6e,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x6d, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x75, 0x73, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x63, 0x68, 0x69, 0x6e, 0x64, 0x30, 0x31,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x63, 0x68, 0x69, 0x6e, 0x64, 0x30,
	0x31, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x75, 0x64, 0x69, 0x74, 0x69, 0x64, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x75, 0x64, 0x69, 0x74, 0x69, 0x64, 0x42, 0x5c, 0x0a, 0x1b, 0x63,
	0x6f, 0x6d, 0x2e, 0x6e, 0x6e, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x63, 0x68, 0x61, 0x72, 0x74, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x08, 0x43, 0x68, 0x61, 0x72,
	0x74, 0x73, 0x44, 0x42, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x59, 0x65, 0x52, 0x75, 0x79, 0x69, 0x2f, 0x64, 0x66, 0x73, 0x2d, 0x67, 0x6f, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x76, 0x72, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x2f, 0x76,
	0x31, 0x3b, 0x63, 0x68, 0x61, 0x72, 0x74, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_svr_chart_v1_chart_db_proto_rawDescOnce sync.Once
	file_svr_chart_v1_chart_db_proto_rawDescData = file_svr_chart_v1_chart_db_proto_rawDesc
)

func file_svr_chart_v1_chart_db_proto_rawDescGZIP() []byte {
	file_svr_chart_v1_chart_db_proto_rawDescOnce.Do(func() {
		file_svr_chart_v1_chart_db_proto_rawDescData = protoimpl.X.CompressGZIP(file_svr_chart_v1_chart_db_proto_rawDescData)
	})
	return file_svr_chart_v1_chart_db_proto_rawDescData
}

var file_svr_chart_v1_chart_db_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_svr_chart_v1_chart_db_proto_goTypes = []interface{}{
	(*Chart)(nil),     // 0: svr.chart.v1.Chart
	(*Status)(nil),    // 1: svr.chart.v1.Status
	(*Entity)(nil),    // 2: svr.chart.v1.Entity
	(*Style)(nil),     // 3: svr.chart.v1.Style
	(*Dimension)(nil), // 4: svr.chart.v1.Dimension
	(*ConInfo)(nil),   // 5: svr.chart.v1.ConInfo
}
var file_svr_chart_v1_chart_db_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_svr_chart_v1_chart_db_proto_init() }
func file_svr_chart_v1_chart_db_proto_init() {
	if File_svr_chart_v1_chart_db_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_svr_chart_v1_chart_db_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chart); i {
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
		file_svr_chart_v1_chart_db_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_svr_chart_v1_chart_db_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entity); i {
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
		file_svr_chart_v1_chart_db_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Style); i {
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
		file_svr_chart_v1_chart_db_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dimension); i {
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
		file_svr_chart_v1_chart_db_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConInfo); i {
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
			RawDescriptor: file_svr_chart_v1_chart_db_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_svr_chart_v1_chart_db_proto_goTypes,
		DependencyIndexes: file_svr_chart_v1_chart_db_proto_depIdxs,
		MessageInfos:      file_svr_chart_v1_chart_db_proto_msgTypes,
	}.Build()
	File_svr_chart_v1_chart_db_proto = out.File
	file_svr_chart_v1_chart_db_proto_rawDesc = nil
	file_svr_chart_v1_chart_db_proto_goTypes = nil
	file_svr_chart_v1_chart_db_proto_depIdxs = nil
}
