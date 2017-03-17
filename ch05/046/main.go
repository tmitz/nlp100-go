package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/tmitz/nlp100-go/ch05/041/chunk"
)

type CasePattern struct {
	Verb      string
	Particles []string
	Terms     []string
}

type CasePatterns []CasePattern

func main() {
	file := os.Args[1:]
	sentences := chunk.List(file[0])
	cps := parseCasePattern(sentences)

	f, err := os.Create("save_case_pattern.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewWriter(f)
	for _, cs := range cps {
		buf.WriteString(fmt.Sprintf("%s\t%s\t%s\n", cs.Verb, strings.Join(cs.Particles, " "), strings.Join(cs.Terms, " ")))
	}
	buf.Flush()
}

func parseCasePattern(sentences chunk.Sentences) CasePatterns {
	cps := make(CasePatterns, 0)
	for _, sentence := range sentences {
		for _, c := range sentence {
			if !c.HasVerb() {
				continue
			}

			var particles []string
			var terms []string
			for _, cc := range sentence {
				if c.Srcs == cc.Dst {
					if cc.HasParticle() {
						particles = append(particles, cc.LastParticle().Base)
					}
					var joinstr []string
					for _, m := range cc.Morphs {
						joinstr = append(joinstr, m.Surface)
					}
					terms = append(terms, strings.Join(joinstr, ""))
				}
			}

			if len(particles) > 0 {
				sort.Strings(particles)
				cps = append(cps, CasePattern{Verb: c.FirstVerb().Base, Particles: particles, Terms: terms})
			}
		}
	}
	return cps
}
