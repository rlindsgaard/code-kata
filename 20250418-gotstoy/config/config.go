package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/knadh/koanf/maps"
	"github.com/spf13/cobra"
)

type Settings struct {
	Ensure              Ensure    `json:"ensure,omitempty"`
	Scope               Scope     `json:"scope,omitempty"`
	UpdateAutomatically *bool     `json:"updateAutomatically,omitempty"`
	UpdateFrequency     Frequency `json:"updateFrequency,omitempty"`
	configPath          string
}

func (s *Settings) GetConfigMap() (map[string]any, error) {
	path, err := s.GetConfigPath()
	if err != nil {
		return nil, err
	}
	return getAppConfigMap(path)
}

func (s *Settings) GetConfigSettings() (Settings, error) {
	config, err := s.GetConfigMap()
	if errors.Is(err, os.ErrNotExist) {
		return Settings{
			Ensure: EnsureAbsent,
			Scope:  s.Scope,
		}, nil
	} else if err != nil {
		return Settings{}, err
	}

	return getAppConfigSettings(s.Scope, config)
}

func getAppConfigSettings(scope Scope, config map[string]any) (Settings, error) {
	maps.IntfaceKeysToStrings(config)

	settings := Settings{
		Scope:  scope,
		Ensure: EnsurePresent,
	}

	updates, ok := config["updates"]
	if ok {
		for key, value := range updates.(map[string]any) {
			switch key {
			case "automatic":
				auto := value.(bool)
				settings.UpdateAutomatically = &auto
			case "checkFrequency":
				intValue := int(value.(float64))
				frequency := Frequency(intValue)
				settings.UpdateFrequency = frequency
			}
		}
	}
	return settings, nil
}

func getAppConfigMap(path string) (map[string]any, error) {
	var config map[string]any

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &config)
	return config, err
}

func (s *Settings) GetConfigPath() (string, error) {
	if s.configPath == "" {
		path, err := getAppConfigPath(s.Scope)
		if err != nil {
			return "", err
		}
		s.configPath = path
	}

	return s.configPath, nil
}

func getAppConfigPath(s Scope) (string, error) {
	args := []string{"show", "path", s.String()}

	output, err := exec.Command("./tstoy", args...).Output()
	if err != nil {
		return "", err
	}

	path := string(output)
	path = strings.Trim(path, "\n")
	path = strings.Trim(path, "\r")

	return path, nil
}

func (s *Settings) Enforce() (*Settings, error) {
	err := s.Validate()
	if err != nil {
		return nil, err
	}
	current, err := s.GetConfigSettings()
	if err != nil {
		return nil, err
	}

	if s.Ensure == EnsureAbsent {
		return s.remove(current)
	}

	if current.Ensure == EnsureAbsent {
		return s.create(current)
	}

	// update the config file
	return s.update(current)
}

func (s *Settings) remove(current Settings) (*Settings, error) {
	if current.Ensure == EnsureAbsent {
		return s, nil
	}

	err := os.Remove(s.configPath)
	if err != nil {
		return &current, err
	}

	return s, nil
}

func (s *Settings) create(current Settings) (*Settings, error) {
	configDir := filepath.Dir(s.configPath)
	if err := os.MkdirAll(configDir, 0750); err != nil {
		return &current, fmt.Errorf(
			"failed to create config directory %s: %s",
			configDir,
			err,
		)
	}

	configFile, err := os.Create(s.configPath)
	if err != nil {
		return &current, fmt.Errorf(
			"failed to create config file '%s': %s",
			s.configPath,
			err,
		)
	}

	settings := make(map[string]any)
	updates := make(map[string]any)
	addUpdates := false
	if s.UpdateAutomatically != nil {
		addUpdates = true
		updates["automatic"] = *s.UpdateAutomatically
	}
	if s.UpdateFrequency != 0 {
		addUpdates = true
		updates["checkFrequency"] = int(s.UpdateFrequency)
	}
	if addUpdates {
		settings["updates"] = updates
	}

	configJson, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		return &current, fmt.Errorf(

			"unable to convert settings to json: %s",
			err,
		)
	}

	_, err = configFile.Write(configJson)
	if err != nil {
		return &current, fmt.Errorf(
			"failed to write config file '%s': %s",
			s.configPath,
			err,
		)
	}
	return s, nil
}

