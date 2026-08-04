package treebuilding

import "sort"

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	tree := Node{ID: 0}

	if len(records) == 1 {
		return &tree, nil
	}

	directParent := 0

	for _, r := range records {
		if r.Parent > 0 {

		}
		if tree.ID != r.ID {
			tree.Children = appendAndSort(&tree, r)
		}
	}
	return &tree, nil

}

func appendAndSort(n *Node, r Record) []*Node {
	n.Children = append(n.Children, &Node{ID: r.ID})
	sort.Slice(n.Children, func(i int, j int) bool {
		return n.Children[i].ID < n.Children[j].ID
	})
	return n.Children
}
