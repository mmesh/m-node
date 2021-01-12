package mmp

import (
	"io"

	"mmesh.dev/m-lib/pkg/logging"
)

type ioPipe struct {
	rp *io.PipeReader
	wp *io.PipeWriter
}

type ioPipes struct {
	in  *ioPipe
	out *ioPipe
}

func newIOPipes() *ioPipes {
	inReaderPipe, inWriterPipe := io.Pipe()
	outReaderPipe, outWriterPipe := io.Pipe()

	return &ioPipes{
		in: &ioPipe{
			rp: inReaderPipe,
			wp: inWriterPipe,
		},
		out: &ioPipe{
			rp: outReaderPipe,
			wp: outWriterPipe,
		},
	}
}

func closeIOPipes(iop *ioPipes) {
	closeIORPipe(iop.in.rp)
	closeIOWPipe(iop.in.wp)
	closeIORPipe(iop.out.rp)
	closeIOWPipe(iop.out.wp)
}

func closeIORPipe(p *io.PipeReader) {
	if p == nil {
		return
	}
	if err := p.Close(); err != nil {
		logging.Errorf("Closing IO PipeReader: %v", err)
	}
}

func closeIOWPipe(p *io.PipeWriter) {
	if p == nil {
		return
	}
	if err := p.Close(); err != nil {
		logging.Errorf("Closing IO PipeWriter: %v", err)
	}
}
