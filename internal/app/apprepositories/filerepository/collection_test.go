// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package filerepository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hyperifyio/gocertcenter/internal/app/apprepositories/filerepository"
	"github.com/hyperifyio/gocertcenter/internal/common/commonmocks"
)

func TestNewCollection(t *testing.T) {

	fileManager := &commonmocks.MockFileManager{}
	certManager := &commonmocks.MockCertificateManager{}

	// Example file path
	filePath := "test_data"

	// Call NewCollection
	collection := filerepository.NewCollection(certManager, fileManager, filePath)

	// Assertions to ensure that the collection and its repositories are not nil
	assert.NotNil(t, collection, "Expected non-nil collection")
	assert.NotNil(t, collection.Organization, "Expected non-nil Organization service")
	assert.NotNil(t, collection.Certificate, "Expected non-nil Certificate service")
	assert.NotNil(t, collection.PrivateKey, "Expected non-nil PrivateKey service")

	// Additional checks can include verifying that the repositories are correctly initialized with the filePath
	// This step requires access to the internal state of the repositories or using reflection if not directly accessible
	// For example:
	// assert.Equal(t, filePath, collection.OrganizationRepository().(*filerepository.OrganizationRepository).filePath, "OrganizationRepository filePath mismatch")
}
