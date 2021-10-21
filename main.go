package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type user struct {
	Id       string `db:"id"`
	Email    string `db:"email"`
	Usertype string `db:"user_type"`
}

type userrequest struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Usertype string `json:"user_type"`
}

type issue struct {
	Id          string    `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Reported_by string    `db:"reported_by"`
	Resolved_by string    `db:"resolved_by"`
	Status      string    `db:"status"`
	Resolved_at time.Time `db:"resolved_at"`
	Created_by  string    `db:"created_by"`
	Created_at  time.Time `db:"created_at"`
	Updated_at  time.Time `db:"updated_at"`
	Belongs_to  string    `db:"belongs_to"`
}

type issuerequest struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ReportedBy  string    `json:"reported_by"`
	ResolvedBy  string    `json:"resolved_by"`
	Status      string    `json:"status"`
	ResolvedAt  time.Time `json:"resolved_at"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	BelongsTo   string    `json:"belongs_to"`
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Complete User Data")
	var userinfo user

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", "postgres", "postgres", "localhost", 5432, "dpay_helpdesk")
	db, err := sqlx.Open("postgres", dbURL)
	if err != nil {
		fmt.Println(1, err.Error())
		return
	}

	// dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", "postgres", "postgres", "localhost", 5432, "dpay_helpdesk")
	// db, err := sqlx.Open("postgres", dbURL)
	// if err != nil {
	// 	fmt.Println(1, err.Error())
	// 	return
	// }
	defer db.Close()

	query := `SELECT * from "user";`
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(2, err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&userinfo.Id, &userinfo.Email, &userinfo.Usertype)
		if err != nil {
			fmt.Println(3, err.Error())
			return
		}
		fmt.Println("\n", userinfo.Id, userinfo.Email, userinfo.Usertype)
	}

	err = rows.Err()
	if err != nil {
		fmt.Println(4, err.Error())
		return
	}
	db.Close()
}

func ListIssues(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Complete Issue Data")
	var issueinfo issue

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", "postgres", "postgres", "localhost", 5432, "dpay_helpdesk")
	db, err := sqlx.Open("postgres", dbURL)
	if err != nil {
		fmt.Println(1, err.Error())
		return
	}
	defer db.Close()

	query := `SELECT * from "issue";`
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(2, err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&issueinfo.Id, &issueinfo.Title, &issueinfo.Description, &issueinfo.Reported_by, &issueinfo.Resolved_by, &issueinfo.Status, &issueinfo.Resolved_at, &issueinfo.Created_by, &issueinfo.Created_at, &issueinfo.Updated_at, &issueinfo.Belongs_to)
		if err != nil {
			fmt.Println(3, err.Error())
			return
		}
		fmt.Println("\n", issueinfo.Id, issueinfo.Title, issueinfo.Description, issueinfo.Reported_by, issueinfo.Resolved_by, issueinfo.Status, issueinfo.Resolved_at, issueinfo.Created_by, issueinfo.Created_at, issueinfo.Updated_at, issueinfo.Belongs_to)
	}

	err = rows.Err()
	if err != nil {
		fmt.Println(4, err.Error())
		return
	}
	db.Close()
}

func GetDetailsByID(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["UserId"]
	fmt.Println(id)

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", "postgres", "postgres", "localhost", 5432, "dpay_helpdesk")
	d, err := sqlx.Open("postgres", dbURL)
	if err != nil {
		return
	}
	ctx := req.Context()
	query := `Select * FROM "user" WHERE id=$1;`
	var userinfo user

	err = d.GetContext(ctx, &userinfo, query, id)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Given the details to the user")
	fmt.Fprintf(res, "email: %v\n", userinfo.Email)
	fmt.Println(userinfo.Email)
	d.Close()
}

func ShowIssueStatus(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["IssueId"]
	fmt.Println(id)

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", "postgres", "postgres", "localhost", 5432, "dpay_helpdesk")
	d, err := sqlx.Open("postgres", dbURL)
	if err != nil {
		return
	}
	ctx := req.Context()
	query := `Select * FROM "issue" WHERE id=$1;`
	var issueinfo issue

	err = d.GetContext(ctx, &issueinfo, query, id)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Details given to the user")
	fmt.Fprintf(res, "Issue_id: %v\n", issueinfo.Id)
	fmt.Fprintf(res, "title: %v\n", issueinfo.Title)
	fmt.Fprintf(res, "Description: %v\n", issueinfo.Description)
	fmt.Fprintf(res, "User_id: %v\n", issueinfo.Belongs_to)
	fmt.Fprintf(res, "Issue Submitted: %v\n", issueinfo.Created_at)
	fmt.Fprintf(res, "Status: %v\n", issueinfo.Status)
	fmt.Fprintf(res, "Updated At: %v\n", issueinfo.Updated_at)
	d.Close()
}

