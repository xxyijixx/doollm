package common

import (
	"context"
	"doollm/repo"
	"doollm/repo/model"
	"strings"

	log "github.com/sirupsen/logrus"
)

func GetUserAndBiuldMap(ctx context.Context) ([]*model.User, map[int64]*model.User, error) {
	// 构建 UserId 到 User 的映射
	users, err := repo.User.WithContext(ctx).Where(repo.User.Bot.Eq(0)).Find()
	if err != nil {
		return nil, nil, err
	}
	userMap := make(map[int64]*model.User)
	for _, user := range users {
		userMap[user.Userid] = user
	}

	return users, userMap, nil
}

// GetUserNames 获取一组用户的名称，使用逗号进行分隔
func GetUserNames(userIds []int64, userMap *map[int64]*model.User) string {
	names := make([]string, len(userIds))
	uMap := *userMap
	for i, userId := range userIds {
		user := uMap[userId]
		if user == nil {
			log.WithField("userId", userId).Warn("查找不到用户信息")
			continue
		}
		names[i] = user.Nickname
		if names[i] == "" {
			names[i] = user.Email
		}

	}
	return strings.Join(names, ",")
}
