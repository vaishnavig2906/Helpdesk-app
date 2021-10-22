package issue

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/vaishnavig2906/Helpdesk-app/issue"
)

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
	var IssueDetails issue.IssueRequest

	err := json.NewDecoder(req.Body).Decode(&IssueDetails)
	if err != nil {
		log.Fatal(err)
		return
	}

	db, err := DB.init_DB()
	if err != nil {
		log.Fatal(err)
		return
	}

	ctx := req.Context()
	query := `INSERT INTO public.issue(
		id, title, description, reported_by, resolved_by, status ,resolved_at, created_by, created_at, updated_at, belongs_to)
		VALUES ($1, $2, $3, $4, 'Not assinged', DEFAULT, CURRENT_TIMESTAMP, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, $6);`

	_, err = db.ExecContext(ctx, query, IssueDetails.Id, IssueDetails.Title, IssueDetails.Description, IssueDetails.ReportedBy, IssueDetails.CreatedBy, IssueDetails.BelongsTo)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Fprintf(res, "Succesfully Submitted the issue")
	db.Close()
}
