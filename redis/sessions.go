package redis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
	"github.com/jmc-quetzal/api/models"
)

type RedisSessions struct {
	Pool *redis.Pool
}

func (rs RedisSessions) Set(session *models.Session) error {
	conn := rs.Pool.Get()
	defer conn.Close()
	_, err := conn.Do("SET", session.SessionId, session.User.ID.String())
	if err != nil {
		return err
	}
	return nil
}

func (rs RedisSessions) Get(id uuid.UUID) (bool, error) {
	conn := rs.Pool.Get()
	defer conn.Close()
	exists, err := redis.Bool(conn.Do("EXISTS", id.String()))
	return exists, err
}

func (rs RedisSessions) Delete(id uuid.UUID) error {
	conn := rs.Pool.Get()
	defer conn.Close()
	_, err := conn.Do("DEL", id.String())
	if err != nil {
		return err
	}
	return nil
}
