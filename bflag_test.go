package bflag

import (
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
			t.Fatalf("Defined bool value not correct. Got %t expected `true`", option.Value)
		}
		if option != bOpt {
			t.Fatal("Value returned by DefineBool does not match bools[\"mybool\"]")
		}
	} else {
		t.Fatal("bools[\"mybool\"] is expected to contain a *BoolOption but does not")
	}

	if option, ok := bf.bools["b"]; ok {
		if option.Value != true {
			t.Fatalf("Defined bool value not correct. Got %t expected `true`", option.Value)
		}
		if option != bOpt {
			t.Fatal("Value returned by DefineBool does not match bools[\"b\"]")
		}
	} else {
		t.Fatal("bools[\"b\"] is expected to contain a *BoolOption but does not")
	}

	// StringOption
	if option, ok := bf.strings["mystr"]; ok {
		if option.Value != "foo" {
			t.Fatalf("Defined string value not correct. Got %s expected `\"foo\"`", option.Value)
		}
		if option != sOpt {
			t.Fatal("Value returned by DefineString does not match strings[\"mystr\"]")
		}
	} else {
		t.Fatal("strings[\"mystr\"] is expected to contain a *StringOption but does not")
	}

	if option, ok := bf.strings["s"]; ok {
		if option.Value != "foo" {
			t.Fatalf("Defined string value not correct. Got %s expected `\"foo\"`", option.Value)
		}
		if option != sOpt {
			t.Fatal("Value returned by DefineString does not match strings[\"s\"]")
		}
	} else {
		t.Fatal("strings[\"s\"] is expected to contain a *StringOption but does not")
	}

	// IntOption
	if option, ok := bf.ints["myint"]; ok {
		if option.Value != 10 {
			t.Fatalf("Defined int value not correct. Got %d expected `10`", option.Value)
		}
		if option != iOpt {
			t.Fatal("Value returned by DefineInt does not match ints[\"myint\"]")
		}
	} else {
		t.Fatal("ints[\"myint\"] is expected to contain a *IntOption but does not")
	}

	if option, ok := bf.ints["i"]; ok {
		if option.Value != 10 {
			t.Fatalf("Defined int value not correct. Got %d expected `10`", option.Value)
		}
		if option != iOpt {
			t.Fatal("Value returned by DefineInt does not match ints[\"i\"]")
		}
	} else {
		t.Fatal("ints[\"i\"] is expected to contain a *IntOption but does not")
	}

	// FloatOption
	if option, ok := bf.floats["myfloat"]; ok {
		if option.Value != 2.25 {
			t.Fatalf("Defined float value not correct. Got %.2f expected `2.25`", option.Value)
		}
		if option != fOpt {
			t.Fatal("Value returned by DefineFloat does not match ints[\"myfloat\"]")
		}
	} else {
		t.Fatal("floats[\"myfloat\"] is expected to contain a *FloatOption but does not")
	}

	if option, ok := bf.floats["f"]; ok {
		if option.Value != 2.25 {
			t.Fatalf("Defined float value not correct. Got %.2f expected `2.25`", option.Value)
		}
		if option != fOpt {
			t.Fatal("Value returned by DefineFloat does not match ints[\"f\"]")
		}
	} else {
		t.Fatal("floats[\"f\"] is expected to contain a *FloatOption but does not")
	}
}

func TestArgs(t *testing.T) {

}

func TestInvalid(t *testing.T) {

}

func TestCombinedShortOptions(t *testing.T) {

}
