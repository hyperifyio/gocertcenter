// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package memoryrepository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hyperifyio/gocertcenter/internal/repositories/memoryrepository"
)

func TestNewCollection(t *testing.T) {
	collection := memoryrepository.NewCollection()

	// Assert collection is not nil
	assert.NotNil(t, collection, "Collection should not be nil")

	// Assert that each repository in the collection is initialized
	// This step assumes that models.Collection exposes its repositories,
	// either as public fields or via getter methods.
	// Adjust the assertions according to the actual implementation.
	assert.NotNil(t, collection.Organization, "Organization should be initialized")
	assert.NotNil(t, collection.Certificate, "Certificate should be initialized")
	assert.NotNil(t, collection.PrivateKey, "PrivateKey should be initialized")
}
