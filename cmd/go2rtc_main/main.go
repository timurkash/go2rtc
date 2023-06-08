package main

import (
	"github.com/timurkash/go2rtc/internal/api"
	"github.com/timurkash/go2rtc/internal/api/ws"
	"github.com/timurkash/go2rtc/internal/app"
	"github.com/timurkash/go2rtc/internal/debug"
	"github.com/timurkash/go2rtc/internal/dvrip"
	"github.com/timurkash/go2rtc/internal/echo"
	"github.com/timurkash/go2rtc/internal/exec"
	"github.com/timurkash/go2rtc/internal/ffmpeg"
	"github.com/timurkash/go2rtc/internal/hass"
	"github.com/timurkash/go2rtc/internal/hls"
	"github.com/timurkash/go2rtc/internal/homekit"
	"github.com/timurkash/go2rtc/internal/http"
	"github.com/timurkash/go2rtc/internal/isapi"
	"github.com/timurkash/go2rtc/internal/ivideon"
	"github.com/timurkash/go2rtc/internal/mjpeg"
	"github.com/timurkash/go2rtc/internal/mp4"
	"github.com/timurkash/go2rtc/internal/mpegts"
	"github.com/timurkash/go2rtc/internal/nest"
	"github.com/timurkash/go2rtc/internal/ngrok"
	"github.com/timurkash/go2rtc/internal/onvif"
	"github.com/timurkash/go2rtc/internal/roborock"
	"github.com/timurkash/go2rtc/internal/rtmp"
	"github.com/timurkash/go2rtc/internal/rtsp"
	"github.com/timurkash/go2rtc/internal/srtp"
	"github.com/timurkash/go2rtc/internal/streams"
	"github.com/timurkash/go2rtc/internal/tapo"
	"github.com/timurkash/go2rtc/internal/webrtc"
	"github.com/timurkash/go2rtc/internal/webtorrent"
	"github.com/timurkash/go2rtc/pkg/shell"
)

func main() {
	// 1. Core modules: app, api/ws, streams

	app.Init() // init config and logs

	api.Init() // init API before all others
	ws.Init()  // init WS API endpoint

	streams.Init() // streams module

	// 2. Main sources and servers

	rtsp.Init()   // rtsp source, RTSP server
	webrtc.Init() // webrtc source, WebRTC server

	// 3. Main API

	mp4.Init()   // MP4 API
	hls.Init()   // HLS API
	mjpeg.Init() // MJPEG API

	// 4. Other sources and servers

	hass.Init()       // hass source, Hass API server
	onvif.Init()      // onvif source, ONVIF API server
	webtorrent.Init() // webtorrent source, WebTorrent module

	// 5. Other sources

	rtmp.Init()     // rtmp source
	exec.Init()     // exec source
	ffmpeg.Init()   // ffmpeg source
	echo.Init()     // echo source
	ivideon.Init()  // ivideon source
	http.Init()     // http/tcp source
	dvrip.Init()    // dvrip source
	tapo.Init()     // tapo source
	isapi.Init()    // isapi source
	mpegts.Init()   // mpegts passive source
	roborock.Init() // roborock source
	homekit.Init()  // homekit source
	nest.Init()     // nest source

	// 6. Helper modules

	ngrok.Init() // Ngrok module
	srtp.Init()  // SRTP server
	debug.Init() // debug API

	// 7. Go

	shell.RunUntilSignal()
}
