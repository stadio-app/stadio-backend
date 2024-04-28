package services

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func (service Service) GraphImageUpload(
	ctx context.Context,
	file graphql.Upload,
	params uploader.UploadParams,
) (*uploader.UploadResult, error) {
	return service.Cloudinary.Upload.Upload(ctx, file.File, params)
}
