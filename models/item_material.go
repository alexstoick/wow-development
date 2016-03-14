package models

type ItemMaterial struct {
	ID         int  `gorm:"primary_key" json:"-"`
	SpellID    int  `sql:"unique_index:spell_and_material" json:"-"`
	MaterialID int  `sql:"unique_index:spell_and_material" json:"-"`
	Material   Item `gorm:"foreignkey:MaterialID"`
	Quantity   int
}
