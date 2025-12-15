package server

// sessionId: str
// chatInput: str
// universityId: str

type DiktiRequest struct {
	SessionId    string `json:"sessionId" validate:"required"`
	ChatInput    string `json:"chatInput" validate:"required"`
	UniversityId string `json:"universityId" validate:"required"`
}
