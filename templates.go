package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"sync"
)

type templateMap map[string]*template.Template

type templates struct {
	sync.Mutex

	base      string
	templates templateMap
}

func newTemplates(base string) *templates {
	return &templates{
		base:      base,
		templates: make(templateMap),
	}
}

func (t *templates) Add(name string, template *template.Template) {
	t.Lock()
	defer t.Unlock()

	t.templates[name] = template
}

func (t *templates) Exec(name string, ctx interface{}) (io.WriterTo, error) {
	t.Lock()
	defer t.Unlock()

	template, ok := t.templates[name]
	if !ok {
		log.Printf("template %s not found", name)
		return nil, fmt.Errorf("no such template: %s", name)
	}

	buf := bytes.NewBuffer([]byte{})
	err := template.ExecuteTemplate(buf, t.base, ctx)
	if err != nil {
		log.Printf("error parsing template %s: %s", name, err)
		return nil, err
	}

	return buf, nil
}
