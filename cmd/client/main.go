package main

import (
	"context"
	"fmt"
	api "github.com/cookienyancloud/testrpckafkapsqlclick/protos/protos"
	"google.golang.org/grpc"
	"log"
	"strings"
	"time"
)

const ClientHost = ":8002"

func main() {
	conn, err := grpc.Dial(ClientHost, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("error connecting to %s : %v", ClientHost, err)
	}
	defer conn.Close()

	client := api.NewUsersClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	for {
		var inp string
		_, err := fmt.Scanln(&inp)
		if err != nil {
			log.Printf("error scanning: %v\n", err)
		}
		input := strings.Split(strings.ToLower(inp), " ")
		if len(input) == 0 {
			log.Println("no input")
			break
		}
		switch input[0] {

		case "create":
			if len(input) != 4 {
				log.Println("need 3 args")
				break
			}
			r, err := client.CreateUser(ctx, &api.CreateRequest{
				Name:     input[1],
				Password: input[2],
				Key:      input[3],
			})
			if err != nil {
				log.Printf("error creating user: %v\n", err)
			}
			log.Printf("response: %v\n", r.State)
		case "delete":
			if len(input) != 3 {
				log.Println("need 2 args")
				break
			}
			r, err := client.DeleteUser(ctx, &api.DeleteRequest{
				Id:  input[1],
				Key: input[2],
			})
			if err != nil {
				log.Printf("error deleting user: %v\n", err)
			}
			log.Printf("response: %v\n", r.State)
		case "find":
			if len(input) != 1 {
				log.Println("need 0 args")
				break
			}
			r, err := client.FindAll(ctx, &api.Empty{})
			if err != nil {
				log.Printf("error deleting user: %v\n", err)
			}
			log.Printf("response: %v\n", r.Users)
		case "exit":
			break
		default:
			log.Println("no such func")
		}
	}

}
