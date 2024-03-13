package main

import (
	"log/slog"
	"os"

	"github.com/lingdor/go-logcar/slogcar"
)

func main() {

	logger := slog.New(slogcar.NewTextHandler(os.Stdout, nil))
	logger.Info("hello world")
	logger.Warn("hello warn")

}
