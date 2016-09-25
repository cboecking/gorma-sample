//************************************************************************//
// API "cellar-aaa": Application Contexts
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/cboecking/gorma-sample/design
// --out=$(GOPATH)src/github.com/cboecking/gorma-sample
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"strconv"
)

// CreateAccountbbbContext provides the accountbbb create action context.
type CreateAccountbbbContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateAccountbbbPayload
}

// NewCreateAccountbbbContext parses the incoming request URL and body, performs validations and creates the
// context used by the accountbbb controller create action.
func NewCreateAccountbbbContext(ctx context.Context, service *goa.Service) (*CreateAccountbbbContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateAccountbbbContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// createAccountbbbPayload is the accountbbb create action payload.
type createAccountbbbPayload struct {
	// Name of accountbbb
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createAccountbbbPayload) Validate() (err error) {
	if payload.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}

	return
}

// Publicize creates CreateAccountbbbPayload from createAccountbbbPayload
func (payload *createAccountbbbPayload) Publicize() *CreateAccountbbbPayload {
	var pub CreateAccountbbbPayload
	if payload.Name != nil {
		pub.Name = *payload.Name
	}
	return &pub
}

// CreateAccountbbbPayload is the accountbbb create action payload.
type CreateAccountbbbPayload struct {
	// Name of accountbbb
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateAccountbbbPayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}

	return
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateAccountbbbContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreateAccountbbbContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// DeleteAccountbbbContext provides the accountbbb delete action context.
type DeleteAccountbbbContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	AccountbbbID int
}

// NewDeleteAccountbbbContext parses the incoming request URL and body, performs validations and creates the
// context used by the accountbbb controller delete action.
func NewDeleteAccountbbbContext(ctx context.Context, service *goa.Service) (*DeleteAccountbbbContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := DeleteAccountbbbContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramAccountbbbID := req.Params["accountbbbID"]
	if len(paramAccountbbbID) > 0 {
		rawAccountbbbID := paramAccountbbbID[0]
		if accountbbbID, err2 := strconv.Atoi(rawAccountbbbID); err2 == nil {
			rctx.AccountbbbID = accountbbbID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("accountbbbID", rawAccountbbbID, "integer"))
		}
	}
	return &rctx, err
}

// NoContent sends a HTTP response with status code 204.
func (ctx *DeleteAccountbbbContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *DeleteAccountbbbContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DeleteAccountbbbContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListAccountbbbContext provides the accountbbb list action context.
type ListAccountbbbContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListAccountbbbContext parses the incoming request URL and body, performs validations and creates the
// context used by the accountbbb controller list action.
func NewListAccountbbbContext(ctx context.Context, service *goa.Service) (*ListAccountbbbContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ListAccountbbbContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListAccountbbbContext) OK(r AccountbbbCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.accountbbb+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *ListAccountbbbContext) OKLink(r AccountbbbLinkCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.accountbbb+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKTiny sends a HTTP response with status code 200.
func (ctx *ListAccountbbbContext) OKTiny(r AccountbbbTinyCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.accountbbb+json; type=collection")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// ShowAccountbbbContext provides the accountbbb show action context.
type ShowAccountbbbContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	AccountbbbID int
}

// NewShowAccountbbbContext parses the incoming request URL and body, performs validations and creates the
// context used by the accountbbb controller show action.
func NewShowAccountbbbContext(ctx context.Context, service *goa.Service) (*ShowAccountbbbContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowAccountbbbContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramAccountbbbID := req.Params["accountbbbID"]
	if len(paramAccountbbbID) > 0 {
		rawAccountbbbID := paramAccountbbbID[0]
		if accountbbbID, err2 := strconv.Atoi(rawAccountbbbID); err2 == nil {
			rctx.AccountbbbID = accountbbbID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("accountbbbID", rawAccountbbbID, "integer"))
		}
		if rctx.AccountbbbID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`accountbbbID`, rctx.AccountbbbID, 1, true))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowAccountbbbContext) OK(r *Accountbbb) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.accountbbb+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *ShowAccountbbbContext) OKLink(r *AccountbbbLink) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.accountbbb+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKTiny sends a HTTP response with status code 200.
func (ctx *ShowAccountbbbContext) OKTiny(r *AccountbbbTiny) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.accountbbb+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *ShowAccountbbbContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowAccountbbbContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// UpdateAccountbbbContext provides the accountbbb update action context.
type UpdateAccountbbbContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	AccountbbbID int
	Payload      *UpdateAccountbbbPayload
}

// NewUpdateAccountbbbContext parses the incoming request URL and body, performs validations and creates the
// context used by the accountbbb controller update action.
func NewUpdateAccountbbbContext(ctx context.Context, service *goa.Service) (*UpdateAccountbbbContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := UpdateAccountbbbContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramAccountbbbID := req.Params["accountbbbID"]
	if len(paramAccountbbbID) > 0 {
		rawAccountbbbID := paramAccountbbbID[0]
		if accountbbbID, err2 := strconv.Atoi(rawAccountbbbID); err2 == nil {
			rctx.AccountbbbID = accountbbbID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("accountbbbID", rawAccountbbbID, "integer"))
		}
	}
	return &rctx, err
}

// updateAccountbbbPayload is the accountbbb update action payload.
type updateAccountbbbPayload struct {
	// Name of accountbbb
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *updateAccountbbbPayload) Validate() (err error) {
	if payload.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}

	return
}

// Publicize creates UpdateAccountbbbPayload from updateAccountbbbPayload
func (payload *updateAccountbbbPayload) Publicize() *UpdateAccountbbbPayload {
	var pub UpdateAccountbbbPayload
	if payload.Name != nil {
		pub.Name = *payload.Name
	}
	return &pub
}

// UpdateAccountbbbPayload is the accountbbb update action payload.
type UpdateAccountbbbPayload struct {
	// Name of accountbbb
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate runs the validation rules defined in the design.
func (payload *UpdateAccountbbbPayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "name"))
	}

	return
}

// NoContent sends a HTTP response with status code 204.
func (ctx *UpdateAccountbbbContext) NoContent() error {
	ctx.ResponseData.WriteHeader(204)
	return nil
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *UpdateAccountbbbContext) BadRequest(r error) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.error")
	return ctx.ResponseData.Service.Send(ctx.Context, 400, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpdateAccountbbbContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}
