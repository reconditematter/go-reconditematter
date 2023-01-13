package randomnames

import (
	"errors"
	"math"
	"math/rand"
	"time"
)

// HumanName consists of family name, given name, gender.
type HumanName struct {
	Family string
	Given  string
	Gender string
}

// Generate generates `count` random human names.
// It returns an error if `count` is greater than `maxCount` (1000).
func Generate(count uint) ([]HumanName, error) {
	const maxCount = 1000
	if count > maxCount {
		return nil, errors.New("randomnames.Generate: invalid count")
	}

	names := make(map[[2]string]string) // [family,given] -> gender
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for len(names) < int(count) {
		i, j := rng.Intn(maxCount), rng.Intn(maxCount)
		if rng.Uint64() < math.MaxUint64/2 {
			name := [2]string{family[i], givenf[j]}
			names[name] = "female"
		} else {
			name := [2]string{family[i], givenm[j]}
			names[name] = "male"
		}
	}

	hnames := make([]HumanName, 0, count)
	for name, gender := range names {
		hn := HumanName{Family: name[0], Given: name[1], Gender: gender}
		hnames = append(hnames, hn)
	}

	return hnames, nil
}