//Handle new User syntax
// {
//     "id":"1",
//     "email":"a@example.com",
//     "type":"Customer"
// }
func HandleNewUSer(res http.ResponseWriter, req *http.Request) {
	var userdetails userrequest
	// var userinfo user

	err := json.NewDecoder(req.Body).Decode(&userdetails)
	if err != nil {
		return
	}

	// fmt.Println(userdetails.Id, userdetails.Email)

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", "postgres", "postgres", "localhost", 5432, "dpay_helpdesk")
	d, err := sqlx.Open("postgres", dbURL)
	if err != nil {
		return
	}

	ctx := req.Context()
	// query := `INSERT INTO public."user"(id, email) VALUES ($1,$2) RETURNING id, email;`

	// err = d.GetContext(ctx, &userinfo, query, userdetails.Id, userdetails.Email)

	query := `INSERT INTO public."user"(id, email, user_type) VALUES ($1,$2,$3);`

	_, err = d.ExecContext(ctx, query, userdetails.Id, userdetails.Email, userdetails.Usertype)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Succesfully Registered")
	fmt.Fprintf(res, "Succesfully Registered")
	d.Close()
}

//Handle new Issue
// {
//     "id":"1",
//     "title":"Server Issue",
//     "description":"my server is not running",
//     "reported_by":"1",
//     "created_by":"1",
//     "belongs_to":"1"
// }
func HandleNewIssue(res http.ResponseWriter, req *http.Request) {
	var issuedetails issuerequest

	err := json.NewDecoder(req.Body).Decode(&issuedetails)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", "postgres", "postgres", "localhost", 5432, "dpay_helpdesk")
	d, err := sqlx.Open("postgres", dbURL)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ctx := req.Context()
	query := `INSERT INTO public.issue(
		id, title, description, reported_by, resolved_by, status ,resolved_at, created_by, created_at, updated_at, belongs_to)
		VALUES ($1, $2, $3, $4, 'Not assinged', DEFAULT, CURRENT_TIMESTAMP, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, $6);`

	_, err = d.ExecContext(ctx, query, issuedetails.Id, issuedetails.Title, issuedetails.Description, issuedetails.ReportedBy, issuedetails.CreatedBy, issuedetails.BelongsTo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Succesfully Submitted the issue")
	fmt.Fprintf(res, "Succesfully Submitted the issue")
	d.Close()
}

func Hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello Welcome to helpdesk")
}

func AssignCustomerCare(res http.ResponseWriter, req *http.Request) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", "postgres", "postgres", "localhost", 5432, "dpay_helpdesk")
	d, err := sqlx.Open("postgres", dbURL)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ctx := req.Context()
	query := `UPDATE public.issue
	SET status='Inprogress', resolved_by='#10', updated_at=CURRENT_TIMESTAMP
	WHERE status='New';`

	_, err = d.ExecContext(ctx, query)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Status set to Inprogress for all the new requests.")
	fmt.Fprintf(res, "Status set to Inprogress for all the new requests.")
	d.Close()
}

//SolveQuery
// {
// 	"id":"1",
// 	"description":"please restart the service"
// }
func SolveQuery(res http.ResponseWriter, req *http.Request) {
	var issuedetails issuerequest

	err := json.NewDecoder(req.Body).Decode(&issuedetails)
	if err != nil {
		return
	}

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", "postgres", "postgres", "localhost", 5432, "dpay_helpdesk")
	d, err := sqlx.Open("postgres", dbURL)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ctx := req.Context()
	query := `UPDATE public.issue
	SET description=$2, status='Closed', resolved_at=CURRENT_TIMESTAMP, updated_at=CURRENT_TIMESTAMP
	WHERE id=$1;`

	_, err = d.ExecContext(ctx, query, issuedetails.Id, issuedetails.Description)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Issue Resolved")
	fmt.Fprintf(res, "Description: %v\n", issuedetails.Description)
	d.Close()
}

func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)

	//GET requests
	r.HandleFunc("/hello", Hello).Methods("GET")                           //Welcome to the helpdesk
	r.HandleFunc("/listusers", ListUsers).Methods("GET")                   //List all the users
	r.HandleFunc("/listissues", ListIssues).Methods("GET")                 //List all the issues
	r.HandleFunc("/user/{UserId}", GetDetailsByID).Methods("GET")          //Get User Status
	r.HandleFunc("/issuestatus/{IssueId}", ShowIssueStatus).Methods("GET") //Get issue status

	//POST requests
	r.HandleFunc("/newuser", HandleNewUSer).Methods("POST")    //Register as a new User
	r.HandleFunc("/postissue", HandleNewIssue).Methods("POST") //Submit a Issue

	//PUT requests
	r.HandleFunc("/assingcustomercare", AssignCustomerCare).Methods("PUT") //Assing Customer Care to a Query
	r.HandleFunc("/solvequery", SolveQuery).Methods("PUT")                 //Solve query and change status, description and update time

	log.Fatal(http.ListenAndServe(":1001", r))
}

func main() {
	handleRequests()
}