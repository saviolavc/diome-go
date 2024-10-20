package main // declarar meu pacote principal

import "fmt" //importr função fmt

const ebulicaoF = 212.0 //declaracao da variavel const do ponto de ebolicao da agua em Faren...

func main() {

	//tempF := ebulicaoF            //variavel do valor da temperatura em grasus F
	tempC := (ebulicaoF - 32) * 5 / 9 //conversão

	fmt.Printf("A temperatura de ebulição em F = %g   e temperatura de ebulição da agua em C = %g \n", ebulicaoF, tempC)

	//fmt.Println("Temperatura de ebulicao da agua em f= ", tempF) //outher method the view
	//fmt.Println("a temperatura da agua em C =", tempC)

}
