package jobs

import (
	"github.com/golang/protobuf/proto"

	"github.com/totoval/framework/helpers/debug"
	"github.com/totoval/framework/job"
	test "totoval/app/jobs/protocol_buffers"
)

func init() {
	job.Add(&ExampleJob{})
}

type ExampleJob struct {
	job.Job
}

func (e *ExampleJob) Retries() uint32 {
	return 3
}

func (e *ExampleJob) Name() string {
	return "example_job"
}

func (e *ExampleJob) ParamProto() proto.Message {
	return &test.ExampleJob{}
}

func (e *ExampleJob) Handle(paramPtr proto.Message) error {
	obj := paramPtr.(*test.ExampleJob)
	debug.Dump(obj)
	return nil
}
