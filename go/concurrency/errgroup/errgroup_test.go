package errgroup

import (
	"golang.org/x/sync/errgroup"
	"testing"
	"time"
)

func TestErrGroup(t *testing.T) {
	g := errgroup.Group{}

	g.Go(func() error {
		t.Log("1")
		return nil
	})
	g.Go(func() error {
		t.Log("2")
		return nil
	})
	g.Go(func() error {
		time.Sleep(time.Second)
		t.Log("3")

		panic("error 3")
	})

	g.Go(func() error {
		time.Sleep(time.Second * 2)
		t.Log("4")

		panic("error 4")
	})

	if err := g.Wait(); err != nil {
		t.Logf("wait err: %v", err)
		return
	}
}
