package routers

import (
	"github.com/labstack/echo"

	"github.com/MeusApps/usuarios/controllers"
)

// App é uma instancia de Echo
var App *echo.Echo

func init() {
	App = echo.New()

	//A pagina inicial
	App.GET("/", controllers.Home)
	App.GET("/add", controllers.Add)
	App.GET("/atualizar/:id", controllers.Update)

	api := App.Group("/v1")

	api.POST("/insert", controllers.Inserir)

	api.DELETE("/delete/:id", controllers.Deletar)
	api.PUT("/update/:id", controllers.Atualizar)

}
