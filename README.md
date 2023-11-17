# Piyaporn_agnos_backend

โปรแกรมตรวจสอบ password พัฒนาด้วย Gin framework ของภาษา Go 
มีการเชื่อมต่อกับฐานข้อมูล Postgresql ซึ่งชื่อของ database คือ testagnos และ script ที่ใช้ในการสร้างฐานข้อมูลจะอยู่ที่ database/_createdatabase.sql

ในการเริ่มต้นเขียนโค้ด จะมีการรันคำสั่งดังนี้
- go mod init Piyaporn_agnos_backend
- go get -u github.com/gin-gonic/gin
โดย default port จะเป็น 8080
- go get github.com/lib/pq
- go get github.com/stretchr/testify/assert
สามารถรันโปรแกรมได้ด้วยการใช้คำสั่ง go run main.go

จากโปรแกรมจะมีการทดสอบ unit test ที่ไฟล์ unit_test.go โดยจะเป็นการทดสอบ function ของ model และ handler สามารถ run unit test ได้ด้วยการใช้คำสั่ง
go test unit_test.go
