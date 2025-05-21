package httpresponse

type BaseResponse[T any] struct {
	Status string `json:"status"`
	Data   T      `json:"data,omitempty"` // Only if status is "success"
	Error  string `json:"error,omitempty"` // Only if status is "error"
}

type ErrorResponse = BaseResponse[struct{}]

func NewErrorResponse(errMsg string) ErrorResponse {
	return ErrorResponse{
		Status: "error",
		Error:  errMsg,
	}
}

func NewSuccessResponse(data any) BaseResponse[any] {
	return BaseResponse[any]{
		Status: "success",
		Data:   data,
	}
}