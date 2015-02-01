package main
import (
	"fmt"
	"path/filepath"
	"os"
	"encoding/json"
	"net/http"
	"time"
)


type Files struct {

	Files []File

}

type File struct {
	Name string
	Size int64
	
	ModTime time.Time
	IsDir bool
	
}


func files (path string)(Files) {


	fileList := []File{}

	filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		 if err != nil {
                      return err
               }

		file := File{Name: path, Size: f.Size(), ModTime: f.ModTime(), IsDir: f.IsDir()}
		fileList = append(fileList, file)
		
		return nil
	})

	return Files{Files:fileList}
	
}

func handler(w http.ResponseWriter, r *http.Request) {

	remPartOfURL := r.URL.Path[len("/path"):]

	js, err := json.Marshal(files(remPartOfURL))



	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	fmt.Println("Server Running")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}

