package goerrors

type context map[string]interface{}

func (c context) ToList() []interface{} {
	res := make([]interface{}, len(c)*2)

	i := 0
	for key, val := range c {
		res[i] = key
		res[i+1] = val
		i += 2
	}

	return res
}
