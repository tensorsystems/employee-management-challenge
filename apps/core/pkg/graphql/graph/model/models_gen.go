// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph_models

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/tensorsystems/employee-management-challenge/apps/core/pkg/models"
)

type Connection interface {
	IsConnection()
	GetTotalCount() int
	GetPageInfo() *PageInfo
}

type CreateDepartmentInput struct {
	Name string `json:"name"`
}

type CreateEmployeeInput struct {
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	Gender       string    `json:"gender"`
	PhoneNo      string    `json:"phoneNo"`
	Email        string    `json:"email"`
	DateOfBirth  time.Time `json:"dateOfBirth"`
	Country      string    `json:"country"`
	Region       *string   `json:"region"`
	City         *string   `json:"city"`
	SubCity      *string   `json:"subCity"`
	Woreda       *string   `json:"woreda"`
	Zone         *string   `json:"zone"`
	Kebele       *string   `json:"kebele"`
	HouseNo      *string   `json:"houseNo"`
	DepartmentID int       `json:"departmentId"`
}

type DepartmentConnection struct {
	TotalCount int               `json:"totalCount"`
	PageInfo   *PageInfo         `json:"pageInfo"`
	Edges      []*DepartmentEdge `json:"edges"`
}

func (DepartmentConnection) IsConnection()               {}
func (this DepartmentConnection) GetTotalCount() int     { return this.TotalCount }
func (this DepartmentConnection) GetPageInfo() *PageInfo { return this.PageInfo }

type DepartmentEdge struct {
	Node *models.Department `json:"node"`
}

type EmployeeConnection struct {
	TotalCount int             `json:"totalCount"`
	PageInfo   *PageInfo       `json:"pageInfo"`
	Edges      []*EmployeeEdge `json:"edges"`
}

func (EmployeeConnection) IsConnection()               {}
func (this EmployeeConnection) GetTotalCount() int     { return this.TotalCount }
func (this EmployeeConnection) GetPageInfo() *PageInfo { return this.PageInfo }

type EmployeeEdge struct {
	Node *models.Employee `json:"node"`
}

type FileUpload struct {
	File graphql.Upload `json:"file"`
	Name string         `json:"name"`
}

type HealthCheckReport struct {
	Health string `json:"health"`
	Db     bool   `json:"db"`
}

type PageInfo struct {
	TotalPages int `json:"totalPages"`
}

type UpdateDepartmentInput struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UpdateEmployeeInput struct {
	ID           int       `json:"id"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	Gender       string    `json:"gender"`
	PhoneNo      string    `json:"phoneNo"`
	Email        string    `json:"email"`
	DateOfBirth  time.Time `json:"dateOfBirth"`
	Country      string    `json:"country"`
	Region       *string   `json:"region"`
	City         *string   `json:"city"`
	SubCity      *string   `json:"subCity"`
	Woreda       *string   `json:"woreda"`
	Zone         *string   `json:"zone"`
	Kebele       *string   `json:"kebele"`
	HouseNo      *string   `json:"houseNo"`
	DepartmentID int       `json:"departmentId"`
}

type OrderDirection string

const (
	OrderDirectionAsc  OrderDirection = "asc"
	OrderDirectionDesc OrderDirection = "desc"
)

var AllOrderDirection = []OrderDirection{
	OrderDirectionAsc,
	OrderDirectionDesc,
}

func (e OrderDirection) IsValid() bool {
	switch e {
	case OrderDirectionAsc, OrderDirectionDesc:
		return true
	}
	return false
}

func (e OrderDirection) String() string {
	return string(e)
}

func (e *OrderDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderDirection", str)
	}
	return nil
}

func (e OrderDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
