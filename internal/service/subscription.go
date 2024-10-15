package service

import (
	"database/sql"
	"doollm/internal/repository"
	"time"
)

// 从 pre_workspace_subscriptions 表中获取订阅类型
func GetSubscriptionType() (string, error) {
	var subscriptionType string
	query := "SELECT type FROM pre_workspace_subscriptions LIMIT 1"
	err := repository.DB.QueryRow(query).Scan(&subscriptionType)
	if err == sql.ErrNoRows {
		return "free", nil
	}
	if err != nil {
		return "", err
	}
	return subscriptionType, nil
}

// 更新或插入 pre_workspace_subscriptions 表中的订阅类型和有效期
func SetSubscriptionType(subscriptionType string, startTime, endTime time.Time, isForever bool) error {
	var exists bool
	queryCheck := "SELECT COUNT(1) FROM pre_workspace_subscriptions LIMIT 1"
	err := repository.DB.QueryRow(queryCheck).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if exists {
		queryUpdate := `
            UPDATE pre_workspace_subscriptions 
            SET type = ?, start_time = ?, end_time = ?, is_forever = ?
            WHERE 1 LIMIT 1`
		_, err := repository.DB.Exec(queryUpdate, subscriptionType, startTime, endTime, isForever)
		if err != nil {
			return err
		}
	} else {
		queryInsert := `
            INSERT INTO pre_workspace_subscriptions (type, start_time, end_time, is_forever) 
            VALUES (?, ?, ?, ?)`
		_, err := repository.DB.Exec(queryInsert, subscriptionType, startTime, endTime, isForever)
		if err != nil {
			return err
		}
	}

	return nil
}

// 将非永久订阅恢复为 free
func checkAndUpdateSubscriptions() error {
	query := `
        UPDATE pre_workspace_subscriptions
        SET type = 'free',
			start_time = NOW(),
			end_time = NULL,
			is_forever = FALSE
        WHERE end_time <= NOW() AND is_forever = FALSE`
	_, err := repository.DB.Exec(query)
	return err
}
