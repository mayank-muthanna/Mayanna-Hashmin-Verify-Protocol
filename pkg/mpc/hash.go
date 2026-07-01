package mpc

import (
	"crypto/sha256"
	"encoding/json"
)

func CommitView(view View) (Commitment, error) {

	data, err := json.Marshal(view)

	if err != nil {
		return nil, err
	}

	hash := sha256.Sum256(data)

	return hash[:], nil
}

func CommitViews(views []*View) ([3]Commitment, error) {

	var commitments [3]Commitment

	for i := 0; i < 3; i++ {

		hash, err := CommitView(
			*views[i],
		)

		if err != nil {
			return commitments, err
		}

		commitments[i] = hash
	}

	return commitments, nil
}
