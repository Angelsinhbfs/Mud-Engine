package Game

type ItemType int

const (
	Loot ItemType = iota
	Weapon
	Resource
	Helmet
	Gloves
	ChestPiece
	Accessory
	Pants
	Boots
)

func (d ItemType) String() string {
	return [...]string{
		"Loot",
		"Weapon",
		"Resource",
		"Helmet",
		"Gloves",
		"ChestPiece",
		"Accessory",
		"Pants",
		"Boots",
	}[d]
}

func GetItemType(s string) (ItemType, bool) {
	switch s {
	case "Loot":
		return Loot, true
	case "Weapon":
		return Weapon, true
	case "Resource":
		return Resource, true
	case "Helmet":
		return Helmet, true
	case "Gloves":
		return Gloves, true
	case "ChestPiece":
		return ChestPiece, true
	case "Accessory":
		return Accessory, true
	case "Pants":
		return Pants, true
	case "Boots":
		return Boots, true
	}
	return -1, false
}
