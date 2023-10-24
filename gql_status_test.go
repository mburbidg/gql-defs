package gqldefs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGQLStatus(t *testing.T) {
	assert.Equal(t, "note: successful completion", NewGQLStatus(SuccessfulCompletion.NoSubclass).Description)
	assert.Equal(t, "warn: (no subclass)", NewGQLStatus(Warning.NoSubclass).Description)
	assert.Equal(t, "warn: string data, right truncation", NewGQLStatus(Warning.StringDataWriteTruncation).Description)
	assert.Equal(t, "warn: graph does not exist", NewGQLStatus(Warning.GraphDoesNotExist).Description)
	assert.Equal(t, "note: no data", NewGQLStatus(NoData.NoSubclass).Description)
	assert.Equal(t, "error: data exception - datetime field overflow", NewGQLStatus(DataException.DatetimeFieldOverflow).Description)
	assert.Equal(t, "error: invalid transaction state - active GQL-transaction", NewGQLStatus(InvalidTransactionState.ActiveGQLTransaction).Description)
	assert.Equal(t, "error: syntax error or access rule violation", NewGQLStatus(SyntaxErrorOrAccessRuleViolation.NoSubclass).Description)
	assert.Equal(t, "error: syntax error or access rule violation - cannot drop the current working schema", NewGQLStatus(SyntaxErrorOrAccessRuleViolation.CannotDropTheCurrentWorkingSchema).Description)
}
