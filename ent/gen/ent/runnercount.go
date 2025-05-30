// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/buildbarn/bb-portal/ent/gen/ent/actionsummary"
	"github.com/buildbarn/bb-portal/ent/gen/ent/runnercount"
)

// RunnerCount is the model entity for the RunnerCount schema.
type RunnerCount struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// ExecKind holds the value of the "exec_kind" field.
	ExecKind string `json:"exec_kind,omitempty"`
	// ActionsExecuted holds the value of the "actions_executed" field.
	ActionsExecuted int64 `json:"actions_executed,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RunnerCountQuery when eager-loading is set.
	Edges                       RunnerCountEdges `json:"edges"`
	action_summary_runner_count *int
	selectValues                sql.SelectValues
}

// RunnerCountEdges holds the relations/edges for other nodes in the graph.
type RunnerCountEdges struct {
	// ActionSummary holds the value of the action_summary edge.
	ActionSummary *ActionSummary `json:"action_summary,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// ActionSummaryOrErr returns the ActionSummary value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RunnerCountEdges) ActionSummaryOrErr() (*ActionSummary, error) {
	if e.ActionSummary != nil {
		return e.ActionSummary, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: actionsummary.Label}
	}
	return nil, &NotLoadedError{edge: "action_summary"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RunnerCount) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case runnercount.FieldID, runnercount.FieldActionsExecuted:
			values[i] = new(sql.NullInt64)
		case runnercount.FieldName, runnercount.FieldExecKind:
			values[i] = new(sql.NullString)
		case runnercount.ForeignKeys[0]: // action_summary_runner_count
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RunnerCount fields.
func (rc *RunnerCount) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case runnercount.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rc.ID = int(value.Int64)
		case runnercount.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				rc.Name = value.String
			}
		case runnercount.FieldExecKind:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field exec_kind", values[i])
			} else if value.Valid {
				rc.ExecKind = value.String
			}
		case runnercount.FieldActionsExecuted:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field actions_executed", values[i])
			} else if value.Valid {
				rc.ActionsExecuted = value.Int64
			}
		case runnercount.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field action_summary_runner_count", value)
			} else if value.Valid {
				rc.action_summary_runner_count = new(int)
				*rc.action_summary_runner_count = int(value.Int64)
			}
		default:
			rc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RunnerCount.
// This includes values selected through modifiers, order, etc.
func (rc *RunnerCount) Value(name string) (ent.Value, error) {
	return rc.selectValues.Get(name)
}

// QueryActionSummary queries the "action_summary" edge of the RunnerCount entity.
func (rc *RunnerCount) QueryActionSummary() *ActionSummaryQuery {
	return NewRunnerCountClient(rc.config).QueryActionSummary(rc)
}

// Update returns a builder for updating this RunnerCount.
// Note that you need to call RunnerCount.Unwrap() before calling this method if this RunnerCount
// was returned from a transaction, and the transaction was committed or rolled back.
func (rc *RunnerCount) Update() *RunnerCountUpdateOne {
	return NewRunnerCountClient(rc.config).UpdateOne(rc)
}

// Unwrap unwraps the RunnerCount entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rc *RunnerCount) Unwrap() *RunnerCount {
	_tx, ok := rc.config.driver.(*txDriver)
	if !ok {
		panic("ent: RunnerCount is not a transactional entity")
	}
	rc.config.driver = _tx.drv
	return rc
}

// String implements the fmt.Stringer.
func (rc *RunnerCount) String() string {
	var builder strings.Builder
	builder.WriteString("RunnerCount(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rc.ID))
	builder.WriteString("name=")
	builder.WriteString(rc.Name)
	builder.WriteString(", ")
	builder.WriteString("exec_kind=")
	builder.WriteString(rc.ExecKind)
	builder.WriteString(", ")
	builder.WriteString("actions_executed=")
	builder.WriteString(fmt.Sprintf("%v", rc.ActionsExecuted))
	builder.WriteByte(')')
	return builder.String()
}

// RunnerCounts is a parsable slice of RunnerCount.
type RunnerCounts []*RunnerCount
