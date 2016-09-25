//************************************************************************//
// API "cellar-aaa": Application Controllers
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
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// AccountbbbController is the controller interface for the Accountbbb actions.
type AccountbbbController interface {
	goa.Muxer
	Create(*CreateAccountbbbContext) error
	Delete(*DeleteAccountbbbContext) error
	List(*ListAccountbbbContext) error
	Show(*ShowAccountbbbContext) error
	Update(*UpdateAccountbbbContext) error
}

// MountAccountbbbController "mounts" a Accountbbb resource controller on the given service.
func MountAccountbbbController(service *goa.Service, ctrl AccountbbbController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateAccountbbbContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateAccountbbbPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/cellar-aaa/accountbbbs", ctrl.MuxHandler("Create", h, unmarshalCreateAccountbbbPayload))
	service.LogInfo("mount", "ctrl", "Accountbbb", "action", "Create", "route", "POST /cellar-aaa/accountbbbs")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteAccountbbbContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	service.Mux.Handle("DELETE", "/cellar-aaa/accountbbbs/:accountbbbID", ctrl.MuxHandler("Delete", h, nil))
	service.LogInfo("mount", "ctrl", "Accountbbb", "action", "Delete", "route", "DELETE /cellar-aaa/accountbbbs/:accountbbbID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListAccountbbbContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/cellar-aaa/accountbbbs", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Accountbbb", "action", "List", "route", "GET /cellar-aaa/accountbbbs")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowAccountbbbContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/cellar-aaa/accountbbbs/:accountbbbID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Accountbbb", "action", "Show", "route", "GET /cellar-aaa/accountbbbs/:accountbbbID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateAccountbbbContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateAccountbbbPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	service.Mux.Handle("PUT", "/cellar-aaa/accountbbbs/:accountbbbID", ctrl.MuxHandler("Update", h, unmarshalUpdateAccountbbbPayload))
	service.LogInfo("mount", "ctrl", "Accountbbb", "action", "Update", "route", "PUT /cellar-aaa/accountbbbs/:accountbbbID")
}

// unmarshalCreateAccountbbbPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateAccountbbbPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createAccountbbbPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateAccountbbbPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateAccountbbbPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &updateAccountbbbPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}
