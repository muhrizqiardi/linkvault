package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"server/dtos"
	"server/entities"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type FolderService struct {
	ctx context.Context
	l   *log.Logger
	pg  *sqlx.DB
}

func NewFolderService(ctx context.Context, l *log.Logger, pg *sqlx.DB) *FolderService {
	return &FolderService{
		ctx,
		l,
		pg,
	}
}

// Create new folder
func (fs *FolderService) Create(ownerId uuid.UUID, payload dtos.CreateFolderDto) (entities.FolderEntity, error) {
	createNewFolderQuery := `
		insert into public.folders (name, owner_id)
			values ($1, $2)
			returning 
				id, name, owner_id, created_at, updated_at;`

	var newFolder entities.FolderEntity
	if err := fs.pg.Get(
		&newFolder,
		createNewFolderQuery,
		payload.Name,
		ownerId.String(),
	); err != nil {
		return entities.FolderEntity{}, err
	}

	return newFolder, nil
}

// Get many folders
func (fs *FolderService) GetManyBelongsToUser(ownerId uuid.UUID, orderBy string, limit int, page int) ([]entities.FolderEntity, error) {
	getManyFoldersQuery := `
		select id, name, owner_id, created_at, updated_at
			from public.folders
			where owner_id = $1
			order by %s
			limit $2 offset $3;
	`
	switch orderBy {
	case "name_ASC":
		getManyFoldersQuery = fmt.Sprintf(getManyFoldersQuery, "name asc")
		break
	case "name_DESC":
		getManyFoldersQuery = fmt.Sprintf(getManyFoldersQuery, "name desc")
		break
	case "createdAt_ASC":
		getManyFoldersQuery = fmt.Sprintf(getManyFoldersQuery, "created_at asc")
		break
	case "createdAt_DESC":
		getManyFoldersQuery = fmt.Sprintf(getManyFoldersQuery, "created_at desc")
		break
	case "updatedAt_ASC":
		getManyFoldersQuery = fmt.Sprintf(getManyFoldersQuery, "updated_at asc")
		break
	case "updatedAt_DESC":
		getManyFoldersQuery = fmt.Sprintf(getManyFoldersQuery, "updated_at desc")
		break
	case "":
		getManyFoldersQuery = fmt.Sprintf(getManyFoldersQuery, "updated_at desc")
		break
	default:
		return []entities.FolderEntity{}, errors.New("Invalid `orderBy` parameter")
	}

	var folders []entities.FolderEntity
	if err := fs.pg.Select(
		&folders,
		getManyFoldersQuery,
		ownerId.String(),
		limit,
		(page-1)*limit,
	); err != nil {
		return []entities.FolderEntity{}, err
	}

	return folders, nil
}

// Get one folder
func (fs *FolderService) GetOne(folderId uuid.UUID) (entities.FolderEntity, error) {
	getOneFolderQuery := `
		select id, name, owner_id, created_at, updated_at
			from public.folders
			where id = $1;
	`

	var folder entities.FolderEntity
	if err := fs.pg.Get(
		&folder,
		getOneFolderQuery,
		folderId.String(),
	); err != nil {
		return entities.FolderEntity{}, err
	}

	return folder, nil
}

func (fs *FolderService) UpdateOne(folderId uuid.UUID, payload dtos.UpdateFolderDto) (entities.FolderEntity, error) {
	updateOneFolderQuery := `
		update public.folders
			set
				name = coalesce($1, name),
				owner_id = coalesce($2, owner_id),
				updated_at = current_timestamp
			where 
				id = $1
			returning
				id, name, owner_id, created_at, updated_at;
	`

	var updatedFolder entities.FolderEntity
	if err := fs.pg.Select(
		&updatedFolder,
		updateOneFolderQuery,
		payload.Name,
		payload.OwnerId.String(),
	); err != nil {
		return entities.FolderEntity{}, err
	}

	return updatedFolder, nil
}

func (fs *FolderService) DeleteOne(folderId uuid.UUID) (entities.FolderEntity, error) {
	deleteOneFolderQuery := `
		delete from public.folders
			where id = $1;
	`

	var deletedFolder entities.FolderEntity
	if err := fs.pg.Select(
		&deletedFolder,
		deleteOneFolderQuery,
		folderId.String(),
	); err != nil {
		return entities.FolderEntity{}, err
	}

	return deletedFolder, nil
}
