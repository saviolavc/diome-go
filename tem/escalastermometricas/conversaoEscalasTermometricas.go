package main // decalração do pacote principal

import "fmt" // importação do fmt de print e scan

const tempK = 273.15 // declaração de constantes de tempk e tempF temperatura fahrenheit e Kelvi.
const tempF = 212

func main() {

	tempC := 0.0
	fmt.Println("Digite a temperatura deseja em Celsius que deseja converter para Kelvin e Fahrenheit")
	fmt.Scanln(&tempC)                 // aqui digitara a temperatura que deseja converter para graus celsius para as demais, kelvi e fahrenheit.
	tempCparaK := tempC + tempK        // calculo de kelvi
	tempCparaF := (tempC * 9 / 5) + 32 // calculo de fahrenheit

	fmt.Printf("Temperatura de Celsius %g em Kelvin =: %g, e em Fahrenheit: %g \n", tempC, tempCparaK, tempCparaF) //resultado em uma linha

}
