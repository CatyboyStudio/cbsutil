package reflext

import (
	"fmt"
	"reflect"
)

type NameWithPkg struct {
	PKG  string
	Name string
}

func NP(pkg, name string) NameWithPkg {
	return NameWithPkg{
		PKG:  pkg,
		Name: name,
	}
}

func NPT(t reflect.Type) NameWithPkg {
	return NameWithPkg{t.PkgPath(), t.Name()}
}

func (o NameWithPkg) IsNil() bool {
	return o.PKG == "" && o.Name == ""
}

func (o NameWithPkg) String() string {
	if o.PKG == "" {
		return o.Name
	}
	return fmt.Sprintf("%s.%s", o.PKG, o.Name)
}

func TypeFullname(typ reflect.Type) string {
	o := NPT(typ)
	if o.IsNil() {
		s := fmt.Sprintf("%v", typ)
		return s
	}
	return o.String()
}
