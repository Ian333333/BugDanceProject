package repository

import (
	"douyinapp/entity"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

type FavoriteDao struct {
}

var (
	favoriteDao  *FavoriteDao
	favoriteOnce sync.Once
)

func NewFavoriteDaoInstance() *FavoriteDao {
	favoriteOnce.Do(func() {
		favoriteDao = &FavoriteDao{}
	})
	return favoriteDao
}

// AddFavorite 在点赞表添加一项数据
func (*FavoriteDao) AddFavorite(userId int64, videoId int64) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/db_douyin?charset=utf8mb4&parseTime=True&loc=Local"

	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	fav := &entity.Favorite{VideoId: videoId, UserId: userId}
	db.Create(fav)
	fmt.Println(fav.Id)
}

// DeleteFavorite 在点赞表删除一条数据
func (*FavoriteDao) DeleteFavorite(userId int64, videoId int64) {

	dsn := "root:123456@tcp(127.0.0.1:3306)/db_douyin?charset=utf8mb4&parseTime=True&loc=Local"

	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db.Delete(entity.Favorite{}, "video_id = ? and user_id = ?", videoId, userId)

}

// FavoriteList 根据用户Id获取该用户点赞的所有视频Id
func (*FavoriteDao) FavoriteList(userId int64) []int64 {
	dsn := "root:123456@tcp(127.0.0.1:3306)/db_douyin?charset=utf8mb4&parseTime=True&loc=Local"

	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	favoriteList := make([]*entity.Favorite, 0)

	db.Where("user_id = ?", userId).Find(&favoriteList)

	videoIdList := make([]int64, 0)
	for _, favorite := range favoriteList {
		videoIdList = append(videoIdList, favorite.VideoId)
	}
	return videoIdList
}
