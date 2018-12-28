package AuthServer

import (
	"fmt"
	"forex/starter"
	"math"
)

type User struct {
	starter.MysqlModel
	Username string
	ParentID int
	Level    int
}

func ModelAddrs() []interface{} {
	return []interface{}{
		&User{},
	}
}

func (m *AuthServer) debugModeData() {
	defer m.Mysql.Connector()()
	level := 10
	admin := User{}
	users := []User{}
	id := 1

	for i := 0; i < level; i++ {
		for j := 0; j < int(math.Pow(float64(2), float64(i))); j++ {
			user := User{
				ParentID: int(math.Round(float64(id)/2/0.5) * 0.5),
				Level:    i + 1,
			}
			fmt.Println(float64(2), float64(i+1))
			user.ID = id
			users = append(users, user)

			id++
		}
	}

	for i := range users {
		m.Mysql.DB.Debug().Create(&users[i])
	}
	admin.Level = 0
	admin.Username = "admin"

	m.Mysql.DB.Debug().Create(&admin)
}

func (m *AuthServer) HierachicalUsersCTE() {

	type ObsTree struct {
		ID       int
		ParentID int
		Level    int
		Tree     string
	}

	result := []ObsTree{}

	defer m.Mysql.Connector()()
	m.Mysql.DB.Raw(`
		WITH RECURSIVE results AS (
			SELECT 
				id, parrent_id, 1 AS level, '/'||CAST(id AS VARCHAR) AS tree 
			FROM 
				users
			WHERE 
				parent_id IS 0
			UNION ALL
				SELECT 
					tab.id, tab.parent_id, tab.level + 1, pri.tree||'/'||CAST(t.id AS VARCHAR)
				FROM 
					users tab 
				JOIN 
					obs_resultstree pri 
				ON
					tab.parent_id = pri.id
		)
		SELECT 
			id, parent_id, level, tree 
		FROM 
			results 
		ORDER BY 
			tree
	`).Scan(&result)

	fmt.Println(result)
}
