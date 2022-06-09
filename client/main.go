package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"weather_service/api"
)

func main() {
	addr := "localhost:8080"
	// establish connection, grpc.WithInsecure() is deprecated... idc about tls man
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	// create client
	client := api.NewWeatherServiceClient(conn)
	// context... super cool grpc thingy
	ctx := context.Background()
	// no need for ListCitiesRequest but fake it mb u improve dis
	resp, err := client.ListCities(ctx, &api.ListCitiesRequest{})
	if err != nil {
		panic(err)
	}

	fmt.Println("cities:")
	for _, city := range resp.Items {
		fmt.Printf("\t%s: %s\n", city.GetCityCode(), city.GetCityName())
	}
	// city code is not verified in the server but who cares
	stream, err := client.QueryWeather(ctx, &api.WeatherRequest{
		CityCode: "tr_ank",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("weather in Ankara:")
	for {
		// receiving end of file...
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Printf("\t temperature: %.2f\n", msg.GetTemperature())
	}
	fmt.Println("server stopped sending")
}
