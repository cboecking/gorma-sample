package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("accountbbb", func() {

	DefaultMedia(Accountccc)
	BasePath("/accountbbbs")

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("Retrieve all accountbbbs.")
		Response(OK, CollectionOf(Accountccc))
	})

	Action("show", func() {
		Routing(
			GET("/:accountbbbID"),
		)
		Description("Retrieve accountbbb with given id.")
		Params(func() {
			Param("accountbbbID", Integer, "accountbbb ID", func() {
				Minimum(1)
			})
		})
		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Create new accountbbb")
		Payload(func() {
			Member("name")
			Required("name")
		})
		Response(Created, "/accountbbbs/[0-9]+")
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Routing(
			PUT("/:accountbbbID"),
		)
		Description("Change accountbbb name")
		Params(func() {
			Param("accountbbbID", Integer, "accountbbb ID")
		})
		Payload(func() {
			Member("name")
			Required("name")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Routing(
			DELETE("/:accountbbbID"),
		)
		Params(func() {
			Param("accountbbbID", Integer, "accountbbb ID")
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})
})
