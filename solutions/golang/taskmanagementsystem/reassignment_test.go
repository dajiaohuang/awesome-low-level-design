package taskmanagementsystem
import ("testing";"time")
func TestCanonicalReassignment(t *testing.T){
 old:=&User{Id:"old"}; next:=&User{Id:"next"}; tm:=GetTaskManager()
 task:=NewTask("task","Task","Details",time.Now(),1,old);tm.CreateTask(task)
 updated:=NewTask("task","Task","Details",time.Now(),1,next);tm.UpdateTask(updated)
 if task.GetAssignedUser()!=next || len(tm.GetTaskHistory(old))!=0 || len(tm.GetTaskHistory(next))!=1 {t.Fatal("canonical task and index disagree")}
 tm.DeleteTask(task.Id); if len(tm.GetTaskHistory(next))!=0 {t.Fatal("deletion used stale assignee")}
}
