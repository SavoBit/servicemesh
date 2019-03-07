// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mixer/template/listentry/template_handler_service.proto

/*
	Package listentry is a generated protocol buffer package.

	The `listentry` template is designed to let you perform list check operations
	with the [list](https://istio.io/docs/reference/config/policy-and-telemetry/adapters/list/) adapter.

	Example config:

	```yaml
	apiVersion: "config.istio.io/v1alpha2"
	kind: listentry
	metadata:
	  name: appversion
	  namespace: istio-system
	spec:
	  value: source.labels["version"]
	```

	The `listentry` template is used to verify the presence/absence of a string
	within a list.

	When writing the configuration, the value for the fields associated with this template can either be a
	literal or an [expression](https://istio.io/docs/reference//config/policy-and-telemetry/expression-language/). Please note that if the datatype of a field is not istio.policy.v1beta1.Value,
	then the expression's [inferred type](https://istio.io/docs/reference//config/policy-and-telemetry/expression-language/#type-checking) must match the datatype of the field.

	It is generated from these files:
		mixer/template/listentry/template_handler_service.proto

	It has these top-level messages:
		HandleListEntryRequest
		InstanceMsg
		Type
		InstanceParam
*/
package listentry

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "istio.io/api/mixer/adapter/model/v1beta1"
import google_protobuf1 "github.com/gogo/protobuf/types"
import istio_mixer_adapter_model_v1beta11 "istio.io/api/mixer/adapter/model/v1beta1"
import istio_policy_v1beta1 "istio.io/api/policy/v1beta1"
import istio_policy_v1beta11 "istio.io/api/policy/v1beta1"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Request message for HandleListEntry method.
type HandleListEntryRequest struct {
	// 'listentry' instance.
	Instance *InstanceMsg `protobuf:"bytes,1,opt,name=instance" json:"instance,omitempty"`
	// Adapter specific handler configuration.
	//
	// Note: Backends can also implement [InfrastructureBackend][https://istio.io/docs/reference/config/mixer/istio.mixer.adapter.model.v1beta1.html#InfrastructureBackend]
	// service and therefore opt to receive handler configuration during session creation through [InfrastructureBackend.CreateSession][TODO: Link to this fragment]
	// call. In that case, adapter_config will have type_url as 'google.protobuf.Any.type_url' and would contain string
	// value of session_id (returned from InfrastructureBackend.CreateSession).
	AdapterConfig *google_protobuf1.Any `protobuf:"bytes,2,opt,name=adapter_config,json=adapterConfig" json:"adapter_config,omitempty"`
	// Id to dedupe identical requests from Mixer.
	DedupId string `protobuf:"bytes,3,opt,name=dedup_id,json=dedupId,proto3" json:"dedup_id,omitempty"`
}

func (m *HandleListEntryRequest) Reset()      { *m = HandleListEntryRequest{} }
func (*HandleListEntryRequest) ProtoMessage() {}
func (*HandleListEntryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorTemplateHandlerService, []int{0}
}

