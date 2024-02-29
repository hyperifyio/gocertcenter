// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appcontrollers

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hyperifyio/gocertcenter/internal/common/managers"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
)

// CertificateController implements appmodels.ICertificateController to control
// operations for certificate models.
//
// It utilizes appmodels.ICertificateService implementation to abstract
// and inject the underlying storage mechanism (e.g., database, memory). This
// design promotes separation of concerns by decoupling the business logic from
// the specific details of data persistence.
type CertificateController struct {
	serialNumber appmodels.ISerialNumber
	model        appmodels.ICertificate

	// parentOrganizationController is the parent organization controller
	parentOrganizationController appmodels.IOrganizationController

	// parentCertificateController  is an optional parent certificate controller
	parentCertificateController appmodels.ICertificateController

	privateKeyController appmodels.IPrivateKeyController

	certManager   managers.ICertificateManager
	randomManager managers.IRandomManager

	certificateRepository appmodels.ICertificateService
	privateKeyRepository  appmodels.IPrivateKeyService

	expiration time.Duration
}

func (r *CertificateController) GetOrganizationController() appmodels.IOrganizationController {
	if r.parentOrganizationController == nil {
		panic("[CertificateController:GetOrganizationController]: No parent organization controller")
	}
	return r.parentOrganizationController
}

func (r *CertificateController) GetApplicationController() appmodels.IApplicationController {
	if r.parentOrganizationController == nil {
		panic("[CertificateController.GetApplicationController]: No parent organization controller")
	}
	return r.parentOrganizationController.GetApplicationController()
}

func (r *CertificateController) GetOrganizationID() string {
	if r.parentOrganizationController == nil {
		panic("[CertificateController.GetOrganizationID]: No parent organization controller")
	}
	return r.parentOrganizationController.GetOrganizationID()
}

func (r *CertificateController) GetOrganizationModel() appmodels.IOrganization {
	return r.parentOrganizationController.GetOrganizationModel()
}

func (r *CertificateController) GetCertificateModel() appmodels.ICertificate {
	return r.model
}

func (r *CertificateController) GetChildCertificateCollection(certificateType string) ([]appmodels.ICertificate, error) {
	organization := r.GetOrganizationID()
	path := append(r.model.GetParents(), r.serialNumber)
	list, err := r.certificateRepository.FindAllByOrganizationAndSerialNumbers(
		organization,
		path,
	)

	if certificateType != "" {
		list = apputils.FilterCertificatesByType(list, certificateType)
	}

	if err != nil {
		return nil, fmt.Errorf("[%s@%s:GetChildCertificateCollection]: failed: %w", r.serialNumber.String(), organization, err)
	}
	return list, nil
}

func (r *CertificateController) GetChildCertificateModel(serialNumber appmodels.ISerialNumber) (appmodels.ICertificate, error) {
	organization := r.GetOrganizationID()
	if r.certificateRepository == nil {
		return nil, fmt.Errorf("[%s@%s:GetChildCertificateModel:%s]: No parent certificateRepository", r.serialNumber.String(), organization, serialNumber.String())
	}
	if r.model == nil {
		return nil, fmt.Errorf("[%s@%s:GetChildCertificateModel:%s]: No parent model", r.serialNumber.String(), organization, serialNumber.String())
	}
	model, err := r.certificateRepository.FindByOrganizationAndSerialNumbers(
		organization,
		append(r.model.GetParents(), r.serialNumber, serialNumber),
	)
	if err != nil {
		return nil, fmt.Errorf("[%s@%s:GetChildCertificateModel:%s]: failed: %w", r.serialNumber.String(), organization, serialNumber.String(), err)
	}
	return model, nil
}

