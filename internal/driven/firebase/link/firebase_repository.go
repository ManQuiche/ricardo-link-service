package firebase

import (
	"context"
	"gitlab.com/ricardo134/link-service/internal/core/entities"
	"gitlab.com/ricardo134/link-service/internal/core/ports"
	"google.golang.org/api/firebasedynamiclinks/v1"
)

type LinkService interface {
	ports.ExternalLinkService
}

type linkService struct {
	fbLinkService *firebasedynamiclinks.Service
}

func NewLinkService(fbLinkService *firebasedynamiclinks.Service) LinkService {
	return linkService{fbLinkService}
}

func (l linkService) Create(ctx context.Context, link entities.Link) (entities.ExternalLink, error) {
	call := l.fbLinkService.ShortLinks.Create(&firebasedynamiclinks.CreateShortDynamicLinkRequest{
		DynamicLinkInfo: &firebasedynamiclinks.DynamicLinkInfo{
			AnalyticsInfo:     nil,
			AndroidInfo:       nil,
			DesktopInfo:       nil,
			DomainUriPrefix:   "",
			DynamicLinkDomain: "",
			IosInfo:           nil,
			Link:              "",
			NavigationInfo:    nil,
			SocialMetaTagInfo: nil,
			ForceSendFields:   nil,
			NullFields:        nil,
		},
		LongDynamicLink: "",
		SdkVersion:      "",
		Suffix: &firebasedynamiclinks.Suffix{
			CustomSuffix:    "",
			Option:          "",
			ForceSendFields: nil,
			NullFields:      nil,
		},
		ForceSendFields: nil,
		NullFields:      nil,
	})

	return entities.ExternalLink{}, nil
}

func (l linkService) Delete(ctx context.Context, extLinkID uint) error {
	// nothing to do here
	return nil
}

func (l linkService) DeleteForLinks(ctx context.Context, link entities.Link) error {
	// nothing to do here
	return nil
}
