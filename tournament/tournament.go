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

type tally struct {
	teams map[string]*team
}

type sortTeams []*team

// Tally the received input to the given output writer.
func Tally(input io.Reader, output io.Writer) error {
	tally := newTally()

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			continue
		}
		if text[0] == byte('#') {
			continue
		}

		s := strings.Split(text, ";")
		if len(s) != 3 {
			return fmt.Errorf("invalid input")
		}

		teamA, teamB, result := tally.getTeam(s[0]), tally.getTeam(s[1]), s[2]

		switch result {
		case "win":
			teamA.recordWin()
			teamB.recordLoss()
		case "loss":
			teamA.recordLoss()
			teamB.recordWin()
		case "draw":
			teamA.recordDraw()
			teamB.recordDraw()
		default:
			return fmt.Errorf("invalid result")
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	tally.write(output)

	return nil
}

func (tm *team) recordWin() {
	tm.matches++
	tm.wins++
	tm.points += 3
}

func (tm *team) recordLoss() {
	tm.matches++
	tm.losses++
}

func (tm *team) recordDraw() {
	tm.matches++
	tm.draws++
	tm.points++
}

func newTally() *tally {
	return &tally{teams: make(map[string]*team)}
}

func (ta *tally) getTeam(name string) *team {
	t := ta.teams[name]
	if t == nil {
		t = &team{name: name}
		ta.teams[name] = t
	}

	return t
}

func (ta *tally) write(output io.Writer) {
	sorted := make(sortTeams, len(ta.teams))
	i := 0
	for _, team := range ta.teams {
		sorted[i] = team
		i++
	}
	sort.Sort(sort.Reverse(sorted))

	fmt.Fprintf(output, "%-31s|%3s |%3s |%3s |%3s |%3s\n", "Team", "MP", "W", "D", "L", "P")
	for _, t := range sorted {
		fmt.Fprintf(output, "%-31s|%3d |%3d |%3d |%3d |%3d\n", t.name, t.matches, t.wins, t.draws, t.losses, t.points)
	}
}

func (s sortTeams) Len() int {
	return len(s)
}

func (s sortTeams) Less(i, j int) bool {
	if s[i].points == s[j].points {
		return s[i].name > s[j].name
	}

	return s[i].points < s[j].points
}

func (s sortTeams) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
