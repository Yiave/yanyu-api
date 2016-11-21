package yanyu

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"regexp"
)

func GetAll(rows *sql.Rows, columns []string) []map[string]interface{} {
	count := len(columns)

	queryData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuesPtrs := make([]interface{}, count) // 指向values的指针

	for rows.Next() {
		// 先做一次指针拷贝
		for i := 0; i < count; i++ {
			valuesPtrs[i] = &values[i]
		}
		rows.Scan(valuesPtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		queryData = append(queryData, entry)
	}
	return queryData
}

func GetPart(row *sql.Row, columns []string) map[string]interface{} {
	count := len(columns)
	values := make([]interface{}, count)
	valuesPtrs := make([]interface{}, count) // 指向values的指针
	// 先做一次指针拷贝
	for i := 0; i < count; i++ {
		valuesPtrs[i] = &values[i]
	}
	if err := row.Scan(valuesPtrs...); err != nil {
		log.Println("Error scanning major row:", err)
		return nil
	} else {
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		return entry
	}
}

const (
	NumberPattern = `^[1-9]+?[0-9]*$`
)

var (
	NumberRegexp = regexp.MustCompile(NumberPattern)
)

func IsNumber(id string) bool {
	return NumberRegexp.Match([]byte(id))
}

func mustPrepare(db *sql.DB, query string) *sql.Stmt {
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatalf("Error when preparing statement %q: %s", query, err)
	}
	return stmt
}
