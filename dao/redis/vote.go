package redis

import (
	"github.com/go-redis/redis"
	"math"
	"time"
	"web_app/models"
)

const (
	OneWeekInSeconds         = 7 * 24 * 3600
	VoteScore        float64 = 432
	PostPerAge               = 20
)

/*

KeyPostInfoHashPrefix = "cl:post:"
KeyPostTimeZSet       = "cl:post:time" 记录每个post的创建时间
KeyPostScoreZSet      = "cl:post:score" 记录每个post的分数

KeyPostVotedZSetPrefix = "cl:post:voted:" 记录每个post的所有user投票情况
KeyCommunityPostSetPrefix = "cl:community:"
*/

func PostVote(userID string,p *models.ParamVoteData)(err error){
	//1.检查帖子的发布时间,超过一星期不允许投票
	postId := p.PostID
	direction := float64( p.Direction )
	postTime := rdb.ZScore(KeyPostTimeZSet,postId ).Val()

	if float64( time.Now().Unix() )-postTime > OneWeekInSeconds{//超过一周了
		return ErrorVoteTimeExpire
	}
	//2.增加帖子的分数
	key := KeyPostVotedZSetPrefix + postId
	now_score := rdb.ZScore(key,userID).Val()   //拿到对应分数
	diff_score := math.Abs( now_score-direction ) //和之前投票形成的差值

	pipline := rdb.TxPipeline()  //开始批量处理命令
	pipline.ZAdd( key,redis.Z{ Score:direction,Member: userID} )
	pipline.ZIncrBy(KeyPostScoreZSet,VoteScore*diff_score*direction,postId )

	switch  diff_score {
	case 1:
		//(1->0)取消投票,投票数-1
		pipline.HIncrBy(KeyPostInfoHashPrefix+postId,"votes",-1)
	case 0:
		//反转投票,投票数不变
	case -1:
		//(0->1)新增投票,投票数+1
		pipline.HIncrBy(KeyPostInfoHashPrefix+postId,"votes",1)
	default:
		return ErrorVoted
	}
	_,err = pipline.Exec()  //统一执行命令
	return
}


//新建port时更新redis中的数据
func CreatePost(postID,userID,title,summary,communityName string)(err error){
	now := float64( time.Now().Unix() )
	vote_key := KeyPostVotedZSetPrefix+postID
	communityKey := KeyCommunityPostSetPrefix+communityName
	postInfo := map[string]interface{}{
		"title":title,
		"post:id":postID,
		"user:id":userID,
		"time":now,
		"votes":1,
		"comments":0,
	}
	//实务操作
	pipeline := rdb.TxPipeline()
	pipeline.ZAdd( vote_key,redis.Z{Score:1,Member: userID } ) //作者默认先投一票
	pipeline.Expire(vote_key,time.Second*OneWeekInSeconds) //一周后过期

	pipeline.HMSet(KeyPostInfoHashPrefix+postID,postInfo) //把port相关信息放入map中
	pipeline.ZAdd( KeyPostScoreZSet,redis.Z{ Score:now+VoteScore,Member: postID} ) //初始化port的分数
	pipeline.SAdd( communityKey,postID ) //保存当前所有存在的post

	_,err = pipeline.Exec()
	return
}

//分页的时候查询帖子
func GetPost(order string,page int64)( []map[string]string ){
	key := KeyPostScoreZSet  //默认按照分数来查询
	if order == "time"{
		key = KeyPostTimeZSet  //按照创建时间来查询
	}
	start := (page-1)*PostPerAge
	end := start+PostPerAge-1
	ids := rdb.ZRevRange(key,start,end).Val()  //查询出对应的帖子id
	postList := make( []map[string]string,0,len(ids)  )
	for _,id := range ids{
		postData := rdb.HGetAll(KeyPostInfoHashPrefix+id).Val() //去对应的详情map查询
		postData["id"] = id
	}
	return postList
}

