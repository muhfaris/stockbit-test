package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	pb "github.com/muhfaris/stockbit-test/soal-2/movies/grpc/gen/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:8989",
		grpc.WithInsecure(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                1 * time.Millisecond,
			Timeout:             2 * time.Millisecond,
			PermitWithoutStream: true,
		}))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := pb.NewMovieAPIClient(conn)
	// search movie
	searchResp, err := client.SearchMovie(context.Background(), &pb.MovieRequest{Searchword: "batman"})
	if err != nil {
		fmt.Printf("error: search movie %v", err)
		return
	}

	fmt.Println("search batman")
	fmt.Println(searchResp)

	imdbID := searchResp.Movies[0].GetImdbID()
	movieResp, err := client.GetDetailMovie(context.Background(), &pb.DetailMovieRequest{Id: imdbID})
	if err != nil {
		fmt.Printf("error: search movie %v", err)
	}

	fmt.Println("search detail movie")
	fmt.Println(movieResp)
}
