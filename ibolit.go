package Lets_GO_ibolit

import (
	"encoding/json"
	"encoding/xml"
	"os"
)

type Patient struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

type patients struct {
	List []Patient `xml:"Patient"`
}

func Do(sourcefile, resultfile string) error {
	f, err := os.Open(sourcefile)
	if err != nil {
		return err
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	p := patients{}
	for dec.More() {
		var a Patient
		err = dec.Decode(&a)
		if err != nil {
			return err
		}
		p.List = append(p.List, a)
	}

	fnew, err := os.Create(resultfile)
	if err != nil {
		return err
	}
	defer fnew.Close()

	_, err = fnew.WriteString(xml.Header)
	if err != nil {
		return err
	}

	enc := xml.NewEncoder(fnew)
	enc.Indent("", " ")
	err = enc.Encode(p)
	if err != nil {
		return err
	}

	return nil
}
