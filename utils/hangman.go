package hangmanweb

import (
	"bufio"
	"log"
	"math/rand"
	"os"
)

type HangManData struct {
	Word             string
	ToFind           string
	Attempts         int
	LettersSuggested []string
	Error            string
	HangmanPositions [10]string
}

func Drawhangman(filname string) [10]string { /// fonction pour dessiner le hangman ///
	var hangmanPositions [10]string
	count := 0
	index := 0
	f, err := os.Open("hangman.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if count%8 == 0 && count != 0 {
			index += 1
		}
		line := "\n" + scanner.Text()
		hangmanPositions[index] += line
		count += 1
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return hangmanPositions
}

func ReadFileName(name string) string { /// fonction qui prend un mot dans le fichier words.txt ///
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	file.Close()
	return txtlines[rand.Intn(len(txtlines)-1)]
}

func 