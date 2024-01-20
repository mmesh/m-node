package hsec

import (
	"bytes"
	"encoding/gob"
	"os"

	"mmesh.dev/m-api-go/grpc/resources/hsec"
	"mmesh.dev/m-lib/pkg/errors"
)

func writeReportFile(r *hsec.Report) error {
	f, err := os.Create(reportFile())
	if err != nil {
		return errors.Wrapf(err, "[%v] function os.Create()", errors.Trace())
	}
	defer f.Close()

	var buf bytes.Buffer

	if err := gob.NewEncoder(&buf).Encode(r); err != nil {
		return errors.Wrapf(err, "[%v] function gob.NewEncoder()", errors.Trace())
	}

	// if err := binary.Write(&buf, binary.BigEndian, r); err != nil {
	// 	return errors.Wrapf(err, "[%v] function binary.Write()", errors.Trace())
	// }

	if _, err := f.Write(buf.Bytes()); err != nil {
		return errors.Wrapf(err, "[%v] function f.Write()", errors.Trace())
	}

	return nil
}

func readReportFile() (*hsec.Report, error) {
	data, err := os.ReadFile(reportFile())
	if err != nil {
		return nil, errors.Wrapf(err, "[%v] function os.ReadFile()", errors.Trace())
	}

	buffer := bytes.NewBuffer(data)
	r := hsec.Report{}

	if err := gob.NewDecoder(buffer).Decode(&r); err != nil {
		return nil, errors.Wrapf(err, "[%v] function gob.NewDecoder()", errors.Trace())
	}

	// if err := binary.Read(buffer, binary.BigEndian, &r); err != nil {
	// 	return nil, errors.Wrapf(err, "[%v] function binary.Read()", errors.Trace())
	// }

	return &r, nil
}
