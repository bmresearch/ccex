package rest

type BaseResponse struct {
	Success bool `json:"success"`
}

func (b *BaseResponse) WasSuccessful() bool {
	return b.Success
}
