package server

import (
	"github.com/wenlaizhou/middleware"
)

var Swagger = middleware.SwaggerBuildModel("", "", "")

func Boot() {

	conf := middleware.LoadConfig("app.properties")

	Swagger.AddPath(middleware.SwaggerBuildPath("/clusterMap", "", "GET", ""))
	middleware.RegisterHandler("/clusterMap", func(context middleware.Context) {
		if !checkAuth(conf, context.GetHeader("Authorization")) {
			context.Error(403, "NO AUTHORITY")
			return
		}

	})

	middleware.RegisterHandler("", func(context middleware.Context) {

	})

	middleware.RegisterHandler("", func(context middleware.Context) {

	})

	middleware.EnableSwagger(Swagger)

	middleware.RegisterIndex(func(context middleware.Context) {
		context.Redirect("/swagger-ui")
	})

	middleware.StartServer("", conf.IntUnsafe("port"))
}
