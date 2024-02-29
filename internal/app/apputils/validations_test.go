// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apputils_test

import (
	"testing"

	"github.com/hyperifyio/gocertcenter/internal/app/appmocks"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
)

func TestValidateRootCertificateCommonName(t *testing.T) {
	tests := []struct {
		name        string
		commonName  string
		wantErr     bool
		expectedErr string
	}{
		{"Valid Name", "Valid Common Name", false, ""},
		{"Empty", "", true, "cannot be empty"},
		{"Starts with Space", " CommonName", true, "cannot start with a space"},
		{"Ends with Space", "CommonName ", true, "cannot end with a space"},
		{"Repeating Spaces", "Common  Name", true, "should not have repeating spaces"},
		{"Full Numbers", "123456", true, "should not be full numbers"},
		{"Invalid Characters", "Common*Name", true, "contains invalid characters"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := apputils.ValidateRootCertificateCommonName(tt.commonName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateRootCertificateCommonName() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != nil && err.Error() != tt.expectedErr {
				t.Errorf("ValidateRootCertificateCommonName() error = %v, expectedErr %v", err, tt.expectedErr)
			}
		})
	}
}

func TestValidateOrganizationName(t *testing.T) {
	testCases := []struct {
		name          string
		expectedError string
	}{
		{"", "must not be empty and must be at least two characters long"},
		{"a", "must not be empty and must be at least two characters long"},
		{" abcdefgh", "must not have leading or trailing spaces"},
		{"abcdefgh ", "must not have leading or trailing spaces"},
		{"validName", ""},
		{"valid Name", ""},
		{"valid  Name", "should not have repeating spaces"},
		{"123", "should not be full numbers"},
		{"valid-name", ""},
		{"valid_name", ""},
		{"valid.name", ""},
		{"invalid#name", "contains invalid characters"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := apputils.ValidateOrganizationName(tc.name)
			if tc.expectedError == "" && err != nil {
				t.Errorf("Expected no error, but got %s", err)
			} else if tc.expectedError != "" && (err == nil || err.Error() != tc.expectedError) {
				t.Errorf("Expected error '%s', but got '%v'", tc.expectedError, err)
			}
		})
	}
}

func TestValidateOrganizationNames(t *testing.T) {
	testCases := []struct {
		name          string
		names         []string
		expectedError string
	}{
		{
			name:  "All Valid Names",
			names: []string{"ValidName1", "Valid Name2", "valid-name3"},
		},
		{
			name:  "Empty Slice",
			names: []string{},
		},
		{
			name:          "Contains Invalid Name",
			names:         []string{"ValidName", "Invalid#Name", "AnotherValidName"},
			expectedError: "contains invalid characters",
		},
		{
			name:          "Contains Empty Name",
			names:         []string{"ValidName", "", "AnotherValidName"},
			expectedError: "must not be empty and must be at least two characters long",
		},
		{
			name:          "Leading Spaces",
			names:         []string{" ValidName"},
			expectedError: "must not have leading or trailing spaces",
		},
		{
			name:          "Trailing Spaces",
			names:         []string{"ValidName "},
			expectedError: "must not have leading or trailing spaces",
		},
		{
			name:          "Repeating Spaces",
			names:         []string{"Valid  Name"},
			expectedError: "should not have repeating spaces",
		},
		{
			name:          "Full Numbers",
			names:         []string{"12345"},
			expectedError: "should not be full numbers",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := apputils.ValidateOrganizationNames(tc.names)
			if tc.expectedError == "" && err != nil {
				t.Errorf("Expected no error, but got %s", err)
			} else if tc.expectedError != "" && (err == nil || err.Error() != tc.expectedError) {
				t.Errorf("Expected error '%s', but got '%v'", tc.expectedError, err)
			}
		})
	}
}

