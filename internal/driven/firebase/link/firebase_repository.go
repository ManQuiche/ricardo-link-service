package firebase

import (
	"context"
	"fmt"
	errorsext "gitlab.com/ricardo-public/errors/v2/pkg/errors"
	"gitlab.com/ricardo134/link-service/internal/core/entities"
	"gitlab.com/ricardo134/link-service/internal/core/ports"
	"google.golang.org/api/firebasedynamiclinks/v1"
	"gorm.io/gorm"
	"net/http"
)

const provider = "firebase"

type LinkService interface {
	ports.ExternalLinkService
}

type linkService struct {
	client        *gorm.DB
	fbLinkService *firebasedynamiclinks.Service
	linkPrefix    string
}

func NewLinkService(client *gorm.DB, fbLinkService *firebasedynamiclinks.Service, linkPrefix string) LinkService {
	return linkService{client, fbLinkService, linkPrefix}
}

func (l linkService) Create(ctx context.Context, linkStr string, linkID uint) (entities.ExternalLink, error) {
	req := firebasedynamiclinks.CreateManagedShortLinkRequest{
		DynamicLinkInfo: &firebasedynamiclinks.DynamicLinkInfo{

			//AndroidInfo: &firebasedynamiclinks.AndroidInfo{
			//	AndroidMinPackageVersionCode: "1.0",
			//	AndroidPackageName:           "comm.ricardo.app",
			//},
			DomainUriPrefix: "https://ricardo.page.link",
			//IosInfo: &firebasedynamiclinks.IosInfo{
			//	IosFallbackLink: "https://www.google.com/" + linkStr + "/ios",
			//},
			Link: fmt.Sprintf("%s/%s", l.linkPrefix, linkStr),
		},
		Name: fmt.Sprintf("%s/%s", l.linkPrefix, linkStr),
	}

	call := l.fbLinkService.ManagedShortLinks.Create(&req)

	res, err := call.Do()
	if err != nil {
		return entities.ExternalLink{},
			fmt.Errorf("firebase create: %s: %w", err, errorsext.ErrInternal)
	}

	if res.HTTPStatusCode == http.StatusOK {
		extlink := &entities.ExternalLink{Provider: provider, URL: res.ManagedShortLink.Link, LinkID: linkID}
		err = l.client.Create(extlink).Error
		return *extlink, err
	} else {
		return entities.ExternalLink{},
			fmt.Errorf("firebase create: %s: %w", err, errorsext.ErrBadRequest)
	}
}

func (l linkService) Delete(ctx context.Context, extLinkID uint) error {
	// nothing to do here
	return nil
}

func (l linkService) DeleteForLinks(ctx context.Context, link entities.Link) error {
	// nothing to do here
	return nil
}
