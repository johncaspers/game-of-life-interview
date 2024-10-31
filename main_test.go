package main

import (
	"testing"
)

func TestNewWorld(t *testing.T) {
	t.Run("Create a random world", func(t *testing.T) {
		world := NewWorld(3, 3)
		if world == nil {
			t.Errorf("World is nil")
		}
	})
}

func TestNeighbors(t *testing.T) {

}

func TestAlive(t *testing.T) {

}

func TestNextGeneration(t *testing.T) {

}
