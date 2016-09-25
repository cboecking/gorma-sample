package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var _ = StorageGroup("cellar-stor", func() {
	Description("This is the global storage group")
	Store("postgres", gorma.Postgres, func() {
		Description("This is the Postgres relational store")
		Model("Accountddd", func() {
			RendersTo(Accountccc)
			Description("cellar-aaa Accountddd")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			Field("name", gorma.String)
		})

	})
})
