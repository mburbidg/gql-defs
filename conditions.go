package gqldefs

// StatusMap maps from a GQLStatus to a class and subclass. The key of this map is the two character class string
// concatenated with the 3 character subclass string.
var StatusMap map[string]StatusSubclass = map[string]StatusSubclass{}

type StatusClass interface {
	Category() string
	Code() string
	Description() string
}

type statusClass struct {
	category    string
	code        string
	description string
}

func (s statusClass) Category() string {
	return s.category
}

func (s statusClass) Code() string {
	return s.code
}

func (s statusClass) Description() string {
	return s.description
}

type StatusSubclass interface {
	Code() string
	Subcode() string
	Description() string
	StatusClass() StatusClass
}

type statusSubclass struct {
	subcode     string
	description string
	statusClass StatusClass
}

func (s statusSubclass) Code() string {
	return s.statusClass.Code() + s.subcode
}

func (s statusSubclass) Subcode() string {
	return s.subcode
}

func (s statusSubclass) Description() string {
	return s.description
}

func (s statusSubclass) StatusClass() StatusClass {
	return s.statusClass
}

func newStatusSubclass(subcode string, description string, statusClass StatusClass) StatusSubclass {
	s := &statusSubclass{subcode: subcode, description: description, statusClass: statusClass}
	StatusMap[statusClass.Code()+subcode] = s
	return s
}

type succesfulCompletion struct {
	statusClass
	NoSubclass StatusSubclass
}

var SuccessfulCompletion *succesfulCompletion

type warning struct {
	statusClass
	NoSubclass                       StatusSubclass
	StringDataWriteTruncation        StatusSubclass
	GraphDoesNotExist                StatusSubclass
	GraphTypeDoesNotExist            StatusSubclass
	NullValueEliminatedInSetFunction StatusSubclass
}

var Warning *warning

type noData struct {
	statusClass
	NoSubclass StatusSubclass
}

var NoData *noData

type informational struct {
	statusClass
	NoSubclass StatusSubclass
}

var Informational *informational

type connectionException struct {
	statusClass
	NoSubclass                   StatusSubclass
	TransactionResolutionUnknown StatusSubclass
}

var ConnectionException *connectionException

type dataException struct {
	statusClass
	NoSubclass                                    StatusSubclass
	StringDataWriteTruncation                     StatusSubclass
	NumericValueOutOfRange                        StatusSubclass
	NullValueNotAllowed                           StatusSubclass
	InvalidDatetimeFormat                         StatusSubclass
	DatetimeFieldOverflow                         StatusSubclass
	SubstringError                                StatusSubclass
	DivisionByZero                                StatusSubclass
	IntervalFieldOverflow                         StatusSubclass
	InvalidCharacterValueForCast                  StatusSubclass
	InvalidArgumentForNaturalLogarithm            StatusSubclass
	InvalidArgumentForPowerFunction               StatusSubclass
	TrimError                                     StatusSubclass
	ArrayDataRightTruncation                      StatusSubclass
	NegativeLimitValue                            StatusSubclass
	InvalidValueType                              StatusSubclass
	ValueNotComparable                            StatusSubclass
	InvalidDatetimeFunctionFieldName              StatusSubclass
	InvalidDatetimeFunctionValue                  StatusSubclass
	InvalidDurationFieldName                      StatusSubclass
	ListDataRightTruncation                       StatusSubclass
	ListElementError                              StatusSubclass
	InvalidNumberOfPathsOrGroups                  StatusSubclass
	InvalidDurationFormat                         StatusSubclass
	MultipleAssignmentsToAGraphElementProperty    StatusSubclass
	NumberOfNodeLabelsBelowSupportedMinimum       StatusSubclass
	NumberOfNodeLabelsExceedsSupportedMaximum     StatusSubclass
	NumberOfEdgeLabelsBelowSupportedMinimum       StatusSubclass
	NumberOfEdgeLabelsExceedsSupportedMaximum     StatusSubclass
	NumberOfNodePropertiesExceedsSupportedMaximum StatusSubclass
	NumberOfEdgePropertiesExceedsSupportedMaximum StatusSubclass
	RecordFieldsDoNotMatch                        StatusSubclass
	ReferenceValueInvalidBaseType                 StatusSubclass
	ReferenceValueInvalidConstrainedType          StatusSubclass
	RecordDataFieldUnassignable                   StatusSubclass
	RecordDataFieldMissing                        StatusSubclass
	MalformedPath                                 StatusSubclass
	PathDataRightTruncation                       StatusSubclass
	ReferenceValueReferentDeleted                 StatusSubclass
	InvalidValueTypeG12                           StatusSubclass
	InvalidGroupVariableValue                     StatusSubclass
	IncompatibleTemporalInstantUnitGroups         StatusSubclass
}