func (r *CertificateController) GetChildCertificateController(serialNumber appmodels.ISerialNumber) (appmodels.ICertificateController, error) {
	organization := r.GetOrganizationID()
	model, err := r.GetChildCertificateModel(serialNumber)
	if err != nil {
		return nil, fmt.Errorf("[%s@%s:GetChildCertificateController:%s]: could not find: %w", r.serialNumber.String(), organization, serialNumber.String(), err)
	}
	return NewCertificateController(
		r.GetOrganizationController(),
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

func (r *CertificateController) GetParentCertificateModel() appmodels.ICertificate {
	if r.parentCertificateController == nil {
		return nil
	}
	return r.parentCertificateController.GetCertificateModel()
}

func (r *CertificateController) GetParentCertificateController() appmodels.ICertificateController {
	if r.parentCertificateController == nil {
		return nil
	}
	return r.parentCertificateController
}

func (r *CertificateController) GetPrivateKeyModel() (appmodels.IPrivateKey, error) {
	organization := r.GetOrganizationID()
	if r.privateKeyRepository == nil {
		return nil, fmt.Errorf("[%s@%s:GetPrivateKeyModel]: no private key repository", r.serialNumber, organization)
	}
	model, err := r.privateKeyRepository.FindByOrganizationAndSerialNumbers(
		organization,
		append(r.model.GetParents(), r.serialNumber),
	)
	if err != nil {
		return nil, fmt.Errorf("[%s@%s:GetPrivateKeyModel]: failed: %w", r.serialNumber, organization, err)
	}
	return model, nil
}

func (r *CertificateController) GetPrivateKeyController() (appmodels.IPrivateKeyController, error) {
	organization := r.GetOrganizationID()
	model, err := r.GetPrivateKeyModel()
	if err != nil {
		return nil, fmt.Errorf("[%s@%s:GetPrivateKeyController]: failed: %w", r.serialNumber, organization, err)
	}
	return NewPrivateKeyController(
		model,
		r,
		r.privateKeyRepository,
	), nil
}

func (r *CertificateController) SetExpirationDuration(expiration time.Duration) {
	r.expiration = expiration
}

func (r *CertificateController) NewIntermediateCertificate(commonName string) (appmodels.ICertificate, appmodels.IPrivateKey, error) {

	organization := r.GetOrganizationID()

	parentPrivateKey, err := r.GetPrivateKeyModel()
	if err != nil {
		return nil, nil, fmt.Errorf("[%s@%s:NewIntermediateCertificate:%s]: failed to get private key: %w", r.serialNumber, organization, commonName, err)
	}

	model := r.GetOrganizationModel()
	parentCertificate := r.GetCertificateModel()

	serialNumber, err := apputils.GenerateSerialNumber(r.randomManager)
	if err != nil {
		return nil, nil, fmt.Errorf("[%s@%s:NewIntermediateCertificate:%s]: failed to create serial number: %w", r.serialNumber, organization, commonName, err)
	}

	newPrivateKey, err := apputils.GeneratePrivateKey(
		organization,
		append(parentCertificate.GetParents(), parentCertificate.GetSerialNumber(), serialNumber),
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

func (r *CertificateController) NewServerCertificate(dnsNames ...string) (appmodels.ICertificate, appmodels.IPrivateKey, error) {

	organization := r.GetOrganizationID()

	if len(dnsNames) <= 0 {
		return nil, nil, fmt.Errorf("[%s@%s:NewServerCertificate:%s]: server certificate must have at least one dns name", r.serialNumber, organization, strings.Join(dnsNames, ","))
	}

	model := r.GetOrganizationModel()

	parentCertificate := r.GetCertificateModel()

	parentPrivateKey, err := r.GetPrivateKeyModel()
	if err != nil {
		return nil, nil, fmt.Errorf("[%s@%s:NewServerCertificate:%s]: failed to fetch private key: %w", r.serialNumber, organization, strings.Join(dnsNames, ","), err)
	}

	serialNumber, err := apputils.GenerateSerialNumber(r.randomManager)
	if err != nil {
		return nil, nil, fmt.Errorf("[%s@%s:NewServerCertificate:%s]: failed to create serial number: %w", r.serialNumber, organization, strings.Join(dnsNames, ","), err)
	}

	newPrivateKey, err := apputils.GeneratePrivateKey(
		model.GetID(),
		append(parentCertificate.GetParents(), parentCertificate.GetSerialNumber(), serialNumber),
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

func (r *CertificateController) NewClientCertificate(commonName string) (appmodels.ICertificate, appmodels.IPrivateKey, error) {

	organization := r.GetOrganizationID()

	parentPrivateKey, err := r.GetPrivateKeyModel()
	if err != nil {
		return nil, nil, fmt.Errorf("[%s@%s:NewClientCertificate:%s]: failed to fetch private key: %w", r.serialNumber, organization, commonName, err)
	}
	log.Printf("[%s@%s:NewClientCertificate:%s]: parentPrivateKey accquired", r.serialNumber, organization, commonName)

	model := r.GetOrganizationModel()
	log.Printf("[%s@%s:NewClientCertificate:%s]: model = %s", r.serialNumber, organization, commonName, model)

	parentCertificate := r.GetCertificateModel()

	serialNumber, err := apputils.GenerateSerialNumber(r.randomManager)
	if err != nil {
		return nil, nil, fmt.Errorf("[%s@%s:NewClientCertificate:%s]: failed to create serial number: %w", r.serialNumber, organization, commonName, err)
	}
	log.Printf("[%s@%s:NewClientCertificate:%s]: serialNumber = %s", r.serialNumber, organization, commonName, serialNumber.String())

	newPrivateKey, err := apputils.GeneratePrivateKey(
		organization,
		append(parentCertificate.GetParents(), parentCertificate.GetSerialNumber(), serialNumber),
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

func (r *CertificateController) UsesCertificateService(service appmodels.ICertificateService) bool {
	return r.certificateRepository == service
}

func (r *CertificateController) GetExistingCertificate(organization string, certificates []appmodels.ISerialNumber) (appmodels.ICertificate, error) {
	return r.certificateRepository.FindByOrganizationAndSerialNumbers(organization, certificates)
}

// NewCertificateController creates a new instance of CertificateController
//
// injecting the specified ICertificateService implementation. This setup
// facilitates the separation of business logic from data access layers,
// aligning with the principles of dependency injection.
//
//   - parentOrganizationController appmodels.IOrganizationController
//   - parentCertificateController appmodels.ICertificateController
//   - serialNumber appmodels.ISerialNumber
//   - model appmodels.ICertificate
//   - certificateRepository is appmodels.ICertificateService
//   - privateKeyRepository is appmodels.IPrivateKeyService
//   - certManager is managers.ICertificateManager
//   - randomManager is  managers.IRandomManager
//   - expiration time.Duration is
//
// Returns *CertificateController
func NewCertificateController(
	parentOrganizationController appmodels.IOrganizationController,
	parentCertificateController appmodels.ICertificateController,
	serialNumber appmodels.ISerialNumber,
	model appmodels.ICertificate,
	certificateRepository appmodels.ICertificateService,
	privateKeyRepository appmodels.IPrivateKeyService,
	certManager managers.ICertificateManager,
	randomManager managers.IRandomManager,
	expiration time.Duration,
) *CertificateController {

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

	return &CertificateController{
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

var _ appmodels.ICertificateController = (*CertificateController)(nil)
