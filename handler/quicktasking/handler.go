package quicktasking

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/eagle/handler/loading"
	"github.com/eagle/handler/quicktask"
	"github.com/eagle/handler/task_manager"
)

// http://localhost:3000/quicktask?site=thebrokenarm&msku=567V&size=random
func quicktaskHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/quicktask" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	msku := r.URL.Query().Get("msku")
	if msku == "" {
		http.Error(w, "msku not provided.", http.StatusNotFound)
		return
	}

	site := r.URL.Query().Get("site")
	if site == "" {
		http.Error(w, "site not provided.", http.StatusNotFound)
		return
	}

	size, _ := url.QueryUnescape(r.URL.Query().Get("size"))
	if size == "" {
		http.Error(w, "size not provided.", http.StatusNotFound)
		return
	}

	for _, taskUUID := range loading.Data.Quicktask.Quicktask[site] {
		taskObject, err := quicktask.GetQuicktask(taskUUID)

		if err != nil {
			fmt.Println("Failed to get task: ", err.Error())
			continue
		}

		taskObject.Size = size
		taskObject.Pid = msku

		if !taskObject.Active {
			go task_manager.RunQuickTask(taskObject)
		} else if taskObject.Done {
			task_manager.StopQuickTask(taskObject)
		}

	}

	_, err := http.ResponseWriter(w).Write([]byte("Quicktask started!"))
	if err != nil {
		return
	}

}
