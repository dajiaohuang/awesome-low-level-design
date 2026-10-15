package librarymanagementsystem

import "testing"

func newTestLibraryManager() *LibraryManager {
	return &LibraryManager{
		catalog:     make(map[string]*Book),
		members:     make(map[string]*Member),
		loanMembers: make(map[string]*Member),
	}
}

func TestReturnBookAfterCatalogRemoval(t *testing.T) {
	manager := newTestLibraryManager()
	book := NewBook("isbn-1", "Book", "Author", 2020)
	member := NewMember("member-1", "Member", "contact")
	manager.AddBook(book)
	manager.RegisterMember(member)
	if err := manager.BorrowBook(member.MemberID, book.ISBN); err != nil {
		t.Fatal(err)
	}

	manager.RemoveBook(book.ISBN)
	if err := manager.ReturnBook(member.MemberID, book.ISBN); err != nil {
		t.Fatalf("return after catalog removal: %v", err)
	}
	if len(member.GetBorrowedBooks()) != 0 || !book.IsAvailable() {
		t.Fatal("return should clear the loan and restore availability")
	}
	if manager.GetBook(book.ISBN) != nil {
		t.Fatal("return should not re-add a removed catalog entry")
	}
}

func TestReturnBookAfterMemberUnregistration(t *testing.T) {
	manager := newTestLibraryManager()
	book := NewBook("isbn-2", "Book", "Author", 2020)
	member := NewMember("member-2", "Member", "contact")
	manager.AddBook(book)
	manager.RegisterMember(member)
	if err := manager.BorrowBook(member.MemberID, book.ISBN); err != nil {
		t.Fatal(err)
	}

	manager.UnregisterMember(member.MemberID)
	if manager.GetMember(member.MemberID) != nil {
		t.Fatal("unregistered member should no longer be registered")
	}
	if err := manager.ReturnBook(member.MemberID, book.ISBN); err != nil {
		t.Fatalf("return after unregistration: %v", err)
	}
	if len(member.GetBorrowedBooks()) != 0 || !book.IsAvailable() {
		t.Fatal("return should clear the retained loan and restore availability")
	}
	if manager.loanMembers[member.MemberID] != nil {
		t.Fatal("retired loan owner should be released after the final return")
	}
}
