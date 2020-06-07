package Game

type Item struct {
	Name        string
	Description string
	iType       ItemType
	CanPickUp   bool
	CanEquip    bool
	contains    []Item
}

func (i Item) Inspect() string {
	if i.contains != nil && len(i.contains) > 0 {
		retVal := i.Description + " containing: "
		for _, ci := range i.contains {
			retVal += ci.Name + " "
		}
	}
	return i.Description
}
