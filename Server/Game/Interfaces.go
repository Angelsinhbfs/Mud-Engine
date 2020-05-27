package Game

type Inspectable interface {
	Inspect() string
}

type Attackable interface {
	Attack() bool
}

type Interactable interface {
	CheckInteractions() []string
}

type Updateable interface {
	Update()
}
