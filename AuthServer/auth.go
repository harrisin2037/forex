package AuthServer

import (
	"forex/library/document"
	"forex/starter"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthServer struct {
	*starter.Content
	apis []string
}

func (m *AuthServer) Starter() error {
	for _, model := range ModelAddrs() {
		m.Mysql.AutoMigrateByAddr(model)
	}
	if m.App.DebugMode {
		m.debugModeData()
	}
	m.TemplateRoutes()
	m.Routes()
	m.Server.Starter(m.Content)
	return nil
}

func (m *AuthServer) Routes() (routes []gin.IRoutes) {

	group := m.Server.Engine.Group("/api")

	routes = []gin.IRoutes{
		m.User(group),
	}
	return
}

func (m *AuthServer) TemplateRoutes() (routes []gin.IRoutes) {

	group := m.Server.Engine.Group("/template")

	routes = []gin.IRoutes{
		m.TemplateHttpConnectionCheck(group),
		m.TemplateUsers(group),
	}
	return
}

func (m *AuthServer) TemplateHttpConnectionCheck(group *gin.RouterGroup) gin.IRoutes {
	type response struct {
		Status string
	}
	resp := response{Status: "Connection successful"}
	doc := document.Doc{
		FilePath:   "",
		API:        "/connectionCheck",
		Method:     "GET",
		StatusCode: http.StatusOK,
		Response:   resp,
		Handler: func(context *gin.Context) {
			context.JSON(http.StatusOK, resp)
		},
	}
	return group.Handle(doc.ToJSONResponse())
}

func (m *AuthServer) TemplateUsers(group *gin.RouterGroup) gin.IRoutes {
	type response struct {
		Users []User
	}
	resp := response{
		Users: []User{
			User{
				Username: "template user 1",
				ParentID: 0,
			},
			User{
				Username: "template user 2",
				ParentID: 0,
			},
		},
	}
	doc := document.Doc{
		FilePath:   "",
		API:        "/user",
		Method:     "GET",
		StatusCode: http.StatusOK,
		Response:   resp,
		Handler: func(context *gin.Context) {
			context.JSON(http.StatusOK, resp)
		},
	}
	return group.Handle(doc.ToJSONResponse())
}

func (m *AuthServer) User(group *gin.RouterGroup) gin.IRoutes {

	return group.Handle("GET", "/user", func(c *gin.Context) {

		defer m.Mysql.Connector()()

		req := struct {
			Unit string `json:"unit"`
		}{}

		resp := struct {
			Message string `json:"message"`
			Users   []User `json:"users"`
		}{}

		c.BindQuery(&req)

		m.Mysql.DB.Find(&resp.Users)
		// m.HierachicalUsersCTE()

		c.JSON(http.StatusOK, gin.H{
			"Response": resp,
		})
	})
}

func (m *AuthServer) UserID(group *gin.RouterGroup) gin.IRoutes {

	return group.Handle("GET", "/user/:id", func(c *gin.Context) {

		defer m.Mysql.Connector()()

		req := struct {
			Unit string `json:"unit"`
		}{}

		resp := struct {
			Message string `json:"message"`
			Users   []User `json:"users"`
		}{}

		c.BindQuery(&req)

		m.Mysql.DB.Where("", nil).Find(&resp.Users)
		// m.HierachicalUsersCTE()

		c.JSON(http.StatusOK, gin.H{
			"Response": resp,
		})
	})
}
