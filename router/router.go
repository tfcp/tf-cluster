package router

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"tf-cluster/app/api/auth"
	"tf-cluster/app/api/cluster"
	_ "tf-cluster/docs" // gin-swagger
	"tf-cluster/internal/middleware/cors"
	"tf-cluster/internal/middleware/jwt"

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
	at := Router.Group("auth")
	at.Use(jwt.JWT())
	at.GET("/test", auth.TestApi)
	at.GET("/hello-list", auth.HelloListApi)
	at.GET("/hello-info", auth.HelloInfoApi)
	at.GET("/user-list", auth.UserListApi)
	at.GET("/user-count", auth.UserCountApi)
	at.GET("/user-info", auth.UserInfoApi)
	at.POST("/user-delete", auth.UserDeleteApi)
	at.POST("/user-change", auth.UserChangeStatusApi)
	at.POST("/user-save", auth.UserSaveApi)
	at.POST("/pwd-edit", auth.UserPwdChangeApi)
	us := Router.Group("user")
	us.Any("/login", auth.LoginApi)
	us.GET("/info", auth.InfoApi)

	// 集群
	cls := Router.Group("cluster")
	// 集群配置
	cls.GET("/config-list", cluster.ConfigListApi)
	cls.GET("/config-count", cluster.ConfigCountApi)
	cls.POST("/config-save", cluster.ConfigSaveApi)
	cls.GET("/config-info", cluster.ConfigInfoApi)

	// 节点分组
	cls.GET("/node-list", cluster.NodeListApi)
	cls.GET("/node-count", cluster.NodeCountApi)
	cls.GET("/node-info", cluster.NodeInfoApi)

	// 负载分组
	deploy := cls.Group("deploy")
	deploy.GET("/list", cluster.DeployListApi)
	deploy.GET("/info", cluster.DeployInfoApi)
	// 任务
	tk := Router.Group("task")
	tk.GET("list")
}
