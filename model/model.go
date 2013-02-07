// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
	Package model defines the basic
	data structures of the docs engine.
*/
package model

import (
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
)

type RepositoryItem struct {
	Path       string
	Files      []RepositoryItemFile
	ChildItems []RepositoryItem
}

type RepositoryItemFile struct {
	Path string
}

func NewRepositoryItemFile(path string) RepositoryItemFile {
	return RepositoryItemFile{
		Path: path,
	}
}

func NewRepositoryItem(path string, files []RepositoryItemFile, childItems []RepositoryItem) RepositoryItem {
	return RepositoryItem{
		Path:       path,
		Files:      files,
		ChildItems: childItems,
	}
}

func (item *RepositoryItem) GetHash() string {
	itemBytes, readFileErr := ioutil.ReadFile(item.Path)
	if readFileErr != nil {
		return ""
	}

	sha1Hash := sha1.New()
	sha1Hash.Write(itemBytes)
	hashBytes := sha1Hash.Sum(nil)

	return string(hex.EncodeToString(hashBytes[0:6]))
}

func (item *RepositoryItem) String() string {
	s := item.Path + "(Hash: " + item.GetHash() + ")\n"

	s += "\n"
	s += "Files:\n"
	if len(item.Files) > 0 {
		for _, file := range item.Files {
			s += " - " + file.Path + "\n"
		}
	} else {
		s += "<none>\n"
	}

	s += "\n"
	s += "ChildItems:\n"
	if len(item.ChildItems) > 0 {
		for _, child := range item.ChildItems {
			s += child.String()
		}
	} else {
		s += "<none>\n"
	}
	s += "\n"

	return s
}
