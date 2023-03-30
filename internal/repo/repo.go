package repo

import (
	"TestTask/internal/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PathRepository interface {
	GetAllPaths(ctx context.Context) ([]models.Path, error)
	GetPathByID(ctx context.Context, id string) (models.Path, error)
	AddPath(ctx context.Context, path models.Path) (string, error)
	UpdatePath(ctx context.Context, id string, path models.Path) error
	DeletePath(ctx context.Context, id string) error
	GetPathByActiveLink(ctx context.Context, link string) (models.Path, bool)
	GetPathByHistoryLink(ctx context.Context, link string) (models.Path, bool)
}

type PathMongo struct {
	path *mongo.Collection
}

func NewPathMongo(db *mongo.Database) *PathMongo {
	return &PathMongo{
		path: db.Collection("paths"),
	}
}

func (repo *PathMongo) GetAllPaths(ctx context.Context) ([]models.Path, error) {

	filter := bson.M{}
	cur, err := repo.path.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var paths []models.Path
	for cur.Next(context.Background()) {
		var user models.Path
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		paths = append(paths, user)
	}
	return paths, nil
}

func (repo *PathMongo) GetPathByID(ctx context.Context, id string) (models.Path, error) {

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Path{}, err
	}

	filter := bson.M{"_id": objID}

	cur := repo.path.FindOne(ctx, filter)
	if cur.Err() != nil {
		return models.Path{}, err
	}

	var path models.Path

	if err = cur.Decode(&path); err != nil {
		return models.Path{}, err
	}

	return path, nil
}

func (repo *PathMongo) AddPath(ctx context.Context, path models.Path) (string, error) {
	return "id", nil
}

func (repo *PathMongo) UpdatePath(ctx context.Context, id string, path models.Path) error {
	return nil
}

func (repo *PathMongo) DeletePath(ctx context.Context, id string) error {
	return nil
}

func (repo *PathMongo) GetPathByActiveLink(ctx context.Context, link string) (models.Path, bool) {
	filter := bson.M{"active_link": link}

	cur := repo.path.FindOne(ctx, filter)
	if cur.Err() != nil {
		return models.Path{}, false
	}

	var path models.Path

	if err := cur.Decode(&path); err != nil {
		return models.Path{}, false
	}

	return path, true
}
func (repo *PathMongo) GetPathByHistoryLink(ctx context.Context, link string) (models.Path, bool) {
	filter := bson.M{"history_link": link}

	cur := repo.path.FindOne(ctx, filter)
	if cur.Err() != nil {
		return models.Path{}, false
	}

	var path models.Path

	if err := cur.Decode(&path); err != nil {
		return models.Path{}, false
	}

	return path, true
}
