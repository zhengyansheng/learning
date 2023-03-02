package testcommon

import (
	"fmt"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	name := "k8s-ent-tcr-online.tencentcloudcr.com/huohua-online/we-sparkle-bff-server:7afe9627"
	fmt.Println(strings.LastIndex(name, ":"))
	fmt.Println(strings.LastIndex(name, "/"))
	fmt.Println(len("k8s-ent-tcr-online.tencentcloudcr.com/huohua-online/we-sparkle-bff-server"))
	fmt.Println(len("k8s-ent-tcr-online.tencentcloudcr.com/huohua-online"))
}
