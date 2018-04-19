package singleflight

import (
	"testing"
	"fmt"
	"errors"
	"go.uber.org/atomic"
	"time"
	"sync"
)

func TestGroup_Do(t *testing.T) {
	var g Group
	v, err := g.Do("key", func() (val interface{}, err error) {
		return "value", nil
	})
	if got, want := fmt.Sprintf("%v (%T)", v, v), "value (string)"; got != want {
		t.Errorf("Do = %v; want %v", got, want)
	}
	if err != nil {
		t.Errorf("Do error = %v", err)
	}
}

func TestGroup_Do2(t *testing.T) {
	var someErr = errors.New("some error")
	var g Group
	v,  err := g.Do("key", func() (val interface{}, err error) {
		return nil, someErr
	})
	if err != someErr {
		t.Errorf("Do error = %v; want someErr", err)
	}
	if v != nil {
		t.Errorf("unexpected non-nil value %#v", v)
	}
}

func TestGroup_Do3(t *testing.T) {
	var g Group
	var times atomic.Int32

	ch := make(chan string)
	fn := func() (interface{}, error){
		times.Add(1)
		time.Sleep(2*time.Second)
		return <-ch, nil
	}

	var wg sync.WaitGroup
	for i:=0; i<10 ; i++ {
		wg.Add(1)
		go func() {
			v, e := g.Do("key", fn)
			if e != nil {
				t.Errorf("Do error: %v", v)
			}
			if v.(string) != "done" {
				t.Errorf("got %q, want %s", v, "done")
			}
			wg.Done()
		}()
	}
	time.Sleep(1*time.Second)
	ch <- "done"
	wg.Wait()
	if got := times.Load(); got != 1 {
		t.Errorf("number of calls = %d; want 1", got)
	}
}
