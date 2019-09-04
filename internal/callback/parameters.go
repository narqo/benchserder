package callback

//proteus:generate
type Parameters map[string]string

func (params Parameters) Copy() Parameters {
	if params == nil {
		return nil
	}

	newParams := make(Parameters, len(params))
	for k, v := range params {
		newParams[k] = v
	}
	return newParams
}
