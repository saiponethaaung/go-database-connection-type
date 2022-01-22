package repository

import (
	"test/model"
)

type SampleRepo interface {
	Fetch() ([]*model.SampleModel, error)
	Create(user *model.SampleModel) (uint64, error)
}
