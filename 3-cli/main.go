package main

import "fmt"

/*

Создать struct Bin и списка BinList с полями
	id - строка
	private - bool
	createdAt - дата и время
	name - строка
Сделать функции создания

*/

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

func (bl *BinList) AddBin(bin Bin) {
	*bl = append(*bl, bin)
}

func main() {
	binList := NewBinList()
	bin1 := NewBin("bin1", true, "2024-06-01T10:00:00Z", "First Bin")
	bin2 := NewBin("bin2", false, "2024-06-02T11:00:00Z", "Second Bin")
	binList.AddBin(bin1)
	binList.AddBin(bin2)
	fmt.Println(binList)
}