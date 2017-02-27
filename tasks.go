package main

func (t *Task) isTaskEnded() bool {
	if t.EndTime.Sub(t.StartTime) > 0 {
		return true
	} else {
		return false
	}
}
