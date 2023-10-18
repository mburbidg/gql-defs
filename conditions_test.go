package gqldefs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSuccessfulCompletion(t *testing.T) {
	assert.Equal(t, StatusMap["00000"], SuccessfulCompletion.NoSubclass)
}

func TestWarning(t *testing.T) {
	assert.Equal(t, StatusMap["01000"], Warning.NoSubclass)
	assert.Equal(t, StatusMap["01004"], Warning.StringDataWriteTruncation)
	assert.Equal(t, StatusMap["01G03"], Warning.GraphDoesNotExist)
	assert.Equal(t, StatusMap["01G04"], Warning.GraphTypeDoesNotExist)
	assert.Equal(t, StatusMap["01G11"], Warning.NullValueEliminatedInSetFunction)
}

func TestNoData(t *testing.T) {
	assert.Equal(t, StatusMap["02000"], NoData.NoSubclass)
}

func TestConnectionException(t *testing.T) {
	assert.Equal(t, StatusMap["08000"], ConnectionException.NoSubclass)
	assert.Equal(t, StatusMap["08007"], ConnectionException.TransactionResolutionUnknown)
}
