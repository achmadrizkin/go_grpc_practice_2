package main

import (
	"fmt"
	pb "go_grpc_client_server/student"
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func getDataStudentByEmail(client pb.DataStudentClient, email string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s := pb.Student{Email: email}
	student, err := client.FindStudentByEmail(ctx, &s)
	if err != nil {
		log.Fatalln("error when get student by email: ", err.Error())
	}

	fmt.Println(student)
}

func main() {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(":1200", opts...)
	if err != nil {
		log.Fatalln("error in dial: ", err.Error())
	}

	defer conn.Close()
	client := pb.NewDataStudentClient(conn)
	getDataStudentByEmail(client, "aceng@gmail.com")
	getDataStudentByEmail(client, "aceng@gmail.com123132123")
}
