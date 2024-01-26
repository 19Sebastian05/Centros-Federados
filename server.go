package main

import (
	"gcff-auth/api"

	"gcff-auth/internal/env"
)

func main() {
	c := env.NewConfiguration()
	api.Start(c.App.Port, c.App.ServiceName, c.App.LoggerHttp, c.App.AllowedDomains)
}
