package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	
	gw "endpoint/proto"
	"srv"
	
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

const (
	grpcAddr = ":8090" // grpc 服务端口
	httpAddr = ":8080" // http 协议端口
)

func main() {
	
	{
		// grpc server
		rpcServer := NewGrpcServer()
		// 注册 grpc 处理函数
		gw.RegisterSayHiServer(rpcServer, srv.NewSayHiSrv())
		
		lis, _ := net.Listen("tcp", grpcAddr)
		go rpcServer.Serve(lis)
	}
	
	{
		// grpc gateway server
		mux := runtime.NewServeMux()
		ctx, _ := context.WithCancel(context.Background())
		opts := []grpc.DialOption{grpc.WithInsecure()}
		
		// 注册反向代理服务
		err := gw.RegisterSayHiHandlerFromEndpoint(ctx, mux, grpcAddr, opts)
		if err != nil {
			fmt.Println("RegisterSayHiHandlerFromEndpoint error:", err.Error())
			return
		}
		// 启动反向代理服务
		http.ListenAndServe(httpAddr, mux)
	}
}
