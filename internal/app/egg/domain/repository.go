package domain

type EggRepository interface {
	Create(egg Egg) Egg
}
