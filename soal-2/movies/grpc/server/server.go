package server

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/muhfaris/stockbit-test/soal-2/movies/configs"
	"github.com/muhfaris/stockbit-test/soal-2/movies/domain"
	"github.com/muhfaris/stockbit-test/soal-2/movies/gateway/structures"
	pb "github.com/muhfaris/stockbit-test/soal-2/movies/grpc/gen/proto"
	"github.com/muhfaris/stockbit-test/soal-2/movies/services"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
)

type MovieAPIServer struct {
	pb.UnimplementedMovieAPIServer
	MovieService services.MovieService
	app          *configs.App
}

func (s *MovieAPIServer) SearchMovie(ctx context.Context, req *pb.MovieRequest) (*pb.ResponsesRequest, error) {
	log.Println("Search")
	searchword := req.GetSearchword()
	pagination := req.GetPagination()

	filter := structures.MovieRead{
		SearchWord: searchword,
		Pagination: int(pagination),
	}

	resp := s.MovieService.SearchMovie(ctx, filter)
	response, ok := resp.Data.(domain.MoviesResponseModel)
	if !ok {
		return &pb.ResponsesRequest{}, nil
	}

	data := response.ToProto()
	return &pb.ResponsesRequest{Movies: data}, nil
}

func (s *MovieAPIServer) GetDetailMovie(ctx context.Context, req *pb.DetailMovieRequest) (*pb.ResponseRequest, error) {
	imdbID := req.GetId()

	resp := s.MovieService.GetMovie(ctx, imdbID)

	response, ok := resp.Data.(domain.MovieResponseModel)
	if !ok {
		return &pb.ResponseRequest{}, nil
	}

	data := response.ToProto()
	return &pb.ResponseRequest{Data: data}, nil
}

const port = 8282

// GRPCServe is declare serve of grpc service
func GRPCServe(app *configs.App, l net.Listener) error {
	app.Logger.SetFormatter(&logrus.JSONFormatter{})
	app.Logger.SetLevel(logrus.DebugLevel)

	s := grpc.NewServer(
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             2 * time.Microsecond,
			PermitWithoutStream: true,
		}),
		grpc_middleware.WithUnaryServerChain(
			grpc_logrus.UnaryServerInterceptor(logrus.NewEntry(app.Logger)),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_logrus.StreamServerInterceptor(logrus.NewEntry(app.Logger)),
		),
	)

	movieAPI := &MovieAPIServer{
		MovieService: services.NewMoviceService(app),
		app:          app,
	}

	pb.RegisterMovieAPIServer(s, movieAPI)

	return s.Serve(l)
}
