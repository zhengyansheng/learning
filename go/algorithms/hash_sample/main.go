package main

import (
	"fmt"

	"github.com/stathat/consistent"
)

func main() {
	// 创建一致性哈希实例
	hashRing := consistent.New()

	// 添加节点
	hashRing.Add("Node1")
	hashRing.Add("Node2")
	hashRing.Add("Node3")
	hashRing.Add("Node4")

	// 根据数据的 key 获取对应的节点
	key := "some_key"
	node, err := hashRing.Get(key)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Key '%s' maps to node '%s'\n", key, node)

	// 移除节点
	hashRing.Remove("Node2")

	// 再次根据相同的 key 获取对应的节点
	node, err = hashRing.Get(key)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("After removing Node2, key '%s' maps to node '%s'\n", key, node)
}
