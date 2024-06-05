package business

type deleteAccountStorage interface {
	DeleteAccount(id string) error
}

type deleteAccountBuisiness struct {
	store deleteAccountStorage
}

func NewDeleteAccountBusiness(store deleteAccountStorage) *deleteAccountBuisiness {
	return &deleteAccountBuisiness{store: store}
}

func (biz *deleteAccountBuisiness) DeleteAccountBiz(id string) error {
	if err_query := biz.store.DeleteAccount(id); err_query != nil {
		return err_query
	}
	return nil
}
