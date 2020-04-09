// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testWESEndpoints(t *testing.T) {
	t.Parallel()

	query := WESEndpoints()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testWESEndpointsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WESEndpoint{}
	if err = randomize.Struct(seed, o, wESEndpointDBTypes, true, wESEndpointColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WESEndpoint struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := WESEndpoints().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWESEndpointsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WESEndpoint{}
	if err = randomize.Struct(seed, o, wESEndpointDBTypes, true, wESEndpointColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WESEndpoint struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := WESEndpoints().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := WESEndpoints().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWESEndpointsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WESEndpoint{}
	if err = randomize.Struct(seed, o, wESEndpointDBTypes, true, wESEndpointColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WESEndpoint struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := WESEndpointSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := WESEndpoints().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testWESEndpointsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WESEndpoint{}
	if err = randomize.Struct(seed, o, wESEndpointDBTypes, true, wESEndpointColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WESEndpoint struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := WESEndpointExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if WESEndpoint exists: %s", err)
	}
	if !e {
		t.Errorf("Expected WESEndpointExists to return true, but got false.")
	}
}

func testWESEndpointsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WESEndpoint{}
	if err = randomize.Struct(seed, o, wESEndpointDBTypes, true, wESEndpointColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WESEndpoint struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	wESEndpointFound, err := FindWESEndpoint(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if wESEndpointFound == nil {
		t.Error("want a record, got nil")
	}
}

func testWESEndpointsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WESEndpoint{}
	if err = randomize.Struct(seed, o, wESEndpointDBTypes, true, wESEndpointColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WESEndpoint struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = WESEndpoints().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testWESEndpointsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WESEndpoint{}
	if err = randomize.Struct(seed, o, wESEndpointDBTypes, true, wESEndpointColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WESEndpoint struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := WESEndpoints().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testWESEndpointsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	wESEndpointOne := &WESEndpoint{}
	wESEndpointTwo := &WESEndpoint{}
	if err = randomize.Struct(seed, wESEndpointOne, wESEndpointDBTypes, false, wESEndpointColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WESEndpoint struct: %s", err)
	}
	if err = randomize.Struct(seed, wESEndpointTwo, wESEndpointDBTypes, false, wESEndpointColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WESEndpoint struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = wESEndpointOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = wESEndpointTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := WESEndpoints().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testWESEndpointsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	wESEndpointOne := &WESEndpoint{}
	wESEndpointTwo := &WESEndpoint{}
	if err = randomize.Struct(seed, wESEndpointOne, wESEndpointDBTypes, false, wESEndpointColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WESEndpoint struct: %s", err)
	}
	if err = randomize.Struct(seed, wESEndpointTwo, wESEndpointDBTypes, false, wESEndpointColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WESEndpoint struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = wESEndpointOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = wESEndpointTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := WESEndpoints().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func wESEndpointBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *WESEndpoint) error {
	*o = WESEndpoint{}
	return nil
}

func wESEndpointAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *WESEndpoint) error {
	*o = WESEndpoint{}
	return nil
}

func wESEndpointAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *WESEndpoint) error {
	*o = WESEndpoint{}
	return nil
}

func wESEndpointBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *WESEndpoint) error {
	*o = WESEndpoint{}
	return nil
}

func wESEndpointAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *WESEndpoint) error {
	*o = WESEndpoint{}
	return nil
}

func wESEndpointBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *WESEndpoint) error {
	*o = WESEndpoint{}
	return nil
}

func wESEndpointAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *WESEndpoint) error {
	*o = WESEndpoint{}
	return nil
}

func wESEndpointBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *WESEndpoint) error {
	*o = WESEndpoint{}
	return nil
}

func wESEndpointAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *WESEndpoint) error {
	*o = WESEndpoint{}
	return nil
}

func testWESEndpointsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &WESEndpoint{}
	o := &WESEndpoint{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, wESEndpointDBTypes, false); err != nil {
		t.Errorf("Unable to randomize WESEndpoint object: %s", err)
	}

	AddWESEndpointHook(boil.BeforeInsertHook, wESEndpointBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	wESEndpointBeforeInsertHooks = []WESEndpointHook{}

	AddWESEndpointHook(boil.AfterInsertHook, wESEndpointAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	wESEndpointAfterInsertHooks = []WESEndpointHook{}

	AddWESEndpointHook(boil.AfterSelectHook, wESEndpointAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	wESEndpointAfterSelectHooks = []WESEndpointHook{}

	AddWESEndpointHook(boil.BeforeUpdateHook, wESEndpointBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	wESEndpointBeforeUpdateHooks = []WESEndpointHook{}

	AddWESEndpointHook(boil.AfterUpdateHook, wESEndpointAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	wESEndpointAfterUpdateHooks = []WESEndpointHook{}

	AddWESEndpointHook(boil.BeforeDeleteHook, wESEndpointBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	wESEndpointBeforeDeleteHooks = []WESEndpointHook{}

	AddWESEndpointHook(boil.AfterDeleteHook, wESEndpointAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	wESEndpointAfterDeleteHooks = []WESEndpointHook{}

	AddWESEndpointHook(boil.BeforeUpsertHook, wESEndpointBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	wESEndpointBeforeUpsertHooks = []WESEndpointHook{}

	AddWESEndpointHook(boil.AfterUpsertHook, wESEndpointAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	wESEndpointAfterUpsertHooks = []WESEndpointHook{}
}

func testWESEndpointsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WESEndpoint{}
	if err = randomize.Struct(seed, o, wESEndpointDBTypes, true, wESEndpointColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WESEndpoint struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := WESEndpoints().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testWESEndpointsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WESEndpoint{}
	if err = randomize.Struct(seed, o, wESEndpointDBTypes, true); err != nil {
		t.Errorf("Unable to randomize WESEndpoint struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(wESEndpointColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := WESEndpoints().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testWESEndpointToManyWESEndpointWorkflowruns(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a WESEndpoint
	var b, c Workflowrun

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, wESEndpointDBTypes, true, wESEndpointColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WESEndpoint struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, workflowrunDBTypes, false, workflowrunColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, workflowrunDBTypes, false, workflowrunColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.WESEndpointID, a.ID)
	queries.Assign(&c.WESEndpointID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.WESEndpointWorkflowruns().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.WESEndpointID, b.WESEndpointID) {
			bFound = true
		}
		if queries.Equal(v.WESEndpointID, c.WESEndpointID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := WESEndpointSlice{&a}
	if err = a.L.LoadWESEndpointWorkflowruns(ctx, tx, false, (*[]*WESEndpoint)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.WESEndpointWorkflowruns); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.WESEndpointWorkflowruns = nil
	if err = a.L.LoadWESEndpointWorkflowruns(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.WESEndpointWorkflowruns); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testWESEndpointToManyAddOpWESEndpointWorkflowruns(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a WESEndpoint
	var b, c, d, e Workflowrun

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, wESEndpointDBTypes, false, strmangle.SetComplement(wESEndpointPrimaryKeyColumns, wESEndpointColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Workflowrun{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, workflowrunDBTypes, false, strmangle.SetComplement(workflowrunPrimaryKeyColumns, workflowrunColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Workflowrun{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddWESEndpointWorkflowruns(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.WESEndpointID) {
			t.Error("foreign key was wrong value", a.ID, first.WESEndpointID)
		}
		if !queries.Equal(a.ID, second.WESEndpointID) {
			t.Error("foreign key was wrong value", a.ID, second.WESEndpointID)
		}

		if first.R.WESEndpoint != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.WESEndpoint != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.WESEndpointWorkflowruns[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.WESEndpointWorkflowruns[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.WESEndpointWorkflowruns().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testWESEndpointToManySetOpWESEndpointWorkflowruns(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a WESEndpoint
	var b, c, d, e Workflowrun

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, wESEndpointDBTypes, false, strmangle.SetComplement(wESEndpointPrimaryKeyColumns, wESEndpointColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Workflowrun{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, workflowrunDBTypes, false, strmangle.SetComplement(workflowrunPrimaryKeyColumns, workflowrunColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetWESEndpointWorkflowruns(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.WESEndpointWorkflowruns().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetWESEndpointWorkflowruns(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.WESEndpointWorkflowruns().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.WESEndpointID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.WESEndpointID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.WESEndpointID) {
		t.Error("foreign key was wrong value", a.ID, d.WESEndpointID)
	}
	if !queries.Equal(a.ID, e.WESEndpointID) {
		t.Error("foreign key was wrong value", a.ID, e.WESEndpointID)
	}

	if b.R.WESEndpoint != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.WESEndpoint != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.WESEndpoint != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.WESEndpoint != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.WESEndpointWorkflowruns[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.WESEndpointWorkflowruns[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testWESEndpointToManyRemoveOpWESEndpointWorkflowruns(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a WESEndpoint
	var b, c, d, e Workflowrun

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, wESEndpointDBTypes, false, strmangle.SetComplement(wESEndpointPrimaryKeyColumns, wESEndpointColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Workflowrun{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, workflowrunDBTypes, false, strmangle.SetComplement(workflowrunPrimaryKeyColumns, workflowrunColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddWESEndpointWorkflowruns(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.WESEndpointWorkflowruns().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveWESEndpointWorkflowruns(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.WESEndpointWorkflowruns().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.WESEndpointID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.WESEndpointID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.WESEndpoint != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.WESEndpoint != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.WESEndpoint != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.WESEndpoint != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.WESEndpointWorkflowruns) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.WESEndpointWorkflowruns[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.WESEndpointWorkflowruns[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testWESEndpointToOneProjectUsingOwner(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local WESEndpoint
	var foreign Project

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, wESEndpointDBTypes, false, wESEndpointColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WESEndpoint struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, projectDBTypes, false, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.OwnerID = foreign.ProjectID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Owner().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ProjectID != foreign.ProjectID {
		t.Errorf("want: %v, got %v", foreign.ProjectID, check.ProjectID)
	}

	slice := WESEndpointSlice{&local}
	if err = local.L.LoadOwner(ctx, tx, false, (*[]*WESEndpoint)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Owner == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Owner = nil
	if err = local.L.LoadOwner(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Owner == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testWESEndpointToOneSetOpProjectUsingOwner(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a WESEndpoint
	var b, c Project

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, wESEndpointDBTypes, false, strmangle.SetComplement(wESEndpointPrimaryKeyColumns, wESEndpointColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, projectDBTypes, false, strmangle.SetComplement(projectPrimaryKeyColumns, projectColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, projectDBTypes, false, strmangle.SetComplement(projectPrimaryKeyColumns, projectColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Project{&b, &c} {
		err = a.SetOwner(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Owner != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.OwnerWESEndpoints[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.OwnerID != x.ProjectID {
			t.Error("foreign key was wrong value", a.OwnerID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.OwnerID))
		reflect.Indirect(reflect.ValueOf(&a.OwnerID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.OwnerID != x.ProjectID {
			t.Error("foreign key was wrong value", a.OwnerID, x.ProjectID)
		}
	}
}

func testWESEndpointsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WESEndpoint{}
	if err = randomize.Struct(seed, o, wESEndpointDBTypes, true, wESEndpointColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WESEndpoint struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testWESEndpointsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WESEndpoint{}
	if err = randomize.Struct(seed, o, wESEndpointDBTypes, true, wESEndpointColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WESEndpoint struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := WESEndpointSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testWESEndpointsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &WESEndpoint{}
	if err = randomize.Struct(seed, o, wESEndpointDBTypes, true, wESEndpointColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WESEndpoint struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := WESEndpoints().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	wESEndpointDBTypes = map[string]string{`ID`: `int`, `Name`: `varchar`, `Link`: `varchar`, `OwnerID`: `varchar`}
	_                  = bytes.MinRead
)

func testWESEndpointsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(wESEndpointPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(wESEndpointAllColumns) == len(wESEndpointPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &WESEndpoint{}
	if err = randomize.Struct(seed, o, wESEndpointDBTypes, true, wESEndpointColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WESEndpoint struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := WESEndpoints().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, wESEndpointDBTypes, true, wESEndpointPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize WESEndpoint struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testWESEndpointsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(wESEndpointAllColumns) == len(wESEndpointPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &WESEndpoint{}
	if err = randomize.Struct(seed, o, wESEndpointDBTypes, true, wESEndpointColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize WESEndpoint struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := WESEndpoints().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, wESEndpointDBTypes, true, wESEndpointPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize WESEndpoint struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(wESEndpointAllColumns, wESEndpointPrimaryKeyColumns) {
		fields = wESEndpointAllColumns
	} else {
		fields = strmangle.SetComplement(
			wESEndpointAllColumns,
			wESEndpointPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := WESEndpointSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testWESEndpointsUpsert(t *testing.T) {
	t.Parallel()

	if len(wESEndpointAllColumns) == len(wESEndpointPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLWESEndpointUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := WESEndpoint{}
	if err = randomize.Struct(seed, &o, wESEndpointDBTypes, false); err != nil {
		t.Errorf("Unable to randomize WESEndpoint struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert WESEndpoint: %s", err)
	}

	count, err := WESEndpoints().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, wESEndpointDBTypes, false, wESEndpointPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize WESEndpoint struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert WESEndpoint: %s", err)
	}

	count, err = WESEndpoints().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
