// Package robotname contains methods to manage robot factory settings.
package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

const max = 26 * 26 * 10 * 10 * 10

var used = make(map[string]bool)

// Robot represents a robot.
type Robot struct {
	name string
}

// Name returns the robots name.
func (r *Robot) Name() (string, error) {
	if len(r.name) > 0 {
		return r.name, nil
	}

	if len(used) >= max {
		return "", fmt.Errorf("namespace exhausted")
	}

	r.name = generateName()
	for used[r.name] {
		r.name = generateName()
	}
	used[r.name] = true

	return r.name, nil
}

// Reset the robot name.
func (r *Robot) Reset() {
	r.name = ""
}

func generateName() string {
	return fmt.Sprintf("%s%s%03d", string('A'+rand.Intn(26)), string('A'+rand.Intn(26)), rand.Intn(1000))
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
