package models

import (
	"database/sql"
	"fmt"
)

type Cliente struct {
	Id       int
	Name     string
	Password string
}

type Compra struct {
	Id_Compra  int `json:"id_compra"`
	Id_Cliente int `json:"id_Cliente"`
}

type Detalle struct {
	Id_Compra   int    `json:"id_Compra"`
	Id_Producto int    `json:"id_Producto"`
	Cantidad    int    `json:"cantidad"`
	Fecha       string `json:"fecha"`
}

type Producto struct {
	Id_Producto         int    `json:"id_Producto"`
	Nombre              string `json:"nombre"`
	Cantidad_Disponible int    `json:"cantidad_disponible"`
	Precio_Unitario     int    `json:"precio_unitario"`
}

func GetProduct() []Producto {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/tarea_1_sd")
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}
	defer db.Close()
	results, err := db.Query("SELECT * FROM producto")

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	producto := []Producto{}
	for results.Next() {
		var prod Producto
		err = results.Scan(&prod.Id_Producto, &prod.Nombre, &prod.Cantidad_Disponible, &prod.Precio_Unitario)
		if err != nil {
			panic(err.Error())
		}
		producto = append(producto, prod)
	}
	return producto
}
func GetProductbyid(id string) *Producto {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/tarea_1_sd")
	prod := &Producto{}
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}
	defer db.Close()
	results, err := db.Query("SELECT * FROM producto where id_producto=?", id)

	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	if results.Next() {
		err = results.Scan(&prod.Id_Producto, &prod.Nombre, &prod.Cantidad_Disponible, &prod.Precio_Unitario)
		if err != nil {
			return nil
		}

	} else {
		return nil
	}

	return prod
}

func Addproduct(product Producto) {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/tarea_1_sd")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	insert, err := db.Query("INSERT INTO producto (nombre,cantidad_disponible,precio_unitario) VALUES (?,?,?)", product.Nombre, product.Cantidad_Disponible, product.Precio_Unitario)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}
func Addcompra(compra Compra) {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/tarea_1_sd")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	insert, err := db.Query("INSERT INTO compra(id_cliente) VALUES (?)", compra.Id_Cliente)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}