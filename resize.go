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
 	"log"
  "github.com/disintegration/imaging"
 )

 func Resize(imagePath, resultPath string, resultWidth, resultHeight int) (err error) {
   // Open a test image.
   src, err := imaging.Open(imagePath)
   if err != nil {
     log.Fatalf("failed to open image: %v", err)
   }

   // Resize the cropped image to width = 200px preserving the aspect ratio.
   src = imaging.Resize(src, resultWidth, resultHeight, imaging.Lanczos)

   // Save the resulting image as JPEG.
   err = imaging.Save(src, resultPath)
   if err != nil {
     log.Fatalf("failed to save image: %v", err)
   }

   return err
 }
