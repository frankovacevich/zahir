package data

const SOURCE_MQTT = "MQTT"

type Variable struct {
	ID           int         `json:"id"`
	Name         string      `json:"name"`
	Description  string      `json:"description"`
	DefaultValue interface{} `json:"default_value"`
}

type VariableValues struct {
	VariableID int           `json:"variable_id"`
	SequenceID int           `json:"sequence_id"`
	Values     []interface{} `json:"values"`
}

type Event struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type EventValues struct {
	EventID   int `json:"event_id"`
	Timestamp int `json:"timestamp"`
}

type Source struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Type        string     `json:"type"`
	Variables   []Variable `json:"variables"`
	Events      []Event    `json:"events"`
}

type Sequence struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Length      int    `json:"length"`
	SourcesIDs  []int  `json:"source_ids"`
}

type dataCollection struct {
	Sources        []Source         `json:"sources"`
	Sequences      []Sequence       `json:"sequences"`
	VariableValues []VariableValues `json:"variable_values"`

	// Internal
	filename          string                   `json:"-"`
	SequencesMap      map[int]*Sequence        `json:"-"`
	SourcesMap        map[int]*Source          `json:"-"`
	VariableValuesMap map[int][]VariableValues `json:"-"` // SequenceID -> VariableValues
}

type PlayerState struct {
	Running      bool    `json:"running"`       // Is the player running
	CurrentIdx   int     `json:"current_idx"`   // Index in the queue
	Queue        []int   `json:"queue"`         // Sequence IDs
	StepDuration float64 `json:"step_duration"` // Duration of each step in seconds
	OnEndStop    bool    `json:"on_end_stop"`   // Stop when the sequence ends
	Step         int     `json:"step"`          // Current step in  running sequence
	Updated      int64   `json:"updated"`       // Last update timestamp
}
