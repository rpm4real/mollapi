package actions

import (
	"fmt"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/rpm4real/mollapi/models"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Person)
// DB Table: Plural (people)
// Resource: Plural (People)
// Path: Plural (/people)
// View Template Folder: Plural (/templates/people/)

// PeopleResource is the resource for the Person model
type PeopleResource struct {
	buffalo.Resource
}

// List gets all People. This function is mapped to the path
// GET /people
func (v PeopleResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	people := &models.People{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all People from the DB
	if err := q.All(people); err != nil {
		return err
	}

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.Auto(c, people))
}

// Show gets the data for one Person. This function is mapped to
// the path GET /people/{person_id}
func (v PeopleResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Person
	person := &models.Person{}

	// To find the Person the parameter person_id is used.
	if err := tx.Find(person, c.Param("person_id")); err != nil {
		return c.Error(404, err)
	}

	return c.Render(200, r.Auto(c, person))
}

// Create adds a Person to the DB. This function is mapped to the
// path POST /people
func (v PeopleResource) Create(c buffalo.Context) error {
	// Allocate an empty Person
	person := &models.Person{}

	// Bind person to the html form elements
	if err := c.Bind(person); err != nil {
		return err
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(person)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, person))
	}

	// and redirect to the people index page
	return c.Render(201, r.Auto(c, person))
}

// Update changes a Person in the DB. This function is mapped to
// the path PUT /people/{person_id}
func (v PeopleResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Person
	person := &models.Person{}

	if err := tx.Find(person, c.Param("person_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Person to the html form elements
	if err := c.Bind(person); err != nil {
		return err
	}

	verrs, err := tx.ValidateAndUpdate(person)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.Auto(c, person))
	}

	// and redirect to the people index page
	return c.Render(200, r.Auto(c, person))
}

// Destroy deletes a Person from the DB. This function is mapped
// to the path DELETE /people/{person_id}
func (v PeopleResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty Person
	person := &models.Person{}

	// To find the Person the parameter person_id is used.
	if err := tx.Find(person, c.Param("person_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(person); err != nil {
		return err
	}

	// Redirect to the people index page
	return c.Render(200, r.Auto(c, person))
}
