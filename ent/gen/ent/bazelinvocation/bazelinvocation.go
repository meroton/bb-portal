// Code generated by ent, DO NOT EDIT.

package bazelinvocation

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the bazelinvocation type in the database.
	Label = "bazel_invocation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldInvocationID holds the string denoting the invocation_id field in the database.
	FieldInvocationID = "invocation_id"
	// FieldStartedAt holds the string denoting the started_at field in the database.
	FieldStartedAt = "started_at"
	// FieldEndedAt holds the string denoting the ended_at field in the database.
	FieldEndedAt = "ended_at"
	// FieldChangeNumber holds the string denoting the change_number field in the database.
	FieldChangeNumber = "change_number"
	// FieldPatchsetNumber holds the string denoting the patchset_number field in the database.
	FieldPatchsetNumber = "patchset_number"
	// FieldSummary holds the string denoting the summary field in the database.
	FieldSummary = "summary"
	// FieldBepCompleted holds the string denoting the bep_completed field in the database.
	FieldBepCompleted = "bep_completed"
	// FieldStepLabel holds the string denoting the step_label field in the database.
	FieldStepLabel = "step_label"
	// FieldRelatedFiles holds the string denoting the related_files field in the database.
	FieldRelatedFiles = "related_files"
	// FieldUserEmail holds the string denoting the user_email field in the database.
	FieldUserEmail = "user_email"
	// FieldUserLdap holds the string denoting the user_ldap field in the database.
	FieldUserLdap = "user_ldap"
	// FieldBuildLogs holds the string denoting the build_logs field in the database.
	FieldBuildLogs = "build_logs"
	// FieldCPU holds the string denoting the cpu field in the database.
	FieldCPU = "cpu"
	// FieldPlatformName holds the string denoting the platform_name field in the database.
	FieldPlatformName = "platform_name"
	// FieldHostname holds the string denoting the hostname field in the database.
	FieldHostname = "hostname"
	// FieldIsCiWorker holds the string denoting the is_ci_worker field in the database.
	FieldIsCiWorker = "is_ci_worker"
	// FieldConfigurationMnemonic holds the string denoting the configuration_mnemonic field in the database.
	FieldConfigurationMnemonic = "configuration_mnemonic"
	// FieldNumFetches holds the string denoting the num_fetches field in the database.
	FieldNumFetches = "num_fetches"
	// FieldProfileName holds the string denoting the profile_name field in the database.
	FieldProfileName = "profile_name"
	// EdgeEventFile holds the string denoting the event_file edge name in mutations.
	EdgeEventFile = "event_file"
	// EdgeBuild holds the string denoting the build edge name in mutations.
	EdgeBuild = "build"
	// EdgeProblems holds the string denoting the problems edge name in mutations.
	EdgeProblems = "problems"
	// EdgeMetrics holds the string denoting the metrics edge name in mutations.
	EdgeMetrics = "metrics"
	// EdgeTestCollection holds the string denoting the test_collection edge name in mutations.
	EdgeTestCollection = "test_collection"
	// EdgeTargets holds the string denoting the targets edge name in mutations.
	EdgeTargets = "targets"
	// EdgeSourceControl holds the string denoting the source_control edge name in mutations.
	EdgeSourceControl = "source_control"
	// Table holds the table name of the bazelinvocation in the database.
	Table = "bazel_invocations"
	// EventFileTable is the table that holds the event_file relation/edge.
	EventFileTable = "bazel_invocations"
	// EventFileInverseTable is the table name for the EventFile entity.
	// It exists in this package in order to avoid circular dependency with the "eventfile" package.
	EventFileInverseTable = "event_files"
	// EventFileColumn is the table column denoting the event_file relation/edge.
	EventFileColumn = "event_file_bazel_invocation"
	// BuildTable is the table that holds the build relation/edge.
	BuildTable = "bazel_invocations"
	// BuildInverseTable is the table name for the Build entity.
	// It exists in this package in order to avoid circular dependency with the "build" package.
	BuildInverseTable = "builds"
	// BuildColumn is the table column denoting the build relation/edge.
	BuildColumn = "build_invocations"
	// ProblemsTable is the table that holds the problems relation/edge.
	ProblemsTable = "bazel_invocation_problems"
	// ProblemsInverseTable is the table name for the BazelInvocationProblem entity.
	// It exists in this package in order to avoid circular dependency with the "bazelinvocationproblem" package.
	ProblemsInverseTable = "bazel_invocation_problems"
	// ProblemsColumn is the table column denoting the problems relation/edge.
	ProblemsColumn = "bazel_invocation_problems"
	// MetricsTable is the table that holds the metrics relation/edge.
	MetricsTable = "metrics"
	// MetricsInverseTable is the table name for the Metrics entity.
	// It exists in this package in order to avoid circular dependency with the "metrics" package.
	MetricsInverseTable = "metrics"
	// MetricsColumn is the table column denoting the metrics relation/edge.
	MetricsColumn = "bazel_invocation_metrics"
	// TestCollectionTable is the table that holds the test_collection relation/edge.
	TestCollectionTable = "test_collections"
	// TestCollectionInverseTable is the table name for the TestCollection entity.
	// It exists in this package in order to avoid circular dependency with the "testcollection" package.
	TestCollectionInverseTable = "test_collections"
	// TestCollectionColumn is the table column denoting the test_collection relation/edge.
	TestCollectionColumn = "bazel_invocation_test_collection"
	// TargetsTable is the table that holds the targets relation/edge.
	TargetsTable = "target_pairs"
	// TargetsInverseTable is the table name for the TargetPair entity.
	// It exists in this package in order to avoid circular dependency with the "targetpair" package.
	TargetsInverseTable = "target_pairs"
	// TargetsColumn is the table column denoting the targets relation/edge.
	TargetsColumn = "bazel_invocation_targets"
	// SourceControlTable is the table that holds the source_control relation/edge.
	SourceControlTable = "source_controls"
	// SourceControlInverseTable is the table name for the SourceControl entity.
	// It exists in this package in order to avoid circular dependency with the "sourcecontrol" package.
	SourceControlInverseTable = "source_controls"
	// SourceControlColumn is the table column denoting the source_control relation/edge.
	SourceControlColumn = "bazel_invocation_source_control"
)

