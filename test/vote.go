package test

import "strings"

func GetWinner(votes []string) string {

	if votes == nil || len(votes) == 0 {
		return ""
	}

	vMap := make(map[string]int, 0)
	voteOrder := make(map[string]int, 0)

	for idx, name := range votes {
		if name == "" {
			continue
		}

		if _, found := vMap[name]; found {
			vMap[name]++
		} else { // new participant
			vMap[name] = 1
			voteOrder[name] = idx
		}
	}

	var maxVote int
	var winner string

	for name, vote := range vMap {
		if vote > maxVote || (vote == maxVote && voteOrder[winner] > voteOrder[name]) {
			maxVote = vote
			winner = name
		}
	}

	return winner
}

type Vote struct {
	votes []string
}

func GetWinnerFromList(votes []*Vote, numberOfVotesPerPerson int) string {
	if votes == nil || len(votes) == 0 {
		return ""
	}

	rankPointMap := make(map[int]int, numberOfVotesPerPerson)
	for i := 0; i < numberOfVotesPerPerson; i++ {
		rankPointMap[i] = numberOfVotesPerPerson - i
	}

	vMap := make(map[string]int, 0)
	voteOrder := make(map[string]map[int]int, 0)

	for _, vote := range votes {
		for rank, name := range vote.votes {
			if name == "" {
				continue
			}

			if _, found := vMap[name]; found {
				vMap[name] += rankPointMap[rank]
				voteOrder[name][rank] += 1

			} else { // new participant
				vMap[name] = rankPointMap[rank]
				voteOrder[name] = make(map[int]int, numberOfVotesPerPerson)
				for j := 0; j < numberOfVotesPerPerson; j++ {
					voteOrder[name][j] = 0
				}

				voteOrder[name][rank] = 1
			}
		}
	}

	var maxVote int
	var winner string

	for name, vote := range vMap {
		if vote > maxVote {
			maxVote = vote
			winner = name
		} else if vote == maxVote {
			tie := true
			// 1. number of 1 st rank
			for k := 0; k < numberOfVotesPerPerson; k++ {
				if voteOrder[winner][k] != voteOrder[name][k] {
					if voteOrder[winner][k] < voteOrder[name][k] {
						maxVote = vote
						winner = name
						tie = false
						break
					}
				}
			}

			// alpha order
			if tie {
				if strings.Compare(name, winner) < 0 {
					maxVote = vote
					winner = name
				}
			}
		}
	}

	return winner
}
