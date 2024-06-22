package business

type startTaskStorage interface {
	StartTask(proid, taskid, assigneeid string) error
}

type startTaskBiz struct {
	store startTaskStorage
}

func NewStartTaskBiz(store startTaskStorage) *startTaskBiz {
	return &startTaskBiz{store: store}
}

func (biz *startTaskBiz) StartTaskBiz(proid, taskid, assigneeid string) error {
	if err_query := biz.store.StartTask(proid, taskid, assigneeid); err_query != nil {
		return err_query
	}
	return nil
}
