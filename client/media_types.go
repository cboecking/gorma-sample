//************************************************************************//
// API "cellar-aaa": Application Media Types
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/cboecking/gorma-sample/design
// --out=$(GOPATH)src/github.com/cboecking/gorma-sample
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package client

import (
	"github.com/goadesign/goa"
	"net/http"
	"time"
)

// A Accountccc (default view)
//
// Identifier: application/vnd.accountbbb+json; view=default
type Accountbbb struct {
	// Date of creation
	CreatedAt time.Time `form:"created_at" json:"created_at" xml:"created_at"`
	// Email of accountbbb owner
	CreatedBy string `form:"created_by" json:"created_by" xml:"created_by"`
	// API href of accountbbb
	Href string `form:"href" json:"href" xml:"href"`
	// ID of accountbbb
	ID int `form:"id" json:"id" xml:"id"`
	// Name of accountbbb
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the Accountbbb media type instance.
func (mt *Accountbbb) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.CreatedBy == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "created_by"))
	}

	if err2 := goa.ValidateFormat(goa.FormatEmail, mt.CreatedBy); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.created_by`, mt.CreatedBy, goa.FormatEmail, err2))
	}
	return
}

// A Accountccc (link view)
//
// Identifier: application/vnd.accountbbb+json; view=link
type AccountbbbLink struct {
	// API href of accountbbb
	Href string `form:"href" json:"href" xml:"href"`
	// ID of accountbbb
	ID int `form:"id" json:"id" xml:"id"`
}

// Validate validates the AccountbbbLink media type instance.
func (mt *AccountbbbLink) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}

	return
}

// A Accountccc (tiny view)
//
// Identifier: application/vnd.accountbbb+json; view=tiny
type AccountbbbTiny struct {
	// API href of accountbbb
	Href string `form:"href" json:"href" xml:"href"`
	// ID of accountbbb
	ID int `form:"id" json:"id" xml:"id"`
	// Name of accountbbb
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the AccountbbbTiny media type instance.
func (mt *AccountbbbTiny) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	return
}

// DecodeAccountbbb decodes the Accountbbb instance encoded in resp body.
func (c *Client) DecodeAccountbbb(resp *http.Response) (*Accountbbb, error) {
	var decoded Accountbbb
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeAccountbbbLink decodes the AccountbbbLink instance encoded in resp body.
func (c *Client) DecodeAccountbbbLink(resp *http.Response) (*AccountbbbLink, error) {
	var decoded AccountbbbLink
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeAccountbbbTiny decodes the AccountbbbTiny instance encoded in resp body.
func (c *Client) DecodeAccountbbbTiny(resp *http.Response) (*AccountbbbTiny, error) {
	var decoded AccountbbbTiny
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// AccountbbbCollection is the media type for an array of Accountbbb (default view)
//
// Identifier: application/vnd.accountbbb+json; type=collection; view=default
type AccountbbbCollection []*Accountbbb

// Validate validates the AccountbbbCollection media type instance.
func (mt AccountbbbCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Href == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "href"))
		}
		if e.Name == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "name"))
		}
		if e.CreatedBy == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "created_by"))
		}

		if err2 := goa.ValidateFormat(goa.FormatEmail, e.CreatedBy); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`response[*].created_by`, e.CreatedBy, goa.FormatEmail, err2))
		}
	}
	return
}

// AccountbbbCollection is the media type for an array of Accountbbb (link view)
//
// Identifier: application/vnd.accountbbb+json; type=collection; view=link
type AccountbbbLinkCollection []*AccountbbbLink

// Validate validates the AccountbbbLinkCollection media type instance.
func (mt AccountbbbLinkCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Href == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "href"))
		}

	}
	return
}

// AccountbbbCollection is the media type for an array of Accountbbb (tiny view)
//
// Identifier: application/vnd.accountbbb+json; type=collection; view=tiny
type AccountbbbTinyCollection []*AccountbbbTiny

// Validate validates the AccountbbbTinyCollection media type instance.
func (mt AccountbbbTinyCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Href == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "href"))
		}
		if e.Name == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "name"))
		}

	}
	return
}

// DecodeAccountbbbCollection decodes the AccountbbbCollection instance encoded in resp body.
func (c *Client) DecodeAccountbbbCollection(resp *http.Response) (AccountbbbCollection, error) {
	var decoded AccountbbbCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeAccountbbbLinkCollection decodes the AccountbbbLinkCollection instance encoded in resp body.
func (c *Client) DecodeAccountbbbLinkCollection(resp *http.Response) (AccountbbbLinkCollection, error) {
	var decoded AccountbbbLinkCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeAccountbbbTinyCollection decodes the AccountbbbTinyCollection instance encoded in resp body.
func (c *Client) DecodeAccountbbbTinyCollection(resp *http.Response) (AccountbbbTinyCollection, error) {
	var decoded AccountbbbTinyCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeErrorResponse decodes the ErrorResponse instance encoded in resp body.
func (c *Client) DecodeErrorResponse(resp *http.Response) (*goa.ErrorResponse, error) {
	var decoded goa.ErrorResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
