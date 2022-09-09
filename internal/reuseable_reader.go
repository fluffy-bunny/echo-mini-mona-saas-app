package internal

import (
	"bytes"
	"io"
)

// https://blog.flexicondev.com/read-go-http-request-body-multiple-times

type reusableReader struct {
	io.Reader
	readBuf *bytes.Buffer
	backBuf *bytes.Buffer
}

func ReusableReader(r io.Reader) (io.Reader, error) {
	readBuf := bytes.Buffer{}
	_, err := readBuf.ReadFrom(r)
	if err != nil {
		return nil, err
	}
	backBuf := bytes.Buffer{}

	return reusableReader{
		io.TeeReader(&readBuf, &backBuf),
		&readBuf,
		&backBuf,
	}, nil
}

func (r reusableReader) Read(p []byte) (int, error) {
	n, err := r.Reader.Read(p)
	if err == io.EOF {
		r.reset()
	}
	return n, err
}

func (r reusableReader) reset() {
	io.Copy(r.readBuf, r.backBuf) // nolint: errcheck
}
