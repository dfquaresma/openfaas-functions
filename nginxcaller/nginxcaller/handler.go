package function

import (
	"fmt"
	"net/http"
)

// Handle a serverless request
func Handle(req []byte) string {
	resp, err := http.Get("http://127.0.0.1:80")
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("Nginx request status: %s", resp.Status)
}
