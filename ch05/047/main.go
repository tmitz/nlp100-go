package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/tmitz/nlp100-go/ch05/041/chunk"
)

func main() {
	file := os.Args[1:]
	sentences := chunk.List(file[0])
	for _, sentence := range sentences {
		for _, c := range sentence {
			if !c.HasVerb() {
				continue
			}

			var strary []string
			var particles []string
			var clauses []string
			for _, cc := range sentence {
				if c.Srcs == cc.Dst {
					if cc.HasSahenConnectionNounPlusWo() {
						strary = append(strary, fmt.Sprint(cc))
					}

					if !cc.HasSahenConnectionNounPlusWo() && cc.HasParticle() {
						clauses = append(clauses, fmt.Sprint(cc))
						particles = append(particles, cc.LastParticle().Base)
					}
				}
			}
			if len(strary) > 0 && len(particles) > 0 {
				zip := make(map[string]string, 0)
				sorted := make(map[string]string, 0)
				for i, p := range particles {
					zip[p] = clauses[i]
				}
				sort.Strings(particles)
				for _, v := range particles {
					sorted[v] = zip[v]
				}
				var keyRes []string
				var valRes []string
				for key, val := range sorted {
					keyRes = append(keyRes, key)
					valRes = append(valRes, val)
				}
				fmt.Printf("%s%s\t%s\t%s\n", strary[0], c.FirstVerb().Base, strings.Join(keyRes, " "), strings.Join(valRes, " "))
			}
		}
	}
}
