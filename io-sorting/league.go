package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

func (l League) find(name string) *Player {
	var p *Player
	for i, player := range l {
		if player.Name == name {
			p = &l[i]
		}
	}
	return p
}

func NewLeague(rdr io.Reader) (League, error) {
	var league League
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing leagur: %v", err)
	}

	return league, err
}
