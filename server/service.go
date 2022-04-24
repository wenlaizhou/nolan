package server

import "github.com/wenlaizhou/middleware"

func Boot() {

	conf := middleware.LoadConfig("app.properties")

	middleware.RegisterHandler("", func(context middleware.Context) {

	})

	middleware.RegisterHandler("", func(context middleware.Context) {

	})

	middleware.RegisterHandler("", func(context middleware.Context) {

	})

	middleware.StartServer("", conf.IntUnsafe("port"))
}
