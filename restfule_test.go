package gorestful

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGORestful_GET(t *testing.T) {
	router := New()
	hit := false
	router.GET("/foo", func(w http.ResponseWriter, r *http.Request) {
		hit = true
		w.WriteHeader(200)
	})
	r, _ := http.NewRequest("GET", "/foo", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	if hit == false {
		t.Errorf("Expected hit is true, got:%v", hit)
	}
}
func TestGORestful_HEAD(t *testing.T) {
	router := New()
	hit := false
	router.HEAD("/foo", func(writer http.ResponseWriter, request *http.Request) {
		hit = true
		writer.WriteHeader(200)
	})
	r, _ := http.NewRequest("HEAD", "/foo", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	if hit == false {
		t.Errorf("Expected hit is true, got:%v", hit)
	}
}
func TestGORestful_POST(t *testing.T) {
	router := New()
	hit := false
	router.POST("/foo", func(writer http.ResponseWriter, request *http.Request) {
		hit = true
		writer.WriteHeader(200)
	})
	r, _ := http.NewRequest("POST", "/foo", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	if hit == false {
		t.Errorf("Expected hit is true, got:%v", hit)
	}
}
func TestGORestful_PUT(t *testing.T) {
	router := New()
	hit := false
	router.PUT("/foo", func(writer http.ResponseWriter, request *http.Request) {
		hit = true
		writer.WriteHeader(200)
	})
	r, _ := http.NewRequest("PUT", "/foo", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	if hit == false {
		t.Errorf("Expected hit is true, got:%v", hit)
	}
}
func TestGORestful_PATCH(t *testing.T) {
	router := New()
	hit := false
	router.PATCH("/foo", func(writer http.ResponseWriter, request *http.Request) {
		hit = true
		writer.WriteHeader(200)
	})
	r, _ := http.NewRequest("PATCH", "/foo", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	if hit == false {
		t.Errorf("Expected hit is true, got:%v", hit)
	}
}
func TestGORestful_DELETE(t *testing.T) {
	router := New()
	hit := false
	router.DELETE("/foo", func(writer http.ResponseWriter, request *http.Request) {
		hit = true
		writer.WriteHeader(200)
	})
	r, _ := http.NewRequest("DELETE", "/foo", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	if hit == false {
		t.Errorf("Expected hit is true, got:%v", hit)
	}
}
func TestGORestful_CONNECT(t *testing.T) {
	router := New()
	hit := false
	router.CONNECT("/foo", func(writer http.ResponseWriter, request *http.Request) {
		hit = true
		writer.WriteHeader(200)
	})
	r, _ := http.NewRequest("CONNECT", "/foo", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	if hit == false {
		t.Errorf("Expected hit is true, got:%v", hit)
	}
}
func TestGORestful_TRACE(t *testing.T) {
	router := New()
	hit := false
	router.TRACE("/foo", func(writer http.ResponseWriter, request *http.Request) {
		hit = true
		writer.WriteHeader(200)
	})
	r, _ := http.NewRequest("TRACE", "/foo", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	if hit == false {
		t.Errorf("Expected hit is true, got:%v", hit)
	}
}
func TestGORestful_Run(t *testing.T) {
	router := New()
	router.GET("/foo", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	})
	ts := httptest.NewServer(router)
	defer ts.Close()
	api := ts.URL
	fmt.Println("api", api)
	result, err := GetOk(api)
	if err != nil {
		t.Error(err)
	}
	if result != "ok" {
		t.Errorf("Expect:ok, got:%s", result)
	}

}

func TestGORestful_HandlerServeHTTP(t *testing.T) {
	router := New()
	got := "hello"
	router.HandlerServeHTTP("GET", "/foo", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(got))
	})
	ts := httptest.NewServer(router)
	defer ts.Close()
	api := ts.URL
	fmt.Println("api", api)
	result, err := GetOk(api)
	if err != nil {
		t.Error(err)
	}
	if result != got {
		t.Errorf("Expect:%s, got:%s", got, result)
	}
}

func TestGORestful_Run2(t *testing.T) {
	SetMode(ModeDebug)
	router := New()
	router.GET("/foo", func(w http.ResponseWriter, r *http.Request) {
	})
	router.POST("/foo1", func(writer http.ResponseWriter, request *http.Request) {
	})

	http.ListenAndServe(":8080", router)
}


func GetOk(api string) (result string, err error) {
	url := fmt.Sprintf("%s/foo", api)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http status code is not 200, status-code:%d", resp.StatusCode)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	result = string(b)
	return
}