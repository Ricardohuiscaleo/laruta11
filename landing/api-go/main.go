package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

type S3Handler struct {
	client *s3.Client
	bucket string
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type Image struct {
	Key      string    `json:"key"`
	URL      string    `json:"url"`
	Name     string    `json:"name"`
	Size     int64     `json:"size"`
	Modified time.Time `json:"modified"`
}

func main() {
	awsKey := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecret := os.Getenv("AWS_SECRET_ACCESS_KEY")
	awsRegion := os.Getenv("S3_REGION")
	bucket := os.Getenv("S3_BUCKET")

	if awsRegion == "" {
		awsRegion = "us-east-1"
	}
	if bucket == "" {
		bucket = "laruta11-images"
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(awsRegion),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(awsKey, awsSecret, "")),
	)
	if err != nil {
		log.Fatal("Error loading AWS config:", err)
	}

	handler := &S3Handler{
		client: s3.NewFromConfig(cfg),
		bucket: bucket,
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.POST("/api/s3", handler.HandleS3)
	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "service": "laruta11-api"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	log.Printf("🚀 API running on port %s", port)
	r.Run(":" + port)
}

func (h *S3Handler) HandleS3(c *gin.Context) {
	action := c.PostForm("action")

	switch action {
	case "list":
		h.listImages(c)
	case "upload":
		h.uploadImage(c)
	case "delete":
		h.deleteImage(c)
	case "test":
		h.testConnection(c)
	default:
		c.JSON(400, Response{Success: false, Error: "Invalid action"})
	}
}

func (h *S3Handler) listImages(c *gin.Context) {
	result, err := h.client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(h.bucket),
	})

	if err != nil {
		c.JSON(500, Response{Success: false, Error: err.Error()})
		return
	}

	images := []Image{}
	for _, obj := range result.Contents {
		images = append(images, Image{
			Key:      *obj.Key,
			URL:      fmt.Sprintf("https://%s.s3.amazonaws.com/%s", h.bucket, *obj.Key),
			Name:     *obj.Key,
			Size:     *obj.Size,
			Modified: *obj.LastModified,
		})
	}

	c.JSON(200, Response{Success: true, Data: gin.H{"images": images}})
}

func (h *S3Handler) uploadImage(c *gin.Context) {
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(400, Response{Success: false, Error: "No image provided"})
		return
	}
	defer file.Close()

	key := c.PostForm("custom_name")
	if key == "" {
		key = header.Filename
	}

	body, _ := io.ReadAll(file)

	_, err = h.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(h.bucket),
		Key:         aws.String(key),
		Body:        io.NopCloser(io.Reader(file)),
		ContentType: aws.String(header.Header.Get("Content-Type")),
	})

	if err != nil {
		c.JSON(500, Response{Success: false, Error: err.Error()})
		return
	}

	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", h.bucket, key)
	c.JSON(200, Response{
		Success: true,
		Message: "Image uploaded",
		Data:    gin.H{"url": url, "key": key, "size": len(body)},
	})
}

func (h *S3Handler) deleteImage(c *gin.Context) {
	key := c.PostForm("key")
	if key == "" {
		c.JSON(400, Response{Success: false, Error: "Key required"})
		return
	}

	_, err := h.client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(h.bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		c.JSON(500, Response{Success: false, Error: err.Error()})
		return
	}

	c.JSON(200, Response{Success: true, Message: "Image deleted"})
}

func (h *S3Handler) testConnection(c *gin.Context) {
	_, err := h.client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket:  aws.String(h.bucket),
		MaxKeys: aws.Int32(1),
	})

	if err != nil {
		c.JSON(500, Response{Success: false, Error: err.Error()})
		return
	}

	c.JSON(200, Response{
		Success: true,
		Message: "S3 OK",
		Data:    gin.H{"bucket": h.bucket},
	})
}
