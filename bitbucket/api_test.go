package bitbucket

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func addAuthHeader(_ context.Context, req *http.Request) error {
	// Basic auth, username "admin", password "admin"
	req.Header.Set("Authorization", "Basic YWRtaW46YWRtaW4=")
	return nil
}

func newClient(t *testing.T) *ClientWithResponses {
	t.Helper()
	client, err := NewClientWithResponses("http://localhost:7990/rest", WithRequestEditorFn(addAuthHeader))
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	return client
}

func TestSanity(t *testing.T) {
	client := newClient(t)
	ctx := context.Background()
	res, err := client.Get2WithResponse(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, 200, res.StatusCode(), string(res.Body))
}

func TestListProjects(t *testing.T) {
	client := newClient(t)
	ctx := context.Background()
	res, err := client.GetProjectsWithResponse(ctx, nil)
	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode(), string(res.Body))
	assert.Equal(t, 0, int(*res.JSON200.Size))
}

func TestCreateAndDeleteProject(t *testing.T) {
	client := newClient(t)
	ctx := context.Background()
	projectName := "test"
	createRes, err := client.CreateProjectWithResponse(ctx, RestProject{Key: &projectName})
	assert.NoError(t, err)
	assert.Equal(t, 201, createRes.StatusCode(), string(createRes.Body))
	assert.Equal(t, *createRes.JSON201.Name, projectName)

	deleteRes, err := client.DeleteProjectWithResponse(ctx, projectName)
	assert.NoError(t, err)
	assert.Equal(t, 204, deleteRes.StatusCode(), string(deleteRes.Body))
}
