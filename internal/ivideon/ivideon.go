package ivideon

import (
	"github.com/timurkash/go2rtc/internal/streams"
	"github.com/timurkash/go2rtc/pkg/core"
	"github.com/timurkash/go2rtc/pkg/ivideon"
	"strings"
)

func Init() {
	streams.HandleFunc("ivideon", func(url string) (core.Producer, error) {
		id := strings.Replace(url[8:], "/", ":", 1)
		prod := ivideon.NewClient(id)
		if err := prod.Dial(); err != nil {
			return nil, err
		}
		return prod, nil
	})
}
