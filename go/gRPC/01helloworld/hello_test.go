package helloworld

import (
	"github.com/zhengyansheng/helloworld/service"
	"google.golang.org/protobuf/proto"
	"testing"
)

func TestHello(t *testing.T) {
	u := &service.UserRequest{
		Username: "zhengyansheng",
		Age:      18,
	}

	marshal, err := proto.Marshal(u)
	if err != nil {
		panic(err)
	}
	t.Logf("pb marshal: %v", marshal)

	var n service.UserRequest
	err = proto.Unmarshal(marshal, &n)
	if err != nil {
		panic(err)
	}
	t.Logf("User: %+v", n)
	t.Logf("User: %v", n.String())
}
