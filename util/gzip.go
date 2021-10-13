package util

import (
	"compress/gzip"
	"io"
	"net/http"
)

type gzreadCloser struct {
	*gzip.Reader
	io.Closer
}

func (gz gzreadCloser) Close() error {
	return gz.Closer.Close()
}

func GzipDecode(resp *http.Response) error {
	if resp.Header.Get("Content-Encoding") == "gzip" {
		resp.Header.Del("Content-Length")
		zr, err := gzip.NewReader(resp.Body)
		if err != nil {
			return err
		}
		resp.Body = gzreadCloser{zr, resp.Body}
	}
	return nil
}
