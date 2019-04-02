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

type tally []*team

func Tally(input io.Reader, output io.Writer) error {
	var teams = make(map[string]*team)

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

		teamA, teamB, result := teams[s[0]], teams[s[1]], s[2]

		if teamA == nil {
			teamA = &team{name: s[0]}
			teams[s[0]] = teamA
		}
		if teamB == nil {
			teamB = &team{name: s[1]}
			teams[s[1]] = teamB
		}

		teamA.matches++
		teamB.matches++

		switch result {
		case "win":
			teamA.points += 3
			teamA.wins++
			teamB.losses++
		case "loss":
			teamB.points += 3
			teamB.wins++
			teamA.losses++
		case "draw":
			teamA.points++
			teamA.draws++
			teamB.points++
			teamB.draws++
		default:
			return fmt.Errorf("invalid result")
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	newTally(teams).write(output)

	return nil
}

func newTally(teams map[string]*team) *tally {
	tally := make(tally, len(teams))

	i := 0
	for _, team := range teams {
		tally[i] = team
		i++
	}

	sort.Sort(sort.Reverse(tally))

	return &tally
}

func (t *tally) write(output io.Writer) {
	fmt.Fprintf(output, "%-31s|%3s |%3s |%3s |%3s |%3s\n", "Team", "MP", "W", "D", "L", "P")

	for _, team := range *t {
		fmt.Fprintf(output, "%-31s|%3d |%3d |%3d |%3d |%3d\n", team.name, team.matches, team.wins, team.draws, team.losses, team.points)
	}
}

func (t tally) Len() int {
	return len(t)
}

func (t tally) Less(i, j int) bool {
	if t[i].points == t[j].points {
		return t[i].name > t[j].name
	}

	return t[i].points < t[j].points
}

func (t tally) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
