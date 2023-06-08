package debug

import (
	"github.com/timurkash/go2rtc/internal/api"
	"github.com/timurkash/go2rtc/internal/streams"
	"github.com/timurkash/go2rtc/pkg/core"
)

func Init() {
	api.HandleFunc("api/stack", stackHandler)

	streams.HandleFunc("null", nullHandler)
}

func nullHandler(string) (core.Producer, error) {
	return nil, nil
}
