package atm

import "testing"

func TestRejectedCashPayoutRestoresDebit(t *testing.T) {
	bank := NewBankingService()
	bank.CreateAccount("customer", 100)
	dispenser := NewCashDispenser(20)
	machine := NewATM(bank, dispenser)
	if err := machine.WithdrawCash("customer", 50); err != ErrInsufficientCashInATM {
		t.Fatalf("expected existing payout failure, got %v", err)
	}
	if balance, _ := machine.CheckBalance("customer"); balance != 100 {
		t.Fatalf("rejected payout left balance %v", balance)
	}
	if err := machine.WithdrawCash("customer", 20); err != nil {
		t.Fatal(err)
	}
	if balance, _ := machine.CheckBalance("customer"); balance != 80 {
		t.Fatalf("successful payout left balance %v", balance)
	}
}
