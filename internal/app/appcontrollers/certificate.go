// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appcontrollers

import (
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/hyperifyio/gocertcenter/internal/common/managers"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
)

// CertCertificateController implements appmodels.CertificateController to control
// operations for certificate models.
//
// It utilizes appmodels.CertificateRepository implementation to abstract
// and inject the underlying storage mechanism (e.g., database, memory). This
// design promotes separation of concerns by decoupling the business logic from
// the specific details of data persistence.
type CertCertificateController struct {
	serialNumber *big.Int
	model        appmodels.Certificate

	// parentOrganizationController is the parent organization controller
	parentOrganizationController appmodels.OrganizationController

	// parentCertificateController  is an optional parent certificate controller
	parentCertificateController appmodels.CertificateController

	privateKeyController appmodels.PrivateKeyController

	certManager   managers.CertificateManager
	randomManager managers.RandomManager

	certificateRepository appmodels.CertificateRepository
	privateKeyRepository  appmodels.PrivateKeyRepository

	expiration time.Duration
}

func (r *CertCertificateController) OrganizationController() appmodels.OrganizationController {
	if r.parentOrganizationController == nil {
		panic("[CertCertificateController:OrganizationController]: No parent organization controller")
	}
	return r.parentOrganizationController
}

func (r *CertCertificateController) ApplicationController() appmodels.ApplicationController {
	if r.parentOrganizationController == nil {
		panic("[CertCertificateController.ApplicationController]: No parent organization controller")
	}
	return r.parentOrganizationController.ApplicationController()
}

func (r *CertCertificateController) OrganizationID() string {
	if r.parentOrganizationController == nil {
		panic("[CertCertificateController.OrganizationID]: No parent organization controller")
	}
	return r.parentOrganizationController.OrganizationID()
}

func (r *CertCertificateController) Organization() appmodels.Organization {
	return r.parentOrganizationController.Organization()
}

func (r *CertCertificateController) Certificate() appmodels.Certificate {
	return r.model
}

func (r *CertCertificateController) ChildCertificateCollection(certificateType string) ([]appmodels.Certificate, error) {
	organization := r.OrganizationID()
	path := append(r.model.Parents(), r.serialNumber)
	list, err := r.certificateRepository.FindAllByOrganizationAndSerialNumbers(
		organization,
		path,
	)

	if certificateType != "" {
		list = apputils.FilterCertificatesByType(list, certificateType)
	}

	if err != nil {
		return nil, fmt.Errorf("[%s@%s:ChildCertificateCollection]: failed: %w", r.serialNumber.String(), organization, err)
	}
	return list, nil
}

func (r *CertCertificateController) ChildCertificate(serialNumber *big.Int) (appmodels.Certificate, error) {
	organization := r.OrganizationID()
	if r.certificateRepository == nil {
		return nil, fmt.Errorf("[%s@%s:ChildCertificate:%s]: No parent certificateRepository", r.serialNumber.String(), organization, serialNumber.String())
	}
	if r.model == nil {
		return nil, fmt.Errorf("[%s@%s:ChildCertificate:%s]: No parent model", r.serialNumber.String(), organization, serialNumber.String())
	}
	model, err := r.certificateRepository.FindByOrganizationAndSerialNumbers(
		organization,
		append(r.model.Parents(), r.serialNumber, serialNumber),
	)
	if err != nil {
		return nil, fmt.Errorf("[%s@%s:ChildCertificate:%s]: failed: %w", r.serialNumber.String(), organization, serialNumber.String(), err)
	}
	return model, nil
}

func (r *CertCertificateController) ChildCertificateController(serialNumber *big.Int) (appmodels.CertificateController, error) {
	organization := r.OrganizationID()
	model, err := r.ChildCertificate(serialNumber)
	if err != nil {
		return nil, fmt.Errorf("[%s@%s:ChildCertificateController:%s]: could not find: %w", r.serialNumber.String(), organization, serialNumber.String(), err)
	}
	return NewCertificateController(
		r.OrganizationController(),
		r,
		serialNumber,
		model,
		r.certificateRepository,
		r.privateKeyRepository,
		r.certManager,
		r.randomManager,
		r.expiration,
	), nil
}

