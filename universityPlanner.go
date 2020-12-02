package main

import( 
  "fmt"
)


//GRAFOS DIRIGIDOS Y NO VALORADOS
/*Son las estructuras que van a permitir definir el grafo*/
// estado sera 1 cuando se ha aprobado el curso y se pone cero cuando no se ha aprobado aun
type nodoGrafo struct{
	nivel string
  	codigo string
  	nombre string
  	credito float64
  	estado int
	aristaEntrada []string
	aristaSalida []string
}

type Grafo struct{
	Node []*nodoGrafo
}

//CREA EL GRAFO VACÍO
func CrearGrafo() *Grafo{
	return &Grafo{
		Node: []*nodoGrafo{},
	}
}
//AÑADE UN NODO AL GRAFO EL CUAL NO PRESENTA ARISTAS DE ENTRADA O SALIDA
func (g *Grafo) AddNode(codigo string,nivel string,nombre string,credito float64){
	
	g.Node = append(g.Node,&nodoGrafo{
		nivel: nivel,
    codigo: codigo,
    nombre: nombre,
    credito: credito,
		aristaEntrada: make([]string,0),
		aristaSalida: make([]string,0),
	})

	return 
}

//AÑADE UNA ARISTA AL GRAFO PARTIENDO DEL NODO n1 HACIA EL NODO n2 EN EL SENTIDO MENCIONADO
func(g *Grafo)AddEdge(n1 string,n2 string)(flag1,flag2 int){
	
	flag1 = 0
	flag2 = 0

	for i := range g.Node{
		
			if(g.Node[i].codigo == n1){
				flag1 = 1
			}else if(g.Node[i].codigo == n2){
				flag2 =1
			}
	}
	if(flag1==0 || flag2==0 ){
		panic("El nodo no se encuentra\n")

	}
	for i := range g.Node{
		
			if(g.Node[i].codigo == n1){
				g.Node[i].aristaSalida = append(g.Node[i].aristaSalida,n2)
			}else if(g.Node[i].codigo == n2){
				g.Node[i].aristaEntrada = append(g.Node[i].aristaEntrada,n1)
			}
	}
	return
}
//IMPRIME LOS NODOS DEL GRAFO
func(g*Grafo)imprimirNodos(){
  var flag int 
  flag=0
	for i := range g.Node{
    credito := fmt.Sprint(g.Node[i].credito)
    if(flag==0){
      fmt.Println("------------Nivel",g.Node[i].nivel,"------------")
      flag=1
    }
		fmt.Println(g.Node[i].codigo, g.Node[i].nivel,g.Node[i].nombre,credito)
    if(g.Node[i].codigo=="1LIN15" || g.Node[i].codigo=="1PSI02"||g.Node[i].codigo=="INF134"){
      flag=0
    }
  }
}
func (grafo1 *Grafo) f(){
  fmt.Println("///////////(////////////////////*******/////////////////*///////****////////////")
  fmt.Println("//////////////////////////////////******///////////////////////*****////////////")
  fmt.Println("(((////((/////////////*****/*/////*******////////////////////**************/////")
  fmt.Println("(////((((((((////////********************///////(//////////////////******///////") 
  fmt.Println("//////((((((((/////****////**********,,,..,*****///////////////////***//////////")
  fmt.Println("////(((((((((((((/////////////******,,,,,,,...,*****////////////////////////////")
  fmt.Println("((((((((((((((((//////////////((((((((/////*********,,,*////////////////////***/")
  fmt.Println("(((((((((//////////////////((((((((//////******///*////*,*////////////////****//")
  fmt.Println("////((((/(((//(/////////((####((((((//(######(//(((#(((/,**((///////////***//***")
  fmt.Println("//////////////////////(#####%###(((((((((////((#%%##((##(#(/(#////////**********")
  fmt.Println("/////////////////////(#%%#############%%%#####(((/////(/###((#(////////*********")
  fmt.Println("///////***///////////(%%%#((/********************//////(/##(###////////*********")
  fmt.Println("/////////////////////((#(((//**/******************//////((##(%%(//////////*/////")
  fmt.Println("/////////////////////(((((//////(#%&%#((//**/////((///(((#%%#%%(//////**////////")
  fmt.Println("/////////////////////((((////(((#%%%##((//*/(###%%###((((#&&%%#/////////////////")
  fmt.Println("///////((////////////////////////(((((//****(##(%&@&#####%&%%#//((////*****/////")
  fmt.Println("/(((((/////////////////(////******//*******/((/////(###(##%#(//////////**///////")
  fmt.Println("(///(((((((((((////////*/////***************//(/**/////((##((((((////**/////////")
  fmt.Println("((((((((((((((((((((///////********/(/////*/((#/***///((##(/////////////////////")
  fmt.Println("(((((((((((((((((/////*///////////(/*///(#((((((//*//((##(////*////////////////(")
  fmt.Println("((((((((((((((((///////////////((/******///////(((/(((##(/////*///////////////((")
  fmt.Println("(((((((((((((//////////,*//////**///(((((####%#((/(((##(/////***////////////((((")
  fmt.Println("(((##((((((///////////..,(////////*//////(((((/////(##(////////////////////(((((")
  fmt.Println("(((((((((((/////////*....,#////////*****////////(((##(/////*****/////////(((((((")
  fmt.Println("((((((((((//////,. ........*((((////****////(((####/*///////////////**///(((((((")
  fmt.Println("//(((((((//*.     ...........*(((((((((((((((####(#(/,,*/////////////((((((###((")
  fmt.Println("/(((((/,       .  .............,/((((((########(((#//*,..  ./////(///((((((((((#")
  fmt.Println("((/,              ................./((/((####(((##(//**.,,.     ,/((((((((((((((")
  fmt.Println(",               .................. ,,/#((##((((((#*,*,,.  ..  .     ,/((((((((((")
  fmt.Println("              ..  .............. ,*..,#(((((((///(,,.,...  .,....      *((((((((")
  fmt.Println(".      ..........  .............,.,,./((////////((,,/*..... .......      *#(((((")
  fmt.Println(" ...    .......................,(,..,((//////////,***,,..... ........ ... ,(((((")
  fmt.Println("   ..... .......................,,/,,/(((((///(*,*,,,,,...... ..........  .,((((")
  fmt.Println(".................................,.,**,*(/(/((,,,,.,..,.................,...,(##")
  fmt.Println("...................................,,,,,,*(/#*,....,..,................,,.,,,,##")
  fmt.Println(".........,..........................,,..,,,(/,....,,...................,*,,,,,./")
  fmt.Println("..........,.,........................,,..,.,,,...,,...................,,*,,,*,,,")
  fmt.Println("........,,,,,.........................,...,,....,,................,,,..,*,,,*,,,")
  fmt.Println("........,,,,,....,,....................,...,...,,................,,,,,,**,,**,,,")

}

