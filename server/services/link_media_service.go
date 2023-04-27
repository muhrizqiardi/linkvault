package services

import (
	"context"
	"log"
	"server/dtos"
	"server/entities"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type LinkMediaService struct {
	ctx context.Context
	l   *log.Logger
	pg  *sqlx.DB
}

func NewLinkMediaService(ctx context.Context, l *log.Logger, pg *sqlx.DB) *LinkMediaService {
	return &LinkMediaService{
		ctx,
		l,
		pg,
	}
}

func (lms *LinkMediaService) Create(payload dtos.CreateLinkMediaDto) (entities.LinkMediaEntity, error) {
	createLinkMediaQuery := `
		insert into public.link_medias (link_id, media_url, owner_id)
			value ($1, $2, $3)
			returning (id, link_id, media_url, owner_id, created_at, updated_at);
	`

	var newLinkMedia entities.LinkMediaEntity
	if err := lms.pg.Select(
		&newLinkMedia,
		createLinkMediaQuery,
		payload.LinkId,
		payload.MediaUrl,
		payload.OwnerId,
	); err != nil {
		return entities.LinkMediaEntity{}, err
	}

	return newLinkMedia, nil
}

func (lms *LinkMediaService) GetMany() ([]entities.LinkMediaEntity, error) {
	getManyLinkMediasQuery := `
		select id, link_id, media_url, owner_id, created_at, updated_at
			from public.link_medias;	
	`

	var linkMedias []entities.LinkMediaEntity
	if err := lms.pg.Select(
		&linkMedias,
		getManyLinkMediasQuery,
	); err != nil {
		return []entities.LinkMediaEntity{}, err
	}

	return linkMedias, nil
}

func (lms *LinkMediaService) GetOne(linkMediaId uuid.UUID) (entities.LinkMediaEntity, error) {
	getOneLinkMediaQuery := `
		select id, link_id, media_url, owner_id, created_at, updated_at
			from public.link_medias
			where id = $1;
	`

	var linkMedia entities.LinkMediaEntity
	if err := lms.pg.Select(
		&linkMedia,
		getOneLinkMediaQuery,
		linkMediaId.String(),
	); err != nil {
		return entities.LinkMediaEntity{}, err
	}

	return linkMedia, nil
}

func (lms *LinkMediaService) UpdateOne(linkMediaId uuid.UUID, payload dtos.UpdateLinkMediaDto) (entities.LinkMediaEntity, error)

func (lms *LinkMediaService) DeleteOne(linkMediaId uuid.UUID) (entities.LinkMediaEntity, error)
