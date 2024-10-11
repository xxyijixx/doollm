package documents

// --------------------------------------------
// This section handles user authentication
// --------------------------------------------
type UploadResponse struct {
	Success   bool        `json:"success"`
	Error     interface{} `json:"error"`
	Documents []Documents `json:"documents"`
}

type Documents struct {
	ID                 string `json:"id"`
	URL                string `json:"url"`
	Title              string `json:"title"`
	DocAuthor          string `json:"docAuthor"`
	Description        string `json:"description"`
	DocSource          string `json:"docSource"`
	ChunkSource        string `json:"chunkSource"`
	Published          string `json:"published"`
	WordCount          int    `json:"wordCount"`
	PageContent        string `json:"pageContent"`
	TokenCountEstimate int    `json:"token_count_estimate"`
	Location           string `json:"location"`
}

type RawTextParams struct {
	TextContent string          `json:"textContent"`
	Metadata    RawTextMetadata `json:"metadata"`
}
type RawTextMetadata struct {
	// 中文不支持，会为“”
	Title string `json:"title"`
	// KeyOne string `json:"keyOne"`
	// KeyTwo string `json:"keyTwo"`
	// Etc    string `json:"etc"`
}

type QueryResponse struct {
	LocalFiles struct {
		Name  string `json:"name"`
		Type  string `json:"type"`
		Items []struct {
			Name  string `json:"name"`
			Type  string `json:"type"`
			Items []struct {
				Name               string        `json:"name"`
				Type               string        `json:"type"`
				ID                 string        `json:"id"`
				URL                string        `json:"url"`
				Title              string        `json:"title"`
				DocAuthor          string        `json:"docAuthor"`
				Description        string        `json:"description"`
				DocSource          string        `json:"docSource"`
				ChunkSource        string        `json:"chunkSource"`
				Published          string        `json:"published"`
				WordCount          int           `json:"wordCount"`
				TokenCountEstimate int           `json:"token_count_estimate"`
				Cached             bool          `json:"cached"`
				PinnedWorkspaces   []interface{} `json:"pinnedWorkspaces"`
				CanWatch           bool          `json:"canWatch"`
				Watched            bool          `json:"watched"`
			} `json:"items"`
		} `json:"items"`
	} `json:"localFiles"`
}
