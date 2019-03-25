// Copyright (C) 2019 Aaron Browne
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// A copy of the license can be found in the LICENSE file and at
//         https://www.gnu.org/licenses/agpl.html

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/api", apiHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("coming soon")); err != nil {
		log.Printf("Write error: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
