package acmp

import (
	"testing"
)

func TestDifficulties(t *testing.T) {
	expectedMap := map[string]float64{
		"https://acmp.ru/index.asp?main=task&id_task=1":             2,
		"https://acmp.ru/index.asp?main=task&id_task=2":             19,
		"https://acmp.ru/index.asp?main=task&id_task=3":             8,
		"https://acmp.ru/index.asp?main=task&id_task=4":             4,
		"https://acmp.ru/index.asp?main=task&id_task=5":             15,
		"https://acmp.ru/index.asp?main=task&id_task=6":             23,
		"https://acmp.ru/index.asp?main=task&id_task=7":             30,
		"https://acmp.ru/index.asp?main=task&id_task=8":             5,
		"https://acmp.ru/index.asp?main=task&id_task=9":             27,
		"https://acmp.ru/index.asp?main=task&id_task=10":            17,
		"https://acmp.ru/index.asp?main=task&id_task=2834928374982": -1,
		"https://golang.org/download":                               -1,
	}
	for url, expected := range expectedMap {
		got := Difficulty(url)
		if got != expected {
			t.Errorf("for %s: expected %.2f, but got %.2f", url, expected, got)
		}
	}

}
