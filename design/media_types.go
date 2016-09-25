package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Accountccc is the accountbbb resource media type.
var Accountccc = MediaType("application/vnd.accountbbb+json", func() {
	Description("A Accountccc")
	Attributes(func() {
		Attribute("id", Integer, "ID of accountbbb", func() {
			Example(1)
		})
		Attribute("href", String, "API href of accountbbb", func() {
			Example("/accountbbb/1")
		})
		Attribute("name", String, "Name of accountbbb", func() {
			Example("test")
		})
		Attribute("created_at", DateTime, "Date of creation")
		Attribute("created_by", String, "Email of accountbbb owner", func() {
			Format("email")
			Example("me@moeboe.io")
		})

		Required("id", "href", "name", "created_at", "created_by")
	})

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("name")
		Attribute("created_at")
		Attribute("created_by")
	})

	View("tiny", func() {
		Description("tiny is the view used to list accounts")
		Attribute("id")
		Attribute("href")
		Attribute("name")
	})

	View("link", func() {
		Attribute("id")
		Attribute("href")
	})
})
