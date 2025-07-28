package class_03

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
SQL语句练习
题目1：基本CRUD操作
假设有一个名为 students 的表，包含字段
id （主键，自增）、
name （学生姓名，字符串类型）、
age （学生年龄，整数类型）、
grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*/

/*
CREATE TABLE `students`
(`id` bigint AUTO_INCREMENT,`name` varchar,`age` smallint,`grade` varchar,PRIMARY KEY (`id`))
*/
type Student struct {
	ID    uint64 `gorm:"type:bigint;primaryKey;autoIncrement"`
	Name  string `gorm:"type:varchar(120);size:120"`
	Age   uint16 `gorm:"type:smallint"`
	Grade string `gorm:"type:varchar(120);size:120"`
}

func BasicSQL(step string) {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Student{})
	switch step {
	case "create":
		s := &Student{Name: "张三", Age: 20, Grade: "三年级"}
		db.Create(s)
	case "query":
		ss := []Student{}
		db.Where("age > ?", 18).Find(&ss)
		fmt.Println(ss)
	case "update":
		db.Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")
	case "delete":
		db.Where("age < ?", 15).Delete(&Student{})
	default:
		panic("Step string is illegal")
	}
}

/*
SQL语句练习
题目2：事务语句
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和
transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，
向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/
type Account struct {
	ID      uint64 `gorm:"type:bigint;primaryKey;autoIncrement"`
	Balance int64  `gorm:"type:bigint"`
}

type Transaction struct {
	ID            uint64 `gorm:"type:bigint;primaryKey;autoIncrement"`
	FromAccountID uint64 `gorm:"type:bigint"`
	ToAccountID   uint64 `gorm:"type:bigint"`
	Amount        int64  `gorm:"type:bigint"`
}

func TransactionSQL() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Transaction{})
	db.Transaction(func(tx *gorm.DB) error {

		var accA Account
		var accB Account
		if err := tx.First(&accA, 1).Error; err != nil {
			return err
		} else if accA.Balance < 10000 {
			return errors.New("账号A余额不足")
		} else if err := tx.First(&accB, 2).Error; err != nil {
			return err
		} else {
			accA.Balance -= 10000
			accB.Balance += 10000
			txRecord := Transaction{FromAccountID: accA.ID, ToAccountID: accB.ID, Amount: 10000}
			if err := tx.Save(accA).Error; err != nil {
				return err
			} else if err := tx.Save(accB).Error; err != nil {
				return err
			} else if err := tx.Create(&txRecord).Error; err != nil {
				return err
			} else {
				return nil
			}
		}
	})
}

/*
进阶gorm
题目1：模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，
其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章），
Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。
*/
type User struct {
	gorm.Model
	Username   string `gorm:"unique"`
	PostsCount int
	Posts      []Post
}

type Post struct {
	gorm.Model
	Title         string
	Content       string
	CommentsCount int
	CommentStatus string
	UserID        uint
	Comments      []Comment
}

type Comment struct {
	gorm.Model
	Title   string
	Content string
	PostID  uint
	UserID  uint
}

func BlogCreateTable() {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Comment{})
}

/*
进阶gorm
题目2：关联查询
基于上述博客系统的模型定义。
要求 ：
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
编写Go代码，使用Gorm查询评论数量最多的文章信息。
*/
func LinkedTableQuery() (*User, *Post) {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	var user User
	db.Preload("Posts.Comments").Preload("Posts").Where("username = ?", "Lucy").First(&user)
	if len(user.Posts) > 0 {
		mostCommented := user.Posts[0]
		maxNum := len(mostCommented.Comments)
		for i := 1; i < len(user.Posts); i++ {
			if len(user.Posts[i].Comments) > maxNum {
				mostCommented = user.Posts[i]
				maxNum = len(user.Posts[i].Comments)
			}
		}
		return &user, &mostCommented
	} else {
		return &user, nil
	}
}

/*
进阶gorm
题目3：钩子函数
继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/
func (p *Post) AfterCreate(tx *gorm.DB) error {
	tx.Model(&User{}).Where("id = ?", p.UserID).Update("posts_count", gorm.Expr("posts_count + 1"))
	return nil
}

func (p *Post) AfterDelete(tx *gorm.DB) error {
	tx.Model(&User{}).Where("id = ?", p.UserID).Update("posts_count", gorm.Expr("posts_count - 1"))
	return nil
}

func (c *Comment) AfterCreate(tx *gorm.DB) error {
	tx.Model(&Post{}).Where("id = ?", c.PostID).
		Updates(map[string]interface{}{"comments_count": gorm.Expr("comments_count + 1"), "comment_status": "有评论"})
	return nil
}

func (c *Comment) AfterDelete(tx *gorm.DB) error {
	var post Post
	tx.First(&post, c.PostID)
	if post.CommentsCount < 2 {
		tx.Model(&Post{}).Where("id = ?", c.PostID).
			Updates(map[string]interface{}{"comments_count": 0, "comment_status": "无评论"})
		return nil
	} else {
		tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comments_count", gorm.Expr("comments_count - 1"))
		return nil
	}
}

func DeleteComment(p *Comment) {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	if p.PostID <= 0 {
		db.First(p)
	}
	db.Delete(p)
}

func CreatePost(p *Post) {
	db, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	db.Create(p)

}
