package grpc

import (
	"context"
	"ride-sharing/services/trip-service/internal/domain"
	pb "ride-sharing/shared/proto/trip"
	"ride-sharing/shared/types"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type gRPCHandler struct {
	pb.UnimplementedTripServiceServer
	service domain.TripService
}

func NewGRPCHandler(server *grpc.Server, service domain.TripService) *gRPCHandler {
	handler := &gRPCHandler{
		service: service,
	}
	pb.RegisterTripServiceServer(server, handler)
	return handler
}

func (h *gRPCHandler) PreviewTrip(ctx context.Context, req *pb.PreviewTripRequest) (*pb.PreviewTripResponse, error) {
	pickupCoord := &types.Coordinate{
		Latitude:  req.Pickup.Latitude,
		Longitude: req.Pickup.Longitude,
	}

	destinationCoord := &types.Coordinate{
		Latitude:  req.Destination.Latitude,
		Longitude: req.Destination.Longitude,
	}

	route, err := h.service.GetRoute(ctx, pickupCoord, destinationCoord)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get route: %v", err)
	}

	return &pb.PreviewTripResponse{
		Route:     route.ToProto(),
		RideFares: []*pb.RideFare{},
	}, nil
}
