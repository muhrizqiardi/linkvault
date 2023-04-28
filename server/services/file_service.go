package services

import (
	"context"
	"log"
	"server/dtos"
	"server/entities"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type FileService struct {
	ctx context.Context
	l   *log.Logger
	pg  *sqlx.DB
}

func NewFileService(ctx context.Context, l *log.Logger, pg *sqlx.DB) *FileService {
	return &FileService{
		ctx,
		l,
		pg,
	}
}

func (lms *FileService) Create(payload dtos.CreateFileDto) (entities.FileEntity, error) {
	createFileQuery := `
		insert into public.files (link_id, file_url, owner_id)
			value ($1, $2, $3)
			returning id, link_id, file_url, owner_id, created_at, updated_at;
	`

	var newFile entities.FileEntity
	if err := lms.pg.Select(
		&newFile,
		createFileQuery,
		payload.LinkId,
		payload.FileUrl,
		payload.OwnerId,
	); err != nil {
		return entities.FileEntity{}, err
	}

	return newFile, nil
}

func (lms *FileService) GetMany() ([]entities.FileEntity, error) {
	getManyFilesQuery := `
		select id, link_id, file_url, owner_id, created_at, updated_at
			from public.files;	
	`

	var linkMedias []entities.FileEntity
	if err := lms.pg.Select(
		&linkMedias,
		getManyFilesQuery,
	); err != nil {
		return []entities.FileEntity{}, err
	}

	return linkMedias, nil
}

func (lms *FileService) GetOne(linkMediaId uuid.UUID) (entities.FileEntity, error) {
	getOneFileQuery := `
		select id, link_id, file_url, owner_id, created_at, updated_at
			from public.files
			where id = $1;
	`

	var linkMedia entities.FileEntity
	if err := lms.pg.Select(
		&linkMedia,
		getOneFileQuery,
		linkMediaId.String(),
	); err != nil {
		return entities.FileEntity{}, err
	}

	return linkMedia, nil
}

func (lms *FileService) UpdateOne(linkMediaId uuid.UUID, payload dtos.UpdateFileDto) (entities.FileEntity, error) {
	updateOneFileQuery := `
		update public.files
			set
				link_id = coalesce($2, link_id),
				file_url = coalesce($3, file_url),
				owner_id = coalesce($4, owner_id),
				updated_at = current_timestamp
			where 
				id = $1;
	`

	var updatedFile entities.FileEntity
	if err := lms.pg.Select(
		&updatedFile,
		updateOneFileQuery,
		linkMediaId.String(),
		payload.LinkId,
		payload.FileUrl,
		payload.OwnerId,
	); err != nil {
		return entities.FileEntity{}, err
	}

	return updatedFile, nil
}

func (lms *FileService) DeleteOne(linkMediaId uuid.UUID) (entities.FileEntity, error) {
	deleteOneFileQuery := `
		delete from public.files
			where id = $1
			returning id, link_id, file_url, owner_id, created_at, updated_at;
	`

	var deletedFile entities.FileEntity
	if err := lms.pg.Select(
		&deletedFile,
		deleteOneFileQuery,
		linkMediaId,
	); err != nil {
		return entities.FileEntity{}, err
	}

	return deletedFile, nil
}
