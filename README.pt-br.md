[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-channels/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-channels/blob/main/README.md)

## Canais
Os canais são mais utilizados em go para sincronizar as ```goroutines```, mas sincronizar elas de um pouco melhor, sem se preocupar com contador ou algo do tipo.  
```go
package main

import (
	"time"
  "fmt"
)

func main() {
	channel := make(chan string)
	go writePhrase("Hello World", channel)
	message:= <-channel // significa que estou recebendo um valor, esperando que o canal receba um valor
  fmt.Println(message) // printa apenas uma vez
}

func writePhrase(text string, c chan string) {
	for k := 0; k < 5; k++ {
		c <- text // channel <- value significa que esta mandando um valor pra dentro do canal
		time.Sleep(time.Second)
	}
}
```
Aqui ele escreveu apenas uma vez na tela, porquê uma vez só.  
Vamos falar da natureza dos ```chan``` e como são usados pra sincronizar, ele possui duas operações básicas e importantes, que é *enviar um dado* e *receber um dado*, o interessante dessa operação é que elas bloqueiam a execução do programa.  
Por exemplo, no código acima estamos criando um ```chan``` e passando ele pra uma função com a clausula ```go```, portanto nosso programa não vai esperar essa função executar para seguir, então ele chamou a função ```writePhrase``` e continuou, ai chegou na linha ```message := <- chanell```, e é exatamente nesse ponto que vai ocorrer a sincronização com a ```goroutine``` do ```writePhrase```, quando dizemos ```message := <- chanell``` estamos falando obrigatoriamente para ele esperar nosso ```chan``` receber um valor, e quadno ele recebe um valor? no nosso caso, é recebido na função ```writePhrase```, portanto na somente após a atribuição de do ```chan``` na nossa ```message``` ela espera, recebe o valor da ```chan``` printa e finaliza a aplicação, pois não estamos pedindo para nosso ```chn``` esperar.<br/><br/>

*Deadlock*   
Quando você não tem mais nenhum lugar que está enviando dados no seu canal, só que seu canal ainida está esperando receber um dado, isso gera um erro com seu programa esperando um valor que nunca vai chegar, em nossa função ```writePhrase``` nós executamos apenas 5 vezes, e irá executar 5 vezes, mas ao passar de 5x ele vai sair do lop e não terá mais nada pra executar, mas como estamos num loop infinito, vamos voltar pra função, só que não vai retornar nenhum valor, ai ocorre o deadlock, quando seu programa espera por algo que nunca vai acontecer.  
*O Deadlock não é pego na compilação do código, só em execução*  

### Fechando canal
Para evitar o ```deadlock``` temos de fechar o nosso ```chan```:  
```go
func main() {
	channel := make(chan string) // create chan (channel)
	go writePhrase("Hello World", channel) // go routine
	fmt.Println("Depois da função escrever ser executada")

  for { //loop infinito
    message, open := <-channel // criando message with chan value and verification is open
    if !open { // caso o nosso chan não esteja aberto
      fmt.Println("chan fechado estado é", open)
      break // finaliza a aplicação
    }
    fmt.Println(message, open) // output Hello world
  }
}

func writePhrase(text string, c chan string) {
	for k := 0; k < 5; k++ {
		c <- text // atribuindo valor ao nosso chan
		time.Sleep(time.Second) // delay de um segundo
	}
	// closing the chan
	close(c) // Fecha o canal
}
```
Aqui estamos adicionado uma nova  variável, junto a ```message```, a variável ```open```, ela recebe o valor booleano se o nosso ```chan``` está aberto ou fechado, logo após fazemos a verificação, caso ele não esteja aberto ```!=true``` nós paramos a aplicação para evitar o deadlock, para isso funcionar foi preciso adicionar a função ```close(chan)``` após nosso loop terminar para ele fechar o canal, e podermos verificar a mudança de estado.  
