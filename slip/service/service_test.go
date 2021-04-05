package service_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/assert.v1"

	"github.com/payboxth/go-slip/mock"
	"github.com/payboxth/go-slip/slip"
	"github.com/payboxth/go-slip/slip/service"
)

func TestServiceCreate(t *testing.T) {
	ctx := mock.Context{}
	repo := mock.SlipRepository{
		CreateFunc: func(ctx context.Context, head *slip.Head) (string, error) {
			assert.NotNil(t, head)
			return "abc", nil
		},
	}
	storage := mock.SlipStorage{
		SaveImageFunc: func(ctx context.Context, image []byte) (string, error) {
			assert.NotNil(t, image)
			return "path/to/image", nil
		},
	}

	s := service.New(&repo, &storage)

	id, url, err := s.Create(&ctx, &slip.Head{})
	assert.NoError(t, err)
	assert.Equal(t, "abc", id)
	assert.Equal(t, "path/to/image", url)
}