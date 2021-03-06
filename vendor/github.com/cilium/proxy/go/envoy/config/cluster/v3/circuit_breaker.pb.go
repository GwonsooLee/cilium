// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/cluster/v3/circuit_breaker.proto

package envoy_config_cluster_v3

import (
	fmt "fmt"
	v3 "github.com/cilium/proxy/go/envoy/config/core/v3"
	v31 "github.com/cilium/proxy/go/envoy/type/v3"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// :ref:`Circuit breaking<arch_overview_circuit_break>` settings can be
// specified individually for each defined priority.
type CircuitBreakers struct {
	// If multiple :ref:`Thresholds<envoy_api_msg_config.cluster.v3.CircuitBreakers.Thresholds>`
	// are defined with the same :ref:`RoutingPriority<envoy_api_enum_config.core.v3.RoutingPriority>`,
	// the first one in the list is used. If no Thresholds is defined for a given
	// :ref:`RoutingPriority<envoy_api_enum_config.core.v3.RoutingPriority>`, the default values
	// are used.
	Thresholds           []*CircuitBreakers_Thresholds `protobuf:"bytes,1,rep,name=thresholds,proto3" json:"thresholds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *CircuitBreakers) Reset()         { *m = CircuitBreakers{} }
func (m *CircuitBreakers) String() string { return proto.CompactTextString(m) }
func (*CircuitBreakers) ProtoMessage()    {}
func (*CircuitBreakers) Descriptor() ([]byte, []int) {
	return fileDescriptor_131e6721a0838844, []int{0}
}

func (m *CircuitBreakers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CircuitBreakers.Unmarshal(m, b)
}
func (m *CircuitBreakers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CircuitBreakers.Marshal(b, m, deterministic)
}
func (m *CircuitBreakers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CircuitBreakers.Merge(m, src)
}
func (m *CircuitBreakers) XXX_Size() int {
	return xxx_messageInfo_CircuitBreakers.Size(m)
}
func (m *CircuitBreakers) XXX_DiscardUnknown() {
	xxx_messageInfo_CircuitBreakers.DiscardUnknown(m)
}

var xxx_messageInfo_CircuitBreakers proto.InternalMessageInfo

func (m *CircuitBreakers) GetThresholds() []*CircuitBreakers_Thresholds {
	if m != nil {
		return m.Thresholds
	}
	return nil
}

// A Thresholds defines CircuitBreaker settings for a
// :ref:`RoutingPriority<envoy_api_enum_config.core.v3.RoutingPriority>`.
// [#next-free-field: 9]
type CircuitBreakers_Thresholds struct {
	// The :ref:`RoutingPriority<envoy_api_enum_config.core.v3.RoutingPriority>`
	// the specified CircuitBreaker settings apply to.
	Priority v3.RoutingPriority `protobuf:"varint,1,opt,name=priority,proto3,enum=envoy.config.core.v3.RoutingPriority" json:"priority,omitempty"`
	// The maximum number of connections that Envoy will make to the upstream
	// cluster. If not specified, the default is 1024.
	MaxConnections *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=max_connections,json=maxConnections,proto3" json:"max_connections,omitempty"`
	// The maximum number of pending requests that Envoy will allow to the
	// upstream cluster. If not specified, the default is 1024.
	MaxPendingRequests *wrappers.UInt32Value `protobuf:"bytes,3,opt,name=max_pending_requests,json=maxPendingRequests,proto3" json:"max_pending_requests,omitempty"`
	// The maximum number of parallel requests that Envoy will make to the
	// upstream cluster. If not specified, the default is 1024.
	MaxRequests *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=max_requests,json=maxRequests,proto3" json:"max_requests,omitempty"`
	// The maximum number of parallel retries that Envoy will allow to the
	// upstream cluster. If not specified, the default is 3.
	MaxRetries *wrappers.UInt32Value `protobuf:"bytes,5,opt,name=max_retries,json=maxRetries,proto3" json:"max_retries,omitempty"`
	// Specifies a limit on concurrent retries in relation to the number of active requests. This
	// parameter is optional.
	//
	// .. note::
	//
	//    If this field is set, the retry budget will override any configured retry circuit
	//    breaker.
	RetryBudget *CircuitBreakers_Thresholds_RetryBudget `protobuf:"bytes,8,opt,name=retry_budget,json=retryBudget,proto3" json:"retry_budget,omitempty"`
	// If track_remaining is true, then stats will be published that expose
	// the number of resources remaining until the circuit breakers open. If
	// not specified, the default is false.
	//
	// .. note::
	//
	//    If a retry budget is used in lieu of the max_retries circuit breaker,
	//    the remaining retry resources remaining will not be tracked.
	TrackRemaining bool `protobuf:"varint,6,opt,name=track_remaining,json=trackRemaining,proto3" json:"track_remaining,omitempty"`
	// The maximum number of connection pools per cluster that Envoy will concurrently support at
	// once. If not specified, the default is unlimited. Set this for clusters which create a
	// large number of connection pools. See
	// :ref:`Circuit Breaking <arch_overview_circuit_break_cluster_maximum_connection_pools>` for
	// more details.
	MaxConnectionPools   *wrappers.UInt32Value `protobuf:"bytes,7,opt,name=max_connection_pools,json=maxConnectionPools,proto3" json:"max_connection_pools,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CircuitBreakers_Thresholds) Reset()         { *m = CircuitBreakers_Thresholds{} }
func (m *CircuitBreakers_Thresholds) String() string { return proto.CompactTextString(m) }
func (*CircuitBreakers_Thresholds) ProtoMessage()    {}
func (*CircuitBreakers_Thresholds) Descriptor() ([]byte, []int) {
	return fileDescriptor_131e6721a0838844, []int{0, 0}
}

func (m *CircuitBreakers_Thresholds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CircuitBreakers_Thresholds.Unmarshal(m, b)
}
func (m *CircuitBreakers_Thresholds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CircuitBreakers_Thresholds.Marshal(b, m, deterministic)
}
func (m *CircuitBreakers_Thresholds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CircuitBreakers_Thresholds.Merge(m, src)
}
func (m *CircuitBreakers_Thresholds) XXX_Size() int {
	return xxx_messageInfo_CircuitBreakers_Thresholds.Size(m)
}
func (m *CircuitBreakers_Thresholds) XXX_DiscardUnknown() {
	xxx_messageInfo_CircuitBreakers_Thresholds.DiscardUnknown(m)
}

var xxx_messageInfo_CircuitBreakers_Thresholds proto.InternalMessageInfo

func (m *CircuitBreakers_Thresholds) GetPriority() v3.RoutingPriority {
	if m != nil {
		return m.Priority
	}
	return v3.RoutingPriority_DEFAULT
}

func (m *CircuitBreakers_Thresholds) GetMaxConnections() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxConnections
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxPendingRequests() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxPendingRequests
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxRequests() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxRequests
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetMaxRetries() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxRetries
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetRetryBudget() *CircuitBreakers_Thresholds_RetryBudget {
	if m != nil {
		return m.RetryBudget
	}
	return nil
}

func (m *CircuitBreakers_Thresholds) GetTrackRemaining() bool {
	if m != nil {
		return m.TrackRemaining
	}
	return false
}

func (m *CircuitBreakers_Thresholds) GetMaxConnectionPools() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxConnectionPools
	}
	return nil
}

type CircuitBreakers_Thresholds_RetryBudget struct {
	// Specifies the limit on concurrent retries as a percentage of the sum of active requests and
	// active pending requests. For example, if there are 100 active requests and the
	// budget_percent is set to 25, there may be 25 active retries.
	//
	// This parameter is optional. Defaults to 20%.
	BudgetPercent *v31.Percent `protobuf:"bytes,1,opt,name=budget_percent,json=budgetPercent,proto3" json:"budget_percent,omitempty"`
	// Specifies the minimum retry concurrency allowed for the retry budget. The limit on the
	// number of active retries may never go below this number.
	//
	// This parameter is optional. Defaults to 3.
	MinRetryConcurrency  *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=min_retry_concurrency,json=minRetryConcurrency,proto3" json:"min_retry_concurrency,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CircuitBreakers_Thresholds_RetryBudget) Reset() {
	*m = CircuitBreakers_Thresholds_RetryBudget{}
}
func (m *CircuitBreakers_Thresholds_RetryBudget) String() string { return proto.CompactTextString(m) }
func (*CircuitBreakers_Thresholds_RetryBudget) ProtoMessage()    {}
func (*CircuitBreakers_Thresholds_RetryBudget) Descriptor() ([]byte, []int) {
	return fileDescriptor_131e6721a0838844, []int{0, 0, 0}
}

func (m *CircuitBreakers_Thresholds_RetryBudget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CircuitBreakers_Thresholds_RetryBudget.Unmarshal(m, b)
}
func (m *CircuitBreakers_Thresholds_RetryBudget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CircuitBreakers_Thresholds_RetryBudget.Marshal(b, m, deterministic)
}
func (m *CircuitBreakers_Thresholds_RetryBudget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CircuitBreakers_Thresholds_RetryBudget.Merge(m, src)
}
func (m *CircuitBreakers_Thresholds_RetryBudget) XXX_Size() int {
	return xxx_messageInfo_CircuitBreakers_Thresholds_RetryBudget.Size(m)
}
func (m *CircuitBreakers_Thresholds_RetryBudget) XXX_DiscardUnknown() {
	xxx_messageInfo_CircuitBreakers_Thresholds_RetryBudget.DiscardUnknown(m)
}

var xxx_messageInfo_CircuitBreakers_Thresholds_RetryBudget proto.InternalMessageInfo

func (m *CircuitBreakers_Thresholds_RetryBudget) GetBudgetPercent() *v31.Percent {
	if m != nil {
		return m.BudgetPercent
	}
	return nil
}

func (m *CircuitBreakers_Thresholds_RetryBudget) GetMinRetryConcurrency() *wrappers.UInt32Value {
	if m != nil {
		return m.MinRetryConcurrency
	}
	return nil
}

func init() {
	proto.RegisterType((*CircuitBreakers)(nil), "envoy.config.cluster.v3.CircuitBreakers")
	proto.RegisterType((*CircuitBreakers_Thresholds)(nil), "envoy.config.cluster.v3.CircuitBreakers.Thresholds")
	proto.RegisterType((*CircuitBreakers_Thresholds_RetryBudget)(nil), "envoy.config.cluster.v3.CircuitBreakers.Thresholds.RetryBudget")
}

func init() {
	proto.RegisterFile("envoy/config/cluster/v3/circuit_breaker.proto", fileDescriptor_131e6721a0838844)
}

var fileDescriptor_131e6721a0838844 = []byte{
	// 610 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4f, 0x6b, 0xd4, 0x4e,
	0x1c, 0xc6, 0xc9, 0xf6, 0xd7, 0xfd, 0x2d, 0x93, 0xba, 0x2d, 0xa9, 0xda, 0xb0, 0x6a, 0x5d, 0xc5,
	0xea, 0x82, 0x38, 0xc1, 0x5d, 0xf0, 0x50, 0x29, 0x85, 0x2c, 0x0a, 0x22, 0x48, 0x88, 0x7f, 0xae,
	0x61, 0x36, 0x3b, 0x4d, 0x87, 0x26, 0x33, 0x71, 0x66, 0x12, 0x37, 0xb7, 0xe2, 0xc9, 0xd7, 0xe0,
	0x4b, 0xf1, 0x2e, 0x78, 0xf5, 0x2d, 0x78, 0xf4, 0x25, 0x78, 0x92, 0x99, 0xc9, 0x6e, 0xda, 0x4a,
	0x21, 0x7a, 0xcb, 0xe4, 0x79, 0x9e, 0x4f, 0xf2, 0x7d, 0x66, 0x06, 0x3c, 0xc2, 0xb4, 0x64, 0x95,
	0x17, 0x33, 0x7a, 0x44, 0x12, 0x2f, 0x4e, 0x0b, 0x21, 0x31, 0xf7, 0xca, 0x89, 0x17, 0x13, 0x1e,
	0x17, 0x44, 0x46, 0x33, 0x8e, 0xd1, 0x09, 0xe6, 0x30, 0xe7, 0x4c, 0x32, 0x67, 0x47, 0xdb, 0xa1,
	0xb1, 0xc3, 0xda, 0x0e, 0xcb, 0xc9, 0xe0, 0xf6, 0x79, 0x0e, 0xe3, 0x58, 0x41, 0x66, 0x48, 0x60,
	0x93, 0x1c, 0xdc, 0x30, 0x06, 0x59, 0xe5, 0x5a, 0xc9, 0x31, 0x8f, 0x31, 0x95, 0xb5, 0xb8, 0x9b,
	0x30, 0x96, 0xa4, 0xd8, 0xd3, 0xab, 0x59, 0x71, 0xe4, 0x7d, 0xe0, 0x28, 0xcf, 0x31, 0x17, 0xb5,
	0x7e, 0xab, 0x98, 0xe7, 0xc8, 0x43, 0x94, 0x32, 0x89, 0x24, 0x61, 0x54, 0x78, 0x42, 0x22, 0x59,
	0x2c, 0xe5, 0x3b, 0x7f, 0xc8, 0x25, 0xe6, 0x82, 0x30, 0x4a, 0x68, 0x52, 0x5b, 0x76, 0x4a, 0x94,
	0x92, 0x39, 0x92, 0xd8, 0x5b, 0x3e, 0x18, 0xe1, 0xee, 0x69, 0x0f, 0x6c, 0x4e, 0xcd, 0xac, 0xbe,
	0x19, 0x55, 0x38, 0xaf, 0x01, 0x90, 0xc7, 0x1c, 0x8b, 0x63, 0x96, 0xce, 0x85, 0x6b, 0x0d, 0xd7,
	0x46, 0xf6, 0x78, 0x02, 0x2f, 0x19, 0x1d, 0x5e, 0x48, 0xc3, 0x37, 0xab, 0x68, 0x78, 0x06, 0x33,
	0xf8, 0xd9, 0x05, 0xa0, 0x91, 0x9c, 0x97, 0xa0, 0x97, 0x73, 0xc2, 0x38, 0x91, 0x95, 0x6b, 0x0d,
	0xad, 0x51, 0x7f, 0xbc, 0x77, 0xe1, 0x0b, 0x8c, 0x63, 0x85, 0x0f, 0x59, 0x21, 0x09, 0x4d, 0x82,
	0xda, 0xec, 0xf7, 0x7e, 0xf9, 0xeb, 0x1f, 0xad, 0xce, 0x96, 0x15, 0xae, 0x00, 0xce, 0x33, 0xb0,
	0x99, 0xa1, 0x45, 0x14, 0x33, 0x4a, 0x71, 0xac, 0x1b, 0x70, 0x3b, 0x43, 0x6b, 0x64, 0x8f, 0x6f,
	0x42, 0xd3, 0x2c, 0x5c, 0x36, 0x0b, 0xdf, 0xbe, 0xa0, 0x72, 0x32, 0x7e, 0x87, 0xd2, 0x02, 0x87,
	0xfd, 0x0c, 0x2d, 0xa6, 0x4d, 0xc6, 0x79, 0x05, 0xae, 0x2a, 0x4c, 0x8e, 0xe9, 0x9c, 0xd0, 0x24,
	0xe2, 0xf8, 0x7d, 0x81, 0x85, 0x14, 0xee, 0x5a, 0x0b, 0x96, 0x93, 0xa1, 0x45, 0x60, 0x82, 0x61,
	0x9d, 0x73, 0x0e, 0xc1, 0x86, 0xe2, 0xad, 0x38, 0xff, 0xb5, 0xe0, 0xd8, 0x19, 0x5a, 0xac, 0x00,
	0x07, 0xc0, 0x36, 0x00, 0xc9, 0x09, 0x16, 0xee, 0x7a, 0x8b, 0x3c, 0xd0, 0x79, 0xed, 0x77, 0x66,
	0x60, 0x43, 0x45, 0xab, 0x68, 0x56, 0xcc, 0x13, 0x2c, 0xdd, 0x9e, 0xce, 0x1f, 0xfe, 0xc3, 0x4e,
	0x42, 0x85, 0xac, 0x7c, 0x8d, 0x09, 0x6d, 0xde, 0x2c, 0x9c, 0x07, 0x60, 0x53, 0x72, 0x14, 0x9f,
	0x44, 0x1c, 0x67, 0x88, 0xa8, 0x13, 0xe7, 0x76, 0x87, 0xd6, 0xa8, 0x17, 0xf6, 0xf5, 0xeb, 0x70,
	0xf9, 0x76, 0x59, 0x6e, 0xb3, 0x47, 0x51, 0xce, 0x58, 0x2a, 0xdc, 0xff, 0x5b, 0x96, 0xdb, 0x6c,
	0x54, 0xa0, 0x72, 0x83, 0x1f, 0x16, 0xb0, 0xcf, 0xfc, 0x95, 0x73, 0x00, 0xfa, 0x66, 0xcc, 0xa8,
	0xbe, 0x5b, 0xfa, 0x58, 0xd9, 0xe3, 0xeb, 0xf5, 0xb8, 0xea, 0xe6, 0xa9, 0x21, 0x03, 0xa3, 0x86,
	0x57, 0x8c, 0xbb, 0x5e, 0x3a, 0x01, 0xb8, 0x96, 0x11, 0x1a, 0x99, 0xbe, 0x62, 0x46, 0xe3, 0x82,
	0x73, 0x4c, 0xe3, 0xaa, 0xd5, 0x41, 0xda, 0xce, 0x08, 0xd5, 0xff, 0x32, 0x6d, 0x82, 0xfb, 0xfe,
	0xe7, 0xaf, 0x9f, 0x76, 0x0f, 0xc0, 0x53, 0xf3, 0x79, 0x94, 0x13, 0x58, 0x8e, 0x57, 0x6d, 0xb7,
	0xab, 0x7a, 0xff, 0x89, 0x62, 0x3c, 0x06, 0xde, 0x5f, 0x32, 0xf6, 0x1f, 0xaa, 0xdc, 0x7d, 0x70,
	0xaf, 0x4d, 0xce, 0x7f, 0xfe, 0xe5, 0xf4, 0xdb, 0xf7, 0x6e, 0x67, 0xab, 0x03, 0xf6, 0x08, 0x33,
	0x6d, 0xe5, 0x9c, 0x2d, 0xaa, 0xcb, 0xce, 0x89, 0xbf, 0x7d, 0x9e, 0x10, 0xa8, 0x4a, 0x02, 0x6b,
	0xd6, 0xd5, 0xdd, 0x4c, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x00, 0x84, 0xf0, 0x54, 0x05,
	0x00, 0x00,
}