func TestValidateOrganizationID(t *testing.T) {
	testCases := []struct {
		name          string
		id            string
		expectedError string
	}{
		{
			name: "Valid ID",
			id:   "valid-id123",
		},
		{
			name: "Valid ID",
			id:   "abc",
		},
		{
			name: "Valid ID",
			id:   "abcd",
		},
		{
			name: "Valid ID",
			id:   "ab",
		},
		{
			name: "Valid ID",
			id:   "helloworld",
		},
		{
			name: "Valid ID",
			id:   "neworg",
		},
		{
			name:          "Valid ID",
			id:            "newOrg",
			expectedError: "contains uppercase characters",
		},
		{
			name:          "Too Short",
			id:            "v",
			expectedError: "must be at least two characters long",
		},
		{
			name:          "Full Numbers",
			id:            "123456",
			expectedError: "should not be full numbers",
		},
		{
			name:          "Invalid Characters",
			id:            "invalid_id!",
			expectedError: "contains invalid characters, or has leading/trailing spaces, '-', or '.'",
		},
		{
			name:          "Leading Spaces",
			id:            " leading-id",
			expectedError: "contains invalid characters, or has leading/trailing spaces, '-', or '.'",
		},
		{
			name:          "Trailing Spaces",
			id:            "trailing-id ",
			expectedError: "contains invalid characters, or has leading/trailing spaces, '-', or '.'",
		},
		{
			name:          "Leading Dash",
			id:            "-leadingdash",
			expectedError: "contains invalid characters, or has leading/trailing spaces, '-', or '.'",
		},
		{
			name:          "Trailing Dash",
			id:            "trailingdash-",
			expectedError: "contains invalid characters, or has leading/trailing spaces, '-', or '.'",
		},
		{
			name:          "Leading Period",
			id:            ".leadingperiod",
			expectedError: "contains invalid characters, or has leading/trailing spaces, '-', or '.'",
		},
		{
			name:          "Trailing Period",
			id:            "trailingperiod.",
			expectedError: "contains invalid characters, or has leading/trailing spaces, '-', or '.'",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := apputils.ValidateOrganizationID(tc.id)
			if tc.expectedError == "" && err != nil {
				t.Errorf("Expected no error for '%s', but got: %s", tc.id, err)
			} else if tc.expectedError != "" && (err == nil || err.Error() != tc.expectedError) {
				t.Errorf("Expected error '%s', but got '%v'", tc.expectedError, err)
			}
		})
	}
}

func TestValidateOrganizationModel(t *testing.T) {
	testCases := []struct {
		name          string
		setupMocks    func(*appmocks.MockOrganization)
		expectedError string
	}{
		{
			name: "Valid Organization",
			setupMocks: func(m *appmocks.MockOrganization) {
				m.On("GetID").Return("valid-id")
				m.On("GetName").Return("Valid Organization")
				m.On("GetNames").Return([]string{"Valid Org", "Another Valid Name"})
			},
		},
		{
			name: "Invalid Organization ID",
			setupMocks: func(m *appmocks.MockOrganization) {
				m.On("GetID").Return("!!!!!!")
				m.On("GetName").Return("Valid Organization")
				m.On("GetNames").Return([]string{"Valid Org", "Another Valid Name"})
			},
			expectedError: "id: '!!!!!!': contains invalid characters, or has leading/trailing spaces, '-', or '.'",
		},
		{
			name: "Invalid Organization Name",
			setupMocks: func(m *appmocks.MockOrganization) {
				m.On("GetID").Return("valid-id")
				m.On("GetName").Return("!!!!!!!")
				m.On("GetNames").Return([]string{"Valid Org", "Another Valid Name"})
			},
			expectedError: "name: '!!!!!!!': contains invalid characters",
		},
		{
			name: "Invalid Organization Names",
			setupMocks: func(m *appmocks.MockOrganization) {
				m.On("GetID").Return("valid-id")
				m.On("GetName").Return("Valid Organization")
				m.On("GetNames").Return([]string{"Valid Org", "!!!!!!"})
			},
			expectedError: "names: '[Valid Org !!!!!!]': contains invalid characters",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockOrg := new(appmocks.MockOrganization)
			tc.setupMocks(mockOrg)

			err := apputils.ValidateOrganizationModel(mockOrg)
			if tc.expectedError == "" && err != nil {
				t.Errorf("Expected no error, but got %s", err)
			} else if tc.expectedError != "" && (err == nil || err.Error() != tc.expectedError) {
				t.Errorf("Expected error '%s', but got '%v'", tc.expectedError, err)
			}
		})
	}
}

