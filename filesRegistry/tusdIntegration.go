package filesRegistry

import (
	"context"

	"github.com/tus/tusd/pkg/filestore"
	tusd "github.com/tus/tusd/pkg/handler"
)

// for deeper integration

type tusdFileStorageImpl struct {
	filestore.FileStore
}

func (t tusdFileStorageImpl) NewUpload(ctx context.Context, info tusd.FileInfo) (upload tusd.Upload, err error) {
	panic("implement me")
}

func (t tusdFileStorageImpl) GetUpload(ctx context.Context, id string) (upload tusd.Upload, err error) {
	panic("implement me")
}
