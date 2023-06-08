package main

import (
	"github.com/timurkash/go2rtc/internal/api"
	"github.com/timurkash/go2rtc/internal/api/ws"
	"github.com/timurkash/go2rtc/internal/app"
	"github.com/timurkash/go2rtc/internal/rtsp"
	"github.com/timurkash/go2rtc/internal/streams"
	"github.com/timurkash/go2rtc/internal/webrtc"
	"github.com/timurkash/go2rtc/pkg/shell"
)

func main() {
	app.Init() // init config and logs

	api.Init() // init API before all others
	ws.Init()  // init WS API endpoint

	streams.Init() // streams module

	// 2. Main sources and servers

	rtsp.Init() // rtsp source, RTSP server
	//rtmp.Init()
	webrtc.Init() // webrtc source, WebRTC server

	shell.RunUntilSignal()
}
