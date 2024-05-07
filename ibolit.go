package Lets_GO_ibolit

import (
	"encoding/json"
	"os"
)

type animal struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func Do(sourcefile, resultfile string) error {
	f, err := os.Open(sourcefile)
	if err != nil {
		return err
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	res := make([]animal, 0, 3)
	for dec.More() {
		var a animal
		err = dec.Decode(&a)
		if err != nil {
			return err
		}
		res = append(res, a)
	}

	fnew, err := os.Create(resultfile)
	if err != nil {
		return err
	}
	defer fnew.Close()

	_, err = fnew.WriteString("[")
	if err != nil {
		return err
	}

	for i, anim := range res {
		jsonData, err := json.Marshal(anim)
		if err != nil {
			return err
		}
		_, err = fnew.Write(jsonData)
		if err != nil {
			return err
		}

		if i < len(res)-1 {
			_, err = fnew.WriteString(",\n")
		}
		if err != nil {
			return err
		}
	}

	_, err = fnew.WriteString("]")
	if err != nil {
		return err
	}

	return nil
}
