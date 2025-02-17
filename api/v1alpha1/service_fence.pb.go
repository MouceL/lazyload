// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: service_fence.proto

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Destinations_Status int32

const (
	Destinations_ACTIVE Destinations_Status = 0
	Destinations_EXPIRE Destinations_Status = 1
	// In order to avoid deleting frequently, add status EXPIREWAIT between ACTIVE and EXPIRE.
	// When new metric does not contain host of ACTIVE status, its status will change to EXPIREWAIT. If new metric does not contain
	// host of EXPIREWAIT status, which means this host is not contained in the last two metrics, the status will change to EXPIRE.
	// Otherwise, EXPIREWAIT status will change back to ACTIVE.
	// Hosts of ACTIVE or EXPIREWAIT status are all valid for sidecar.
	// For prometheus metric source, as metric can continuously be watched, we can set status update interval in the future version,
	// refer to RecentlyCalled of RecyclingStrategy. But for accesslog metric source, metric only stores in lazyload controller memory.
	// Metric can not continuously produce after host added to sidecar. So after lazyload controller rebooting, we can not tell whether
	// old host is valid or not until it is removed from sidecar and goes to global-sidecar again.
	// We do not have a proper solution to do same thing for accesslog metric source so far. Need further thinking.
	Destinations_EXPIREWAIT Destinations_Status = 2
)

var Destinations_Status_name = map[int32]string{
	0: "ACTIVE",
	1: "EXPIRE",
	2: "EXPIREWAIT",
}

var Destinations_Status_value = map[string]int32{
	"ACTIVE":     0,
	"EXPIRE":     1,
	"EXPIREWAIT": 2,
}

func (x Destinations_Status) String() string {
	return proto.EnumName(Destinations_Status_name, int32(x))
}

func (Destinations_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b4b8d9f0db3c7310, []int{4, 0}
}

