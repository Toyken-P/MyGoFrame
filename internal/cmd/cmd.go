package cmd

import (
	"context"
	"myGoFrame/internal/consts"
	"myGoFrame/internal/controller"
	"myGoFrame/internal/controller/backend"
	"myGoFrame/internal/controller/frontend"
	"myGoFrame/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  consts.ProjectName,
		Usage: consts.ProjectUsage,
		Brief: consts.ProjectBrief,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// 启动管理后台gtoken
			gfAdminToken, err := StartBackendGToken()
			if err != nil {
				return err
			}
			//管理后台路由组
			s.Group("/backend", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().CORS, // 跨域认证
					service.Middleware().Ctx,  // contex链路追踪、数值传递
					service.Middleware().ResponseHandler,
				)
				//不需要登录的路由组绑定
				group.Bind(
					controller.Admin.Create, // 管理员
					controller.Login,        // 登录
				)
				//需要登录的路由组绑定
				group.Group("/", func(group *ghttp.RouterGroup) {
					err := gfAdminToken.Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					group.Bind(
						controller.Data,         // 数据大屏相关
						controller.Role,         // 角色
						controller.Permission,   // 权限
						controller.Admin.List,   // 管理员 增删改查
						controller.Admin.Update, // 管理员 增删改查
						controller.Admin.Delete, // 管理员 增删改查
						controller.Admin.Info,   // 管理员 增删改查
						controller.Rotation,     // 轮播图
						controller.Position,     // 手工位
						controller.File,         // 文件入库
						controller.Category,     // 商品分类管理
						controller.Coupon,       // 商品优惠券管理
						controller.UserCoupon,   // 商品优惠券管理
						controller.Goods,        // 商品管理
						controller.GoodsOptions, // 商品规格管理
						controller.Address,      // 城市地址管理
						controller.Order.List,   // 订单列表
						controller.Order.Detail, // 订单详情
						backend.Article,         // 文章管理&CMS
					)
				})
			})
			// 启动前台项目gtoken
			frontendToken, err := StartFrontendGToken()
			if err != nil {
				return err
			}
			//前台项目路由组
			s.Group("/frontend", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().CORS,
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				//不需要登录的路由组绑定
				group.Bind(
					controller.User.Register, // 用户注册
					controller.Goods,         // 商品
				)
				//需要登录鉴权的路由组
				group.Group("/", func(group *ghttp.RouterGroup) {
					err := frontendToken.Middleware(ctx, group)
					if err != nil {
						return
					}
					//需要登录鉴权的接口放到这里
					group.Bind(
						controller.User.Info,           // 当前登录用户的信息
						controller.User.UpdatePassword, // 当前用户修改密码
						controller.Collection,          // 收藏
						controller.Praise,              // 点赞
						controller.Comment,             // 评论
						controller.Cart,                // 购物车
						controller.Order.Add,           // 下单
						controller.OrderGoodsComments,  // 订单评价
						frontend.Article,               // 文章
						frontend.Refund,                // 售后
					)
				})
			})
			s.SetPort(8000) // 设置端口
			s.Run()
			return nil
		},
	}
)
