package sample

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	ASetting string `md:"aSetting,required"`
}

type Input struct {
	Args  string `md:"args,required"`
	Flags string `md:"flags,required"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	strVal, err := coerce.ToString(values["args"])
	if err != nil {
		return err
	}
	r.Args = strVal

	strVal, err = coerce.ToString(values["flags"])
	if err != nil {
		return err
	}
	r.Flags = strVal

	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"args":  r.Args,
		"flags": r.Flags,
	}
}

type Output struct {
	AnOutput string `md:"anOutput"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["anOutput"])
	o.AnOutput = strVal
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"anOutput": o.AnOutput,
	}
}