type Timestamp struct {
	// Represents seconds of UTC time since Unix epoch
	// 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to
	// 9999-12-31T23:59:59Z inclusive.
	Seconds int64 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	// Non-negative fractions of a second at nanosecond resolution. Negative
	// second values with fractions must still have non-negative nanos values
	// that count forward in time. Must be from 0 to 999,999,999
	// inclusive.
	Nanos                int32    `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Timestamp) Reset()         { *m = Timestamp{} }
func (m *Timestamp) String() string { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()    {}
func (*Timestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b8d9f0db3c7310, []int{0}
}
func (m *Timestamp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Timestamp.Unmarshal(m, b)
}
func (m *Timestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Timestamp.Marshal(b, m, deterministic)
}
func (m *Timestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timestamp.Merge(m, src)
}
func (m *Timestamp) XXX_Size() int {
	return xxx_messageInfo_Timestamp.Size(m)
}
func (m *Timestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_Timestamp.DiscardUnknown(m)
}

var xxx_messageInfo_Timestamp proto.InternalMessageInfo

func (m *Timestamp) GetSeconds() int64 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *Timestamp) GetNanos() int32 {
	if m != nil {
		return m.Nanos
	}
	return 0
}

// Spec Example
//   spec:
//    enable: true
//    host:
//      reviews.default.svc.cluster.local: # stable dependency of reviews.default service
//        stable:
//    namespaceSelector: # Match namespace names, multiple namespaces are 'or' relations, static dependency
//      - foo
//      - bar
//    labelSelector: # Match service label, multiple selectors are 'or' relationship, static dependency
//      - selector:
//          project: back
//      - selector: # labels in same selector are 'and' relationship
//          project: front
//          group: web
type ServiceFenceSpec struct {
	Host map[string]*RecyclingStrategy `protobuf:"bytes,1,rep,name=host,proto3" json:"host,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Switch to render servicefence as sidecar
	Enable bool `protobuf:"varint,2,opt,name=enable,proto3" json:"enable,omitempty"`
	// services in these namespaces are all static dependency, will not expire
	NamespaceSelector []string `protobuf:"bytes,3,rep,name=namespaceSelector,proto3" json:"namespaceSelector,omitempty"`
	// services match one selector of the label selector are all static dependency, will not expire
	LabelSelector        []*Selector `protobuf:"bytes,4,rep,name=labelSelector,proto3" json:"labelSelector,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ServiceFenceSpec) Reset()         { *m = ServiceFenceSpec{} }
func (m *ServiceFenceSpec) String() string { return proto.CompactTextString(m) }
func (*ServiceFenceSpec) ProtoMessage()    {}
func (*ServiceFenceSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b8d9f0db3c7310, []int{1}
}
func (m *ServiceFenceSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceFenceSpec.Unmarshal(m, b)
}
func (m *ServiceFenceSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceFenceSpec.Marshal(b, m, deterministic)
}
func (m *ServiceFenceSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceFenceSpec.Merge(m, src)
}
func (m *ServiceFenceSpec) XXX_Size() int {
	return xxx_messageInfo_ServiceFenceSpec.Size(m)
}
func (m *ServiceFenceSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceFenceSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceFenceSpec proto.InternalMessageInfo

func (m *ServiceFenceSpec) GetHost() map[string]*RecyclingStrategy {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *ServiceFenceSpec) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

func (m *ServiceFenceSpec) GetNamespaceSelector() []string {
	if m != nil {
		return m.NamespaceSelector
	}
	return nil
}

func (m *ServiceFenceSpec) GetLabelSelector() []*Selector {
	if m != nil {
		return m.LabelSelector
	}
	return nil
}

type Selector struct {
	Selector             map[string]string `protobuf:"bytes,1,rep,name=selector,proto3" json:"selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Selector) Reset()         { *m = Selector{} }
func (m *Selector) String() string { return proto.CompactTextString(m) }
func (*Selector) ProtoMessage()    {}
func (*Selector) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b8d9f0db3c7310, []int{2}
}
func (m *Selector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Selector.Unmarshal(m, b)
}
func (m *Selector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Selector.Marshal(b, m, deterministic)
}
func (m *Selector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Selector.Merge(m, src)
}
func (m *Selector) XXX_Size() int {
	return xxx_messageInfo_Selector.Size(m)
}
func (m *Selector) XXX_DiscardUnknown() {
	xxx_messageInfo_Selector.DiscardUnknown(m)
}

var xxx_messageInfo_Selector proto.InternalMessageInfo

func (m *Selector) GetSelector() map[string]string {
	if m != nil {
		return m.Selector
	}
	return nil
}

