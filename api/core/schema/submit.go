package schema

type Submission struct {
	ID        int    `json:"id,omitempty"`
	Link      string `json:"link"`
	Image     string `json:"image,omitempty"`
	CreatedAt string `json:"created_at"`
	Author    string `json:"author"`
	Status    int    `json:"status"`
}

type SubmissionQuery struct {
	Author   string `query:"author"`
	Limit    string `query:"limit"`
	Page     string `query:"page"`
	Approved string `query:"approved"`
}

type ResponseSubmission struct {
	ID         string     `json:"id,omitempty"`
	Message    string     `json:"message,omitempty"`
	Submission Submission `json:"submission,omitempty"`
	AuthorPage string     `json:"author_page,omitempty"`
}