var DataException *dataException

type invalidTransactionState struct {
	statusClass
	NoSubclass                                StatusSubclass
	ActiveGQLTransaction                      StatusSubclass
	CatalogAndDataStatementMixingNotSupported StatusSubclass
	ReadOnlyTransaction                       StatusSubclass
	AccessingMultipleGraphsNotSupported       StatusSubclass
}

var InvalidTransactionState *invalidTransactionState

type invalidTransactionTermination struct {
	statusClass
	NoSubclass StatusSubclass
}

var InvalidTransactionTermination *invalidTransactionTermination

type transactionRollback struct {
	statusClass
	NoSubclass                   StatusSubclass
	IntegrityConstraintViolation StatusSubclass
	StatementCompletionUnknown   StatusSubclass
}

var TransactionRollback *transactionRollback

type syntaxErrorOrAccessRuleViolation struct {
	statusClass
	NoSubclass                                      StatusSubclass
	InvalidSyntax                                   StatusSubclass
	InvalidReference                                StatusSubclass
	CannotDropTheCurrentWorkingSchema               StatusSubclass
	UseOfVisuallyConfusableIdentifiers              StatusSubclass
	EndpointNodeTypeNotDefinedInGraphTypeDefinition StatusSubclass
	NumberOfEdgeLabelsBelowSupportedMinimum         StatusSubclass
	NumberOfEdgeLabelsExceedsSupportedMaximum       StatusSubclass
	NumberOfEdgePropertiesExceedsSupportedMinimum   StatusSubclass
	NumberOfNodeLabelsBelowSupportedMinimum         StatusSubclass
	NumberOfNodeLabelsExceedsSupportedMaximum       StatusSubclass
	NumberOfNodePropertiesExceedsSupportedMinimum   StatusSubclass
}

var SyntaxErrorOrAccessRuleViolation *syntaxErrorOrAccessRuleViolation

type dependentObjectError struct {
	statusClass
	NoSubclass      StatusSubclass
	EdgesStillExist StatusSubclass
}

var DependentObjectError *dependentObjectError

type graphTypeViolation struct {
	statusClass
	NoSubclass StatusSubclass
}

var GraphTypeViolation *graphTypeViolation

func Code(gqlStatus string) string {
	return FirstN(gqlStatus, 2)
}

func Subcode(gqlStatus string) string {
	return LastN(gqlStatus, 3)
}

