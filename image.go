// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/prontogui/golib/field"
	"github.com/prontogui/golib/key"
)

type ImageWith struct {
	Embodiment string
	Image      []byte
	FromFile   string
}

// Makes a new Image with specified field values.
func (w ImageWith) Make() *Image {
	image := &Image{}
	image.embodiment.Set(w.Embodiment)

	if len(w.FromFile) > 0 {
		loadedImage := loadImageFromFile(w.FromFile)
		if loadedImage != nil {
			image.image.Set(loadedImage.Pix)
		}
	} else {
		image.image.Set(w.Image)
	}
	return image
}

type Image struct {
	// Mix-in the common guts for primitives
	Reserved

	embodiment field.String
	image      field.Blob
}

func (image *Image) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	image.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &image.embodiment},
			{key.FKey_Image, &image.image},
		}
	})
}

func (image *Image) Embodiment() string {
	return image.embodiment.Get()
}

func (image *Image) SetEmbodiment(s string) {
	image.embodiment.Set(s)
}

func (image *Image) Image() []byte {
	return image.image.Get()
}

func (image *Image) SetImage(data []byte) {
	image.image.Set(data)
}

func loadImageFromFile(filePath string) *image.RGBA {
	imgFile, err := os.Open(filePath)
	if err != nil {
		log.Println("Cannot read file:", err)
		return nil
	}
	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		log.Println("Cannot decode file:", err)
		return nil
	}

	return img.(*image.RGBA)
}
