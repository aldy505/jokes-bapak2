package schema

type ImageAPI struct {
	Data    ImageAPIData `json:"data"`
	Success bool         `json:"success"`
	Status  int          `json:"status"`
}

type ImageAPIData struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	URLViewer  string `json:"url_viewer"`
	URL        string `json:"url"`
	DisplayURL string `json:"display_url"`
}
