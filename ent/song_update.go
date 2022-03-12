// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Zubayear/song-ql/ent/artist"
	"github.com/Zubayear/song-ql/ent/predicate"
	"github.com/Zubayear/song-ql/ent/song"
	"github.com/google/uuid"
)

// SongUpdate is the builder for updating Song entities.
type SongUpdate struct {
	config
	hooks    []Hook
	mutation *SongMutation
}

// Where appends a list predicates to the SongUpdate builder.
func (su *SongUpdate) Where(ps ...predicate.Song) *SongUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetTitle sets the "title" field.
func (su *SongUpdate) SetTitle(s string) *SongUpdate {
	su.mutation.SetTitle(s)
	return su
}

// SetDuration sets the "duration" field.
func (su *SongUpdate) SetDuration(f float64) *SongUpdate {
	su.mutation.ResetDuration()
	su.mutation.SetDuration(f)
	return su
}

// AddDuration adds f to the "duration" field.
func (su *SongUpdate) AddDuration(f float64) *SongUpdate {
	su.mutation.AddDuration(f)
	return su
}

// SetLyricsExits sets the "lyricsExits" field.
func (su *SongUpdate) SetLyricsExits(b bool) *SongUpdate {
	su.mutation.SetLyricsExits(b)
	return su
}

// SetNillableLyricsExits sets the "lyricsExits" field if the given value is not nil.
func (su *SongUpdate) SetNillableLyricsExits(b *bool) *SongUpdate {
	if b != nil {
		su.SetLyricsExits(*b)
	}
	return su
}

// AddArtistIDs adds the "artists" edge to the Artist entity by IDs.
func (su *SongUpdate) AddArtistIDs(ids ...uuid.UUID) *SongUpdate {
	su.mutation.AddArtistIDs(ids...)
	return su
}

// AddArtists adds the "artists" edges to the Artist entity.
func (su *SongUpdate) AddArtists(a ...*Artist) *SongUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return su.AddArtistIDs(ids...)
}

// Mutation returns the SongMutation object of the builder.
func (su *SongUpdate) Mutation() *SongMutation {
	return su.mutation
}

// ClearArtists clears all "artists" edges to the Artist entity.
func (su *SongUpdate) ClearArtists() *SongUpdate {
	su.mutation.ClearArtists()
	return su
}

// RemoveArtistIDs removes the "artists" edge to Artist entities by IDs.
func (su *SongUpdate) RemoveArtistIDs(ids ...uuid.UUID) *SongUpdate {
	su.mutation.RemoveArtistIDs(ids...)
	return su
}

// RemoveArtists removes "artists" edges to Artist entities.
func (su *SongUpdate) RemoveArtists(a ...*Artist) *SongUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return su.RemoveArtistIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SongUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SongMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SongUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SongUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SongUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SongUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   song.Table,
			Columns: song.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: song.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldTitle,
		})
	}
	if value, ok := su.mutation.Duration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: song.FieldDuration,
		})
	}
	if value, ok := su.mutation.AddedDuration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: song.FieldDuration,
		})
	}
	if value, ok := su.mutation.LyricsExits(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: song.FieldLyricsExits,
		})
	}
	if su.mutation.ArtistsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   song.ArtistsTable,
			Columns: []string{song.ArtistsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: artist.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedArtistsIDs(); len(nodes) > 0 && !su.mutation.ArtistsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   song.ArtistsTable,
			Columns: []string{song.ArtistsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: artist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ArtistsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   song.ArtistsTable,
			Columns: []string{song.ArtistsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: artist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{song.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// SongUpdateOne is the builder for updating a single Song entity.
type SongUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SongMutation
}

// SetTitle sets the "title" field.
func (suo *SongUpdateOne) SetTitle(s string) *SongUpdateOne {
	suo.mutation.SetTitle(s)
	return suo
}

// SetDuration sets the "duration" field.
func (suo *SongUpdateOne) SetDuration(f float64) *SongUpdateOne {
	suo.mutation.ResetDuration()
	suo.mutation.SetDuration(f)
	return suo
}

// AddDuration adds f to the "duration" field.
func (suo *SongUpdateOne) AddDuration(f float64) *SongUpdateOne {
	suo.mutation.AddDuration(f)
	return suo
}

// SetLyricsExits sets the "lyricsExits" field.
func (suo *SongUpdateOne) SetLyricsExits(b bool) *SongUpdateOne {
	suo.mutation.SetLyricsExits(b)
	return suo
}

// SetNillableLyricsExits sets the "lyricsExits" field if the given value is not nil.
func (suo *SongUpdateOne) SetNillableLyricsExits(b *bool) *SongUpdateOne {
	if b != nil {
		suo.SetLyricsExits(*b)
	}
	return suo
}

// AddArtistIDs adds the "artists" edge to the Artist entity by IDs.
func (suo *SongUpdateOne) AddArtistIDs(ids ...uuid.UUID) *SongUpdateOne {
	suo.mutation.AddArtistIDs(ids...)
	return suo
}

// AddArtists adds the "artists" edges to the Artist entity.
func (suo *SongUpdateOne) AddArtists(a ...*Artist) *SongUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return suo.AddArtistIDs(ids...)
}

// Mutation returns the SongMutation object of the builder.
func (suo *SongUpdateOne) Mutation() *SongMutation {
	return suo.mutation
}

// ClearArtists clears all "artists" edges to the Artist entity.
func (suo *SongUpdateOne) ClearArtists() *SongUpdateOne {
	suo.mutation.ClearArtists()
	return suo
}

// RemoveArtistIDs removes the "artists" edge to Artist entities by IDs.
func (suo *SongUpdateOne) RemoveArtistIDs(ids ...uuid.UUID) *SongUpdateOne {
	suo.mutation.RemoveArtistIDs(ids...)
	return suo
}

// RemoveArtists removes "artists" edges to Artist entities.
func (suo *SongUpdateOne) RemoveArtists(a ...*Artist) *SongUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return suo.RemoveArtistIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SongUpdateOne) Select(field string, fields ...string) *SongUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Song entity.
func (suo *SongUpdateOne) Save(ctx context.Context) (*Song, error) {
	var (
		err  error
		node *Song
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SongMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SongUpdateOne) SaveX(ctx context.Context) *Song {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SongUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SongUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SongUpdateOne) sqlSave(ctx context.Context) (_node *Song, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   song.Table,
			Columns: song.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: song.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Song.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, song.FieldID)
		for _, f := range fields {
			if !song.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != song.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: song.FieldTitle,
		})
	}
	if value, ok := suo.mutation.Duration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: song.FieldDuration,
		})
	}
	if value, ok := suo.mutation.AddedDuration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: song.FieldDuration,
		})
	}
	if value, ok := suo.mutation.LyricsExits(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: song.FieldLyricsExits,
		})
	}
	if suo.mutation.ArtistsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   song.ArtistsTable,
			Columns: []string{song.ArtistsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: artist.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedArtistsIDs(); len(nodes) > 0 && !suo.mutation.ArtistsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   song.ArtistsTable,
			Columns: []string{song.ArtistsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: artist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ArtistsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   song.ArtistsTable,
			Columns: []string{song.ArtistsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: artist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Song{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{song.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}