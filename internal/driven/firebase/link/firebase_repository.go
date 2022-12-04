package firebase

import (
	"context"
	"errors"
	"fmt"
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
	req := firebasedynamiclinks.CreateShortDynamicLinkRequest{
		DynamicLinkInfo: &firebasedynamiclinks.DynamicLinkInfo{
			//AndroidInfo: &firebasedynamiclinks.AndroidInfo{
			//	AndroidMinPackageVersionCode: "1.0",
			//	AndroidPackageName:           "comm.ricardo.app",
			//},
			DomainUriPrefix: "https://ricardo.page.link",
			//IosInfo: &firebasedynamiclinks.IosInfo{
			//	IosFallbackLink: "https://www.google.com/" + linkStr + "/ios",
			//},
			Link:            fmt.Sprintf("%s/%s", l.linkPrefix, linkStr),
			ForceSendFields: []string{"DomainUriPrefix"},
		},
	}

	call := l.fbLinkService.ShortLinks.Create(&req)

	res, err := call.Do()
	if err != nil {
		return entities.ExternalLink{},
			errors.New(fmt.Sprintf("firebase dynamic link creation: %s", err))
	}

	if res.HTTPStatusCode == http.StatusOK {
		extlink := &entities.ExternalLink{Provider: provider, URL: res.ShortLink, LinkID: linkID}
		err = l.client.Create(extlink).Error
		return *extlink, err
	} else {
		return entities.ExternalLink{},
			errors.New(fmt.Sprintf("firebase dynamic link creation: %s", err))
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
