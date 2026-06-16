package handlers

import (
	"fmt"
	"path/filepath"
	"strconv"
	"time"
	"video-platform/backend/internal/models"
	"video-platform/backend/internal/services"

	"github.com/gofiber/fiber/v2"
)

type VideoHandler struct {
	videoService *services.VideoService
}

func NewVideoHandler(videoService *services.VideoService) *VideoHandler {
	return &VideoHandler{
		videoService: videoService,
	}
}

func (h *VideoHandler) UploadVideo(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)

	// Parse multipart form
	file, err := c.FormFile("video")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "video file is required",
		})
	}

	// Validate file size (max 2GB for now)
	maxSize := int64(2 * 1024 * 1024 * 1024) // 2GB
	if file.Size > maxSize {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "file size exceeds 2GB limit",
		})
	}

	// Validate file extension
	ext := filepath.Ext(file.Filename)
	allowedExts := map[string]bool{
		".mp4": true, ".mov": true, ".avi": true,
		".mkv": true, ".webm": true,
	}
	if !allowedExts[ext] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid file format. Allowed: mp4, mov, avi, mkv, webm",
		})
	}

	// Get metadata from form
	title := c.FormValue("title")
	description := c.FormValue("description")

	if title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "title is required",
		})
	}

	// Generate unique filename
	timestamp := time.Now().Unix()
	filename := fmt.Sprintf("%d_%d%s", userID, timestamp, ext)
	uploadPath := fmt.Sprintf("uploads/originals/%s", filename)

	// Save file to temporary location (in production, this would be MinIO/S3)
	if err := c.SaveFile(file, uploadPath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to save file",
		})
	}

	// Create video record
	video, err := h.videoService.CreateVideo(userID, &models.UploadVideoRequest{
		Title:       title,
		Description: description,
	}, uploadPath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create video record",
		})
	}

	// Queue for processing
	if err := h.videoService.QueueVideoForProcessing(video.ID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to queue video for processing",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(video)
}

func (h *VideoHandler) GetVideo(c *fiber.Ctx) error {
	videoID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid video ID",
		})
	}

	video, err := h.videoService.GetVideoByID(uint(videoID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "video not found",
		})
	}

	// Increment views
	go h.videoService.IncrementViews(uint(videoID))

	return c.JSON(video)
}

func (h *VideoHandler) GetVideosList(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "20"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	var userID *uint
	if userIDParam := c.Query("user_id"); userIDParam != "" {
		if id, err := strconv.ParseUint(userIDParam, 10, 32); err == nil {
			uid := uint(id)
			userID = &uid
		}
	}

	response, err := h.videoService.GetVideosList(page, pageSize, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to fetch videos",
		})
	}

	return c.JSON(response)
}

func (h *VideoHandler) DeleteVideo(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uint)
	videoID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid video ID",
		})
	}

	if err := h.videoService.DeleteVideo(uint(videoID), userID); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
