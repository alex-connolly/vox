package vox

import "github.com/end-r/efp"

// Vox ...
func Vox(e *efp.Element) (*Division, []string) {
	d := new(Division)
	errs := make([]string, 0)
	// do the top level division manually - works differently
	d.Name = e.FirstField("name").Value()
	d.Description = e.FirstField("description").Value()
	divs := e.FirstElement("divisions")
	d.DivisionsName = divs.Parameter(0).Value()
	for _, e := range divs.Elements("divisions") {
		d, es := parseDivision(e)
		if errs != nil {
			errs = append(errs, es...)
		}
		d.Divisions[e.Key()] = d
	}
	return d, nil
}

func parseDivision(e *efp.Element) (*Division, []string) {
	d := new(Division)
	return d, nil
}
