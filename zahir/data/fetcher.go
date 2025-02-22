package data

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var d = dataCollection{}

func Load(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteValue, &d)
	if err != nil {
		return err
	}

	// Create maps
	d.SourcesMap = map[int]*Source{}
	d.SequencesMap = map[int]*Sequence{}
	d.VariableValuesMap = map[int][]VariableValues{}
	for i, source := range d.Sources {
		d.SourcesMap[source.ID] = &d.Sources[i]
	}
	for i, sequence := range d.Sequences {
		if sequence.SourcesIDs == nil {
			sequence.SourcesIDs = []int{}
		}
		d.SequencesMap[sequence.ID] = &d.Sequences[i]
		d.VariableValuesMap[sequence.ID] = *fillVariableValues(&d.Sequences[i])
	}

	d.filename = filename
	return nil
}

func fillVariableValues(sequence *Sequence) *[]VariableValues {
	sequenceValues := []VariableValues{}

	// Get all variables for the sequence, with the default values
	defaultValues := map[int]interface{}{}
	for _, sourceID := range sequence.SourcesIDs {
		source, err := GetSource(sourceID)
		if err != nil {
			continue
		}
		for _, variable := range source.Variables {
			defaultValues[variable.ID] = variable.DefaultValue
		}
	}

	// Fix already defined variable values with sequence length
	definedValues := map[int]bool{}
	for i := range d.VariableValues {
		v := d.VariableValues[i]
		if v.SequenceID != sequence.ID {
			continue
		}
		definedValues[v.VariableID] = true

		if len(v.Values) > sequence.Length {
			v.Values = v.Values[:sequence.Length]
		}
		if len(v.Values) < sequence.Length {
			for {
				if len(v.Values) >= sequence.Length {
					break
				}
				v.Values = append(v.Values, defaultValues[v.VariableID])
			}
		}
		sequenceValues = append(sequenceValues, v)
	}

	// Fill missing variable values
	for variableID := range defaultValues {
		if definedValues[variableID] {
			continue
		}
		newV := VariableValues{VariableID: variableID, SequenceID: sequence.ID}
		for i := 0; i < sequence.Length; i++ {
			newV.Values = append(newV.Values, defaultValues[variableID])
		}
		sequenceValues = append(sequenceValues, newV)
	}

	return &sequenceValues
}

func LoadAs(filenameFrom, filenameTo string) error {
	err := Load(filenameFrom)
	if err != nil {
		return err
	}
	d.filename = filenameTo
	return nil
}

func Save() error {
	if d.filename == "" {
		return fmt.Errorf("must load data before saving")
	}

	jsonData, err := json.MarshalIndent(d, "", "    ")
	if err != nil {
		return err
	}

	err = os.WriteFile(d.filename, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Get all
func GetSources() ([]Source, error) {
	return d.Sources, nil
}

func GetSequences() ([]Sequence, error) {
	return d.Sequences, nil
}

func GetVariables() ([]Variable, error) {
	variables := []Variable{}
	for _, source := range d.Sources {
		variables = append(variables, source.Variables...)
	}
	return variables, nil
}

// Get by ID
func GetSource(sourceID int) (*Source, error) {
	source, exists := d.SourcesMap[sourceID]
	if !exists {
		return nil, fmt.Errorf("Source %d not found", sourceID)
	}
	return source, nil
}

func GetSequence(sequenceID int) (*Sequence, error) {
	sequence, exists := d.SequencesMap[sequenceID]
	if !exists {
		return nil, fmt.Errorf("Sequence %d not found", sequenceID)

	}
	return sequence, nil
}

func GetVariableInSequence(variableID, sequenceID int) (*Variable, error) {
	sequence, err := GetSequence(sequenceID)
	if err != nil {
		return nil, err
	}

	for _, sourceID := range sequence.SourcesIDs {
		source, err := GetSource(sourceID)
		if err != nil {
			continue
		}
		for j := range source.Variables {
			if source.Variables[j].ID == variableID {
				return &source.Variables[j], nil
			}
		}
	}

	return nil, fmt.Errorf("Variable %d not found in sequence %d", variableID, sequenceID)
}

func GetSequenceSources(sequenceID int) ([]Source, error) {
	sequence, err := GetSequence(sequenceID)
	if err != nil {
		return nil, err
	}

	sources := []Source{}
	for _, sourceID := range sequence.SourcesIDs {
		source, err := GetSource(sourceID)
		if err == nil {
			sources = append(sources, *source)
		}
	}

	return sources, nil
}

func GetVariableValues(sequenceID int) ([]VariableValues, error) {
	values, exists := d.VariableValuesMap[sequenceID]
	if !exists {
		return nil, fmt.Errorf("Sequence %d not found", sequenceID)
	}
	return values, nil
}

func SetSequenceLength(sequenceID, length int) error {
	sequence, err := GetSequence(sequenceID)
	if err != nil {
		return err
	}

	sequence.Length = length
	d.VariableValuesMap[sequenceID] = *fillVariableValues(sequence)
	return Save()
}

func SetVariableValues(sequenceID, variableID int, values []interface{}) error {

	// Get sequence and check that variable exists in it
	sequence, err := GetSequence(sequenceID)
	if err != nil {
		return err
	}
	_, err = GetVariableInSequence(variableID, sequenceID)
	if err != nil {
		return err
	}

	// Check that values length is the same as sequence length
	if len(values) != sequence.Length {
		return fmt.Errorf("values length must be %d", sequence.Length)
	}

	// If variable already defined, update it
	for i := range d.VariableValues {
		v := &d.VariableValues[i]
		if v.SequenceID == sequenceID && v.VariableID == variableID {
			v.Values = values
			d.VariableValuesMap[sequenceID] = *fillVariableValues(sequence)
			return Save()
		}
	}

	// If variable not defined, create it
	newV := VariableValues{SequenceID: sequenceID, VariableID: variableID, Values: values}
	d.VariableValues = append(d.VariableValues, newV)
	d.VariableValuesMap[sequence.ID] = append(d.VariableValuesMap[sequence.ID], newV)
	return Save()
}

func CreateSource() error {
	source := Source{ID: len(d.Sources) + 1, Name: "New Source", Type: SOURCE_MQTT}
	d.Sources = append(d.Sources, source)
	d.SourcesMap[source.ID] = &source
	return Save()
}

func CreateVariable(sourceID int) error {
	source, err := GetSource(sourceID)
	if err != nil {
		return err
	}
	variable := Variable{ID: len(source.Variables) + 1}
	source.Variables = append(source.Variables, variable)
	return Save()
}

func CreateSequence() error {
	sequence := Sequence{ID: len(d.Sequences) + 1}
	d.Sequences = append(d.Sequences, sequence)
	d.SequencesMap[sequence.ID] = &sequence
	return Save()
}

func AddSourceToSequence(sequenceID, sourceID int) error {
	sequence, err := GetSequence(sequenceID)
	if err != nil {
		return err
	}

	_, err = GetSource(sourceID)
	if err != nil {
		return err
	}

	sequence.SourcesIDs = append(sequence.SourcesIDs, sourceID)
	return Save()
}
