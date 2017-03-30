package hydraweb

import (
	"fmt"
	"math/rand"
	"net/http"
)

type testhandler struct {
	r int
}

func newHandler() testhandler {
	return testhandler{
		r: rand.Int(),
	}
}

func (h testhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "Welcome to the Hydra software system")
	case "/testhandle":
		fmt.Fprint(w, "test handle object with random number: ", h.r)
	}
	fmt.Println(r.URL.Query())
}
