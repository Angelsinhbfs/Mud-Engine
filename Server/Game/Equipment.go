package Game

type Equipment struct {
	Weapon    EqSlot
	Helmet    EqSlot
	Gloves    EqSlot
	Chest     EqSlot
	Accessory EqSlot
	Pants     EqSlot
	Boots     EqSlot
}

func (e *Equipment) Init() {
	e.Weapon.IType = Weapon
	e.Helmet.IType = Helmet
	e.Gloves.IType = Gloves
	e.Chest.IType = ChestPiece
	e.Accessory.IType = Accessory
	e.Pants.IType = Pants
	e.Boots.IType = Boots
}

type EqSlot struct {
	IType    ItemType
	Equipped *Item
}

func (e *Equipment) Equip(i *Item, s *EqSlot) bool {
	if s.IType == i.iType && s.Equipped == nil {
		s.Equipped = i
		return true
	}
	return false
}
