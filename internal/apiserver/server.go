package apiserver

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/ividernvi/algohub/internal/apiserver/config"
	"github.com/ividernvi/algohub/internal/apiserver/middlewares"
	"github.com/ividernvi/algohub/internal/pkg/options"
)

type RESTfulServer struct {
	*gin.Engine

	healthz bool

	SecureServeOptions   *options.SecureServeOptions
	InsecureServeOptions *options.InsecureServeOptions
}

func NewServer(cfg *config.Config) *RESTfulServer {
	e := gin.Default()
	return &RESTfulServer{
		Engine:               e,
		healthz:              true,
		SecureServeOptions:   cfg.Options.SecureServeOptions,
		InsecureServeOptions: cfg.Options.InsecureServeOptions,
	}
}

func (rest *RESTfulServer) InitMiddleware() {
	rest.Use(middlewares.Logrus())
	rest.Use(gin.Recovery())
}

func (rest *RESTfulServer) RegisteAPI() {
	rest.Engine.GET("/version", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"version": version,
		})
	})

	rest.Engine.GET("healthz", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"healthz": "true",
		})
	})
}

func (rest *RESTfulServer) Setup() *RESTfulServer {
	rest.InitMiddleware()
	rest.RegisteAPI()
	RegisterRoutes(rest.Engine)
	return rest
}

func (rest *RESTfulServer) Run() error {
	ctx := context.Background()
	go func() {
		if err := rest.ListenAndServe(ctx); err != nil {
			panic(err)
		}
	}()

	if err := rest.ListenAndServeTLS(ctx); err != nil {
		panic(err)
	}

	return nil
}

func (rest *RESTfulServer) ListenAndServe(ctx context.Context) error {
	if err := rest.Engine.Run(rest.InsecureServeOptions.Port()); err != nil {
		return err
	}
	return nil
}

func (rest *RESTfulServer) ListenAndServeTLS(ctx context.Context) error {
	if err := rest.Engine.RunTLS(rest.SecureServeOptions.Port(), rest.SecureServeOptions.CertFile(), rest.SecureServeOptions.KeyFile()); err != nil {
		return err
	}
	return nil
}
