package router

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"tf-cluster/app/api/auth"
	"tf-cluster/app/api/cluster"
	_ "tf-cluster/docs" // gin-swagger
	"tf-cluster/internal/middleware/cors"

	//_ "tf-watcher/internal/rice"
	//"tf-watcher/library/utils"
	//rice "github.com/GeertJohan/go.rice"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

var (
	Router *gin.Engine
)

func RegisterRouter() {
	Router = gin.Default()
	// pprof
	pprof.Register(Router)
	Router.Use(cors.Cors())
	//fs := utils.EmbeddingFileSystem(rice.MustFindBox(enum.RicePath).HTTPBox())
	//Router.Use(utils.Serve("", fs))
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	dm := Router.Group("auth")
	//dm.Use(jwt.JWT()) # jwt auth
	dm.GET("/test", auth.TestApi)
	dm.GET("/hello-list", auth.HelloListApi)
	dm.GET("/hello-info", auth.HelloInfoApi)
	dm.GET("/user-list", auth.UserListApi)
	dm.GET("/user-count", auth.UserCountApi)
	dm.GET("/user-detail", auth.UserDetailApi)
	dm.POST("/user-delete", auth.UserDeleteApi)
	dm.POST("/user-change", auth.UserChangeStatusApi)
	dm.POST("/user-save", auth.UserSaveApi)
	us := Router.Group("user")
	us.Any("/login", auth.LoginApi)
	us.GET("/info", auth.InfoApi)

	// 集群
	cls := Router.Group("watcher")
	// 节点分组
	node := cls.Group("node")
	node.GET("/list", cluster.NodeListApi)
	node.GET("/info", cluster.NodeInfoApi)
	// 负载分组
	deploy := cls.Group("deploy")
	deploy.GET("/list", cluster.DeployListApi)
	deploy.GET("/info", cluster.DeployInfoApi)
	// 任务
	tk := Router.Group("task")
	tk.GET("list")
}
