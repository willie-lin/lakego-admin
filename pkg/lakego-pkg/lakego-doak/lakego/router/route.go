package router

import (
    "sync"
)

var instanceRoute *Route
var onceRoute sync.Once

func NewRoute() *Route {
    onceRoute.Do(func() {
        instanceRoute = &Route{}
    })

    return instanceRoute
}

/**
 * 缓存路由信息
 *
 * @create 2021-9-7
 * @author deatil
 */
type Route struct {
    // 路由
    routeEngine *Engine
}

// 设置
func (this *Route) With(engine *Engine) {
    this.routeEngine = engine
}

// 设置
func (this *Route) Get() *Engine {
    return this.routeEngine
}

// 路由信息
/*
type RouteInfo struct {
    Method      string
    Path        string
    Handler     string
    HandlerFunc HandlerFunc
}
RoutesInfo []RouteInfo
*/
func (this *Route) GetRoutes() RoutesInfo {
    return this.routeEngine.Routes()
}

// 路由信息
func (this *Route) GetRouteMap() map[string]interface{} {
    routes := this.GetRoutes()

    newRoutes := make(map[string]interface{})
    for _, v := range routes {
        if newRoute, ok := newRoutes[v.Method]; ok {
            newRoute = append(newRoute.([]string), v.Path)
            newRoutes[v.Method] = newRoute
        } else {
            newRoutes[v.Method] = []string{v.Path}
        }
    }

    return newRoutes
}

// 最后一个
func (this *Route) GetLastRoute() RouteInfo {
    routes := this.routeEngine.Routes()

    return routes[len(routes) - 1]
}
