package helper

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Response
	Data map[string]interface{} `json:"data"`
}

type ErrorResponse struct {
	Response
	Errors map[string]interface{} `json:"errors"`
}

func NewSuccessResponse(data map[string]interface{}) SuccessResponse {

	response := SuccessResponse{}

	response.Status = true
	response.Message = "Process completed successfully"

	if len(data) > 0 {
		response.Data = data
	}

	return response
}

func NewErrorResponse(errors map[string]interface{}) ErrorResponse {

	response := ErrorResponse{}

	response.Status = false
	response.Message = "Process failed"

	if len(errors) > 0 {
		response.Errors = errors
	}

	return response
}
