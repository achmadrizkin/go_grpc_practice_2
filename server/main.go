package main

import (
	"context"
	"encoding/json"
	"fmt"
	pb "go_grpc_client_server/student"
	"io/ioutil"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type dataStudentServer struct {
	pb.UnimplementedDataStudentServer
	mu       sync.Mutex
	students []*pb.Student // dao
}

func (d *dataStudentServer) FindStudentByEmail(ctx context.Context, student *pb.Student) (*pb.Student, error) {
	for _, v := range d.students {
		if v.Email == student.Email {
			return v, nil
		}
	}

	return nil, nil
}

func (d *dataStudentServer) loadData() {
	data, err := ioutil.ReadFile("student.json")
	if err != nil {
		log.Fatalln("Error in read file: ", err.Error())
	}

	if err := json.Unmarshal(data, &d.students); err != nil {
		log.Fatalln("Error unmarshall file: ", err.Error())
	}

	fmt.Println(data)
}

func newServer() *dataStudentServer {
	s := dataStudentServer{}
	s.loadData()
	return &s
}

func main() {
	listen, err := net.Listen("tcp", ":1200")
	if err != nil {
		log.Fatalln("Error in listen: ", err.Error())
	}

	grpcServer := grpc.NewServer()
	pb.RegisterDataStudentServer(grpcServer, newServer())

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalln("Error in serve grpc: ", err.Error())
	}
}
