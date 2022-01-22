package repository

import (
	"test/driver"
	"test/model"
)

func NewSQLSampleRepo(Conn *driver.DB) SampleRepo {
	return &mysqlSampleRepo{
		Conn: Conn,
	}
}

type mysqlSampleRepo struct {
	Conn *driver.DB
}

func (m *mysqlSampleRepo) Fetch() ([]*model.SampleModel, error) {
	rows, err := m.Conn.SQL.Query("SELECT * FROM sample")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	results := make([]*model.SampleModel, 0)

	for rows.Next() {
		data := new(model.SampleModel)

		err := rows.Scan(
			&data.ID,
			&data.Data,
		)

		if err != nil {
			return nil, err
		}

		results = append(results, data)
	}

	return results, nil
}

func (m *mysqlSampleRepo) Create(sample *model.SampleModel) (uint64, error) {
	return 1, nil
}
