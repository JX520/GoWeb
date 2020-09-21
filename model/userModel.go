package model

// binding对格式进行校验
// user 结构体
//type UserModel struct {
//	Id int `form:id`
//	Email string `form:"email" binding:"email"`
//	Pwd string `form:"pwd" `
//	//PwdAgain string `form:"pwdAgain" binding:"eqfield=Pwd"`
//	Avatar sql.NullString `form:"avatar"` //string类型无法接受null类型，用Nullstring处理
//}
//
////user保存到数据库的方法
//func (user *UserModel) Save() int64  {
//	result, err := initDB.Db.Exec("insert into go.user (email,password) values(?,?);", user.Email,user.Pwd)
//	if err != nil{
//		log.Panicln("用户插入错误！", err.Error())
//	}
//	id, e := result.LastInsertId();
//	if e != nil {
//		log.Panicln("用户插入id错误！", e.Error())
//	}
//	return id
//}
//
////user通过email查询数据库用户
//func (user *UserModel) QueryUser() UserModel  {
//	u := UserModel{}
//	row := initDB.Db.QueryRow("select * from user where email = ?;", user.Email)
//	e := row.Scan(&u.Id, &u.Email, &u.Pwd, &u.Avatar)
//	if e != nil {
//		log.Panicln(e)
//	}
//	return u
//}
//
////user通过id查询数据库用户
//func (user *UserModel) QueryById(id int) (UserModel, error)  {
//	u := UserModel{}
//	row := initDB.Db.QueryRow("select * from user where id = ?;", id)
//	e := row.Scan(&u.Id, &u.Email, &u.Pwd, &u.Avatar)
//	if e != nil {
//		log.Panicln(e)
//	}
//	return u, e
//}
//
////更新密码和头像
//func (user *UserModel) Update(id int) error  {
//	var stmt, err = initDB.Db.Prepare("update user set password=?,avatar=? where id=? ")
//	if err != nil {
//		log.Panicln("发生了错误", err.Error())
//	}
//	_, err = stmt.Exec(user.Pwd, user.Avatar.String, user.Id)
//	if err != nil {
//		log.Panicln("错误 e", err.Error())
//	}
//
//	return err
//}