type RecyclingStrategy struct {
	// Configuration that will not be cleaned up
	Stable *RecyclingStrategy_Stable `protobuf:"bytes,1,opt,name=stable,proto3" json:"stable,omitempty"`
	// Configurations that expire after expiration
	Deadline *RecyclingStrategy_Deadline `protobuf:"bytes,2,opt,name=deadline,proto3" json:"deadline,omitempty"`
	// Deprecated
	Auto                 *RecyclingStrategy_Auto `protobuf:"bytes,3,opt,name=auto,proto3" json:"auto,omitempty"`
	RecentlyCalled       *Timestamp              `protobuf:"bytes,4,opt,name=RecentlyCalled,proto3" json:"RecentlyCalled,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *RecyclingStrategy) Reset()         { *m = RecyclingStrategy{} }
func (m *RecyclingStrategy) String() string { return proto.CompactTextString(m) }
func (*RecyclingStrategy) ProtoMessage()    {}
func (*RecyclingStrategy) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b8d9f0db3c7310, []int{3}
}
func (m *RecyclingStrategy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecyclingStrategy.Unmarshal(m, b)
}
func (m *RecyclingStrategy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecyclingStrategy.Marshal(b, m, deterministic)
}
func (m *RecyclingStrategy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecyclingStrategy.Merge(m, src)
}
func (m *RecyclingStrategy) XXX_Size() int {
	return xxx_messageInfo_RecyclingStrategy.Size(m)
}
func (m *RecyclingStrategy) XXX_DiscardUnknown() {
	xxx_messageInfo_RecyclingStrategy.DiscardUnknown(m)
}

var xxx_messageInfo_RecyclingStrategy proto.InternalMessageInfo

func (m *RecyclingStrategy) GetStable() *RecyclingStrategy_Stable {
	if m != nil {
		return m.Stable
	}
	return nil
}

func (m *RecyclingStrategy) GetDeadline() *RecyclingStrategy_Deadline {
	if m != nil {
		return m.Deadline
	}
	return nil
}

func (m *RecyclingStrategy) GetAuto() *RecyclingStrategy_Auto {
	if m != nil {
		return m.Auto
	}
	return nil
}

func (m *RecyclingStrategy) GetRecentlyCalled() *Timestamp {
	if m != nil {
		return m.RecentlyCalled
	}
	return nil
}

type RecyclingStrategy_Stable struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecyclingStrategy_Stable) Reset()         { *m = RecyclingStrategy_Stable{} }
func (m *RecyclingStrategy_Stable) String() string { return proto.CompactTextString(m) }
func (*RecyclingStrategy_Stable) ProtoMessage()    {}
func (*RecyclingStrategy_Stable) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b8d9f0db3c7310, []int{3, 0}
}
func (m *RecyclingStrategy_Stable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecyclingStrategy_Stable.Unmarshal(m, b)
}
func (m *RecyclingStrategy_Stable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecyclingStrategy_Stable.Marshal(b, m, deterministic)
}
func (m *RecyclingStrategy_Stable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecyclingStrategy_Stable.Merge(m, src)
}
func (m *RecyclingStrategy_Stable) XXX_Size() int {
	return xxx_messageInfo_RecyclingStrategy_Stable.Size(m)
}
func (m *RecyclingStrategy_Stable) XXX_DiscardUnknown() {
	xxx_messageInfo_RecyclingStrategy_Stable.DiscardUnknown(m)
}

var xxx_messageInfo_RecyclingStrategy_Stable proto.InternalMessageInfo

type RecyclingStrategy_Deadline struct {
	Expire               *Timestamp `protobuf:"bytes,1,opt,name=expire,proto3" json:"expire,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RecyclingStrategy_Deadline) Reset()         { *m = RecyclingStrategy_Deadline{} }
func (m *RecyclingStrategy_Deadline) String() string { return proto.CompactTextString(m) }
func (*RecyclingStrategy_Deadline) ProtoMessage()    {}
func (*RecyclingStrategy_Deadline) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b8d9f0db3c7310, []int{3, 1}
}
func (m *RecyclingStrategy_Deadline) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecyclingStrategy_Deadline.Unmarshal(m, b)
}
func (m *RecyclingStrategy_Deadline) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecyclingStrategy_Deadline.Marshal(b, m, deterministic)
}
func (m *RecyclingStrategy_Deadline) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecyclingStrategy_Deadline.Merge(m, src)
}
func (m *RecyclingStrategy_Deadline) XXX_Size() int {
	return xxx_messageInfo_RecyclingStrategy_Deadline.Size(m)
}
func (m *RecyclingStrategy_Deadline) XXX_DiscardUnknown() {
	xxx_messageInfo_RecyclingStrategy_Deadline.DiscardUnknown(m)
}

var xxx_messageInfo_RecyclingStrategy_Deadline proto.InternalMessageInfo

func (m *RecyclingStrategy_Deadline) GetExpire() *Timestamp {
	if m != nil {
		return m.Expire
	}
	return nil
}

