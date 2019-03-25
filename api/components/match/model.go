package match

import (
	score "github.com/AndreBtt/Computsal/api/components/score"
)

type PreviousMatchList struct {
	ID     int
	Team1  string
	Team2  string
	Score1 int
	Score2 int
	Type   int
}

type PreviousMatch struct {
	ID          int
	Team1       string
	Team2       string
	YellowCard1 int
	YellowCard2 int
	RedCard1    int
	RedCard2    int
	Score1      int
	Score2      int
	Type        int
	Group       int
	PlayerScore []score.PlayerScore
}

/*
select
	previous_match.match_type,
	previous_match.match_phase,
    previous_match.fk_match_team1 as team1,
    previous_match.fk_match_team2 as team2,
    COALESCE(player.fk_player_team,"flag_no_team") as team,
    previous_match.id as match_id,
    COALESCE(sum(player_score.quantity),0) as team_score
from match
	left join player_score
		on previous_match.id = player_score.fk_score_match
	left join player
		on player.id = player_score.fk_score_player
group by match_type, team1, team2, team, match_id
order by match_id
*/
