package prefly

import (
	"errors"
	"log"
	"net/http"
	"net/url"
)

type preflyee struct {
	url.URL
	Dest string
}

func (pf *preflyee) Go() {
	res, err := http.Get(pf.String())
	if err != nil {
		log.Fatal(err)
	}
	pf.Dest = res.Request.URL.String()
}

func NewPreflyee(targetURL string) (*preflyee, error) {
	pf := &preflyee{}
	urlPtr, err := (pf.Parse(targetURL))
	if err != nil {
		return nil, err
	}
	if urlPtr.Scheme == "" {
		return nil, errors.New("請提供協定 ex: http://, https://")
	}
	pf.URL = *urlPtr
	return pf, nil
}
