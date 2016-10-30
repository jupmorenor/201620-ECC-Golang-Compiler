package main

import (
  "fmt"
  "strconv"
)

/**
 * Estructura de un arbol 
 */
type Arbol struct {
  Izquierda *Arbol
  Derecha *Arbol
  Valor string
}

/**
 * Recorrido del arbol inorden
 */
func RecorrerInorden(t *Arbol) {
  if t == nil {
    return
  }
  RecorrerInorden(t.Izquierda)
  fmt.Print(t.Valor, " ")
  RecorrerInorden(t.Derecha)
}

/**
 * Recorrido del arbol preorden
 */
func RecorrerPreorden(t *Arbol) {
  if t == nil {
    return
  }
  fmt.Print(t.Valor, " ")
  RecorrerInorden(t.Izquierda)
  RecorrerInorden(t.Derecha)
}

/**
 * Recorrido del arbol postorden
 */
func RecorrerPostorden(t *Arbol) {
  if t == nil {
    return
  }
  RecorrerInorden(t.Izquierda)
  RecorrerInorden(t.Derecha)
  fmt.Print(t.Valor, " ")
}

/**
 * Operacion del arbol inorden
 */
func OperarArbol(t *Arbol) int{
	if t == nil {
		return 0
	}
	switch t.Valor{
	case "+":
		return OperarArbol(t.Izquierda) + OperarArbol(t.Derecha)
	case "-":
		return OperarArbol(t.Izquierda) - OperarArbol(t.Derecha)
	case "*":
		return OperarArbol(t.Izquierda) * OperarArbol(t.Derecha)
	case "/":
		return OperarArbol(t.Izquierda) / OperarArbol(t.Derecha)
	case "%":
		return OperarArbol(t.Izquierda) % OperarArbol(t.Derecha)
	default:
		var val, _ = strconv.Atoi(t.Valor)
		return val
  }
}

func main() {
  t1 := &Arbol{&Arbol{&Arbol{&Arbol{nil, nil, "2"}, &Arbol{nil, nil, "3"}, "*"}, &Arbol{&Arbol{nil, nil, "9"}, &Arbol{nil, nil, "3"}, "/"}, "+"}, &Arbol{&Arbol{nil, nil, "6"}, &Arbol{nil, nil, "1"}, "-"}, "%"}
  //(2*3)+(9/3)%(5*1)
  RecorrerInorden(t1)
  fmt.Print(" =  ")
  fmt.Println(OperarArbol(t1), " <- Resultado de evaluar el arbol")
}
