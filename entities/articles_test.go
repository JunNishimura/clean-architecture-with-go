// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package entities

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

func testArticles(t *testing.T) {
	t.Parallel()

	query := Articles()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testArticlesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Article{}
	if err = randomize.Struct(seed, o, articleDBTypes, true, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
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

	count, err := Articles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testArticlesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Article{}
	if err = randomize.Struct(seed, o, articleDBTypes, true, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Articles().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Articles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testArticlesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Article{}
	if err = randomize.Struct(seed, o, articleDBTypes, true, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ArticleSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Articles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testArticlesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Article{}
	if err = randomize.Struct(seed, o, articleDBTypes, true, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ArticleExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Article exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ArticleExists to return true, but got false.")
	}
}

func testArticlesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Article{}
	if err = randomize.Struct(seed, o, articleDBTypes, true, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	articleFound, err := FindArticle(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if articleFound == nil {
		t.Error("want a record, got nil")
	}
}

func testArticlesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Article{}
	if err = randomize.Struct(seed, o, articleDBTypes, true, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Articles().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testArticlesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Article{}
	if err = randomize.Struct(seed, o, articleDBTypes, true, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Articles().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testArticlesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	articleOne := &Article{}
	articleTwo := &Article{}
	if err = randomize.Struct(seed, articleOne, articleDBTypes, false, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}
	if err = randomize.Struct(seed, articleTwo, articleDBTypes, false, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = articleOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = articleTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Articles().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testArticlesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	articleOne := &Article{}
	articleTwo := &Article{}
	if err = randomize.Struct(seed, articleOne, articleDBTypes, false, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}
	if err = randomize.Struct(seed, articleTwo, articleDBTypes, false, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = articleOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = articleTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Articles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func articleBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Article) error {
	*o = Article{}
	return nil
}

func articleAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Article) error {
	*o = Article{}
	return nil
}

func articleAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Article) error {
	*o = Article{}
	return nil
}

func articleBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Article) error {
	*o = Article{}
	return nil
}

func articleAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Article) error {
	*o = Article{}
	return nil
}

func articleBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Article) error {
	*o = Article{}
	return nil
}

func articleAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Article) error {
	*o = Article{}
	return nil
}

func articleBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Article) error {
	*o = Article{}
	return nil
}

func articleAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Article) error {
	*o = Article{}
	return nil
}

func testArticlesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Article{}
	o := &Article{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, articleDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Article object: %s", err)
	}

	AddArticleHook(boil.BeforeInsertHook, articleBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	articleBeforeInsertHooks = []ArticleHook{}

	AddArticleHook(boil.AfterInsertHook, articleAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	articleAfterInsertHooks = []ArticleHook{}

	AddArticleHook(boil.AfterSelectHook, articleAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	articleAfterSelectHooks = []ArticleHook{}

	AddArticleHook(boil.BeforeUpdateHook, articleBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	articleBeforeUpdateHooks = []ArticleHook{}

	AddArticleHook(boil.AfterUpdateHook, articleAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	articleAfterUpdateHooks = []ArticleHook{}

	AddArticleHook(boil.BeforeDeleteHook, articleBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	articleBeforeDeleteHooks = []ArticleHook{}

	AddArticleHook(boil.AfterDeleteHook, articleAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	articleAfterDeleteHooks = []ArticleHook{}

	AddArticleHook(boil.BeforeUpsertHook, articleBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	articleBeforeUpsertHooks = []ArticleHook{}

	AddArticleHook(boil.AfterUpsertHook, articleAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	articleAfterUpsertHooks = []ArticleHook{}
}

func testArticlesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Article{}
	if err = randomize.Struct(seed, o, articleDBTypes, true, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Articles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testArticlesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Article{}
	if err = randomize.Struct(seed, o, articleDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(articleColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Articles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testArticlesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Article{}
	if err = randomize.Struct(seed, o, articleDBTypes, true, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
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

func testArticlesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Article{}
	if err = randomize.Struct(seed, o, articleDBTypes, true, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ArticleSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testArticlesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Article{}
	if err = randomize.Struct(seed, o, articleDBTypes, true, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Articles().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	articleDBTypes = map[string]string{`ID`: `bigint`, `Title`: `varchar`, `Body`: `text`}
	_              = bytes.MinRead
)

func testArticlesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(articlePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(articleAllColumns) == len(articlePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Article{}
	if err = randomize.Struct(seed, o, articleDBTypes, true, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Articles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, articleDBTypes, true, articlePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testArticlesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(articleAllColumns) == len(articlePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Article{}
	if err = randomize.Struct(seed, o, articleDBTypes, true, articleColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Articles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, articleDBTypes, true, articlePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(articleAllColumns, articlePrimaryKeyColumns) {
		fields = articleAllColumns
	} else {
		fields = strmangle.SetComplement(
			articleAllColumns,
			articlePrimaryKeyColumns,
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

	slice := ArticleSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testArticlesUpsert(t *testing.T) {
	t.Parallel()

	if len(articleAllColumns) == len(articlePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLArticleUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Article{}
	if err = randomize.Struct(seed, &o, articleDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Article: %s", err)
	}

	count, err := Articles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, articleDBTypes, false, articlePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Article struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Article: %s", err)
	}

	count, err = Articles().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
