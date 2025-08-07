package domainModel

type Rarity string

const (
	RarityCommon   Rarity = "common"
	RarityUncommon Rarity = "uncommon"
	RarityRare     Rarity = "rare"
)

type ItemObject struct {
	Id       string
	Name     string
	Price    int
	ReqLevel uint8
	MaxStack uint8
	Rarity   Rarity
	ItemType *ItemTypeObject
}
