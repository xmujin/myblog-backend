package main

import (
	"github.com/xmujin/myblog-backend/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
