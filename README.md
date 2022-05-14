go简单项目骨架
===============
个人自用项目骨架，非最佳实践，只是自己用习惯，也许会经常调整，如果被你看到了，请勿直接用于生产环境~

支持：
- 命令行工具开发(基于urfave/cli框架)
- HTTP服务开发(基于gin框架)

目录结构

```
bin #存放编译生成的程序或者本地测试文件，不会进入版本控制
commands # 命令行，每一个命令新建一个子目录
core # 全项目共享的代码
    config # 配置解析和管理
    libs # 自定义类库
    models # 数据库模型
    services # 服务
        database # 数据库服务
        redis # redis服务
server # 一个内置的HTTP服务，也是注册到commands里面，单独拧出来比较方便
    constants # 自定义常量，比如JWT的secret key
    controllers # 控制器目录
    middlewares # 中间件目录
    routes # 路由定义目录
    views # 视图模板目录
    server.go # 服务初始化
main.go # 主入口
config.toml.example # 配置文件示例
```
