package model

import "Goweb/initDB"

type Article struct {
	Id      int    `json:"id"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

//article的插入方法
//func (article Article) Insert() int {
//	result, e := initDB.Db.Exec("insert into `article` (type, content) values (?, ?);", article.Type, article.Content)
//	//create := initDB.Db.Create(&article)
//	if e != nil {
//		log.Panicln("文章添加失败", e.Error())
//	}
//	i, _ := result.LastInsertId()
//	//i, _ := create.Last()
//	return int(i)
//}



//article的查询所有方法
//func (article Article) FindAll() []Article {
//	rows, e := initDB.Db.Query("select * from ` article`;")
//	if e != nil {
//		log.Panicln("查询数据失败")
//	}
//
//	var articles []Article
//
//	for rows.Next() {
//		var a Article
//		if e := rows.Scan(&a.Id, &a.Type, &a.Content); e == nil {
//			articles = append(articles, a)
//		}
//	}
//	return articles
//}
//
////article的通过ID查询
//func (article Article) FindById() Article {
//	row := initDB.Db.QueryRow("select * from ` article` where id=?;", article.Id)
//	if e := row.Scan(&article.Id, &article.Type, &article.Content); e != nil {
//		log.Panicln("绑定发生错误", e.Error())
//	}
//	return article
//}
//
////article的通过ID删除
//func (article Article) DeleteOne() {
//	if _, e := initDB.Db.Exec("delete from ` article` where id = ?", article.Id); e != nil {
//		log.Panicln("数据发生错误，无法删除")
//	}
//}


// 通过gorm进行数据库操作
func (article Article) TableName() string {
	return "article"
}

func (article Article) Insert() int  {
	create := initDB.Db.Create(&article)
	if create.Error != nil{
		return 0
	}
	return 1

}

func (article Article) FindAll() []Article {
	var articles []Article
	initDB.Db.Find(&articles)
	return articles
}

func (article Article) FindById() Article {
	initDB.Db.First(&article, article.Id)
	return article
}

func (article Article) DeleteOne()  {
	initDB.Db.Delete(article)
}