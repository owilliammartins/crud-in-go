package models

import "github.com/MeusApps/usuarios/lib"

//Usuarios representa a tabela de usuários no banco de dados
type Usuarios struct {
	ID    int    `db:"id" json:"id"`
	Nome  string `db:"nome" json:"nome"`
	Email string `db:"email" json:"email"`
}

// UsuarioModel recebe a tabela Usuarios do banco de dados
var UsuarioModel = lib.Sess.Collection("usuarios")
