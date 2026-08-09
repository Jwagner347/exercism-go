package treebuilding

import (
	"errors"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	ParentID int
	Children []*Node
	// feel free to add fields as you see fit
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	allNodes := make([]*Node, len(records))

	for _, r := range records {
		if r.ID >= len(records) || r.Parent >= len(records) {
			return nil, errors.New("Invalid records")
		}
		isRoot := r.ID == 0 && r.Parent == 0
		isOwnParent := r.ID == r.Parent

		if !isRoot && (isOwnParent || r.ID <= r.Parent) {
			return nil, errors.New("Cannot have ParentID greater than ID")
		}

		if allNodes[r.ID] != nil {
			return nil, errors.New("already indexed on this record")
		}
		allNodes[r.ID] = &Node{ID: r.ID, ParentID: r.Parent}
	}

	for _, n := range allNodes {
		if n.ParentID != n.ID {
			allNodes[n.ParentID].Children = append(allNodes[n.ParentID].Children, n)
		}

	}

	return allNodes[0], nil

}
