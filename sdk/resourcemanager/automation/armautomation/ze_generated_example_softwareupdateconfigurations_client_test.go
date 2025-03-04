//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/softwareUpdateConfiguration/createSoftwareUpdateConfiguration.json
func ExampleSoftwareUpdateConfigurationsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomation.NewSoftwareUpdateConfigurationsClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<automation-account-name>",
		"<software-update-configuration-name>",
		armautomation.SoftwareUpdateConfiguration{
			Properties: &armautomation.SoftwareUpdateConfigurationProperties{
				ScheduleInfo: &armautomation.SUCScheduleProperties{
					AdvancedSchedule: &armautomation.AdvancedSchedule{
						WeekDays: []*string{
							to.StringPtr("Monday"),
							to.StringPtr("Thursday")},
					},
					ExpiryTime: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-11-09T11:22:57+00:00"); return t }()),
					Frequency:  armautomation.ScheduleFrequency("Hour").ToPtr(),
					Interval:   to.Int64Ptr(1),
					StartTime:  to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-10-19T12:22:57+00:00"); return t }()),
					TimeZone:   to.StringPtr("<time-zone>"),
				},
				Tasks: &armautomation.SoftwareUpdateConfigurationTasks{
					PostTask: &armautomation.TaskProperties{
						Source: to.StringPtr("<source>"),
					},
					PreTask: &armautomation.TaskProperties{
						Parameters: map[string]*string{
							"COMPUTERNAME": to.StringPtr("Computer1"),
						},
						Source: to.StringPtr("<source>"),
					},
				},
				UpdateConfiguration: &armautomation.UpdateConfiguration{
					AzureVirtualMachines: []*string{
						to.StringPtr("/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-01"),
						to.StringPtr("/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-02"),
						to.StringPtr("/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-03")},
					Duration: to.StringPtr("<duration>"),
					NonAzureComputerNames: []*string{
						to.StringPtr("box1.contoso.com"),
						to.StringPtr("box2.contoso.com")},
					OperatingSystem: armautomation.OperatingSystemTypeWindows.ToPtr(),
					Targets: &armautomation.TargetProperties{
						AzureQueries: []*armautomation.AzureQueryProperties{
							{
								Locations: []*string{
									to.StringPtr("Japan East"),
									to.StringPtr("UK South")},
								Scope: []*string{
									to.StringPtr("/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources"),
									to.StringPtr("/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067")},
								TagSettings: &armautomation.TagSettingsProperties{
									FilterOperator: armautomation.TagOperatorsAll.ToPtr(),
									Tags: map[string][]*string{
										"tag1": {
											to.StringPtr("tag1Value1"),
											to.StringPtr("tag1Value2"),
											to.StringPtr("tag1Value3")},
										"tag2": {
											to.StringPtr("tag2Value1"),
											to.StringPtr("tag2Value2"),
											to.StringPtr("tag2Value3")},
									},
								},
							}},
						NonAzureQueries: []*armautomation.NonAzureQueryProperties{
							{
								FunctionAlias: to.StringPtr("<function-alias>"),
								WorkspaceID:   to.StringPtr("<workspace-id>"),
							},
							{
								FunctionAlias: to.StringPtr("<function-alias>"),
								WorkspaceID:   to.StringPtr("<workspace-id>"),
							}},
					},
					Windows: &armautomation.WindowsProperties{
						ExcludedKbNumbers: []*string{
							to.StringPtr("168934"),
							to.StringPtr("168973")},
						IncludedUpdateClassifications: armautomation.WindowsUpdateClasses("Critical").ToPtr(),
						RebootSetting:                 to.StringPtr("<reboot-setting>"),
					},
				},
			},
		},
		&armautomation.SoftwareUpdateConfigurationsClientCreateOptions{ClientRequestID: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SoftwareUpdateConfigurationsClientCreateResult)
}

// x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/softwareUpdateConfiguration/getSoftwareUpdateConfigurationByName.json
func ExampleSoftwareUpdateConfigurationsClient_GetByName() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomation.NewSoftwareUpdateConfigurationsClient("<subscription-id>", cred, nil)
	res, err := client.GetByName(ctx,
		"<resource-group-name>",
		"<automation-account-name>",
		"<software-update-configuration-name>",
		&armautomation.SoftwareUpdateConfigurationsClientGetByNameOptions{ClientRequestID: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SoftwareUpdateConfigurationsClientGetByNameResult)
}

// x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/softwareUpdateConfiguration/deleteSoftwareUpdateConfiguration.json
func ExampleSoftwareUpdateConfigurationsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomation.NewSoftwareUpdateConfigurationsClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<automation-account-name>",
		"<software-update-configuration-name>",
		&armautomation.SoftwareUpdateConfigurationsClientDeleteOptions{ClientRequestID: nil})
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/softwareUpdateConfiguration/listSoftwareUpdateConfigurations.json
func ExampleSoftwareUpdateConfigurationsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armautomation.NewSoftwareUpdateConfigurationsClient("<subscription-id>", cred, nil)
	res, err := client.List(ctx,
		"<resource-group-name>",
		"<automation-account-name>",
		&armautomation.SoftwareUpdateConfigurationsClientListOptions{ClientRequestID: nil,
			Filter: nil,
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SoftwareUpdateConfigurationsClientListResult)
}