func TestValidateClientCertificateCommonName(t *testing.T) {
	testCases := []struct {
		name          string
		commonName    string
		expectedError string
	}{
		{
			name:          "Valid Common Name",
			commonName:    "john.doe@example.com",
			expectedError: "",
		},
		{
			name:          "Empty Common Name",
			commonName:    "",
			expectedError: "cannot be empty",
		},
		{
			name:          "Starts With Space",
			commonName:    " john.doe@example.com",
			expectedError: "cannot start with a space",
		},
		{
			name:          "Ends With Space",
			commonName:    "john.doe@example.com ",
			expectedError: "cannot end with a space",
		},
		{
			name:          "Repeating Spaces",
			commonName:    "john  doe@example.com",
			expectedError: "should not have repeating spaces",
		},
		{
			name:          "Invalid Characters",
			commonName:    "john*doe@example.com",
			expectedError: "contains invalid characters",
		},
		{
			name:          "Valid With Hyphen",
			commonName:    "john-doe@example.com",
			expectedError: "",
		},
		{
			name:          "Valid With Underscore",
			commonName:    "john_doe@example.com",
			expectedError: "",
		},
		{
			name:          "Valid With Dot",
			commonName:    "john.doe@example.com",
			expectedError: "",
		},
		{
			name:          "Valid With @",
			commonName:    "john.doe@example.com",
			expectedError: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := apputils.ValidateClientCertificateCommonName(tc.commonName)
			if tc.expectedError == "" && err != nil {
				t.Errorf("Expected no error, but got %s", err)
			} else if tc.expectedError != "" && (err == nil || err.Error() != tc.expectedError) {
				t.Errorf("Expected error '%s', but got '%v'", tc.expectedError, err)
			}
		})
	}
}

func TestValidateServerCertificateCommonName(t *testing.T) {
	testCases := []struct {
		name          string
		commonName    string
		expectedError string
	}{
		{
			name:          "Valid Common Name",
			commonName:    "www.example.com",
			expectedError: "",
		},
		{
			name:          "Empty Common Name",
			commonName:    "",
			expectedError: "cannot be empty",
		},
		{
			name:          "Starts With Hyphen",
			commonName:    "-example.com",
			expectedError: "cannot start or end with '-' or '.'",
		},
		{
			name:          "Ends With Hyphen",
			commonName:    "example.com-",
			expectedError: "cannot start or end with '-' or '.'",
		},
		{
			name:          "Starts With Period",
			commonName:    ".example.com",
			expectedError: "cannot start or end with '-' or '.'",
		},
		{
			name:          "Ends With Period",
			commonName:    "example.com.",
			expectedError: "cannot start or end with '-' or '.'",
		},
		{
			name:          "Repeating Periods",
			commonName:    "example..com",
			expectedError: "cannot have repeating '.'",
		},
		{
			name:          "Invalid Characters",
			commonName:    "example*com",
			expectedError: "contains invalid characters",
		},
		{
			name:          "Valid With Wildcard",
			commonName:    "*.example.com",
			expectedError: "",
		},
		{
			name:          "Short Last TLD",
			commonName:    "example.c",
			expectedError: "last TLD must be at least two characters",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := apputils.ValidateServerCertificateCommonName(tc.commonName)
			if tc.expectedError == "" && err != nil {
				t.Errorf("Expected no error for '%s', but got: %s", tc.commonName, err)
			} else if tc.expectedError != "" && (err == nil || err.Error() != tc.expectedError) {
				t.Errorf("Expected error '%s' for '%s', but got: %v", tc.expectedError, tc.commonName, err)
			}
		})
	}
}

func TestValidateDNSNames(t *testing.T) {
	testCases := []struct {
		name          string
		dnsNames      []string
		expectedError string
	}{
		{
			name:          "All Valid DNS Names",
			dnsNames:      []string{"www.example.com", "*.example.com", "mail.example.com"},
			expectedError: "",
		},
		{
			name:          "Contains Empty DNS Name",
			dnsNames:      []string{"www.example.com", "", "mail.example.com"},
			expectedError: "cannot be empty",
		},
		{
			name:          "Contains DNS Name Starting With Hyphen",
			dnsNames:      []string{"www.example.com", "-example.com"},
			expectedError: "cannot start or end with '-' or '.'",
		},
		{
			name:          "Contains DNS Name With Repeating Periods",
			dnsNames:      []string{"example..com"},
			expectedError: "cannot have repeating '.'",
		},
		{
			name:          "Contains DNS Name With Invalid Characters",
			dnsNames:      []string{"www.example.com", "example*com"},
			expectedError: "contains invalid characters",
		},
		{
			name:          "Contains DNS Name With Short Last TLD",
			dnsNames:      []string{"www.example.com", "example.c"},
			expectedError: "last TLD must be at least two characters",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := apputils.ValidateDNSNames(tc.dnsNames)
			if tc.expectedError == "" && err != nil {
				t.Errorf("Expected no error, but got: %s", err)
			} else if tc.expectedError != "" && (err == nil || err.Error() != tc.expectedError) {
				t.Errorf("Expected error '%s', but got: %v", tc.expectedError, err)
			}
		})
	}
}
