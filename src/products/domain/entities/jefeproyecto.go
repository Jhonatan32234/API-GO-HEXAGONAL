package entities


type JefeProyecto struct {
	Idjefeproyecto int32
	Nombrejefe string
    Telefono string
    Correo string
    Salario int32
    Aniosexperiencia int32
}

func NewJefeProyecto(nombrejefe string,telefono string,correo string,salario int32,aniosexperiencia int32) *JefeProyecto{
	return &JefeProyecto{Idjefeproyecto: 1, Nombrejefe: nombrejefe, Telefono: telefono,Correo: correo,Salario: salario,Aniosexperiencia: aniosexperiencia}
}
