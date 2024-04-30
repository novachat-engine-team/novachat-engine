/*
 * Copyright (c) 2021-present,  NovaChat-Engine.
 *  All rights reserved.
 *
 * @Author: Coder (coderxw@gmail.com)
 * @Time : 2021/7/2 9:32
 * @File : fru.go
 */

package partfs

type FRU interface {
	Config(dbPath string) error
	PartFile(fileName string) error
	PartDir(path string) error
	ClearFile(fileName string) error
	ClearDir(path string) error
	UnclearFile(fileName string) error
	UnclearDir(path string) error
}

func DefaultFRU() FRU {
	return &_fru{}
}

type _fru struct {}
func (m *_fru) Config(dbPath string) error {
	return nil
}

func (m *_fru) PartFile(fileName string) error {
	return nil
}

func (m *_fru) PartDir(path string) error {
	return nil
}

func (m *_fru) ClearFile(fileName string) error {
	return nil
}

func (m *_fru) ClearDir(path string) error {
	return nil
}

func (m *_fru) UnclearFile(fileName string) error {
	return nil
}

func (m *_fru) UnclearDir(path string) error {
	return nil
}