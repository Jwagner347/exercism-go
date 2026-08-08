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
	var err error
	if len(records) == 0 {
		return nil, nil
	}

	allNodes := make([]*Node, len(records))

	for _, r := range records {
		if r.ID >= len(records) || r.Parent >= len(records) {
			err = errors.New("Invalid records")
			return nil, err
		}

		if r.Parent == r.ID && r.ID != 0 {
			err = errors.New("Only root node can be its own parent")
			return nil, err
		}

		if r.Parent > r.ID {
			err = errors.New("Cannot have ParentID greater than ID")
			return nil, err
		}

		if allNodes[r.ID] != nil {
			err = errors.New("already indexed on this record")
			return nil, err
		}
		allNodes[r.ID] = &Node{ID: r.ID, ParentID: r.Parent}
	}

	for _, n := range allNodes {
		if n.ParentID != n.ID {
			allNodes[n.ParentID].Children = append(allNodes[n.ParentID].Children, n)
		}

	}

	return allNodes[0], err

}
