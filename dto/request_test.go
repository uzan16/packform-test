package dto_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	. "github.com/uzan16/packform-test/dto"
)

func TestNewPaginationRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	assert := assert.New(t)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())

	u := url.Values{}
	u.Add("pageNumber", "2")
	u.Add("pageSize", "3")
	u.Add("sortBy", "createdAt")
	u.Add("sortDirection", "asc")
	u.Add("q", "test")

	c.Request = &http.Request{
		Header: make(http.Header),
		URL: &url.URL{
			RawQuery: u.Encode(),
		},
	}

	request := NewPaginationRequest(c)

	if assert.NotNil(request) {
		assert.Equal(2, request.PageNumber)
		assert.Equal(3, request.PageSize)
		assert.Equal("created_at", request.SortBy)
		assert.Equal("asc", request.SortDirection)
		assert.Equal("test", request.Query)
	}

	assert.Equal("created_at asc", request.Order())
	assert.Equal(3, request.Offset())
}
