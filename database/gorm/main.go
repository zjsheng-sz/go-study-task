package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	// task1()
	// task2()
	task3()

}

/*
题目1：模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。
*/

type User struct {
	gorm.Model
	Name      string `gorm: "size:100; not null"`
	Email     string `gorm: "size:100; unique; not null"`
	Password  string `gorm: "size:100; not null"`
	PostCount int
	Posts     []Post //一对多关系： has many
}

type Post struct {
	gorm.Model
	Title    string    `gorm: "size:100; not null"`
	Content  string    `gorm: "type:text; not null"`
	UserID   uint      `gorm: not null`
	User     User      // belongs to
	Comments []Comment //一对多关系： has many
}

type Comment struct {
	gorm.Model
	Content string `gorm: "type:text; not null"`
	UserID  uint   `gorm: not null`
	PostID  uint   `gorm: not null`
	User    User   // belongs to
	Post    Post   // belongs to
}

func task1() {

	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("数据库创建成功")

}

func crateData() {

	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	users := []User{User{Name: "zhangsan", Email: "1@qq.com", Password: "123456"},
		User{Name: "lisi", Email: "2@qq.com", Password: "123456"},
		User{Name: "wanger", Email: "3@qq.com", Password: "123456"},
	}
	db.Create(&users)

	posts := []Post{Post{Title: "title1", Content: "post content1", UserID: 1},
		Post{Title: "title2", Content: "post content2", UserID: 1},
		Post{Title: "title3", Content: "post content3", UserID: 2},
		Post{Title: "title4", Content: "post content4", UserID: 2},
		Post{Title: "title5", Content: "post content5", UserID: 3},
		Post{Title: "title6", Content: "post content6", UserID: 3},
	}
	db.Create(&posts)

	comments := []Comment{Comment{Content: "coment1", UserID: 1, PostID: 1},
		Comment{Content: "comment content1", UserID: 1, PostID: 1},
		Comment{Content: "comment content2", UserID: 1, PostID: 2},
		Comment{Content: "comment content3", UserID: 1, PostID: 3},
		Comment{Content: "comment content4", UserID: 2, PostID: 4},
		Comment{Content: "comment content5", UserID: 2, PostID: 5},
		Comment{Content: "comment content6", UserID: 2, PostID: 6},
		Comment{Content: "comment content7", UserID: 2, PostID: 1},
	}
	db.Create(&comments)
}

/*
题目2：关联查询
基于上述博客系统的模型定义。
要求 ：
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
编写Go代码，使用Gorm查询评论数量最多的文章信息。
*/

func task2() {

	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// //查询用户1 的所有文章及其评论
	// var posts []Post
	// db.Where("user_id", 1).Preload("Comments").Find(&posts)
	// for _, v := range posts {
	// 	fmt.Printf("%v\n", v.Title)
	// }

	//查询评论数最多的文章信息
	var post Post

	db.Model(&Post{}).Select("posts.*, Count(comments.id) as comment_count").Joins("left join comments on comments.post_id = posts.id").Group("posts.id").Order("comment_count desc").First(&post)

	fmt.Println(post.Title, post.UserID, post)
}

/*
题目3：钩子函数
继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/
func task3() {

	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	var post = Post{Title: "title8", Content: "post content8", UserID: 1}
	db.Create(&post)

}

func (post *Post) AfterCreate(tx *gorm.DB) (err error) {

	var count int
	tx.Model(&Post{}).Select("count(*) count").Group("user_id").Find(&count)
	fmt.Println("afterCreate ", count)
	tx.Model(&User{}).Where("id=?", post.UserID).Update("post_count", count)
	return

}
