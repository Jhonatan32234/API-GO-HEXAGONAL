package models

import (
	"demo/src/core"
	"demo/src/products/domain/entities"
	"fmt"
	"log"
)

type MySQLP struct {
	conn *core.Conn_MySQL
}

func NewMySQLP() *MySQLP {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &MySQLP{conn: conn}
}

func (mysql *MySQLP) Save(product entities.Product) error{
	query := "INSERT INTO product (name, price) VALUES (?, ?)"

	result, err := mysql.conn.ExecutePreparedQuery(query, product.Name, product.Price)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta: %w", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error al obtener el ID insertado: %w", err)
	}

	product.Id = int32(lastInsertID)
	log.Printf("%s", string(product.Id))

	log.Printf("[MySQL] - Producto guardado: ID: %d, Nombre: %s, Precio: %.2f", product.Id, product.Name, product.Price)
	return nil
}

func (mysql *MySQLP) GetAll() ([]entities.Product, error) {
	query := "SELECT * FROM product"
	rows := mysql.conn.FetchRows(query)
	defer rows.Close()
	var products []entities.Product
    for rows.Next() {
        var product entities.Product
        if err := rows.Scan(&product.Id, &product.Name, &product.Price); err != nil {
            return nil, fmt.Errorf("error al escanear la fila: %w", err)
        }
        products = append(products, product)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterando sobre las filas: %w", err)
    }

    return products, nil
}


func (mysql *MySQLP) Update(product entities.Product) error {
	query := "UPDATE product SET name = ?, price = ? WHERE id = ?"

	_, err := mysql.conn.ExecutePreparedQuery(query, product.Name, product.Price, product.Id)
	if err != nil {
		return fmt.Errorf("error al actualizar el producto: %w", err)
	}

	log.Printf("[MySQL] - Producto actualizado: ID: %d, Nombre: %s, Precio: %.2f", product.Id, product.Name, product.Price)
	return nil
}



func (mysql *MySQLP) Delete(productID int32) error {
	query := "DELETE FROM product WHERE id = ?"

	_, err := mysql.conn.ExecutePreparedQuery(query, productID)
	if err != nil {
		return fmt.Errorf("error al eliminar el producto: %w", err)
	}

	log.Printf("[MySQL] - Producto eliminado: ID: %d", productID)
	return nil
}
