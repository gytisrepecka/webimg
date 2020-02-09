/*
 * Copyright © 2020 Gytis Repečka (gytis@repecka.com)
 *
 * This file is part of webimg.
 *
 * webimg is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License, included
 * in the LICENSE file in this source code package.
 */

package webimg

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

func PrintHello() {
	fmt.Println("Hello there!")
}

func Watermark(imagePath, watermarkPath, resultPath string, offsetX, offsetY int, watermarkAlpha uint8) (err error) {
	// Input image
	// To-do: need to sanitize image paths!
	image1, err := os.Open(imagePath)
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}
	first, err := jpeg.Decode(image1)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer image1.Close()

	// Watermark image
	image2, err := os.Open(watermarkPath)
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}
	second, err := png.Decode(image2)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer image2.Close()

	// x, y from top-left corner
	// offset := image.Pt(100, 50)
	// log.Printf("Point: %V (%T)\n", offset, offset)

	// Bounds of input image
	inBounds := first.Bounds()
	// Bounds of watermark image
	watermarkBounds := second.Bounds()

	// To-do: offset should be configurable!
	watermarkCoordX1 := inBounds.Max.X - offsetX - watermarkBounds.Max.X // Input image x2 - offset x - watermark x2
	watermarkCoordY1 := inBounds.Max.Y - offsetY - watermarkBounds.Max.Y // Input image y2 - offset y - watermark y2

	offset := image.Pt(watermarkCoordX1, watermarkCoordY1)

	outBoundsOff := watermarkBounds.Add(offset)
	// outBoundsOff := inBounds.Sub(offset)
	/*
	  log.Printf("Bounds (input): %V (%T)\n", inBounds, inBounds)
	  log.Printf("Bounds (watermark): %V (%T)\n", watermarkBounds, watermarkBounds)
	  log.Printf("Bounds (watermark with offset): %V (%T)\n", outBoundsOff, outBoundsOff)
	*/
	log.Printf("Watermark x1: %d, y1: %d, x2: %d, y2: %d.", outBoundsOff.Min.X, outBoundsOff.Min.Y, outBoundsOff.Max.X, outBoundsOff.Max.Y)

	// Create new (output) image with size (bounds) of input image
	image3 := image.NewRGBA(inBounds)

	// Transparency
	// 0 - 255
	mask := image.NewUniform(color.Alpha{watermarkAlpha})

	draw.Draw(image3, inBounds, first, image.ZP, draw.Src)

	draw.DrawMask(image3, outBoundsOff, second, image.ZP, mask, image.ZP, draw.Over)

	third, err := os.Create(resultPath)
	if err != nil {
		log.Fatalf("failed to create: %s", err)
	}
	// jpeg.Encode(third, image3, &jpeg.Options{jpeg.DefaultQuality})
	jpeg.Encode(third, image3, &jpeg.Options{85})
	defer third.Close()

	return err
}
