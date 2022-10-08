package routes

import (
	"encoding/gob"
	"io"
	userapi "myproject3/api/route"
	"myproject3/middleware"
	"myproject3/model"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

type Option func(engine *gin.Engine)

var Options = []Option{}

// 注册app的路由配置
func Include(opts ...Option) {
	Options = append(Options, opts...)
}

// 初始化
func Init() *gin.Engine {
	// 加载多路由
	Include(userapi.Routers)
	// 初始化日志
	gin.DisableConsoleColor()
	//// 创建日志文件
	f, _ := os.Create("myproject.log")
	//// 写入日志
	gin.DefaultWriter = io.MultiWriter(f)
	// 创建一个默认路由
	r := gin.Default()
	r.Use(cors.Default())

	r.Use(middleware.LoggerToFile())
	// //注册Users结构体，使其可以跨路由存取
	gob.Register(model.Users{})
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	store.Options(sessions.Options{
		MaxAge: 24 * 60 * 60,
		Path:   "/",
	})
	r.Use(sessions.Sessions("session", store))
	r.Use(gin.Recovery())
	// 抛出指针
	// 加载注册的app路由
	for _, opt := range Options {
		opt(r)
	}
	return r
}
