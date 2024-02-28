// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package commonmocks

import (
	"os"

	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/common/managers"
)

// MockFileManager is a mock implementation of managers.IFileManager for testing
type MockFileManager struct {
	mock.Mock
}

func (m *MockFileManager) Rename(oldpath, newpath string) error {
	args := m.Called(oldpath, newpath)
	return args.Error(0)
}

func (m *MockFileManager) ReadFile(fileName string) ([]byte, error) {
	args := m.Called(fileName)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockFileManager) MkdirAll(dir string, dirPerms os.FileMode) error {
	args := m.Called(dir, dirPerms)
	return args.Error(0)
}

func (m *MockFileManager) CreateTemp(dir, pattern string) (managers.IFile, error) {
	args := m.Called(dir, pattern)
	return args.Get(0).(managers.IFile), args.Error(1)
}

func (m *MockFileManager) Remove(name string) error {
	args := m.Called(name)
	return args.Error(0)
}

func (m *MockFileManager) Chmod(file string, mode os.FileMode) error {
	args := m.Called(file, mode)
	return args.Error(0)
}

func NewMockFileManager() *MockFileManager {
	return &MockFileManager{}
}

var _ managers.IFileManager = (*MockFileManager)(nil)
