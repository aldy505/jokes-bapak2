package submit

type Submission struct {
	ID        int    `json:"id,omitempty" db:"id"`
	Link      string `json:"link" form:"link" db:"link"`
	Image     string `json:"image,omitempty" form:"image"`
	CreatedAt string `json:"created_at" db:"created_at"`
	Author    string `json:"author" form:"author" db:"author"`
	Status    int    `json:"status" db:"status"`
}

type SubmissionQuery struct {
	Author   string `query:"author"`
	Limit    string `query:"limit"`
	Page     string `query:"page"`
	Approved string `query:"approved"`
}

type ResponseSubmission struct {
	ID      string     `json:"id,omitempty"`
	Message string     `json:"message,omitempty"`
	Data    Submission `json:"data,omitempty"`
}

type Error struct {
	Error string `json:"error"`
}
