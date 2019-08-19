// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Auths", testAuths)
	t.Run("People", testPeople)
}

func TestDelete(t *testing.T) {
	t.Run("Auths", testAuthsDelete)
	t.Run("People", testPeopleDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Auths", testAuthsQueryDeleteAll)
	t.Run("People", testPeopleQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Auths", testAuthsSliceDeleteAll)
	t.Run("People", testPeopleSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Auths", testAuthsExists)
	t.Run("People", testPeopleExists)
}

func TestFind(t *testing.T) {
	t.Run("Auths", testAuthsFind)
	t.Run("People", testPeopleFind)
}

func TestBind(t *testing.T) {
	t.Run("Auths", testAuthsBind)
	t.Run("People", testPeopleBind)
}

func TestOne(t *testing.T) {
	t.Run("Auths", testAuthsOne)
	t.Run("People", testPeopleOne)
}

func TestAll(t *testing.T) {
	t.Run("Auths", testAuthsAll)
	t.Run("People", testPeopleAll)
}

func TestCount(t *testing.T) {
	t.Run("Auths", testAuthsCount)
	t.Run("People", testPeopleCount)
}

func TestHooks(t *testing.T) {
	t.Run("Auths", testAuthsHooks)
	t.Run("People", testPeopleHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Auths", testAuthsInsert)
	t.Run("Auths", testAuthsInsertWhitelist)
	t.Run("People", testPeopleInsert)
	t.Run("People", testPeopleInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("AuthToPersonUsingPerson", testAuthToOnePersonUsingPerson)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {
	t.Run("PersonToAuthUsingAuth", testPersonOneToOneAuthUsingAuth)
}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("AuthToPersonUsingAuth", testAuthToOneSetOpPersonUsingPerson)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {
	t.Run("PersonToAuthUsingAuth", testPersonOneToOneSetOpAuthUsingAuth)
}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Auths", testAuthsReload)
	t.Run("People", testPeopleReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Auths", testAuthsReloadAll)
	t.Run("People", testPeopleReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Auths", testAuthsSelect)
	t.Run("People", testPeopleSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Auths", testAuthsUpdate)
	t.Run("People", testPeopleUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Auths", testAuthsSliceUpdateAll)
	t.Run("People", testPeopleSliceUpdateAll)
}
