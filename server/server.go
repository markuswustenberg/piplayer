package server

import (
	"context"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/maragudk/errors"

	"piplayer/player"
)

const shutdownTimeout = time.Minute

// Server handles HTTP requests.
type Server struct {
	host   string
	log    *log.Logger
	mux    chi.Router
	player *player.Player
	port   int
	server *http.Server
}

// Options for New.
type Options struct {
	Host   string
	Log    *log.Logger
	Player *player.Player
	Port   int
}

// New Server with the given Options.
func New(opts Options) *Server {
	if opts.Log == nil {
		opts.Log = log.New(ioutil.Discard, "", 0)
	}
	return &Server{
		host:   opts.Host,
		log:    opts.Log,
		mux:    chi.NewMux(),
		player: opts.Player,
		port:   opts.Port,
	}
}

func (s *Server) Start() error {
	s.setupRoutes()

	http.DefaultClient.Timeout = 10 * time.Second

	address := net.JoinHostPort(s.host, strconv.Itoa(s.port))
	s.server = &http.Server{
		Addr:              address,
		Handler:           s.mux,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       5 * time.Second,
		ErrorLog:          log.New(ioutil.Discard, "", 0),
	}

	s.log.Println("Starting on http://" + address)
	if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return errors.Wrap(err, "error serving http")
	}

	return nil
}

func (s *Server) Stop() error {
	s.log.Println("Stopping")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