func (s *Settings) update(current Settings) (*Settings, error) {
	writeConfig := false

	currentMap, err := s.GetConfigMap()
	if err != nil {
		return nil, err
	}

	maps.IntfaceKeysToStrings(currentMap)

	updates, ok := currentMap["updates"]

	if !ok {
		currentMap["updates"] = make(map[string]any)
		updates = currentMap["updates"]
	}

	shouldSetUA := false
	if s.UpdateAutomatically != nil {
		if current.UpdateAutomatically == nil {
			shouldSetUA = true
		} else if *s.UpdateAutomatically != *current.UpdateAutomatically {
			shouldSetUA = true
		}
	}

	if shouldSetUA {
		writeConfig = true
		updates.(map[string]any)["automatic"] = *s.UpdateAutomatically
	} else if current.UpdateAutomatically != nil {
		updates.(map[string]any)["automatic"] = *current.UpdateAutomatically
	}

	if s.UpdateFrequency != 0 && s.UpdateFrequency != current.UpdateFrequency {
		writeConfig = true
		updates.(map[string]any)["checkFrequency"] = s.UpdateFrequency
	} else if current.UpdateFrequency != 0 {
		updates.(map[string]any)["checkFrequency"] = current.UpdateFrequency
	}

	if !writeConfig {
		return s, nil
	}

	currentMap["updates"] = updates.(map[string]any)

	configJson, err := json.MarshalIndent(currentMap, "", "  ")
	if err != nil {
		return &current, fmt.Errorf(
			"unable to convert settings to json: %s",
			err,
		)
	}

	err = os.WriteFile(s.configPath, configJson, 0644)
	if err != nil {
		return &current, fmt.Errorf(
			"failed to write config file %s: %s",
			s.configPath,
			err,
		)
	}

	return s, nil
}

func (s *Settings) Validate() error {
	if s.Scope == ScopeUndefined {
		return fmt.Errorf(
			"the Scope setting isn't defined. Must define a Scope for Settings",
		)
	}

	if s.Ensure == EnsureAbsent {
		return nil
	}

	if s.UpdateFrequency != 0 {
		return s.UpdateFrequency.Validate()
	}

	return nil
}

func (s *Settings) Print() error {
	configJson, err := json.Marshal(s)
	if err != nil {
		return err
	}

	fmt.Println(string(configJson))
	return nil
}

/* Begin define Ensure */
type Ensure int

const (
	EnsureUndefined Ensure = iota
	EnsurePresent
	EnsureAbsent
)

func (e Ensure) String() string {
	switch e {
	case EnsurePresent:
		return "present"
	case EnsureAbsent:
		return "absent"
	default:
		return "undefined"
	}
}

func ParseEnsure(s string) (Ensure, error) {
	switch strings.ToLower(s) {
	case "absent":
		return EnsureAbsent, nil
	case "present":
		return EnsurePresent, nil
	}
	return EnsureUndefined, fmt.Errorf(
		"unable to convert '%s' to type Ensure, must be one of 'absent', 'present'",
		s,
	)
}

func (e Ensure) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

func (ensure *Ensure) UnmarshalJSON(data []byte) (err error) {
	var e string
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	if *ensure, err = ParseEnsure(e); err != nil {
		return err
	}
	return nil
}

var EnsureMap = map[Ensure][]string{
	EnsureAbsent:  {"absent"},
	EnsurePresent: {"present"},
}

func EnsureFlagCompletion(
	cmd *cobra.Command,
	args []string,
	toComplete string,
) ([]string, cobra.ShellCompDirective) {
	completions := []string{
		"absent\tThe configuration file shouldn't exist.",
		"present\tThe configuration file should exist.",
	}
	return completions, cobra.ShellCompDirectiveNoFileComp
}

/* End define Ensure */

/* Begin define Scope */
type Scope int

const (
	ScopeUndefined Scope = iota
	ScopeMachine
	ScopeUser
)

func (s Scope) String() string {
	switch s {
	case ScopeMachine:
		return "machine"
	case ScopeUser:
		return "user"
	}
	return "undefined"
}

func (s Scope) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (scope *Scope) UnmarshalJSON(data []byte) (err error) {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if *scope, err = ParseScope(s); err != nil {
		return err
	}
	return nil
}

func ParseScope(s string) (Scope, error) {
	switch strings.ToLower(s) {
	case "machine":
		return ScopeMachine, nil
	case "user":
		return ScopeUser, nil
	}
	return ScopeUndefined, fmt.Errorf(
		"unable to convert '%s' to type Scope, must be one of 'machine', 'user'",
		s,
	)
}

var ScopeMap = map[Scope][]string{
	ScopeMachine: {"machine"},
	ScopeUser:    {"user"},
}

func ScopeFlagCompletion(
	cmd *cobra.Command,
	args []string,
	toComplete string,
) ([]string, cobra.ShellCompDirective) {
	completions := []string{
		"machine\tThe configuration file is machine specific.",
		"user\tThe configuration file is user specific.",
	}
	return completions, cobra.ShellCompDirectiveNoFileComp
}

/* End define Scope */

/* Begin define Frequency */
type Frequency int

func (f Frequency) Validate() error {
	v := int(f)
	if v < 1 || v > 90 {
		return fmt.Errorf(
			"invalid frequency value %d, must be between 1 and 90",
			v,
		)
	}
	return nil
}

func (f *Frequency) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return err
	}

	*f = Frequency(v)

	return f.Validate()
}

func (f *Frequency) Type() string {
	return "int"
}

func (f *Frequency) String() string {
	return strconv.Itoa(int(*f))
}

/* End define Frequency */
