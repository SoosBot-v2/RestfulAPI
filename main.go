package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

type Sad struct {
	ID			string      `json:"id,omitempty"`
    URL        string   `json:"url,omitempty"`
}
var Sads []Sad

type StartSad struct {
	StartWith			string      `json:"StartWith,omitempty"`
    EndWith        string   `json:"EndWith,omitempty"`
}

func main() {


	Sads = append(Sads, Sad{ID: "1", URL: "https://i.ytimg.com/vi/db9sbgN_gCw/maxresdefault.jpg"})
	Sads = append(Sads, Sad{ID: "2", URL: "https://i.ytimg.com/vi/ONIrJ8fOJqs/maxresdefault.jpg"})
	Sads = append(Sads, Sad{ID: "3", URL: "https://1.bp.blogspot.com/-jp09Xr4XNmc/WvGVTuAa40I/AAAAAAAAQiw/zk4FBOnXByoqBS9HyV5OufCg3Nqr7bEwwCLcBGAs/s1600/thumb_6b1e356d-7cbd-4646-8ad5-5657aca16592.jpg"})
	Sads = append(Sads, Sad{ID: "4", URL: "https://78.media.tumblr.com/b8b6c02ce2fc56d7821c38dbfae236bf/tumblr_o2syq3usxw1ub9qlao1_500.gif"})

	router := mux.NewRouter()
	router.HandleFunc("/sad", GetStartSad).Methods("GET")
	router.HandleFunc("/sad/{id}", GetSad).Methods("GET")
	
    log.Fatal(http.ListenAndServe(":8000", router))
}

func GetStartSad(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(&StartSad{StartWith: "1", EndWith: "4"})
}

func GetSad(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range Sads {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Sad{})
}