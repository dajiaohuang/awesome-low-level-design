package pubsubsystem
import ("testing"; "time")
type oneShot struct { topic *Topic; called chan struct{} }
func (s *oneShot) OnMessage(m *Message) { s.topic.RemoveSubscriber(s); close(s.called) }
func TestCallbackCanUnsubscribe(t *testing.T) {
 topic:=NewTopic("updates"); s:=&oneShot{topic,make(chan struct{})}; topic.AddSubscriber(s)
 done:=make(chan struct{}); go func(){topic.Publish(&Message{});close(done)}()
 select {case <-done: case <-time.After(time.Second):t.Fatal("callback unsubscribe deadlocked")}
 select {case <-s.called:default:t.Fatal("callback not invoked")}
 topic.Publish(&Message{})
}
