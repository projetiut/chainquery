// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

func testOutputs(t *testing.T) {
	t.Parallel()

	query := Outputs(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testOutputsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	output := &Output{}
	if err = randomize.Struct(seed, output, outputDBTypes, true, outputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Output struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = output.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = output.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Outputs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOutputsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	output := &Output{}
	if err = randomize.Struct(seed, output, outputDBTypes, true, outputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Output struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = output.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Outputs(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Outputs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOutputsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	output := &Output{}
	if err = randomize.Struct(seed, output, outputDBTypes, true, outputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Output struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = output.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := OutputSlice{output}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Outputs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testOutputsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	output := &Output{}
	if err = randomize.Struct(seed, output, outputDBTypes, true, outputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Output struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = output.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := OutputExists(tx, output.ID)
	if err != nil {
		t.Errorf("Unable to check if Output exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OutputExistsG to return true, but got false.")
	}
}
func testOutputsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	output := &Output{}
	if err = randomize.Struct(seed, output, outputDBTypes, true, outputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Output struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = output.Insert(tx); err != nil {
		t.Error(err)
	}

	outputFound, err := FindOutput(tx, output.ID)
	if err != nil {
		t.Error(err)
	}

	if outputFound == nil {
		t.Error("want a record, got nil")
	}
}
func testOutputsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	output := &Output{}
	if err = randomize.Struct(seed, output, outputDBTypes, true, outputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Output struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = output.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Outputs(tx).Bind(output); err != nil {
		t.Error(err)
	}
}

func testOutputsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	output := &Output{}
	if err = randomize.Struct(seed, output, outputDBTypes, true, outputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Output struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = output.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Outputs(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOutputsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	outputOne := &Output{}
	outputTwo := &Output{}
	if err = randomize.Struct(seed, outputOne, outputDBTypes, false, outputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Output struct: %s", err)
	}
	if err = randomize.Struct(seed, outputTwo, outputDBTypes, false, outputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Output struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = outputOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = outputTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Outputs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOutputsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	outputOne := &Output{}
	outputTwo := &Output{}
	if err = randomize.Struct(seed, outputOne, outputDBTypes, false, outputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Output struct: %s", err)
	}
	if err = randomize.Struct(seed, outputTwo, outputDBTypes, false, outputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Output struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = outputOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = outputTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Outputs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testOutputsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	output := &Output{}
	if err = randomize.Struct(seed, output, outputDBTypes, true, outputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Output struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = output.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Outputs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOutputsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	output := &Output{}
	if err = randomize.Struct(seed, output, outputDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Output struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = output.Insert(tx, outputColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Outputs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOutputToManyAddresses(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Output
	var b, c Address

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, outputDBTypes, true, outputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Output struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, addressDBTypes, false, addressColumnsWithDefault...)
	randomize.Struct(seed, &c, addressDBTypes, false, addressColumnsWithDefault...)

	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	_, err = tx.Exec("insert into `output_addresses` (`output_id`, `address_id`) values (?, ?)", a.ID, b.ID)
	if err != nil {
		t.Fatal(err)
	}
	_, err = tx.Exec("insert into `output_addresses` (`output_id`, `address_id`) values (?, ?)", a.ID, c.ID)
	if err != nil {
		t.Fatal(err)
	}

	address, err := a.Addresses(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range address {
		if v.ID == b.ID {
			bFound = true
		}
		if v.ID == c.ID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := OutputSlice{&a}
	if err = a.L.LoadAddresses(tx, false, (*[]*Output)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Addresses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Addresses = nil
	if err = a.L.LoadAddresses(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Addresses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", address)
	}
}

func testOutputToManyAddOpAddresses(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Output
	var b, c, d, e Address

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, outputDBTypes, false, strmangle.SetComplement(outputPrimaryKeyColumns, outputColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Address{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, addressDBTypes, false, strmangle.SetComplement(addressPrimaryKeyColumns, addressColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Address{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddAddresses(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if first.R.Outputs[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}
		if second.R.Outputs[0] != &a {
			t.Error("relationship was not added properly to the slice")
		}

		if a.R.Addresses[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Addresses[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Addresses(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testOutputToManySetOpAddresses(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Output
	var b, c, d, e Address

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, outputDBTypes, false, strmangle.SetComplement(outputPrimaryKeyColumns, outputColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Address{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, addressDBTypes, false, strmangle.SetComplement(addressPrimaryKeyColumns, addressColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.SetAddresses(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Addresses(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetAddresses(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Addresses(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	// The following checks cannot be implemented since we have no handle
	// to these when we call Set(). Leaving them here as wishful thinking
	// and to let people know there's dragons.
	//
	// if len(b.R.Outputs) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	// if len(c.R.Outputs) != 0 {
	// 	t.Error("relationship was not removed properly from the slice")
	// }
	if d.R.Outputs[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}
	if e.R.Outputs[0] != &a {
		t.Error("relationship was not added properly to the slice")
	}

	if a.R.Addresses[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Addresses[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testOutputToManyRemoveOpAddresses(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Output
	var b, c, d, e Address

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, outputDBTypes, false, strmangle.SetComplement(outputPrimaryKeyColumns, outputColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Address{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, addressDBTypes, false, strmangle.SetComplement(addressPrimaryKeyColumns, addressColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.AddAddresses(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Addresses(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveAddresses(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Addresses(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if len(b.R.Outputs) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if len(c.R.Outputs) != 0 {
		t.Error("relationship was not removed properly from the slice")
	}
	if d.R.Outputs[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Outputs[0] != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if len(a.R.Addresses) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Addresses[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Addresses[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testOutputToOneTransactionUsingTransaction(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Output
	var foreign Transaction

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, outputDBTypes, false, outputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Output struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, transactionDBTypes, false, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.TransactionID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Transaction(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := OutputSlice{&local}
	if err = local.L.LoadTransaction(tx, false, (*[]*Output)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.Transaction == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Transaction = nil
	if err = local.L.LoadTransaction(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Transaction == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testOutputToOneInputUsingSpentByInput(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Output
	var foreign Input

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, outputDBTypes, true, outputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Output struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, inputDBTypes, false, inputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Input struct: %s", err)
	}

	local.SpentByInputID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.SpentByInputID.Uint64 = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.SpentByInput(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := OutputSlice{&local}
	if err = local.L.LoadSpentByInput(tx, false, (*[]*Output)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.SpentByInput == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.SpentByInput = nil
	if err = local.L.LoadSpentByInput(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.SpentByInput == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testOutputToOneSetOpTransactionUsingTransaction(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Output
	var b, c Transaction

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, outputDBTypes, false, strmangle.SetComplement(outputPrimaryKeyColumns, outputColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, transactionDBTypes, false, strmangle.SetComplement(transactionPrimaryKeyColumns, transactionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, transactionDBTypes, false, strmangle.SetComplement(transactionPrimaryKeyColumns, transactionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Transaction{&b, &c} {
		err = a.SetTransaction(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Transaction != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Outputs[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.TransactionID != x.ID {
			t.Error("foreign key was wrong value", a.TransactionID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.TransactionID))
		reflect.Indirect(reflect.ValueOf(&a.TransactionID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.TransactionID != x.ID {
			t.Error("foreign key was wrong value", a.TransactionID, x.ID)
		}
	}
}
func testOutputToOneSetOpInputUsingSpentByInput(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Output
	var b, c Input

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, outputDBTypes, false, strmangle.SetComplement(outputPrimaryKeyColumns, outputColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, inputDBTypes, false, strmangle.SetComplement(inputPrimaryKeyColumns, inputColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, inputDBTypes, false, strmangle.SetComplement(inputPrimaryKeyColumns, inputColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Input{&b, &c} {
		err = a.SetSpentByInput(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.SpentByInput != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.SpentByInputOutputs[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.SpentByInputID.Uint64 != x.ID {
			t.Error("foreign key was wrong value", a.SpentByInputID.Uint64)
		}

		zero := reflect.Zero(reflect.TypeOf(a.SpentByInputID.Uint64))
		reflect.Indirect(reflect.ValueOf(&a.SpentByInputID.Uint64)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.SpentByInputID.Uint64 != x.ID {
			t.Error("foreign key was wrong value", a.SpentByInputID.Uint64, x.ID)
		}
	}
}

func testOutputToOneRemoveOpInputUsingSpentByInput(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Output
	var b Input

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, outputDBTypes, false, strmangle.SetComplement(outputPrimaryKeyColumns, outputColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, inputDBTypes, false, strmangle.SetComplement(inputPrimaryKeyColumns, inputColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetSpentByInput(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveSpentByInput(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.SpentByInput(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.SpentByInput != nil {
		t.Error("R struct entry should be nil")
	}

	if a.SpentByInputID.Valid {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.SpentByInputOutputs) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testOutputsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	output := &Output{}
	if err = randomize.Struct(seed, output, outputDBTypes, true, outputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Output struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = output.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = output.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testOutputsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	output := &Output{}
	if err = randomize.Struct(seed, output, outputDBTypes, true, outputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Output struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = output.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := OutputSlice{output}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testOutputsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	output := &Output{}
	if err = randomize.Struct(seed, output, outputDBTypes, true, outputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Output struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = output.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Outputs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	outputDBTypes = map[string]string{`AddressList`: `text`, `Created`: `datetime`, `Hash160`: `varchar`, `ID`: `bigint`, `IsSpent`: `tinyint`, `Modified`: `datetime`, `RequiredSignatures`: `int`, `ScriptPubKeyAsm`: `text`, `ScriptPubKeyHex`: `text`, `SpentByInputID`: `bigint`, `TransactionHash`: `varchar`, `TransactionID`: `bigint`, `Type`: `varchar`, `Value`: `decimal`, `Vout`: `int`}
	_             = bytes.MinRead
)

func testOutputsUpdate(t *testing.T) {
	t.Parallel()

	if len(outputColumns) == len(outputPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	output := &Output{}
	if err = randomize.Struct(seed, output, outputDBTypes, true, outputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Output struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = output.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Outputs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, output, outputDBTypes, true, outputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Output struct: %s", err)
	}

	if err = output.Update(tx); err != nil {
		t.Error(err)
	}
}

func testOutputsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(outputColumns) == len(outputPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	output := &Output{}
	if err = randomize.Struct(seed, output, outputDBTypes, true, outputColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Output struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = output.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Outputs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, output, outputDBTypes, true, outputPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Output struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(outputColumns, outputPrimaryKeyColumns) {
		fields = outputColumns
	} else {
		fields = strmangle.SetComplement(
			outputColumns,
			outputPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(output))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := OutputSlice{output}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testOutputsUpsert(t *testing.T) {
	t.Parallel()

	if len(outputColumns) == len(outputPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	output := Output{}
	if err = randomize.Struct(seed, &output, outputDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Output struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = output.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Output: %s", err)
	}

	count, err := Outputs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &output, outputDBTypes, false, outputPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Output struct: %s", err)
	}

	if err = output.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Output: %s", err)
	}

	count, err = Outputs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
