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

type GetMove func(characterData []byte, moveName string) (s string, err error)

func FindData(message string, getMove GetMove) (s string, err error) {
	// Get the character name
	// Message must have at least Character + Move name
	explode := strings.SplitN(message, " ", 2)
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
	if !util.ContainsIgnoreCaseAndWithoutSpace(charactersFile[:], characterFile) {
		return
	}

	// Get the move of the character
	characterData, err := ioutil.ReadFile("json/" + strings.ToLower(characterFile))

	// Call the GetMove from the implementation
	s, err = getMove(characterData, explode[1])

	return
}
