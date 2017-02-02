package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Welcome!")
}

func GetDocument(w http.ResponseWriter, r *http.Request) {

	type Configuration struct {
		COLLECTORS_PATH []string
	}

	// TODO reading of config should happen on instatiation of webserver
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Fatal("error:", err)
	}

	vars := mux.Vars(r)
	ediParam := vars["ediParam"]
	matParam := vars["matParam"]

	if len(configuration.COLLECTORS_PATH) == 0 {
		log.Fatal("error: varible COLLECTORS_PATH not set")
	}
	python_path := configuration.COLLECTORS_PATH[0]
	//script := python_path + "getMateria.py"
	version := python_path + "getVersion.py"
	//arg1 := fmt.Sprintf("--edition=%v", ediParam)
	//arg2 := fmt.Sprintf("--document=%v", matParam)
	//out, err := exec.Command("C:/Anaconda2/python", script, arg1, arg2).Output()
	out, err := exec.Command("C:/Anaconda2/python", version).Output()
	if err != nil {
		//TODO Don't just give up here. Return some meaningful error message and carry on
		log.Fatal(err, "Error running script:", version)
	}

	output := fmt.Sprintf("%q", out)
	materias := Materias{
		Materia{Edition: ediParam, Document: matParam, Text: output},
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(materias); err != nil {
		panic(err)
	}
}