// Columns holds all SQL columns for bazelinvocation fields.
var Columns = []string{
	FieldID,
	FieldInvocationID,
	FieldStartedAt,
	FieldEndedAt,
	FieldChangeNumber,
	FieldPatchsetNumber,
	FieldSummary,
	FieldBepCompleted,
	FieldStepLabel,
	FieldRelatedFiles,
	FieldUserEmail,
	FieldUserLdap,
	FieldBuildLogs,
	FieldCPU,
	FieldPlatformName,
	FieldHostname,
	FieldIsCiWorker,
	FieldConfigurationMnemonic,
	FieldNumFetches,
	FieldProfileName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "bazel_invocations"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"build_invocations",
	"event_file_bazel_invocation",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the BazelInvocation queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByInvocationID orders the results by the invocation_id field.
func ByInvocationID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInvocationID, opts...).ToFunc()
}

// ByStartedAt orders the results by the started_at field.
func ByStartedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartedAt, opts...).ToFunc()
}

// ByEndedAt orders the results by the ended_at field.
func ByEndedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndedAt, opts...).ToFunc()
}

// ByChangeNumber orders the results by the change_number field.
func ByChangeNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChangeNumber, opts...).ToFunc()
}

// ByPatchsetNumber orders the results by the patchset_number field.
func ByPatchsetNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPatchsetNumber, opts...).ToFunc()
}

// ByBepCompleted orders the results by the bep_completed field.
func ByBepCompleted(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBepCompleted, opts...).ToFunc()
}

// ByStepLabel orders the results by the step_label field.
func ByStepLabel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStepLabel, opts...).ToFunc()
}

// ByUserEmail orders the results by the user_email field.
func ByUserEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserEmail, opts...).ToFunc()
}

// ByUserLdap orders the results by the user_ldap field.
func ByUserLdap(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserLdap, opts...).ToFunc()
}

// ByBuildLogs orders the results by the build_logs field.
func ByBuildLogs(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBuildLogs, opts...).ToFunc()
}

// ByCPU orders the results by the cpu field.
func ByCPU(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCPU, opts...).ToFunc()
}

// ByPlatformName orders the results by the platform_name field.
func ByPlatformName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlatformName, opts...).ToFunc()
}

// ByHostname orders the results by the hostname field.
func ByHostname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHostname, opts...).ToFunc()
}

// ByIsCiWorker orders the results by the is_ci_worker field.
func ByIsCiWorker(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsCiWorker, opts...).ToFunc()
}

// ByConfigurationMnemonic orders the results by the configuration_mnemonic field.
func ByConfigurationMnemonic(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldConfigurationMnemonic, opts...).ToFunc()
}

// ByNumFetches orders the results by the num_fetches field.
func ByNumFetches(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNumFetches, opts...).ToFunc()
}

// ByProfileName orders the results by the profile_name field.
func ByProfileName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProfileName, opts...).ToFunc()
}

// ByEventFileField orders the results by event_file field.
func ByEventFileField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEventFileStep(), sql.OrderByField(field, opts...))
	}
}

// ByBuildField orders the results by build field.
func ByBuildField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBuildStep(), sql.OrderByField(field, opts...))
	}
}

// ByProblemsCount orders the results by problems count.
func ByProblemsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProblemsStep(), opts...)
	}
}

// ByProblems orders the results by problems terms.
func ByProblems(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProblemsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMetricsField orders the results by metrics field.
func ByMetricsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMetricsStep(), sql.OrderByField(field, opts...))
	}
}

// ByTestCollectionCount orders the results by test_collection count.
func ByTestCollectionCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTestCollectionStep(), opts...)
	}
}

// ByTestCollection orders the results by test_collection terms.
func ByTestCollection(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTestCollectionStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTargetsCount orders the results by targets count.
func ByTargetsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTargetsStep(), opts...)
	}
}

// ByTargets orders the results by targets terms.
func ByTargets(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTargetsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySourceControlField orders the results by source_control field.
func BySourceControlField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSourceControlStep(), sql.OrderByField(field, opts...))
	}
}
func newEventFileStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EventFileInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, EventFileTable, EventFileColumn),
	)
}
func newBuildStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BuildInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, BuildTable, BuildColumn),
	)
}
func newProblemsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProblemsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ProblemsTable, ProblemsColumn),
	)
}
func newMetricsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MetricsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, MetricsTable, MetricsColumn),
	)
}
func newTestCollectionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TestCollectionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TestCollectionTable, TestCollectionColumn),
	)
}
func newTargetsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TargetsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TargetsTable, TargetsColumn),
	)
}
func newSourceControlStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SourceControlInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, SourceControlTable, SourceControlColumn),
	)
}
