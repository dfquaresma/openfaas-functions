package function

import (
	"fmt"
	"net/http"
)

// Handle a serverless request
func Handle(req []byte) string {
	resp, _ := http.Get("http://127.0.0.1")
	return fmt.Sprintf("Nginx request status: %s", resp.Status)
}
