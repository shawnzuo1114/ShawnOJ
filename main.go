package main

import (
	"ShawnOJ/internal/router"
	"ShawnOJ/utils"
	"ShawnOJ/utils/database"
	"ShawnOJ/utils/setting"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	if err := utils.UtilsInit(); err != nil {
		zap.L().Error("utils init err", zap.Error(err))
		return
	}
	setupService()
}

func setupService() {
	//注册路由
	r := routers.Setup(setting.Conf.Mode)
	//启动服务
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", setting.Conf.Host, setting.Conf.Port),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1)                      // 创建一个接收信号的通道
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	defer database.CloseMysql()
	defer database.CloseRedis()

	zap.L().Info("Server exiting")
}
