package di

import (
	"io/fs"
	"myapp/core/config"
	"myapp/core/libs/echopongo2"
	"myapp/core/services/database"
	"myapp/core/services/redis"
	"sync"
)

type Container struct {
	serviceSet sync.Map
}

var defaultContainer *Container

func init() {
	defaultContainer = NewDI()
}

func Default() *Container {
	return defaultContainer
}

func NewDI() *Container {
	return &Container{}
}

func (container *Container) get(name string) interface{} {
	v, _ := container.serviceSet.Load(name)
	return v
}

func (container *Container) set(name string, v interface{}) interface{} {
	actual, _ := container.serviceSet.LoadOrStore(name, v)
	return actual
}

// GetTplRenderer 模板渲染引擎
func (container *Container) GetTplRenderer(name string, f fs.FS) *echopongo2.Renderer {
	serviceName := "TplRenderer#" + name
	if v := container.get(serviceName); v != nil {
		return v.(*echopongo2.Renderer)
	}
	renderer := echopongo2.NewRendererOverFs(serviceName, f)
	return container.set(serviceName, renderer).(*echopongo2.Renderer)
}

// GetDB 获取数据库连接 connectName对应配置文件database.xxx
func (container *Container) GetDB(connectName string) *database.Connection {
	if connectName == "" {
		return nil
	}
	serviceName := "database#" + connectName
	if v := container.get(serviceName); v != nil {
		return v.(*database.Connection)
	}
	host := config.GetString("databases.%s.host", connectName)
	dbname := config.GetString("databases.%s.dbname", connectName)
	user := config.GetString("databases.%s.user", connectName)
	password := config.GetString("databases.%s.password", connectName)
	port := config.GetInt("databases.%s.port", connectName)
	connection := database.Connect(host, dbname, user, password, port)
	return container.set(serviceName, connection).(*database.Connection)
}

// GetRedis 获取redis服务
func (container *Container) GetRedis() *redis.RedisClient {
	serviceName := "Redis"
	if v := container.get(serviceName); v != nil {
		return v.(*redis.RedisClient)
	}
	redisClient := redis.NewRedisClient()
	return container.set(serviceName, redisClient).(*redis.RedisClient)
}
