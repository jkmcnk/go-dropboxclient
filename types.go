package dropboxclient

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type DropboxFile struct {
	Bytes    int64  `json:"bytes"`
	Modified string `json:"modified"`
	Reader   io.ReadCloser
}

func DropboxFileFromHeaders(path string, headers http.Header) (file *DropboxFile) {
	metadata := headers.Get("x-dropbox-metadata")

	file = &DropboxFile{}

	_ = json.Unmarshal([]byte(metadata), file)

	contentLength, _ := strconv.ParseInt(headers.Get("Content-Length"), 10, 0)

	file.Bytes = contentLength

	return
}