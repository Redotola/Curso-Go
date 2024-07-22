package main

import (
	"fmt"
	"math"
	"strconv"

	"rsc.io/quote"
)

// Declaraciones de constantes
const Pi float32 = 3.14

// Declarar varias constantes
const (
	x = 100
	y = 0b1010 // binario
	z = 0o12 // octal
	w = 0xFF // hexadecimal
)

// Definimos constantes que se incrementan con la palabra reservada "iota"
const (
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {

	manejo_fmt()
}

func imprimir_datos(){

		// Imprimir el paquete de Go
		fmt.Println("Hola Mundo")
		fmt.Println(quote.Go())
	
		// Declaracion de variables
		var firstName, lastName string
		var age int
	
		// Otra forma de declarar variables
		var (
			firstName1, lastName1 	string = "Axl", "Castrejón"
			age1 					int	= 21
		)
	
		// Otra forma de declarar
		var (
			firstName2, lastName2, age2 = "Axl", "Castrejón", 21 // Inicializando variables en la misma línea
		)
	
		// Declarar variable nueva sin usar var, solo dentro de las funciones
		firstName3, lastName3 := "Axl", "Castrejón" 
		age3 :=  21
	
		firstName = "Axl"
		lastName = "Castrejón"
		age = 21
	
		fmt.Println(firstName, lastName, age)
		fmt.Println(firstName1, lastName1, age1)
		fmt.Println(firstName2, lastName2, age2)
		fmt.Println(firstName3, lastName3, age3)
		fmt.Println(Pi)
		fmt.Println(x, y, z, w)
	
		/*Tipos de datos*/
	
		// Numericos
		// Cuando la variable se declara "int" el sistema decide el tamaño conforme la arquitectura
		var integer int8 = 12
	
		// Uint solo almacena numeros enteros positivos
		var integeruint uint = 3
	
		// Flotantes 
		var float float64 = 3.34
	
		// El menor y mayor rango que se puede alcanzar con los tipos de datos
		fmt.Println(math.MinInt64, math.MaxInt64)
		fmt.Println(integer, integeruint, float)
		fmt.Println(math.MaxFloat32, math.MaxUint32)
		
		// Formateo de valores en string
		fullName := "Alex Roel \t(alias \"roelcode\")\n"
	
		// Booleanos
		var valueBool bool = false
	
		// Bits
		var a byte = 'a'
	
		// Impresion de contenido
		fmt.Println(fullName, valueBool, a)
	
		s := "hola"
	
		fmt.Println(s[0])
	
		// Running
	
		var r rune = '🖤'
		fmt.Println(r)
}


func manejo_de_datos(){

	// Manejo de conversiones de datos para operaciones
	var integer16 int16 = 50
	var integer32 int32 = 100

	// Impresion de la suma
	fmt.Println(integer16 + int16(integer32))

	// Conversion de string a entero con error controlado
	s := "100"
	i, _ := strconv.Atoi(s)
	fmt.Println("Impresion de  el valor sumado y convertido: ")
	fmt.Println(i + 1)

	// Conversion de entero a strig
	n := 42
	s = strconv.Itoa(n)
	fmt.Println("Conversion de entero a string: ", s)
	fmt.Println("Suma de cadenas: ", s+s)


}

func manejo_fmt(){
	fmt.Println("Hola Mundo")
	fmt.Println(quote.Go())

	name := "Axl"
	age := 22

	fmt.Printf("Hola, me llamo %s y tengo %d años. \n", name, age)
	
	greeting := fmt.Sprintf("Hola, me llamo %s y tengo %d años. \n", name, age)
	fmt.Println(greeting)

	var name1, name2 string
	var age1 int

	fmt.Printf("Ingrese su nombre: ")
	fmt.Scanln(&name1, &name2)
	fmt.Printf("Ingrese su edad: ")
	fmt.Scanln(&age1)

    // Formateo con placeholder %s para strings y %d para enteros
	fmt.Printf("Hola, me llamo %s %s y tengo %d años.\n", name1, name2, age1)

	// Solo en caso donde el dato no se especifique
	fmt.Printf("Hola, me llamo %v y tengo %v años \n", name1, age1)

	// Imprimir el tipo de dato que es
	fmt.Printf("El tipo de la variable name1 es: %T y de la variable age1 es: %T \n", name1, age1)
}