package domain

import (
	"context"
	"fmt"
	"github.com/zhaoche27/colago/core/event"
	"reflect"
	"testing"
)

type DemoEvent struct {
	name string
}

type DemoEventHandler struct {
}

func (*DemoEventHandler) Execute(ctx context.Context, event event.Event) error {
	fmt.Println(event.(*DemoEvent).name)
	return nil
}

func (*DemoEventHandler) EventType() reflect.Type {
	return reflect.TypeOf(&DemoEvent{})
}

func init() {
	event.Hub.Register(&DemoEventHandler{})
}

func Test_eventPublisher_Publish(t *testing.T) {
	type args struct {
		ctx   context.Context
		event event.Event
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "demoEvent", args: args{ctx: context.TODO(), event: &DemoEvent{name: "demo"}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := EventPublisher.Publish(tt.args.ctx, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("eventPublisher.Publish() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
