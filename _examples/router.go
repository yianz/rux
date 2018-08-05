package main

import (
	"github.com/gookit/sux"
	"fmt"
	"net/http"
	"time"
	"log"
)

func main() {
	r := sux.New()
	r.Use()

	r.GET("/", func(ctx *sux.Context) {
		ctx.WriteBytes([]byte("hello, in " + ctx.URL().Path))
	})

	r.GET("/about[.html]", defHandle)

	r.GET("/hi-{name}", defHandle)

	r.Group("/users", func(sub *sux.Router) {
		sub.GET("/", func(ctx *sux.Context) {
			ctx.WriteBytes([]byte("hello, in " + ctx.URL().Path))
		})

		sub.GET("/{id}", func(ctx *sux.Context) {
			ctx.WriteBytes([]byte("hello, in " + ctx.URL().Path))
		})
	})

	r.Controller("/site", &SiteController{})

	fmt.Println(r)

	ret := r.Match("GET", "/hi-tom")
	ret1 := r.Match("GET", "/hi-john")

	fmt.Println(ret)
	fmt.Println(ret1)

	// r.RunServe(":8090")
}

func defHandle(ctx *sux.Context) {
	ctx.WriteBytes([]byte("hello, in " + ctx.URL().Path))
}

type SiteController struct {
}

func (c *SiteController) AddRoutes(r *sux.Router) {
	r.GET("{id}", c.Get)
	r.POST("", c.Post)
}

func (c *SiteController) Get(ctx *sux.Context) {
	ctx.WriteBytes([]byte("hello, in " + ctx.URL().Path))
	ctx.WriteBytes([]byte("\n ok"))
}

func (c *SiteController) Post(ctx *sux.Context) {
	ctx.WriteBytes([]byte("hello, in " + ctx.URL().Path))
}


func customServer() {
	r := sux.New()

	// add routes
	r.GET("/", func(ctx *sux.Context) {
		ctx.WriteString("hello")
	})

	s := &http.Server{
		Addr:    ":8080",
		Handler: r,

		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
