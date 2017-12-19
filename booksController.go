package main
import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)
func putHandle(encoder *json.Encoder, w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pathSplits := strings.Split(path, "/")
	if len(pathSplits) == 3 && len(pathSplits[2]) == 8 {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for i := 0; i < len(books); i++ {
			if books[i].Id == pathSplits[2] {
				if len(r.Form["isRent"]) == 1 {
					isRent, error := strconv.ParseBool(r.Form["isRent"][0])
					if error != nil {
						return
					}
					books[i].IsRent = isRent
				}
				if len(r.Form["name"]) == 1 {
					name := r.Form["name"][0]
					if err != nil {
						return
					}
					books[i].Name = name
				}
				encoder.Encode(books)
			}
		}
	}
}
func postHandle(encoder *json.Encoder, w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(r.Form["name"]) == 1 {
		newBook := book{
			Name:   string(r.Form["name"][0]),
			IsRent: false,
			Id:     RandStringRunes(8),
		}
		books = append(books, newBook)
		encoder.Encode(books)
	}
	return
}
