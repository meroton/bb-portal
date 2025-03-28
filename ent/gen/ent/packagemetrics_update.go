// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/buildbarn/bb-portal/ent/gen/ent/metrics"
	"github.com/buildbarn/bb-portal/ent/gen/ent/packageloadmetrics"
	"github.com/buildbarn/bb-portal/ent/gen/ent/packagemetrics"
	"github.com/buildbarn/bb-portal/ent/gen/ent/predicate"
)

// PackageMetricsUpdate is the builder for updating PackageMetrics entities.
type PackageMetricsUpdate struct {
	config
	hooks    []Hook
	mutation *PackageMetricsMutation
}

// Where appends a list predicates to the PackageMetricsUpdate builder.
func (pmu *PackageMetricsUpdate) Where(ps ...predicate.PackageMetrics) *PackageMetricsUpdate {
	pmu.mutation.Where(ps...)
	return pmu
}

// SetPackagesLoaded sets the "packages_loaded" field.
func (pmu *PackageMetricsUpdate) SetPackagesLoaded(i int64) *PackageMetricsUpdate {
	pmu.mutation.ResetPackagesLoaded()
	pmu.mutation.SetPackagesLoaded(i)
	return pmu
}

// SetNillablePackagesLoaded sets the "packages_loaded" field if the given value is not nil.
func (pmu *PackageMetricsUpdate) SetNillablePackagesLoaded(i *int64) *PackageMetricsUpdate {
	if i != nil {
		pmu.SetPackagesLoaded(*i)
	}
	return pmu
}

// AddPackagesLoaded adds i to the "packages_loaded" field.
func (pmu *PackageMetricsUpdate) AddPackagesLoaded(i int64) *PackageMetricsUpdate {
	pmu.mutation.AddPackagesLoaded(i)
	return pmu
}

// ClearPackagesLoaded clears the value of the "packages_loaded" field.
func (pmu *PackageMetricsUpdate) ClearPackagesLoaded() *PackageMetricsUpdate {
	pmu.mutation.ClearPackagesLoaded()
	return pmu
}

// SetMetricsID sets the "metrics" edge to the Metrics entity by ID.
func (pmu *PackageMetricsUpdate) SetMetricsID(id int) *PackageMetricsUpdate {
	pmu.mutation.SetMetricsID(id)
	return pmu
}

// SetNillableMetricsID sets the "metrics" edge to the Metrics entity by ID if the given value is not nil.
func (pmu *PackageMetricsUpdate) SetNillableMetricsID(id *int) *PackageMetricsUpdate {
	if id != nil {
		pmu = pmu.SetMetricsID(*id)
	}
	return pmu
}

// SetMetrics sets the "metrics" edge to the Metrics entity.
func (pmu *PackageMetricsUpdate) SetMetrics(m *Metrics) *PackageMetricsUpdate {
	return pmu.SetMetricsID(m.ID)
}

// AddPackageLoadMetricIDs adds the "package_load_metrics" edge to the PackageLoadMetrics entity by IDs.
func (pmu *PackageMetricsUpdate) AddPackageLoadMetricIDs(ids ...int) *PackageMetricsUpdate {
	pmu.mutation.AddPackageLoadMetricIDs(ids...)
	return pmu
}

// AddPackageLoadMetrics adds the "package_load_metrics" edges to the PackageLoadMetrics entity.
func (pmu *PackageMetricsUpdate) AddPackageLoadMetrics(p ...*PackageLoadMetrics) *PackageMetricsUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pmu.AddPackageLoadMetricIDs(ids...)
}

// Mutation returns the PackageMetricsMutation object of the builder.
func (pmu *PackageMetricsUpdate) Mutation() *PackageMetricsMutation {
	return pmu.mutation
}

// ClearMetrics clears the "metrics" edge to the Metrics entity.
func (pmu *PackageMetricsUpdate) ClearMetrics() *PackageMetricsUpdate {
	pmu.mutation.ClearMetrics()
	return pmu
}

// ClearPackageLoadMetrics clears all "package_load_metrics" edges to the PackageLoadMetrics entity.
func (pmu *PackageMetricsUpdate) ClearPackageLoadMetrics() *PackageMetricsUpdate {
	pmu.mutation.ClearPackageLoadMetrics()
	return pmu
}

// RemovePackageLoadMetricIDs removes the "package_load_metrics" edge to PackageLoadMetrics entities by IDs.
func (pmu *PackageMetricsUpdate) RemovePackageLoadMetricIDs(ids ...int) *PackageMetricsUpdate {
	pmu.mutation.RemovePackageLoadMetricIDs(ids...)
	return pmu
}

