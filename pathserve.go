package main
import (
	"fmt"
	"path/filepath"
	"os"
	"encoding/json"
	"net/http"
	"time"
)


type Directory struct {

	Files []File

}

type File struct {
	Name string
	Size int64
	
	ModTime time.Time
	IsDir bool
	
}


func files (path string)(Directory) {


	fileList := []File{}

	filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		file := File{Name: path, Size: f.Size(), ModTime: f.ModTime(), IsDir: f.IsDir()}
		fileList = append(fileList, file)
		
		return nil
	})

	return Directory{Files:fileList}
	
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

