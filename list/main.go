package main

import "database/sql"
import "fmt"
import _ "github.com/mattn/go-sqlite3"
import "log"
import "os"
import "path/filepath"

func main() {
	root := os.Args[1]
	os.Remove("./files.db")
	db, err := sql.Open("sqlite3", "./files.db")
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	createTable(db)

	tx, err := db.Begin()
	if err != nil {
		log.Panic(err)
	}
	findStmt, err := tx.Prepare("select count(*) from file where name = ?")
	defer findStmt.Close()
	if err != nil {
		log.Panic(err)
	}
	fileStmt, err := tx.Prepare("insert into file (name, size) values (?, ?)")
	defer fileStmt.Close()
	if err != nil {
		log.Panic(err)
	}
	pathStmt, err := tx.Prepare("insert into path (name, path) values (?, ?)")
	defer pathStmt.Close()
	if err != nil {
		log.Panic(err)
	}
	defer tx.Rollback()
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Panic(err)
		}
		if info.IsDir() {
			return nil
		}

		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		rows, err := findStmt.Query(info.Name())
		defer rows.Close()
		if err != nil {
			return err
		}
		var count int
		for rows.Next() {
			err = rows.Scan(&count)
			if err != nil {
				return err
			}
		}
		err = rows.Err()
		if err != nil {
			return err
		}
		if count < 1 {
			fmt.Print("Name:", info.Name(), ", Size:", info.Size(), ", ")
			_, err = fileStmt.Exec(info.Name(), info.Size())
			if err != nil {
				return err
			}
		}
		fullpath, err := filepath.Abs(rel)
		if err != nil {
			return err
		}
		_, err = pathStmt.Exec(info.Name(), fullpath)
		fmt.Println("path:", fullpath)
		return nil
	})
	tx.Commit()
	if err != nil {
		log.Panic(err)
	}
}

func createTable(db *sql.DB) error {
	stmt := "create table file (name text not null primary key, size integer, checksum text)"
	_, err := db.Exec(stmt)
	if err != nil {
		return err
	}
	stmt = "create table path (name text not null, path text not null, primary key (name, path))"
	_, err = db.Exec(stmt)
	if err != nil {
		return err
	}
	return nil
}
