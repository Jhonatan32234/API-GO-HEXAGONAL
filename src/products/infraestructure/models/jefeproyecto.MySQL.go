package models

import (
	"demo/src/core"
	"demo/src/products/domain/entities"
	"fmt"
	"log"
)

type MySQLJP struct {
	conn *core.Conn_MySQL
}

func NewMySQLJP() *MySQLJP {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &MySQLJP{conn: conn}
}

func (mysql *MySQLJP) Save(jefeproyecto entities.JefeProyecto) error{
	query := "INSERT INTO jefeproyecto (nombrejefe,telefono,correo,salario,aniosexperiencia) VALUES (?, ?,?,?,?)"

	result, err := mysql.conn.ExecutePreparedQuery(query, jefeproyecto.Nombrejefe,jefeproyecto.Telefono,jefeproyecto.Correo,jefeproyecto.Salario,jefeproyecto.Aniosexperiencia)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta: %w", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error al obtener el ID insertado: %w", err)
	}

	jefeproyecto.Idjefeproyecto = int32(lastInsertID)
	log.Printf("%s", string(jefeproyecto.Idjefeproyecto))

	log.Printf("[MySQL] - Trabajador guardado: ID: %d, Nombre: %s, Teléfono: %s, Correo: %s, Salario: %d, Años de Experiencia: %d", jefeproyecto.Idjefeproyecto,jefeproyecto.Nombrejefe,jefeproyecto.Telefono,jefeproyecto.Correo,jefeproyecto.Salario,jefeproyecto.Aniosexperiencia)

	return nil
}

func (mysql *MySQLJP) GetAll() ([]entities.JefeProyecto, error) {
	query := "SELECT * FROM jefeproyecto"
	rows := mysql.conn.FetchRows(query)
	defer rows.Close()
	var jefeproyectos []entities.JefeProyecto
    for rows.Next() {
        var jefeproyecto entities.JefeProyecto
		log.Printf("%+v", jefeproyecto)
        if err := rows.Scan(&jefeproyecto.Idjefeproyecto, &jefeproyecto.Nombrejefe,&jefeproyecto.Telefono,&jefeproyecto.Correo,&jefeproyecto.Salario,&jefeproyecto.Aniosexperiencia); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %w", err)
        }
        jefeproyectos = append(jefeproyectos, jefeproyecto)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterando sobre las filas: %w", err)
    }

    return jefeproyectos, nil
}


func (mysql *MySQLJP) Update(jefeproyecto entities.JefeProyecto) error {
	query := "UPDATE jefeproyecto SET nombrejefe=?,telefono=?,correo=?,salario=?,aniosexperiencia=? WHERE idjefeproyecto = ?"

	_, err := mysql.conn.ExecutePreparedQuery(query, jefeproyecto.Nombrejefe,jefeproyecto.Telefono,jefeproyecto.Correo,jefeproyecto.Salario,jefeproyecto.Aniosexperiencia, jefeproyecto.Idjefeproyecto)
	if err != nil {
		return fmt.Errorf("error al actualizar el jefeproyecto: %w", err)
	}

	log.Printf("[MySQL] - JefeProyecto actualizado: ID: %d, Nombre: %s, Teléfono: %s, Correo: %s, Salario: %d, Años de Experiencia: %d", jefeproyecto.Idjefeproyecto,jefeproyecto.Nombrejefe,jefeproyecto.Telefono,jefeproyecto.Correo,jefeproyecto.Salario,jefeproyecto.Aniosexperiencia)
	return nil
}



func (mysql *MySQLJP) Delete(jefeproyectoID int32) error {
	query := "DELETE FROM jefeproyecto WHERE idjefeproyecto = ?"

	_, err := mysql.conn.ExecutePreparedQuery(query, jefeproyectoID)
	if err != nil {
		return fmt.Errorf("error al eliminar el JefeProyecto: %w", err)
	}

	log.Printf("[MySQL] - JefeProyecto eliminado: ID: %d", jefeproyectoID)
	return nil
}
