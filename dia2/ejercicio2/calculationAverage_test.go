package main

import (
	"github.com/stretchr/testify/assert"
    "testing"
)

func TestAverageNotes(t *testing.T) {
	t.Run("Average notes not empty", func(t *testing.T) {
			// arrange
			note1 := 1
			note2 := 5

			expectResult := 3

			// act
			result := calculationAverange(note1, note2)

			// assert

			assert.Equal(t, expectResult, result)
	})

	t.Run("Average note empty", func(t *testing.T) {
			// arrange
			expectResult := 0
			// act
			result := calculationAverange(1, 2)

			// assert
			assert.NotEqual(t, expectResult, result)

	})
}