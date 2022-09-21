package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("deployments", func() {
	Title("Api for deployments")
	Version("0.0.1")
	Server("http", func() {
		Host("development", func() {
			URI("http://localhost:8080")
		})
	})
})

var Parameter = Type("Parameter", func() {
	Attribute("name", String, func() {
		MinLength(1)
	})
	Attribute("source", func() {
		Enum("static", "generated", "configurable")
	})
	Attribute("type", func() {
		Enum("string", "int", "bool", "password")
	})
	Required("name", "source", "type")
})

var ParameterUsage = Type("ParameterUsage", func() {
	Attribute("name", String, func() {
		MinLength(1)
	})
	Attribute("environment", ArrayOf(String))
	Attribute("files", ArrayOf(String))
})

var Container = Type("Container", func() {
	Attribute("name", String)
	Attribute("image", String)
	Attribute("usedParams", ArrayOf(ParameterUsage))
	Attribute("services", ArrayOf(String))
	Required("name", "image")
})

var Deployment = Type("Deployment", func() {
	Attribute("name", String)
	Attribute("containers", ArrayOf(Container), func() {
		MinLength(1)
	})
	Attribute("params", ArrayOf(Parameter))
	Required("name", "containers")
})

var _ = Service("deployments", func() {
	Description("used to manage active deployments")

	Method("upsert", func() {
		Payload(Deployment)
		HTTP(func() {
			POST("/api/deployments")
			Response(204)
		})
	})

	Method("list", func() {
		Result(ArrayOf(String))
		HTTP(func() {
			GET("/api/deployments")
			Response(200)
		})
	})

	Method("get", func() {
		Payload(String, "Name of the Deployment")
		Result(Deployment)
		Error("DeploymentNotFound")
		HTTP(func() {
			GET("/api/deployments/{deploymentName}")
			GET("/api/deployments/{deploymentName}/view")
			Response(200)
			Response("DeploymentNotFound", 404)
		})
	})

	Method("delete", func() {
		Payload(String, "Name of the Deployment")
		Error("DeploymentNotFound")
		HTTP(func() {
			GET("/api/deployments/{deploymentName}/delete")
			Response("DeploymentNotFound", 404)
		})
	})

	Files("/openapi3.json", "./gen/http/openapi3.json")
})
