package tapo

import (
	"github.com/timurkash/go2rtc/internal/streams"
	"github.com/timurkash/go2rtc/pkg/core"
	"github.com/timurkash/go2rtc/pkg/tapo"
)

func Init() {
	streams.HandleFunc("tapo", handle)
}

func handle(url string) (core.Producer, error) {
	conn := tapo.NewClient(url)
	if err := conn.Dial(); err != nil {
		return nil, err
	}
	return conn, nil
}
