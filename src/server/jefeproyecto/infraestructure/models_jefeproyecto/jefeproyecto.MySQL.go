package models_jefeproyecto

import (
	"demo/src/core"
	"demo/src/server/jefeproyecto/domain/entities"
	"fmt"
	"log"
	"time"
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


func (mysql *MySQLJP) GetExperience(jefeproyectoID int32) (int32, error) {
    var experienciaActual int32
    query := "SELECT aniosexperiencia FROM jefeproyecto WHERE idjefeproyecto = ?"

    row := mysql.conn.DB.QueryRow(query, jefeproyectoID)
    if err := row.Scan(&experienciaActual); err != nil {
        return 0, fmt.Errorf("error al escanear la fila: %w", err)
    }

    timeout := time.After(30 * time.Second)
    ticker := time.NewTicker(2 * time.Second) 
    defer ticker.Stop()

    for {
        select {
        case <-timeout:
            return experienciaActual, nil 

        case <-ticker.C:
            var nuevaExperiencia int32
            row := mysql.conn.DB.QueryRow(query, jefeproyectoID)
            if err := row.Scan(&nuevaExperiencia); err != nil {
                return 0, fmt.Errorf("error al escanear la fila: %w", err)
            }

            if nuevaExperiencia != experienciaActual {
                return nuevaExperiencia, nil 
            }
        }
    }
}
