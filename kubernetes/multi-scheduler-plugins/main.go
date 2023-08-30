package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/zhengyansheng/multi_scheduler_plugins/pkg/plugins"
	_ "k8s.io/component-base/logs/json/register" // for JSON log format registration
	_ "k8s.io/component-base/metrics/prometheus/clientgo"
	_ "k8s.io/component-base/metrics/prometheus/version" // for version metric registration
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	command := app.NewSchedulerCommand(
		// 注册插件
		app.WithPlugin(plugins.Name, plugins.New),
	)

	if err := command.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

}
