package ops

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/common/status"
	"mmesh.dev/m-api-go/grpc/network/mmsp/command"
	"mmesh.dev/m-api-go/grpc/resources/ops/workflow"
	"x6a.dev/pkg/xlog"
)

func runWorkflowAction(wf *workflow.Workflow, a *workflow.Action) *workflow.Operation {
	//var cmdOut, cmdErr bytes.Buffer
	var statusMsg string
	var statusCode int32
	var resultStatus command.CommandResultStatus

	if a.Command == nil {
		return nil
	}

	xlog.Infof("Executing workflow %s, action %s", wf.WorkflowID, a.Name)

	mmID := viper.GetString("mm.id")

	c := a.Command

	// execute the command
	cmd := exec.Command(c.Cmd, c.Args...)
	//cmd.Stderr = &cmdErr
	//cmd.Stdout = &cmdOut
	cmd.Stdin = nil

	t1 := time.Now()
	//if err := cmd.Run(); err != nil {
	out, err := cmd.CombinedOutput()
	if err != nil {
		statusCode = -1
		statusMsg = err.Error()
		resultStatus = command.CommandResultStatus_FAILED

		xlog.Errorf("Unable to run command %s: %v", c.Cmd, err)
	} else {
		statusCode = 0
		statusMsg = "OK"
		resultStatus = command.CommandResultStatus_EXECUTED
	}

	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)

	return &workflow.Operation{
		AccountID:    wf.AccountID,
		AccountToken: wf.AccountToken,
		ProjectID:    wf.ProjectID,
		WorkflowID:   wf.WorkflowID,
		OperationID:  fmt.Sprintf("%s:%d@%s", a.Name, timestamp, mmID),
		Description:  fmt.Sprintf("Action %s executed on %s at %s", a.Name, mmID, tm.String()),
		NodeID:       mmID,
		Status: &status.StatusResponse{
			SourceID:  mmID,
			Code:      statusCode,
			Message:   statusMsg,
			Timestamp: timestamp,
		},
		Result: &command.CommandResult{
			Status:   resultStatus,
			Duration: int64(time.Since(t1).Seconds()),
		},
		//Stdout: cmdOut.Bytes(),
		//Stderr: cmdErr.Bytes(),
		StdoutStderr: out,
	}
}