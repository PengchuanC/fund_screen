package rpc

import (
	"fmt"
	"fund_screen/common/settings"
	"fund_screen/services"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
	"net"
)

type ScreenRpcServer struct {}

var db, tx *gorm.DB


func Serve() {
	config := settings.GetSettings()
	rpc := config.Rpc
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", rpc.Port))
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	services.RegisterScreenRpcServerServer(server, &ScreenRpcServer{})
	reflection.Register(server)
	register()
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}


func register() {
	config := settings.GetSettings()
	consul := config.Consul
	rpc := config.Rpc
	var consulConfig = api.Config{Address: fmt.Sprintf("%s:%d", consul.Host, consul.Port)}
	client, err := api.NewClient(&consulConfig)
	if err != nil {
		panic(err)
	}

	reg := &api.AgentServiceRegistration{
		ID: rpc.Name,
		Name: rpc.Name,
		Port: rpc.Port,
		Address: rpc.Host,
		Check: &api.AgentServiceCheck{
			Interval: "10s",
			Timeout: "60s",
			DeregisterCriticalServiceAfter: "60s", // 注销时间，相当于过期时间
			TCP: fmt.Sprintf("%s:%d", rpc.Host, rpc.Port),
		},
	}
	if err := client.Agent().ServiceRegister(reg); err != nil {
		panic(err)
	}
	fmt.Printf("RPC服务 %s@%s 已注册\n", rpc.Name, fmt.Sprintf("%s:%d", rpc.Host, rpc.Port))
}