func (r *CertCertificateController) ParentCertificate() appmodels.Certificate {
	if r.parentCertificateController == nil {
		return nil
	}
	return r.parentCertificateController.Certificate()
}

func (r *CertCertificateController) ParentCertificateController() appmodels.CertificateController {
	if r.parentCertificateController == nil {
		return nil
	}
	return r.parentCertificateController
}

func (r *CertCertificateController) PrivateKey() (appmodels.PrivateKey, error) {
	organization := r.OrganizationID()
	if r.privateKeyRepository == nil {
		return nil, fmt.Errorf("[%s@%s:PrivateKey]: no private key repository", r.serialNumber, organization)
	}
	model, err := r.privateKeyRepository.FindByOrganizationAndSerialNumbers(
		organization,
		append(r.model.Parents(), r.serialNumber),
	)
	if err != nil {
		return nil, fmt.Errorf("[%s@%s:PrivateKey]: failed: %w", r.serialNumber, organization, err)
	}
	return model, nil
}

func (r *CertCertificateController) PrivateKeyController() (appmodels.PrivateKeyController, error) {
	organization := r.OrganizationID()
	model, err := r.PrivateKey()
	if err != nil {
		return nil, fmt.Errorf("[%s@%s:PrivateKeyController]: failed: %w", r.serialNumber, organization, err)
	}
	return NewPrivateKeyController(
		model,
		r,
		r.privateKeyRepository,
	), nil
}

func (r *CertCertificateController) SetExpirationDuration(expiration time.Duration) {
	r.expiration = expiration
}

func (r *CertCertificateController) NewIntermediateCertificate(commonName string) (appmodels.Certificate, appmodels.PrivateKey, error) {

	organization := r.OrganizationID()

	parentPrivateKey, err := r.PrivateKey()
	if err != nil {
		return nil, nil, fmt.Errorf("[%s@%s:NewIntermediateCertificate:%s]: failed to get private key: %w", r.serialNumber, organization, commonName, err)
	}

	model := r.Organization()
	parentCertificate := r.Certificate()

	serialNumber, err := apputils.GenerateSerialNumber(r.randomManager)
	if err != nil {
		return nil, nil, fmt.Errorf("[%s@%s:NewIntermediateCertificate:%s]: failed to create serial number: %w", r.serialNumber, organization, commonName, err)
	}

	newPrivateKey, err := apputils.GeneratePrivateKey(
		organization,
		append(parentCertificate.Parents(), parentCertificate.SerialNumber(), serialNumber),
		appmodels.ECDSA_P384,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("[%s@%s:NewIntermediateCertificate:%s]: failed to create private key: %w", r.serialNumber, organization, commonName, err)
	}

	cert, err := apputils.NewIntermediateCertificate(
		r.certManager,
		serialNumber,
		model,
		r.expiration,
		newPrivateKey,
		parentCertificate,
		parentPrivateKey,
		commonName,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("[%s@%s:NewIntermediateCertificate:%s]: failed: %w", r.serialNumber, organization, commonName, err)
	}
	log.Printf("[%s@%s:NewIntermediateCertificate:%s]: Certificate generated", r.serialNumber, organization, commonName)

	// savedPrivateKey, err := r.privateKeyRepository.Save(newPrivateKey)
	// if err != nil {
	// 	return nil, nil, fmt.Errorf("[%s@%s:NewIntermediateCertificate:%s]: could not save private key: %w", r.serialNumber, organization, commonName, err)
	// }
	// log.Printf("[%s@%s:NewIntermediateCertificate:%s]: Private key saved", r.serialNumber, organization, commonName)

	savedModel, err := r.certificateRepository.Save(cert)
	if err != nil {
		return nil, nil, fmt.Errorf("[%s@%s:NewIntermediateCertificate:%s]: could not save certificate: %w", r.serialNumber, organization, commonName, err)
	}
	log.Printf("[%s@%s:NewIntermediateCertificate:%s]: Certificate saved", r.serialNumber, organization, commonName)

	return savedModel, newPrivateKey, nil
}

