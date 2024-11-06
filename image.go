// Copyright 2024 ProntoGUI, LLC.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package golib

import (
	"image"
	_ "image/png" // Question:  how much will this impact load performance for those not using Image?
	"log"
	"os"

	"github.com/prontogui/golib/key"
)

// An image for displaying a graphic to display on the screen.  (EXPERIMENTAL)
type ImageWith struct {
	Embodiment string
	Image      []byte
	FromFile   string
	Tag        string
}

// Makes a new Image with specified field values.
func (w ImageWith) Make() *Image {
	image := &Image{}
	image.embodiment.Set(CanonizeEmbodiment(w.Embodiment))

	if len(w.FromFile) > 0 {
		loadedImage := loadImageFromFile(w.FromFile)
		if loadedImage != nil {
			image.image.Set(loadedImage.Pix)
		}
	} else {
		image.image.Set(w.Image)
	}

	image.tag.Set(w.Tag)
	return image
}

// An image for displaying a graphic to display on the screen.  (EXPERIMENTAL)
type Image struct {
	// Mix-in the common guts for primitives
	PrimitiveBase

	embodiment StringField
	image      BlobField
	tag        StringField
}

// Creates a new Image from a file.  (EXPERIMENTAL)
func NewImage(fromFile string) *Image {
	return ImageWith{FromFile: fromFile}.Make()
}

// Prepares the primitive for tracking pending updates to send to the app and
// for injesting updates from the app.  This is used internally by this library
// and normally should not be called by users of the library.
func (image *Image) PrepareForUpdates(pkey key.PKey, onset key.OnSetFunction) {

	image.InternalPrepareForUpdates(pkey, onset, func() []FieldRef {
		return []FieldRef{
			{key.FKey_Embodiment, &image.embodiment},
			{key.FKey_Image, &image.image},
			{key.FKey_Tag, &image.tag},
		}
	})
}

// Returns a JSON string specifying the embodiment to use for this primitive.
func (image *Image) Embodiment() string {
	return image.embodiment.Get()
}

// Sets a JSON string specifying the embodiment to use for this primitive.
func (image *Image) SetEmbodiment(s string) *Image {
	image.embodiment.Set(CanonizeEmbodiment(s))
	return image
}

// Returns the binary data for the image
func (image *Image) Image() []byte {
	return image.image.Get()
}

// Sets the binary data for the image
func (image *Image) SetImage(data []byte) *Image {
	image.image.Set(data)
	return image
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

// Returns an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Commands as Table cells.
func (image *Image) Tag() string {
	return image.tag.Get()
}

// Sets an optional and arbitrary string to keep with this primitive.  This is useful for
// identification later on, such as using Commands as Table cells.
func (image *Image) SetTag(s string) *Image {
	image.tag.Set(s)
	return image
}
