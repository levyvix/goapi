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

type UpdateOpeningRequest struct {
	Role     string  `json:"role,omitempty"`
	Company  string  `json:"company,omitempty"`
	Location string  `json:"location,omitempty"`
	Remote   *bool   `json:"remote,omitempty"`
	Link     string  `json:"link,omitempty"`
	Salary   float64 `json:"salary,omitempty"`
}

func (r *UpdateOpeningRequest) Validate() error {
	// if any field is not null then it is an update

	if r.Role == "" && r.Company == "" && r.Location == "" && r.Remote == nil && r.Link == "" && r.Salary == 0 {
		return fmt.Errorf("malformed request body. At least one field is required")
	}

	return nil
}
