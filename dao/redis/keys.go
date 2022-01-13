package redis

/*
	Redis Key
*/

const (
	KeyPostInfoHashPrefix = "cl:post:" //map
	KeyPostTimeZSet       = "cl:post:time" //zset
	KeyPostScoreZSet      = "cl:post:score" //zset

	KeyPostVotedZSetPrefix = "cl:post:voted:"  //zset

	KeyCommunityPostSetPrefix = "cl:community:" //set
)