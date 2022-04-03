package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"time"
)

type Blog struct {
	ID          string
	Title       string
	Ingredients string
	Content     string
	Time        string
	AuthorName  string
}

var blog []Blog

func main() {

	//mock data - implement DB
	blog = append(blog, Blog{
		ID:          uuid.NewString(),
		Title:       "Pilaf Rice",
		Ingredients: "3/4 cup unsalted raw cashews\n1 tablespoon extra-virgin olive oil\n3 finely chopped, plus 2 whole small garlic cloves\n3/4 pound (about 18 spears), ends trimmed, cut into 2-inch pieces",
		Content: "How To Make a Simple Rice Pilaf:" + " " +
			"Place the rice in a strainer and rinse it thoroughly under cool water. The water running through the rice will look milky at first, " +
			"but will then become clearer and only lightly clouded. It's fine if there's still some haze in the water. There is no need to dry the rice before cooking; " +
			"a bit of moisture on the rice is fine. " +
			"Set the strainer of rice aside while you cook the onion.",
		Time:       time.Now().Format(time.RFC822),
		AuthorName: "Sidney Sheldon",
	})
	blog = append(blog, Blog{
		ID:          uuid.NewString(),
		Title:       "Concoction Rice",
		Ingredients: "3/4 cup unsalted raw cashews\n1 tablespoon extra-virgin olive oil\n3 finely chopped, plus 2 whole small garlic cloves\n3/4 pound (about 18 spears), ends trimmed, cut into 2-inch pieces",
		Content:     "In a small saucepan, cover the cashews with water to a depth of about 2 inches. Bring to a boil over medium-high heat. Remove from the heat and let soak for at least 10 minutes. Alternatively, soak the cashews in room temperature water overnight in the refrigerator.",
		Time:        time.Now().Format(time.RFC822),
		AuthorName:  "Cecilia",
	})
	blog = append(blog, Blog{
		ID:          uuid.NewString(),
		Title:       "Gizzdodo",
		Ingredients: "3 tablespoons fresh lemon juice (from about 1 large lemon), divided\n3/4 cup unsweetened non-dairy milk\n1 tablespoon white miso\n3/4 teaspoon onion powder\n1/2 (14-ounce) can quartered artichoke hearts, drained, and coarsely chopped",
		Content: "In a large skillet over medium heat, warm the oil. Cook the chopped garlic, stirring constantly, for about 30 seconds, until fragrant. Add the asparagus and cook, stirring constantly, for about 2 minute more, until bright green; season with salt and pepper." +
			"Transfer the drained cashews to a blender. Add the milk, miso, onion powder, 2 whole garlic cloves, the remaining 2 tablespoons of the lemon juice, and ¼ cup water. Blend on high speed, adding more water if needed, until completely smooth. ",
		Time:       time.Now().Format(time.RFC822),
		AuthorName: "Lovey",
	})

	//init router
	r := mux.NewRouter()

	//route handlers/ endpoints
	r.HandleFunc("/", getBlogs).Methods("GET")
	r.HandleFunc("/add-post", addFile).Methods("GET")
	r.HandleFunc("/add-post", postPost).Methods("POST")
	r.HandleFunc("/view-post/{id}", viewPost).Methods("GET")
	r.HandleFunc("/edit-post/{id}", editPost).Methods("GET")
	r.HandleFunc("/update-post", updatePost).Methods("POST")
	r.HandleFunc("/delete-post/{id}", deletePost).Methods("GET")

	//handling the css file
	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	//to run the server
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		return
	}
}

//-------------------------GET/VIEW A POST-----------------------------------------------
func getBlogs(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")

	if err != nil {
		fmt.Println(err)
	}
	_ = t.Execute(w, blog)
	//t.ExecuteTemplate(w, "index.html", p)
}

func viewPost(w http.ResponseWriter, r *http.Request) {
	blogInstance := Blog{}
	vars := mux.Vars(r)
	id := vars["id"]
	for _, v := range blog {
		if id == v.ID {
			blogInstance = v
		}
	}
	t, err := template.ParseFiles("templates/view.html")

	if err != nil {
		fmt.Println(err)
	}
	_ = t.Execute(w, blogInstance)
}

//-------------------------CREATE A POST-----------------------------------------------
func postPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	content := r.PostFormValue("content")

	blogPost := Blog{
		ID:         uuid.NewString(),
		Title:      title,
		Content:    content,
		Time:       time.Now().Format(time.RFC822),
		AuthorName: author,
	}
	blog = append(blog, blogPost)
	http.Redirect(w, r, "/", 302)
}

func addFile(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/add.html")

	if err != nil {
		fmt.Println(err)
	}
	_ = t.Execute(w, nil)
}

//-------------------------EDIT A POST HANDLER-----------------------------------------------
func editPost(w http.ResponseWriter, r *http.Request) {
	editInstance := Blog{}
	vars := mux.Vars(r)
	id := vars["id"]
	for _, v := range blog {
		if id == v.ID {
			editInstance = v
		}
	}
	t, err := template.ParseFiles("templates/edit.html")

	if err != nil {
		fmt.Println(err)
	}
	_ = t.Execute(w, editInstance)
}

//-------------------------UPDATE A POST-----------------------------------------------
func updatePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.PostFormValue("id")
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	content := r.PostFormValue("content")

	postUpdate := Blog{
		ID:         id,
		Title:      title,
		Content:    content,
		Time:       time.Now().Format(time.RFC822),
		AuthorName: author,
	}
	for i, v := range blog {
		if id == v.ID {
			blog[i] = postUpdate
		}
	}
	http.Redirect(w, r, "/", 302)
}

//delete a post
func deletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for i, v := range blog {
		if id == v.ID {
			blog = append(blog[:i], blog[i+1:]...)
		}
	}
	http.Redirect(w, r, "/", 302)
}
