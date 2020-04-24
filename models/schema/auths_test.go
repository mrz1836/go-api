// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package schema

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

func testAuths(t *testing.T) {
	t.Parallel()

	query := Auths()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testAuthsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Auth{}
	if err = randomize.Struct(seed, o, authDBTypes, true, authColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Auth struct: %s", err)
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

	count, err := Auths().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Auth{}
	if err = randomize.Struct(seed, o, authDBTypes, true, authColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Auth struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Auths().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Auths().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Auth{}
	if err = randomize.Struct(seed, o, authDBTypes, true, authColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Auth struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AuthSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Auths().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAuthsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Auth{}
	if err = randomize.Struct(seed, o, authDBTypes, true, authColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Auth struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := AuthExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Auth exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AuthExists to return true, but got false.")
	}
}

func testAuthsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Auth{}
	if err = randomize.Struct(seed, o, authDBTypes, true, authColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Auth struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	authFound, err := FindAuth(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if authFound == nil {
		t.Error("want a record, got nil")
	}
}

func testAuthsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Auth{}
	if err = randomize.Struct(seed, o, authDBTypes, true, authColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Auth struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Auths().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testAuthsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Auth{}
	if err = randomize.Struct(seed, o, authDBTypes, true, authColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Auth struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Auths().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAuthsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	authOne := &Auth{}
	authTwo := &Auth{}
	if err = randomize.Struct(seed, authOne, authDBTypes, false, authColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Auth struct: %s", err)
	}
	if err = randomize.Struct(seed, authTwo, authDBTypes, false, authColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Auth struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = authOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = authTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Auths().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAuthsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	authOne := &Auth{}
	authTwo := &Auth{}
	if err = randomize.Struct(seed, authOne, authDBTypes, false, authColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Auth struct: %s", err)
	}
	if err = randomize.Struct(seed, authTwo, authDBTypes, false, authColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Auth struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = authOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = authTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Auths().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func authBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Auth) error {
	*o = Auth{}
	return nil
}

func authAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Auth) error {
	*o = Auth{}
	return nil
}

func authAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Auth) error {
	*o = Auth{}
	return nil
}

func authBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Auth) error {
	*o = Auth{}
	return nil
}

func authAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Auth) error {
	*o = Auth{}
	return nil
}

func authBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Auth) error {
	*o = Auth{}
	return nil
}

func authAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Auth) error {
	*o = Auth{}
	return nil
}

func authBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Auth) error {
	*o = Auth{}
	return nil
}

func authAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Auth) error {
	*o = Auth{}
	return nil
}

func testAuthsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Auth{}
	o := &Auth{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, authDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Auth object: %s", err)
	}

	AddAuthHook(boil.BeforeInsertHook, authBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	authBeforeInsertHooks = []AuthHook{}

	AddAuthHook(boil.AfterInsertHook, authAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	authAfterInsertHooks = []AuthHook{}

	AddAuthHook(boil.AfterSelectHook, authAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	authAfterSelectHooks = []AuthHook{}

	AddAuthHook(boil.BeforeUpdateHook, authBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	authBeforeUpdateHooks = []AuthHook{}

	AddAuthHook(boil.AfterUpdateHook, authAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	authAfterUpdateHooks = []AuthHook{}

	AddAuthHook(boil.BeforeDeleteHook, authBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	authBeforeDeleteHooks = []AuthHook{}

	AddAuthHook(boil.AfterDeleteHook, authAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	authAfterDeleteHooks = []AuthHook{}

	AddAuthHook(boil.BeforeUpsertHook, authBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	authBeforeUpsertHooks = []AuthHook{}

	AddAuthHook(boil.AfterUpsertHook, authAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	authAfterUpsertHooks = []AuthHook{}
}

func testAuthsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Auth{}
	if err = randomize.Struct(seed, o, authDBTypes, true, authColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Auth struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Auths().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAuthsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Auth{}
	if err = randomize.Struct(seed, o, authDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Auth struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(authColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Auths().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAuthToOnePersonUsingPerson(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Auth
	var foreign Person

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, authDBTypes, false, authColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Auth struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, personDBTypes, false, personColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Person struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.PersonID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Person().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := AuthSlice{&local}
	if err = local.L.LoadPerson(ctx, tx, false, (*[]*Auth)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Person == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Person = nil
	if err = local.L.LoadPerson(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Person == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testAuthToOneSetOpPersonUsingPerson(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Auth
	var b, c Person

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, authDBTypes, false, strmangle.SetComplement(authPrimaryKeyColumns, authColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, personDBTypes, false, strmangle.SetComplement(personPrimaryKeyColumns, personColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, personDBTypes, false, strmangle.SetComplement(personPrimaryKeyColumns, personColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Person{&b, &c} {
		err = a.SetPerson(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Person != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Auth != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.PersonID != x.ID {
			t.Error("foreign key was wrong value", a.PersonID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.PersonID))
		reflect.Indirect(reflect.ValueOf(&a.PersonID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PersonID != x.ID {
			t.Error("foreign key was wrong value", a.PersonID, x.ID)
		}
	}
}

func testAuthsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Auth{}
	if err = randomize.Struct(seed, o, authDBTypes, true, authColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Auth struct: %s", err)
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

func testAuthsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Auth{}
	if err = randomize.Struct(seed, o, authDBTypes, true, authColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Auth struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := AuthSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testAuthsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Auth{}
	if err = randomize.Struct(seed, o, authDBTypes, true, authColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Auth struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Auths().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	authDBTypes = map[string]string{`ID`: `bigint`, `PersonID`: `bigint`, `CreatedAt`: `timestamp`, `ModifiedAt`: `timestamp`, `PasswordDigest`: `varchar`, `YubikeyDigest`: `varchar`, `YubikeyBackupDigest`: `varchar`, `Email`: `varchar`, `EmailConfirmToken`: `char`, `EmailConfirmed`: `tinyint`, `EmailConfirmTime`: `timestamp`, `LastIPAddress`: `varchar`, `LastLoginAt`: `timestamp`, `LastUserAgent`: `varchar`, `LoginCount`: `int`, `ResetForce`: `tinyint`, `ResetPasswordTime`: `timestamp`, `ResetPasswordToken`: `char`, `ResetTokenExpiresAt`: `timestamp`, `Locked`: `tinyint`, `LockedTime`: `timestamp`, `LockedByUserID`: `bigint`, `IsDeleted`: `tinyint`}
	_           = bytes.MinRead
)

func testAuthsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(authPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(authAllColumns) == len(authPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Auth{}
	if err = randomize.Struct(seed, o, authDBTypes, true, authColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Auth struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Auths().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, authDBTypes, true, authPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Auth struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testAuthsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(authAllColumns) == len(authPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Auth{}
	if err = randomize.Struct(seed, o, authDBTypes, true, authColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Auth struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Auths().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, authDBTypes, true, authPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Auth struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(authAllColumns, authPrimaryKeyColumns) {
		fields = authAllColumns
	} else {
		fields = strmangle.SetComplement(
			authAllColumns,
			authPrimaryKeyColumns,
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

	slice := AuthSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testAuthsUpsert(t *testing.T) {
	t.Parallel()

	if len(authAllColumns) == len(authPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLAuthUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Auth{}
	if err = randomize.Struct(seed, &o, authDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Auth struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Auth: %s", err)
	}

	count, err := Auths().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, authDBTypes, false, authPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Auth struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Auth: %s", err)
	}

	count, err = Auths().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
