package services

import (
	"context"
	"log"
	"server/dtos"
	"server/entities"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type LinkService struct {
	ctx context.Context
	l   *log.Logger
	pg  *sqlx.DB
}

func NewLinkService(ctx context.Context, l *log.Logger, pg *sqlx.DB) *LinkService {
	return &LinkService{
		ctx,
		l,
		pg,
	}
}

func (ls *LinkService) Create(payload dtos.CreateLinkDto) (entities.LinkEntity, error) {
	createNewLinkQuery := `
		insert into public.links (
					url, 
					excerpt, 
					cover_url, 
					owner_id, 
					folder_id 
				)			
			values 
				($1, $2, $3, $4, $5, $6, $7)
			returning 
				url,
				excerpt,
				cover_url,
				owner_id,
				folder_id,
				created_at,
				update_at;
	`
	var newLink entities.LinkEntity
	if err := ls.pg.Select(
		&newLink,
		createNewLinkQuery,
		payload.Url,
		payload.Excerpt,
		payload.CoverUrl,
		payload.OwnerId,
		payload.OwnerId,
		payload.FolderId,
	); err != nil {
		return entities.LinkEntity{}, err
	}

	return newLink, nil
}

func (ls *LinkService) GetOne(linkId uuid.UUID) (entities.LinkEntity, error) {
	getOneLinkQuery := `
		select 
				id, 
				url, 
				excerpt,
				cover_url,
				owner_id,
				folder_id,
				created_at,
				updated_at
			from 
				public.links
			where
				id = $1; 
	`
	var link entities.LinkEntity
	if err := ls.pg.Get(
		&link,
		getOneLinkQuery,
		linkId.String(),
	); err != nil {
		return entities.LinkEntity{}, err
	}
	return link, nil
}

func (ls *LinkService) GetMany() ([]entities.LinkEntity, error) {
	getManyLinkQuery := `
		select 
				id, 
				url, 
				excerpt,
				cover_url,
				owner_id,
				folder_id,
				created_at,
				updated_at
			from 
				public.links;
	`

	var links []entities.LinkEntity
	if err := ls.pg.Select(
		&links,
		getManyLinkQuery,
	); err != nil {
		return []entities.LinkEntity{}, err
	}
	return links, nil
}

func (ls *LinkService) UpdateOne(linkId uuid.UUID, payload dtos.UpdateLinkDto) (entities.LinkEntity, error) {
	updateOneLinkQuery := `
		update public.links
			set 
				id = coalesce($1, id), 
				url = coalesce($2, url), 
				excerpt = coalesce($3, excerpt),
				cover_url = coalesce($4, cover_url),
				owner_id = coalesce($5, owner_id),
				folder_id = coalesce($6, folder_id),
				created_at = coalesce($7, created_at),
				updated_at = current_timestamp
			returning
				id, 			
				url, 			
				excerpt, 
				cover_url, 
				owner_id, 
				folder_id, 
				created_at, 
				updated_at;
	`

	var updatedLink entities.LinkEntity
	if err := ls.pg.Select(
		&updatedLink,
		updateOneLinkQuery,
		linkId.String(),
	); err != nil {
		return entities.LinkEntity{}, err
	}
	return updatedLink, nil
}

func (ls *LinkService) DeleteOne(linkId uuid.UUID) (entities.LinkEntity, error) {
	deleteOneLinkQuery := `
		delete from public.link
			where
				id = $1;
	`

	var deletedLink entities.LinkEntity
	if err := ls.pg.Select(
		&deletedLink,
		deleteOneLinkQuery,
		linkId.String(),
	); err != nil {
		return entities.LinkEntity{}, err
	}
	return deletedLink, nil
}
