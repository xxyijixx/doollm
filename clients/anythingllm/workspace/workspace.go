package workspace

import "time"

type CreateParams struct {
	Name string `json:"name"`
}

type CreateResponse struct {
	Workspace Workspace   `json:"workspace"`
	Message   interface{} `json:"message"`
}
type Threads struct {
	UserID int    `json:"user_id"`
	Slug   string `json:"slug"`
}

type ListResponse struct {
	Workspaces []Workspace `json:"workspaces"`
}

type GetResponse struct {
	Workspace []Workspace `json:"workspace"`
}

type Documents struct {
	ID            int       `json:"id"`
	DocID         string    `json:"docId"`
	Filename      string    `json:"filename"`
	Docpath       string    `json:"docpath"`
	WorkspaceID   int       `json:"workspaceId"`
	Metadata      string    `json:"metadata"`
	Pinned        bool      `json:"pinned"`
	Watched       bool      `json:"watched"`
	CreatedAt     time.Time `json:"createdAt"`
	LastUpdatedAt time.Time `json:"lastUpdatedAt"`
}

type Workspace struct {
	ID                   int         `json:"id"`
	Name                 string      `json:"name"`
	Slug                 string      `json:"slug"`
	VectorTag            interface{} `json:"vectorTag"`
	CreatedAt            time.Time   `json:"createdAt"`
	OpenAiTemp           interface{} `json:"openAiTemp"`
	OpenAiHistory        int         `json:"openAiHistory"`
	LastUpdatedAt        time.Time   `json:"lastUpdatedAt"`
	OpenAiPrompt         interface{} `json:"openAiPrompt"`
	SimilarityThreshold  float64     `json:"similarityThreshold"`
	ChatProvider         interface{} `json:"chatProvider"`
	ChatModel            interface{} `json:"chatModel"`
	TopN                 int         `json:"topN"`
	ChatMode             string      `json:"chatMode"`
	PfpFilename          interface{} `json:"pfpFilename"`
	AgentProvider        interface{} `json:"agentProvider"`
	AgentModel           interface{} `json:"agentModel"`
	QueryRefusalResponse interface{} `json:"queryRefusalResponse"`
	Threads              []Threads   `json:"threads"`
}

type UpdateEmbeddingsParams struct {
	Adds    []string `json:"adds"`
	Deletes []string `json:"deletes"`
}
