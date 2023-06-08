package homekit

import (
	"github.com/rs/zerolog"
	"github.com/timurkash/go2rtc/internal/api"
	"github.com/timurkash/go2rtc/internal/app"
	"github.com/timurkash/go2rtc/internal/srtp"
	"github.com/timurkash/go2rtc/internal/streams"
	"github.com/timurkash/go2rtc/pkg/core"
	"github.com/timurkash/go2rtc/pkg/homekit"
)

func Init() {
	log = app.GetLogger("homekit")

	streams.HandleFunc("homekit", streamHandler)

	api.HandleFunc("api/homekit", apiHandler)
}

var log zerolog.Logger

func streamHandler(url string) (core.Producer, error) {
	conn, err := homekit.NewClient(url, srtp.Server)
	if err != nil {
		return nil, err
	}
	if err = conn.Dial(); err != nil {
		return nil, err
	}
	return conn, nil
}
