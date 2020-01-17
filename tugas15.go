package tugas15

import (
	"fmt"
	"net/http"
)

func Tugas15(w http.ResponseWriter, r *http.Request) {
	for i := 1; i <= 100; i++ {
		fmt.Fprintln(w, i)
	}
}
