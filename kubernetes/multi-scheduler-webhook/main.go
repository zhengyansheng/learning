package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhengyansheng/common"
	"k8s.io/klog/v2"
	extenderv1 "k8s.io/kube-scheduler/extender/v1"
)

const (
	apiPrefix = "/scheduler/extender"
)

type Predicate struct {
}

func NewPredicate() *Predicate {
	return &Predicate{}
}

func (p *Predicate) Filter(args extenderv1.ExtenderArgs) *extenderv1.ExtenderFilterResult {
	pod := args.Pod
	if pod == nil {
		return &extenderv1.ExtenderFilterResult{Error: fmt.Sprintf("pod is nil")}
	}
	scheduleNodes := make([]string, 0)
	failedNodes := make(map[string]string)
	for _, nodeName := range *args.NodeNames {
		scheduleNodes = append(scheduleNodes, nodeName)
	}

	klog.Infof("filter localstorage pods on nodes: %v", scheduleNodes)
	return &extenderv1.ExtenderFilterResult{
		NodeNames:   &scheduleNodes,
		Nodes:       args.Nodes,
		FailedNodes: failedNodes,
	}
}

type Prioritize struct {
}

func NewPrioritize() *Prioritize {
	return &Prioritize{}
}

func (p *Prioritize) Score(args extenderv1.ExtenderArgs) *extenderv1.HostPriorityList {
	pod := args.Pod
	if pod == nil {
		klog.Errorf("pod is nil")
		return nil
	}

	nodeNames := *args.NodeNames
	klog.Infof("scoring nodes %v", nodeNames)

	hostPriorityList := make(extenderv1.HostPriorityList, len(nodeNames))

	for i, nodeName := range nodeNames {
		hostPriorityList[i] = extenderv1.HostPriority{
			Host:  nodeName,
			Score: int64(rand.Intn(10)),
		}
	}

	klog.Infof("score localstorage pods on nodes: %v", hostPriorityList)
	return &hostPriorityList
}

func main() {
	p1 := NewPredicate()
	p2 := NewPrioritize()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST(apiPrefix+"/filter", func(c *gin.Context) {
		var (
			extenderArgs         extenderv1.ExtenderArgs
			extenderFilterResult *extenderv1.ExtenderFilterResult
		)
		if err := c.ShouldBindJSON(&extenderArgs); err != nil {
			extenderFilterResult = &extenderv1.ExtenderFilterResult{Error: err.Error()}
		} else {
			extenderFilterResult = p1.Filter(extenderArgs)
		}
		common.Indent(extenderFilterResult)
		c.Header("Content-Type", "application/json")
		result, err := json.Marshal(extenderFilterResult)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.String(http.StatusOK, string(result))
		return

	})

	r.POST(apiPrefix+"/prioritize", func(c *gin.Context) {
		var (
			extenderArgs     extenderv1.ExtenderArgs
			hostPriorityList *extenderv1.HostPriorityList
		)
		if err := c.ShouldBindJSON(&extenderArgs); err != nil {
			hostPriorityList = &extenderv1.HostPriorityList{}
		} else {
			hostPriorityList = p2.Score(extenderArgs)
		}
		common.Indent(hostPriorityList)
		c.Header("Content-Type", "application/json")
		result, err := json.Marshal(hostPriorityList)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.String(http.StatusOK, string(result))
		return
	})

	r.Run("0.0.0.0:8000")
}
