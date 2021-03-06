package parser

import (
	"testing"

	"github.com/freeconf/yang/meta"
	"github.com/freeconf/yang/val"
)

func TestTypeResolve(t *testing.T) {
	yang := `
module ff {
	namespace "ns";

	description "mod";

	revision 99-99-9999 {
	  description "bingo";
	}

	leaf x {
		type int32;
	}
	typedef q {
		type string;
	}
	list y {
		key "id";
		leaf id {
			type string;
		}
	}
	container z {
	  description "z1";
	  leaf z1 {
	    type leafref {
	    	path "../../x";
	    }
	  }
	  leaf z2 {
	    type leafref {
	    	path "../../y/id";
	    }
	  }
		leaf z3 {
			type q;
		}
	}
}
`
	m, err := LoadModuleFromString(nil, yang)
	if err != nil {
		t.Fatal(err)
	}
	z1 := meta.Find(m, "z/z1")
	if z1 == nil {
		t.Error("No z1")
	}
	dt := z1.(meta.HasType).Type()
	if dt.Format() != val.FmtLeafRef {
		t.Errorf("actual type %s", dt.Format())
	}

	dt = z1.(meta.HasType).Type().Resolve()
	if dt.Format() != val.FmtInt32 {
		t.Errorf("actual type %s", dt.Format())
	}
	z3 := meta.Find(m, "z/z3")
	if z3 == nil {
		t.Error("no z3")
	}
	dt = z3.(meta.HasType).Type().Resolve()
	if dt.Format() != val.FmtString {
		t.Errorf("actual type %s", dt.Format())
	}
}
