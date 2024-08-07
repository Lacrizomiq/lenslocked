package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
	"github.com/joncalhoun/lenslocked/controllers"
	"github.com/joncalhoun/lenslocked/models"
	"github.com/joncalhoun/lenslocked/templates"
	"github.com/joncalhoun/lenslocked/views"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("Request: %s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
		fmt.Printf("Request processed in %s\n", time.Since(start))
	})
}

func main() {
	r := chi.NewRouter()

	r.Use(LoggingMiddleware)

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS,
		"home.gohtml", "tailwind.gohtml",
	))))
	r.Get("/contact", controllers.StaticHandler(views.Must(
		views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))
	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))

	// Setup a database connection
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Setup our model services
	userService := models.UserService{
		DB: db,
	}

	// Setup our controllers
	usersC := controllers.Users{
		UserService: &userService,
	}
	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS, "signup.gohtml", "tailwind.gohtml"))
	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)

	usersC.Templates.SignIn = views.Must(views.ParseFS(
		templates.FS, "signin.gohtml", "tailwind.gohtml"))
	r.Get("/signin", usersC.SignIn)
	r.Post("/signin", usersC.ProcessSignIn)
	r.Get("/users/me", usersC.CurrentUser)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	csrfKey := "32-byte-long-auth-key"
	csrfMiddleware := csrf.Protect([]byte(csrfKey), csrf.Secure(false))
	http.ListenAndServe(":3000", csrfMiddleware(r))
}
