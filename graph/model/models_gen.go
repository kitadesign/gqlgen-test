// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/kitadesign/gqlgen-test/models"
)

type Node interface {
	IsNode()
	GetID() string
}

type Pagination interface {
	IsPagination()
	GetPageInfo() *PaginationInfo
	GetNodes() []Node
}

type CompanyPagination struct {
	PageInfo *PaginationInfo   `json:"pageInfo"`
	Nodes    []*models.Company `json:"nodes"`
}

func (CompanyPagination) IsPagination()                     {}
func (this CompanyPagination) GetPageInfo() *PaginationInfo { return this.PageInfo }
func (this CompanyPagination) GetNodes() []Node {
	if this.Nodes == nil {
		return nil
	}
	interfaceSlice := make([]Node, 0, len(this.Nodes))
	for _, concrete := range this.Nodes {
		interfaceSlice = append(interfaceSlice, concrete)
	}
	return interfaceSlice
}

type CreateCompanyInput struct {
	CompanyName    string `json:"companyName"`
	Representative string `json:"representative"`
	PhoneNumber    string `json:"phoneNumber"`
}

type CreateDepartmentInput struct {
	DepartmentName string `json:"departmentName"`
	Email          string `json:"email"`
}

type CreateEmployeeInput struct {
	Name          string        `json:"name"`
	Gender        models.Gender `json:"gender"`
	Email         string        `json:"email"`
	DependentsNum int           `json:"dependentsNum"`
	IsManager     bool          `json:"isManager"`
}

type DepartmentPagination struct {
	PageInfo *PaginationInfo      `json:"pageInfo"`
	Nodes    []*models.Department `json:"nodes"`
}

func (DepartmentPagination) IsPagination()                     {}
func (this DepartmentPagination) GetPageInfo() *PaginationInfo { return this.PageInfo }
func (this DepartmentPagination) GetNodes() []Node {
	if this.Nodes == nil {
		return nil
	}
	interfaceSlice := make([]Node, 0, len(this.Nodes))
	for _, concrete := range this.Nodes {
		interfaceSlice = append(interfaceSlice, concrete)
	}
	return interfaceSlice
}

type EmployeePagination struct {
	PageInfo *PaginationInfo    `json:"pageInfo"`
	Nodes    []*models.Employee `json:"nodes"`
}

func (EmployeePagination) IsPagination()                     {}
func (this EmployeePagination) GetPageInfo() *PaginationInfo { return this.PageInfo }
func (this EmployeePagination) GetNodes() []Node {
	if this.Nodes == nil {
		return nil
	}
	interfaceSlice := make([]Node, 0, len(this.Nodes))
	for _, concrete := range this.Nodes {
		interfaceSlice = append(interfaceSlice, concrete)
	}
	return interfaceSlice
}

type PaginationInfo struct {
	Page             int  `json:"page"`
	PaginationLength int  `json:"paginationLength"`
	HasNextPage      bool `json:"hasNextPage"`
	HasPreviousPage  bool `json:"hasPreviousPage"`
	Count            int  `json:"count"`
	TotalCount       int  `json:"totalCount"`
}

type UpdateCompanyInput struct {
	ID             string  `json:"id"`
	CompanyName    *string `json:"companyName,omitempty"`
	Representative *string `json:"representative,omitempty"`
	PhoneNumber    *string `json:"phoneNumber,omitempty"`
}

type UpdateDepartmentInput struct {
	ID             string  `json:"id"`
	DepartmentName *string `json:"departmentName,omitempty"`
	Email          *string `json:"email,omitempty"`
}

type UpdateEmployeeInput struct {
	ID            string         `json:"id"`
	Name          *string        `json:"name,omitempty"`
	Gender        *models.Gender `json:"gender,omitempty"`
	Email         *string        `json:"email,omitempty"`
	DependentsNum *int           `json:"dependentsNum,omitempty"`
	IsManager     *bool          `json:"isManager,omitempty"`
}
