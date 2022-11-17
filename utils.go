package Utils

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

type HangManData struct {
	Word             string
	ToFind           string
	Attempts         int
	LettersSuggested []string
	Error            string
	HangmanPositions [10]string
}

func Drawhangman(filename string) [10]string {
	var hangmanPositions [10]string
	count := 0
	index := 0

	f, err := os.Open(filename)

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

func Randomword(fileName string) string {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
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

func Indexreveal(data *HangManData, letter string) []int {
	var results []int

	index := len(data.Word)
	tmp := data.Word
	for {
		match := strings.LastIndex(tmp[0:index], letter)
		if match == -1 {
			break
		} else {
			index = match
			results = append(results, match)
		}
	}

	return results
}

func Revealletters(data *HangManData, indexes []int) {
	tmp := []rune(data.ToFind)
	for _, c := range indexes {
		tmp[c] = rune(data.Word[c])
	}
	data.ToFind = string(tmp)
}

func RevealRandomLetter(data *HangManData) {
	nbr := len(data.Word)/2 - 1
	var indexes []int = make([]int, nbr)
	for i := 0; i < nbr; i++ {
		indexes[i] = rand.Intn(len(data.Word) - 1)
	}
	Revealletters(data, indexes)
}

func Letterispresent(data *HangManData, l string) bool {
	for _, letter := range data.LettersSuggested {
		if letter == l {
			return true
		}
	}
	return false
}

func HangMan(data *HangManData, l string) string {

	data.Error = ""
	if len(l) > 1 {
		if l == data.Word {
			data.ToFind = data.Word
		} else {
			data.Attempts -= 2
			data.Error = fmt.Sprintf("The word is not %s, %d attempts remaining\n", l, data.Attempts)
		}
	} else if !Letterispresent(data, l) {
		data.LettersSuggested = append(data.LettersSuggested, l)
		indexes := Indexreveal(data, l)
		if len(indexes) > 0 {
			Revealletters(data, indexes)
		} else {
			data.Attempts -= 1
			data.Error = fmt.Sprintf("Not present in the word, %d attempts remaining\n", data.Attempts)
		}
	} else {
		data.Error = fmt.Sprintf("Letter '%s' already suggested\n", l)
	}
	return data.Error
}
