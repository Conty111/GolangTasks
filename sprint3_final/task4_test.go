package sprint3_final

import (
	"fmt"
	"log"
	"os"
	"slices"
	"testing"
)

func TestNumbersGen(t *testing.T) {
	filename := "testfile"
	_, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			log.Println(err)
		}
	}(filename)

	t.Run("default positive", func(t *testing.T) {
		file, err := os.OpenFile(filename, os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal("error open file")
		}
		for i := 0; i < 100; i++ {
			_, err = file.WriteString(fmt.Sprintf("%d\n", i))
			if err != nil {
				return
			}
		}
		file.WriteString("aaa")
		file.Close()

		expected := 0
		getted := make([]bool, 100)
		for num := range NumbersGen(filename) {
			if num == expected {
				getted[num] = true
				expected++
			} else {
				t.Errorf("expected: %v, got: %v", expected, num)
			}
		}
		if slices.Contains(getted, false) {
			t.Errorf("not all numbers getted: %v", getted)
		}
	})

	t.Run("invalid filename", func(t *testing.T) {
		expected := 0
		for num := range NumbersGen("aaa") {
			t.Errorf("expected: %v, got: %v", expected, num)
		}
	})
}
