package handler

import (
	"fmt"
	"log"
	"net"

	"github.com/muhfaris/stockbit-test/soal-2/movies/configs"
	"golang.org/x/sync/errgroup"

	"github.com/soheilhy/cmux"

	"github.com/muhfaris/stockbit-test/soal-2/movies/grpc/server"
)

// InitRouter is create new handler
func InitRouter(app *configs.App) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", app.Port))
	if err != nil {
		log.Fatal(err)
	}

	m := cmux.New(listener)
	//grpcListener := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	grpcListener := m.Match(cmux.Any())
	httpListener := m.Match(cmux.HTTP1Fast())

	g := new(errgroup.Group)
	g.Go(func() error { return server.GRPCServe(app, grpcListener) })
	g.Go(func() error { return HTTPServe(app, httpListener) })
	g.Go(func() error { return m.Serve() })
	log.Println("run server at", app.Port, g.Wait())
}
