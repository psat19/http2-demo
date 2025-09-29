package main
import (
	"fmt"
	"net/http"
)
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Server 1!")
}
func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server 1 running on :8081")
	http.ListenAndServe(":8081", nil)
}
