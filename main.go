package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	printIntroLine("Вы родились в бедной семье фермеров в небольшой деревушке на окраине королевства.")
	printIntroLine("Но сбор урожая и уход за скотом никогда вас не радовал. Вас манил запах золота.")
	printIntroLine("Собрав все свое имущество вы отправились в путешествие в поисках богатств и лучшей жизни.")
	printIntroLine("Ваш путь был непрост, и заняв два месяца он закончился когда вы увидели ворота большого города.")
	printIntroLine("На рыночной площади вы и решили разбирть свою торговую лавку.")

	for {
		buyerName := generateBuyerName()
		fmt.Printf("\nК вашей лавке подошел %s.\n\n", buyerName)

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Что вы хотите продать? : ")
		input, _ := reader.ReadString('\n')

		fmt.Printf("Похоже, что %s не заинтересован в \"%s\"\n", buyerName, strings.TrimSpace(input))
	}
}

func printIntroLine(message string) {
	fmt.Println(message)
	timer := time.NewTimer(0 * time.Second)
	<-timer.C
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
