package models

import (
	"encoding/json"
	"strings"
)

func CreateTable() {
	db, _ := Connet()
	db.Exec(`
	CREATE TABLE product
	(id INT AUTO_INCREMENT PRIMARY KEY,
		name_product TEXT NOT NULL, 
		price INT NOT NULL
		)`)
}

func Insert(rows ...interface{}) {
	db, _ := Connet()
	stmp, _ := db.Prepare("INSERT INTO product VALUES(NULL, ?,?)")
	stmp.Exec(rows...)
}
type Product struct {
	Id           int    `json:"id"`
	Name_Product string `json:"name_product"`
	Price        string `json:"price"`
}

func ReadAlldata() string {
	var product Product
	var productall []string
	db, _ := Connet()
	rows, _ := db.Query("SELECT * FROM product")
	for rows.Next() {
		rows.Scan(&product.Id, &product.Name_Product, &product.Price)
		data, _ := json.Marshal(product)
		productall = append(productall, string(data))
	}

	str_join_json := strings.Join(productall, ",")
	return "[" + str_join_json + "]"
}

func DeleteData(id string) (string, int) {
	db, _ := Connet()
	var product Product
	db.QueryRow("SELECT * FROM product  WHERE id = ?", id).Scan(&product.Id, &product.Name_Product, &product.Price)
	if product.Id == 0 {

		return `{"msg": "not fount"}`, 404
	}
	db.Exec("DELETE FROM product WHERE id=?", id)
	return `{"msg":"sucess Delete Data"}`,200
}

func DetailsData(id string) (string, int) {
	db, _ := Connet()
	var product Product
	db.QueryRow("SELECT * FROM product  WHERE id = ?", id).Scan(&product.Id, &product.Name_Product, &product.Price)
	if product.Id == 0 {
		return `{"msg": "id not fount"}`, 404
	}
	json_resp, _ := json.Marshal(product)
	return string(json_resp), 200
}


func UpdateData(rows ...interface{}) ( string, int) { 
	db, _ := Connet()
	var product Product
	db.QueryRow("SELECT * FROM product  WHERE id = ?", rows[2]).Scan(&product.Id, &product.Name_Product, &product.Price)
	if product.Id == 0 {
		return `{"msg": "id not fount"}`, 404
	}
	db.Exec("UPDATE product SET name_product=?, price=? WHERE id=?", rows...)
	return "ok", 201
}