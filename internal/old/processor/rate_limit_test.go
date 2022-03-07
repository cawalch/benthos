package processor

import (
	"context"
	"errors"
	"reflect"
	"sync/atomic"
	"testing"
	"time"

	"github.com/Jeffail/benthos/v3/internal/component/metrics"
	"github.com/Jeffail/benthos/v3/internal/log"
	"github.com/Jeffail/benthos/v3/internal/manager/mock"
	"github.com/Jeffail/benthos/v3/internal/message"
)

func TestRateLimitBasic(t *testing.T) {
	var hits int32
	rlFn := func(context.Context) (time.Duration, error) {
		atomic.AddInt32(&hits, 1)
		return 0, nil
	}

	mgr := mock.NewManager()
	mgr.RateLimits["foo"] = rlFn

	conf := NewConfig()
	conf.Type = "rate_limit"
	conf.RateLimit.Resource = "foo"
	proc, err := New(conf, mgr, log.Noop(), metrics.Noop())
	if err != nil {
		t.Fatal(err)
	}

	input := message.QuickBatch([][]byte{
		[]byte(`{"key":"1","value":"foo 1"}`),
		[]byte(`{"key":"2","value":"foo 2"}`),
		[]byte(`{"key":"1","value":"foo 3"}`),
	})

	output, res := proc.ProcessMessage(input)
	if res != nil {
		t.Fatal(res)
	}

	if len(output) != 1 {
		t.Fatalf("Wrong count of result messages: %v", len(output))
	}

	if exp, act := message.GetAllBytes(input), message.GetAllBytes(output[0]); !reflect.DeepEqual(exp, act) {
		t.Errorf("Wrong result messages: %s != %s", act, exp)
	}

	if exp, act := int32(3), atomic.LoadInt32(&hits); exp != act {
		t.Errorf("Wrong count of rate limit hits: %v != %v", act, exp)
	}
}

func TestRateLimitErroredOut(t *testing.T) {
	rlFn := func(context.Context) (time.Duration, error) {
		return 0, errors.New("omg foo")
	}

	mgr := mock.NewManager()
	mgr.RateLimits["foo"] = rlFn

	conf := NewConfig()
	conf.Type = "rate_limit"
	conf.RateLimit.Resource = "foo"
	proc, err := New(conf, mgr, log.Noop(), metrics.Noop())
	if err != nil {
		t.Fatal(err)
	}

	input := message.QuickBatch([][]byte{
		[]byte(`{"key":"1","value":"foo 1"}`),
		[]byte(`{"key":"2","value":"foo 2"}`),
		[]byte(`{"key":"1","value":"foo 3"}`),
	})

	closedChan := make(chan struct{})
	go func() {
		output, res := proc.ProcessMessage(input)
		if res != nil {
			t.Error(res)
		}

		if len(output) != 1 {
			t.Errorf("Wrong count of result messages: %v", len(output))
		}

		if exp, act := message.GetAllBytes(input), message.GetAllBytes(output[0]); !reflect.DeepEqual(exp, act) {
			t.Errorf("Wrong result messages: %s != %s", act, exp)
		}
		close(closedChan)
	}()

	proc.CloseAsync()
	select {
	case <-closedChan:
	case <-time.After(time.Second):
		t.Error("Timed out")
	}
}

func TestRateLimitBlocked(t *testing.T) {
	rlFn := func(context.Context) (time.Duration, error) {
		return time.Second * 10, nil
	}

	mgr := mock.NewManager()
	mgr.RateLimits["foo"] = rlFn

	conf := NewConfig()
	conf.Type = "rate_limit"
	conf.RateLimit.Resource = "foo"
	proc, err := New(conf, mgr, log.Noop(), metrics.Noop())
	if err != nil {
		t.Fatal(err)
	}

	input := message.QuickBatch([][]byte{
		[]byte(`{"key":"1","value":"foo 1"}`),
		[]byte(`{"key":"2","value":"foo 2"}`),
		[]byte(`{"key":"1","value":"foo 3"}`),
	})

	closedChan := make(chan struct{})
	go func() {
		output, res := proc.ProcessMessage(input)
		if res != nil {
			t.Error(res)
		}

		if len(output) != 1 {
			t.Errorf("Wrong count of result messages: %v", len(output))
		}

		if exp, act := message.GetAllBytes(input), message.GetAllBytes(output[0]); !reflect.DeepEqual(exp, act) {
			t.Errorf("Wrong result messages: %s != %s", act, exp)
		}
		close(closedChan)
	}()

	proc.CloseAsync()
	select {
	case <-closedChan:
	case <-time.After(time.Second):
		t.Error("Timed out")
	}
}