package streaming

import (
	"io"

	"mmesh.dev/m-lib/pkg/logging"
)

type IOPipe struct {
	RP *io.PipeReader
	WP *io.PipeWriter
}

type IOPipes struct {
	In  *IOPipe
	Out *IOPipe
}

func NewIOPipes() *IOPipes {
	inReaderPipe, inWriterPipe := io.Pipe()
	outReaderPipe, outWriterPipe := io.Pipe()

	return &IOPipes{
		In: &IOPipe{
			RP: inReaderPipe,
			WP: inWriterPipe,
		},
		Out: &IOPipe{
			RP: outReaderPipe,
			WP: outWriterPipe,
		},
	}
}

func CloseIOPipes(iop *IOPipes) {
	closeIORPipe(iop.In.RP)
	closeIOWPipe(iop.In.WP)
	closeIORPipe(iop.Out.RP)
	closeIOWPipe(iop.Out.WP)
}

func closeIORPipe(p *io.PipeReader) {
	if p == nil {
		return
	}
	if err := p.Close(); err != nil {
		logging.Errorf("Unable to close IO PipeReader: %v", err)
	}
}

func closeIOWPipe(p *io.PipeWriter) {
	if p == nil {
		return
	}
	if err := p.Close(); err != nil {
		logging.Errorf("Unable to close IO PipeWriter: %v", err)
	}
}
