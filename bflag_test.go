package bflag

import (
	"os"
	"testing"
)

func TestDefine(t *testing.T) {
	bOpt := DefineBool("mybool", 'b', true)
	sOpt := DefineString("mystr", 's', "foo")
	iOpt := DefineInt("myint", 'i', 10)
	fOpt := DefineFloat("myfloat", 'f', 2.25)
	Parse()

	// BoolOption
	if option, ok := bf.bools["mybool"]; ok {
		if option.Value != true {
			t.Errorf("Defined bool value not correct. Got %t expected `true`", option.Value)
		}
		if option != bOpt {
			t.Error("Value returned by DefineBool does not match bools[\"mybool\"]")
		}
	} else {
		t.Error("bools[\"mybool\"] is expected to contain a *BoolOption but does not")
	}

	if option, ok := bf.bools["b"]; ok {
		if option.Value != true {
			t.Errorf("Defined bool value not correct. Got %t expected `true`", option.Value)
		}
		if option != bOpt {
			t.Error("Value returned by DefineBool does not match bools[\"b\"]")
		}
	} else {
		t.Error("bools[\"b\"] is expected to contain a *BoolOption but does not")
	}

	// StringOption
	if option, ok := bf.strings["mystr"]; ok {
		if option.Value != "foo" {
			t.Errorf("Defined string value not correct. Got %s expected `\"foo\"`", option.Value)
		}
		if option != sOpt {
			t.Error("Value returned by DefineString does not match strings[\"mystr\"]")
		}
	} else {
		t.Error("strings[\"mystr\"] is expected to contain a *StringOption but does not")
	}

	if option, ok := bf.strings["s"]; ok {
		if option.Value != "foo" {
			t.Errorf("Defined string value not correct. Got %s expected `\"foo\"`", option.Value)
		}
		if option != sOpt {
			t.Error("Value returned by DefineString does not match strings[\"s\"]")
		}
	} else {
		t.Error("strings[\"s\"] is expected to contain a *StringOption but does not")
	}

	// IntOption
	if option, ok := bf.ints["myint"]; ok {
		if option.Value != 10 {
			t.Errorf("Defined int value not correct. Got %d expected `10`", option.Value)
		}
		if option != iOpt {
			t.Error("Value returned by DefineInt does not match ints[\"myint\"]")
		}
	} else {
		t.Error("ints[\"myint\"] is expected to contain a *IntOption but does not")
	}

	if option, ok := bf.ints["i"]; ok {
		if option.Value != 10 {
			t.Errorf("Defined int value not correct. Got %d expected `10`", option.Value)
		}
		if option != iOpt {
			t.Error("Value returned by DefineInt does not match ints[\"i\"]")
		}
	} else {
		t.Error("ints[\"i\"] is expected to contain a *IntOption but does not")
	}

	// FloatOption
	if option, ok := bf.floats["myfloat"]; ok {
		if option.Value != 2.25 {
			t.Errorf("Defined float value not correct. Got %.2f expected `2.25`", option.Value)
		}
		if option != fOpt {
			t.Error("Value returned by DefineFloat does not match ints[\"myfloat\"]")
		}
	} else {
		t.Error("floats[\"myfloat\"] is expected to contain a *FloatOption but does not")
	}

	if option, ok := bf.floats["f"]; ok {
		if option.Value != 2.25 {
			t.Errorf("Defined float value not correct. Got %.2f expected `2.25`", option.Value)
		}
		if option != fOpt {
			t.Error("Value returned by DefineFloat does not match ints[\"f\"]")
		}
	} else {
		t.Error("floats[\"f\"] is expected to contain a *FloatOption but does not")
	}
	Reset()
}

func TestArgs(t *testing.T) {
	os.Args = []string{
		"program_name", // ignored by Parse()
		"arg1",
		"--mybool",
		"-s", "foo",
		"arg2",
		"--myint=10",
		"-f=2.25",
	}

	bOpt := DefineBool("mybool", 'b', false)
	sOpt := DefineString("mystr", 's', "bar")
	iOpt := DefineInt("myint", 'i', 20)
	fOpt := DefineFloat("myfloat", 'f', 5.25)
	Parse()

	// BoolOption
	if option, ok := bf.bools["mybool"]; ok {
		if option.Value != true {
			t.Errorf("Defined bool value not correct. Got %t expected `true`", option.Value)
		}
		if option != bOpt {
			t.Error("Value returned by DefineBool does not match bools[\"mybool\"]")
		}
	} else {
		t.Error("bools[\"mybool\"] is expected to contain a *BoolOption but does not")
	}

	if option, ok := bf.bools["b"]; ok {
		if option.Value != true {
			t.Errorf("Defined bool value not correct. Got %t expected `true`", option.Value)
		}
		if option != bOpt {
			t.Error("Value returned by DefineBool does not match bools[\"b\"]")
		}
	} else {
		t.Error("bools[\"b\"] is expected to contain a *BoolOption but does not")
	}

	// StringOption
	if option, ok := bf.strings["mystr"]; ok {
		if option.Value != "foo" {
			t.Errorf("Defined string value not correct. Got %s expected `\"foo\"`", option.Value)
		}
		if option != sOpt {
			t.Error("Value returned by DefineString does not match strings[\"mystr\"]")
		}
	} else {
		t.Error("strings[\"mystr\"] is expected to contain a *StringOption but does not")
	}

	if option, ok := bf.strings["s"]; ok {
		if option.Value != "foo" {
			t.Errorf("Defined string value not correct. Got %s expected `\"foo\"`", option.Value)
		}
		if option != sOpt {
			t.Error("Value returned by DefineString does not match strings[\"s\"]")
		}
	} else {
		t.Error("strings[\"s\"] is expected to contain a *StringOption but does not")
	}

	// IntOption
	if option, ok := bf.ints["myint"]; ok {
		if option.Value != 10 {
			t.Errorf("Defined int value not correct. Got %d expected `10`", option.Value)
		}
		if option != iOpt {
			t.Error("Value returned by DefineInt does not match ints[\"myint\"]")
		}
	} else {
		t.Error("ints[\"myint\"] is expected to contain a *IntOption but does not")
	}

	if option, ok := bf.ints["i"]; ok {
		if option.Value != 10 {
			t.Errorf("Defined int value not correct. Got %d expected `10`", option.Value)
		}
		if option != iOpt {
			t.Error("Value returned by DefineInt does not match ints[\"i\"]")
		}
	} else {
		t.Error("ints[\"i\"] is expected to contain a *IntOption but does not")
	}

	// FloatOption
	if option, ok := bf.floats["myfloat"]; ok {
		if option.Value != 2.25 {
			t.Errorf("Defined float value not correct. Got %.2f expected `2.25`", option.Value)
		}
		if option != fOpt {
			t.Error("Value returned by DefineFloat does not match ints[\"myfloat\"]")
		}
	} else {
		t.Error("floats[\"myfloat\"] is expected to contain a *FloatOption but does not")
	}

	if option, ok := bf.floats["f"]; ok {
		if option.Value != 2.25 {
			t.Errorf("Defined float value not correct. Got %.2f expected `2.25`", option.Value)
		}
		if option != fOpt {
			t.Error("Value returned by DefineFloat does not match ints[\"f\"]")
		}
	} else {
		t.Error("floats[\"f\"] is expected to contain a *FloatOption but does not")
	}
	Reset()
}

