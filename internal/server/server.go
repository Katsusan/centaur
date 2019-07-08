package server

import (
	"context"

	"github.com/Katsusan/centaur/internal/config"
)

func Start(c context.Context, conf *config.Config) {
	if !conf.Debug() {

	}

}
