// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "appengine": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/akm/goa_gae_datastore_example/design
// --out=$(GOPATH)/src/github.com/akm/goa_gae_datastore_example
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// CreateUserContext provides the User create action context.
type CreateUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateUserPayload
}

// NewCreateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the User controller create action.
func NewCreateUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*CreateUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CreateUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// createUserPayload is the User create action payload.
type createUserPayload struct {
	// name
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createUserPayload) Validate() (err error) {
	if payload.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	return
}

// Publicize creates CreateUserPayload from createUserPayload
func (payload *createUserPayload) Publicize() *CreateUserPayload {
	var pub CreateUserPayload
	if payload.Name != nil {
		pub.Name = *payload.Name
	}
	return &pub
}

// CreateUserPayload is the User create action payload.
type CreateUserPayload struct {
	// name
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateUserPayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	return
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateUserContext) Created(r *User) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.user+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 201, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreateUserContext) BadRequest(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *CreateUserContext) Unauthorized(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 401, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *CreateUserContext) NotFound(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CreateUserContext) InternalServerError(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// DeleteUserContext provides the User delete action context.
type DeleteUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID interface{}
}

// NewDeleteUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the User controller delete action.
func NewDeleteUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*DeleteUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := DeleteUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteUserContext) NoContent(r *User) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.user+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 204, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *DeleteUserContext) BadRequest(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *DeleteUserContext) Unauthorized(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 401, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteUserContext) NotFound(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *DeleteUserContext) InternalServerError(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// ListUserContext provides the User list action context.
type ListUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Name string
}

// NewListUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the User controller list action.
func NewListUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramName := req.Params["name"]
	if len(paramName) == 0 {
		err = goa.MergeErrors(err, goa.MissingParamError("name"))
	} else {
		rawName := paramName[0]
		rctx.Name = rawName
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListUserContext) OK(r UserCollection) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.user+json; type=collection")
	}
	if r == nil {
		r = UserCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ListUserContext) BadRequest(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *ListUserContext) Unauthorized(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 401, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListUserContext) NotFound(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ListUserContext) InternalServerError(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// ShowUserContext provides the User show action context.
type ShowUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID interface{}
}

// NewShowUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the User controller show action.
func NewShowUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowUserContext) OK(r *User) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.user+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ShowUserContext) BadRequest(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *ShowUserContext) Unauthorized(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 401, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowUserContext) NotFound(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ShowUserContext) InternalServerError(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}

// UpdateUserContext provides the User update action context.
type UpdateUserContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID      interface{}
	Payload *UpdateUserPayload
}

// NewUpdateUserContext parses the incoming request URL and body, performs validations and creates the
// context used by the User controller update action.
func NewUpdateUserContext(ctx context.Context, r *http.Request, service *goa.Service) (*UpdateUserContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := UpdateUserContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		rctx.ID = rawID
	}
	return &rctx, err
}

// updateUserPayload is the User update action payload.
type updateUserPayload struct {
	// name
	Name *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *updateUserPayload) Validate() (err error) {
	if payload.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	return
}

// Publicize creates UpdateUserPayload from updateUserPayload
func (payload *updateUserPayload) Publicize() *UpdateUserPayload {
	var pub UpdateUserPayload
	if payload.Name != nil {
		pub.Name = *payload.Name
	}
	return &pub
}

// UpdateUserPayload is the User update action payload.
type UpdateUserPayload struct {
	// name
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
}

// Validate runs the validation rules defined in the design.
func (payload *UpdateUserPayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *UpdateUserContext) OK(r *User) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.user+json")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *UpdateUserContext) BadRequest(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// Unauthorized sends a HTTP response with status code 401.
func (ctx *UpdateUserContext) Unauthorized(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 401, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateUserContext) NotFound(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 404, r)
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *UpdateUserContext) InternalServerError(r error) error {
	if ctx.ResponseData.Header().Get("Content-Type") == "" {
		ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 500, r)
}