package services

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/cloudflare/cloudflare-go"
)

func (service Service) CloudflareImageUpload(ctx context.Context, data cloudflare.UploadImageParams) (cloudflare.Image, error) {
	cf_resource_container := cloudflare.ResourceContainer{
		Level: cloudflare.UserRouteLevel,
		Type: cloudflare.UserType,
		Identifier: "<PASTE_IDENTIFIER_HERE>",
	}
	return service.CloudflareApi.UploadImage(ctx, &cf_resource_container, data)
}

func (service Service) GraphImageUpload(ctx context.Context, file graphql.Upload) (cloudflare.Image, error) {
	return service.CloudflareImageUpload(ctx, cloudflare.UploadImageParams{
		File: nil,
	})
}
