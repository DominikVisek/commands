package util

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cheggaaa/pb"
)

// DownloadFile retrieves a file.
func DownloadFile(fp string, url string, progressBar bool) (err error) {
	// Create the file
	out, err := os.Create(fp)
	if err != nil {
		return err
	}
	defer CheckClose(out)

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer CheckClose(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download link %s returned wrong status code: got %v want %v", url, resp.StatusCode, http.StatusOK)
	}
	reader := resp.Body
	if progressBar {

		bar := pb.New(int(resp.ContentLength)).SetUnits(pb.U_BYTES).Prefix(filepath.Base(fp))
		bar.Start()

		// create proxy reader
		reader = bar.NewProxyReader(resp.Body)
		// Writer the body to file
		_, err = io.Copy(out, reader)
		bar.Finish()
	} else {
		_, err = io.Copy(out, reader)
	}

	if err != nil {
		return err
	}

	return nil
}
