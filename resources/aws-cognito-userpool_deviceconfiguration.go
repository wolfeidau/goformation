package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Cognito::UserPool.DeviceConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-deviceconfiguration.html
type AWSCognitoUserPool_DeviceConfiguration struct {

	// ChallengeRequiredOnNewDevice AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-deviceconfiguration.html#cfn-cognito-userpool-deviceconfiguration-challengerequiredonnewdevice

	ChallengeRequiredOnNewDevice bool `json:"ChallengeRequiredOnNewDevice"`

	// DeviceOnlyRememberedOnUserPrompt AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-deviceconfiguration.html#cfn-cognito-userpool-deviceconfiguration-deviceonlyrememberedonuserprompt

	DeviceOnlyRememberedOnUserPrompt bool `json:"DeviceOnlyRememberedOnUserPrompt"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPool_DeviceConfiguration) AWSCloudFormationType() string {
	return "AWS::Cognito::UserPool.DeviceConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoUserPool_DeviceConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCognitoUserPool_DeviceConfigurationResources retrieves all AWSCognitoUserPool_DeviceConfiguration items from a CloudFormation template
func GetAllAWSCognitoUserPool_DeviceConfiguration(template *Template) map[string]*AWSCognitoUserPool_DeviceConfiguration {

	results := map[string]*AWSCognitoUserPool_DeviceConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSCognitoUserPool_DeviceConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCognitoUserPool_DeviceConfigurationWithName retrieves all AWSCognitoUserPool_DeviceConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCognitoUserPool_DeviceConfiguration(name string, template *Template) (*AWSCognitoUserPool_DeviceConfiguration, error) {

	result := &AWSCognitoUserPool_DeviceConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCognitoUserPool_DeviceConfiguration{}, errors.New("resource not found")

}