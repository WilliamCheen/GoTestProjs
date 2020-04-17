package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Index!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	if err := r.ParseForm(); err != nil {
		fmt.Fprintln(w, "解析错误", err)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	username := r.PostFormValue("username")

	// body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	// if err != nil {
	// 	panic(err)
	// }

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	fmt.Printf("拿到了用户名===》" + username)
	if len(username) > 0 {
		fmt.Printf("拿到了用户名")
		t := RepoCreateTodoWithName(todo, username)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(t); err != nil {
			panic(err)
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		http.Error(w, "用户名不能为空", 422)
	}

	/*
		if err := json.Unmarshal(body, &todo); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(422)
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}

			t := RepoCreateTodoWithName(todo, username)
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.WriteHeader(http.StatusCreated)

			if err := json.NewEncoder(w).Encode(t); err != nil {
				panic(err)
			}
		}*/
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todId := vars["todoId"]
	fmt.Println(w, "Todo show:", todId)
}
