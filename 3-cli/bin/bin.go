package bin
// Bin package defines the Bin struct and BinList type

type Bin struct {
	ID        string
	Private   bool
	CreatedAt string // В реальном коде лучше использовать time.Time
	Name      string
}
type BinList []Bin

func NewBin(id string, private bool, createdAt string, name string) Bin {
	return Bin{
		ID:        id,
		Private:   private,
		CreatedAt: createdAt,
		Name:      name,
	}
}

func NewBinList() BinList {
	return BinList{}
}