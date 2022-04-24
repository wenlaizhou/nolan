package server

import (
	"github.com/wenlaizhou/middleware"
)

var Swagger = middleware.SwaggerBuildModel("", "", "")

func Boot() {

	conf := middleware.LoadConfig("app.properties")

	Swagger.AddPath(middleware.SwaggerBuildPath("", "", "", ""))
	middleware.RegisterHandler("/clusterMap", func(context middleware.Context) {
		if !checkAuth(conf, context.GetHeader("Authorization")) {
			context.Error(404, "")
			return
		}

	})

	middleware.RegisterHandler("", func(context middleware.Context) {

	})

	middleware.RegisterHandler("", func(context middleware.Context) {

	})

	middleware.StartServer("", conf.IntUnsafe("port"))
}
