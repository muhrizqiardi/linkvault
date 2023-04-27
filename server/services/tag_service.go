package services

import (
	"context"
	"log"
	"server/dtos"
	"server/entities"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type TagService struct {
	ctx context.Context
	l   *log.Logger
	pg  *sqlx.DB
}

func NewTagService(ctx context.Context, l *log.Logger, pg *sqlx.DB) *TagService {
	return &TagService{
		ctx,
		l,
		pg,
	}
}

func (ts *TagService) Create(payload dtos.CreateTagDto) (entities.TagEntity, error) {
	createTagQuery := `
		insert into public.tags (name, link_id, owner_id)
			values ($1, $2, $3)
			returning  id, name, link_id, owner_id, created_at, updated_at;
	`

	var newTag entities.TagEntity
	if err := ts.pg.Select(
		&newTag,
		createTagQuery,
		payload.Name,
		payload.LinkId,
		payload.OwnerId,
	); err != nil {
		return entities.TagEntity{}, err
	}

	return newTag, nil
}

func (ts *TagService) GetMany() ([]entities.TagEntity, error)

func (ts *TagService) GetOne(tagId uuid.UUID) (entities.TagEntity, error)

func (ts *TagService) UpdateOne(tagId uuid.UUID, payload dtos.UpdateTagDto) (entities.TagEntity, error)

func (ts *TagService) DeleteOne(tagId uuid.UUID) (entities.TagEntity, error)
