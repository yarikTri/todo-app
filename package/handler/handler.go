package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yarikTri/todo-app/package/service"
)

// Handler describes a multiplexer for the app
type Handler struct {
	services *service.Service
}

// NewHandler basicly inits new Handler
func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// InitRoutes does exactly what it supposed to do
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListByID)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/items/", h.createItem)
				items.GET("/items/", h.getAllItems)
				lists.GET("/items/:item_id", h.getItemByID)
				lists.PUT("/items/:item_id", h.updateItem)
				lists.DELETE("/items/:item_id", h.deleteItem)
			}
		}
	}

	return router
}
