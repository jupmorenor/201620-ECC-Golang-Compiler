//Juan Pablo Moreno Rico - 20111020059
package main

import (
  "fmt"
  "strconv"
  "strings"
)

const CARACTERES_VALIDOS = "0123456789+-*/%"

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
  RecorrerPreorden(t.Izquierda)
  RecorrerPreorden(t.Derecha)
}

/**
 * Recorrido del arbol postorden
 */
func RecorrerPostorden(t *Arbol) {
  if t == nil {
    return
  }
  RecorrerPostorden(t.Izquierda)
  RecorrerPostorden(t.Derecha)
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

/**
 * validacion del arbol
 */
func EsArbolValido(t *Arbol) bool { 
	if t == nil {
		return true
	} else {
		for i := range t.Valor { 
			//fmt.Println(t.Valor[i], string(t.Valor[i]))
			if !strings.Contains(CARACTERES_VALIDOS, string(t.Valor[i])) {
				return false
			}
		}
		return EsArbolValido(t.Izquierda) && EsArbolValido(t.Derecha)
	}
}

func main() {
  t1 := &Arbol{&Arbol{&Arbol{&Arbol{nil, nil, "120"}, &Arbol{nil, nil, "70"}, "+"}, &Arbol{&Arbol{nil, nil, "150"}, &Arbol{nil, nil, "30"}, "-"}, "*"}, &Arbol{&Arbol{nil, nil, "140"}, &Arbol{nil, nil, "50"}, "%"}, "/"}
  //(120+70)*(150-30)/(140%50) = 570 
  if EsArbolValido(t1) {
	  fmt.Println("El arbol es valido")
	  RecorrerInorden(t1)
	  fmt.Print(" =  ")
	  fmt.Println(OperarArbol(t1), " <- Resultado de evaluar el arbol")
  } else {
	  fmt.Println("El arbol ingresado NO es valido")
  }
}
