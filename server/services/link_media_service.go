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

func (lms *LinkMediaService) Create(payload dtos.CreateLinkMediaDto) (entities.LinkMediaEntity, error)

func (lms *LinkMediaService) GetMany() (entities.LinkMediaEntity, error)

func (lms *LinkMediaService) GetOne(linkMediaId uuid.UUID) (entities.LinkMediaEntity, error)

func (lms *LinkMediaService) UpdateOne(linkMediaId uuid.UUID, payload dtos.UpdateLinkMediaDto) (entities.LinkMediaEntity, error)

func (lms *LinkMediaService) DeleteOne(linkMediaId uuid.UUID) (entities.LinkMediaEntity, error)
