// Package tournament implements methods to tally the results of a small football competition.
package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type team struct {
	name                                 string
	matches, wins, losses, draws, points int
}

// Tally the received input to the given output writer.
func Tally(input io.Reader, output io.Writer) error {
	tally := make(map[string]team)

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			continue
		}
		if strings.HasPrefix(text, "#") {
			continue
		}

		s := strings.Split(text, ";")
		if len(s) != 3 {
			return fmt.Errorf("invalid input")
		}

		teamA, teamB := tally[s[0]], tally[s[1]]

		teamA.matches++
		teamB.matches++

		switch s[2] {
		case "win":
			teamA.wins++
			teamA.points += 3
			teamB.losses++
		case "loss":
			teamA.losses++
			teamB.wins++
			teamB.points += 3
		case "draw":
			teamA.draws++
			teamA.points++
			teamB.draws++
			teamB.points++
		default:
			return fmt.Errorf("invalid result")
		}

		tally[s[0]], tally[s[1]] = teamA, teamB
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	sorted := make([]team, len(tally))
	i := 0
	for n, t := range tally {
		t.name = n
		sorted[i] = t
		i++
	}
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].points == sorted[j].points {
			return sorted[i].name < sorted[j].name
		}
		return sorted[i].points > sorted[j].points
	})

	fmt.Fprintf(output, "%-31s|%3s |%3s |%3s |%3s |%3s\n", "Team", "MP", "W", "D", "L", "P")
	for _, t := range sorted {
		fmt.Fprintf(output, "%-31s|%3d |%3d |%3d |%3d |%3d\n", t.name, t.matches, t.wins, t.draws, t.losses, t.points)
	}

	return nil
}
