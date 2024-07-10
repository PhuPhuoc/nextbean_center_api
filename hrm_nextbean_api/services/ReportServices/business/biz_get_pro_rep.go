package business

type getProjectReportStorage interface {
	GetProjectReportInOJT(oid string) ([]byte, error)
}

type getProjectReportBiz struct {
	store getProjectReportStorage
}

func NewGetProjectReportBiz(store getProjectReportStorage) *getProjectReportBiz {
	return &getProjectReportBiz{store: store}
}

func (biz *getProjectReportBiz) GetProjectReportBiz(oid string) ([]byte, error) {
	data, err_query := biz.store.GetProjectReportInOJT(oid)
	if err_query != nil {
		return nil, err_query
	}
	return data, nil
}
