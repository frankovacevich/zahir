package data

//
//import "fmt"
//
//func CreateSequence() error {
//	sequence := Sequence{ID: len(_sequences) + 1, Length: 120}
//	_sequences = append(_sequences, sequence)
//	_sequencesMap[sequence.ID] = &sequence
//	return nil
//}
//
//func CreateSource() error {
//	source := Source{ID: len(_sources) + 1, Type: SOURCE_MQTT}
//	_sources = append(_sources, source)
//	_sourcesMap[source.ID] = &source
//	return nil
//}
//
//func CreateVariable() error {
//	// TODO
//	return nil
//}
//
//func CreateEvent() error {
//	// TODO
//	return nil
//}
//
//func UpdateSequence(sequence Sequence) error {
//	s, exists := _sequencesMap[sequence.ID]
//	if !exists {
//		return fmt.Errorf("Sequence %d not found", sequence.ID)
//	}
//	s.Name = sequence.Name
//	s.Description = sequence.Description
//	s.Length = sequence.Length
//	s.SourcesIDs = sequence.SourcesIDs
//	return nil
//}
//
//func UpdateSource(source Source) error {
//	s, exists := _sourcesMap[source.ID]
//	if !exists {
//		return fmt.Errorf("Source %d not found", source.ID)
//	}
//	s.Name = source.Name
//	s.Description = source.Description
//	s.Type = source.Type
//	return nil
//}
//
//func UpdateVariable(variable Variable) error {
//	// TODO
//	return nil
//}
//
//func UpdateEvent(event Event) error {
//	// TODO
//	return nil
//}
//
//func DeleteSequence(sequence Sequence) error {
//	// TODO
//	return nil
//}
//
//func DeleteSource(source Source) error {
//	// TODO
//	return nil
//}
//
//func DeleteVariable(variable Variable) error {
//	// TODO
//	return nil
//}
//
//func DeleteEvent(event Event) error {
//	// TODO
//	return nil
//}
//
