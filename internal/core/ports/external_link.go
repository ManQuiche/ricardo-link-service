package ports

import (
	"context"
	"gitlab.com/ricardo134/link-service/internal/core/entities"
)

type ExternalLinkService interface {
	Create(ctx context.Context, linkStr string, linkID uint) (entities.ExternalLink, error)
	Delete(ctx context.Context, extLinkID uint) error
	DeleteForLinks(ctx context.Context, link entities.Link) error
}
