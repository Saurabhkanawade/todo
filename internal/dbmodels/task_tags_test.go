// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package dbmodels

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testTaskTags(t *testing.T) {
	t.Parallel()

	query := TaskTags()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTaskTagsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskTag{}
	if err = randomize.Struct(seed, o, taskTagDBTypes, true, taskTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskTag struct: %s", err)
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

	count, err := TaskTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTaskTagsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskTag{}
	if err = randomize.Struct(seed, o, taskTagDBTypes, true, taskTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := TaskTags().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TaskTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTaskTagsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskTag{}
	if err = randomize.Struct(seed, o, taskTagDBTypes, true, taskTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TaskTagSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := TaskTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTaskTagsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskTag{}
	if err = randomize.Struct(seed, o, taskTagDBTypes, true, taskTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := TaskTagExists(ctx, tx, o.TaskTagID)
	if err != nil {
		t.Errorf("Unable to check if TaskTag exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TaskTagExists to return true, but got false.")
	}
}

func testTaskTagsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskTag{}
	if err = randomize.Struct(seed, o, taskTagDBTypes, true, taskTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	taskTagFound, err := FindTaskTag(ctx, tx, o.TaskTagID)
	if err != nil {
		t.Error(err)
	}

	if taskTagFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTaskTagsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskTag{}
	if err = randomize.Struct(seed, o, taskTagDBTypes, true, taskTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = TaskTags().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testTaskTagsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskTag{}
	if err = randomize.Struct(seed, o, taskTagDBTypes, true, taskTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := TaskTags().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTaskTagsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	taskTagOne := &TaskTag{}
	taskTagTwo := &TaskTag{}
	if err = randomize.Struct(seed, taskTagOne, taskTagDBTypes, false, taskTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskTag struct: %s", err)
	}
	if err = randomize.Struct(seed, taskTagTwo, taskTagDBTypes, false, taskTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = taskTagOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = taskTagTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TaskTags().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTaskTagsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	taskTagOne := &TaskTag{}
	taskTagTwo := &TaskTag{}
	if err = randomize.Struct(seed, taskTagOne, taskTagDBTypes, false, taskTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskTag struct: %s", err)
	}
	if err = randomize.Struct(seed, taskTagTwo, taskTagDBTypes, false, taskTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = taskTagOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = taskTagTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TaskTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func taskTagBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *TaskTag) error {
	*o = TaskTag{}
	return nil
}

func taskTagAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *TaskTag) error {
	*o = TaskTag{}
	return nil
}

func taskTagAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *TaskTag) error {
	*o = TaskTag{}
	return nil
}

func taskTagBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *TaskTag) error {
	*o = TaskTag{}
	return nil
}

func taskTagAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *TaskTag) error {
	*o = TaskTag{}
	return nil
}

func taskTagBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *TaskTag) error {
	*o = TaskTag{}
	return nil
}

func taskTagAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *TaskTag) error {
	*o = TaskTag{}
	return nil
}

func taskTagBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *TaskTag) error {
	*o = TaskTag{}
	return nil
}

func taskTagAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *TaskTag) error {
	*o = TaskTag{}
	return nil
}

func testTaskTagsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &TaskTag{}
	o := &TaskTag{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, taskTagDBTypes, false); err != nil {
		t.Errorf("Unable to randomize TaskTag object: %s", err)
	}

	AddTaskTagHook(boil.BeforeInsertHook, taskTagBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	taskTagBeforeInsertHooks = []TaskTagHook{}

	AddTaskTagHook(boil.AfterInsertHook, taskTagAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	taskTagAfterInsertHooks = []TaskTagHook{}

	AddTaskTagHook(boil.AfterSelectHook, taskTagAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	taskTagAfterSelectHooks = []TaskTagHook{}

	AddTaskTagHook(boil.BeforeUpdateHook, taskTagBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	taskTagBeforeUpdateHooks = []TaskTagHook{}

	AddTaskTagHook(boil.AfterUpdateHook, taskTagAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	taskTagAfterUpdateHooks = []TaskTagHook{}

	AddTaskTagHook(boil.BeforeDeleteHook, taskTagBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	taskTagBeforeDeleteHooks = []TaskTagHook{}

	AddTaskTagHook(boil.AfterDeleteHook, taskTagAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	taskTagAfterDeleteHooks = []TaskTagHook{}

	AddTaskTagHook(boil.BeforeUpsertHook, taskTagBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	taskTagBeforeUpsertHooks = []TaskTagHook{}

	AddTaskTagHook(boil.AfterUpsertHook, taskTagAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	taskTagAfterUpsertHooks = []TaskTagHook{}
}

func testTaskTagsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskTag{}
	if err = randomize.Struct(seed, o, taskTagDBTypes, true, taskTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TaskTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTaskTagsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskTag{}
	if err = randomize.Struct(seed, o, taskTagDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TaskTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(taskTagColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := TaskTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTaskTagToOneTagUsingTag(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local TaskTag
	var foreign Tag

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, taskTagDBTypes, false, taskTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskTag struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, tagDBTypes, false, tagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Tag struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.TagID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Tag().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddTagHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Tag) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := TaskTagSlice{&local}
	if err = local.L.LoadTag(ctx, tx, false, (*[]*TaskTag)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Tag == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Tag = nil
	if err = local.L.LoadTag(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Tag == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testTaskTagToOneTaskUsingTask(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local TaskTag
	var foreign Task

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, taskTagDBTypes, false, taskTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskTag struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, taskDBTypes, false, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.TaskID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Task().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddTaskHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Task) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := TaskTagSlice{&local}
	if err = local.L.LoadTask(ctx, tx, false, (*[]*TaskTag)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Task == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Task = nil
	if err = local.L.LoadTask(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Task == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testTaskTagToOneSetOpTagUsingTag(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a TaskTag
	var b, c Tag

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, taskTagDBTypes, false, strmangle.SetComplement(taskTagPrimaryKeyColumns, taskTagColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, tagDBTypes, false, strmangle.SetComplement(tagPrimaryKeyColumns, tagColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, tagDBTypes, false, strmangle.SetComplement(tagPrimaryKeyColumns, tagColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Tag{&b, &c} {
		err = a.SetTag(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Tag != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.TaskTags[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.TagID != x.ID {
			t.Error("foreign key was wrong value", a.TagID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.TagID))
		reflect.Indirect(reflect.ValueOf(&a.TagID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.TagID != x.ID {
			t.Error("foreign key was wrong value", a.TagID, x.ID)
		}
	}
}
func testTaskTagToOneSetOpTaskUsingTask(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a TaskTag
	var b, c Task

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, taskTagDBTypes, false, strmangle.SetComplement(taskTagPrimaryKeyColumns, taskTagColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, taskDBTypes, false, strmangle.SetComplement(taskPrimaryKeyColumns, taskColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, taskDBTypes, false, strmangle.SetComplement(taskPrimaryKeyColumns, taskColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Task{&b, &c} {
		err = a.SetTask(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Task != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.TaskTags[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.TaskID != x.ID {
			t.Error("foreign key was wrong value", a.TaskID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.TaskID))
		reflect.Indirect(reflect.ValueOf(&a.TaskID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.TaskID != x.ID {
			t.Error("foreign key was wrong value", a.TaskID, x.ID)
		}
	}
}

func testTaskTagsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskTag{}
	if err = randomize.Struct(seed, o, taskTagDBTypes, true, taskTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskTag struct: %s", err)
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

func testTaskTagsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskTag{}
	if err = randomize.Struct(seed, o, taskTagDBTypes, true, taskTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := TaskTagSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testTaskTagsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &TaskTag{}
	if err = randomize.Struct(seed, o, taskTagDBTypes, true, taskTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := TaskTags().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	taskTagDBTypes = map[string]string{`TaskID`: `uuid`, `TagID`: `uuid`, `CreatedAt`: `timestamp without time zone`, `TaskTagID`: `uuid`}
	_              = bytes.MinRead
)

func testTaskTagsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(taskTagPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(taskTagAllColumns) == len(taskTagPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TaskTag{}
	if err = randomize.Struct(seed, o, taskTagDBTypes, true, taskTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TaskTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, taskTagDBTypes, true, taskTagPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TaskTag struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testTaskTagsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(taskTagAllColumns) == len(taskTagPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &TaskTag{}
	if err = randomize.Struct(seed, o, taskTagDBTypes, true, taskTagColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize TaskTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := TaskTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, taskTagDBTypes, true, taskTagPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TaskTag struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(taskTagAllColumns, taskTagPrimaryKeyColumns) {
		fields = taskTagAllColumns
	} else {
		fields = strmangle.SetComplement(
			taskTagAllColumns,
			taskTagPrimaryKeyColumns,
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

	slice := TaskTagSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testTaskTagsUpsert(t *testing.T) {
	t.Parallel()

	if len(taskTagAllColumns) == len(taskTagPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := TaskTag{}
	if err = randomize.Struct(seed, &o, taskTagDBTypes, true); err != nil {
		t.Errorf("Unable to randomize TaskTag struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TaskTag: %s", err)
	}

	count, err := TaskTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, taskTagDBTypes, false, taskTagPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize TaskTag struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert TaskTag: %s", err)
	}

	count, err = TaskTags().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
