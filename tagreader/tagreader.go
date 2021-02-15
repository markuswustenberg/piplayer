package tagreader

import (
	"bytes"
	"encoding/binary"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/maragudk/errors"
)

type TagReader struct {
	debug bool
	f     *os.File
	log   *log.Logger
}

type Options struct {
	Debug bool
	Log   *log.Logger
}

func New(opts Options) *TagReader {
	if opts.Log == nil {
		opts.Log = log.New(ioutil.Discard, "", 0)
	}
	return &TagReader{
		debug: opts.Debug,
		log:   opts.Log,
	}
}

func (t *TagReader) Open(path string) error {
	var err error
	t.f, err = os.Open(path)
	if err != nil {
		return err
	}
	return nil
}

func (t *TagReader) Read() (string, error) {
	var id string
	for {
		var err error
		var event InputEvent
		buffer := make([]byte, eventSize)
		if _, err = t.f.Read(buffer); err != nil && !errors.Is(err, io.EOF) {
			return "", err
		}

		b := bytes.NewBuffer(buffer)
		if err := binary.Read(b, binary.LittleEndian, &event); err != nil {
			return "", err
		}

		if t.debug {
			t.log.Println(event.String())
		}

		if event.Type == EV_KEY && event.Value == KeyDown {
			if event.Code == KEY_ENTER {
				break
			}
			id += strings.TrimPrefix(KEY[int(event.Code)], "KEY_")
		}
	}

	return id, nil
}

func (t *TagReader) Close() error {
	return t.f.Close()
}
