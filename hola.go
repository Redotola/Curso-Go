package main

import (
	"fmt"
    "rsc.io/quote"
	"math"
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

}