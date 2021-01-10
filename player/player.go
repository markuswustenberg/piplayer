package player

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/maragudk/errors"
)

type Options struct {
	Dir      string
	Host     string
	Log      *log.Logger
	Password string
	Port     int
}

type Player struct {
	client   *http.Client
	dir      string
	endpoint string
	log      *log.Logger
	password string
}

func New(opts Options) *Player {
	if opts.Log == nil {
		opts.Log = log.New(ioutil.Discard, "", 0)
	}
	return &Player{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		dir:      opts.Dir,
		endpoint: "http://" + net.JoinHostPort(opts.Host, strconv.Itoa(opts.Port)),
		log:      opts.Log,
		password: opts.Password,
	}
}

func (p *Player) Play(ctx context.Context, path string) error {
	p.log.Println("Play " + path)

	return p.request(ctx, "in_play&input="+escape(path))
}

func (p *Player) Pause(ctx context.Context) error {
	p.log.Println("Pause")

	return p.request(ctx, "pl_pause")
}

func (p *Player) request(ctx context.Context, command string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, p.endpoint+"/requests/status.xml?command="+command, nil)
	if err != nil {
		return errors.Wrap(err, "could not create request for media player")
	}
	req.SetBasicAuth("", p.password)

	res, err := p.client.Do(req)
	if err != nil {
		return errors.Wrap(err, "could not contact media player")
	}
	if res.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("got status code %v from media player", res.StatusCode))
	}
	return nil
}

// escape a path so VLC understands it.
func escape(path string) string {
	// For some reason VLC does not understand the format used by url.QueryEscape
	return strings.ReplaceAll(url.PathEscape(path), "&", "%26")
}
