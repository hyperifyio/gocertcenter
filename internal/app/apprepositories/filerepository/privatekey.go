// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package filerepository

import (
	"errors"
	"fmt"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
	"github.com/hyperifyio/gocertcenter/internal/common/managers"
)

// PrivateKeyRepository implements models.IPrivateKeyService for a file system
type PrivateKeyRepository struct {
	filePath    string
	certManager managers.ICertificateManager
	fileManager managers.IFileManager
}

func (r *PrivateKeyRepository) GetFilePath() string {
	return r.filePath
}

func (r *PrivateKeyRepository) GetExistingPrivateKey(
	organization string,
	certificates []appmodels.ISerialNumber,
) (appmodels.IPrivateKey, error) {
	if len(certificates) <= 0 {
		return nil, errors.New("no certificate serial numbers provided")
	}
	fileName := GetPrivateKeyPemPath(r.filePath, organization, certificates)
	privkey, keyType, err := ReadPrivateKeyFile(r.fileManager, r.certManager, fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key: %w", err)
	}
	return appmodels.NewPrivateKey(organization, certificates, keyType, privkey), nil
}

func (r *PrivateKeyRepository) CreatePrivateKey(key appmodels.IPrivateKey) (appmodels.IPrivateKey, error) {

	organization := key.GetOrganizationID()
	certificates := key.GetParents()
	serialNumber := key.GetSerialNumber()

	// Determine file path for the private key
	fileName := GetPrivateKeyPemPath(r.filePath, organization, append(certificates, serialNumber))

	// Serialize the private key into PEM format
	pemData, err := apputils.MarshalPrivateKeyAsPEM(r.certManager, key.GetPrivateKey())
	if err != nil {
		return nil, fmt.Errorf("failed to serialize private key to PEM: %w", err)
	}

	// Save the PEM-encoded private key to file
	err = r.fileManager.SaveBytes(fileName, pemData, 0600, 0700)
	if err != nil {
		return nil, fmt.Errorf("failed to save private key: %w", err)
	}

	// Return the original key if saving was successful
	return key, nil
}

// NewPrivateKeyRepository creates a file based repository for private keys
func NewPrivateKeyRepository(
	certManager managers.ICertificateManager,
	fileManager managers.IFileManager,
	filePath string,
) *PrivateKeyRepository {
	return &PrivateKeyRepository{
		fileManager: fileManager,
		certManager: certManager,
		filePath:    filePath,
	}
}

var _ appmodels.IPrivateKeyService = (*PrivateKeyRepository)(nil)
