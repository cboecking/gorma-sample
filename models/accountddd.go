//************************************************************************//
// API "cellar-aaa": Models
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

// cellar-aaa Accountddd
type Accountddd struct {
	ID        int `gorm:"primary_key"` // primary key
	CreatedAt time.Time
	DeletedAt *time.Time
	Name      string
	UpdatedAt time.Time
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m Accountddd) TableName() string {
	return "accountddds"

}

// AccountdddDB is the implementation of the storage interface for
// Accountddd.
type AccountdddDB struct {
	Db *gorm.DB
}

// NewAccountdddDB creates a new storage type.
func NewAccountdddDB(db *gorm.DB) *AccountdddDB {
	return &AccountdddDB{Db: db}
}

// DB returns the underlying database.
func (m *AccountdddDB) DB() interface{} {
	return m.Db
}

// AccountdddStorage represents the storage interface.
type AccountdddStorage interface {
	DB() interface{}
	List(ctx context.Context) ([]*Accountddd, error)
	Get(ctx context.Context, id int) (*Accountddd, error)
	Add(ctx context.Context, accountddd *Accountddd) error
	Update(ctx context.Context, accountddd *Accountddd) error
	Delete(ctx context.Context, id int) error

	ListAccountbbb(ctx context.Context) []*app.Accountbbb
	OneAccountbbb(ctx context.Context, id int) (*app.Accountbbb, error)

	ListAccountbbbLink(ctx context.Context) []*app.AccountbbbLink
	OneAccountbbbLink(ctx context.Context, id int) (*app.AccountbbbLink, error)

	ListAccountbbbTiny(ctx context.Context) []*app.AccountbbbTiny
	OneAccountbbbTiny(ctx context.Context, id int) (*app.AccountbbbTiny, error)
}

// TableName overrides the table name settings in Gorm to force a specific table name
// in the database.
func (m *AccountdddDB) TableName() string {
	return "accountddds"

}

// CRUD Functions

// Get returns a single Accountddd as a Database Model
// This is more for use internally, and probably not what you want in  your controllers
func (m *AccountdddDB) Get(ctx context.Context, id int) (*Accountddd, error) {
	defer goa.MeasureSince([]string{"goa", "db", "accountddd", "get"}, time.Now())

	var native Accountddd
	err := m.Db.Table(m.TableName()).Where("id = ?", id).Find(&native).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	}

	return &native, err
}

// List returns an array of Accountddd
func (m *AccountdddDB) List(ctx context.Context) ([]*Accountddd, error) {
	defer goa.MeasureSince([]string{"goa", "db", "accountddd", "list"}, time.Now())

	var objs []*Accountddd
	err := m.Db.Table(m.TableName()).Find(&objs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return objs, nil
}

// Add creates a new record.
func (m *AccountdddDB) Add(ctx context.Context, model *Accountddd) error {
	defer goa.MeasureSince([]string{"goa", "db", "accountddd", "add"}, time.Now())

	err := m.Db.Create(model).Error
	if err != nil {
		goa.LogError(ctx, "error adding Accountddd", "error", err.Error())
		return err
	}

	return nil
}

// Update modifies a single record.
func (m *AccountdddDB) Update(ctx context.Context, model *Accountddd) error {
	defer goa.MeasureSince([]string{"goa", "db", "accountddd", "update"}, time.Now())

	obj, err := m.Get(ctx, model.ID)
	if err != nil {
		goa.LogError(ctx, "error updating Accountddd", "error", err.Error())
		return err
	}
	err = m.Db.Model(obj).Updates(model).Error

	return err
}

// Delete removes a single record.
func (m *AccountdddDB) Delete(ctx context.Context, id int) error {
	defer goa.MeasureSince([]string{"goa", "db", "accountddd", "delete"}, time.Now())

	var obj Accountddd

	err := m.Db.Delete(&obj, id).Error

	if err != nil {
		goa.LogError(ctx, "error deleting Accountddd", "error", err.Error())
		return err
	}

	return nil
}
