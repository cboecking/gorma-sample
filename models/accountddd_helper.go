//************************************************************************//
// API "cellar-aaa": Model Helpers
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/cboecking/gorma-sample/design
// --out=$(GOPATH)src/github.com/cboecking/gorma-sample
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package models

import (
	"github.com/cboecking/gorma-sample/app"
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"golang.org/x/net/context"
	"time"
)

// MediaType Retrieval Functions

// ListAccountbbb returns an array of view: default.
func (m *AccountdddDB) ListAccountbbb(ctx context.Context) []*app.Accountbbb {
	defer goa.MeasureSince([]string{"goa", "db", "accountbbb", "listaccountbbb"}, time.Now())

	var native []*Accountddd
	var objs []*app.Accountbbb
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Accountddd", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.AccountdddToAccountbbb())
	}

	return objs
}

// AccountdddToAccountbbb loads a Accountddd and builds the default view of media type Accountbbb.
func (m *Accountddd) AccountdddToAccountbbb() *app.Accountbbb {
	accountddd := &app.Accountbbb{}
	accountddd.CreatedAt = m.CreatedAt
	accountddd.ID = m.ID
	accountddd.Name = m.Name

	return accountddd
}

// OneAccountbbb loads a Accountddd and builds the default view of media type Accountbbb.
func (m *AccountdddDB) OneAccountbbb(ctx context.Context, id int) (*app.Accountbbb, error) {
	defer goa.MeasureSince([]string{"goa", "db", "accountbbb", "oneaccountbbb"}, time.Now())

	var native Accountddd
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Accountddd", "error", err.Error())
		return nil, err
	}

	view := *native.AccountdddToAccountbbb()
	return &view, err
}

// MediaType Retrieval Functions

// ListAccountbbbLink returns an array of view: link.
func (m *AccountdddDB) ListAccountbbbLink(ctx context.Context) []*app.AccountbbbLink {
	defer goa.MeasureSince([]string{"goa", "db", "accountbbb", "listaccountbbblink"}, time.Now())

	var native []*Accountddd
	var objs []*app.AccountbbbLink
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Accountddd", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.AccountdddToAccountbbbLink())
	}

	return objs
}

// AccountdddToAccountbbbLink loads a Accountddd and builds the link view of media type Accountbbb.
func (m *Accountddd) AccountdddToAccountbbbLink() *app.AccountbbbLink {
	accountddd := &app.AccountbbbLink{}
	accountddd.ID = m.ID

	return accountddd
}

// OneAccountbbbLink loads a Accountddd and builds the link view of media type Accountbbb.
func (m *AccountdddDB) OneAccountbbbLink(ctx context.Context, id int) (*app.AccountbbbLink, error) {
	defer goa.MeasureSince([]string{"goa", "db", "accountbbb", "oneaccountbbblink"}, time.Now())

	var native Accountddd
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Accountddd", "error", err.Error())
		return nil, err
	}

	view := *native.AccountdddToAccountbbbLink()
	return &view, err
}

// MediaType Retrieval Functions

// ListAccountbbbTiny returns an array of view: tiny.
func (m *AccountdddDB) ListAccountbbbTiny(ctx context.Context) []*app.AccountbbbTiny {
	defer goa.MeasureSince([]string{"goa", "db", "accountbbb", "listaccountbbbtiny"}, time.Now())

	var native []*Accountddd
	var objs []*app.AccountbbbTiny
	err := m.Db.Scopes().Table(m.TableName()).Find(&native).Error

	if err != nil {
		goa.LogError(ctx, "error listing Accountddd", "error", err.Error())
		return objs
	}

	for _, t := range native {
		objs = append(objs, t.AccountdddToAccountbbbTiny())
	}

	return objs
}

// AccountdddToAccountbbbTiny loads a Accountddd and builds the tiny view of media type Accountbbb.
func (m *Accountddd) AccountdddToAccountbbbTiny() *app.AccountbbbTiny {
	accountddd := &app.AccountbbbTiny{}
	accountddd.ID = m.ID
	accountddd.Name = m.Name

	return accountddd
}

// OneAccountbbbTiny loads a Accountddd and builds the tiny view of media type Accountbbb.
func (m *AccountdddDB) OneAccountbbbTiny(ctx context.Context, id int) (*app.AccountbbbTiny, error) {
	defer goa.MeasureSince([]string{"goa", "db", "accountbbb", "oneaccountbbbtiny"}, time.Now())

	var native Accountddd
	err := m.Db.Scopes().Table(m.TableName()).Where("id = ?", id).Find(&native).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		goa.LogError(ctx, "error getting Accountddd", "error", err.Error())
		return nil, err
	}

	view := *native.AccountdddToAccountbbbTiny()
	return &view, err
}
