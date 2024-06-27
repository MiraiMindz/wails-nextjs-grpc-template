package api

import (
	"backend/proto/out"
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Response struct {
	Message  string `json:"message" xml:"message"`
}

type Request struct {
	Message  string `json:"message" xml:"message"`
}

type GrpcClient struct {
	conn *grpc.ClientConn
}

func NewGrpcClient(address string, opts []grpc.DialOption) (*GrpcClient, error) {
	connection, err := grpc.NewClient(address, opts...)
	if err != nil {
		return nil, err
	}

	return &GrpcClient{conn: connection}, nil
}

func (gc *GrpcClient) TestFunc(msg string) (*out.TestResponse, error) {
	testClient := out.NewTestClient(gc.conn)
	res, err := testClient.TestFunc(context.Background(), &out.TestRequest{Message: msg})
	return res, err
}

func Server() {
	gc,err := NewGrpcClient("localhost:50051", []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	})
	if err != nil {
		log.Fatalln(err.Error())
	}


	e := echo.New()
	e.Use(middleware.CORS())
	log.Println("Starting Echo HTTP Server")
	e.GET("/api/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, &Response{
			Message: "GO GET RESPONSE",
		})
	})

	e.POST("/api/test", func(c echo.Context) error {
		r := new(Request)
		if err := c.Bind(r); err != nil {
			return c.JSON(http.StatusBadRequest, &Response{
				Message: "BAD REQUEST",
			})
		}

		response, err := gc.TestFunc(r.Message)
		if err != nil {
			log.Fatalln(err.Error())
		}

		return c.JSON(http.StatusOK, response)
	})
	e.Logger.Fatal(e.Start(":1323"))
}