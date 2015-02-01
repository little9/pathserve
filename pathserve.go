package main
import (
	"fmt"
	"path/filepath"
	"os"
	"encoding/json"
	"net/http"
)


type Directory struct {

	Paths []string
	
}

func files (path string)(Directory) {

	fileList := []string{}
	
	filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)

		return nil
	})

	return Directory{Paths:fileList}
	
}

func handler(w http.ResponseWriter, r *http.Request) {

	remPartOfURL := r.URL.Path[len("/path"):]

	js, err := json.Marshal(files(remPartOfURL))

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
}

func main() {
	fmt.Println("Server Running")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}