// Contains instance payload for 'listentry' template. This is passed to infrastructure backends during request-time
// through HandleListEntryService.HandleListEntry.
type InstanceMsg struct {
	// Name of the instance as specified in configuration.
	Name string `protobuf:"bytes,72295727,opt,name=name,proto3" json:"name,omitempty"`
	// Specifies the entry to verify in the list. This value can either be a string or an IP address.
	Value *istio_policy_v1beta11.Value `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *InstanceMsg) Reset()      { *m = InstanceMsg{} }
func (*InstanceMsg) ProtoMessage() {}
func (*InstanceMsg) Descriptor() ([]byte, []int) {
	return fileDescriptorTemplateHandlerService, []int{1}
}

// Contains inferred type information about specific instance of 'listentry' template. This is passed to
// infrastructure backends during configuration-time through [InfrastructureBackend.CreateSession][TODO: Link to this fragment].
type Type struct {
	// Specifies the entry to verify in the list. This value can either be a string or an IP address.
	Value istio_policy_v1beta1.ValueType `protobuf:"varint,1,opt,name=value,proto3,enum=istio.policy.v1beta1.ValueType" json:"value,omitempty"`
}

func (m *Type) Reset()                    { *m = Type{} }
func (*Type) ProtoMessage()               {}
func (*Type) Descriptor() ([]byte, []int) { return fileDescriptorTemplateHandlerService, []int{2} }

// Represents instance configuration schema for 'listentry' template.
type InstanceParam struct {
	// Specifies the entry to verify in the list. This value can either be a string or an IP address.
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *InstanceParam) Reset()      { *m = InstanceParam{} }
func (*InstanceParam) ProtoMessage() {}
func (*InstanceParam) Descriptor() ([]byte, []int) {
	return fileDescriptorTemplateHandlerService, []int{3}
}

func init() {
	proto.RegisterType((*HandleListEntryRequest)(nil), "listentry.HandleListEntryRequest")
	proto.RegisterType((*InstanceMsg)(nil), "listentry.InstanceMsg")
	proto.RegisterType((*Type)(nil), "listentry.Type")
	proto.RegisterType((*InstanceParam)(nil), "listentry.InstanceParam")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for HandleListEntryService service

type HandleListEntryServiceClient interface {
	// HandleListEntry is called by Mixer at request-time to deliver 'listentry' instances to the backend.
	HandleListEntry(ctx context.Context, in *HandleListEntryRequest, opts ...grpc.CallOption) (*istio_mixer_adapter_model_v1beta11.CheckResult, error)
}

type handleListEntryServiceClient struct {
	cc *grpc.ClientConn
}

func NewHandleListEntryServiceClient(cc *grpc.ClientConn) HandleListEntryServiceClient {
	return &handleListEntryServiceClient{cc}
}

func (c *handleListEntryServiceClient) HandleListEntry(ctx context.Context, in *HandleListEntryRequest, opts ...grpc.CallOption) (*istio_mixer_adapter_model_v1beta11.CheckResult, error) {
	out := new(istio_mixer_adapter_model_v1beta11.CheckResult)
	err := grpc.Invoke(ctx, "/listentry.HandleListEntryService/HandleListEntry", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HandleListEntryService service

type HandleListEntryServiceServer interface {
	// HandleListEntry is called by Mixer at request-time to deliver 'listentry' instances to the backend.
	HandleListEntry(context.Context, *HandleListEntryRequest) (*istio_mixer_adapter_model_v1beta11.CheckResult, error)
}

func RegisterHandleListEntryServiceServer(s *grpc.Server, srv HandleListEntryServiceServer) {
	s.RegisterService(&_HandleListEntryService_serviceDesc, srv)
}

func _HandleListEntryService_HandleListEntry_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleListEntryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HandleListEntryServiceServer).HandleListEntry(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/listentry.HandleListEntryService/HandleListEntry",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HandleListEntryServiceServer).HandleListEntry(ctx, req.(*HandleListEntryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HandleListEntryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "listentry.HandleListEntryService",
	HandlerType: (*HandleListEntryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleListEntry",
			Handler:    _HandleListEntryService_HandleListEntry_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mixer/template/listentry/template_handler_service.proto",
}

func (m *HandleListEntryRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HandleListEntryRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Instance != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(m.Instance.Size()))
		n1, err := m.Instance.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.AdapterConfig != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(m.AdapterConfig.Size()))
		n2, err := m.AdapterConfig.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if len(m.DedupId) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(len(m.DedupId)))
		i += copy(dAtA[i:], m.DedupId)
	}
	return i, nil
}

func (m *InstanceMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InstanceMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Value != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(m.Value.Size()))
		n3, err := m.Value.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0xfa
		i++
		dAtA[i] = 0xd2
		i++
		dAtA[i] = 0xe4
		i++
		dAtA[i] = 0x93
		i++
		dAtA[i] = 0x2
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	return i, nil
}

func (m *Type) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Type) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Value != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(m.Value))
	}
	return i, nil
}

func (m *InstanceParam) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InstanceParam) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTemplateHandlerService(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	return i, nil
}

func encodeVarintTemplateHandlerService(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *HandleListEntryRequest) Size() (n int) {
	var l int
	_ = l
	if m.Instance != nil {
		l = m.Instance.Size()
		n += 1 + l + sovTemplateHandlerService(uint64(l))
	}
	if m.AdapterConfig != nil {
		l = m.AdapterConfig.Size()
		n += 1 + l + sovTemplateHandlerService(uint64(l))
	}
	l = len(m.DedupId)
	if l > 0 {
		n += 1 + l + sovTemplateHandlerService(uint64(l))
	}
	return n
}

func (m *InstanceMsg) Size() (n int) {
	var l int
	_ = l
	if m.Value != nil {
		l = m.Value.Size()
		n += 1 + l + sovTemplateHandlerService(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 5 + l + sovTemplateHandlerService(uint64(l))
	}
	return n
}

func (m *Type) Size() (n int) {
	var l int
	_ = l
	if m.Value != 0 {
		n += 1 + sovTemplateHandlerService(uint64(m.Value))
	}
	return n
}

func (m *InstanceParam) Size() (n int) {
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovTemplateHandlerService(uint64(l))
	}
	return n
}

func sovTemplateHandlerService(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTemplateHandlerService(x uint64) (n int) {
	return sovTemplateHandlerService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *HandleListEntryRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HandleListEntryRequest{`,
		`Instance:` + strings.Replace(fmt.Sprintf("%v", this.Instance), "InstanceMsg", "InstanceMsg", 1) + `,`,
		`AdapterConfig:` + strings.Replace(fmt.Sprintf("%v", this.AdapterConfig), "Any", "google_protobuf1.Any", 1) + `,`,
		`DedupId:` + fmt.Sprintf("%v", this.DedupId) + `,`,
		`}`,
	}, "")
	return s
}
func (this *InstanceMsg) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&InstanceMsg{`,
		`Value:` + strings.Replace(fmt.Sprintf("%v", this.Value), "Value", "istio_policy_v1beta11.Value", 1) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Type) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Type{`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`}`,
	}, "")
	return s
}
func (this *InstanceParam) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&InstanceParam{`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTemplateHandlerService(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *HandleListEntryRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTemplateHandlerService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HandleListEntryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HandleListEntryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Instance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Instance == nil {
				m.Instance = &InstanceMsg{}
			}
			if err := m.Instance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdapterConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AdapterConfig == nil {
				m.AdapterConfig = &google_protobuf1.Any{}
			}
			if err := m.AdapterConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DedupId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DedupId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTemplateHandlerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *InstanceMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTemplateHandlerService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InstanceMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InstanceMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Value == nil {
				m.Value = &istio_policy_v1beta11.Value{}
			}
			if err := m.Value.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 72295727:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTemplateHandlerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Type) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTemplateHandlerService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Type: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Type: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= (istio_policy_v1beta1.ValueType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTemplateHandlerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *InstanceParam) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTemplateHandlerService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InstanceParam: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InstanceParam: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemplateHandlerService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTemplateHandlerService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTemplateHandlerService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTemplateHandlerService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTemplateHandlerService
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTemplateHandlerService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTemplateHandlerService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTemplateHandlerService
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTemplateHandlerService
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTemplateHandlerService(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTemplateHandlerService = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTemplateHandlerService   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("mixer/template/listentry/template_handler_service.proto", fileDescriptorTemplateHandlerService)
}

var fileDescriptorTemplateHandlerService = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x93, 0xb5, 0x6a, 0x77, 0x4a, 0x2b, 0x0c, 0x6b, 0xd9, 0xae, 0x30, 0xad, 0x0b, 0x62,
	0x0f, 0x32, 0x61, 0x57, 0xc4, 0x83, 0x78, 0xd0, 0x52, 0xb0, 0xa0, 0x20, 0x51, 0xf4, 0xb8, 0xcc,
	0x26, 0xaf, 0xe9, 0x60, 0x32, 0x13, 0x33, 0x93, 0xa5, 0xb9, 0x88, 0xf8, 0x09, 0x84, 0x7e, 0x05,
	0x0f, 0xde, 0xfc, 0x02, 0x7e, 0x80, 0xe2, 0xa9, 0x78, 0xf2, 0x22, 0xb8, 0xb1, 0x07, 0x8f, 0x3d,
	0x7a, 0x94, 0x4c, 0x66, 0xd3, 0xa5, 0x96, 0xde, 0xf2, 0xe6, 0xfd, 0xf2, 0x9f, 0xf7, 0x7f, 0xf3,
	0x47, 0xf7, 0x13, 0xbe, 0x0f, 0x99, 0xa7, 0x21, 0x49, 0x63, 0xa6, 0xc1, 0x8b, 0xb9, 0xd2, 0x20,
	0x74, 0x56, 0x34, 0x47, 0xa3, 0x3d, 0x26, 0xc2, 0x18, 0xb2, 0x91, 0x82, 0x6c, 0xc2, 0x03, 0xa0,
	0x69, 0x26, 0xb5, 0xc4, 0xed, 0x86, 0xec, 0x75, 0x22, 0x19, 0x49, 0x73, 0xea, 0x55, 0x5f, 0x35,
	0xd0, 0xbb, 0x53, 0x2b, 0xb3, 0x90, 0xa5, 0x1a, 0x32, 0x2f, 0x91, 0x21, 0xc4, 0xde, 0x64, 0x30,
	0x06, 0xcd, 0x06, 0x1e, 0xec, 0x6b, 0x10, 0x8a, 0x4b, 0xa1, 0x2c, 0xbd, 0x16, 0x49, 0x19, 0xc5,
	0xe0, 0x99, 0x6a, 0x9c, 0xef, 0x7a, 0x4c, 0x14, 0xb6, 0x75, 0xfb, 0x22, 0xa1, 0x60, 0x0f, 0x82,
	0x37, 0x16, 0x5c, 0x4f, 0x65, 0xcc, 0x83, 0xa2, 0xe9, 0x4d, 0x58, 0x9c, 0xc3, 0x48, 0x17, 0x29,
	0xcc, 0x2e, 0x39, 0x03, 0x9c, 0xb6, 0xfa, 0x9f, 0x5c, 0xb4, 0xfa, 0xc4, 0x18, 0x7d, 0xca, 0x95,
	0xde, 0xae, 0x7c, 0xf9, 0xf0, 0x36, 0x07, 0xa5, 0xf1, 0x10, 0x2d, 0x72, 0xa1, 0x34, 0x13, 0x01,
	0x74, 0xdd, 0x0d, 0x77, 0x73, 0x69, 0xb8, 0x4a, 0x1b, 0xf3, 0x74, 0xc7, 0xb6, 0x9e, 0xa9, 0xc8,
	0x6f, 0x38, 0xfc, 0x00, 0xad, 0xd8, 0x79, 0x47, 0x81, 0x14, 0xbb, 0x3c, 0xea, 0xb6, 0xcc, 0x9f,
	0x1d, 0x5a, 0xfb, 0xa4, 0x33, 0x9f, 0xf4, 0x91, 0x28, 0xfc, 0x65, 0xcb, 0x6e, 0x19, 0x14, 0xaf,
	0xa1, 0xc5, 0x10, 0xc2, 0x3c, 0x1d, 0xf1, 0xb0, 0x7b, 0x69, 0xc3, 0xdd, 0x6c, 0xfb, 0x57, 0x4d,
	0xbd, 0x13, 0xf6, 0x5f, 0xa3, 0xa5, 0xb9, 0x0b, 0xf1, 0x00, 0x5d, 0x36, 0x26, 0xed, 0x5c, 0x37,
	0x28, 0x57, 0x9a, 0x4b, 0x5a, 0xdb, 0xa4, 0xd6, 0x26, 0x7d, 0x55, 0x21, 0x7e, 0x4d, 0xe2, 0xeb,
	0x68, 0x41, 0xb0, 0x04, 0xba, 0x5f, 0xbe, 0x7d, 0xed, 0x1b, 0x6d, 0x53, 0xf6, 0x1f, 0xa2, 0x85,
	0x97, 0x45, 0x0a, 0xf8, 0xde, 0xbc, 0xe2, 0xca, 0x70, 0xfd, 0x02, 0xc5, 0x8a, 0xb7, 0xaa, 0xfd,
	0x5b, 0x68, 0x79, 0x36, 0xd7, 0x73, 0x96, 0xb1, 0x04, 0x77, 0xe6, 0x75, 0xda, 0x16, 0x1b, 0xbe,
	0xfb, 0x6f, 0xc9, 0x2f, 0xea, 0x50, 0xe1, 0x10, 0x5d, 0x3b, 0xd3, 0xc1, 0x37, 0xe7, 0xb6, 0x7c,
	0xfe, 0xd3, 0xf4, 0xa8, 0x1d, 0xcf, 0x24, 0x84, 0xda, 0x2d, 0x52, 0x93, 0x90, 0x66, 0xd6, 0xad,
	0x2a, 0x21, 0x3e, 0xa8, 0x3c, 0xd6, 0x8f, 0xb7, 0x0f, 0xa7, 0xc4, 0x39, 0x9a, 0x12, 0xe7, 0xc7,
	0x94, 0x38, 0x27, 0x53, 0xe2, 0xbc, 0x2f, 0x89, 0xfb, 0xb9, 0x24, 0xce, 0x61, 0x49, 0xdc, 0xa3,
	0x92, 0xb8, 0xbf, 0x4a, 0xe2, 0xfe, 0x29, 0x89, 0x73, 0x52, 0x12, 0xf7, 0xe3, 0x6f, 0xe2, 0xfc,
	0xfd, 0x7e, 0x7c, 0xd0, 0x72, 0x3e, 0xfc, 0x3c, 0x3e, 0x68, 0x9d, 0x06, 0x7e, 0x7c, 0xc5, 0xbc,
	0xde, 0xdd, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1e, 0xee, 0xf4, 0x64, 0x3d, 0x03, 0x00, 0x00,
}