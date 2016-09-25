package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// This is the cellar-aaa application API design used by goa to generate
// the application code, client, tests, documentation etc.
var _ = API("cellar-aaa", func() {
	Title("cellar-aaa title")
	Description("cellar-aaa description")
	Contact(func() {
		Name("moeboe team")
		Email("chuck@moeboe.io")
		URL("http://moeboe.io")
	})
	Host("localhost:8081")
	Scheme("http")
	BasePath("/cellar-aaa")

	ResponseTemplate(Created, func(pattern string) {
		Description("Resource created")
		Status(201)
		Headers(func() {
			Header("Location", String, "href to created resource", func() {
				Pattern(pattern)
			})
		})
	})
})
