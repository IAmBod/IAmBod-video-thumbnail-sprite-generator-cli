package main

import (
	"encoding/json"
	"fmt"
	vtsg "github.com/IAmBod/video-thumbnail-sprite-generator"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()

	fileName := os.Args[1]
	interval, err := strconv.Atoi(os.Args[2])

	if err != nil {
		log.Println("Invalid argument `interval`:" + err.Error())
		os.Exit(128)
	}

	maxTileWidth, err := strconv.Atoi(os.Args[3])

	if err != nil {
		log.Println("Invalid argument `maxTileWidth`:" + err.Error())
		os.Exit(128)
	}

	maxTileHeight, err := strconv.Atoi(os.Args[4])

	if err != nil {
		log.Println("Invalid argument `maxTileHeight`:" + err.Error())
		os.Exit(128)
	}

	maxColumns, err := strconv.Atoi(os.Args[5])

	if err != nil {
		log.Println("Invalid argument `maxColumns`:" + err.Error())
		os.Exit(128)
	}

	outputFileName := os.Args[6]
	outputMetadataFileName := os.Args[7]

	videoMetadata, err := vtsg.GetMetadata(fileName)

	if err != nil {
		log.Fatalln("Error while reading video metadata: " + err.Error())
	}

	duration := videoMetadata.Duration
	tileCount := int(math.Ceil(duration / float64(interval)))
	tileWidth, tileHeight, err := vtsg.CalculateTileDimensions(videoMetadata.Width, videoMetadata.Height, maxTileWidth, maxTileHeight)

	if err != nil {
		log.Fatalln("Error calculating tile dimensions: " + err.Error())
	}

	columns := min(tileCount, maxColumns)
	rows := int(math.Ceil(float64(tileCount) / float64(maxColumns)))
	storyboardBuffer, err := vtsg.GenerateStoryboardImage(fileName, interval, columns, rows, tileWidth, tileHeight)

	if err != nil {
		log.Fatalln("Error creating storyboard: " + err.Error())
	}

	err = os.WriteFile(outputFileName, storyboardBuffer.Bytes(), 0777)

	if err != nil {
		log.Fatalln("Error writing storyboard: " + err.Error())
	}

	storyboardMetadata := vtsg.GenerateStoryboardMetadata("%url%", interval, columns, tileWidth, tileHeight, duration)
	storyBoardJson, err := json.Marshal(storyboardMetadata)

	if err != nil {
		log.Fatalln("Error writing storyboard metadata: " + err.Error())
	}

	err = os.WriteFile(outputMetadataFileName, storyBoardJson, 0777)

	elapsed := time.Since(start)
	seconds := elapsed.Seconds()
	fmt.Printf("Time taken: %.2f seconds\n", seconds)
}
