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

func (ts *TagService) GetMany() ([]entities.TagEntity, error) {
	getManyTagQuery := `
		select id, name, link_id, owner_id, created_at, updated_at
			from public.tags;
	`

	var tags []entities.TagEntity
	if err := ts.pg.Select(
		&tags,
		getManyTagQuery,
	); err != nil {
		return []entities.TagEntity{}, err
	}

	return tags, nil
}

func (ts *TagService) GetOne(tagId uuid.UUID) (entities.TagEntity, error) {
	getOneTagQuery := `
		select id, name, link_id, owner_id, created_at, updated_at
			from public.tags
			where id = $1;
	`

	var tag entities.TagEntity
	if err := ts.pg.Select(
		&tag,
		getOneTagQuery,
	); err != nil {
		return entities.TagEntity{}, err
	}

	return tag, nil
}

func (ts *TagService) UpdateOne(tagId uuid.UUID, payload dtos.UpdateTagDto) (entities.TagEntity, error) {
	updateOneTagQuery := `
		update public.tags
			set 
				name = coalesce($2, name),
				link_id = coalesce($3, link_id),
				owner_id = coalesce($4, owner_id)
			where id = $1;
	`

	var updatedTag entities.TagEntity
	if err := ts.pg.Select(
		&updatedTag,
		updateOneTagQuery,
		tagId,
		payload.Name,
		payload.LinkId,
		payload.OwnerId.String(),
	); err != nil {
		return entities.TagEntity{}, err
	}

	return updatedTag, nil
}

func (ts *TagService) DeleteOne(tagId uuid.UUID) (entities.TagEntity, error)
