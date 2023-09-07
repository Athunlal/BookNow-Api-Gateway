// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BookingManagementClient is the client API for BookingManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookingManagementClient interface {
	SearchTrain(ctx context.Context, in *SearchTrainRequest, opts ...grpc.CallOption) (*SearchTrainResponse, error)
	SearchCompartment(ctx context.Context, in *SearchCompartmentRequest, opts ...grpc.CallOption) (*SearchCompartmentResponse, error)
	Checkout(ctx context.Context, in *CheckoutRequest, opts ...grpc.CallOption) (*CheckoutResponse, error)
	Payment(ctx context.Context, in *PaymentRequest, opts ...grpc.CallOption) (*PaymentResponse, error)
	CreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*CreateWalletResponse, error)
	UpdateAmount(ctx context.Context, in *UpdateAmountRequest, opts ...grpc.CallOption) (*UpdateAmountResponse, error)
	ViewTicket(ctx context.Context, in *ViewTicketRequest, opts ...grpc.CallOption) (*ViewTicketResponse, error)
	CancelTicket(ctx context.Context, in *CancelTicketRequest, opts ...grpc.CallOption) (*CancelTicketResponse, error)
	BookingHistory(ctx context.Context, in *BookingHistroyRequest, opts ...grpc.CallOption) (*BookingHistoryResponse, error)
}

type bookingManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewBookingManagementClient(cc grpc.ClientConnInterface) BookingManagementClient {
	return &bookingManagementClient{cc}
}

