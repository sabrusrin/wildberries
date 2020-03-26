package theory

// BookList ...
type BookList interface {
	AppendBookList(string) string
}

type bookList struct {
	book         []string
	theoryHeader string
}

// AppendBookList appends the book list and returns the plan for theory
func (t *bookList) AppendBookList(s string) string {
	var iter int
	var bookList, buffer string
	if len(s) != 0 {
		t.book = append(t.book, s)
	}
	for iter, buffer = range t.book {
		if iter == 0 {
			bookList += " " + buffer
		} else {
			bookList += ", " + buffer
		}
	}
	return t.theoryHeader + bookList + "\n"
}

// NewBookList ...
func NewBookList(s string) BookList {
	return &bookList{
		theoryHeader: s,
	}
}
