package main

import (
	"fmt"
	"log"

	"github.com/Azizbek-Qodirov/hr_platform_api_service/api"
	"github.com/Azizbek-Qodirov/hr_platform_api_service/api/handler"
	pb "github.com/Azizbek-Qodirov/hr_platform_api_service/genprotos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	UserConn, err := grpc.NewClient(fmt.Sprintf("localhost%s", ":8085"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while NEwclient: ", err.Error())
	}
	defer UserConn.Close()

	CompanyConn, err := grpc.NewClient(fmt.Sprintf("com%s", ":8086"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while NEwclient: ", err.Error())
	}

	defer CompanyConn.Close()
	us := pb.NewAuthServiceClient(UserConn)
	ev := pb.NewEvaulationServiceClient(CompanyConn)
	gu := pb.NewGuideServiceClient(CompanyConn)
	nt := pb.NewServiceNotificationClient(CompanyConn)

	h := handler.NewHandler(us, ev, gu, nt)
	r := api.NewGin(h)

	fmt.Println("Server started on port:8080")
	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Error while running server: ", err.Error())
	}
}
