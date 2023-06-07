package response

import "github.com/alwinihza/talent-connect-be/model"

type EvaluationResponse struct {
	Mid   []model.Evaluation
	Final []model.Evaluation
}
