package handler

import "fmt"

func ErrorParamRequired(name, typ string) error {
	return fmt.Errorf("param %s is required for type %s", name, typ)
}

type CreateOpeningRequest struct {
	Role     string  `json:"role"`
	Company  string  `json:"company"`
	Location string  `json:"location"`
	Remote   *bool   `json:"remote"`
	Link     string  `json:"link"`
	Salary   float64 `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Remote == nil && r.Link == "" && r.Salary == 0 {
		return fmt.Errorf("malformed request body")
	}
	if r.Role == "" {
		return ErrorParamRequired("role", "string")
	}

	if r.Company == "" {
		return ErrorParamRequired("company", "string")
	}

	if r.Location == "" {
		return ErrorParamRequired("location", "string")
	}

	if r.Remote == nil {
		return ErrorParamRequired("remote", "bool")
	}

	if r.Link == "" {
		return ErrorParamRequired("link", "string")
	}

	if r.Salary <= 0 {
		return ErrorParamRequired("salary", "float64")
	}

	return nil
}
