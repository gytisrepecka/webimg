/*
 * Copyright © 2020 Gytis Repečka (gytis@repecka.com)
 *
 * This file is part of webimg.
 *
 * webimg is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License, included
 * in the LICENSE file in this source code package.
 */

package main

import (
	"code.gyt.is/webimg"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	// webimg.PrintHello()

	// Set some parameters
	timestampLayout := "2006-01-02 15:04:05.000 (MST Z07:00)"
	currentTime := time.Now()

	// Logging
	logFile, err := os.OpenFile("webimg.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	// Log beginning of program
	log.Printf("Started webimg: %s.\n", currentTime.Format(timestampLayout))

	cmdLineArgs := os.Args
	/*
	   Subcommands
	   Args[0] - executable path;
	   Args[1] - subcommand.
	*/

	// Main subcommands
	watermarkCommand := flag.NewFlagSet("watermark", flag.ExitOnError)

	// Subcommand "watermark" flag pointers
	watermarkImagePtr := watermarkCommand.String("image", "", "Path to image (JPG) file to watermark.")

	// General options usable with other commands
	outputVersion := flag.Bool("v", false, "Output the current version")
	outputHelp := flag.Bool("h", false, "Show help information")

	flag.Parse()

	if *outputVersion {
		webimg.OutputVersion()
		os.Exit(0)
	} else if *outputHelp {
		fmt.Println("General commands:")
		flag.PrintDefaults()
		fmt.Println("Watermark commands:")
		watermarkCommand.PrintDefaults()

		// Check if subcommand (watermark) was provided
	} else if len(cmdLineArgs) >= 2 {
		// Switch based on subcommand
		switch cmdLineArgs[1] {
		case "watermark":
			watermarkCommand.Parse(cmdLineArgs[2:])

			if *watermarkImagePtr != "" {
				imageToWatermark := fmt.Sprintf("%s", *watermarkImagePtr)
				fmt.Printf("Image to watermark: %s\n", imageToWatermark)

				// Input image, watermark image, result image, bottom-right offset X, bottom-right offset Y, watermark alpha
				doWatermark := webimg.Watermark(imageToWatermark, "watermark_inretio-logo.png", "result_img.jpg", 30, 30, 70)
				if doWatermark != nil {
					fmt.Println("There was an error watermarking image...")
				}
			} else {
				fmt.Println("No image given to watermark!")
			}
		default:
			fmt.Println("Wrong subcommand provided!")
		}
	} else {
		fmt.Println("No work given. Holiday time!")
	}

	// Log end of program
	// time.Sleep(2 * time.Second)
	currentTime = time.Now()
	log.Printf("Finished webimg: %s.\n", currentTime.Format(timestampLayout))
	log.Println("--------------------")
}
