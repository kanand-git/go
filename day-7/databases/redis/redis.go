package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

//https://redis.uptrace.dev/guide/go-redis.html

// RedisClient is a wrapper for the Redis client.
type RedisClient struct {
	client *redis.Client
}

func main() {
	opts := &redis.Options{
		Addr:         "localhost:6379", // Use the appropriate address if Redis is running elsewhere.
		Username:     "default",
		Password:     "",              // no password set
		MinIdleConns: 3,               // Minimum number of idle connections.
		DialTimeout:  5 * time.Second, // Timeout for establishing new connections.
		ReadTimeout:  3 * time.Second, // Timeout for socket reads.
		WriteTimeout: 3 * time.Second, // Timeout for socket writes.

	}
	// Create a new Redis client.
	client := redis.NewClient(opts)
	defer client.Close()

	// Create RedisClient instance.
	redisClient := &RedisClient{
		client: client,
	}

	// Create a user
	userID := "user:1"
	ctx := context.Background()
	err := redisClient.CreateUser(ctx, userID, "Alice", "alice@example.com")
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}
	fmt.Printf("User %s created\n", userID)

	// Create posts for the user
	postID1, err := redisClient.CreatePost(ctx, userID, "This is my first post!")
	if err != nil {
		log.Fatalf("Failed to create post: %v", err)
	}

	fmt.Printf("Post %s created by user %s\n", postID1, userID)

	postID2, err := redisClient.CreatePost(ctx, userID, "Hello, Redis!")
	if err != nil {
		log.Fatalf("Failed to create post: %v", err)
	}
	fmt.Printf("Post %s created by user %s\n", postID2, userID)

	// View user's posts
	posts, err := redisClient.GetUserPosts(ctx, userID)
	if err != nil {
		log.Fatalf("Failed to get user posts: %v", err)
	}
	fmt.Println("User posts:", posts)

	// Like a post
	err = redisClient.LikePost(ctx, postID1)
	if err != nil {
		log.Fatalf("Failed to like post: %v", err)
	}
	fmt.Printf("Post %s liked\n", postID1)

	// Retrieve top liked posts
	topPosts, err := redisClient.GetTopLikedPosts(ctx, 10)
	if err != nil {
		log.Fatalf("Failed to get top liked posts: %v", err)
	}
	fmt.Println("Top liked posts:", topPosts)

}

// CreateUser creates a new user.
func (r *RedisClient) CreateUser(ctx context.Context, userID, name, email string) error {
	err := r.client.HSet(ctx, userID, "name", name, "email", email).Err()
	if err != nil {
		return err
	}
	return nil
}

// CreatePost creates a new post and associates it with a user.
func (r *RedisClient) CreatePost(ctx context.Context, userID, content string) (string, error) {
	postID := fmt.Sprintf("post:%s", uuid.NewString()) // Unique post ID
	err := r.client.Set(ctx, postID, content, 0).Err()
	if err != nil {
		return "", err
	}
	err = r.client.RPush(ctx, userID+":posts", postID).Err()
	if err != nil {
		return "", err
	}
	return postID, nil
}

// GetUserPosts retrieves posts made by a user.
func (r *RedisClient) GetUserPosts(ctx context.Context, userID string) ([]string, error) {
	return r.client.LRange(ctx, userID+":posts", 0, -1).Result()
}

// LikePost increments the like count of a post.
func (r *RedisClient) LikePost(ctx context.Context, postID string) error {
	return r.client.ZIncrBy(ctx, "posts:likes", 1, postID).Err()
}

// GetTopLikedPosts retrieves the top liked posts.
func (r *RedisClient) GetTopLikedPosts(ctx context.Context, count int64) ([]string, error) {
	return r.client.ZRevRange(ctx, "posts:likes", 0, count-1).Result()
}
