package storage

import (
	"context"

	"github.com/todoapps2021/telegrambot/internal/models"
)

type Creator interface {
	CreateTask(ctx context.Context, car models.Task) error
}

type Reader interface {
	GetTasks(ctx context.Context, limit, offset uint64) ([]*models.Task, error)
}

type Interface interface {
	Creator
	Reader
}

type TestStorage struct {
}

func (t *TestStorage) CreateTask(ctx context.Context, task models.Task) error {
	return nil
}

func (t *TestStorage) GetTasks(ctx context.Context, limit, offset uint64) ([]*models.Task, error) {
	cars := []*models.Task{
		{Name: "Test1", ID: "id_fdasfsad", Content: "FDSAfhdsjafhdslajfdsal"},
		{Name: "Test1", ID: "id_fdas3213", Content: "Fdsafdsafdshafjsda"},
		{Name: "Test1", ID: "id_fdasf5454", Content: "Gfdsagfsdklgjfdsk"},
		{Name: "Test1", ID: "id_fdasfsgfsdg", Content: "FDSAFdsafdsahfsa"},
		{Name: "Test1", ID: "id_fdasfsagfds", Content: "FDSAFdsafhdsahfdjsak"},
		{Name: "Test1", ID: "id_fdasf12355", Content: "FDSAFdsafdshafhdsajfhdsa"},
		{Name: "Test1", ID: "id_fdbvxc", Content: "fdsahfjkdshlafjdshafjdsa"},
		{Name: "Test1", ID: "id_fdasffdassad", Content: "Hfdsafdshafkdshajfhdsajkfhdsaj"},
		{Name: "Test1", ID: "id_f123xzxjdasfsad", Content: "Hellofdsafdsafdsafdsafdsafsdafdsa"},
		{Name: "Test1", ID: "id_fdasfsfdsaq3ad", Content: "Altus omnia virtualiter visums calceus est."},
	}
	if offset > uint64(len(cars)) {
		offset = uint64(len(cars))
	}
	if limit > uint64(len(cars))-offset {
		limit = uint64(len(cars)) - offset
	}

	return cars[offset : offset+limit], nil
}
