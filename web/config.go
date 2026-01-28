package web

import (
	"embed"
	"os"
	"strings"
)

type config struct {
	WebServerAddr       string
	ProductsServiceAddr string
}

var Config config

//go:embed *.env*
var embededEnvFile embed.FS

func init() {
	file, err := embededEnvFile.ReadFile(".env")
	if err != nil {
		file, err = embededEnvFile.ReadFile(".env.example")
		if err != nil {
			return
		}
	}

	origEnvMap := make(map[string]bool)
	for _, line := range os.Environ() {
		key, _, found := strings.Cut(line, "=")
		if !found {
			continue
		}
		origEnvMap[key] = true
	}

	for line := range strings.SplitSeq(string(file), "\n") {
		key, val, found := strings.Cut(line, "=")
		if origEnvMap[key] || !found {
			continue
		}
		os.Setenv(key, strings.ReplaceAll(val, "\"", ""))
	}

	Config = config{
		WebServerAddr:       os.Getenv("WEB_SERVER_ADDR"),
		ProductsServiceAddr: os.Getenv("PRODUCTS_SERVICE_ADDR"),
	}
}
