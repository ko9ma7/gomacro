// this file was generated by gomacro command: import _b "path/filepath"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	filepath "path/filepath"
)

// reflection: allow interpreted code to import "path/filepath"
func init() {
	Packages["path/filepath"] = Package{
		Name: "filepath",
		Binds: map[string]Value{
			"Abs":	ValueOf(filepath.Abs),
			"Base":	ValueOf(filepath.Base),
			"Clean":	ValueOf(filepath.Clean),
			"Dir":	ValueOf(filepath.Dir),
			"ErrBadPattern":	ValueOf(&filepath.ErrBadPattern).Elem(),
			"EvalSymlinks":	ValueOf(filepath.EvalSymlinks),
			"Ext":	ValueOf(filepath.Ext),
			"FromSlash":	ValueOf(filepath.FromSlash),
			"Glob":	ValueOf(filepath.Glob),
			"HasPrefix":	ValueOf(filepath.HasPrefix),
			"IsAbs":	ValueOf(filepath.IsAbs),
			"Join":	ValueOf(filepath.Join),
			"ListSeparator":	ValueOf(filepath.ListSeparator),
			"Match":	ValueOf(filepath.Match),
			"Rel":	ValueOf(filepath.Rel),
			"Separator":	ValueOf(filepath.Separator),
			"SkipDir":	ValueOf(&filepath.SkipDir).Elem(),
			"Split":	ValueOf(filepath.Split),
			"SplitList":	ValueOf(filepath.SplitList),
			"ToSlash":	ValueOf(filepath.ToSlash),
			"VolumeName":	ValueOf(filepath.VolumeName),
			"Walk":	ValueOf(filepath.Walk),
			"WalkDir":	ValueOf(filepath.WalkDir),
		}, Types: map[string]Type{
			"WalkFunc":	TypeOf((*filepath.WalkFunc)(nil)).Elem(),
		}, Untypeds: map[string]string{
			"ListSeparator":	"rune:58",
			"Separator":	"rune:47",
		}, 
	}
}
