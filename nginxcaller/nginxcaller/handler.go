package function

import (
	"fmt"
	"net/http"
)

// Handle a serverless request
func Handle(req []byte) string {
	resp, _ := http.Get("http://127.0.0.1:80")
	return fmt.Sprintf("Outcaller status: %s", resp.Status)
}
