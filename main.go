package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	fmt.Println("Вы родились в бедной семье фермеров в небольшой деревушке на окраине королевства.")
	fmt.Println("Но сбор урожая и уход за скотом никогда вас не радовал. Вас манил запах золота.")
	fmt.Println("Собрав все свое имущество вы отправились в путешествие в поисках богатств и лучшей жизни.")
	fmt.Println("Ваш путь был непрост, и заняв два месяца он закончился когда вы увидели ворота большого города.")
	fmt.Println("На рыночной площади вы и решили разбирть свою торговую лавку.")

	const SecondsToNextClient = 1
	ticker := time.NewTicker(SecondsToNextClient * time.Second)
	for range ticker.C {
		buyerName := generateBuyerName()
		fmt.Printf("К вашей лавке подошел %s.\n\n", buyerName)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
	}
}

func generateBuyerName() string {
	mobNames := []string{
		"шахтер",
		"искатель приключений",
		"горожанин",
		"алхимик",
	}

	n := rand.Int() % len(mobNames)

	return mobNames[n]
}
