package stackoverflow

import (
	"testing"
	"time"
)

func TestUserContentMethodsUpdateReputationWithoutDeadlock(t *testing.T) {
	operations := []struct {
		name  string
		delta int
		run   func(*User)
	}{
		{"question", QuestionReputation, func(user *User) { user.AddQuestion(NewQuestion(user, "title", "body", nil)) }},
		{"answer", AnswerReputation, func(user *User) { user.AddAnswer(NewAnswer(user, nil, "body")) }},
		{"comment", CommentReputation, func(user *User) { user.AddComment(NewComment(user, "body")) }},
	}

	for _, operation := range operations {
		t.Run(operation.name, func(t *testing.T) {
			user := NewUser(operation.name, operation.name, "")
			done := make(chan struct{})
			go func() {
				operation.run(user)
				close(done)
			}()
			select {
			case <-done:
			case <-time.After(time.Second):
				t.Fatal("user content operation deadlocked")
			}
			if got := user.GetReputation(); got != operation.delta {
				t.Fatalf("reputation = %d, want %d", got, operation.delta)
			}
		})
	}
}
