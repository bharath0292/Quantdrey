package redisFactory

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/rueidis"
)

type RedisConfig rueidis.ClientOption

type RedisClient struct {
	client rueidis.Client
}

func NewRedisClient(config RedisConfig) (*RedisClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := rueidis.NewClient(rueidis.ClientOption(config))
	if err != nil {
		return nil, fmt.Errorf("connection failed: %w", err)
	}

	// PING to validate connection
	if err := client.Do(ctx, client.B().Ping().Build()).Error(); err != nil {
		return nil, fmt.Errorf("ping failed: %w", err)
	}

	return &RedisClient{client: client}, nil
}

func (r *RedisClient) Close() {
	r.client.Close()
}

func (r *RedisClient) KeyExists(key string) (bool, error) {
	ctx := context.Background()
	cmd := r.client.B().Exists().Key(key).Build()
	resp := r.client.Do(ctx, cmd)
	if resp.Error() != nil {
		return false, resp.Error()
	}
	count, _ := resp.ToInt64()
	return count > 0, nil
}

func (r *RedisClient) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cmd := r.client.B().Get().Key(key).Build()
	resp := r.client.Do(ctx, cmd)

	if resp.Error() != nil {
		return "", fmt.Errorf("error retrieving key '%s': %w", key, resp.Error())
	}

	return resp.ToString()
}

func (r *RedisClient) Set(key string, value string) (string, error) {
	ctx := context.Background()

	cmd := r.client.B().Set().Key(key).Value(value).Build()
	resp := r.client.Do(ctx, cmd)

	if resp.Error() != nil {
		return "", fmt.Errorf("error setting key '%s': %w", key, resp.Error())
	}

	return resp.ToString()
}

func (r *RedisClient) AddToSet(key string, members ...string) error {
	ctx := context.Background()
	cmd := r.client.B().Sadd().Key(key).Member(members...).Build()
	resp := r.client.Do(ctx, cmd)
	if resp.Error() != nil {
		return fmt.Errorf("error adding to set '%s': %w", key, resp.Error())
	}
	return nil
}

func (r *RedisClient) IsMemberOfSet(key string, member string) (bool, error) {
	ctx := context.Background()
	cmd := r.client.B().Sismember().Key(key).Member(member).Build()
	resp := r.client.Do(ctx, cmd)
	if resp.Error() != nil {
		return false, fmt.Errorf("error checking member in set '%s': %w", key, resp.Error())
	}
	val, _ := resp.ToInt64()
	return val == 1, nil
}

// Get all members of Set (SMEMBERS)
func (r *RedisClient) GetSetMembers(key string) ([]string, error) {
	ctx := context.Background()
	cmd := r.client.B().Smembers().Key(key).Build()
	resp := r.client.Do(ctx, cmd)
	if resp.Error() != nil {
		return nil, fmt.Errorf("error getting members from set '%s': %w", key, resp.Error())
	}

	return resp.AsStrSlice()
}

func (r *RedisClient) IncrementCount(key string) (int, error) {
	ctx := context.Background()
	cmd := r.client.B().Incr().Key(key).Build()
	resp, err := r.client.Do(ctx, cmd).ToInt64()
	if err != nil {
		return 0, err
	}
	return int(resp), nil
}

func (r *RedisClient) DecrementCount(key string) (int, error) {
	ctx := context.Background()
	cmd := r.client.B().Decr().Key(key).Build()
	resp, err := r.client.Do(ctx, cmd).ToInt64()
	if err != nil {
		return 0, err
	}
	return int(resp), nil
}

func (r *RedisClient) DeleteKeys(keys ...string) error {
	ctx := context.Background()
	cmd := r.client.B().Del().Key(keys...).Build()
	resp := r.client.Do(ctx, cmd)
	if resp.Error() != nil {
		return resp.Error()
	}
	return nil
}
