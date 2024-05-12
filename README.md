[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-channels/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-channels/blob/main/README.md)

## Channels
Channels are most used in go to synchronize ```goroutines```, but synchronize them a little better, without worrying about counters or anything like that.
```go
package main

import (
	"time"
  "fmt"
)

func main() {
	channel := make(chan string)
	go writePhrase("Hello World", channel)
	message:= <-channel // means I'm receiving a value, waiting for the channel to receive a value
  fmt.Println(message) // print only once// print only once
}

func writePhrase(text string, c chan string) {
	for k := 0; k < 5; k++ {
		c <- text // channel <- value means that a value is being sent into the channel
		time.Sleep(time.Second)
	}
}
```
Here he wrote only once on the screen, because only once.  
Let's talk about the nature of ```chan``` and how they are used to synchronize, it has two basic and important operations, which is *send data* and *receive data*, the interesting thing about this operation is that they block the program execution.  
For example, in the code above we are creating a ```chan``` and passing it to a function with the ```go``` clause, so our program will not wait for this function to execute before proceeding, so it called the function ```writePhrase``` and continued, then arrived at the ```message := <- chanell``` line, and it is exactly at this point that synchronization with ```goroutine``` of ``` will occur writePhrase```, when we say ```message := <- chanell``` are we necessarily saying that it waits for our ```chan``` to receive a value, and when it receives a value? in our case, it is received in the ```writePhrase``` function, so only after the assignment of ```chan``` in our ```message``` does it wait, it receives the value of ``` chan``` prints and finishes the application, as we are not asking our ```chn``` to wait.<br/><br/>

*Deadlock*  
When you no longer have anywhere that is sending data on your channel, but your channel is still waiting to receive data, this generates an error with your program expecting a value that will never arrive, in our ```writePhrase`` function. ` we only execute it 5 times, and it will execute 5 times, but when it passes 5x it will exit the lop and there will be nothing left to execute, but as we are in an infinite loop, we will return to the function, but it will not return any value , then deadlock occurs, when your program waits for something that will never happen.  
*Deadlock is not caught when compiling the code, only when executing it*  

### Closing channel
To avoid the ```deadlock``` we have to close our ```chan```:  
```go
func main() {
	channel := make(chan string) // create chan (channel)
	go writePhrase("Hello World", channel) // go routine
	fmt.Println("After the write function is executed")

  for { //infinity loop
    message, open := <-channel // creating message with chan value and verification is open
    if !open { // if our chan is not open
      fmt.Println("chan closed state is", open)
      break // finish application
    }
    fmt.Println(message, open) // output Hello world
  }
}

func writePhrase(text string, c chan string) {
	for k := 0; k < 5; k++ {
		c <- text // assigning value to our chan
		time.Sleep(time.Second) // one second delay
	}
	// closing the chan
	close(c)
}
```
Here we have added a new variable, next to ```message```, the variable ```open```, it receives the Boolean value if our ```chan``` is open or closed, then we do the check, if it is not open ```!=true``` we stop the application to avoid the deadlock, for this to work we had to add the function ```close(chan)``` after our loop ends for it close the channel, and we can check the state change.