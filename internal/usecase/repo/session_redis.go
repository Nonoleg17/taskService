package repo

import (
	"taskService/internal/entity"
	"taskService/pkg/redis"
	"time"
)

type SessionRepo struct {
	db *redis.Redis
}

func NewSessionRepo(db *redis.Redis) *SessionRepo {
	return &SessionRepo{
		db: db,
	}
}

func (sr *SessionRepo) Set(token, username string) (*entity.Session, error) {

	err := sr.db.DbConnect.Set(token, username, 24*time.Hour).Err()
	if err != nil {
		return nil, err
	}
	session := entity.Session{
		Username: username,
		Value:    token,
		Expire:   24 * time.Hour,
	}
	return &session, nil
}

func (sr *SessionRepo) Get(token string) error {
	_, err := sr.db.DbConnect.Get(token).Result()
	if err != nil {
		return err
	}
	return nil
}
