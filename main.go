package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

//GIT自动拉取服务
func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Query().Get("path")
		log.Println("拉取", path)

		if path == "" {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		cmd := fmt.Sprintf("cd %s && git pull", path)
		out, err := exec.Command("bash", "-c", cmd).CombinedOutput()
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			_, _ = writer.Write([]byte(err.Error()))
			return
		}
		result := string(out)
		log.Println(result)

		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write(out)
	})

	log.Fatalln(http.ListenAndServe(":81", nil))
}
