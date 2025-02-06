package models

import (
	"demo/src/core"
	"demo/src/server/domain/entities"
	"fmt"
	"log"
)

type MySQLT struct {
	conn *core.Conn_MySQL
}

func NewMySQLT() *MySQLT {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &MySQLT{conn: conn}
}

func (mysql *MySQLT) Save(trabajador entities.Trabajador) error {
	query := "INSERT INTO trabajador (nombretrabajador, posicion,telefono,correo,salario,aniosexperiencia) VALUES (?, ?,?,?,?,?)"
	fmt.Printf("Insertando trabajador: %+v\n", trabajador)

	result, err := mysql.conn.ExecutePreparedQuery(query, trabajador.Nombretrabajador, trabajador.Posicion, trabajador.Telefono, trabajador.Correo, trabajador.Salario, trabajador.Aniosexperiencia)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta: %w", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error al obtener el ID insertado: %w", err)
	}

	trabajador.Idtrabajador = int32(lastInsertID)
	log.Printf("%s", string(trabajador.Idtrabajador))

	log.Printf("[MySQL] - Trabajador guardado: ID: %d, Nombre: %s, Posición: %s, Teléfono: %s, Correo: %s, Salario: %d, Años de Experiencia: %d", trabajador.Idtrabajador, trabajador.Nombretrabajador, trabajador.Posicion, trabajador.Telefono, trabajador.Correo, trabajador.Salario, trabajador.Aniosexperiencia)

	return nil
}

func (mysql *MySQLT) GetAll() ([]entities.Trabajador, error) {
	query := "SELECT * FROM prueba.trabajador"
	rows := mysql.conn.FetchRows(query)
	defer rows.Close()
	var trabajadores []entities.Trabajador
	for rows.Next() {
		var trabajador entities.Trabajador
		if err := rows.Scan(&trabajador.Idtrabajador, &trabajador.Nombretrabajador, &trabajador.Posicion, &trabajador.Telefono, &trabajador.Correo, &trabajador.Salario, &trabajador.Aniosexperiencia); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %w", err)
		}
		trabajadores = append(trabajadores, trabajador)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando sobre las filas: %w", err)
	}

	return trabajadores, nil
}

func (mysql *MySQLT) Update(trabajador entities.Trabajador) error {
	query := "UPDATE trabajador SET nombretrabajador=?,posicion=?,telefono=?,correo=?,salario=?,aniosexperiencia=? WHERE idtrabajador = ?"

	_, err := mysql.conn.ExecutePreparedQuery(query, trabajador.Nombretrabajador, trabajador.Posicion, trabajador.Telefono, trabajador.Correo, trabajador.Salario, trabajador.Aniosexperiencia, trabajador.Idtrabajador)
	if err != nil {
		return fmt.Errorf("error al actualizar el trabajador: %w", err)
	}

	log.Printf("[MySQL] - Trabajador actualizado: ID: %d, Nombre: %s, Posición: %s, Teléfono: %s, Correo: %s, Salario: %d, Años de Experiencia: %d", trabajador.Idtrabajador, trabajador.Nombretrabajador, trabajador.Posicion, trabajador.Telefono, trabajador.Correo, trabajador.Salario, trabajador.Aniosexperiencia)
	return nil
}

func (mysql *MySQLT) Delete(trabajadorID int32) error {
	query := "DELETE FROM trabajador WHERE idtrabajador = ?"

	_, err := mysql.conn.ExecutePreparedQuery(query, trabajadorID)
	if err != nil {
		return fmt.Errorf("error al eliminar el trabajador: %w", err)
	}

	log.Printf("[MySQL] - Trabajador eliminado: ID: %d", trabajadorID)
	return nil
}




func (mysql *MySQLT) GetSalary(trabajadorID int32) (int32, error) {
    query := "SELECT salario FROM trabajador WHERE idtrabajador = ?"
    row := mysql.conn.DB.QueryRow(query, trabajadorID)

    var salario int32
    if err := row.Scan(&salario); err != nil {
        return 0, fmt.Errorf("error al escanear la fila: %w", err)
    }

    return salario, nil
}



func (mysql *MySQLT) GetPosition(trabajadorID int32) (string, error) {
    query := "SELECT posicion FROM trabajador WHERE idtrabajador = ?"
    row := mysql.conn.DB.QueryRow(query, trabajadorID)

    var posicion string
    if err := row.Scan(&posicion); err != nil {
        return "", fmt.Errorf("error al escanear la fila: %w", err)
    }

    return posicion, nil
}