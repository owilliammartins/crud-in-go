package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/MeusApps/usuarios/models"
)

// Home é a pagina inicial da minha aplicação
func Home(c echo.Context) error {

	var usuarios []models.Usuarios

	if err := models.UsuarioModel.Find().All(&usuarios); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"mensagem": "Erro no retorno do registro do banco de dados.",
		})
	}

	data := map[string]interface{}{
		"titulo":   "Lista de usuários",
		"usuarios": usuarios,
	}

	return c.Render(http.StatusOK, "index.html", data)

	//return c.String(http.StatusOK, "Tudo OK?")
}

// Add é a pagina para cadastro de um novo usuário
func Add(c echo.Context) error {
	return c.Render(http.StatusOK, "add.html", nil)
}

//Inserir é a chamada para gravar os dados no MySql
func Inserir(c echo.Context) error {
	nome := c.FormValue("nome")
	email := c.FormValue("email")

	var usuario models.Usuarios

	usuario.Nome = nome
	usuario.Email = email

	if nome != "" && email != "" {
		if _, err := models.UsuarioModel.Insert(usuario); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"messagem": "Não foi possível adicionar o registro no banco de dados.",
			})
		}

		return c.Redirect(http.StatusFound, "/")
		/*return c.JSON(http.StatusCreated, map[string]string{
			"messagem": "O registro foi gravado com sucesso!",
		}) */

	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"mensagem": "Os campos devem ser preenchidos",
	})

}

//Deletar é chamada para excluir um usuario do banco de dados
func Deletar(c echo.Context) error {
	usuarioID, _ := strconv.Atoi(c.Param("id"))

	resultado := models.UsuarioModel.Find("id=?", usuarioID)

	//s := models.UsuarioModel.Find("id=?", usuarioID).String()
	//s := models.UsuarioModel.Find("id=?", 99).String(

	if count, _ := resultado.Count(); count < 1 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"mensagem": "Não foi possível encontrar o registro para ser excluído.",
			//"mensagem": fmt.Sprintf(" Não foi possível excluir o registro com o código %d", usuarioID),
			//"mensagem": fmt.Sprintf(" Não foi possível encontrar o registro com o código %s", c.Param("ID")),
		})
	}

	if err := resultado.Delete(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"mensagem": "Não foi possível excluir o registro.",
		})
	}

	return c.JSON(http.StatusAccepted, map[string]string{
		"mensagem": "Registro excluído com sucesso.",
	})
}

//Atualizar é chamada para atualizar os dados do usuário no banco de dados
func Atualizar(c echo.Context) error {
	usuarioID, _ := strconv.Atoi(c.Param("id"))
	nome := c.FormValue("nome")
	email := c.FormValue("email")

	usuario := models.Usuarios{
		ID:    usuarioID,
		Nome:  nome,
		Email: email,
	}

	resultado := models.UsuarioModel.Find("id=?", usuarioID)

	if count, _ := resultado.Count(); count < 1 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"mensagem": "Registro não encontrado.",
		})
	}

	if err := resultado.Update(usuario); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"mensagem": "Erro ao tentar atualizar o cadastro.",
		})

	}

	return c.JSON(http.StatusAccepted, usuario)
}

//Update é chamada para recuperar os dados do usuario para edição no browser
func Update(c echo.Context) error {
	var usuarioID, _ = strconv.Atoi(c.Param("id"))

	var usuario models.Usuarios

	resultado := models.UsuarioModel.Find("id=?", usuarioID)

	if count, _ := resultado.Count(); count < 1 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"mensagem": "Usuario não foi encontrado.",
		})
	}

	if err := resultado.One(&usuario); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"mensagem": "Ocorreu um erro ao tentar atualizar o cadastro.",
		})
	}

	var data = map[string]interface{}{
		"usuario": usuario,
	}

	return c.Render(http.StatusOK, "update.html", data)
}
