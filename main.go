package main

import flag "github.com/spf13/pflag"
import "cloudgo/service"
import "os"

const (
	PORT string = "8001"
)

func main() {
	var port string
	//获得环境变量
	port = os.Getenv("PORT")
	//如果获得的环境变量不存在，则赋值PORT
	if len(port) <= 0 {
		port = PORT
	}
	//定义服务器端口
	flag.StringVarP(&port, "port", "p", "8001", "define server port")
	flag.Parse()
	server := service.NewServer()
	server.Run(":" + port)
}
