//************************************************************************//
// API "cellar-aaa": accountbbb TestHelpers
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/cboecking/gorma-sample/design
// --out=$(GOPATH)src/github.com/cboecking/gorma-sample
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package test

import (
	"bytes"
	"fmt"
	"github.com/cboecking/gorma-sample/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"golang.org/x/net/context"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
)

// CreateAccountbbbBadRequest runs the method Create of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func CreateAccountbbbBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.AccountbbbController, payload *app.CreateAccountbbbPayload) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		if e.ResponseStatus() != 400 {
			t.Errorf("unexpected payload validation error: %+v", e)
		}
		return nil, e
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/cellar-aaa/accountbbbs"),
	}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AccountbbbTest"), rw, req, prms)
	createCtx, err := app.NewCreateAccountbbbContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	createCtx.Payload = payload

	// Perform action
	err = ctrl.Create(createCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var ok bool
		mt, ok = resp.(error)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of error", resp)
		}
	}

	// Return results
	return rw, mt
}

// CreateAccountbbbCreated runs the method Create of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func CreateAccountbbbCreated(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.AccountbbbController, payload *app.CreateAccountbbbPayload) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		if e.ResponseStatus() != 201 {
			t.Errorf("unexpected payload validation error: %+v", e)
		}
		return nil
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/cellar-aaa/accountbbbs"),
	}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AccountbbbTest"), rw, req, prms)
	createCtx, err := app.NewCreateAccountbbbContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	createCtx.Payload = payload

	// Perform action
	err = ctrl.Create(createCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 201 {
		t.Errorf("invalid response status code: got %+v, expected 201", rw.Code)
	}

	// Return results
	return rw
}

// DeleteAccountbbbBadRequest runs the method Delete of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func DeleteAccountbbbBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.AccountbbbController, accountbbbID int) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/cellar-aaa/accountbbbs/%v", accountbbbID),
	}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountbbbID"] = []string{fmt.Sprintf("%v", accountbbbID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AccountbbbTest"), rw, req, prms)
	deleteCtx, err := app.NewDeleteAccountbbbContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.Delete(deleteCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var ok bool
		mt, ok = resp.(error)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of error", resp)
		}
	}

	// Return results
	return rw, mt
}

// DeleteAccountbbbNoContent runs the method Delete of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func DeleteAccountbbbNoContent(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.AccountbbbController, accountbbbID int) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/cellar-aaa/accountbbbs/%v", accountbbbID),
	}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountbbbID"] = []string{fmt.Sprintf("%v", accountbbbID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AccountbbbTest"), rw, req, prms)
	deleteCtx, err := app.NewDeleteAccountbbbContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.Delete(deleteCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 204 {
		t.Errorf("invalid response status code: got %+v, expected 204", rw.Code)
	}

	// Return results
	return rw
}

// DeleteAccountbbbNotFound runs the method Delete of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func DeleteAccountbbbNotFound(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.AccountbbbController, accountbbbID int) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/cellar-aaa/accountbbbs/%v", accountbbbID),
	}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountbbbID"] = []string{fmt.Sprintf("%v", accountbbbID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AccountbbbTest"), rw, req, prms)
	deleteCtx, err := app.NewDeleteAccountbbbContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.Delete(deleteCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

	// Return results
	return rw
}

