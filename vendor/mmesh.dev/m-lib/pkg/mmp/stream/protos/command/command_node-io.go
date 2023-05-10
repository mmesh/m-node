package command

import (
	"context"
	"io"
	"time"

	"mmesh.dev/m-api-go/grpc/common/status"
	"mmesh.dev/m-api-go/grpc/network/mmsp"
	"mmesh.dev/m-api-go/grpc/network/mmsp/stream"
	"mmesh.dev/m-api-go/grpc/resources/ops"
	"mmesh.dev/m-lib/pkg/errors"
	"mmesh.dev/m-lib/pkg/mmp/stream/utils/streaming"
	"mmesh.dev/m-lib/pkg/mmp/queuing"
	"mmesh.dev/m-lib/pkg/xlog"
)

func newShellOutput(srcID string, p *mmsp.Payload) *mmsp.Payload {
	cReq := p.CommandPDU.Command.Request

	return &mmsp.Payload{
		SrcID: srcID,
		DstID: p.SrcID,
		Type: mmsp.PDUType_COMMAND,
		CommandPDU: &mmsp.CommandPDU{
			Type: mmsp.CommandMsgType_COMMAND_OUTPUT,
			Command: &stream.Command{
				Response: &stream.CommandResponse{
					RequestedCommand: cReq,
					Result:           &ops.CommandResult{},
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
		},
	}

	// return &mmsp.Payload{
	// 	SrcID:       srcID,
	// 	DstID:       p.SrcID,
	// 	PayloadType: mmsp.PayloadType_COMMAND_SHELL_OUTPUT,
	// 	Command: &command.Command{
	// 		CommandResponse: &command.CommandResponse{
	// 			RequestedCommand: cReq,
	// 			Result:           &command.CommandResult{},
	// 			Stdout:           nil,
	// 			Stderr:           nil,
	// 			Status: &status.StatusResponse{
	// 				SourceID:  srcID,
	// 				Code:      0,
	// 				Message:   "",
	// 				Timestamp: time.Now().Unix(),
	// 			},
	// 		},
	// 	},
	// }
}

func shellWriteOutput(mmID string, payload *mmsp.Payload, outrp *io.PipeReader) {
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

		p.CommandPDU.Command.Response.Stdout = buffer[0:n]
		// p.CommandPDU.Command.Response.Stderr = data
		p.CommandPDU.Command.Response.Result.Status = ops.CommandResultStatus_RUNNING

		queuing.TxControlQueue <- p
	}
}

func shellReadInput(ctx context.Context, p *mmsp.Payload) error {
	if p.CommandPDU.Command == nil {
		return nil
	}

	if p.CommandPDU.Command.Request == nil {
		return nil
	}

	if len(p.CommandPDU.Command.Request.Stdin) == 0 {
		xlog.Trace("len(p.CommandPDU.Command.Request.Stdin) == 0")
		return nil
	}

	// logging.Tracef("Processing data from node %s", p.CommandPDU.Metadata.SrcID)

	sID := p.SrcID

	iop := shs.GetIOSessionIO(sID)
	if iop == nil {
		return errors.Errorf("shell io pipes not found for session from %s", sID)
	}
	if iop.In.WP == nil {
		return errors.Errorf("shell input writer pipe not found for session from %s", sID)
	}

	n, err := iop.In.WP.Write(p.CommandPDU.Command.Request.Stdin)
	if err != nil {
		// xlog.Errorf("Unable to read stdin data from node %s: %v", sID, err)
		return errors.Wrapf(err, "[%v] function iop.in.wp.Write()", errors.Trace())
	}
	xlog.Tracef("Received %d bytes from node %s stdin", n, sID)

	return nil
}