type RecyclingStrategy_Auto struct {
	Duration             *Timestamp `protobuf:"bytes,1,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RecyclingStrategy_Auto) Reset()         { *m = RecyclingStrategy_Auto{} }
func (m *RecyclingStrategy_Auto) String() string { return proto.CompactTextString(m) }
func (*RecyclingStrategy_Auto) ProtoMessage()    {}
func (*RecyclingStrategy_Auto) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b8d9f0db3c7310, []int{3, 2}
}
func (m *RecyclingStrategy_Auto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecyclingStrategy_Auto.Unmarshal(m, b)
}
func (m *RecyclingStrategy_Auto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecyclingStrategy_Auto.Marshal(b, m, deterministic)
}
func (m *RecyclingStrategy_Auto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecyclingStrategy_Auto.Merge(m, src)
}
func (m *RecyclingStrategy_Auto) XXX_Size() int {
	return xxx_messageInfo_RecyclingStrategy_Auto.Size(m)
}
func (m *RecyclingStrategy_Auto) XXX_DiscardUnknown() {
	xxx_messageInfo_RecyclingStrategy_Auto.DiscardUnknown(m)
}

var xxx_messageInfo_RecyclingStrategy_Auto proto.InternalMessageInfo

func (m *RecyclingStrategy_Auto) GetDuration() *Timestamp {
	if m != nil {
		return m.Duration
	}
	return nil
}

type Destinations struct {
	// Deprecated
	RecentlyCalled       *Timestamp          `protobuf:"bytes,1,opt,name=RecentlyCalled,proto3" json:"RecentlyCalled,omitempty"`
	Hosts                []string            `protobuf:"bytes,2,rep,name=hosts,proto3" json:"hosts,omitempty"`
	Status               Destinations_Status `protobuf:"varint,3,opt,name=status,proto3,enum=slime.microservice.lazyload.v1alpha1.Destinations_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Destinations) Reset()         { *m = Destinations{} }
func (m *Destinations) String() string { return proto.CompactTextString(m) }
func (*Destinations) ProtoMessage()    {}
func (*Destinations) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b8d9f0db3c7310, []int{4}
}
func (m *Destinations) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Destinations.Unmarshal(m, b)
}
func (m *Destinations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Destinations.Marshal(b, m, deterministic)
}
func (m *Destinations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Destinations.Merge(m, src)
}
func (m *Destinations) XXX_Size() int {
	return xxx_messageInfo_Destinations.Size(m)
}
func (m *Destinations) XXX_DiscardUnknown() {
	xxx_messageInfo_Destinations.DiscardUnknown(m)
}

var xxx_messageInfo_Destinations proto.InternalMessageInfo

func (m *Destinations) GetRecentlyCalled() *Timestamp {
	if m != nil {
		return m.RecentlyCalled
	}
	return nil
}

func (m *Destinations) GetHosts() []string {
	if m != nil {
		return m.Hosts
	}
	return nil
}

func (m *Destinations) GetStatus() Destinations_Status {
	if m != nil {
		return m.Status
	}
	return Destinations_ACTIVE
}

type ServiceFenceStatus struct {
	Domains              map[string]*Destinations `protobuf:"bytes,1,rep,name=domains,proto3" json:"domains,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	MetricStatus         map[string]string        `protobuf:"bytes,3,rep,name=metricStatus,proto3" json:"metricStatus,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Visitor              map[string]bool          `protobuf:"bytes,2,rep,name=visitor,proto3" json:"visitor,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ServiceFenceStatus) Reset()         { *m = ServiceFenceStatus{} }
