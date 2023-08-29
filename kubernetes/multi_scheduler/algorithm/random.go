package algorithm

import (
	"math/rand"

	corev1 "k8s.io/api/core/v1"
)

func RandomPredicate(node *corev1.Node, pod *corev1.Pod) bool {
	r := rand.Intn(2)
	return r == 0
}

func RandomPriority(node *corev1.Node, pod *corev1.Pod) int {
	return rand.Intn(100)
}
