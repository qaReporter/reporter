package projects

import (
	"time"

	"github.com/google/uuid"
	"github.com/icrowley/fake"
)

type Controller struct {
	// Dependencies...
}

// Project struct
type Project struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Active      bool      `json:"active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ProjectResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}

// Index shows a list of projects
// GET /projects
func (c *Controller) Index() (prjs []*ProjectResponse) {

	for i := 0; i < 5; i++ {
		id, err := uuid.NewUUID()
		if err != nil {
			return prjs
		}
		p := &ProjectResponse{
			ID:          id.String(),
			Name:        fake.ProductName(),
			Description: fake.Words(),
			Active:      true,
		}
		prjs = append(prjs, p)
	}

	return prjs
}

// New project page
// GET /projects/new
func (c *Controller) New() {}

// Create a new project
// POST /projects
func (c *Controller) Create(name string, age int) (prj *Project, err error) {

	return
}

// Show a project
// GET /projects/:id
func (c *Controller) Show(id int) (prj *Project, err error) {

	return
}

// Update a project
// PATCH /projects/:id
func (c *Controller) Update(id int, name string, age int) (err error) {

	return
}

// Delete a project
// DELETE /projects/:id
func (c *Controller) Delete(id int) (err error) {

	return
}

// Edit project page
// GET /projects/:id/edit
func (c *Controller) Edit(id int) (prj *Project, err error) {

	return
}
