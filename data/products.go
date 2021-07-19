package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Product defines the structure of an API product
type Product struct {
	ID          int `json:"id"`
	Name        string ` json:"name"`
	Description string `json:"description"`
	Price       float32 `json:"price"`
	SKU         string `json:"sku"`
	CreatedOn   string `json:"_"`
	UpdatedOn   string `json:"_"`
	DeletedOn   string `json:"_"`
}

func (p *Product) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

// Prroducts is a collection of Product
type Products []*Product 

// ToJSON serializes the contents of the collection to JSON
// NewEncoder provides better prformance than json.Unmarshal as it does not has to buffer
// the out put into an in memory slicw of bytes which reduce allocation and overheads of the services 
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return productList
}

func AddProduct (p *Product) {
	p.ID = getNextId()
	productList = append(productList, p)
}

func getNextId() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p
	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")
func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Ikogoki",
		Description: "Black strong coffee",
		Price:       4.49,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Owiedo",
		Description: "Milky brownie coffee",
		Price:       3.99,
		SKU:         "xyz987",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   time.Now().UTC().String(),
	},
}
