package main

import (
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from reader", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "Vane", "Wins": 10},
			{"Name": "Oscar", "Wins": 33}
			]`)

		store := FileSystemPlayerStore{database}

		got := store.GetLeague()

		want := []Player{
			{"Vane", 10},
			{"Oscar", 33},
		}

		assertLeague(t, got, want)
	})
}
