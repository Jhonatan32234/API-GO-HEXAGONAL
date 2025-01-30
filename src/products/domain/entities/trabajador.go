package entities


type Trabajador struct {
	Idtrabajador      int32   
    Nombretrabajador  string  
    Posicion          string  
    Telefono          string  
    Correo            string 
    Salario           int32 
    Aniosexperiencia  int32   
}

func NewTrabajador(nombretrabajador string,posicion string,telefono string,correo string,salario int32,aniosexperiencia int32) *Trabajador{
	return &Trabajador{Idtrabajador: 1, Nombretrabajador: nombretrabajador, Posicion: posicion, Telefono: telefono,Correo: correo,Salario: salario,Aniosexperiencia: aniosexperiencia}
}
