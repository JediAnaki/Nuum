package processor

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type VideoProcessor struct {
	outputDir string
}

type ProcessingOptions struct {
	InputPath  string
	OutputDir  string
	VideoID    uint
}

type ProcessedResult struct {
	Quality string
	Path    string
	Size    int64
	Bitrate int
}

func NewVideoProcessor(outputDir string) *VideoProcessor {
	return &VideoProcessor{
		outputDir: outputDir,
	}
}

// Process video into multiple qualities
func (p *VideoProcessor) Process(opts ProcessingOptions) ([]ProcessedResult, error) {
	results := []ProcessedResult{}

	// Define quality presets
	qualities := []struct {
		Name       string
		Resolution string
		Bitrate    string
		VideoBR    int
	}{
		{"360p", "640x360", "800k", 800},
		{"480p", "854x480", "1400k", 1400},
		{"720p", "1280x720", "2800k", 2800},
		{"1080p", "1920x1080", "5000k", 5000},
	}

	// Create output directory for this video
	videoOutputDir := filepath.Join(opts.OutputDir, fmt.Sprintf("video_%d", opts.VideoID))
	if err := os.MkdirAll(videoOutputDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create output directory: %w", err)
	}

	// Process each quality
	for _, q := range qualities {
		outputPath := filepath.Join(videoOutputDir, fmt.Sprintf("%s.mp4", q.Name))

		// FFmpeg command for transcoding
		// -i: input file
		// -vf scale: resize video
		// -c:v libx264: use H.264 codec
		// -preset: encoding speed/quality tradeoff
		// -b:v: video bitrate
		// -c:a aac: audio codec
		// -b:a: audio bitrate
		cmd := exec.Command("ffmpeg",
			"-i", opts.InputPath,
			"-vf", fmt.Sprintf("scale=%s", q.Resolution),
			"-c:v", "libx264",
			"-preset", "fast",
			"-b:v", q.Bitrate,
			"-c:a", "aac",
			"-b:a", "128k",
			"-movflags", "+faststart", // Enable streaming
			"-y", // Overwrite output file
			outputPath,
		)

		log.Printf("Processing %s quality for video %d", q.Name, opts.VideoID)

		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Printf("FFmpeg error for %s: %s", q.Name, string(output))
			continue // Skip this quality but continue with others
		}

		// Get file size
		fileInfo, err := os.Stat(outputPath)
		if err != nil {
			log.Printf("Failed to stat output file %s: %v", outputPath, err)
			continue
		}

		results = append(results, ProcessedResult{
			Quality: q.Name,
			Path:    outputPath,
			Size:    fileInfo.Size(),
			Bitrate: q.VideoBR,
		})

		log.Printf("Successfully processed %s quality for video %d", q.Name, opts.VideoID)
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("failed to process any quality variants")
	}

	return results, nil
}

// GenerateThumbnail creates a thumbnail from the video at specific timestamp
func (p *VideoProcessor) GenerateThumbnail(inputPath string, outputPath string, timestamp string) error {
	// Extract frame at specific time (default: 00:00:01)
	if timestamp == "" {
		timestamp = "00:00:01"
	}

	cmd := exec.Command("ffmpeg",
		"-i", inputPath,
		"-ss", timestamp,
		"-vframes", "1",
		"-vf", "scale=320:180",
		"-y",
		outputPath,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to generate thumbnail: %w, output: %s", err, string(output))
	}

	return nil
}

// GetVideoMetadata extracts video metadata using ffprobe
func (p *VideoProcessor) GetVideoMetadata(inputPath string) (map[string]interface{}, error) {
	cmd := exec.Command("ffprobe",
		"-v", "quiet",
		"-print_format", "json",
		"-show_format",
		"-show_streams",
		inputPath,
	)

	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get video metadata: %w", err)
	}

	// Parse JSON output
	// For simplicity, returning raw for now
	// In production, parse this into structured data
	metadata := map[string]interface{}{
		"raw": string(output),
	}

	return metadata, nil
}