// ListAccountbbbOK runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListAccountbbbOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.AccountbbbController) (http.ResponseWriter, app.AccountbbbCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/cellar-aaa/accountbbbs"),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AccountbbbTest"), rw, req, prms)
	listCtx, err := app.NewListAccountbbbContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.List(listCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.AccountbbbCollection
	if resp != nil {
		var ok bool
		mt, ok = resp.(app.AccountbbbCollection)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.AccountbbbCollection", resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}

// ListAccountbbbOKLink runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListAccountbbbOKLink(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.AccountbbbController) (http.ResponseWriter, app.AccountbbbLinkCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/cellar-aaa/accountbbbs"),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AccountbbbTest"), rw, req, prms)
	listCtx, err := app.NewListAccountbbbContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.List(listCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.AccountbbbLinkCollection
	if resp != nil {
		var ok bool
		mt, ok = resp.(app.AccountbbbLinkCollection)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.AccountbbbLinkCollection", resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}

// ListAccountbbbOKTiny runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListAccountbbbOKTiny(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.AccountbbbController) (http.ResponseWriter, app.AccountbbbTinyCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/cellar-aaa/accountbbbs"),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AccountbbbTest"), rw, req, prms)
	listCtx, err := app.NewListAccountbbbContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.List(listCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.AccountbbbTinyCollection
	if resp != nil {
		var ok bool
		mt, ok = resp.(app.AccountbbbTinyCollection)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.AccountbbbTinyCollection", resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}

// ShowAccountbbbBadRequest runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowAccountbbbBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.AccountbbbController, accountbbbID int) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/cellar-aaa/accountbbbs/%v", accountbbbID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountbbbID"] = []string{fmt.Sprintf("%v", accountbbbID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AccountbbbTest"), rw, req, prms)
	showCtx, err := app.NewShowAccountbbbContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.Show(showCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var ok bool
		mt, ok = resp.(error)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of error", resp)
		}
	}

	// Return results
	return rw, mt
}

// ShowAccountbbbNotFound runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowAccountbbbNotFound(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.AccountbbbController, accountbbbID int) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/cellar-aaa/accountbbbs/%v", accountbbbID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountbbbID"] = []string{fmt.Sprintf("%v", accountbbbID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AccountbbbTest"), rw, req, prms)
	showCtx, err := app.NewShowAccountbbbContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.Show(showCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

	// Return results
	return rw
}

// ShowAccountbbbOK runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowAccountbbbOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.AccountbbbController, accountbbbID int) (http.ResponseWriter, *app.Accountbbb) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/cellar-aaa/accountbbbs/%v", accountbbbID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountbbbID"] = []string{fmt.Sprintf("%v", accountbbbID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AccountbbbTest"), rw, req, prms)
	showCtx, err := app.NewShowAccountbbbContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.Show(showCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.Accountbbb
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.Accountbbb)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.Accountbbb", resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}

// ShowAccountbbbOKLink runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowAccountbbbOKLink(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.AccountbbbController, accountbbbID int) (http.ResponseWriter, *app.AccountbbbLink) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/cellar-aaa/accountbbbs/%v", accountbbbID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountbbbID"] = []string{fmt.Sprintf("%v", accountbbbID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AccountbbbTest"), rw, req, prms)
	showCtx, err := app.NewShowAccountbbbContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.Show(showCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.AccountbbbLink
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.AccountbbbLink)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.AccountbbbLink", resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}

// ShowAccountbbbOKTiny runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowAccountbbbOKTiny(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.AccountbbbController, accountbbbID int) (http.ResponseWriter, *app.AccountbbbTiny) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/cellar-aaa/accountbbbs/%v", accountbbbID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountbbbID"] = []string{fmt.Sprintf("%v", accountbbbID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AccountbbbTest"), rw, req, prms)
	showCtx, err := app.NewShowAccountbbbContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	// Perform action
	err = ctrl.Show(showCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.AccountbbbTiny
	if resp != nil {
		var ok bool
		mt, ok = resp.(*app.AccountbbbTiny)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of app.AccountbbbTiny", resp)
		}
		err = mt.Validate()
		if err != nil {
			t.Errorf("invalid response media type: %s", err)
		}
	}

	// Return results
	return rw, mt
}

// UpdateAccountbbbBadRequest runs the method Update of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func UpdateAccountbbbBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.AccountbbbController, accountbbbID int, payload *app.UpdateAccountbbbPayload) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		if e.ResponseStatus() != 400 {
			t.Errorf("unexpected payload validation error: %+v", e)
		}
		return nil, e
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/cellar-aaa/accountbbbs/%v", accountbbbID),
	}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountbbbID"] = []string{fmt.Sprintf("%v", accountbbbID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AccountbbbTest"), rw, req, prms)
	updateCtx, err := app.NewUpdateAccountbbbContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	updateCtx.Payload = payload

	// Perform action
	err = ctrl.Update(updateCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var ok bool
		mt, ok = resp.(error)
		if !ok {
			t.Fatalf("invalid response media: got %+v, expected instance of error", resp)
		}
	}

	// Return results
	return rw, mt
}

// UpdateAccountbbbNoContent runs the method Update of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func UpdateAccountbbbNoContent(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.AccountbbbController, accountbbbID int, payload *app.UpdateAccountbbbPayload) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		if e.ResponseStatus() != 204 {
			t.Errorf("unexpected payload validation error: %+v", e)
		}
		return nil
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/cellar-aaa/accountbbbs/%v", accountbbbID),
	}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountbbbID"] = []string{fmt.Sprintf("%v", accountbbbID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AccountbbbTest"), rw, req, prms)
	updateCtx, err := app.NewUpdateAccountbbbContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	updateCtx.Payload = payload

	// Perform action
	err = ctrl.Update(updateCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 204 {
		t.Errorf("invalid response status code: got %+v, expected 204", rw.Code)
	}

	// Return results
	return rw
}

// UpdateAccountbbbNotFound runs the method Update of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func UpdateAccountbbbNotFound(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.AccountbbbController, accountbbbID int, payload *app.UpdateAccountbbbPayload) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		if e.ResponseStatus() != 404 {
			t.Errorf("unexpected payload validation error: %+v", e)
		}
		return nil
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/cellar-aaa/accountbbbs/%v", accountbbbID),
	}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["accountbbbID"] = []string{fmt.Sprintf("%v", accountbbbID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AccountbbbTest"), rw, req, prms)
	updateCtx, err := app.NewUpdateAccountbbbContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	updateCtx.Payload = payload

	// Perform action
	err = ctrl.Update(updateCtx)

	// Validate response
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

	// Return results
	return rw
}
