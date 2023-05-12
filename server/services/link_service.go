package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"server/dtos"
	"server/entities"
	"server/utils"

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

func (ls *LinkService) Create(ownerId uuid.UUID, folderId uuid.UUID, payload dtos.CreateLinkDto) (entities.LinkEntity, error) {
	// TODO: create query that checks whether folder belongs to user
	createNewLinkQuery := `
		insert into public.links (url, title, excerpt, cover_url, owner_id, folder_id)
			values ($1, $2, $3, $4, $5, $6)
			returning url, title, excerpt, cover_url, owner_id, folder_id, created_at, updated_at;
	`
	var newLink entities.LinkEntity
	if err := ls.pg.QueryRowx(
		createNewLinkQuery,
		payload.Url,
		payload.Title,
		payload.Excerpt,
		payload.CoverUrl,
		ownerId.String(),
		folderId.String(),
	).StructScan(&newLink); err != nil {
		return entities.LinkEntity{}, err
	}
	return newLink, nil
}

func (ls *LinkService) GetOne(linkId uuid.UUID) (entities.LinkEntity, error) {
	getOneLinkQuery := `
		select id, title, url, excerpt, cover_url, owner_id, folder_id, created_at, updated_at
			from public.links
			where id = $1; 
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

func (ls *LinkService) GetManyBelongsToUser(ownerId uuid.UUID, title string, excerpt string, orderBy string, limit int, page int) ([]entities.LinkEntity, error) {
	getManyLinkQuery := `
		select id, url, title, excerpt, cover_url, owner_id, folder_id, created_at, updated_at
			from public.links
			where 
				owner_id = $1 and
				($2::text is null or $2 = '' or excerpt ilike $2) and
				($3::text is null or $3 = '' or title ilike $3)
			order by %s 
			limit $4 offset $5;
	`

	if err := utils.ValidateEnumString(orderBy, "title_ASC", "title_DESC", "createdAt_ASC", "createdAt_DESC", "updatedAt_ASC", "updatedAt_DESC"); err != nil {
		return []entities.LinkEntity{}, errors.New("Invalid `orderBy` param")
	}
	switch orderBy {
	case "title_ASC":
		getManyLinkQuery = fmt.Sprintf(getManyLinkQuery, "title asc")
		break
	case "title_DESC":
		getManyLinkQuery = fmt.Sprintf(getManyLinkQuery, "title desc")
		break
	case "createdAt_ASC":
		getManyLinkQuery = fmt.Sprintf(getManyLinkQuery, "created_at asc")
		break
	case "createdAt_DESC":
		getManyLinkQuery = fmt.Sprintf(getManyLinkQuery, "created_at desc")
		break
	case "updatedAt_ASC":
		getManyLinkQuery = fmt.Sprintf(getManyLinkQuery, "updated_at asc")
		break
	case "updatedAt_DESC":
		getManyLinkQuery = fmt.Sprintf(getManyLinkQuery, "updated_at desc")
		break
	}

	var links []entities.LinkEntity
	if err := ls.pg.Select(&links, getManyLinkQuery, ownerId.String(), title, excerpt, limit, (page-1)*limit); err != nil {
		return []entities.LinkEntity{}, err
	}
	return links, nil
}

func (ls *LinkService) GetManyBelongsToUserInFolder(ownerId uuid.UUID, folderId uuid.UUID, title string, excerpt string, orderBy string, limit int, page int) ([]entities.LinkEntity, error) {
	getManyLinkQuery := `
		select id, url, title, excerpt, cover_url, owner_id, folder_id, created_at, updated_at
			from public.links
			where 
				owner_id = $1 and
				folder_id = $2 and
				($3::text is null or $3 = '' or title ilike $3) and
				($4::text is null or $4 = '' or excerpt ilike $4)
			order by %s 
			limit $5 offset $6;
	`

	if err := utils.ValidateEnumString(orderBy, "title_ASC", "title_DESC", "createdAt_ASC", "createdAt_DESC", "updatedAt_ASC", "updatedAt_DESC"); err != nil {
		return []entities.LinkEntity{}, errors.New("Invalid `orderBy` param")
	}
	switch orderBy {
	case "title_ASC":
		getManyLinkQuery = fmt.Sprintf(getManyLinkQuery, "title asc")
		break
	case "title_DESC":
		getManyLinkQuery = fmt.Sprintf(getManyLinkQuery, "title desc")
		break
	case "createdAt_ASC":
		getManyLinkQuery = fmt.Sprintf(getManyLinkQuery, "created_at asc")
		break
	case "createdAt_DESC":
		getManyLinkQuery = fmt.Sprintf(getManyLinkQuery, "created_at desc")
		break
	case "updatedAt_ASC":
		getManyLinkQuery = fmt.Sprintf(getManyLinkQuery, "updated_at asc")
		break
	case "updatedAt_DESC":
		getManyLinkQuery = fmt.Sprintf(getManyLinkQuery, "updated_at desc")
		break
	}

	var links []entities.LinkEntity
	if err := ls.pg.Select(&links, getManyLinkQuery, ownerId.String(), folderId, title, excerpt, limit, (page-1)*limit); err != nil {
		return []entities.LinkEntity{}, err
	}
	return links, nil
}

func (ls *LinkService) UpdateOne(linkId uuid.UUID, payload dtos.UpdateLinkDto) (entities.LinkEntity, error) {
	updateOneLinkQuery := `
		update public.links
			set 
				title = coalesce($2, title), 
				excerpt = coalesce($3, excerpt),
				cover_url = coalesce($4, cover_url),
				updated_at = current_timestamp
			where id = $1
			returning
				id, url, title, excerpt, cover_url, owner_id, folder_id, created_at, updated_at;
	`

	var updatedLink entities.LinkEntity
	if err := ls.pg.Get(
		&updatedLink,
		updateOneLinkQuery,
		linkId.String(),
		payload.Title,
		payload.Excerpt,
		payload.CoverUrl,
	); err != nil {
		return entities.LinkEntity{}, err
	}
	return updatedLink, nil
}

func (ls *LinkService) DeleteOne(linkId uuid.UUID) (entities.LinkEntity, error) {
	deleteOneLinkQuery := `delete from public.links where id = $1;`

	var deletedLink entities.LinkEntity
	if _, err := ls.pg.Exec(deleteOneLinkQuery, linkId.String()); err != nil {
		return entities.LinkEntity{}, err
	}
	return deletedLink, nil
}
