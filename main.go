package main

import (
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

	maxWidth, err := strconv.Atoi(os.Args[3])

	if err != nil {
		log.Println("Invalid argument `maxWidth`:" + err.Error())
		os.Exit(128)
	}

	maxHeight, err := strconv.Atoi(os.Args[4])

	if err != nil {
		log.Println("Invalid argument `maxHeight`:" + err.Error())
		os.Exit(128)
	}

	maxColumns, err := strconv.Atoi(os.Args[5])

	if err != nil {
		log.Println("Invalid argument `maxColumns`:" + err.Error())
		os.Exit(128)
	}

	outputFileName := os.Args[6]

	metadata, err := vtsg.GetMetadata(fileName)

	if err != nil {
		log.Fatalln("Error while reading video metadata: " + err.Error())
	}

	frameCount := metadata.Duration / interval
	frameWidth, frameHeight, err := vtsg.CalculateFrameDimensions(metadata.Width, metadata.Height, maxWidth, maxHeight)

	if err != nil {
		log.Fatalln("Error calculating sprite frame dimensions: " + err.Error())
	}

	gridColumns := min(frameCount, maxColumns)
	gridRows := int(math.Ceil(float64(frameCount) / float64(maxColumns)))
	spriteBuffer, err := vtsg.GenerateSprite(fileName, interval, gridColumns, gridRows, frameWidth, frameHeight)

	if err != nil {
		log.Fatalln("Error calculating creating sprite: " + err.Error())
	}

	err = os.WriteFile(outputFileName, spriteBuffer.Bytes(), 0777)

	if err != nil {
		log.Fatalln("Error writing file: " + err.Error())
	}

	elapsed := time.Since(start)
	seconds := elapsed.Seconds()
	fmt.Printf("Time taken: %.2f seconds\n", seconds)
}
