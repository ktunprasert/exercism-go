package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Team struct {
	name                                string
	played, wins, draws, losses, points int
}

func (t Team) String() string {
	return fmt.Sprintf("%-31s|  %d |  %d |  %d |  %d |  %d", t.name, t.played, t.wins, t.draws, t.losses, t.points)
}

func (t *Team) Win() {
	t.wins++
	t.points += 3
	t.played++
}

func (t *Team) Draw() {
	t.draws++
	t.points += 1
	t.played++
}

func (t *Team) Lost() {
	t.losses++
	t.played++
}

func getOrCreateTeam(m map[string]*Team, name string) *Team {
    if m[name] == nil {
        m[name] = &Team{name: name}
    }

    return m[name]
}

func Tally(reader io.Reader, writer io.Writer) error {
	games := make(map[string]*Team)
	scan := bufio.NewScanner(reader)

	for scan.Scan() {
		line := scan.Text()

		if len(line) == 0 || line[0] == '#' {
			continue
		}

		match := strings.Split(line, ";")
		if len(match) != 3 {
			return errors.New("invalid match")
		}

		first, second, status := match[0], match[1], match[2]

		t1 := getOrCreateTeam(games, first)
		t2 := getOrCreateTeam(games, second)

		switch status {
		case "win":
			t1.Win()
			t2.Lost()
		case "draw":
			t1.Draw()
			t2.Draw()
		case "loss":
			t2.Win()
			t1.Lost()
		default:
			return errors.New("invalid status")
		}
	}

	teams := make([]*Team, 0)
	for _, t := range games {
		teams = append(teams, t)
	}

	sort.Slice(teams, func(i, j int) bool {
		if teams[i].points == teams[j].points {
			return teams[i].name < teams[j].name
		}

		return teams[i].points > teams[j].points
	})

	_, err := writer.Write([]byte(fmt.Sprintf("%-31s| MP |  W |  D |  L |  P\n", "Team")))
	if err != nil {
		return err
	}
	for _, t := range teams {
		_, err := writer.Write([]byte(fmt.Sprintf("%s\n", t.String())))
		if err != nil {
			return err
		}
	}

	return nil
}
