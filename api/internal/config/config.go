package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	System  zrpc.RpcClientConf
	Achieve zrpc.RpcClientConf
	Storage zrpc.RpcClientConf
	JwtConf struct {
		SecretKey string
		Expire    string
		Buffer    string
		Issuer    string
		Audience  string
	}
}

var GlobalConfig Config
