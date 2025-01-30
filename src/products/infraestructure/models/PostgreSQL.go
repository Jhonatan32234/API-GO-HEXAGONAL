package models

import "fmt"

type PostgreSQL struct {}

func NewPostgreSQL() *PostgreSQL {
	return &PostgreSQL{}
}

func (ps *PostgreSQL) Save() {
	fmt.Println("[PS] - Producto guardado")
}

func (ps *PostgreSQL) GetAll() {
	fmt.Println("[PS] - Lista de productos")
}