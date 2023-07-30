/*
Immudb Cloud Service

Testing DocumentsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package storage

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/WladOd/auction.git/storage"
)

func Test_storage_DocumentsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DocumentsAPIService AuditDocument", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ledger string
		var collection string
		var documentId string

		resp, httpRes, err := apiClient.DocumentsAPI.AuditDocument(context.Background(), ledger, collection, documentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DocumentsAPIService CountDocuments", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ledger string
		var collection string

		resp, httpRes, err := apiClient.DocumentsAPI.CountDocuments(context.Background(), ledger, collection).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DocumentsAPIService DiffDocument", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ledger string
		var collection string
		var documentId string

		resp, httpRes, err := apiClient.DocumentsAPI.DiffDocument(context.Background(), ledger, collection, documentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DocumentsAPIService DocumentCreate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ledger string
		var collection string

		resp, httpRes, err := apiClient.DocumentsAPI.DocumentCreate(context.Background(), ledger, collection).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DocumentsAPIService DocumentCreateMany", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ledger string
		var collection string

		resp, httpRes, err := apiClient.DocumentsAPI.DocumentCreateMany(context.Background(), ledger, collection).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DocumentsAPIService GetDocumentProof", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ledger string
		var collection string
		var documentId string

		resp, httpRes, err := apiClient.DocumentsAPI.GetDocumentProof(context.Background(), ledger, collection, documentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DocumentsAPIService SearchDocument", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ledger string
		var collection string

		resp, httpRes, err := apiClient.DocumentsAPI.SearchDocument(context.Background(), ledger, collection).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DocumentsAPIService UpdateDocument", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ledger string
		var collection string

		resp, httpRes, err := apiClient.DocumentsAPI.UpdateDocument(context.Background(), ledger, collection).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
