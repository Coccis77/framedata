/*
Package data analyses messages, parses and reads characters move data
 */
package data

import (
	"strings"
	"fmt"
	"io/ioutil"
	"github.com/fmicaelli/framedata/util"
)

/*
GetMove will be displayed by the bot
 */
type Move interface {
	GetMove(characterData []byte, moveName string) (s string, err error)
}

func FindData(message string, move Move) (s string, err error) {
	// Get the character name
	// Message must have at least Character + Move name
	explode := strings.Fields(message)
	if len(explode) < 2 {
		return
	}
	character := explode[0]
	characterFile := character + ".json"

	// Get all the characters in the folder "json"
	fileInfo, err := ioutil.ReadDir("json")
	if err != nil {
		err = fmt.Errorf("Error while reading folder \"json\"")
	}

	var charactersFile []string
	for _, file := range fileInfo {
		charactersFile = append(charactersFile, file.Name())
	}

	// Check if the character exists
	if !util.ContainsIgnoreCase(charactersFile[:], characterFile) {
		return
	}

	// Get the move of the character
	characterData, err := ioutil.ReadFile("json/" + characterFile)

	// Call the GetMove from the implementation
	s, err = move.GetMove(characterData, explode[1])

	return
}
