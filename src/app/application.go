package app

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gitlab.com/chertokdmitry/surfproto/src/controllers/cameras"
	"gitlab.com/chertokdmitry/surfproto/src/controllers/spots"
	"gitlab.com/chertokdmitry/surfproto/src/controllers/subscriptions"
	"gitlab.com/chertokdmitry/surfproto/src/controllers/weather"
	pb "gitlab.com/chertokdmitry/surfproto/src/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

var (
	router = gin.Default()
)

func StartApplication() {
	listener, err := net.Listen("tcp", ":5300")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterApiServer(grpcServer, &server{})
	grpcServer.Serve(listener)
}

type server struct{}

func (s *server) WeatherBySpotId (c context.Context, request *pb.WeatherRequest) (response *pb.WeatherResponse, err error) {

	weather := weather.Get(request.SpotId)

	response = &pb.WeatherResponse{
		SpotId: weather.SpotId,
		Title: weather.Title,
		Hourly: weather.Hourly,
	}

	return response, nil
}

func (s *server) GetSpotsByRegionId (c context.Context, request *pb.SpotsRequest) (response *pb.SpotsResponse, err error) {

	spots, _ := json.Marshal(spots.GetByRegionId(request.Id))

	response = &pb.SpotsResponse{
		Spots: string(spots),
	}

	return response, nil
}

func (s *server) GetCamerasBySpotId (c context.Context, request *pb.CamerasRequest) (response *pb.CamerasResponse, err error) {

	cameras, _ := json.Marshal(cameras.GetBySpot(request.SpotId))

	response = &pb.CamerasResponse{
		Cameras: string(cameras),
	}

	return response, nil
}

func (s *server) CreateSubscription (c context.Context, request *pb.CreateSubscriptionRequest) (response *pb.CreateSubscriptionResponse, err error) {

	subscriptions.Create(request.Sub)

	response = &pb.CreateSubscriptionResponse{
		Result: "success",
	}

	return response, nil
}

func (s *server) GetAllSpots (c context.Context, request *pb.AllSpotsRequest) (response *pb.AllSpotsResponse, err error) {

	spots, _ := json.Marshal(spots.GetAll())

	response = &pb.AllSpotsResponse{
		Spots: string(spots),
	}

	return response, nil
}

func (s *server) GetSubsByChatId (c context.Context, request *pb.SubsRequest) (response *pb.SubsResponse, err error) {

	subs, _ := json.Marshal(subscriptions.GetByChatId(request.ChatId))
	response = &pb.SubsResponse{
		Subs: string(subs),
	}

	return response, nil
}

func (s *server) DeleteSub (c context.Context, request *pb.DeleteSubRequest) (response *pb.DeleteSubResponse, err error) {

	subscriptions.Delete(request.SubId)

	response = &pb.DeleteSubResponse{
		Result: "success",
	}

	return response, nil
}