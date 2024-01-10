package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"dispatchAst/internal/controller/hello"
	"dispatchAst/internal/controller/qian_extension"
	"dispatchAst/internal/controller/qian_force"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			moduleHello := hello.New()
			moduleQianlue := qian_extension.New()
			moduleForce := qian_force.New()

			s.BindObject("/", moduleHello)

			// 路由嵌套分组 ip:port/qianlue/
			s.Group("/qianlue", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Group("/force", func(g1 *ghttp.RouterGroup) {
					g1.Middleware(ghttp.MiddlewareHandlerResponse)

					g1.REST("/", moduleForce)
				})
				group.Group("/api", func(g2 *ghttp.RouterGroup) {
					g2.Middleware(ghttp.MiddlewareHandlerResponse)

					g2.Bind( // Bind 可以让控制器里的所有方法都能访问
						moduleQianlue,
					)

					//	g2.REST("/", moduleQianlue)
				})
			})

			s.SetServerRoot("resource/public") // 开启静态资源服务
			s.SetPort(2024)                    // default :8000
			s.Run()
			return nil
		},
	}
)
