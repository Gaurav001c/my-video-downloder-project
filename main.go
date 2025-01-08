package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kkdai/youtube/v2"
)

func main() {
	// Ask user for the video link
	var videoURL string
	fmt.Print("Enter the YouTube video URL: ")
	fmt.Scanln(&videoURL)

	// Initialize the YouTube client
	client := youtube.Client{}

	// Fetch video information
	video, err := client.GetVideo(videoURL)
	if err != nil {
		log.Fatalf("Error fetching video information: %v\n", err)
	}

	// Filter formats to ensure we get both audio and video
	format := video.Formats.WithAudioChannels() // Ensures we pick a format with audio
	if len(format) == 0 {
		log.Fatalf("No suitable format with audio found for the video.")
	}

	// Select the first available format with both video and audio
	selectedFormat := format[0]

	// Get the video stream
	stream, _, err := client.GetStream(video, &selectedFormat)
	if err != nil {
		log.Fatalf("Error getting video stream: %v\n", err)
	}

	// Create a file to save the video
	fileName := fmt.Sprintf("%s.mp4", video.Title)
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Error creating file: %v\n", err)
	}
	defer file.Close()

	// Download and save the video
	_, err = file.ReadFrom(stream)
	if err != nil {
		log.Fatalf("Error saving video: %v\n", err)
	}

	fmt.Printf("Video downloaded successfully as: %s\n", fileName)
}
