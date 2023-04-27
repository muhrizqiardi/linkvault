package services

import (
	"context"
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
func (fs *FolderService) Create(payload dtos.CreateFolderDto) (entities.FolderEntity, error) {
	createNewFolderQuery := `
		insert into public.folders (name, owner_id)
			values ($1, $2)
			returning 
				id, name, owner_id, created_at, updated_at;`

	var newFolder entities.FolderEntity
	if err := fs.pg.Select(
		&newFolder,
		createNewFolderQuery,
		payload.Name,
		payload.OwnerId,
	); err != nil {
		return entities.FolderEntity{}, err
	}

	return newFolder, nil
}

// Get many folders
func (fs *FolderService) GetMany() ([]entities.FolderEntity, error) {
	getManyFoldersQuery := `
		select id, name, owner_id, created_at, updated_at
			from public.folders;
	`

	var folders []entities.FolderEntity
	if err := fs.pg.Select(
		&folders,
		getManyFoldersQuery,
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
	if err := fs.pg.Select(
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
