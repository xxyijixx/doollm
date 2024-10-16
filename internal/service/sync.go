package service

import (
	"doollm/clients/anythingllm"
	"doollm/internal/repository"
	"log"
	"time"
)

var anythingllmClient = anythingllm.NewClient()

// 设置定时
// func SyncUsers() {
// 	ticker := time.NewTicker(30 * time.Second) // 每 30 秒执行一次
// 	defer ticker.Stop()

// 	for range ticker.C {
// 		log.Println("Syncing users...")
// 		syncUsers()
// 		checkAndUpdateSubscriptions()
// 	}
// }

// 同步 pre_users 表用户
func SyncUsers() {
	currentTime := time.Now().Format(time.DateTime)
	tx, err := repository.DB.Begin()
	if err != nil {
		log.Println("Failed to begin transaction:", err)
		return
	}

	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM pre_workspace_permissions WHERE user_id NOT IN (SELECT userid FROM pre_users)")
	if err != nil {
		log.Println("Failed to delete users:", err)
		return
	}

	_, err = tx.Exec(`
        INSERT INTO pre_workspace_permissions (user_id, create_time, update_time)
        SELECT userid, ?, ? FROM pre_users
        WHERE userid NOT IN (SELECT user_id FROM pre_workspace_permissions)
    `, currentTime, currentTime)
	if err != nil {
		log.Println("Failed to insert new users:", err)
		return
	}

	_, err = tx.Exec(`
        UPDATE pre_workspace_permissions 
        SET update_time = ?
        WHERE user_id IN (SELECT userid FROM pre_users)
    `, currentTime)
	if err != nil {
		log.Println("Failed to update users:", err)
		return
	}

	if err := tx.Commit(); err != nil {
		log.Println("Failed to commit transaction:", err)
		return
	}

	//log.Println("Sync completed successfully")
}
