// Copyright 2014 The gorest Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jroimartin/gorest"
)

func main() {
	s := gorest.NewServer("localhost:8080")

	s.Route("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello world!")
	}))

	log.Fatalln(s.ListenAndServe())
}
