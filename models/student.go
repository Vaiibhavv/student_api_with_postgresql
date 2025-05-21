// Student represents a student record.
// swagger:model
package models

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name" validate:"required"`
	Age   int    `json:"age" validate:"required"`
	Grade string `json:"grade" validate:"required"`
}
