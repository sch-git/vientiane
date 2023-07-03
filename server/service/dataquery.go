package service

type dataQuery struct {
}

type DataQueryService interface {
}

func NewDataQueryService() DataQueryService {
	return &dataQuery{}
}
