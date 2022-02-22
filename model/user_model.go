package user_model

type UserModel struct {
	Uid int "json:uid"
	Age int "json:age"
	Name string "json:uname"
}

type PostRequestBody struct {
	IsTest bool "json:isTest"
	Uid int "json:uid"
	Uname string "json:uname"
}

type Users struct {
	Id int
	Age int
	Name string
}
