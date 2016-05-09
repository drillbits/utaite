package app

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"google.golang.org/appengine/aetest"
)

func TestMemberAPIList(t *testing.T) {
	inst, err := aetest.NewInstance(&aetest.Options{
		StronglyConsistentDatastore: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer inst.Close()

	r, err := inst.NewRequest("GET", "/api/member", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(w, r)

	if w.Code != 200 {
		t.Fatalf("expeced: `%v`, but got: `%v`", 200, w.Code)
	}
	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != `{"list":[]}` {
		t.Fatalf("expeced: `%v`, but got: `%v`", `{"list":[]}`, string(body))
	}
}
