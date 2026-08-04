package treebuilding

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
		allNodes[r.ID] = &Node{ID: r.ID, ParentID: r.Parent}
	}

	for _, n := range allNodes {
		if n.ParentID != n.ID {
			allNodes[n.ParentID].Children = append(allNodes[n.ParentID].Children, n)
		}

	}

	return allNodes[0], nil

}
