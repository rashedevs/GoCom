package repo

type Product struct {
	ID          int     `json:"id"` //add tag
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageUrl"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(pId int) (*Product, error)
	List() ([]*Product, error)
	Update(p Product) (*Product, error)
	Delete(pId int) error
}

type productRepo struct {
	productList []*Product
}

// constructor function
func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	generateInitialProducts(repo)
	return repo
}

func (r productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil
}

func (r productRepo) Get(pId int) (*Product, error) {
	for _, product := range r.productList {
		if product.ID == pId {
			return product, nil
		}
	}
	return nil, nil
}

func (r productRepo) List() ([]*Product, error) {
	return r.productList, nil
}

func (r productRepo) Update(p Product) (*Product, error) {
	for idx, product := range r.productList {
		if product.ID == p.ID {
			r.productList[idx] = &p
		}
	}
	return &p, nil
}

func (r productRepo) Delete(pId int) error {
	var tempList []*Product
	for _, product := range r.productList {
		if product.ID != pId {
			tempList = append(tempList, product)
			// tempList[idx] = product

		}
	}
	r.productList = tempList
	return nil
}

func generateInitialProducts(r *productRepo) {
	prd1 := &Product{
		ID:          1,
		Title:       "Orange",
		Description: "This is orange.",
		Price:       150.00,
		ImgURL:      "https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcRwJf4xPOftZLLgkWjr2eMMumu9XuRdKiGc1eZXFku9OmA4lSymIPBm_vU0bFIM_BjpYOsI_pX7O9c9sRUxfkXxE0N1x_bWSERw7SXHP3A",
	}
	prd2 := &Product{
		ID:          2,
		Title:       "Apple",
		Description: "This is apple.",
		Price:       120.00,
		ImgURL:      "https://encrypted-tbn3.gstatic.com/images?q=tbn:ANd9GcSbWqLu3A4mBXzZNi1POVqPrGV4UAuPw3Bl9RwTDjJqjCQrMiid5VgbDGHbAIA0rcFkAWdpjba5sqFdLdIvRXZ27_H_dPxXDpebJuqHiHQ",
	}
	prd3 := &Product{
		ID:          3,
		Title:       "Banana",
		Description: "This is banana.",
		Price:       40.00,
		ImgURL:      "https://www.allrecipes.com/thmb/lc7nSL9L5zMHXz9t6PMAVm9biNM=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/ar-new-banana-adobe-ar-2x1-917fdde58d194b529b41042ebff1c031.jpg",
	}
	prd4 := &Product{
		ID:          4,
		Title:       "Grapes",
		Description: "This is grapes.",
		Price:       140.00,
		ImgURL:      "https://d2t8nl1y0ie1km.cloudfront.net/images/6706bbb12b01cd80b29a0488_Grapes-Red-N-500gm_1.webp",
	}
	prd5 := &Product{
		ID:          5,
		Title:       "Mango",
		Description: "This is mango.",
		Price:       80.00,
		ImgURL:      "https://belescooverseas.com/wp-content/uploads/2024/04/Mango.jpeg",
	}

	r.productList = append(r.productList, prd1)
	r.productList = append(r.productList, prd2)
	r.productList = append(r.productList, prd3)
	r.productList = append(r.productList, prd4)
	r.productList = append(r.productList, prd5)
}