func (m *ServiceFenceStatus) String() string { return proto.CompactTextString(m) }
func (*ServiceFenceStatus) ProtoMessage()    {}
func (*ServiceFenceStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4b8d9f0db3c7310, []int{5}
}
func (m *ServiceFenceStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceFenceStatus.Unmarshal(m, b)
}
func (m *ServiceFenceStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceFenceStatus.Marshal(b, m, deterministic)
}
func (m *ServiceFenceStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceFenceStatus.Merge(m, src)
}
func (m *ServiceFenceStatus) XXX_Size() int {
	return xxx_messageInfo_ServiceFenceStatus.Size(m)
}
func (m *ServiceFenceStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceFenceStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceFenceStatus proto.InternalMessageInfo

func (m *ServiceFenceStatus) GetDomains() map[string]*Destinations {
	if m != nil {
		return m.Domains
	}
	return nil
}

func (m *ServiceFenceStatus) GetMetricStatus() map[string]string {
	if m != nil {
		return m.MetricStatus
	}
	return nil
}

func (m *ServiceFenceStatus) GetVisitor() map[string]bool {
	if m != nil {
		return m.Visitor
	}
	return nil
}

func init() {
	proto.RegisterEnum("slime.microservice.lazyload.v1alpha1.Destinations_Status", Destinations_Status_name, Destinations_Status_value)
	proto.RegisterType((*Timestamp)(nil), "slime.microservice.lazyload.v1alpha1.Timestamp")
	proto.RegisterType((*ServiceFenceSpec)(nil), "slime.microservice.lazyload.v1alpha1.ServiceFenceSpec")
	proto.RegisterMapType((map[string]*RecyclingStrategy)(nil), "slime.microservice.lazyload.v1alpha1.ServiceFenceSpec.HostEntry")
	proto.RegisterType((*Selector)(nil), "slime.microservice.lazyload.v1alpha1.Selector")
	proto.RegisterMapType((map[string]string)(nil), "slime.microservice.lazyload.v1alpha1.Selector.SelectorEntry")
	proto.RegisterType((*RecyclingStrategy)(nil), "slime.microservice.lazyload.v1alpha1.RecyclingStrategy")
	proto.RegisterType((*RecyclingStrategy_Stable)(nil), "slime.microservice.lazyload.v1alpha1.RecyclingStrategy.Stable")
	proto.RegisterType((*RecyclingStrategy_Deadline)(nil), "slime.microservice.lazyload.v1alpha1.RecyclingStrategy.Deadline")
	proto.RegisterType((*RecyclingStrategy_Auto)(nil), "slime.microservice.lazyload.v1alpha1.RecyclingStrategy.Auto")
	proto.RegisterType((*Destinations)(nil), "slime.microservice.lazyload.v1alpha1.Destinations")
	proto.RegisterType((*ServiceFenceStatus)(nil), "slime.microservice.lazyload.v1alpha1.ServiceFenceStatus")
	proto.RegisterMapType((map[string]*Destinations)(nil), "slime.microservice.lazyload.v1alpha1.ServiceFenceStatus.DomainsEntry")
	proto.RegisterMapType((map[string]string)(nil), "slime.microservice.lazyload.v1alpha1.ServiceFenceStatus.MetricStatusEntry")
	proto.RegisterMapType((map[string]bool)(nil), "slime.microservice.lazyload.v1alpha1.ServiceFenceStatus.VisitorEntry")
}

func init() { proto.RegisterFile("service_fence.proto", fileDescriptor_b4b8d9f0db3c7310) }

var fileDescriptor_b4b8d9f0db3c7310 = []byte{
	// 696 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0xc1, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0x49, 0xdb, 0x75, 0xe9, 0x5b, 0x37, 0x75, 0x66, 0x42, 0x55, 0x4e, 0x55, 0xc5, 0xa1,
	0x87, 0x29, 0x65, 0xe5, 0x00, 0x6c, 0x13, 0x6c, 0x6c, 0x85, 0x0d, 0x34, 0x69, 0x38, 0xd5, 0x36,
	0x21, 0xa4, 0xc9, 0x4b, 0xcc, 0x16, 0x91, 0xd8, 0x51, 0xec, 0x54, 0x94, 0x23, 0xff, 0x09, 0x37,
	0xee, 0xfc, 0x6f, 0x9c, 0x51, 0xec, 0x24, 0x4a, 0xb7, 0x49, 0xb4, 0x85, 0x9b, 0xad, 0xf4, 0xfd,
	0xde, 0xfb, 0x3e, 0xdb, 0x9f, 0x0a, 0x0f, 0x05, 0x8d, 0xc7, 0xbe, 0x4b, 0x2f, 0x3f, 0x53, 0xe6,
	0x52, 0x3b, 0x8a, 0xb9, 0xe4, 0xe8, 0xb1, 0x08, 0xfc, 0x90, 0xda, 0xa1, 0xef, 0xc6, 0x3c, 0xfb,
	0x6e, 0x07, 0xe4, 0xdb, 0x24, 0xe0, 0xc4, 0xb3, 0xc7, 0x5b, 0x24, 0x88, 0x6e, 0xc8, 0x56, 0x77,
	0x07, 0x1a, 0x23, 0x3f, 0xa4, 0x42, 0x92, 0x30, 0x42, 0x6d, 0x58, 0x16, 0xd4, 0xe5, 0xcc, 0x13,
	0x6d, 0xa3, 0x63, 0xf4, 0xaa, 0x38, 0xdf, 0xa2, 0x0d, 0x58, 0x62, 0x84, 0x71, 0xd1, 0xae, 0x74,
	0x8c, 0xde, 0x12, 0xd6, 0x9b, 0xee, 0xef, 0x0a, 0xb4, 0x1c, 0x8d, 0x7e, 0x93, 0x76, 0x76, 0x22,
	0xea, 0xa2, 0x11, 0xd4, 0x6e, 0xb8, 0x90, 0x6d, 0xa3, 0x53, 0xed, 0xad, 0x0c, 0xf6, 0xec, 0x59,
	0xc6, 0xb0, 0x6f, 0x53, 0xec, 0x23, 0x2e, 0xe4, 0x90, 0xc9, 0x78, 0x82, 0x15, 0x0d, 0x3d, 0x82,
	0x3a, 0x65, 0xe4, 0x2a, 0xa0, 0x6a, 0x02, 0x13, 0x67, 0x3b, 0xb4, 0x09, 0xeb, 0x8c, 0x84, 0x54,
	0x44, 0xc4, 0xa5, 0x0e, 0x0d, 0xa8, 0x2b, 0x79, 0xdc, 0xae, 0x76, 0xaa, 0xbd, 0x06, 0xbe, 0xfb,
	0x01, 0x8d, 0x60, 0x35, 0x20, 0x57, 0x34, 0x28, 0x7e, 0x59, 0x53, 0x43, 0xda, 0xb3, 0x0e, 0xa9,
	0xab, 0xf0, 0x34, 0xc4, 0x8a, 0xa0, 0x51, 0x8c, 0x8b, 0x5a, 0x50, 0xfd, 0x42, 0x27, 0xca, 0xbf,
	0x06, 0x4e, 0x97, 0xe8, 0x04, 0x96, 0xc6, 0x24, 0x48, 0xf4, 0xe4, 0x2b, 0x83, 0x67, 0xb3, 0x35,
	0xc3, 0xd4, 0x9d, 0xb8, 0x81, 0xcf, 0xae, 0x1d, 0x19, 0x13, 0x49, 0xaf, 0x27, 0x58, 0x53, 0xb6,
	0x2b, 0xcf, 0x8d, 0xee, 0x0f, 0x03, 0xcc, 0x42, 0xd4, 0x05, 0x98, 0x22, 0xd7, 0xa3, 0x4d, 0xdf,
	0x9d, 0x4f, 0x4f, 0xb1, 0xd0, 0x86, 0x17, 0x34, 0x6b, 0x07, 0x56, 0xa7, 0x3e, 0xdd, 0x23, 0x6e,
	0xa3, 0x2c, 0xae, 0x51, 0x9e, 0xf1, 0x67, 0x0d, 0xd6, 0xef, 0x88, 0x40, 0x67, 0x50, 0x17, 0x52,
	0x9d, 0xa3, 0xa1, 0xdc, 0x78, 0xb9, 0xa0, 0x1b, 0xb6, 0xa3, 0x28, 0x38, 0xa3, 0xa1, 0x4f, 0x60,
	0x7a, 0x94, 0x78, 0x81, 0xcf, 0x72, 0x9f, 0xf7, 0x16, 0x25, 0x1f, 0x66, 0x1c, 0x5c, 0x10, 0xd1,
	0x29, 0xd4, 0x48, 0x22, 0x79, 0xbb, 0xaa, 0xc8, 0xbb, 0x8b, 0x92, 0xf7, 0x13, 0xc9, 0xb1, 0x22,
	0xa1, 0x73, 0x58, 0xc3, 0xd4, 0xa5, 0x4c, 0x06, 0x93, 0x03, 0x12, 0x04, 0xd4, 0x6b, 0xd7, 0x14,
	0xbb, 0x3f, 0x1b, 0xbb, 0x78, 0xb3, 0xf8, 0x16, 0xc6, 0x32, 0xa1, 0xae, 0xad, 0xb1, 0x1c, 0x30,
	0x73, 0x29, 0xe8, 0x2d, 0xd4, 0xe9, 0xd7, 0xc8, 0x8f, 0x73, 0xdb, 0xe7, 0x6e, 0x93, 0x95, 0x5b,
	0x0e, 0xd4, 0x52, 0x15, 0xe8, 0x3d, 0x98, 0x5e, 0x12, 0x13, 0xe9, 0x73, 0xb6, 0x28, 0xb2, 0x00,
	0x74, 0xbf, 0x57, 0xa0, 0x79, 0x48, 0x85, 0xf4, 0x99, 0xda, 0x8b, 0x7b, 0xdc, 0x31, 0xfe, 0x8b,
	0x3b, 0xe9, 0x75, 0x4d, 0xe3, 0x24, 0xcd, 0xb1, 0x34, 0x22, 0xf4, 0x06, 0x7d, 0x50, 0x97, 0x52,
	0x26, 0x42, 0x1d, 0xf0, 0xda, 0xe0, 0xc5, 0x6c, 0x6d, 0xca, 0x23, 0xa7, 0xf7, 0x51, 0x26, 0x02,
	0x67, 0xa0, 0xee, 0x13, 0x75, 0x0c, 0x32, 0x11, 0x08, 0xa0, 0xbe, 0x7f, 0x30, 0x3a, 0x3e, 0x1b,
	0xb6, 0x1e, 0xa4, 0xeb, 0xe1, 0xc5, 0xe9, 0x31, 0x1e, 0xb6, 0x0c, 0xb4, 0x06, 0xa0, 0xd7, 0xe7,
	0xfb, 0xc7, 0xa3, 0x56, 0xa5, 0xfb, 0xab, 0x06, 0x68, 0x2a, 0x06, 0x75, 0xf9, 0x25, 0x2c, 0x7b,
	0x3c, 0x24, 0x3e, 0x13, 0xd9, 0xe3, 0x1e, 0x2e, 0x90, 0xa8, 0x0a, 0x65, 0x1f, 0x6a, 0x8e, 0x7e,
	0xe5, 0x39, 0x15, 0x31, 0x68, 0x86, 0x54, 0xc6, 0xbe, 0xeb, 0xe4, 0x16, 0xa4, 0x5d, 0xde, 0x2d,
	0xdc, 0xe5, 0xa4, 0x04, 0xd3, 0xad, 0xa6, 0xf8, 0xa9, 0xa0, 0xb1, 0x2f, 0xfc, 0x34, 0xad, 0x2a,
	0xff, 0x28, 0xe8, 0x4c, 0x73, 0x32, 0x41, 0x19, 0xd5, 0x62, 0xd0, 0x2c, 0x2b, 0xbd, 0x27, 0xb4,
	0x8e, 0xa6, 0x13, 0x79, 0x30, 0xff, 0x71, 0x97, 0x82, 0xce, 0x7a, 0x05, 0xeb, 0x77, 0x34, 0xcf,
	0x93, 0x94, 0xd6, 0x36, 0x34, 0xcb, 0x4a, 0xfe, 0x56, 0x6b, 0x96, 0x6a, 0x5f, 0xdb, 0x1f, 0x37,
	0xf5, 0xf0, 0x3e, 0xef, 0xab, 0x45, 0x3f, 0xe4, 0x5e, 0x12, 0x50, 0xd1, 0xcf, 0x05, 0xf4, 0x49,
	0xe4, 0xf7, 0x73, 0x11, 0x57, 0x75, 0xf5, 0xe7, 0xe0, 0xe9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x5d, 0x52, 0x38, 0x22, 0x33, 0x08, 0x00, 0x00,
}
