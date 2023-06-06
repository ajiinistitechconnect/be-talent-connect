package response

import "github.com/alwinihza/talent-connect-be/model"

type ProgramListResponse struct {
	Admin       []model.Program `json:"admin,omitempty"`
	Panelist    []model.Program `json:"panelist,omitempty"`
	Participant []model.Program `json:"participant,omitempty"`
	Mentor      []model.Program `json:"mentor,omitempty"`
}