func (c *bookingManagementClient) SearchTrain(ctx context.Context, in *SearchTrainRequest, opts ...grpc.CallOption) (*SearchTrainResponse, error) {
	out := new(SearchTrainResponse)
	err := c.cc.Invoke(ctx, "/Booking.BookingManagement/SearchTrain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingManagementClient) SearchCompartment(ctx context.Context, in *SearchCompartmentRequest, opts ...grpc.CallOption) (*SearchCompartmentResponse, error) {
	out := new(SearchCompartmentResponse)
	err := c.cc.Invoke(ctx, "/Booking.BookingManagement/SearchCompartment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingManagementClient) Checkout(ctx context.Context, in *CheckoutRequest, opts ...grpc.CallOption) (*CheckoutResponse, error) {
	out := new(CheckoutResponse)
	err := c.cc.Invoke(ctx, "/Booking.BookingManagement/Checkout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingManagementClient) Payment(ctx context.Context, in *PaymentRequest, opts ...grpc.CallOption) (*PaymentResponse, error) {
	out := new(PaymentResponse)
	err := c.cc.Invoke(ctx, "/Booking.BookingManagement/Payment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingManagementClient) CreateWallet(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*CreateWalletResponse, error) {
	out := new(CreateWalletResponse)
	err := c.cc.Invoke(ctx, "/Booking.BookingManagement/CreateWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingManagementClient) UpdateAmount(ctx context.Context, in *UpdateAmountRequest, opts ...grpc.CallOption) (*UpdateAmountResponse, error) {
	out := new(UpdateAmountResponse)
	err := c.cc.Invoke(ctx, "/Booking.BookingManagement/UpdateAmount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingManagementClient) ViewTicket(ctx context.Context, in *ViewTicketRequest, opts ...grpc.CallOption) (*ViewTicketResponse, error) {
	out := new(ViewTicketResponse)
	err := c.cc.Invoke(ctx, "/Booking.BookingManagement/ViewTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingManagementClient) CancelTicket(ctx context.Context, in *CancelTicketRequest, opts ...grpc.CallOption) (*CancelTicketResponse, error) {
	out := new(CancelTicketResponse)
	err := c.cc.Invoke(ctx, "/Booking.BookingManagement/CancelTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingManagementClient) BookingHistory(ctx context.Context, in *BookingHistroyRequest, opts ...grpc.CallOption) (*BookingHistoryResponse, error) {
	out := new(BookingHistoryResponse)
	err := c.cc.Invoke(ctx, "/Booking.BookingManagement/BookingHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookingManagementServer is the server API for BookingManagement service.
// All implementations must embed UnimplementedBookingManagementServer
// for forward compatibility
type BookingManagementServer interface {
	SearchTrain(context.Context, *SearchTrainRequest) (*SearchTrainResponse, error)
	SearchCompartment(context.Context, *SearchCompartmentRequest) (*SearchCompartmentResponse, error)
	Checkout(context.Context, *CheckoutRequest) (*CheckoutResponse, error)
	Payment(context.Context, *PaymentRequest) (*PaymentResponse, error)
	CreateWallet(context.Context, *CreateWalletRequest) (*CreateWalletResponse, error)
	UpdateAmount(context.Context, *UpdateAmountRequest) (*UpdateAmountResponse, error)
	ViewTicket(context.Context, *ViewTicketRequest) (*ViewTicketResponse, error)
	CancelTicket(context.Context, *CancelTicketRequest) (*CancelTicketResponse, error)
	BookingHistory(context.Context, *BookingHistroyRequest) (*BookingHistoryResponse, error)
	mustEmbedUnimplementedBookingManagementServer()
}

// UnimplementedBookingManagementServer must be embedded to have forward compatible implementations.
type UnimplementedBookingManagementServer struct {
}

func (UnimplementedBookingManagementServer) SearchTrain(context.Context, *SearchTrainRequest) (*SearchTrainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchTrain not implemented")
}
func (UnimplementedBookingManagementServer) SearchCompartment(context.Context, *SearchCompartmentRequest) (*SearchCompartmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchCompartment not implemented")
}
func (UnimplementedBookingManagementServer) Checkout(context.Context, *CheckoutRequest) (*CheckoutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Checkout not implemented")
}
func (UnimplementedBookingManagementServer) Payment(context.Context, *PaymentRequest) (*PaymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Payment not implemented")
}
func (UnimplementedBookingManagementServer) CreateWallet(context.Context, *CreateWalletRequest) (*CreateWalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWallet not implemented")
}
func (UnimplementedBookingManagementServer) UpdateAmount(context.Context, *UpdateAmountRequest) (*UpdateAmountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAmount not implemented")
}
func (UnimplementedBookingManagementServer) ViewTicket(context.Context, *ViewTicketRequest) (*ViewTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ViewTicket not implemented")
}
func (UnimplementedBookingManagementServer) CancelTicket(context.Context, *CancelTicketRequest) (*CancelTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelTicket not implemented")
}
func (UnimplementedBookingManagementServer) BookingHistory(context.Context, *BookingHistroyRequest) (*BookingHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BookingHistory not implemented")
}
func (UnimplementedBookingManagementServer) mustEmbedUnimplementedBookingManagementServer() {}

// UnsafeBookingManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookingManagementServer will
// result in compilation errors.
type UnsafeBookingManagementServer interface {
	mustEmbedUnimplementedBookingManagementServer()
}

func RegisterBookingManagementServer(s grpc.ServiceRegistrar, srv BookingManagementServer) {
	s.RegisterService(&BookingManagement_ServiceDesc, srv)
}

func _BookingManagement_SearchTrain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchTrainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingManagementServer).SearchTrain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Booking.BookingManagement/SearchTrain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingManagementServer).SearchTrain(ctx, req.(*SearchTrainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingManagement_SearchCompartment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchCompartmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingManagementServer).SearchCompartment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Booking.BookingManagement/SearchCompartment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingManagementServer).SearchCompartment(ctx, req.(*SearchCompartmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingManagement_Checkout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingManagementServer).Checkout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Booking.BookingManagement/Checkout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingManagementServer).Checkout(ctx, req.(*CheckoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingManagement_Payment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingManagementServer).Payment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Booking.BookingManagement/Payment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingManagementServer).Payment(ctx, req.(*PaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingManagement_CreateWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingManagementServer).CreateWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Booking.BookingManagement/CreateWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingManagementServer).CreateWallet(ctx, req.(*CreateWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingManagement_UpdateAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAmountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingManagementServer).UpdateAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Booking.BookingManagement/UpdateAmount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingManagementServer).UpdateAmount(ctx, req.(*UpdateAmountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingManagement_ViewTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ViewTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingManagementServer).ViewTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Booking.BookingManagement/ViewTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingManagementServer).ViewTicket(ctx, req.(*ViewTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingManagement_CancelTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingManagementServer).CancelTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Booking.BookingManagement/CancelTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingManagementServer).CancelTicket(ctx, req.(*CancelTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingManagement_BookingHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookingHistroyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingManagementServer).BookingHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Booking.BookingManagement/BookingHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingManagementServer).BookingHistory(ctx, req.(*BookingHistroyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookingManagement_ServiceDesc is the grpc.ServiceDesc for BookingManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookingManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Booking.BookingManagement",
	HandlerType: (*BookingManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchTrain",
			Handler:    _BookingManagement_SearchTrain_Handler,
		},
		{
			MethodName: "SearchCompartment",
			Handler:    _BookingManagement_SearchCompartment_Handler,
		},
		{
			MethodName: "Checkout",
			Handler:    _BookingManagement_Checkout_Handler,
		},
		{
			MethodName: "Payment",
			Handler:    _BookingManagement_Payment_Handler,
		},
		{
			MethodName: "CreateWallet",
			Handler:    _BookingManagement_CreateWallet_Handler,
		},
		{
			MethodName: "UpdateAmount",
			Handler:    _BookingManagement_UpdateAmount_Handler,
		},
		{
			MethodName: "ViewTicket",
			Handler:    _BookingManagement_ViewTicket_Handler,
		},
		{
			MethodName: "CancelTicket",
			Handler:    _BookingManagement_CancelTicket_Handler,
		},
		{
			MethodName: "BookingHistory",
			Handler:    _BookingManagement_BookingHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/booking/pb/booking.proto",
}
