package function

import (
	"fmt"
	"net/http"
)

// Handle a serverless request
func Handle(req []byte) string {
	resp, err := http.Get("http://google.com/")
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	return fmt.Sprintf("Outcaller status: %s", resp.Status)
}
