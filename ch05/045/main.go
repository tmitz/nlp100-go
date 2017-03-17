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
		buf.WriteString(fmt.Sprintf("%s\t%s\n", cs.Verb, strings.Join(cs.Particles, "")))
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
			for _, cc := range sentence {
				if c.Srcs == cc.Dst && cc.HasParticle() {
					particles = append(particles, cc.LastParticle().Base)
				}
			}
			if len(particles) > 0 {
				sort.Strings(particles)
				cps = append(cps, CasePattern{Verb: c.FirstVerb().Base, Particles: particles})
			}
		}
	}
	return cps
}
