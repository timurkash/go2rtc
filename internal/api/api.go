package api

import (
	"encoding/json"
	"github.com/rs/zerolog"
	"github.com/timurkash/go2rtc/internal/app"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

func Init() {
	var cfg struct {
		Mod struct {
			Listen    string `yaml:"listen"`
			Username  string `yaml:"username"`
			Password  string `yaml:"password"`
			BasePath  string `yaml:"base_path"`
			StaticDir string `yaml:"static_dir"`
			Origin    string `yaml:"origin"`
		} `yaml:"api"`
	}

	// default config
	cfg.Mod.Listen = ":1984"

	// load config from YAML
	app.LoadConfig(&cfg)

	if cfg.Mod.Listen == "" {
		return
	}

	basePath = cfg.Mod.BasePath
	log = app.GetLogger("api")

	initStatic(cfg.Mod.StaticDir)

	HandleFunc("api", apiHandler)
	HandleFunc("api/config", configHandler)
	HandleFunc("api/exit", exitHandler)

	// ensure we can listen without errors
	listener, err := net.Listen("tcp", cfg.Mod.Listen)
	if err != nil {
		log.Fatal().Err(err).Msg("[api] listen")
		return
	}

	log.Info().Str("addr", cfg.Mod.Listen).Msg("[api] listen")

	Handler = http.DefaultServeMux // 4th

	if cfg.Mod.Origin == "*" {
		Handler = middlewareCORS(Handler) // 3rd
	}

	if cfg.Mod.Username != "" {
		Handler = middlewareAuth(cfg.Mod.Username, cfg.Mod.Password, Handler) // 2nd
	}

	if log.Trace().Enabled() {
		Handler = middlewareLog(Handler) // 1st
	}

	go func() {
		s := http.Server{}
		s.Handler = Handler
		if err = s.Serve(listener); err != nil {
			log.Fatal().Err(err).Msg("[api] serve")
		}
	}()
}

var Handler http.Handler

// HandleFunc handle pattern with relative path:
// - "api/streams" => "{basepath}/api/streams"
// - "/streams"    => "/streams"
func HandleFunc(pattern string, handler http.HandlerFunc) {
	if len(pattern) == 0 || pattern[0] != '/' {
		pattern = basePath + "/" + pattern
	}
	log.Trace().Str("path", pattern).Msg("[api] register path")
	http.HandleFunc(pattern, handler)
}

// ResponseJSON important always add Content-Type
// so go won't need to call http.DetectContentType
func ResponseJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}

func ResponsePrettyJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	_ = enc.Encode(v)
}

func ResponseRawJSON(w http.ResponseWriter, s string) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(s))
}

func ResponseText(w http.ResponseWriter, b []byte) {
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write(b)
}

const StreamNotFound = "stream not found"

var basePath string
var log zerolog.Logger

func middlewareLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Trace().Msgf("[api] %s %s %s", r.Method, r.URL, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

func middlewareAuth(username, password string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.RemoteAddr, "127.") && !strings.HasPrefix(r.RemoteAddr, "[::1]") {
			user, pass, ok := r.BasicAuth()
			if !ok || user != username || pass != password {
				w.Header().Set("Www-Authenticate", `Basic realm="go2rtc"`)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

func middlewareCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization")
		next.ServeHTTP(w, r)
	})
}

var mu sync.Mutex

func apiHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	app.Info["host"] = r.Host
	mu.Unlock()

	ResponseJSON(w, app.Info)
}

func exitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	s := r.URL.Query().Get("code")
	code, _ := strconv.Atoi(s)
	os.Exit(code)
}

type Stream struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func ResponseStreams(w http.ResponseWriter, streams []Stream) {
	if len(streams) == 0 {
		http.Error(w, "no streams", http.StatusNotFound)
		return
	}

	var response = struct {
		Streams []Stream `json:"streams"`
	}{
		Streams: streams,
	}
	ResponseJSON(w, response)
}
