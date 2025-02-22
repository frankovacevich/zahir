package data

import (
	"fmt"
	"os"
	"testing"
	"zahir/utils"
)

func setUp() func() {
	err := LoadAs("../fixtures.json", "temp.json")
	if err != nil {
		panic(err)
	}
	return func() {
		os.Remove("temp.json")
	}
}

func TestLoadAndSaveData(t *testing.T) {
	defer setUp()()

	utils.AssertIntEqual(3, len(d.Sources), t)
	utils.AssertIntEqual(3, len(d.Sequences), t)
	utils.AssertIntEqual(3, len(d.SequencesMap), t)

	err := Save()
	utils.AssertErrorIsNil(err, t)
}

func TestGetAll(t *testing.T) {
	defer setUp()()

	sources, err := GetSources()
	utils.AssertErrorIsNil(err, t)
	utils.AssertIntEqual(3, len(sources), t)

	sequences, err := GetSequences()
	utils.AssertErrorIsNil(err, t)
	utils.AssertIntEqual(3, len(sequences), t)

	variables, err := GetVariables()
	utils.AssertErrorIsNil(err, t)
	utils.AssertIntEqual(7, len(variables), t)
}

func TestGetByID(t *testing.T) {
	defer setUp()()

	id := 1
	source, err := GetSource(id)
	utils.AssertErrorIsNil(err, t)
	utils.AssertIntEqual(id, source.ID, t)
	utils.AssertPointersEqual(&d.Sources[0], source, t)

	sequence, err := GetSequence(id)
	utils.AssertErrorIsNil(err, t)
	utils.AssertIntEqual(id, sequence.ID, t)
	utils.AssertPointersEqual(&d.Sequences[0], sequence, t)
}

func TestGetVariable(t *testing.T) {
	defer setUp()()

	variable, err := GetVariableInSequence(1, 1)
	utils.AssertErrorIsNil(err, t)
	utils.AssertIntEqual(1, variable.ID, t)
	utils.AssertPointersEqual(&d.Sources[0].Variables[0], variable, t)

	_, err = GetVariableInSequence(7, 1)
	utils.AssertError(err, t)
}

func TestGetSequenceSources(t *testing.T) {
	defer setUp()()

	id := 1
	sources, err := GetSequenceSources(id)
	utils.AssertErrorIsNil(err, t)
	utils.AssertIntEqual(2, len(sources), t)
	utils.AssertIntEqual(1, sources[0].ID, t)
	utils.AssertIntEqual(2, sources[1].ID, t)
}

func TestGetVariableValues(t *testing.T) {
	defer setUp()()

	values, err := GetVariableValues(1)
	fmt.Println(values)
	utils.AssertErrorIsNil(err, t)
	utils.AssertIntEqual(6, len(values), t)
}

func TestSetSequenceLength(t *testing.T) {
	defer setUp()()
	newLength := 5

	err := SetSequenceLength(1, newLength)
	utils.AssertErrorIsNil(err, t)
	utils.AssertIntEqual(newLength, d.Sequences[0].Length, t)

	values, err := GetVariableValues(1)
	utils.AssertErrorIsNil(err, t)
	utils.AssertIntEqual(6, len(values), t)
	utils.AssertIntEqual(newLength, len(values[0].Values), t)
	fmt.Println(values)
}

func TestCreateSource(t *testing.T) {
	defer setUp()()

	err := CreateSource()
	utils.AssertErrorIsNil(err, t)
	utils.AssertIntEqual(4, len(d.Sources), t)
}

func TestCreateVariable(t *testing.T) {
	defer setUp()()

	err := CreateVariable(1)
	utils.AssertErrorIsNil(err, t)
	utils.AssertIntEqual(4, len(d.Sources[0].Variables), t)

	vars, err := GetVariables()
	utils.AssertErrorIsNil(err, t)
	utils.AssertIntEqual(8, len(vars), t)
}

func TestCreateSequence(t *testing.T) {
	defer setUp()()

	err := CreateSequence()
	utils.AssertErrorIsNil(err, t)
	utils.AssertIntEqual(4, len(d.Sequences), t)
}

func TestAddSourceToSequence(t *testing.T) {
	defer setUp()()

	err := AddSourceToSequence(1, 3)
	utils.AssertErrorIsNil(err, t)
	utils.AssertIntEqual(3, len(d.Sequences[0].SourcesIDs), t)
}
