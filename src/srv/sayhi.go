package srv

import (
	"context"
	"fmt"
	
	"endpoint/proto"
)

func NewSayHiSrv() *SayHiSrv {
	return &SayHiSrv{}
}

type SayHiSrv struct {
}

// 服务实现接口
func (s *SayHiSrv) SayHi(c context.Context, req *endpoint.HiReq) (*endpoint.HiResp, error) {
	fmt.Println("request name:", req.Name)
	msg := fmt.Sprintf("%s nice to meet u.", req.Name)
	resp := &endpoint.HiResp{Echo: msg}
	fmt.Println(msg)
	return resp, nil
}