func TestNArgs(t *testing.T) {
	os.Args = []string{
		"program_name", // ignored by Parse()
		"arg1",
		"--mybool",
		"-s", "foo",
		"arg2",
		"--myint=10",
		"-f=2.25",
	}

	DefineBool("mybool", 'b', false)
	DefineString("mystr", 's', "bar")
	DefineInt("myint", 'i', 20)
	DefineFloat("myfloat", 'f', 5.25)
	Parse()

	numArgs := NArgs()
	if numArgs != 2 {
		t.Errorf("NArgs() returned %d but expected 2.", numArgs)
	}
	Reset()
}

func TestNOptions(t *testing.T) {
	os.Args = []string{
		"program_name", // ignored by Parse()
		"arg1",
		"--mybool",
		"-s", "foo",
		"arg2",
		"--myint=10",
		"-f=2.25",
	}

	DefineBool("mybool", 'b', false)
	DefineString("mystr", 's', "bar")
	DefineInt("myint", 'i', 20)
	DefineFloat("myfloat", 'f', 5.25)
	Parse()

	numOpts := NOptions()
	if numOpts != 4 {
		t.Errorf("NOptions() returned %d but expected 4.", numOpts)
	}
	Reset()
}

func TestInvalid(t *testing.T) {
	os.Args = []string{
		"program_name", // ignored by Parse()
		"arg1",
		"--mybool",
		"-foo=bar", // invalid (wrong format)
		"-b=misc",  // invalid (wrong format for bool)
		"-s", "foo",
		"arg2",
		"--bar", // invalid (not defined)
		"--myint=10",
		"-f=2.25",
	}

	DefineBool("mybool", 'b', false)
	DefineString("mystr", 's', "bar")
	DefineInt("myint", 'i', 20)
	DefineFloat("myfloat", 'f', 5.25)
	Parse()

	numInvalid := len(Invalid())
	if numInvalid != 3 {
		t.Errorf("Invalid() returned %d but expected 3.", numInvalid)
	}
	Reset()
}

func TestCombinedShortOptions(t *testing.T) {
	os.Args = []string{
		"program_name", // ignored by Parse()
		"arg1",
		"-abc",
		"arg2",
	}

	DefineBool("boola", 'a', false)
	DefineBool("boolb", 'b', false)
	DefineBool("boolc", 'c', false)
	Parse()

	if option, ok := bf.bools["boola"]; ok {
		if option.Value != true {
			t.Errorf("Defined bool value not correct. Got %t expected `true`", option.Value)
		}
	} else {
		t.Error("bools[\"boola\"] is expected to contain a *BoolOption but does not")
	}

	if option, ok := bf.bools["a"]; ok {
		if option.Value != true {
			t.Errorf("Defined bool value not correct. Got %t expected `true`", option.Value)
		}
	} else {
		t.Error("bools[\"a\"] is expected to contain a *BoolOption but does not")
	}

	if option, ok := bf.bools["boolb"]; ok {
		if option.Value != true {
			t.Errorf("Defined bool value not correct. Got %t expected `true`", option.Value)
		}
	} else {
		t.Error("bools[\"boolb\"] is expected to contain a *BoolOption but does not")
	}

	if option, ok := bf.bools["b"]; ok {
		if option.Value != true {
			t.Errorf("Defined bool value not correct. Got %t expected `true`", option.Value)
		}
	} else {
		t.Error("bools[\"b\"] is expected to contain a *BoolOption but does not")
	}

	if option, ok := bf.bools["boolc"]; ok {
		if option.Value != true {
			t.Errorf("Defined bool value not correct. Got %t expected `true`", option.Value)
		}
	} else {
		t.Error("bools[\"boolc\"] is expected to contain a *BoolOption but does not")
	}

	if option, ok := bf.bools["c"]; ok {
		if option.Value != true {
			t.Errorf("Defined bool value not correct. Got %t expected `true`", option.Value)
		}
	} else {
		t.Error("bools[\"c\"] is expected to contain a *BoolOption but does not")
	}

	Reset()
}
