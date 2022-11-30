package db

type AnimalUnit struct {
	AnimalID int    `db:"animal_id"`
	Name     string `db:"name"`
	Age      int    `db:"age"`
	Type     string `db:"type"`
}

type OwnerUnit struct {
	OwnerID int `db:"owner_id"`
	Name    string
}