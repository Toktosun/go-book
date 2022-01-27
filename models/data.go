package models

var DB []Book

type Book struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	PublishedYear int    `json:"published_year"`
	Author        Author `json:"author"`
}

type Author struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func init() {
	book1 := Book{
		ID:            1,
		Name:          "Собачье сердце",
		Description:   "Потрясающая история :0 ",
		PublishedYear: 1925,
		Author: Author{
			ID:        1,
			FirstName: "Михаил",
			LastName:  "Булгаков",
		},
	}
	DB = append(DB, book1)
}

func FindBookById(id int) (Book, bool) {
	var (
		book  Book
		found bool
	)
	for _, b := range DB {
		if b.ID == id {
			book = b
			found = true
			break
		}
	}
	return book, found
}
