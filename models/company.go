package models

type Company struct {
	ID             string `json:"id"`
	CompanyName    string `json:"companyName"`
	Representative string `json:"representative"`
	PhoneNumber    string `json:"phoneNumber"`
	DepartmentID   string `json:"department_id"`
	EmployeeID     string `json:"employee_id"`
}

func (Company) IsNode()         {}
func (c Company) GetID() string { return c.ID }

func (Company) TableName() string {
	return "m_company"
}
