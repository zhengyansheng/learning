package app

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"github.com/zhengyansheng/sample-kubelet/app/client"
	"github.com/zhengyansheng/sample-kubelet/app/informer"
	"github.com/zhengyansheng/sample-kubelet/app/node"
	"github.com/zhengyansheng/sample-kubelet/app/node/lease"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog/v2"
)

func NewKubeletCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "kubelet",
		RunE: func(cmd *cobra.Command, args []string) error {

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			nodeName := "my-kind-worker3"
			// 1.client客户端
			clientSet := client.InitClient(filepath.Join(os.Getenv("HOME"), ".kube", "config"))

			// 2.注册node
			registeredNode := node.RegisterNode(clientSet, nodeName)

			// 3.定时上报node信息
			go wait.Until(func() {
				err := node.ReportNodeStatus(clientSet, registeredNode)
				if err != nil {
					panic(err)
				}
			}, time.Minute*5, wait.NeverStop)

			// 4.启动lease租约控制器
			lease.StartLeaseController(clientSet, nodeName)

			// 5.启动informer
			informer.InitInformer(clientSet, nodeName)

			klog.Infoln("start sample kubelet...")

			select {
			case <-ctx.Done():
				break
			case <-wait.NeverStop:
				break
			}
			return nil
		},
	}
	return cmd
}
