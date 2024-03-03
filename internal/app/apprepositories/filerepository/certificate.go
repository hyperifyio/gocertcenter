// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package filerepository

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/common/managers"
)

// FileCertificateRepository implements models.CertificateRepository for a file system
type FileCertificateRepository struct {
	filePath    string
	certManager managers.CertificateManager
	fileManager managers.FileManager
}

func (r *FileCertificateRepository) FindAllByOrganizationAndSignedBy(organization string, certificate *big.Int) ([]appmodels.Certificate, error) {
	// TODO implement me
	panic("implement me")
}

func (r *FileCertificateRepository) FindAllByOrganization(organization string) ([]appmodels.Certificate, error) {
	// TODO implement me
	panic("implement me")
}

func (r *FileCertificateRepository) FindByOrganizationAndSerialNumber(
	organization string,
	certificate *big.Int,
) (appmodels.Certificate, error) {

	if certificate == nil {
		return nil, errors.New("no certificate serial number provided")
	}

	fileName := CertificatePemPath(r.filePath, organization, certificate)
	cert, err := ReadCertificateFile(r.fileManager, r.certManager, fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate: %w", err)
	}

	return appmodels.NewCertificate(organization, certificate, cert), nil
}

func (r *FileCertificateRepository) Save(certificate appmodels.Certificate) (appmodels.Certificate, error) {
	organization := certificate.OrganizationID()
	serialNumber := certificate.SerialNumber()
	fileName := CertificatePemPath(
		r.filePath,
		organization,
		serialNumber,
	)
	err := SaveCertificateFile(r.fileManager, r.certManager, fileName, certificate.Certificate())
	if err != nil {
		return nil, fmt.Errorf("failed to save certificate: %w", err)
	}
	return r.FindByOrganizationAndSerialNumber(organization, serialNumber)
}

// NewCertificateRepository creates a file based repository
func NewCertificateRepository(
	certManager managers.CertificateManager,
	fileManager managers.FileManager,
	filePath string,
) *FileCertificateRepository {
	return &FileCertificateRepository{
		fileManager: fileManager,
		certManager: certManager,
		filePath:    filePath,
	}
}

var _ appmodels.CertificateRepository = (*FileCertificateRepository)(nil)
