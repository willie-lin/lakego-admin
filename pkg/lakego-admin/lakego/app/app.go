package app

import (
    "net"
    "sync"

    "github.com/spf13/cobra"
    "github.com/gin-gonic/gin"

    "github.com/deatil/lakego-admin/lakego/route"
    "github.com/deatil/lakego-admin/lakego/support/path"
    "github.com/deatil/lakego-admin/lakego/middleware/event"
    "github.com/deatil/lakego-admin/lakego/facade/config"
    "github.com/deatil/lakego-admin/lakego/facade/router"
    providerInterface "github.com/deatil/lakego-admin/lakego/provider/interfaces"
)

// App结构体
func New() *App {
    return &App{
        Runned: false,
    }
}

var serviceProviderLock = new(sync.RWMutex)

// 注册的服务提供者
var serviceProviders = []func() providerInterface.ServiceProvider{}

// 已使用服务提供者
var usedServiceProviders []providerInterface.ServiceProvider

/**
 * App结构体
 *
 * @create 2021-6-19
 * @author deatil
 */
type App struct {
    // 运行状态
    Runned bool

    // 运行在命令行
    RunInConsole bool

    // 路由
    RouteEngine *gin.Engine

    // 根脚本
    RootCmd *cobra.Command

    // 启动前
    BootingCallbacks []func()

    // 启动后
    BootedCallbacks []func()

    // 自定义运行监听
    NetListener net.Listener
}

// 运行
func (app *App) Run() {
    // 运行
    app.runApp()
}

// 注册服务提供者
func (app *App) Register(f func() providerInterface.ServiceProvider) {
    serviceProviderLock.Lock()
    defer serviceProviderLock.Unlock()

    serviceProviders = append(serviceProviders, f)

    // 启动后注册，直接注册
    if app.Runned {
        p := f()

        // 绑定 app 结构体
        p.WithApp(app)

        // 路由
        p.WithRoute(app.RouteEngine)

        // 注册
        p.Register()

        // 引导
        app.BootService(p)
    }
}

// 批量导入
func (app *App) Registers(providers []func() providerInterface.ServiceProvider) {
    if len(providers) > 0 {
        for _, provider := range providers {
            app.Register(provider)
        }
    }
}

// 加载服务提供者
func (app *App) loadServiceProvider() {
    if len(serviceProviders) > 0 {
        for _, provider := range serviceProviders {
            p := provider()

            // 绑定 app 结构体
            p.WithApp(app)

            // 路由
            p.WithRoute(app.RouteEngine)

            p.Register()

            usedServiceProviders = append(usedServiceProviders, p)
        }
    }

    // 启动前
    app.CallBootingCallbacks()

    if len(usedServiceProviders) > 0 {
        for _, sp := range usedServiceProviders {
            app.BootService(sp)
        }
    }

    // 启动后
    app.CallBootedCallbacks()
}

// 引导服务
func (app *App) BootService(s providerInterface.ServiceProvider) {
    s.CallBootingCallback()

    // 启动
    s.Boot()

    s.CallBootedCallback()
}

// 设置启动前函数
func (app *App) WithBooting(f func()) {
    app.BootingCallbacks = append(app.BootingCallbacks, f)
}

// 设置启动后函数
func (app *App) WithBooted(f func()) {
    app.BootedCallbacks = append(app.BootedCallbacks, f)
}

// 启动前回调
func (app *App) CallBootingCallbacks() {
    for _, callback := range app.BootingCallbacks {
        callback()
    }
}

// 启动后回调
func (app *App) CallBootedCallbacks() {
    for _, callback := range app.BootedCallbacks {
        callback()
    }
}

// 设置根脚本
func (app *App) WithRootCmd(root *cobra.Command) {
    app.RootCmd = root
}

// 获取根脚本
func (app *App) GetRootCmd() *cobra.Command {
    return app.RootCmd
}

// 设置命令行状态
func (app *App) WithRunningInConsole(console bool) {
    app.RunInConsole = console
}

// 获取命令行状态
func (app *App) RunningInConsole() bool {
    return app.RunInConsole
}

// 设置自定义监听
func (app *App) WithNetListener(listener net.Listener) *App {
    app.NetListener = listener

    return app
}

// 初始化路由
func (app *App) runApp() {
    var r *gin.Engine

    if !app.RunInConsole {
        // 模式
        mode := config.New("admin").GetString("Mode")
        if mode != "dev" {
            gin.SetMode(gin.ReleaseMode)

            // 路由
            r = gin.New()

            // 使用默认处理机制
            r.Use(gin.Recovery())
        } else {
            gin.SetMode(gin.DebugMode)

            // 路由
            r = gin.Default()
        }
    } else {
        // 脚本取消调试模式
        gin.SetMode(gin.ReleaseMode)

        // 路由
        r = gin.New()
    }

    // 事件
    r.Use(event.Handler())

    // 全局中间件
    globalMiddlewares := router.GetGlobalMiddlewares()

    // 设置全局中间件
    r.Use(globalMiddlewares...)

    // 缓存路由信息
    route.New().With(r)

    // 绑定路由
    app.RouteEngine = r

    // 设置已启动
    app.Runned = true

    // 加载服务提供者
    app.loadServiceProvider()

    // 不是命令行运行
    if !app.RunInConsole {
        app.ServerRun()
    }
}

// 服务运行
func (app *App) ServerRun() {
    conf := config.New("server")

    // 运行方式
    runType := conf.GetString("RunType")
    switch runType {
        case "Http":
            // 运行端口
            httpPort := conf.GetString("Types.Http.Port")

            app.RouteEngine.Run(httpPort)

        case "TLS":
            // 运行端口
            httpPort := conf.GetString("Types.TLS.Port")

            certFile := conf.GetString("Types.TLS.CertFile")
            keyFile := conf.GetString("Types.TLS.KeyFile")

            // 格式化
            certFile = app.FormatPath(certFile)
            keyFile = app.FormatPath(keyFile)

            app.RouteEngine.RunTLS(httpPort, certFile, keyFile)

        case "Unix":
            // 文件
            file := conf.GetString("Types.Unix.File")

            // 格式化
            file = app.FormatPath(file)

            app.RouteEngine.RunUnix(file)

        case "Fd":
            // fd
            fd := conf.GetInt("Types.Fd.Fd")

            app.RouteEngine.RunFd(fd)

        case "NetListener":
            // 监听
            listen := conf.GetBool("Types.NetListener.Listen")

            if listen && app.NetListener != nil {
                app.RouteEngine.RunListener(app.NetListener)
            } else {
                panic("NetListener 服务启动错误")
            }
        default:
            panic("服务启动错误")
    }
}

// 格式化文件路径
func (app *App) FormatPath(file string) string {
    filename := path.FormatPath(file)

    return filename
}