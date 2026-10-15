package splitwise

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestSettleBalanceReportsDebtorPayingPayer(t *testing.T) {
	payer := NewUser("payer", "Payer", "")
	debtor := NewUser("debtor", "Debtor", "")
	payer.Balances["payer:debtor"] = 25
	debtor.Balances["debtor:payer"] = -25
	service := &SplitwiseService{
		users: map[string]*User{"payer": payer, "debtor": debtor},
	}

	read, write, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	originalStdout := os.Stdout
	os.Stdout = write
	service.SettleBalance("payer", "debtor")
	write.Close()
	os.Stdout = originalStdout
	output, err := io.ReadAll(read)
	read.Close()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(output), "Debtor pays Payer") {
		t.Fatalf("settlement direction = %q", output)
	}
}

func TestConcurrentAddExpenseKeepsBalancesConsistent(t *testing.T) {
	payer := NewUser("payer", "Payer", "")
	debtor := NewUser("debtor", "Debtor", "")
	group := NewGroup("group", "Group")
	service := &SplitwiseService{
		users:  map[string]*User{"payer": payer, "debtor": debtor},
		groups: map[string]*Group{"group": group},
	}

	const expenseCount = 20
	var wg sync.WaitGroup
	for i := 0; i < expenseCount; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			expense := NewExpense(string(rune('a'+i)), 10, "shared", payer)
			expense.AddSplit(NewEqualSplit(payer))
			expense.AddSplit(NewEqualSplit(debtor))
			service.AddExpense(group.ID, expense)
		}(i)
	}
	wg.Wait()
	if got := payer.Balances["payer:debtor"]; got != expenseCount*5 {
		t.Fatalf("payer credit = %v, want %d", got, expenseCount*5)
	}
	if got := debtor.Balances["debtor:payer"]; got != -expenseCount*5 {
		t.Fatalf("debtor balance = %v, want %d", got, -expenseCount*5)
	}
	if len(group.Expenses) != expenseCount {
		t.Fatalf("group expense count = %d, want %d", len(group.Expenses), expenseCount)
	}
}
