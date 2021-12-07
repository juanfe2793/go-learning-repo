package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		names := r.URL.Query()["name"]
		var name string

		if names == nil {

			w.Write([]byte("Please add a name in your query. ie: /?name=myname"))

		} else {

			if len(names) == 1 {
				name = names[0]
			} else {

				w.Write([]byte("We can handle only one name, for now"))
			}

			m := map[string]string{"name": name}
			enc := json.NewEncoder(w)

			w.Write([]byte("This is the Json encoded:  "))
			enc.Encode(m)
		}

	})

	err := http.ListenAndServe(":4000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