func (grafo1 *Grafo) c(cantJala int ){
  if(cantJala<=10){
    fmt.Println("USTED TIENE CONDICION DE ALUMNO")
  }else{
    fmt.Println("USTED ESTA ELIMINADO")
  }
}

func (grafo1 *Grafo) a(){

}
func (grafo1 *Grafo) b(){
  
}
func (grafo1 *Grafo) d(){
  
}
func (grafo1 *Grafo) e(){
  
}


func (grafo1 *Grafo) AddCourses(){

  // First Cicle
  grafo1.AddNode("1MAT04","1","ALGEBRA MATRICIAL Y GEOMETRIA ANALITICA",4.50)
  grafo1.AddNode("1MAT05","1","CALCULO",4.50)
  grafo1.AddNode("1FIS01","1","FISICA",3.50)
  grafo1.AddNode("1QUI01","1","QUIMICA",3.50)
  grafo1.AddNode("1QUI02","1","LABORATORIO DE QUIMICA",0.50)
  grafo1.AddNode("1LIN15","1","COMUNICACION ACADEMICA",3.00)


  // Second Cicle
  grafo1.AddNode("1MAT06","2","CALCULO DIFERENCIAL",4.50)
  grafo1.AddNode("1FIS02","2","FISICA 1",4.50)
  grafo1.AddNode("1FIS03","2","LABORATORIO DE FISICA 1",0.50)
  grafo1.AddNode("1ING02","2","DIBUJO EN INGENIERIA",4.50)
  grafo1.AddNode("1LIN16","2","TRABAJO ACADEMICO",3.00)
  grafo1.AddNode("1FIL01","2","CIENCIA Y FILOSOFIA",3.00)
  grafo1.AddNode("1PSI02","2","MOTIVACION Y LIDERAZGO",2.00)
  
  //THIRD CYCLE
  grafo1.AddNode("1MAT07","3","CALCULO INTEGRAL",4.50)
  grafo1.AddNode("1MAT08","3","CALCULO EN VARIAS VARIABLES",4.50)
  grafo1.AddNode("1FIS04","3","FISICA 2",4.50)
  grafo1.AddNode("1FIS05","3","LABORATORIO DE FISICA 2",0.50)
  grafo1.AddNode("1INF01","3","FUNDAMENTOS DE PROGRAMACION",3.00)
  grafo1.AddNode("INF134","3","ESTRUCTURAS DISCRETAS",4.50)

  //FOURTH CYCLE
  grafo1.AddNode("1MAT09","4","CALCULO APLICADO",4.50)
  grafo1.AddNode("1FIS06","4","FISICA 3",4.50)
  grafo1.AddNode("1FIS07","4","LABORATORIO DE FISICA 3",0.50)
  grafo1.AddNode("INF144","4","TECNICAS DE PROGRAMACION",5.00)
  grafo1.AddNode("1SOC01","4","SOCIOLOGIA",3.00)
  grafo1.AddNode("CDR123","4","PENSAMIENTO CRISTIANO",3.50)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main(){
	//var credito float64;
  	//INFORMACIÓN GUARDADA EN LA BASE DE DATOS DE LA PUCP
	var nivelmax int
  var jalado string
  var cantJala int
  var cantNoLlevado int
  var codigo string
  var list []string
  var opcion string

	grafo1 :=CrearGrafo()
	grafo1.AddCourses()

  fmt.Println("----------Bienvenido al programa----------")
  fmt.Println("Malla curricular: INGENIERÍA INFORMATICA")
	grafo1.imprimirNodos()
	fmt.Println("¿Cual es el nivel maximo de un curso que ha llevado?")
	fmt.Scanf("%d", &nivelmax) 

  // de esta manera tendremos hasta el nivel maximo d
  fmt.Println("¿Ha jalado usted algun curso?(s/n)")
  fmt.Scanf("%s", &jalado)
  if(jalado=="s"){
    fmt.Println("¿Cuantos cursos ha jalado y aun no ha aprobado (CONTAR SOLO DESDE EL NIVEL MAXIMO HACIA ATRAS)")
    fmt.Scanf("%d", &cantJala)
    fmt.Println("¿Cuantos cursos aun no ha llevado(CONTAR SOLO DESDE EL NIVEL MAXIMO HACIA ATRAS)")
    fmt.Scanf("%d", &cantNoLlevado)
    fmt.Println("Ingrese el codigo del/los cursos que ha jalado")
    for i := 0; i < cantJala+cantNoLlevado; i++ {
        fmt.Scanf("%s", &codigo)
        list = append(list, codigo)
  	}
    fmt.Println("Procesando.......")
  }else{
    fmt.Println("Procesando.......")
  } 

  fmt.Println("-------------------OPCIONES A CALCULAR--------------- ")
  fmt.Println("OPCION A: CONSULTA DE UN CURSO QUE DESEA LLEVAR  ")
  fmt.Println("OPCION B: CONDICION DE EGRESO")
  fmt.Println("OPCION C: CONDICION DE ALUMNO")
  fmt.Println("OPCION D: CANTIDAD DE CREDITOS APROBADOS")
  fmt.Println("OPCION E: CURSOS FALTANTES PARA LLEVAR")
  fmt.Println("OPCION F: ¿SOY DIGNO?(¿APROBÓ TP?)")
  fmt.Println("¿Que desea calcular?: ")
  fmt.Scanf("%s", &opcion)

  switch os := opcion; os {
	case "A":
		grafo1.a()

    // la opcion A desencadenara la funcion que les hable ayer, esa que consiste en que chequee si es que los cursos requisitos ( contenidos en la lista de entrada) y si es que uno de esos cursos requisitos no esta aprobado, entonces lo añadimos dentro deuna lista  y mandamos a  este curso de vuelta a la funcion A , de esta manera tambien chequeando si es que  a este curso le faltan otros cursos que aprobar para poder llevarse ,de ser el caso los añade , sino ,solo terminaria y ya , devolviendo un mensaje de " usted puede lllevar el curso D " , solo si es que la lista de cursos faltantes esta vacia , si es que no lo esta , entonces  nos salta una repsuesta de " le faltan los siguientes cursos " e imprimiendo la lista completa que se encontro en la funcion A
	case "B":
		grafo1.b()
  case "C":
		grafo1.c(cantJala)
	case "D":
		grafo1.d()
  case "E":
		grafo1.e()
	case "F":
		grafo1.f()
	default:
	  fmt.Println("CHAU RUSO")
	}
}
