package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.38

import (
	"context"
	"fmt"

	"github.com/kitadesign/gqlgen-test/graph/model"
)

// CreateEmployee is the resolver for the createEmployee field.
func (r *mutationResolver) CreateEmployee(ctx context.Context, input model.CreateEmployeeInput) (*model.Employee, error) {
	panic(fmt.Errorf("not implemented: CreateEmployee - createEmployee"))
}

// UpdateEmployee is the resolver for the updateEmployee field.
func (r *mutationResolver) UpdateEmployee(ctx context.Context, input model.UpdateEmployeeInput) (*model.Employee, error) {
	panic(fmt.Errorf("not implemented: UpdateEmployee - updateEmployee"))
}

// DeleteEmployee is the resolver for the deleteEmployee field.
func (r *mutationResolver) DeleteEmployee(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteEmployee - deleteEmployee"))
}

// MergeEmployee is the resolver for the mergeEmployee field.
func (r *mutationResolver) MergeEmployee(ctx context.Context, id1 string, id2 string) (bool, error) {
	panic(fmt.Errorf("not implemented: MergeEmployee - mergeEmployee"))
}

// Employee is the resolver for the employee field.
func (r *queryResolver) Employee(ctx context.Context, id string) (*model.Employee, error) {
	panic(fmt.Errorf("not implemented: Employee - employee"))
}

// Employees is the resolver for the employees field.
func (r *queryResolver) Employees(ctx context.Context, limit int, offset *int, email *string, gender *model.Gender, isManager *bool, hasDependent *bool) (*model.EmployeePagination, error) {
	panic(fmt.Errorf("not implemented: Employees - employees"))
}