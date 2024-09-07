package retry

import (
	"testing"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/util/retry"
)

// TestRetryOnConflict 冲突重试
func TestRetryOnConflict(t *testing.T) {
	// 如果匿名函数返回nil，则退出；如果匿名函数返回errors.NewConflict 则继续重试
	err := retry.RetryOnConflict(retry.DefaultBackoff, func() (err error) {
		t.Log("hello retry")
		//return nil
		return errors.NewConflict(schema.GroupResource{Resource: "Deployments"}, "appset-monkey", nil)
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Done")
}
