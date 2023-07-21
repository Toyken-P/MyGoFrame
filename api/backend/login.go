package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"myGoFrame/internal/model/entity"
	"time"
)

// for jwt
type LoginDoReq struct {
	// g.Meta   `path:"/login" tags:"login" method:"post"  summary:"执行登录请求"`
	Name     string `json:"name" v:"required#请输入账号名"   dc:"账号"`
	Password string `json:"password" v:"required#请输入密码"   dc:"密码(明文)"`
}

type LoginDoRes struct {
	// Referer string `json:"referer" dc:"引导客户端跳转地址"`
	// Info interface{} `json:"info"`
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

// for gtoken
type LoginRes struct {
	Type        string                  `json:"type"`
	Token       string                  `json:"token"`
	ExpireIn    int                     `json:"expire_in"`
	IsAdmin     int                     `json:"is_admin"`    //是否超管
	RoleIds     string                  `json:"role_ids"`    //角色
	Permissions []entity.PermissionInfo `json:"permissions"` //权限列表
}

type RefreshTokenReq struct {
	g.Meta `path:"/refresh_token" tags:"login" method:"post"`
}

type RefreshTokenRes struct {
	Token  string    `json:"token"`
	Expire time.Time `json:"expire"`
}

type LogoutReq struct {
	g.Meta `path:"/logout" tags:"login" method:"post"`
}

type LogoutRes struct {
}
