package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

//------------------------------TEST VIEW HOMEPAGE----------------------------------
func TestGetBlogs(t *testing.T) {
	route := mux.NewRouter()
	route.HandleFunc("/", getBlogs).Methods("GET")
	request, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		return
	}
	request.Header.Set("content-type", "text/html")
	response := httptest.NewRecorder()
	route.ServeHTTP(response, request)
	text := `<p>By Sidney Sheldon</p>`
	if !strings.Contains(response.Body.String(), text) {
		t.Errorf("Expected ")
	}
	if response.Code != 200 {
		t.Errorf("expected %v but got %v", 200, response.Code)
	}
	//log.Println(response.Body.String())
}


//------------------------------TEST ADD FILE----------------------------------
func TestAddFile(t *testing.T) {
	route := mux.NewRouter()
	testServer := httptest.NewServer(route)
	req, err := http.NewRequest(http.MethodGet, testServer.URL+"/add-post", nil)
	if err != nil {
		t.Fatalf("%v", err)
	}

	response := httptest.NewRecorder()
	addFile(response, req)
	log.Println(response.Body)

	submit := `<input type="submit" id="btn" value="submit">`
	if !strings.Contains(response.Body.String(), submit) {
		t.Errorf("expected %v but got %v", response.Body.String(), submit)
	}
	content := `<textarea name="content"></textarea>`
	if !strings.Contains(response.Body.String(), content) {
		t.Errorf("expected %v but got %v", response.Body.String(), content)
	}
	//if response.Code != http.StatusFound {
	//	t.Errorf("expected %v but got %v", http.StatusFound, response.Code)
	//}
}


func TestMakeAPost(t *testing.T) {
    route := mux.NewRouter()
	form := url.Values{}
	form.Set("id", "1")
	form.Set("title", "blogTitle")
	form.Set("authorName", "nameOfAuthor")
	form.Set("content", "contentOfBlog")
	route.HandleFunc("/add-post", postPost).Methods("POST")
	request, err := http.NewRequest(http.MethodPost, "/add-post", strings.NewReader(form.Encode()))
	if err != nil {
		return
	}
	request.Header.Set("content-type", "text.html")
	response := httptest.NewRecorder()
	route.ServeHTTP(response, request)
	//fmt.Println(response.Body.String())
	//fmt.Println(response.Code)

	if response.Code != 302 {
		t.Errorf("expected %v but got %v", 302, response.Code)
	}
}
//------------------------------TEST VIEW POST----------------------------------
func TestViewPost(t *testing.T) {
	route := mux.NewRouter()
	route.HandleFunc("/view-post/{id}", viewPost).Methods("GET")
	request, err := http.NewRequest(http.MethodGet, "/view-post/1", nil)
	if err != nil {
		return
	}
	request.Header.Set("content-type", "text/html")
	response := httptest.NewRecorder()
	route.ServeHTTP(response, request)
	//viewPost(response, request)
	fmt.Println(response.Body.String())

	title := `<h3>Pilaf Rice</h3>`
	if !strings.Contains(response.Body.String(), title) {
		t.Errorf("expected %s but got %s", title, response.Body.String())
	}
}


//------------------------------TEST EDIT POST----------------------------------
func TestEditPost(t *testing.T) {
	route := mux.NewRouter()

	route.HandleFunc("/edit-post/{id}", editPost).Methods("GET")
	request, err := http.NewRequest(http.MethodGet, "/edit-post/1", nil)
	if err != nil {
		return
	}
	request.Header.Set("content-type", "text.html")
	response := httptest.NewRecorder()
	route.ServeHTTP(response, request)
	//editPost(response, request)
	fmt.Println(response.Body.String())

	authorName := `<input type="text" name="author" value="Sidney Sheldon">`
	if !strings.Contains(response.Body.String(), authorName) {
		t.Errorf("expected %s but got %s", authorName, response.Body.String())
	}
}


//------------------------------TEST UPDATE POST----------------------------------
func TestUpdatePost(t *testing.T) {
	route := mux.NewRouter()
	form := url.Values{}
	form.Set("id", "1")
	form.Set("title", "blogTitle")
	form.Set("authorName", "nameOfAuthor")
	form.Set("content", "contentOfBlog")
	route.HandleFunc("/update-post", updatePost).Methods("POST")
	request, err := http.NewRequest(http.MethodPost, "/update-post", strings.NewReader(form.Encode()))
	if err != nil {
		return
	}
	request.Header.Set("content-type", "text.html")
	response := httptest.NewRecorder()
	route.ServeHTTP(response, request)
	//fmt.Println(response.Body.String())
	//fmt.Println(response.Code)

	if response.Code != 302 {
		t.Errorf("expected %v but got %v", 302, response.Code)
	}
}


//------------------------------TEST DELETE POST----------------------------------
func TestDeletePost(t *testing.T) {
	route := mux.NewRouter()

	route.HandleFunc("/delete-post/{id}", deletePost).Methods("GET")
	request, err := http.NewRequest(http.MethodGet, "/delete-post/1", nil)
	if err != nil {
		return
	}
	request.Header.Set("content-type", "text.html")
	response := httptest.NewRecorder()
	route.ServeHTTP(response, request)
	fmt.Println(response.Body.String())

	if response.Code != 302 {
		t.Errorf("expected %v but got %v", 302, response.Code)
	}
}
