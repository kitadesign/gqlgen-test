package models

type Department struct {
	ID             string `json:"id"`
	DepartmentName string `json:"departmentName"`
	Email          string `json:"email"`
}

func (Department) IsNode()         {}
func (d Department) GetID() string { return d.ID }
