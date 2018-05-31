package controller

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"

	"github.com/goadesign/goa"

	"github.com/akm/goa_gae_datastore_example/app"
	"github.com/akm/goa_gae_datastore_example/model"
)

// UserController implements the User resource.
type UserController struct {
	*goa.Controller
}

// NewUserController creates a User controller.
func NewUserController(service *goa.Service) *UserController {
	return &UserController{Controller: service.NewController("UserController")}
}

// Create runs the create action.
func (c *UserController) Create(ctx *app.CreateUserContext) error {
	// UserController_Create: start_implement

	// Put your logic here
	appCtx := appengine.NewContext(ctx.Request)
	udb := model.UserDB{}
	u := &model.User{
		Name: ctx.Payload.Name,
	}
	addu, err := udb.Add(appCtx, u)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}

	return ctx.Created(addu.UserToUser())

	// UserController_Create: end_implement
}

// Delete runs the delete action.
func (c *UserController) Delete(ctx *app.DeleteUserContext) error {
	// UserController_Delete: start_implement

	// Put your logic here
	appCtx := appengine.NewContext(ctx.Request)
	udb := model.UserDB{}
	int64ID, err := model.ConvertIdIntoInt64(ctx.ID)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}
	err = udb.Delete(appCtx, int64ID)
	if err == datastore.ErrNoSuchEntity {
		return ctx.NotFound(goa.ErrNotFound(err))
	} else if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}

	return nil

	// UserController_Delete: end_implement
}

// List runs the list action.
func (c *UserController) List(ctx *app.ListUserContext) error {
	// UserController_List: start_implement

	// Put your logic here
	appCtx := appengine.NewContext(ctx.Request)
	udb := model.UserDB{}
	us, err := udb.GetFindByName(appCtx, ctx.Name)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}

	return ctx.OK(us)

	// UserController_List: end_implement
}

// Show runs the show action.
func (c *UserController) Show(ctx *app.ShowUserContext) error {
	// UserController_Show: start_implement

	// Put your logic here
	appCtx := appengine.NewContext(ctx.Request)
	udb := model.UserDB{}
	int64ID, err := model.ConvertIdIntoInt64(ctx.ID)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}
	getu, err := udb.Get(appCtx, int64ID)
	if err == datastore.ErrNoSuchEntity {
		return ctx.NotFound(goa.ErrNotFound(err))
	} else if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}

	return ctx.OK(getu.UserToUser())

	// UserController_Show: end_implement
}

// Update runs the update action.
func (c *UserController) Update(ctx *app.UpdateUserContext) error {
	// UserController_Update: start_implement

	// Put your logic here
	appCtx := appengine.NewContext(ctx.Request)
	udb := model.UserDB{}
	int64ID, err := model.ConvertIdIntoInt64(ctx.ID)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}
	u := &model.User{
		ID:   int64ID,
		Name: ctx.Payload.Name,
	}
	u, err = udb.Update(appCtx, int64ID, u)
	if err == datastore.ErrNoSuchEntity {
		return ctx.NotFound(goa.ErrNotFound(err))
	} else if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}

	return ctx.OK(u.UserToUser())

	// UserController_Update: end_implement
}
