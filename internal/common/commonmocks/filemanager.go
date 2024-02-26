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

func (m *MockFileManager) ReadBytes(fileName string) ([]byte, error) {
	args := m.Called(fileName)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockFileManager) SaveBytes(fileName string, data []byte, filePerms, dirPerms os.FileMode) error {
	args := m.Called(fileName, data, filePerms, dirPerms)
	return args.Error(0)
}

func NewMockFileManager() *MockFileManager {
	return &MockFileManager{}
}

var _ managers.IFileManager = (*MockFileManager)(nil)
