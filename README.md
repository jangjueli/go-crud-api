# go-crud-api

Assignment Golang

## Overview

- CRUD สำหรับ User
- เชื่อมต่อ PostgreSQL ด้วย Gorm
- รันผ่าน Docker Compose พร้อม DB
- มี unit test ที่ user_test.go

## Database Design

CREATE TABLE user (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

## API Performance Improve

- มี Middleware สำหรับ log ข้อมูล request/response
- รองรับ environment config ผ่าน .env

## Installation

git clone https://github.com/jangjueli/go-crud-api.git
cd go-crud-api

## Run Project (Windows CMD)

docker compose up -d

- ทดสอบ API ที่ `http://localhost:3000/api/user`
- stop container `docker compose down`
