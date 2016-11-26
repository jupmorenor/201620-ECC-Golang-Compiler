//Juan Pablo Moreno Rico - 20111020059
package main

import (
  "fmt"
  "strconv"
  "strings"
  "bufio"
  "os"
  "regexp"
)

const CARACTERES string = " 0123456789+-*/%"
const ASIGNADOR string = ":=" 

/**
 * Estructura de una pila de tokens 
 */
type pilaTokens struct {
	contenido [] *Token
	cantidad int
}

/**
 * Estructura de un token
 */
type Token struct {
	tipo string
	simbolo string
}

/**
 * Funcion que agrega tokens a la pila
 */
func (pil *pilaTokens) agregarAPila(token *Token) {
	pil.contenido = append(pil.contenido[:pil.cantidad], token)
	pil.cantidad++
}

/**
 * Funcion que remueve tokens de la pila
 */
func (pil *pilaTokens) removerDePila() *Token {
	if pil.cantidad == 0 {
		return nil
	}
	pil.cantidad--
	return pil.contenido[pil.cantidad]
}

/**
 * Imprime un token
 */
func (t *Token) String() string {
	return fmt.Sprint(t.tipo, " -> ", t.simbolo)
}

/**
 * Estructura de una cola de expresiones
 */
type Expresiones struct {
    nodes []*string
    size  int
    head  int
    tail  int
    count int
}

/**
 * Funcion que agrega expresiones a la cola
 */
func (q *Expresiones) agregarExpresion(n *string) {
    if q.head == q.tail && q.count > 0 {
        nodes := make([]*string, len(q.nodes)+q.size)
        copy(nodes, q.nodes[q.head:])
        copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
        q.head = 0
        q.tail = len(q.nodes)
        q.nodes = nodes
    }
    q.nodes[q.tail] = n
    q.tail = (q.tail + 1) % len(q.nodes)
    q.count++
}

/**
 * Funcion que remueve expresiones de la pila
 */
func (q *Expresiones) removerExpresion() *string {
    if q.count == 0 {
        return nil
    }
    node := q.nodes[q.head]
    q.head = (q.head + 1) % len(q.nodes)
    q.count--
    return node
}

func crearColaExpresiones(size int) *Expresiones {
    return &Expresiones{
        nodes: make([]*string, size),
        size:  size,
    }
}

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
	elementos := strings.Split(expresion, " ")
	for _, char := range elementos {
		if strings.ContainsAny(CARACTERES[:11], string(char)) {
			pila.agregarAPila(&Arbol{nil, nil, string(char)})
		} else if strings.Contains(CARACTERES[11:], string(char)) {
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

/**
 * Construye una expresion a partir de una pila de subexpresiones que incluyen variables
 */
func construirExpresion(expresiones *Expresiones) string{
	var exprFinal string
	for {
		expr1 := *expresiones.removerExpresion()
		if !strings.HasSuffix(expr1, ASIGNADOR) {
			exprFinal = expr1
			break
		}
		expresiones.agregarExpresion(&expr1)
	}
	for expresiones.count > 0 {
		expr := *expresiones.removerExpresion()
		if strings.HasSuffix(expr, ASIGNADOR) {
			varName := string(expr[len(expr)-4])
			if strings.Contains(exprFinal, varName) {
				exprFinal = strings.Replace(exprFinal, varName, string(expr[:len(expr)-5]), -1)
				expresiones.agregarExpresion(&expr)
			} else {
				expresiones.agregarExpresion(&expr)
			}
		} else {
			expresiones.agregarExpresion(&expr)
		}
		if esExpresionFinal(exprFinal) {
			break
		}
	}
	
	return exprFinal
}

func esExpresionFinal(expr string) bool {
	for _, char := range expr {
		if !strings.Contains(CARACTERES, string(char)) {
			return false
		}
	}
	return true
}

func generarTablaSim(pila *pilaTokens, expr string) {
	expresion := strings.Split(expr, " ")
	for _, char := range expresion {
		switch {
			case strings.ContainsAny(CARACTERES[:11], string(char)):
				pila.agregarAPila(&Token{"VAL", string(char)})
			case strings.Contains(CARACTERES[11:], string(char)):
				pila.agregarAPila(&Token{"OP", string(char)})
			case strings.Contains(ASIGNADOR, string(char)):
				pila.agregarAPila(&Token{"OP", string(char)})
			default:
				pila.agregarAPila(&Token{"ID", string(char)})
		}
	}
}

func main() {
  //t1 := &Arbol{&Arbol{&Arbol{&Arbol{nil, nil, "2"}, &Arbol{nil, nil, "3"}, "*"}, &Arbol{&Arbol{nil, nil, "9"}, &Arbol{nil, nil, "3"}, "/"}, "+"}, &Arbol{&Arbol{nil, nil, "6"}, &Arbol{nil, nil, "1"}, "-"}, "+"}
  //(2*3)+(9/3)%(5*1)
  rIDs, _ := regexp.Compile("^[a-zA-Z]+[a-zA-Z0-9]*")
  rVal, _ := regexp.Compile("^[0-9]+$")
  rOps, _ := regexp.Compile("^[\\+|\\-|\\*|\\/|\\%]|:=$")
  expresiones := crearColaExpresiones(100)
  tablaSimbolos := &pilaTokens{}
  pilaAux := &pilaTokens{}
  var token *Token
  var continuar string = "S" 
  reader := bufio.NewReader(os.Stdin)
  for continuar == "S" {
	  var expr string
	  fmt.Print("Ingrese una expresion: ")
	  expr, _ = reader.ReadString('\n')
	  expr = strings.TrimSpace(expr) // \r\n -> win; \n -> otros
	  expresiones.agregarExpresion(&expr)
	  generarTablaSim(tablaSimbolos, expr)
	  fmt.Print("Desea ingresar otra expresion? (S/N): ")
	  fmt.Scanln(&continuar)
  }
  for tablaSimbolos.cantidad > 0 {
	  token = tablaSimbolos.removerDePila()
	  fmt.Println(token.tipo, token.simbolo)
	  switch token.tipo {
		case "VAL":
			if !rVal.MatchString(token.simbolo) {
				fmt.Println(token.tipo, token.simbolo, "no valido")
			}
		case "OP":
			if !rOps.MatchString(token.simbolo) {
				fmt.Println(token.tipo, token.simbolo, "no valido")
			}
		case "ID":
			if !rIDs.MatchString(token.simbolo) {
				fmt.Println(token.tipo, token.simbolo, "no valido")
			}
	  }
	  pilaAux.agregarAPila(token)
  }
  for pilaAux.cantidad > 0 {
	  tablaSimbolos.agregarAPila(pilaAux.removerDePila())
  }
  var expFinal string = construirExpresion(expresiones)
  t1 := construirArbol(expFinal)
  if t1!=nil {
	  //RecorrerPreorden(t1)
	  //fmt.Println()
	  //RecorrerPostorden(t1)
	  //fmt.Println()
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