func (r *CertCertificateController) NewServerCertificate(dnsNames ...string) (appmodels.Certificate, appmodels.PrivateKey, error) {

	organization := r.OrganizationID()

	if len(dnsNames) <= 0 {
		return nil, nil, fmt.Errorf("[%s@%s:NewServerCertificate:%s]: server certificate must have at least one dns name", r.serialNumber, organization, strings.Join(dnsNames, ","))
	}

	model := r.Organization()

	parentCertificate := r.Certificate()

	parentPrivateKey, err := r.PrivateKey()
	if err != nil {
		return nil, nil, fmt.Errorf("[%s@%s:NewServerCertificate:%s]: failed to fetch private key: %w", r.serialNumber, organization, strings.Join(dnsNames, ","), err)
	}

	serialNumber, err := apputils.GenerateSerialNumber(r.randomManager)
	if err != nil {
		return nil, nil, fmt.Errorf("[%s@%s:NewServerCertificate:%s]: failed to create serial number: %w", r.serialNumber, organization, strings.Join(dnsNames, ","), err)
	}

	newPrivateKey, err := apputils.GeneratePrivateKey(
		model.ID(),
		append(parentCertificate.Parents(), parentCertificate.SerialNumber(), serialNumber),
		appmodels.ECDSA_P384,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("[%s@%s:NewServerCertificate:%s]: failed to create private key: %w", r.serialNumber, organization, strings.Join(dnsNames, ","), err)
	}

	cert, err := apputils.NewServerCertificate(
		r.certManager,
		serialNumber,
		model,
		r.expiration,
		newPrivateKey,
		parentCertificate,
		parentPrivateKey,
		dnsNames[0],
		dnsNames...,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("[%s@%s:NewServerCertificate:%s]: failed to create intermediate certificate: %w", r.serialNumber, organization, strings.Join(dnsNames, ","), err)
	}
	log.Printf("[%s@%s:NewServerCertificate:%s]: Certificate generated", r.serialNumber, organization, strings.Join(dnsNames, ","))

	// savedPrivateKey, err := r.privateKeyRepository.Save(newPrivateKey)
	// if err != nil {
	// 	return nil, nil, fmt.Errorf("[%s@%s:NewServerCertificate:%s]: could not save private key: %w", r.serialNumber, organization, strings.Join(dnsNames, ","), err)
	// }
	// log.Printf("[%s@%s:NewServerCertificate:%s]: Private key saved", r.serialNumber, organization, strings.Join(dnsNames, ","))

	savedModel, err := r.certificateRepository.Save(cert)
	if err != nil {
		return nil, nil, fmt.Errorf("[%s@%s:NewServerCertificate:%s]: could not save certificate: %w", r.serialNumber, organization, strings.Join(dnsNames, ","), err)
	}
	log.Printf("[%s@%s:NewServerCertificate:%s]: Certificate saved", r.serialNumber, organization, strings.Join(dnsNames, ","))

	return savedModel, newPrivateKey, nil
}

