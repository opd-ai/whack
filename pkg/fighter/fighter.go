// Package fighter provides moveset trees and combo logic.
package fighter

import "github.com/opd-ai/whack/pkg/engine"

// Fighter represents a generated fighter entity with its moveset.
type Fighter struct {
	Entity engine.Entity
	Name   string
}

// ComboTree represents a fighter's combo tree structure.
type ComboTree struct {
	Root *ComboNode
}

// ComboNode represents a node in the combo tree.
type ComboNode struct {
	MoveName string
	Children []*ComboNode
}

// NewFighter creates a new fighter associated with an entity.
func NewFighter(entity engine.Entity, name string) *Fighter {
	return &Fighter{
		Entity: entity,
		Name:   name,
	}
}
