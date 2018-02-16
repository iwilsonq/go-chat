package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path"
)

func uploaderHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.FormValue("userid")
	fmt.Println("userID", userID)
	file, header, err := r.FormFile("avatarFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filename := path.Join("avatars", userID+path.Ext(header.Filename))
	err = ioutil.WriteFile(filename, data, 0777)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	io.WriteString(w, "Successful")
}
