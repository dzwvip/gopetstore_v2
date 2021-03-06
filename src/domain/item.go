package domain

type Item struct {
	ItemId     string  `db:"itemid"`
	ProductId  string  `db:"productid"`
	ListPrice  float64 `db:"listprice"`
	UnitCost   float64 `db:"unitcost"`
	SupplierId int     `db:"supplier"`
	Status     string  `db:"status"`
	Attr1      string  `db:"attr1"`
	Attr2      string  `db:"attr2"`
	Attr3      string  `db:"attr3"`
	Attr4      string  `db:"attr4"`
	Attr5      string  `db:"attr5"`
	Quantity   int     `db:"quantity"`
	*Product
}
