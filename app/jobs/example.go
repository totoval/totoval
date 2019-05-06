package jobs

import (
	"github.com/golang/protobuf/proto"

	"github.com/totoval/framework/helpers/debug"
	"github.com/totoval/framework/job"
	pbs "totoval/app/jobs/protocol_buffers"
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
	return "example-job"
}

func (e *ExampleJob) ParamProto() proto.Message {
	return &pbs.ExampleJob{}
}

func (e *ExampleJob) Handle(paramPtr proto.Message) error {
	obj := paramPtr.(*pbs.ExampleJob)
	debug.Dump(obj)
	return nil
}
