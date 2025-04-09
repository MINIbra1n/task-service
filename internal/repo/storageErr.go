package repo

type FailedInsert struct{}

func (b *FailedInsert) Error() string {
	return "failed to insert task"
}

type EmtyStorage struct{}

func (e *EmtyStorage) Error() string {
	return "Empty storage"
}

type NoSuchTaskStorage struct{}

func (n *NoSuchTaskStorage) Error() string {
	return "No such task in storage"
}

type FailedToUpdate struct{}

func (f *FailedToUpdate) Error() string {
	return "Failed to update task"
}

type NothingToDeleteById struct{}

func (n *NothingToDeleteById) Error() string {
	return "Nothing to delete by id"
}