func init() {
	SuccessfulCompletion = &succesfulCompletion{
		statusClass: statusClass{category: "S", code: "00", description: "successful completion"},
	}
	SuccessfulCompletion.NoSubclass = newStatusSubclass("000", "(no subclass)", SuccessfulCompletion)

	Warning = &warning{
		statusClass: statusClass{category: "W", code: "01", description: "warning"},
	}
	Warning.NoSubclass = newStatusSubclass("000", "(no subclass)", Warning)
	Warning.StringDataWriteTruncation = newStatusSubclass("004", "string data, right truncation", Warning)
	Warning.GraphDoesNotExist = newStatusSubclass("G03", "graph does not exist", Warning)
	Warning.GraphTypeDoesNotExist = newStatusSubclass("G04", "graph type does not exist", Warning)
	Warning.NullValueEliminatedInSetFunction = newStatusSubclass("G11", "null value eliminated in set function", Warning)

	NoData = &noData{
		statusClass: statusClass{category: "N", code: "02", description: "no data"},
	}
	NoData.NoSubclass = newStatusSubclass("000", "(no subclass)", NoData)

	Informational = &informational{
		statusClass: statusClass{category: "I", code: "03", description: "no data"},
	}
	Informational.NoSubclass = newStatusSubclass("000", "(no subclass)", Informational)

	ConnectionException = &connectionException{
		statusClass: statusClass{category: "X", code: "08", description: "connection exception"},
	}
	ConnectionException.NoSubclass = newStatusSubclass("000", "(no subclass)", ConnectionException)
	ConnectionException.TransactionResolutionUnknown = newStatusSubclass("007", "transaction resolution unknown", ConnectionException)

	DataException = &dataException{
		statusClass: statusClass{category: "X", code: "22", description: "data exception"},
	}
	DataException.NoSubclass = newStatusSubclass("000", "(no subclass)", DataException)
	DataException.StringDataWriteTruncation = newStatusSubclass("001", "string data, right truncation", DataException)
	DataException.NumericValueOutOfRange = newStatusSubclass("003", "numeric value out of range", DataException)
	DataException.NullValueNotAllowed = newStatusSubclass("004", "null value not allowed", DataException)
	DataException.InvalidDatetimeFormat = newStatusSubclass("007", "invalid datetime format", DataException)
	DataException.DatetimeFieldOverflow = newStatusSubclass("008", "datetime field overflow", DataException)
	DataException.SubstringError = newStatusSubclass("011", "substring error", DataException)
	DataException.DivisionByZero = newStatusSubclass("012", "division by zero", DataException)
	DataException.IntervalFieldOverflow = newStatusSubclass("015", "interval field overflow", DataException)
	DataException.InvalidCharacterValueForCast = newStatusSubclass("018", "invalid character value for cast", DataException)
	DataException.InvalidArgumentForNaturalLogarithm = newStatusSubclass("01E", "invalid argument for natural logarithm", DataException)
	DataException.InvalidArgumentForPowerFunction = newStatusSubclass("01F", "invalid argument for power function", DataException)
	DataException.TrimError = newStatusSubclass("027", "trim error", DataException)
	DataException.ArrayDataRightTruncation = newStatusSubclass("02F", "array data, right truncation", DataException)
	DataException.NegativeLimitValue = newStatusSubclass("G02", "negative limit value", DataException)
	DataException.InvalidValueType = newStatusSubclass("G03", "invalid value type", DataException)
	DataException.ValueNotComparable = newStatusSubclass("G04", "alues not comparable", DataException)
	DataException.InvalidDatetimeFunctionFieldName = newStatusSubclass("G05", "invalid datetime function field name", DataException)
	DataException.InvalidDatetimeFunctionValue = newStatusSubclass("G06", "invalid datetime function value", DataException)
	DataException.InvalidDurationFieldName = newStatusSubclass("G07", "invalid duration function field name", DataException)
	DataException.ListDataRightTruncation = newStatusSubclass("G0B", "list data, right truncation", DataException)
	DataException.ListElementError = newStatusSubclass("G0C", "list element error", DataException)
	DataException.InvalidNumberOfPathsOrGroups = newStatusSubclass("G0F", "invalid number of paths or groups", DataException)
	DataException.InvalidDurationFormat = newStatusSubclass("G0H", "invalid duration format", DataException)
	DataException.MultipleAssignmentsToAGraphElementProperty = newStatusSubclass("G0M", "multiple assignments to a graph element property", DataException)
	DataException.NumberOfNodeLabelsBelowSupportedMinimum = newStatusSubclass("G0N", "number of node labels below supported minimum", DataException)
	DataException.NumberOfNodeLabelsExceedsSupportedMaximum = newStatusSubclass("G0P", "number of node labels exceeds supported maximum", DataException)
	DataException.NumberOfEdgeLabelsBelowSupportedMinimum = newStatusSubclass("G0Q", "number of edge labels below supported minimum", DataException)
	DataException.NumberOfEdgeLabelsExceedsSupportedMaximum = newStatusSubclass("G0R", "number of edge labels exceeds supported maximum", DataException)
	DataException.NumberOfNodePropertiesExceedsSupportedMaximum = newStatusSubclass("G0S", "number of node properties exceeds supported maximum", DataException)
	DataException.NumberOfEdgePropertiesExceedsSupportedMaximum = newStatusSubclass("G0T", "number of edge properties exceeds supported maximum", DataException)
	DataException.RecordFieldsDoNotMatch = newStatusSubclass("G0U", "record fields do not match", DataException)
	DataException.ReferenceValueInvalidBaseType = newStatusSubclass("G0V", "reference value, invalid base type", DataException)
	DataException.ReferenceValueInvalidConstrainedType = newStatusSubclass("G0W", "reference value, invalid constrained type", DataException)
	DataException.RecordDataFieldUnassignable = newStatusSubclass("G0X", "record data, field unassignable", DataException)
	DataException.RecordDataFieldMissing = newStatusSubclass("G0Y", "record data, field missing", DataException)
	DataException.MalformedPath = newStatusSubclass("G0Z", "malformed path", DataException)
	DataException.PathDataRightTruncation = newStatusSubclass("G10", "path data, right truncation", DataException)
	DataException.ReferenceValueReferentDeleted = newStatusSubclass("G11", "reference value, referent deleted", DataException)
	DataException.InvalidValueTypeG12 = newStatusSubclass("G12", "invalid value type", DataException)
	DataException.InvalidGroupVariableValue = newStatusSubclass("G13", "invalid group variable value", DataException)
	DataException.IncompatibleTemporalInstantUnitGroups = newStatusSubclass("G14", "incompatible temporal instant unit groups", DataException)

	InvalidTransactionState = &invalidTransactionState{
		statusClass: statusClass{category: "X", code: "25", description: "invalid transaction state"},
	}
	InvalidTransactionState.NoSubclass = newStatusSubclass("000", "(no subclass)", InvalidTransactionState)
	InvalidTransactionState.ActiveGQLTransaction = newStatusSubclass("G01", "active GQL-transaction", InvalidTransactionState)
	InvalidTransactionState.CatalogAndDataStatementMixingNotSupported = newStatusSubclass("G02", "catalog and data statement mixing not supported", InvalidTransactionState)
	InvalidTransactionState.ReadOnlyTransaction = newStatusSubclass("G03", "read-only transaction", InvalidTransactionState)
	InvalidTransactionState.AccessingMultipleGraphsNotSupported = newStatusSubclass("G04", "accessing multiple graphs not supported", InvalidTransactionState)

	InvalidTransactionTermination = &invalidTransactionTermination{
		statusClass: statusClass{category: "X", code: "2D", description: "invalid transaction termination"},
	}
	InvalidTransactionTermination.NoSubclass = newStatusSubclass("000", "(no subclass)", InvalidTransactionTermination)

	TransactionRollback = &transactionRollback{
		statusClass: statusClass{category: "X", code: "40", description: "transaction rollback"},
	}
	TransactionRollback.NoSubclass = newStatusSubclass("000", "(no subclass)", TransactionRollback)
	TransactionRollback.IntegrityConstraintViolation = newStatusSubclass("002", "integrity constraint violation", TransactionRollback)
	TransactionRollback.StatementCompletionUnknown = newStatusSubclass("003", "statement completion unknown", TransactionRollback)

	SyntaxErrorOrAccessRuleViolation = &syntaxErrorOrAccessRuleViolation{
		statusClass: statusClass{category: "X", code: "42", description: "syntax error or access rule violation"},
	}
	SyntaxErrorOrAccessRuleViolation.NoSubclass = newStatusSubclass("000", "(no subclass)", SyntaxErrorOrAccessRuleViolation)
	SyntaxErrorOrAccessRuleViolation.InvalidSyntax = newStatusSubclass("001", "invalid syntax", SyntaxErrorOrAccessRuleViolation)
	SyntaxErrorOrAccessRuleViolation.InvalidReference = newStatusSubclass("002", "invalid reference", SyntaxErrorOrAccessRuleViolation)
	SyntaxErrorOrAccessRuleViolation.CannotDropTheCurrentWorkingSchema = newStatusSubclass("003", "cannot drop the current working schema", SyntaxErrorOrAccessRuleViolation)
	SyntaxErrorOrAccessRuleViolation.UseOfVisuallyConfusableIdentifiers = newStatusSubclass("004", "use of visually confusable identifiers", SyntaxErrorOrAccessRuleViolation)
	SyntaxErrorOrAccessRuleViolation.EndpointNodeTypeNotDefinedInGraphTypeDefinition = newStatusSubclass("005", "endpoint node type not defined in graph type definition", SyntaxErrorOrAccessRuleViolation)
	SyntaxErrorOrAccessRuleViolation.NumberOfEdgeLabelsBelowSupportedMinimum = newStatusSubclass("006", "number of edge labels below supported minimum", SyntaxErrorOrAccessRuleViolation)
	SyntaxErrorOrAccessRuleViolation.NumberOfEdgeLabelsExceedsSupportedMaximum = newStatusSubclass("007", "number of edge labels exceeds supported maximum", SyntaxErrorOrAccessRuleViolation)
	SyntaxErrorOrAccessRuleViolation.NumberOfEdgePropertiesExceedsSupportedMinimum = newStatusSubclass("008", "number of edge properties exceeds supported maximum", SyntaxErrorOrAccessRuleViolation)
	SyntaxErrorOrAccessRuleViolation.NumberOfNodeLabelsBelowSupportedMinimum = newStatusSubclass("009", "number of node labels below supported minimum", SyntaxErrorOrAccessRuleViolation)
	SyntaxErrorOrAccessRuleViolation.NumberOfNodeLabelsExceedsSupportedMaximum = newStatusSubclass("010", "number of node labels exceeds supported maximum", SyntaxErrorOrAccessRuleViolation)
	SyntaxErrorOrAccessRuleViolation.NumberOfNodePropertiesExceedsSupportedMinimum = newStatusSubclass("000", "number of node properties exceeds supported maximum", SyntaxErrorOrAccessRuleViolation)

	DependentObjectError = &dependentObjectError{
		statusClass: statusClass{category: "X", code: "G1", description: "dependent object error"},
	}
	DependentObjectError.NoSubclass = newStatusSubclass("000", "(no subclass)", DependentObjectError)
	DependentObjectError.EdgesStillExist = newStatusSubclass("001", "edges still exist", DependentObjectError)

	GraphTypeViolation = &graphTypeViolation{
		statusClass: statusClass{category: "X", code: "G2", description: "graph type violation"},
	}
	GraphTypeViolation.NoSubclass = newStatusSubclass("000", "(no subclass)", GraphTypeViolation)
}
