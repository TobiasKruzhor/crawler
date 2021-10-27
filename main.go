package main

import (
	"github.com/TobiasKruzhor/crawler/model"
	"github.com/TobiasKruzhor/crawler/routues"
)

func main() {
	// 引用数据库
	model.InitDb()

	routues.InitRouter()
}
