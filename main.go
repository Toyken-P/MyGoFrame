package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	_ "myGoFrame/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"myGoFrame/internal/cmd"
	_ "myGoFrame/internal/logic"
)

func main() {
	cmd.Main.Run(gctx.New())
}
