package lasttime

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func ExpBufio(count int, usebufio bool) {
	var writer io.Writer
	const buffsize = 65536

	if usebufio {
		bufWriter := bufio.NewWriterSize(os.Stdout, buffsize)
		defer bufWriter.Flush()
		writer = bufWriter
	} else {
		writer = os.Stdout
	}
	for i := 0; i < count; i++ {
		fmt.Fprintln(writer, strings.Repeat("x", i+1))
	}
}

func ExpGetURL(urlStr string) (io.ReadCloser, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	switch u.Scheme {
	case "http", "https":
		return getHTTP(u)
	case "file":
		return getFile(u)
	default:
		return nil, fmt.Errorf("URL scheme must be http,https or file %s", urlStr)
	}
}

func getFile(u *url.URL) (io.ReadCloser, error) {
	return os.Open(u.Path)
}

func getHTTP(u *url.URL) (io.ReadCloser, error) {
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "GoSampleProgram/003")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func ExpSaveFile(reader io.ReadCloser) error {
	tmp, err := ioutil.TempFile(os.TempDir(), "tmp")
	if err != nil {
		return err
	}

	hash := sha256.New()
	w := io.MultiWriter(tmp, hash)

	written, err := io.Copy(w, reader)
	if err != nil {
		return err
	}

	fmt.Printf("Wrote %d bytes to %s\nSHA256:%x\n",
		written,
		tmp.Name(),
		hash.Sum(nil))
	return nil
}
