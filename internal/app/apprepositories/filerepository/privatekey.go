// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package filerepository

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
	"github.com/hyperifyio/gocertcenter/internal/common/fsutils"
	"github.com/hyperifyio/gocertcenter/internal/common/managers"
)

// FilePrivateKeyRepository implements models.PrivateKeyRepository for a file system
type FilePrivateKeyRepository struct {
	filePath    string
	certManager managers.CertificateManager
	fileManager managers.FileManager
}

func (r *FilePrivateKeyRepository) FilePath() string {
	return r.filePath
}

func (r *FilePrivateKeyRepository) FindByOrganizationAndSerialNumber(
	organization string,
	certificate *big.Int,
) (appmodels.PrivateKey, error) {
	if certificate == nil {
		return nil, errors.New("no certificate serial number provided")
	}
	fileName := PrivateKeyPemPath(r.filePath, organization, certificate)
	privkey, keyType, err := ReadPrivateKeyFile(r.fileManager, r.certManager, fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read private key: %w", err)
	}
	return appmodels.NewPrivateKey(organization, certificate, keyType, privkey), nil
}

func (r *FilePrivateKeyRepository) Save(key appmodels.PrivateKey) (appmodels.PrivateKey, error) {

	organization := key.OrganizationID()
	serialNumber := key.SerialNumber()

	// Determine file path for the private key
	fileName := PrivateKeyPemPath(r.filePath, organization, serialNumber)

	// Serialize the private key into PEM format
	pemData, err := apputils.MarshalPrivateKeyAsPEM(r.certManager, key.PrivateKey())
	if err != nil {
		return nil, fmt.Errorf("failed to serialize private key to PEM: %w", err)
	}

	// Save the PEM-encoded private key to file
	err = fsutils.SaveBytes(r.fileManager, fileName, pemData, 0600, 0700)
	if err != nil {
		return nil, fmt.Errorf("failed to save private key: %w", err)
	}

	// Return the original key if saving was successful
	return key, nil
}

// NewPrivateKeyRepository creates a file based repository for private keys
func NewPrivateKeyRepository(
	certManager managers.CertificateManager,
	fileManager managers.FileManager,
	filePath string,
) *FilePrivateKeyRepository {
	return &FilePrivateKeyRepository{
		fileManager: fileManager,
		certManager: certManager,
		filePath:    filePath,
	}
}

var _ appmodels.PrivateKeyRepository = (*FilePrivateKeyRepository)(nil)