// RemovePackageLoadMetrics removes "package_load_metrics" edges to PackageLoadMetrics entities.
func (pmu *PackageMetricsUpdate) RemovePackageLoadMetrics(p ...*PackageLoadMetrics) *PackageMetricsUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pmu.RemovePackageLoadMetricIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pmu *PackageMetricsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pmu.sqlSave, pmu.mutation, pmu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pmu *PackageMetricsUpdate) SaveX(ctx context.Context) int {
	affected, err := pmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pmu *PackageMetricsUpdate) Exec(ctx context.Context) error {
	_, err := pmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pmu *PackageMetricsUpdate) ExecX(ctx context.Context) {
	if err := pmu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pmu *PackageMetricsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(packagemetrics.Table, packagemetrics.Columns, sqlgraph.NewFieldSpec(packagemetrics.FieldID, field.TypeInt))
	if ps := pmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pmu.mutation.PackagesLoaded(); ok {
		_spec.SetField(packagemetrics.FieldPackagesLoaded, field.TypeInt64, value)
	}
	if value, ok := pmu.mutation.AddedPackagesLoaded(); ok {
		_spec.AddField(packagemetrics.FieldPackagesLoaded, field.TypeInt64, value)
	}
	if pmu.mutation.PackagesLoadedCleared() {
		_spec.ClearField(packagemetrics.FieldPackagesLoaded, field.TypeInt64)
	}
	if pmu.mutation.MetricsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   packagemetrics.MetricsTable,
			Columns: []string{packagemetrics.MetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metrics.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pmu.mutation.MetricsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   packagemetrics.MetricsTable,
			Columns: []string{packagemetrics.MetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metrics.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pmu.mutation.PackageLoadMetricsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   packagemetrics.PackageLoadMetricsTable,
			Columns: []string{packagemetrics.PackageLoadMetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageloadmetrics.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pmu.mutation.RemovedPackageLoadMetricsIDs(); len(nodes) > 0 && !pmu.mutation.PackageLoadMetricsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   packagemetrics.PackageLoadMetricsTable,
			Columns: []string{packagemetrics.PackageLoadMetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageloadmetrics.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pmu.mutation.PackageLoadMetricsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   packagemetrics.PackageLoadMetricsTable,
			Columns: []string{packagemetrics.PackageLoadMetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageloadmetrics.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{packagemetrics.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pmu.mutation.done = true
	return n, nil
}

// PackageMetricsUpdateOne is the builder for updating a single PackageMetrics entity.
type PackageMetricsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PackageMetricsMutation
}

// SetPackagesLoaded sets the "packages_loaded" field.
func (pmuo *PackageMetricsUpdateOne) SetPackagesLoaded(i int64) *PackageMetricsUpdateOne {
	pmuo.mutation.ResetPackagesLoaded()
	pmuo.mutation.SetPackagesLoaded(i)
	return pmuo
}

// SetNillablePackagesLoaded sets the "packages_loaded" field if the given value is not nil.
func (pmuo *PackageMetricsUpdateOne) SetNillablePackagesLoaded(i *int64) *PackageMetricsUpdateOne {
	if i != nil {
		pmuo.SetPackagesLoaded(*i)
	}
	return pmuo
}

// AddPackagesLoaded adds i to the "packages_loaded" field.
func (pmuo *PackageMetricsUpdateOne) AddPackagesLoaded(i int64) *PackageMetricsUpdateOne {
	pmuo.mutation.AddPackagesLoaded(i)
	return pmuo
}

// ClearPackagesLoaded clears the value of the "packages_loaded" field.
func (pmuo *PackageMetricsUpdateOne) ClearPackagesLoaded() *PackageMetricsUpdateOne {
	pmuo.mutation.ClearPackagesLoaded()
	return pmuo
}

// SetMetricsID sets the "metrics" edge to the Metrics entity by ID.
func (pmuo *PackageMetricsUpdateOne) SetMetricsID(id int) *PackageMetricsUpdateOne {
	pmuo.mutation.SetMetricsID(id)
	return pmuo
}

// SetNillableMetricsID sets the "metrics" edge to the Metrics entity by ID if the given value is not nil.
func (pmuo *PackageMetricsUpdateOne) SetNillableMetricsID(id *int) *PackageMetricsUpdateOne {
	if id != nil {
		pmuo = pmuo.SetMetricsID(*id)
	}
	return pmuo
}

// SetMetrics sets the "metrics" edge to the Metrics entity.
func (pmuo *PackageMetricsUpdateOne) SetMetrics(m *Metrics) *PackageMetricsUpdateOne {
	return pmuo.SetMetricsID(m.ID)
}

// AddPackageLoadMetricIDs adds the "package_load_metrics" edge to the PackageLoadMetrics entity by IDs.
func (pmuo *PackageMetricsUpdateOne) AddPackageLoadMetricIDs(ids ...int) *PackageMetricsUpdateOne {
	pmuo.mutation.AddPackageLoadMetricIDs(ids...)
	return pmuo
}

// AddPackageLoadMetrics adds the "package_load_metrics" edges to the PackageLoadMetrics entity.
func (pmuo *PackageMetricsUpdateOne) AddPackageLoadMetrics(p ...*PackageLoadMetrics) *PackageMetricsUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pmuo.AddPackageLoadMetricIDs(ids...)
}

// Mutation returns the PackageMetricsMutation object of the builder.
func (pmuo *PackageMetricsUpdateOne) Mutation() *PackageMetricsMutation {
	return pmuo.mutation
}

// ClearMetrics clears the "metrics" edge to the Metrics entity.
func (pmuo *PackageMetricsUpdateOne) ClearMetrics() *PackageMetricsUpdateOne {
	pmuo.mutation.ClearMetrics()
	return pmuo
}

// ClearPackageLoadMetrics clears all "package_load_metrics" edges to the PackageLoadMetrics entity.
func (pmuo *PackageMetricsUpdateOne) ClearPackageLoadMetrics() *PackageMetricsUpdateOne {
	pmuo.mutation.ClearPackageLoadMetrics()
	return pmuo
}

// RemovePackageLoadMetricIDs removes the "package_load_metrics" edge to PackageLoadMetrics entities by IDs.
func (pmuo *PackageMetricsUpdateOne) RemovePackageLoadMetricIDs(ids ...int) *PackageMetricsUpdateOne {
	pmuo.mutation.RemovePackageLoadMetricIDs(ids...)
	return pmuo
}

// RemovePackageLoadMetrics removes "package_load_metrics" edges to PackageLoadMetrics entities.
func (pmuo *PackageMetricsUpdateOne) RemovePackageLoadMetrics(p ...*PackageLoadMetrics) *PackageMetricsUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pmuo.RemovePackageLoadMetricIDs(ids...)
}

// Where appends a list predicates to the PackageMetricsUpdate builder.
func (pmuo *PackageMetricsUpdateOne) Where(ps ...predicate.PackageMetrics) *PackageMetricsUpdateOne {
	pmuo.mutation.Where(ps...)
	return pmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pmuo *PackageMetricsUpdateOne) Select(field string, fields ...string) *PackageMetricsUpdateOne {
	pmuo.fields = append([]string{field}, fields...)
	return pmuo
}

// Save executes the query and returns the updated PackageMetrics entity.
func (pmuo *PackageMetricsUpdateOne) Save(ctx context.Context) (*PackageMetrics, error) {
	return withHooks(ctx, pmuo.sqlSave, pmuo.mutation, pmuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pmuo *PackageMetricsUpdateOne) SaveX(ctx context.Context) *PackageMetrics {
	node, err := pmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pmuo *PackageMetricsUpdateOne) Exec(ctx context.Context) error {
	_, err := pmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pmuo *PackageMetricsUpdateOne) ExecX(ctx context.Context) {
	if err := pmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pmuo *PackageMetricsUpdateOne) sqlSave(ctx context.Context) (_node *PackageMetrics, err error) {
	_spec := sqlgraph.NewUpdateSpec(packagemetrics.Table, packagemetrics.Columns, sqlgraph.NewFieldSpec(packagemetrics.FieldID, field.TypeInt))
	id, ok := pmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PackageMetrics.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, packagemetrics.FieldID)
		for _, f := range fields {
			if !packagemetrics.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != packagemetrics.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pmuo.mutation.PackagesLoaded(); ok {
		_spec.SetField(packagemetrics.FieldPackagesLoaded, field.TypeInt64, value)
	}
	if value, ok := pmuo.mutation.AddedPackagesLoaded(); ok {
		_spec.AddField(packagemetrics.FieldPackagesLoaded, field.TypeInt64, value)
	}
	if pmuo.mutation.PackagesLoadedCleared() {
		_spec.ClearField(packagemetrics.FieldPackagesLoaded, field.TypeInt64)
	}
	if pmuo.mutation.MetricsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   packagemetrics.MetricsTable,
			Columns: []string{packagemetrics.MetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metrics.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pmuo.mutation.MetricsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   packagemetrics.MetricsTable,
			Columns: []string{packagemetrics.MetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(metrics.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pmuo.mutation.PackageLoadMetricsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   packagemetrics.PackageLoadMetricsTable,
			Columns: []string{packagemetrics.PackageLoadMetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageloadmetrics.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pmuo.mutation.RemovedPackageLoadMetricsIDs(); len(nodes) > 0 && !pmuo.mutation.PackageLoadMetricsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   packagemetrics.PackageLoadMetricsTable,
			Columns: []string{packagemetrics.PackageLoadMetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageloadmetrics.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pmuo.mutation.PackageLoadMetricsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   packagemetrics.PackageLoadMetricsTable,
			Columns: []string{packagemetrics.PackageLoadMetricsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(packageloadmetrics.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PackageMetrics{config: pmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{packagemetrics.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pmuo.mutation.done = true
	return _node, nil
}
