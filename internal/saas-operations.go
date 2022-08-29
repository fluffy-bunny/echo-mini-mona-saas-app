package internal

import (
	"github.com/fluffy-bunny/echo_mini_mona_saas_app/internal/operations_apis"
	"github.com/fluffy-bunny/echo_mini_mona_saas_app/models"
)

func SendOperationUpdateSuccess(subID, opID string) error {
	operationAPIs := operations_apis.New(AzureHttpClient)
	err := operationAPIs.Update(subID, opID, &models.OperationsUpdateRequest{
		Status: "Success", // "Success" or "Failure"
	})
	return err
}
func SendOperationUpdateFailure(subID, opID string) error {
	operationAPIs := operations_apis.New(AzureHttpClient)
	err := operationAPIs.Update(subID, opID, &models.OperationsUpdateRequest{
		Status: "Failure", // "Success" or "Failure"
	})
	return err
}
