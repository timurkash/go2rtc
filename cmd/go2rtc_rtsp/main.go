package main

import (
	"github.com/timurkash/go2rtc/internal/app"
	"github.com/timurkash/go2rtc/internal/rtsp"
	"github.com/timurkash/go2rtc/internal/streams"
	"github.com/timurkash/go2rtc/pkg/shell"
)

func main() {
	app.Init()
	streams.Init()

	rtsp.Init()

	shell.RunUntilSignal()
}
