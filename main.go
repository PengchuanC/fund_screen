package main

import (
	"fmt"
	"fund_screen/app/routers"
	"fund_screen/app/rpc"
	"fund_screen/common/settings"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func main() {
	s := settings.GetSettings()
	gin.ForceConsoleColor()
	var r *gin.Engine
	r = gin.New()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST"},
		AllowWebSockets: true,
	}))
	r.Use(TlsHandler())
	routers.RegisterRouter(r)

	go rpc.Serve()

	go func() {
		err := r.RunTLS(fmt.Sprintf("%s:%d", s.Host, s.Ssl), "./external/cert/server.crt", "./external/cert/server.key")
		if err != nil {
			panic(err)
		}
	}()

	err := r.Run(fmt.Sprintf("%s:%d", s.Host, s.Port))
	if err != nil {
		panic(err)
	}
}

func TlsHandler() gin.HandlerFunc {
	s := settings.GetSettings()
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     fmt.Sprintf("%s:%d", s.Host, s.Ssl),
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}
