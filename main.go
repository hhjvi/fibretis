package main

import (
	// stackoverflow.com/q/10687627
	"./grass"
	"./soil"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	// Index page [grass/index.go]
	r.HandleFunc("/", grass.IndexHandler)
	// Accounts-related [grass/accounts.go]
	r.HandleFunc("/login", grass.LoginHandler).Methods("GET", "POST")
	r.HandleFunc("/login/{return:.*}", grass.LoginHandler).Methods("GET", "POST")
	r.HandleFunc("/signup", grass.SignupHandler).Methods("GET", "POST")
	r.HandleFunc("/signup/{return:.*}", grass.SignupHandler).Methods("GET", "POST")
	r.HandleFunc("/logout", grass.LogoutHandler)
	r.HandleFunc("/logout/{return:.*}", grass.LogoutHandler)
	r.HandleFunc("/notifications", grass.NotificationsHandler)
	r.HandleFunc("/profedit", grass.ProfEditHandler)
	// Projects-related [grass/projects.go]
	r.HandleFunc("/projects", grass.ProjectsHandler)
	r.HandleFunc("/projects/{page:[0-9]+}", grass.ProjectsHandler)
	r.HandleFunc("/project_create", grass.ProjectEditHandler)
	r.HandleFunc("/project/{prjid:[0-9]+}", grass.ProjectPageHandler)
	r.HandleFunc("/project_edit/{prjid:[0-9]+}", grass.ProjectEditHandler)
	r.HandleFunc("/invite/{prjid:[0-9]+}", grass.InviteHandler)
	r.HandleFunc("/invite/{prjid:[0-9]+}/{aid:[0-9]+}", grass.InviteHandler)
	r.HandleFunc("/answer_invitation/{token}", grass.AnswerInvitationHandler)
	// Posts-related [grass/posts.go]
	r.HandleFunc("/post_create/{prjid:[0-9]+}", grass.PostEditHandler)
	r.HandleFunc("/post_edit/{pstid:[0-9]+}", grass.PostEditHandler)
	r.HandleFunc("/post/{pstid:[0-9]+}", grass.PostPageHandler)
	r.HandleFunc("/comment/{pstid:[0-9]+}", grass.CommentHandler).Methods("POST")
	// Sights-related [grass/sights.go]
	r.HandleFunc("/sight", grass.SightHandler).Methods("POST")
	// Searching [grass/search.go]
	r.HandleFunc("/account_search/invite/{prjid:[0-9]+}/{q}", grass.AccountSearchHandler)
	// Error pages [grass/errors.go]
	r.HandleFunc("/error", grass.ErrorHandler)
	// Static file server
	r.PathPrefix("/leaves").Handler(http.FileServer(http.Dir("./stalks/")))
	r.PathPrefix("/uploads").Handler(http.FileServer(http.Dir("./")))
	http.Handle("/", r)
	soil.InitDatabase()
	fmt.Println("Now serving")
	http.ListenAndServe(":14706", nil)
}
