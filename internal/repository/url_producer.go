package repository

import (
	"database/sql"

	"github.com/tiago123456789/pub-sub-wook-api/internal/model"
)

type IUrlProducerRepository interface {
	Create(model.UrlProducer) error
}

type UrlProducerRepository struct {
	db *sql.DB
}

func NewUrlProducerRepository(db *sql.DB) *UrlProducerRepository {
	return &UrlProducerRepository{
		db: db,
	}
}

func (repository *UrlProducerRepository) Create(data model.UrlProducer) (model.UrlProducer, error) {

	err := repository.db.QueryRow(
		`
		INSERT INTO urls_producers(enabled, key) 
		VALUES ($1, $2) RETURNING id
	`,
		data.Enabled, data.Key,
	).Scan(&data.Id)

	if err != nil {
		return model.UrlProducer{}, err
	}

	return data, nil
}

func (repository *UrlProducerRepository) GetById(id int) (model.UrlProducer, error) {
	var urlProducer model.UrlProducer
	err := repository.db.QueryRow(
		`
		SELECT id, key FROM urls_producers where id = $1 and enabled = true LIMIT 1
	`,
		id,
	).Scan(&urlProducer.Id, &urlProducer.Key)

	if err != nil {
		return model.UrlProducer{}, err
	}

	return urlProducer, nil
}
