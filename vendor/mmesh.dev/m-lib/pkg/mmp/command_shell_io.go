package mmp

import (
	"context"
	"io"
	"time"

	"mmesh.dev/m-api-go/grpc/common/status"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-api-go/grpc/network/mmsp/command"
	"mmesh.dev/m-lib/pkg/mmp/streaming"
	"x6a.dev/pkg/errors"
	"x6a.dev/pkg/xlog"
)

func newShellOutput(srcID string, p *mmsp.Payload) *mmsp.Payload {
	cReq := p.Command.CommandRequest

	return &mmsp.Payload{
		SrcID:       srcID,
		DstID:       p.SrcID,
		PayloadType: mmsp.PayloadType_COMMAND_SHELL_OUTPUT,
		Command: &command.Command{
			CommandResponse: &command.CommandResponse{
				RequestedCommand: cReq,
				Result:           &command.CommandResult{},
				Stdout:           nil,
				Stderr:           nil,
				Status: &status.StatusResponse{
					SourceID:  srcID,
					Code:      0,
					Message:   "",
					Timestamp: time.Now().Unix(),
				},
			},
		},
	}
}

func shellWriteOutput(mmID string, payload *mmsp.Payload, outrp *io.PipeReader) {
	// t1 := time.Now()
	for {
		p := newShellOutput(mmID, payload)
		buffer := make([]byte, streaming.BufferLen)
		n, err := outrp.Read(buffer)
		if err == io.EOF {
			// xlog.Tracef("OUT Pipe Reader EOF: %v", err)
			outrp.Close()
			break
		}
		if err != nil {
			// xlog.Tracef("OUT Pipe Reader: %v", err)
			outrp.Close()
			break
		}

		p.Command.CommandResponse.Stdout = buffer[0:n]
		//p.Command.CommandResponse.Stderr = data
		p.Command.CommandResponse.Result.Status = command.CommandResultStatus_RUNNING
		// payload.CommandResponse.Result.Duration = int64(time.Since(t1).Seconds())

		TxControlQueue <- p
	}
}

func shellReadInput(ctx context.Context, payload *mmsp.Payload) error {
	if payload.Command == nil {
		return nil
	}

	if payload.Command.CommandRequest == nil {
		return nil
	}

	if len(payload.Command.CommandRequest.Stdin) == 0 {
		xlog.Trace("len(payload.Command.CommandRequest.Stdin) == 0")
		return nil
	}

	// logging.Tracef("Processing data from node %s", payload.SrcID)

	sID := payload.SrcID

	iop := shs.GetIOSessionIO(sID)
	if iop == nil {
		return errors.Errorf("shell io pipes not found for session from %s", sID)
	}
	if iop.In.WP == nil {
		return errors.Errorf("shell input writer pipe not found for session from %s", sID)
	}

	n, err := iop.In.WP.Write(payload.Command.CommandRequest.Stdin)
	if err != nil {
		// xlog.Errorf("Unable to read stdin data from node %s: %v", payload.SrcID, err)
		return errors.Wrapf(err, "[%v] function iop.in.wp.Write(payload.Command.CommandRequest.Stdin)", errors.Trace())
	}
	xlog.Tracef("Received %d bytes from node %s stdin", n, payload.SrcID)

	return nil
}
