package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"math/rand"
	"net"
	"time"
	"weather_service/api"
)

func main() {
	// listen port 8080
	// used tcp for http2
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	// create grpc server
	srv := grpc.NewServer()
	// register weather service server
	api.RegisterWeatherServiceServer(srv, &myWeatherService{})
	fmt.Println("starting server...")
	panic(srv.Serve(lis))
}

type myWeatherService struct {
	// api.UnsafeWeatherServiceServer can be used too
	api.UnimplementedWeatherServiceServer
}

func (m *myWeatherService) ListCities(ctx context.Context,
	req *api.ListCitiesRequest) (*api.ListCitiesResponse, error) {

	return &api.ListCitiesResponse{
		Items: []*api.CityEntry{
			&api.CityEntry{CityCode: "tr_ank",
				CityName: "Ankara"},
			&api.CityEntry{CityCode: "tr_ist",
				CityName: "Istanbul"},
		},
	}, nil
}

func (m *myWeatherService) QueryWeather(req *api.WeatherRequest,
	resp api.WeatherService_QueryWeatherServer) error {

	// generates random temperatures and sends continuously
	for {
		err := resp.Send(&api.WeatherResponse{Temperature: rand.Float32()*10 + 10})
		if err != nil {
			break
		}
		time.Sleep(time.Second)
	}
	return nil
}
