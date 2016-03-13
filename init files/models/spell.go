package models

type Spell struct {
	SpellID       int `gorm:"primary_key" sql:"unique_index:spell_and_item" json:"-"`
	SpellName     string
	Profession    string
	ItemID        int            `sql:"unique_index:spell_and_item" json:"-"`
	ItemMaterials []ItemMaterial `gorm:"foreignkey:SpellID"`
}