func (r *CertCertificateController) NewClientCertificate(commonName string) (appmodels.Certificate, appmodels.PrivateKey, error) {

	organization := r.OrganizationID()

	parentPrivateKey, err := r.PrivateKey()
	if err != nil {
		return nil, nil, fmt.Errorf("[%s@%s:NewClientCertificate:%s]: failed to fetch private key: %w", r.serialNumber, organization, commonName, err)
	}
	log.Printf("[%s@%s:NewClientCertificate:%s]: parentPrivateKey accquired", r.serialNumber, organization, commonName)

	model := r.Organization()
	log.Printf("[%s@%s:NewClientCertificate:%s]: model = %s", r.serialNumber, organization, commonName, model)

	parentCertificate := r.Certificate()

	serialNumber, err := apputils.GenerateSerialNumber(r.randomManager)
	if err != nil {
		return nil, nil, fmt.Errorf("[%s@%s:NewClientCertificate:%s]: failed to create serial number: %w", r.serialNumber, organization, commonName, err)
	}
	log.Printf("[%s@%s:NewClientCertificate:%s]: serialNumber = %s", r.serialNumber, organization, commonName, serialNumber.String())

	newPrivateKey, err := apputils.GeneratePrivateKey(
		organization,
		append(parentCertificate.Parents(), parentCertificate.SerialNumber(), serialNumber),
		appmodels.ECDSA_P384,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("[%s@%s:NewClientCertificate:%s]: failed to create private key: %w", r.serialNumber, organization, commonName, err)
	}
	log.Printf("[%s@%s:NewClientCertificate:%s]: Private key generated", r.serialNumber, organization, commonName)

	cert, err := apputils.NewClientCertificate(
		r.certManager,
		serialNumber,
		model,
		r.expiration,
		newPrivateKey,
		parentCertificate,
		parentPrivateKey,
		commonName,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("[%s@%s:NewClientCertificate:%s]: failed to create intermediate certificate: %w", r.serialNumber, organization, commonName, err)
	}
	log.Printf("[%s@%s:NewClientCertificate:%s]: Certificate generated", r.serialNumber, organization, commonName)

	// savedPrivateKey, err := r.privateKeyRepository.Save(newPrivateKey)
	// if err != nil {
	// 	return nil, nil, fmt.Errorf("[%s@%s:NewClientCertificate:%s]: could not save private key: %w", r.serialNumber, organization, commonName, err)
	// }
	// log.Printf("[%s@%s:NewClientCertificate:%s]: Private key saved", r.serialNumber, organization, commonName)

	savedModel, err := r.certificateRepository.Save(cert)
	if err != nil {
		return nil, nil, fmt.Errorf("[%s@%s:NewClientCertificate:%s]: could not save certificate: %w", r.serialNumber, organization, commonName, err)
	}
	log.Printf("[%s@%s:NewClientCertificate:%s]: Certificate saved", r.serialNumber, organization, commonName)

	return savedModel, newPrivateKey, nil
}

func (r *CertCertificateController) UsesCertificateService(service appmodels.CertificateRepository) bool {
	return r.certificateRepository == service
}

func (r *CertCertificateController) FindCertificate(organization string, certificates []*big.Int) (appmodels.Certificate, error) {
	return r.certificateRepository.FindByOrganizationAndSerialNumbers(organization, certificates)
}

// NewCertificateController creates a new instance of CertCertificateController
//
// injecting the specified CertificateRepository implementation. This setup
// facilitates the separation of business logic from data access layers,
// aligning with the principles of dependency injection.
//
//   - parentOrganizationController appmodels.OrganizationController
//   - parentCertificateController appmodels.CertificateController
//   - serialNumber *big.Int
//   - model appmodels.Certificate
//   - certificateRepository is appmodels.CertificateRepository
//   - privateKeyRepository is appmodels.PrivateKeyRepository
//   - certManager is managers.CertificateManager
//   - randomManager is  managers.RandomManager
//   - expiration time.Duration is
//
// Returns *CertCertificateController
func NewCertificateController(
	parentOrganizationController appmodels.OrganizationController,
	parentCertificateController appmodels.CertificateController,
	serialNumber *big.Int,
	model appmodels.Certificate,
	certificateRepository appmodels.CertificateRepository,
	privateKeyRepository appmodels.PrivateKeyRepository,
	certManager managers.CertificateManager,
	randomManager managers.RandomManager,
	expiration time.Duration,
) *CertCertificateController {

	if parentOrganizationController == nil {
		panic("NewCertificateController: parentOrganizationController not defined")
	}

	if certificateRepository == nil {
		panic("NewCertificateController: certificateRepository not defined")
	}

	if privateKeyRepository == nil {
		panic("NewCertificateController: privateKeyRepository not defined")
	}

	if certManager == nil {
		panic("NewCertificateController: certManager not defined")
	}

	if randomManager == nil {
		panic("NewCertificateController: randomManager not defined")
	}

	return &CertCertificateController{
		serialNumber:                 serialNumber,
		model:                        model,
		certificateRepository:        certificateRepository,
		privateKeyRepository:         privateKeyRepository,
		expiration:                   expiration,
		certManager:                  certManager,
		randomManager:                randomManager,
		parentOrganizationController: parentOrganizationController,
		parentCertificateController:  parentCertificateController,
	}
}

var _ appmodels.CertificateController = (*CertCertificateController)(nil)
