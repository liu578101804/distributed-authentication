package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/multitemplate"
	"github.com/liu578101804/distributed-authentication/backend/web/controllers"
	"path/filepath"
)

func main() {

	router := gin.Default()

	//加载模板渲染
	router.HTMLRender = loadTemplates("./web/views")

	//加载静态文件
	router.Static("assets", "web/assets")
	router.StaticFile("/favicon.ico", "web/assets/favicon.ico")

	//注册控制器
	controllers.RegisterController(router.RouterGroup)

	router.Run(":8083")
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layout,err := filepath.Glob(templatesDir + "/shared/layout.html")
	if err != nil {
		panic(err.Error())
	}
	views,err := filepath.Glob(templatesDir + "/home/*.html")
	if err != nil {
		panic(err.Error())
	}

	for _,view := range views {
		layoutCopy := make([]string, len(views))
		copy(layoutCopy, layout)
		files := append(layoutCopy, view)
		r.AddFromFiles(filepath.Base(view), files...)
	}
	return r
}
