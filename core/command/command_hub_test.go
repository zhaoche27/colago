package command

import (
	"context"
	"fmt"
	"github.com/zhaoche27/colago/common/extension"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/zhaoche27/colago/common/dto"
)

type CostTimeIntercept struct {
	StartTime time.Time
}

func (cti *CostTimeIntercept) PreIntercept(ctx context.Context, command dto.Commander) error {
	cti.StartTime = time.Now()
	return nil
}

func (cti *CostTimeIntercept) PostInterceptor(ctx context.Context, command dto.Commander, response *dto.Response) error {
	log.Printf("biz:%s, desc:%s, cost time:%f s", command.BizScenarioInfo(), command.Desc(), time.Now().Sub(cti.StartTime).Seconds())
	return nil
}

type DemoCommand struct {
	*dto.Command
	Operator int64
}

func (dc *DemoCommand) Desc() string {
	return fmt.Sprintf("Operator:%d", dc.Operator)
}

type DemoCommandExecute struct {
}

func (dce *DemoCommandExecute) CommandType() reflect.Type {
	return reflect.TypeOf(&DemoCommand{})
}

func (dce *DemoCommandExecute) Execute(ctx context.Context, command dto.Commander) (*dto.Response, error) {
	return dto.NewResponseOfSuccess("demo"), nil
}

func init() {
	cti := &CostTimeIntercept{}
	CommandHub.AddPreIntercepts(cti)
	CommandHub.AddPostIntercepts(cti)
	CommandHub.PutCommandExecutor(&DemoCommandExecute{})
}

func Test_commandHub_Send(t *testing.T) {
	type args struct {
		ctx     context.Context
		command dto.Commander
	}
	demoCommand := &DemoCommand{Operator: 123, Command: &dto.Command{BizScenario: extension.NewBizScenarioWithoutBizId("demoUserCase", "demoScenario")}}
	tests := []struct {
		name         string
		args         args
		wantResponse *dto.Response
	}{
		{name: "demo", args: args{ctx: context.TODO(), command: demoCommand}, wantResponse: dto.NewResponseOfSuccess("demo")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResponse := CommandHub.Send(tt.args.ctx, tt.args.command); !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("commandHub.Send() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}
