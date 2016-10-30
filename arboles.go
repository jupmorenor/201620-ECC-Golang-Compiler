//Juan Pablo Moreno Rico - 20111020059
package main

import (
  "fmt"
  "strconv"
  "strings"
)

const CARACTERES = "0123456789+-*/%"

/**
 * Estructura de una pila de arboles
 */
type Pila struct {
	datos [] *Arbol
	cantidad int
}

/**
 * Funcion que agrega arboles a la pila
 */
func (pil *Pila) agregarAPila(arbol *Arbol) {
	pil.datos = append(pil.datos[:pil.cantidad], arbol)
	pil.cantidad++
}

/**
 * Funcion que remueve arboles de la pila
 */
func (pil *Pila) removerDePila() *Arbol {
	if pil.cantidad == 0 {
		return nil
	}
	pil.cantidad--
	return pil.datos[pil.cantidad]
}

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
 * Validacion de la correctitud del arbol
 */
func esArbolValido(t *Arbol) bool {
  if t == nil {
    return true
  } else {
    for _, char := range t.Valor {
      if !strings.Contains(CARACTERES, string(char)) {
        return false
      }
    }
    return esArbolValido(t.Izquierda) && esArbolValido(t.Derecha)
  }
}

/**
 * Construye un arbol a partir de una expresion postfija
 */ 
func construirArbol(expresion string) *Arbol {
	pila := &Pila{}
	for _, char := range expresion {
		fmt.Println(string(char), pila.cantidad)
		if strings.Contains(CARACTERES[:10], string(char)) {
			pila.agregarAPila(&Arbol{nil, nil, string(char)})
		} else if strings.Contains(CARACTERES[10:], string(char)) {
			if pila.cantidad >= 2 {
				t1 := pila.removerDePila()
				t2 := pila.removerDePila()
				pila.agregarAPila(&Arbol{t2, t1, string(char)})
			} else {
				return nil
			}
		} else {
			return nil
		}
	}
	return pila.removerDePila()
}

func main() {
  //t1 := &Arbol{&Arbol{&Arbol{&Arbol{nil, nil, "2"}, &Arbol{nil, nil, "3"}, "*"}, &Arbol{&Arbol{nil, nil, "9"}, &Arbol{nil, nil, "3"}, "/"}, "+"}, &Arbol{&Arbol{nil, nil, "6"}, &Arbol{nil, nil, "1"}, "-"}, "+"}
  //(2*3)+(9/3)%(5*1)
  var expr string
  fmt.Scanln(&expr)
  t1 := construirArbol(expr)
  if t1!=nil {
	  RecorrerPreorden(t1)
	  fmt.Println()
	  RecorrerPostorden(t1)
	  fmt.Println()
	  if esArbolValido(t1) {
		RecorrerInorden(t1)
		fmt.Print(" =  ")
		fmt.Println(OperarArbol(t1), " <- Resultado de evaluar el arbol")
	  } else {
		fmt.Println("Arbol invalido")
	  }
  } else {
	  fmt.Println("Expresion incorrecta")
  }
}
