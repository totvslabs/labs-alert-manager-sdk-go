/*
alertmanager

Testing HealthCheckAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package labs_alert_manager_client

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_labs_alert_manager_client_HealthCheckAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test HealthCheckAPIService GetHealthCheck", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.HealthCheckAPI.GetHealthCheck(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}