package theory

// BookList ...
type BookList interface {
	AppendBookList(string) (string, error)
}

type bookList struct {
	book         []string
	theoryHeader string
}

// AppendBookList appends the book list and returns the plan for theory
func (t *bookList) AppendBookList(s string) (res string, err error) {
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
	res = t.theoryHeader + bookList + "\n"
	err = nil
	return
}

// NewBookList ...
func NewBookList(s string) BookList {
	return &bookList{
		theoryHeader: s,
	}
}
