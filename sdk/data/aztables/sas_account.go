// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package aztables

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"time"
)

var SASVersion = "2019-02-02"

// AccountSASSignatureValues is used to generate a Shared Access Signature (SAS) for an Azure Storage account.
// For more information, see https://docs.microsoft.com/rest/api/storageservices/constructing-an-account-sas
type AccountSASSignatureValues struct {
	Version       string      `param:"sv"`  // If not specified, this defaults to SASVersion
	Protocol      SASProtocol `param:"spr"` // See the SASProtocol* constants
	StartTime     time.Time   `param:"st"`  // Not specified if IsZero
	ExpiryTime    time.Time   `param:"se"`  // Not specified if IsZero
	Permissions   string      `param:"sp"`  // Create by initializing a AccountSASPermissions and then call String()
	IPRange       IPRange     `param:"sip"`
	Services      string      `param:"ss"`  // Create by initializing AccountSASServices and then call String()
	ResourceTypes string      `param:"srt"` // Create by initializing AccountSASResourceTypes and then call String()
}

// Sign uses an account's SharedKeyCredential to sign this signature values to produce the proper SAS query parameters.
func (v AccountSASSignatureValues) Sign(sharedKeyCredential *SharedKeyCredential) (SASQueryParameters, error) {
	// https://docs.microsoft.com/en-us/rest/api/storageservices/Constructing-an-Account-SAS
	if v.ExpiryTime.IsZero() || v.Permissions == "" || v.ResourceTypes == "" || v.Services == "" {
		return SASQueryParameters{}, errors.New("account SAS is missing at least one of these: ExpiryTime, Permissions, Service, or ResourceType")
	}
	if v.Version == "" {
		v.Version = SASVersion
	}
	perms := &AccountSASPermissions{}
	if err := perms.Parse(v.Permissions); err != nil {
		return SASQueryParameters{}, err
	}
	v.Permissions = perms.String()

	startTime, expiryTime := FormatTimesForSASSigning(v.StartTime, v.ExpiryTime)

	stringToSign := strings.Join([]string{
		sharedKeyCredential.AccountName(),
		v.Permissions,
		v.Services,
		v.ResourceTypes,
		startTime,
		expiryTime,
		v.IPRange.String(),
		string(v.Protocol),
		v.Version,
		""}, // That right, the account SAS requires a terminating extra newline
		"\n")

	signature, err := sharedKeyCredential.computeHMACSHA256(stringToSign)
	if err != nil {
		return SASQueryParameters{}, err
	}
	p := SASQueryParameters{
		// Common SAS parameters
		version:     v.Version,
		protocol:    v.Protocol,
		startTime:   v.StartTime,
		expiryTime:  v.ExpiryTime,
		permissions: v.Permissions,
		ipRange:     v.IPRange,

		// Account-specific SAS parameters
		services:      v.Services,
		resourceTypes: v.ResourceTypes,

		// Calculated SAS signature
		signature: signature,
	}

	return p, nil
}

// The AccountSASPermissions type simplifies creating the permissions string for an Azure Storage Account SAS.
// Initialize an instance of this type and then call its String method to set AccountSASSignatureValues's Permissions field.
type AccountSASPermissions struct {
	Read    bool
	Write   bool
	Delete  bool
	List    bool
	Add     bool
	Create  bool
	Update  bool
	Process bool
}

// String produces the SAS permissions string for an Azure Storage account.
// Call this method to set AccountSASSignatureValues's Permissions field.
func (p AccountSASPermissions) String() string {
	var buffer bytes.Buffer
	if p.Read {
		buffer.WriteRune('r')
	}
	if p.Write {
		buffer.WriteRune('w')
	}
	if p.Delete {
		buffer.WriteRune('d')
	}
	if p.List {
		buffer.WriteRune('l')
	}
	if p.Add {
		buffer.WriteRune('a')
	}
	if p.Create {
		buffer.WriteRune('c')
	}
	if p.Update {
		buffer.WriteRune('u')
	}
	if p.Process {
		buffer.WriteRune('p')
	}
	return buffer.String()
}

// Parse initializes the AccountSASPermissions's fields from a string.
func (p *AccountSASPermissions) Parse(s string) error {
	*p = AccountSASPermissions{} // Clear out the flags
	for _, r := range s {
		switch r {
		case 'r':
			p.Read = true
		case 'w':
			p.Write = true
		case 'd':
			p.Delete = true
		case 'l':
			p.List = true
		case 'a':
			p.Add = true
		case 'c':
			p.Create = true
		case 'u':
			p.Update = true
		case 'p':
			p.Process = true
		case 'x':
			p.Process = true
		default:
			return fmt.Errorf("invalid permission character: '%v'", r)
		}
	}
	return nil
}

// The AccountSASResourceTypes type simplifies creating the resource types string for an Azure Storage Account SAS.
// Initialize an instance of this type and then call its String method to set AccountSASSignatureValues's ResourceTypes field.
type AccountSASResourceTypes struct {
	Service   bool
	Container bool
	Object    bool
}

// String produces the SAS resource types string for an Azure Storage account.
// Call this method to set AccountSASSignatureValues's ResourceTypes field.
func (rt AccountSASResourceTypes) String() string {
	var buffer bytes.Buffer
	if rt.Service {
		buffer.WriteRune('s')
	}
	if rt.Container {
		buffer.WriteRune('c')
	}
	if rt.Object {
		buffer.WriteRune('o')
	}
	return buffer.String()
}

// Parse initializes the AccountSASResourceType's fields from a string.
func (rt *AccountSASResourceTypes) Parse(s string) error {
	*rt = AccountSASResourceTypes{} // Clear out the flags
	for _, r := range s {
		switch r {
		case 's':
			rt.Service = true
		case 'c':
			rt.Container = true
		case 'o':
			rt.Object = true
		default:
			return fmt.Errorf("invalid resource type: '%v'", r)
		}
	}
	return nil
}
