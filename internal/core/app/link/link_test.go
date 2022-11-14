package link

import (
	"context"
	"gitlab.com/ricardo134/link-service/internal/core/entities"
	"gitlab.com/ricardo134/link-service/internal/core/ports"
	"reflect"
	"testing"
)

func TestNewService(t *testing.T) {
	type args struct {
		repo       ports.LinkRepository
		extlink    ports.ExternalLinkService
		extlinkURL string
		secret     []byte
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.repo, tt.args.extlink, tt.args.extlinkURL, tt.args.secret); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_IsValid(t *testing.T) {
	type fields struct {
		repo       ports.LinkRepository
		extlink    ports.ExternalLinkService
		extlinkURL string
		secret     []byte
	}
	type args struct {
		ctx context.Context
		m   entities.MagicLink
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := service{
				repo:       tt.fields.repo,
				extlink:    tt.fields.extlink,
				extlinkURL: tt.fields.extlinkURL,
				secret:     tt.fields.secret,
			}
			got, err := p.IsValid(tt.args.ctx, tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsValid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsValid() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_ToMagic(t *testing.T) {
	type fields struct {
		repo       ports.LinkRepository
		extlink    ports.ExternalLinkService
		extlinkURL string
		secret     []byte
	}
	type args struct {
		ctx  context.Context
		link entities.Link
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.MagicLink
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := service{
				repo:       tt.fields.repo,
				extlink:    tt.fields.extlink,
				extlinkURL: tt.fields.extlinkURL,
				secret:     tt.fields.secret,
			}
			got, err := p.ToMagic(tt.args.ctx, tt.args.link)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToMagic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMagic() got = %v, want %v", got, tt.want)
			}
		})
	}
}
