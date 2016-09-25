package main

import (
	"github.com/cboecking/gorma-sample/app"
	"github.com/cboecking/gorma-sample/models"
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
)

//ErrDatabaseError - error returned when DB Query fails.
var ErrDatabaseError = goa.NewErrorClass("db error", 500)

// AccountbbbController implements the accountbbb resource.
type AccountbbbController struct {
	*goa.Controller
}

// NewAccountbbbController creates a accountbbb controller.
func NewAccountbbbController(service *goa.Service) *AccountbbbController {
	return &AccountbbbController{Controller: service.NewController("AccountbbbController")}
}

// Create runs the create action.
func (c *AccountbbbController) Create(ctx *app.CreateAccountbbbContext) error {
	// AccountbbbController_Create: start_implement
	a := models.Accountddd{}
	a.Name = ctx.Payload.Name
	err := accountModel.Add(ctx.Context, &a)
	if err != nil {
		return ErrDatabaseError(err)
	}
	ctx.ResponseData.Header().Set("location", app.AccountbbbHref(a.ID))
	return ctx.Created()

}

// Delete runs the delete action.
func (c *AccountbbbController) Delete(ctx *app.DeleteAccountbbbContext) error {
	// AccountbbbController_Delete: start_implement
	err := accountModel.Delete(ctx.Context, ctx.AccountbbbID)
	if err != nil {
		return ErrDatabaseError(err)
	}
	return ctx.NoContent()

}

// List runs the list action.
func (c *AccountbbbController) List(ctx *app.ListAccountbbbContext) error {
	// AccountbbbController_List: start_implement
	a := accountModel.ListAccountbbb(ctx.Context)
	return ctx.OK(a)

}

// Show runs the show action.
func (c *AccountbbbController) Show(ctx *app.ShowAccountbbbContext) error {
	// AccountbbbController_Show: start_implement
	a, err := accountModel.OneAccountbbb(ctx.Context, ctx.AccountbbbID)
	if err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	} else if err != nil {
		return ErrDatabaseError(err)
	}
	a.Href = app.AccountbbbHref(a.ID)
	return ctx.OK(a)

}

// Update runs the update action.
func (c *AccountbbbController) Update(ctx *app.UpdateAccountbbbContext) error {
	// AccountbbbController_Update: start_implement
	a, err := accountModel.Get(ctx.Context, ctx.AccountbbbID)
	if err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	} else if err != nil {
		return ErrDatabaseError(err)
	}
	a.Name = ctx.Payload.Name
	err = accountModel.Update(ctx, a)
	if err != nil {
		return ErrDatabaseError(err)
	}
	return ctx.NoContent()
}
