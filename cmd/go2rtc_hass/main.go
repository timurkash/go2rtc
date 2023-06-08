package main

import (
	"github.com/timurkash/go2rtc/internal/api"
	"github.com/timurkash/go2rtc/internal/app"
	"github.com/timurkash/go2rtc/internal/hass"
	"github.com/timurkash/go2rtc/internal/streams"
	"github.com/timurkash/go2rtc/pkg/shell"
)

func main() {
	app.Init()
	streams.Init()

	api.Init()

	hass.Init()

	shell.RunUntilSignal()
}
