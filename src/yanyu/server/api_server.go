package yanyu

import (
	"github.com/kataras/iris"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)



const (
	connectionString = "root:Terminal@207@tcp(localhost:3306)/yanyu?charset=utf8"
	maxConnectionCount = 40
)

var (
	db *sql.DB
	// select method
	universityByIdStmp *sql.Stmt // stmt: statement
	universityByNameStmp *sql.Stmt
	majorByIdStmp *sql.Stmt
	majorByNameStmp *sql.Stmt
	questionByIdStmp *sql.Stmt
	answerByIdStmp *sql.Stmt
	shareByIdStmp *sql.Stmt

	// insert method
	questionInsertStmp *sql.Stmt
	answerInsertStmp *sql.Stmt
	shareInsertStmp *sql.Stmt


)

func ServerMain() {
	// ###### database configuration ######
	var err error
	if db, err = sql.Open("mysql", "root:Terminal@207@tcp(localhost:3306)/yanyu?charset=utf8"); err != nil {
		log.Fatalf("Error connect to db: %s", err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatalf("Cannot connect to db: %s", err)
	}
	universityByIdStmp = mustPrepare(db, "SELECT name FROM yanyu_university where id = ?")
	universityByNameStmp = mustPrepare(db, "SELECT id FROM yanyu_university where name = ?")
	majorByIdStmp = mustPrepare(db, "SELECT major FROM yanyu_major where id = ?")
	majorByNameStmp = mustPrepare(db, "SELECT id FROM yanyu_major where major = ?")
	questionByIdStmp = mustPrepare(db, "SELECT * FROM yanyu_question where id = ?")
	answerByIdStmp = mustPrepare(db, "SELECT * FROM yanyu_answer where id = ?")
	shareByIdStmp = mustPrepare(db, "SELECT * FROM yanyu_share where id = ?")
	questionInsertStmp = mustPrepare(db, "INSERT INTO yanyu_question(user_id, " +
		"tag_universities, tag_majors, content, ask_date) values(?, ?, ?, ?, ?)")
	answerInsertStmp = mustPrepare(db, "INSERT INTO yanyu_answer(question_id, " +
		"user_id, content, answer_date, read_count, like_count, fee) values(?, ?, ?, ?, ?, ?, ?)")
	shareInsertStmp = mustPrepare(db, "INSERT INTO yanyu_share(user_id, " +
		"tag_universities, tag_majors, content, post_date, fee, read_count, like_count) values(?, ?, ?, ?, ?, ?, ?, ?)")

	// ###### server configuration ######
	server := iris.New()

	// ###### routers configuration ######
	server.Get("/universities", GetUniversities)
	server.Get("/universities/:id", GetUniversityById)
	//server.Get("/universities/:name", GetUniversityByName)
	server.Get("/majors", GetMajors)
	server.Get("/majors/:id", GetMajorById)
	//server.Get("/majors/:name", GetMajorByName)
	server.Get("/questions", GetQuestions)
	server.Get("/questions/:id", GetQuestionById)
	server.Post("/question", PostQuestion)
	server.Get("/answers", GetAnswers)
	server.Get("/answers/:id", GetAnswerById)
	server.Post("/answer", PostAnswer)
	server.Get("/shares", GetShares)
	server.Get("/shares/:id", GetShareById)
	server.Post("/share", PostShare)

	// ###### server configuration ######
	server.Listen(":8888")
}

func GetUniversities(ctx *iris.Context) {
	rows, err := db.Query("SELECT id, name FROM yanyu_university")
	if err != nil {
		log.Println("Error query university:", err)
	}
	defer rows.Close()
	columns := []string {"id", "name"}
	data := GetAll(rows, columns)
	ctx.JSON(iris.StatusOK, data)
}

func GetUniversityById(ctx *iris.Context) {
	id := ctx.Param("id")
	if IsNumber(id) {  // by id
		name := ""
		if err := universityByIdStmp.QueryRow(id).Scan(&name); err != nil {
			log.Println("Error scanning university row:", err)
		}
		if name != "" {
			ctx.JSON(iris.StatusOK, iris.Map{"name": name})
		} else {
			ctx.EmitError(iris.StatusNotFound)
		}
	} else {  // by name
		name := id
		id := -1
		if err := universityByNameStmp.QueryRow(name).Scan(&id); err != nil {
			log.Println("Error scanning university row:", err)
		}
		if id != -1 {
			ctx.JSON(iris.StatusOK, iris.Map{"id": id})
		} else {
			ctx.EmitError(iris.StatusNotFound)
		}
	}
}

func GetUniversityByName(ctx *iris.Context) {
	name := ctx.Param("name")
	id := -1
	if err := universityByNameStmp.QueryRow(name).Scan(&id); err != nil {
		log.Println("Error scanning university row:", err)
	}
	if id != -1 {
		ctx.JSON(iris.StatusOK, iris.Map{"id": id})
	} else {
		ctx.EmitError(iris.StatusNotFound)
	}
}

func GetMajors(ctx *iris.Context) {
	rows, err := db.Query("SELECT * FROM yanyu_major")
	if err != nil {
		log.Println("Error query major:", err)
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		log.Println("Error get Columns:", err)
	}
	data := GetAll(rows, columns)
	ctx.JSON(iris.StatusOK, data)
}

func GetMajorById(ctx *iris.Context) {
	id := ctx.Param("id")
	if IsNumber(id) { // by id
		name := ""
		if err := majorByIdStmp.QueryRow(id).Scan(&name); err != nil {
			log.Println("Error scanning major row:", err)
		}
		if name != "" {
			ctx.JSON(iris.StatusOK, iris.Map{"name": name})
		} else {
			ctx.EmitError(iris.StatusNotFound)
		}
	} else {  // by name
		name := id
		id := -1
		if err := majorByNameStmp.QueryRow(name).Scan(&id); err != nil {
			log.Println("Error scanning university row:", err)
		}
		if id != -1 {
			ctx.JSON(iris.StatusOK, iris.Map{"id": id})
		} else {
			ctx.EmitError(iris.StatusNotFound)
		}
	}
}

func GetMajorByName(ctx *iris.Context) {
	name := ctx.Param("name")
	id := -1
	if err := majorByNameStmp.QueryRow(name).Scan(&id); err != nil {
		log.Println("Error scanning university row:", err)
	}
	if id != -1 {
		ctx.JSON(iris.StatusOK, iris.Map{"id": id})
	} else {
		ctx.EmitError(iris.StatusNotFound)
	}
}

func GetQuestions(ctx *iris.Context) {
	rows, err := db.Query("SELECT * FROM yanyu_question")
	if err != nil {
		log.Println("Error query university:", err)
	}
	defer rows.Close()
	columns := []string {"id", "user_id", "tag_universities", "tag_majors", "content", "ask_date"}
	data := GetAll(rows, columns)
	ctx.JSON(iris.StatusOK, data)
}

func GetQuestionById(ctx *iris.Context) {
	id := ctx.Param("id")
	if IsNumber(id) {
		columns := []string {"id", "user_id", "tag_universities", "tag_majors", "content", "ask_date"}
		row := questionByIdStmp.QueryRow(id)
		entry := GetPart(row, columns)
		if len(entry) != 0 {
			ctx.JSON(iris.StatusOK, entry)
		} else {
			ctx.EmitError(iris.StatusNotFound)
		}
	} else {
		ctx.EmitError(iris.StatusNotFound)
	}
}

type User struct {
	Name 	string `json:"name" xml:"name" form:"name"`
	Email 	string `json:"email" xml:"email" form:"email"`
}

type Question struct {
	User_id 			int `json:"user_id" xml:"user_id" form:"user_id"`
	Tag_universities   	string `json:"tag_universities" xml:"tag_universities" form:"tag_universities"`
	Tag_majors			string `json:"tag_majors" xml:"tag_majors" form:"tag_majors"`
	Content 			string `json:"content" xml:"content" form:"content"`
	Ask_date 			string `json:"ask_date" xml:"ask_date" form:"ask_date"`
}

type Answer struct {
	Question_id 		int `json:"question_id" xml:"question_id" form:"question_id"`
	User_id				int `json:"user_id" xml:"user_id" form:"user_id"`
	Content 			string `json:"content" xml:"content" form:"content"`
	Answer_date 		string `json:"answer_date" xml:"answer_date" form:"answer_date"`
	Read_count 			int `json:"read_count" xml:"read_count" form:"read_count"`
	Like_count			int `json:"like_count" xml:"like_count" form:"like_count"`
	Fee 				float32 `json:"fee" xml:"fee" form:"fee"`
}

type Share struct {
	User_id				int `json:"user_id" xml:"user_id" form:"user_id"`
	Tag_universities   	string `json:"tag_universities" xml:"tag_universities" form:"tag_universities"`
	Tag_majors			string `json:"tag_majors" xml:"tag_majors" form:"tag_majors"`
	Content 			string `json:"content" xml:"content" form:"content"`
	Post_date 			string `json:"post_date" xml:"post_date" form:"post_date"`
	Fee 				float32 `json:"fee" xml:"fee" form:"fee"`
	Read_count 			int `json:"read_count" xml:"read_count" form:"read_count"`
	Like_count			int `json:"like_count" xml:"like_count" form:"like_count"`
}

func PostQuestion(ctx *iris.Context) {
	question := new(Question)
	if err := ctx.ReadJSON(question); err != nil {
		ctx.EmitError(iris.StatusInternalServerError)
	}

	_, err := questionInsertStmp.Exec(question.User_id, question.Tag_universities, question.Tag_majors,
		question.Content, question.Ask_date)
	if err != nil {
		log.Println("Error insert question:", err)
	} else {
		ctx.JSON(iris.StatusOK, question)
	}
}

func GetAnswers(ctx *iris.Context) {
	rows, err := db.Query("SELECT * FROM yanyu_answer")
	if err != nil {
		log.Println("Error query university:", err)
	}
	defer rows.Close()
	columns := []string {"id", "question_id", "user_id", "content", "answer_date", "read_count", "like_count", "fee"}
	data := GetAll(rows, columns)
	ctx.JSON(iris.StatusOK, data)

}

func GetAnswerById(ctx *iris.Context) {
	id := ctx.Param("id")
	if IsNumber(id) {
		columns := []string {"id", "question_id", "user_id", "content", "answer_date", "read_count", "like_count", "fee"}
		row := answerByIdStmp.QueryRow(id)
		entry := GetPart(row, columns)
		if len(entry) != 0 {
			ctx.JSON(iris.StatusOK, entry)
		} else {
			ctx.EmitError(iris.StatusNotFound)
		}
	} else {
		ctx.EmitError(iris.StatusNotFound)
	}
}

func PostAnswer(ctx *iris.Context) {
	answer := new(Answer)
	if err := ctx.ReadJSON(answer); err != nil {
		ctx.EmitError(iris.StatusInternalServerError)
	}

	_, err := answerInsertStmp.Exec(answer.Question_id, answer.User_id, answer.Content,
		answer.Answer_date, answer.Read_count, answer.Like_count, answer.Fee)
	if err != nil {
		log.Println("Error insert answer:", err)
	} else {
		ctx.JSON(iris.StatusOK, answer)
	}
}

func GetShares(ctx *iris.Context) {
	rows, err := db.Query("SELECT * FROM yanyu_share")
	if err != nil {
		log.Println("Error query university:", err)
	}
	defer rows.Close()
	columns := []string {"id", "user_id", "tag_universities", "tag_majors",
		"content", "post_date", "fee", "read_count", "like_count"}
	data := GetAll(rows, columns)
	ctx.JSON(iris.StatusOK, data)
}

func GetShareById(ctx *iris.Context) {
	id := ctx.Param("id")
	if IsNumber(id) {
		columns := []string {"id", "user_id", "tag_universities", "tag_majors",
			"content", "post_date", "fee", "read_count", "like_count"}
		row := shareByIdStmp.QueryRow(id)
		entry := GetPart(row, columns)
		if len(entry) != 0 {
			ctx.JSON(iris.StatusOK, entry)
		} else {
			ctx.EmitError(iris.StatusNotFound)
		}
	} else {
		ctx.EmitError(iris.StatusNotFound)
	}
}

func PostShare(ctx *iris.Context) {
	share := new(Share)
	if err := ctx.ReadJSON(share); err != nil {
		ctx.EmitError(iris.StatusInternalServerError)
	}

	_, err := shareInsertStmp.Exec(share.User_id, share.Tag_universities, share.Tag_majors, share.Content,
		share.Post_date, share.Fee, share.Read_count, share.Like_count)
	if err != nil {
		log.Println("Error insert share:", err)
	} else {
		ctx.JSON(iris.StatusOK, share)
	}

}
