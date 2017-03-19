package bflag

import (
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type bflag struct {
	bools       map[string]*BoolOption
	strings     map[string]*StringOption
	ints        map[string]*IntOption
	floats      map[string]*FloatOption
	args        []string
	invalidArgs []string
}

var bf *bflag

func init() {
	bf = &bflag{
		bools:       make(map[string]*BoolOption),
		strings:     make(map[string]*StringOption),
		ints:        make(map[string]*IntOption),
		floats:      make(map[string]*FloatOption),
		args:        make([]string, 0),
		invalidArgs: make([]string, 0),
	}
}

// BoolOption type
type BoolOption struct {
	Name      string
	ShortName byte
	Value     bool
	HelpText  string
}

// StringOption type
type StringOption struct {
	Name      string
	ShortName byte
	Value     string
	HelpText  string
}

// IntOption type
type IntOption struct {
	Name      string
	ShortName byte
	Value     int
	HelpText  string
}

// FloatOption type
type FloatOption struct {
	Name      string
	ShortName byte
	Value     float64
	HelpText  string
}

// DefineBool func
func DefineBool(name string, shortName byte, defaultVal bool) *BoolOption {
	data := &BoolOption{
		Name:      name,
		ShortName: shortName,
		Value:     defaultVal,
		HelpText:  "",
	}
	bf.bools[name] = data
	bf.bools[string(shortName)] = data
	return data
}

// DefineString func
func DefineString(name string, shortName byte, defaultVal string) *StringOption {
	data := &StringOption{
		Name:      name,
		ShortName: shortName,
		Value:     defaultVal,
		HelpText:  "",
	}
	bf.strings[name] = data
	bf.strings[string(shortName)] = data
	return data
}

// DefineInt func
func DefineInt(name string, shortName byte, defaultVal int) *IntOption {
	data := &IntOption{
		Name:      name,
		ShortName: shortName,
		Value:     defaultVal,
		HelpText:  "",
	}
	bf.ints[name] = data
	bf.ints[string(shortName)] = data
	return data
}

// DefineFloat func
func DefineFloat(name string, shortName byte, defaultVal float64) *FloatOption {
	data := &FloatOption{
		Name:      name,
		ShortName: shortName,
		Value:     defaultVal,
		HelpText:  "",
	}
	bf.floats[name] = data
	bf.floats[string(shortName)] = data
	return data
}

// Args returns non-option arguments
func Args() []string {
	return bf.args
}

// Invalid returns invalid options
func Invalid() []string {
	return bf.invalidArgs
}

// GetBool returns the specified *BoolOption if defined
func GetBool(name string) *BoolOption {
	if b, ok := bf.bools[name]; ok {
		return b
	}
	return nil
}

// SetBool func
func SetBool(name string, value bool) error {
	if obj, ok := bf.bools[name]; ok {
		obj.Value = value
	} else {
		return errors.New("cannot set undefined bool: " + name)
	}
	return nil
}

// GetString returns the specified *StringOption if defined
func GetString(name string) *StringOption {
	if s, ok := bf.strings[name]; ok {
		return s
	}
	return nil
}

// SetString func
func SetString(name string, value string) error {
	if s, ok := bf.strings[name]; ok {
		s.Value = value
	} else {
		return errors.New("cannot set undefined string: " + name)
	}
	return nil
}

// GetInt returns the specified *IntOption if defined
func GetInt(name string) *IntOption {
	if i, ok := bf.ints[name]; ok {
		return i
	}
	return nil
}

// SetInt func
func SetInt(name string, value int) error {
	if i, ok := bf.ints[name]; ok {
		i.Value = value
	} else {
		return errors.New("cannot set undefined int: " + name)
	}
	return nil
}

// GetFloat returns the specified *FloatOption if defined
func GetFloat(name string) *FloatOption {
	if f, ok := bf.floats[name]; ok {
		return f
	}
	return nil
}

// SetFloat func
func SetFloat(name string, value float64) error {
	if f, ok := bf.floats[name]; ok {
		f.Value = value
	} else {
		return errors.New("cannot set undefined float: " + name)
	}
	return nil
}

// Parse will merge the command-line arguments into user-defined options
func Parse() {
	args := os.Args[1:]
	removeIndexes := []int{}
	invalidArgs := []string{}

	for i, arg := range args {
		if isValidOption(arg) {
			// Check to see if there is a valid option set
			name, val := parseOption(arg)

			// Parse boolean option
			if boolOpt := GetBool(name); boolOpt != nil {
				if isValidFlag(arg) {
					if val == "true" {
						SetBool(name, true)
					} else if val == "false" {
						SetBool(name, false)
					} else {
						SetBool(name, true)
					}
					removeIndexes = append(removeIndexes, i)
					continue
				} else {
					invalidArgs = append(invalidArgs, "invalid option '"+arg+"'")
					removeIndexes = append(removeIndexes, i)
					continue
				}
			}

			// Check `option=value` vs. `option value` for non-boolean flags
			if strings.Index(arg, "=") == -1 {
				// must get the next command-line argument for value
				if i == len(args)-1 {
					invalidArgs = append(invalidArgs, "option '"+arg+"' is missing a value")
					removeIndexes = append(removeIndexes, i)
					continue
				}
				val = args[i+1]
				removeIndexes = append(removeIndexes, i, i+1)
			} else {
				removeIndexes = append(removeIndexes, i)
			}

			// Parse string option
			if strOpt := GetString(name); strOpt != nil {
				SetString(name, val)
				continue
			}

			// Parse int option
			if intOpt := GetInt(name); intOpt != nil {
				if intVal, err := strconv.Atoi(val); err != nil {
					invalidArgs = append(invalidArgs, "option '"+arg+"' is an invalid int ("+val+")")
				} else {
					SetInt(name, intVal)
				}
				continue
			}

			// Parse float option
			if floatOpt := GetFloat(name); floatOpt != nil {
				if floatVal, err := strconv.ParseFloat(val, 64); err != nil {
					invalidArgs = append(invalidArgs, "option '"+arg+"' is an invalid float ("+val+")")
				} else {
					SetFloat(name, floatVal)
				}
				continue
			}

		} else {
			// Invalid option found, check to see if it begins with - or --
			isOption, _ := regexp.MatchString("\\A-{1,2}", arg)
			if isOption {
				invalidArgs = append(invalidArgs, "'"+arg+"' is not a valid option")
				removeIndexes = append(removeIndexes, i)
			}
		}
	}

	// Populate Args with all other non-option command-line arguments
	for i := len(removeIndexes) - 1; i >= 0; i-- {
		j := removeIndexes[i]
		if j == len(args)-1 {
			args = args[:j]
		} else {
			args = append(args[:j], args[j+1:]...)
		}
	}
	bf.args = args
	bf.invalidArgs = invalidArgs
}
