package main

import (
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"
	"unicode"
	"strconv"
	"math"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/gorilla/sessions"
	"net/smtp"
	"math/rand"
	"time"
	
	_ "github.com/go-sql-driver/mysql"

	"os"

	_ "github.com/heroku/x/hmetrics/onload"



)

var store  = sessions.NewCookieStore([]byte("t0p-s3cr3t"))
type user struct {
	iduse    primitive.ObjectID `json:"id,omitempty" bson:"id,omitempty"`
	username string             `json:"username,omitempty" bson:"username,omitempty"`
	password string             `json:"password,omitempty" bson:"password,omitempty"`
	email    string             `json:"email,omitempty" bson:"email,omitempty"`
}
type khachang struct {
	taikhoan     user
	tenkhachhang string
	diachi       string
	sdt          string
}

type diachigiaohang struct{
	Madiachi int
	Tenkhachhang string
	Hokhachhang string
	Thanhpho string
	Huyen string
	Diachi string
	Sodienthoai string
	Idkhachhang int
}
type anhcomment struct {
	 Idanh int 
	 Loaianh int
	 Duongdan string
}
type loaianhcomment struct {
	Idloaianh int 
	Tenloaianh int
	Duongdan string
}
type taikhoan struct{
	Idtaikhoan int
	Email  string
	Tennguoidung string
	Ho string
	Matkhau  string
	Loaitaikhoan int
}
type loaitaikhoan struct{
	Maloaitaikhoan int
	Tenloai string
}
type sanpham struct {
	Idsp        int    ` json:"id" `
	Tensanpham  string `json:"tensanpham"`
	Gia         string ` json:"gia"`
	Loaisanpham int    `json:"loaisanpham"`
	Mota        string `json:"mota"`
	Hinhanh1    string `json:"hinhanh1"`
	Hinhanh2    string `json:"hinhanh2"`
	Hinhanh3    string `json:"hinhanh3"`
}
type sanpham1 struct {
	Idsp1        int     `json:"idsp1"`
	Ngaycapnhat  string  `json:"Ngaycapnhat"`
	Cauhinh      string  `json:"cauhinh"`
	Soluongton   int     `json:"soluongton"`
	Luotxem      int     `json:"luotxem"`
	Luotbinhchon int     `json:"luotbinhchon"`
	Manhasanxuat int     `json:"manhasanxuat"`
	Manhacungcap int     `json:"manhacungcap"`
	Giamgia      int `json:"giamgia"`
	Idsp         int     `json:"idsp"`
}
type hinhanh struct {
	Hinhanh string `json:"hinhanh3"`
	idsp    int    `json:"int"`
}

type loaisanpham struct {
	Idloaisp       int    `json:"id" bson:"_id,omitempty"`
	Tenloaisanpham string `bson:"tenloaisanpham"`
	Mota           string `bson:"mota"`
	// thuoctinh      thuoctinhloaisanpham
}
type thuoctinhloaisanpham struct {
	Idthuoctinh   int
	Idloaisanpham int
	Tenthuoctinh  string
	Mota          string
}

type giohang struct {
	Idgiohang int
	Username  int
}
type donhang struct {
	Madh int
	Ten  string
	Ho string
	Diachi string
	Email  string
	Sdt     string
	Huyen  string
	TP    string
	Note string
	Tinhtrang int
	Ngaylap string
}
type itemhoadon struct{
	Iditemhoadon int
	Idsanpham  int
	Soluong     int
	Tongtien   int
	Mahd      int
}
type structhoadon struct {
	Mahd int
	Madh int
	Iduser  int
	Ship int
}

type itemgiohang struct {
	Iditem  int
	Idgiohang int
	Idsanpham int
	Soluong  int
}

type name struct {
	name string
}
type danhgia struct {
	Iddanhgia int
	Noidung  string
	Idsanpham  int
	Iduser int
	Tenkhachhang string
	Sao int
	Ngaycapnhat string
}
type binhluan struct {
	Idbinhluan int
	Noidung  string
	Idsanpham  int
	Tenkhachhang string
	Sdt string
	Email  string
	Ngaycapnhat string
	Iduser int
}
type binhluans struct {
	Idbinhluans int
	Noidung  string
	Idbinhluan  int
	Tenkhachhang string
	Sdt string
	Email  string
	Ngaycapnhat string
	Iduser int
}
type nhasanxuat struct {
	Idnhasanxuat  int
	Tennhasanxuat string
	Thongtin      string
	Logo          string
}

type nhacungcap struct {
	Idnhacungcap  int
	TenNhaCungCap string
	Diachi        string
	Email         string
	Sdt           string
	Fax           string
}

type File interface {
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Closer
}

func themchitietsanpham(w http.ResponseWriter, r *http.Request) {
	idsanpham := r.FormValue("idsanpham")
	
  if r.Method == "POST" && r.FormValue("chinhsua") != "1" {
	giamgia := r.FormValue("txttensanpham")
	nhasanxuat := r.FormValue("cbbloaisanpham")
	nhacungcap    := r.FormValue("txtmanhacungcap")
	soluongton  := r.FormValue("txtsoluongton")
	luotxem    := r.FormValue("txtluotxem")
	luotbinhchon    := r.FormValue("txtluotbinhchon")
	mota  := r.FormValue("txtmota")
	ngaycapnhat :=r.FormValue("ngaycapnhat")


	insert, err := db1.Query("INSERT INTO `sanpham1` VALUES ('" + giamgia + "','" + string(nhasanxuat) + "','" + string(nhacungcap) +"','"+ngaycapnhat+"','"+luotxem+ "','"+luotbinhchon+"','"+soluongton+"','"+mota+"','"+string(idsanpham)+"')")
	if err != nil {
		panic(err.Error())
	}
	insert.Close()


     }
if r.Method == "POST" && r.FormValue("chinhsua") == "1" {
	tennhacungcap := r.FormValue("txttensanpham")
	diachi := r.FormValue("txtdiachi")
	fax    := r.FormValue("txtfax")
	sdt    := r.FormValue("txtsdt")
	email  := r.FormValue("txtemail")
	insert, err := db1.Query("Update `nhacungcap` set tennhacungcap='"+tennhacungcap+"',diachi='"+diachi+"',email='"+email+"'"+",sdt='"+sdt+"',fax='"+fax+"' where idnhacungcap='"+ r.FormValue("idsp") +"'")
	if err != nil {
		panic(err.Error())
	}
	insert.Close()
}
var Nhasanxuat [] nhasanxuat
var Nhacungcap [] nhacungcap
rows, err := db1.Query("select * from nhasanxuat ")
		for rows.Next() {
			var sanpham1 nhasanxuat
			err = rows.Scan(&sanpham1.Idnhasanxuat, &sanpham1.Tennhasanxuat, &sanpham1.Thongtin, &sanpham1.Logo)
			if err != nil {
				panic(err.Error())
			}
		

			Nhasanxuat = append(Nhasanxuat, sanpham1)

		}

		row, err := db1.Query("select * from nhacungcap ")
		for row.Next() {
			var sanpham1 nhacungcap
			err = row.Scan(&sanpham1.Idnhacungcap, &sanpham1.TenNhaCungCap,&sanpham1.Diachi, &sanpham1.Email, &sanpham1.Sdt,&sanpham1.Fax)
			if err != nil {
				panic(err.Error())
			}

			Nhacungcap = append(Nhacungcap, sanpham1)
		}
		row.Close()
		var sanphamtp []sanpham
		sanphama, err := db1.Query("select * from sanpham  ")
		for sanphama.Next() {
			var sanpham1 sanpham
			err = sanphama.Scan(&sanpham1.Idsp, &sanpham1.Tensanpham, &sanpham1.Loaisanpham, &sanpham1.Gia, &sanpham1.Mota, &sanpham1.Hinhanh1, &sanpham1.Hinhanh2, &sanpham1.Hinhanh3)
			if err != nil {
				panic(err.Error())
			}

			sanphamtp = append(sanphamtp, sanpham1)

		}
		sanphama.Close()
tpl, err := template.ParseFiles("static/quanlythongtinchitietsanpham.html", "static/quanly.html")
if err != nil {
	panic(err.Error())
}
tpl.ExecuteTemplate(w, "quanlythongtinchitietsanpham.html",struct{
	Nhacungcap []nhacungcap
	Nhasanxuat []nhasanxuat
	Sanpham []sanpham
}{Nhacungcap,Nhasanxuat,sanphamtp})

}

func themnhacungcap(w http.ResponseWriter, r *http.Request) {
	session,_ := store.Get(r,"session")
  
	var SS  int
	
	
	if session.Values["iduser"]!=nil{
	  var S= session.Values["iduser"]
	   SS = S.(int)
	  
	
	   
	}else{SS=0
  
	}
  
	selectloaitaikhoan,err := db1.Query("select loaitaikhoan from taikhoan where idtaikhoan='"+strconv.Itoa(SS)+"'")
  
	if err!=nil{
	  panic(err.Error())
	}
	var loaitaikhoanxacnhan int
	for selectloaitaikhoan.Next(){
	 err = selectloaitaikhoan.Scan(&loaitaikhoanxacnhan)
	 if err!=nil{
	  panic(err.Error())
	}
	}
  if loaitaikhoanxacnhan !=3 && SS !=0 {
  
	var mang []nhacungcap
	if r.Method == "POST" && r.FormValue("chinhsua") != "1" {
		tennhacungcap := r.FormValue("txtnhacungcap")
		diachi := r.FormValue("txtdiachi")
		fax    := r.FormValue("txtfax")
		sdt    := r.FormValue("txtsdt")
		email  := r.FormValue("txtemail")
	

		insert, err := db1.Query("INSERT INTO `nhacungcap` VALUES ('" + string(" ") + "','" + tennhacungcap + "','" + diachi + "','" + email +"','"+sdt+"','"+fax+ "')")
		if err != nil {
			panic(err.Error())
		}
		insert.Close()

	}
	if r.Method == "POST" && r.FormValue("chinhsua") == "1" {
		tennhacungcap := r.FormValue("txtnhacungcap")
		diachi := r.FormValue("txtdiachi")
		fax    := r.FormValue("txtfax")
		sdt    := r.FormValue("txtsdt")
		email  := r.FormValue("txtemail")
		insert, err := db1.Query("Update `nhacungcap` set tennhacungcap='"+tennhacungcap+"',diachi='"+diachi+"',email='"+email+"'"+",sdt='"+sdt+"',fax='"+fax+"' where idnhacungcap='"+ r.FormValue("idsp") +"'")
		if err != nil {
			panic(err.Error())
		}
		insert.Close()
	}

	if r.Method == "GET" {

		var idspd = r.FormValue("haha")
		delete, err := db1.Query("delete from nhacungcap where idnhacungcap='" + idspd + "'")
		if err != nil {
			panic(err.Error())
		}
		delete.Close()
	}
	if r.Method == "GET" && r.FormValue("search") != "" {
		sr := r.FormValue("search")
		rows, err := db1.Query("select * from nhacungcap")
		for rows.Next() {

			var sanpham1 nhacungcap
			err = rows.Scan(&sanpham1.Idnhacungcap, &sanpham1.TenNhaCungCap,&sanpham1.Diachi, &sanpham1.Email, &sanpham1.Sdt,&sanpham1.Fax)
			if err != nil {
				panic(err.Error())
			}
			if 0 < strings.Count(sanpham1.TenNhaCungCap, sr) {
				mang = append(mang, sanpham1)
			}

		}
		rows.Close()
	} else {

		rows, err := db1.Query("select * from nhacungcap ")
		for rows.Next() {
			var sanpham1 nhacungcap
			err = rows.Scan(&sanpham1.Idnhacungcap, &sanpham1.TenNhaCungCap,&sanpham1.Diachi, &sanpham1.Email, &sanpham1.Sdt,&sanpham1.Fax)
			if err != nil {
				panic(err.Error())
			}

			mang = append(mang, sanpham1)

		}
		rows.Close()
	}
	tpl, err := template.ParseFiles("static/nhacungcap.html", "static/quanly.html")
	if err != nil {
		panic(err.Error())
	}
	tpl.ExecuteTemplate(w, "nhacungcap.html", struct {
		Nhacungcap []nhacungcap
	}{mang})



}else{
	
	if SS ==0 {

		tpl,err:= template.New("static/page-login.html").Funcs(template.FuncMap{"tinhgiasale":tinhgiasale,"add": add,"floatchuoi":floatchuoi}).ParseFiles("static/page-login.html")
		if err != nil{
			panic(err.Error())
		}
		tpl.ExecuteTemplate(w, "page-login.html", "")
	
	}else{

		tpl, err := template.ParseFiles("static/chuyentrang.html", "static/chuyentrang.html")
		if err != nil {
			panic(err.Error())
		}
		tpl.ExecuteTemplate(w, "chuyentrang.html", "")
	  }	
	}
}




func themnhasanxuat(w http.ResponseWriter, r *http.Request) {

	session,_ := store.Get(r,"session")
	
	var SS  int
	
	
	if session.Values["iduser"]!=nil{
		var S= session.Values["iduser"]
		 SS = S.(int)
		
	
		 
	}else{SS=0

	}

	selectloaitaikhoan,err := db1.Query("select loaitaikhoan from taikhoan where idtaikhoan='"+strconv.Itoa(SS)+"'")

	if err!=nil{
		panic(err.Error())
	}
	var loaitaikhoanxacnhan int
	for selectloaitaikhoan.Next(){
	 err = selectloaitaikhoan.Scan(&loaitaikhoanxacnhan)
	 if err!=nil{
		panic(err.Error())
	}
	}
if loaitaikhoanxacnhan !=3  && SS !=0 {
	var mang []nhasanxuat
	if r.Method == "POST" && r.FormValue("chinhsua") != "1" {
		tennhasanxuat := r.FormValue("txttenhasanxuat")
		thongtin := r.FormValue("txtthongtin")
		anh, handler1, err := r.FormFile("txtanh3")
		defer anh.Close()
		tempfile, err := ioutil.TempFile("static/temp-images", "upload-*.png")
		defer tempfile.Close()
	
       fmt.Print(handler1.Filename)


		fileBytes, err := ioutil.ReadAll(anh)
		tempfile.Write(fileBytes)
		ten :=tempfile.Name()
		
		ten =ten[19:]
		insert, err := db1.Query("INSERT INTO `nhasanxuat` VALUES ('" + string("") + "','" + tennhasanxuat + "','" + thongtin + "','" + ten + "')")
		if err != nil {
			panic(err.Error())
		}
		insert.Close()

	}
	if r.Method == "POST" && r.FormValue("chinhsua") == "1" {
		tennhasanxuat := r.FormValue("txttenhasanxuat")
		thongtin := r.FormValue("txtthongtin")
		anh, handler1, err := r.FormFile("txtanh3")
		if handler1 !=nil{
		defer anh.Close()
	
		if err != nil {
			panic(err.Error())
		}
		defer anh.Close()
		tempfile, err := ioutil.TempFile("static/temp-images", "upload-*.png")
		ten :=tempfile.Name()
		defer tempfile.Close()
		fileBytes, err := ioutil.ReadAll(anh)
		tempfile.Write(fileBytes)
	
		if err != nil {
			panic(err.Error())
		}
		ten =ten[19:]
		tempfile.Write(fileBytes)
	
	
		insert, err := db1.Query("Update `nhasanxuat` set tennhasanxuat='"+tennhasanxuat+"',thongtin='"+thongtin+"',logo='"+ten+"'"+"where idnhasanxuat='"+ r.FormValue("idsp") +"'")
		if err != nil {
			panic(err.Error())
		}
		insert.Close()
	}else{
		insert, err := db1.Query("Update `nhasanxuat` set tennhasanxuat='"+tennhasanxuat+"',thongtin='"+thongtin+"'"+"where idnhasanxuat='"+ r.FormValue("idsp") +"'")
		if err != nil {
			panic(err.Error())
		}
		insert.Close()
	}
	}

	if r.Method == "GET" {

		var idspd = r.FormValue("haha")
		delete, err := db1.Query("delete from nhasanxuat where idnhasanxuat='" + idspd + "'")
		if err != nil {
			panic(err.Error())
		}
		delete.Close()
	}
	if r.Method == "GET" && r.FormValue("search") != "" {
		sr := r.FormValue("search")
		rows, err := db1.Query("select * from nhasanxuat")
		for rows.Next() {

			var sanpham1 nhasanxuat
			err = rows.Scan(&sanpham1.Idnhasanxuat, &sanpham1.Tennhasanxuat, &sanpham1.Thongtin, &sanpham1.Logo)
			if err != nil {
				panic(err.Error())
			}
			if 0 < strings.Count(sanpham1.Tennhasanxuat, sr) {
				mang = append(mang, sanpham1)
			}

		}
		rows.Close()
	} else {

		rows, err := db1.Query("select * from nhasanxuat ")
		for rows.Next() {
			var sanpham1 nhasanxuat
			err = rows.Scan(&sanpham1.Idnhasanxuat, &sanpham1.Tennhasanxuat, &sanpham1.Thongtin, &sanpham1.Logo)
			if err != nil {
				panic(err.Error())
			}

			mang = append(mang, sanpham1)

		}
		rows.Close()
	}
	tpl, err := template.ParseFiles("static/themnhasanxuat.html", "static/quanly.html")
	if err != nil {
		panic(err.Error())
	}
	tpl.ExecuteTemplate(w, "themnhasanxuat.html", struct {
		Nhasanxuat []nhasanxuat
	}{mang})

}else{

	if SS ==0 {

		tpl,err:= template.New("static/page-login.html").Funcs(template.FuncMap{"tinhgiasale":tinhgiasale,"add": add,"floatchuoi":floatchuoi}).ParseFiles("static/page-login.html")
		if err != nil{
			panic(err.Error())
		}
		tpl.ExecuteTemplate(w, "page-login.html", "")
	
	}else{
	tpl, err := template.ParseFiles("static/chuyentrang.html")
	if err != nil {
		panic(err.Error())
	}
	tpl.ExecuteTemplate(w, "chuyentrang.html", "")
     }
}
}
func themsanpham(w http.ResponseWriter, r *http.Request) {
	session,_ := store.Get(r,"session")
	
	var SS  int
	
	
	if session.Values["iduser"]!=nil{
		var S= session.Values["iduser"]
		 SS = S.(int)
		
	
		 
	}else{SS=0

	}

	selectloaitaikhoan,err := db1.Query("select loaitaikhoan from taikhoan where idtaikhoan='"+strconv.Itoa(SS)+"'")

	if err!=nil{
		panic(err.Error())
	}
	var loaitaikhoanxacnhan int
	for selectloaitaikhoan.Next(){
	 err = selectloaitaikhoan.Scan(&loaitaikhoanxacnhan)
	 if err!=nil{
		panic(err.Error())
	}
	}
if loaitaikhoanxacnhan !=3 && SS !=0 {
	var mang []sanpham
	
  if r.Method == "POST" && r.FormValue("chinhsua") != "1" && r.FormValue("idsanpham") !="" {
	 
	idsanpham :=  r.FormValue("idsanpham")
	giamgia := r.FormValue("txttensanpham")
	nhasanxuat := r.FormValue("cbbloaisanpham")
	nhacungcap    := r.FormValue("txtmanhacungcap")
	soluongton  := r.FormValue("txtsoluongton")
	luotxem    := r.FormValue("txtluotxem")
	luotbinhchon    := r.FormValue("txtluotbinhchon")
	mota  := r.FormValue("txtmota")
	ngaycapnhat :=r.FormValue("ngaycapnhat")


	insert, err := db1.Query("INSERT INTO `sanpham1` VALUES ('" + giamgia + "','" + string(nhasanxuat) + "','" + string(nhacungcap) +"','"+ngaycapnhat+"','"+luotxem+ "','"+luotbinhchon+"','"+soluongton+"','"+mota+"','"+string(idsanpham)+"')")
	if err != nil {
		panic(err.Error())
	}
	insert.Close()


	 }
	
	if r.Method == "POST" && r.FormValue("chinhsua") != "1" && r.FormValue("idsanpham") ==""{
		tensanpham := r.FormValue("txttensanpham")
		loaisanpham := r.FormValue("cbbloaisanpham")
		gia := r.FormValue("txtgia")
		anh1, handler1, err := r.FormFile("txtanh1")
		anh2, handler2, err := r.FormFile("txtanh2")
		anh3, handler3, err := r.FormFile("txtanh3")
		mota := r.FormValue("txtmota")
		defer anh1.Close()
		defer anh2.Close()
		defer anh3.Close()
		tempfile, err := ioutil.TempFile("static/temp-images", "upload-*.png")
		if err != nil {
			panic(err.Error())
		}
		ten1 :=tempfile.Name()
		
		ten1 =ten1[19:]
		defer tempfile.Close()
		
		fileBytes, err := ioutil.ReadAll(anh1)
		if err != nil {
			panic(err.Error())
		}
		tempfile.Write(fileBytes)



		tempfile1, err := ioutil.TempFile("static/temp-images", "upload-*.png")
		if err != nil {
			panic(err.Error())
		}
		ten2 :=tempfile1.Name()
		
		ten2 =ten2[19:]
		defer tempfile1.Close()
		
		fileBytes1, err := ioutil.ReadAll(anh2)
		if err != nil {
			panic(err.Error())
		}
		tempfile1.Write(fileBytes1)





		tempfile2, err := ioutil.TempFile("static/temp-images", "upload-*.png")
		if err != nil {
			panic(err.Error())
		}
		ten3 :=tempfile2.Name()
		
		ten3 =ten3[19:]
		defer tempfile2.Close()
		fileBytes2, err := ioutil.ReadAll(anh3)
		if err != nil {
			panic(err.Error())
		}
		tempfile2.Write(fileBytes2)


	      fmt.Print(handler1.Filename+handler2.Filename+handler3.Filename)
		insert, err := db1.Query("insert into sanpham values('"+string(" ")+"','"+tensanpham+"','"+string(loaisanpham)+"','"+gia+"','"+mota+"','"+ten1+"','"+ten2+"','"+ten3+"')")
		if err != nil {
			panic(err.Error())
		}
	
		insert.Close()
        var idsanphamtmp1 int
		slidsp, err := db1.Query("select idsp from sanpham where anh1='"+ten1+"' and anh2='"+ten2+"'")
		if err != nil {
			panic(err.Error())
		}
		for slidsp.Next(){
			err  = slidsp.Scan(&idsanphamtmp1)
			if err != nil {
				panic(err.Error())
			}
		}
		insert.Close()
		

        

			giamgia := r.FormValue("giamgia")
			nhasanxuat := r.FormValue("cbnhasanxuat")
			nhacungcap    := r.FormValue("cbnhacungcap")
			soluongton  := r.FormValue("soluongton")
			luotxem    := r.FormValue("luotxem")
			luotbinhchon    := r.FormValue("luotbinhchon")
			mota  = r.FormValue("txtsoluot")
			ngaycapnhat :=r.FormValue("ngaycapnhat")


			themsanpham1, err := db1.Query("INSERT INTO `sanpham1` VALUES ('" + giamgia + "','" + string(nhasanxuat) + "','" + string(nhacungcap) +"','"+ngaycapnhat+"','"+luotxem+ "','"+luotbinhchon+"','"+soluongton+"','"+mota+"','"+strconv.Itoa(idsanphamtmp1)+"')")
			if err != nil {
				panic(err.Error())
			}
		
			themsanpham1.Close()

	}



	if r.Method == "POST" && r.FormValue("chinhsua") == "1" {
		tensanpham := r.FormValue("txttensanpham")
		loaisanpham := r.FormValue("cbbloaisanpham")
		gia := r.FormValue("txtgia")
		mota := r.FormValue("txtmota")
	fmt.Print(r.FormValue("thaydoianh") )
		if  r.FormValue("thaydoianh") == "on"{
		anh1, handler1, err := r.FormFile("txtanh1")
	
		anh2, handler2, err := r.FormFile("txtanh2")
		
		anh3, handler3, err := r.FormFile("txtanh3")
		
		fmt.Print(handler1.Filename+handler3.Filename)
		
	
			fmt.Print("trong")
		defer anh1.Close()
		defer anh2.Close()
		defer anh3.Close()
		tempfile, err := ioutil.TempFile("static/temp-images", "upload-*.png")
		ten1 :=tempfile.Name()
		
		ten1 =ten1[19:]
		defer tempfile.Close()
		
		fileBytes, err := ioutil.ReadAll(anh1)
		tempfile.Write(fileBytes)
		tempfile1, err := ioutil.TempFile("static/temp-images", "upload-*.png")
		ten2 :=tempfile1.Name()
		
		ten2 =ten2[19:]
		defer tempfile1.Close()
		
		fileBytes1, err := ioutil.ReadAll(anh2)
		tempfile1.Write(fileBytes1)

		tempfile2, err := ioutil.TempFile("static/temp-images", "upload-*.png")
		ten3 :=tempfile2.Name()
		
		ten3 =ten3[19:]
		defer tempfile2.Close()
		fileBytes2, err := ioutil.ReadAll(anh3)
		tempfile2.Write(fileBytes2)

		fmt.Print(handler2.Filename)
		fmt.Print("day la id")
		fmt.Print(r.FormValue("idsp"))
		insert, err := db1.Query("UPDATE `sanpham` SET tensanpham='" + tensanpham + "', loaisanpham='" + loaisanpham + "',gia='" + gia + "',mota='" + mota + "',anh1='" + ten1 + "',anh2='" + ten2 + "',anh3='" + ten3 + "' where idsp='" + r.FormValue("idsp") + "'")
		if err != nil {
			panic(err.Error())
		}
	
		insert.Close()
	}else{
		insert, err := db1.Query("UPDATE `sanpham` SET tensanpham='" + tensanpham + "', loaisanpham='" + loaisanpham + "',gia='" + gia + "',mota='" + mota + "' where idsp='" + r.FormValue("idsp") + "'")
		if err != nil {
			panic(err.Error())
		}
	
		insert.Close()
	}

		     giamgia := r.FormValue("giamgia")
			nhasanxuat := r.FormValue("cbnhasanxuat")
			nhacungcap    := r.FormValue("cbnhacungcap")
			soluongton  := r.FormValue("soluongton")
			luotxem    := r.FormValue("luotxem")
			luotbinhchon    := r.FormValue("luotbinhchon")
			mota  = r.FormValue("txtsoluot")
			ngaycapnhat :=r.FormValue("ngaycapnhat")

			updatesp1, err := db1.Query("Update `sanpham1` set giamgia='"+giamgia+"',manhasanxuat='"+nhasanxuat+"',manhacungcap='"+nhacungcap+"',ngaycapnhat='"+ngaycapnhat+"',luotxem='"+luotxem+"',"+"luotbinhchon='"+luotbinhchon+"' ,soluongton='"+soluongton+"',cauhinh='"+mota+"'"  +"where idsp='"+ r.FormValue("idsp") +"'")
			if err != nil {
				panic(err.Error())
			}
			updatesp1.Close()
	}
	var loaisp []loaisanpham

	row, err := db1.Query("select * from  loaisanpham")
	for row.Next() {

		var lsp loaisanpham
		err = row.Scan(&lsp.Idloaisp, &lsp.Tenloaisanpham, &lsp.Mota)
		if err != nil {
			panic(err.Error())
		}
		loaisp = append(loaisp, lsp)
	}
	row.Close()
	if r.Method == "GET" {

		var idspd = r.FormValue("haha")
		delete, err := db1.Query("delete from sanpham where idsp='" + idspd + "'")
		if err != nil {
			panic(err.Error())
		}
		delete.Close()
		deletesp1, err := db1.Query("delete from sanpham1 where idsp='" + idspd + "'")
		if err != nil {
			panic(err.Error())
		}
		deletesp1.Close()
	}
	if r.Method == "GET" && r.FormValue("search") != "" {
		sr := r.FormValue("search")
		rows, err := db1.Query("select * from sanpham tensanpham ")
		for rows.Next() {

			var sanpham1 sanpham
			err = rows.Scan(&sanpham1.Idsp, &sanpham1.Tensanpham, &sanpham1.Loaisanpham, &sanpham1.Gia, &sanpham1.Mota, &sanpham1.Hinhanh1, &sanpham1.Hinhanh2, &sanpham1.Hinhanh3)
			if err != nil {
				panic(err.Error())
			}
			if 0 < strings.Count(sanpham1.Tensanpham, sr) || 0 < strings.Count(string(sanpham1.Gia), sr) {
				mang = append(mang, sanpham1)
			}

		}
		rows.Close()
	} else {

		rows, err := db1.Query("select * from sanpham  ")
		for rows.Next() {
			var sanpham1 sanpham
			err = rows.Scan(&sanpham1.Idsp, &sanpham1.Tensanpham, &sanpham1.Loaisanpham, &sanpham1.Gia, &sanpham1.Mota, &sanpham1.Hinhanh1, &sanpham1.Hinhanh2, &sanpham1.Hinhanh3)
			if err != nil {
				panic(err.Error())
			}

			mang = append(mang, sanpham1)

		}
		rows.Close()
	}
				var Nhasanxuat [] nhasanxuat
			var Nhacungcap [] nhacungcap
			rows1, err := db1.Query("select * from nhasanxuat ")
					for rows1.Next() {
						var sanpham1 nhasanxuat
						err = rows1.Scan(&sanpham1.Idnhasanxuat, &sanpham1.Tennhasanxuat, &sanpham1.Thongtin, &sanpham1.Logo)
						if err != nil {
							panic(err.Error())
						}
					

						Nhasanxuat = append(Nhasanxuat, sanpham1)

					}
					rows1.Close()
					row1, err := db1.Query("select * from nhacungcap ")
					for row1.Next() {
						var sanpham1 nhacungcap
						err = row1.Scan(&sanpham1.Idnhacungcap, &sanpham1.TenNhaCungCap,&sanpham1.Diachi, &sanpham1.Email, &sanpham1.Sdt,&sanpham1.Fax)
						if err != nil {
							panic(err.Error())
						}

						Nhacungcap = append(Nhacungcap, sanpham1)
					}
					row1.Close()

					var arrsanpham1 []sanpham1
					sanphama, err := db1.Query("select * from sanpham1  ")
					for sanphama.Next() {
						var sanpham1 sanpham1
						err = sanphama.Scan(&sanpham1.Giamgia, &sanpham1.Manhasanxuat, &sanpham1.Manhacungcap, &sanpham1.Ngaycapnhat, &sanpham1.Luotxem, &sanpham1.Luotbinhchon, &sanpham1.Soluongton, &sanpham1.Cauhinh,&sanpham1.Idsp1)
						if err != nil {
							panic(err.Error())
						}
			
						arrsanpham1 = append(arrsanpham1, sanpham1)
			
					}
					sanphama.Close()
	tpl, err := template.ParseFiles("static/quanlysanpham.html", "static/quanly.html")
	tpl.ExecuteTemplate(w, "quanlysanpham.html", struct {
		Sanpham []sanpham
	
		Loaisanpham []loaisanpham
		Nhacungcap []nhacungcap
		Nhasanxuat []nhasanxuat
		Sanpham1 []sanpham1
	}{mang,loaisp,Nhacungcap,Nhasanxuat,arrsanpham1})
}else{
	
	if SS ==0 {

		tpl,err:= template.New("static/page-login.html").Funcs(template.FuncMap{"tinhgiasale":tinhgiasale,"add": add,"floatchuoi":floatchuoi}).ParseFiles("static/page-login.html")
		if err != nil{
			panic(err.Error())
		}
		tpl.ExecuteTemplate(w, "page-login.html", "")
	
	}else{

	tpl, err := template.ParseFiles("static/chuyentrang.html", "static/chuyentrang.html")
	if err != nil {
		panic(err.Error())
	}
	tpl.ExecuteTemplate(w, "chuyentrang.html", "")
}
}
}

func themloaisanpham(w http.ResponseWriter, r *http.Request) {
	session,_ := store.Get(r,"session")
  
	var SS  int
	
	
	if session.Values["iduser"]!=nil{
	  var S= session.Values["iduser"]
	   SS = S.(int)
	  
	
	   
	}else{SS=0
  
	}
  
	selectloaitaikhoan,err := db1.Query("select loaitaikhoan from taikhoan where idtaikhoan='"+strconv.Itoa(SS)+"'")
  
	if err!=nil{
	  panic(err.Error())
	}
	var loaitaikhoanxacnhan int
	for selectloaitaikhoan.Next(){
	 err = selectloaitaikhoan.Scan(&loaitaikhoanxacnhan)
	 if err!=nil{
	  panic(err.Error())
	}
	}
  if loaitaikhoanxacnhan !=3 && SS !=0 {
  
  
	var mang []loaisanpham
	if r.Method == "POST" && r.FormValue("chinhsua") != "1" {
		tensanpham := r.FormValue("txttenloaisp")
		mota := r.FormValue("txtmota")
		insert, err := db1.Query("INSERT INTO `loaisanpham` VALUES ('" + string("") + "','" + mota + "','" + tensanpham + "')")
		if err != nil {
			panic(err.Error())
		}
		insert.Close()

	}
	if r.Method == "POST" && r.FormValue("chinhsua") == "1" {
		tensanpham := r.FormValue("txttenloaisp")
		mota := r.FormValue("txtmota")
		insert, err := db1.Query("UPDATE `loaisanpham` SET tenloaisanpham='" + mota + "', mota='" + tensanpham + "' where idlsp='" + r.FormValue("idsp") + "'")
		if err != nil {
			panic(err.Error())
		}
	
		insert.Close()

	}
	if r.Method == "GET" {

		var idspd = r.FormValue("haha")
		delete, err := db1.Query("delete from loaisanpham where idlsp='" + idspd + "'")
		if err != nil {
			panic(err.Error())
		}
		delete.Close()
	}
	if r.Method == "GET" && r.FormValue("search") != "" {
		sr := r.FormValue("search")
		rows, err := db1.Query("select * from loaisanpham  ")
		for rows.Next() {

			var sanpham1 loaisanpham
			err = rows.Scan(&sanpham1.Idloaisp, &sanpham1.Tenloaisanpham, &sanpham1.Mota)
			if err != nil {
				panic(err.Error())
			}
			if 0 < strings.Count(sanpham1.Tenloaisanpham, sr) {
				mang = append(mang, sanpham1)
			}

		}
		rows.Close()
	} else {

		rows, err := db1.Query("select * from loaisanpham  ")
		for rows.Next() {
			var sanpham1 loaisanpham
			err = rows.Scan(&sanpham1.Idloaisp, &sanpham1.Tenloaisanpham, &sanpham1.Mota)
			if err != nil {
				panic(err.Error())
			}

			mang = append(mang, sanpham1)

		}
		rows.Close()
	}
	tpl, err := template.ParseFiles("static/quanlyloaisanpham.html", "static/quanly.html")
	if err != nil {
		panic(err.Error())
	}
	tpl.ExecuteTemplate(w, "quanlyloaisanpham.html", struct {
		Loaisanpham []loaisanpham
	}{mang})
	}else{

		if SS ==0 {

			tpl,err:= template.New("static/page-login.html").Funcs(template.FuncMap{"tinhgiasale":tinhgiasale,"add": add,"floatchuoi":floatchuoi}).ParseFiles("static/page-login.html")
			if err != nil{
				panic(err.Error())
			}
			tpl.ExecuteTemplate(w, "page-login.html", "")
		
		}else{
		tpl, err := template.ParseFiles("static/chuyentrang.html", "static/chuyentrang.html")
		if err != nil {
			panic(err.Error())
		}
		tpl.ExecuteTemplate(w, "chuyentrang.html", "")
	  }
	}
}



func quanlyuser(w http.ResponseWriter, r *http.Request) {
	session,_ := store.Get(r,"session")
  
	var SS  int
	
	
	if session.Values["iduser"]!=nil{
	  var S= session.Values["iduser"]
	   SS = S.(int)
	  
	
	   
	}else{SS=0
  
	}
  
	selectloaitaikhoan,err := db1.Query("select loaitaikhoan from taikhoan where idtaikhoan='"+strconv.Itoa(SS)+"'")
  
	if err!=nil{
	  panic(err.Error())
	}
	var loaitaikhoanxacnhan int
	for selectloaitaikhoan.Next(){
	 err = selectloaitaikhoan.Scan(&loaitaikhoanxacnhan)
	 if err!=nil{
	  panic(err.Error())
	}
	}
  if loaitaikhoanxacnhan !=3 && loaitaikhoanxacnhan !=2 && SS !=0{
  
  
  
	var mang []loaitaikhoan

	var mangtaikhoan []taikhoan
		rows, err := db1.Query("select * from loaitaikhoan  ")
		for rows.Next() {
			var loaitk loaitaikhoan
			err = rows.Scan(&loaitk.Maloaitaikhoan,&loaitk.Tenloai)
			if err != nil {
				panic(err.Error())
			}

			mang = append(mang, loaitk)

		}

		rows.Close()
		taikhoans,err := db1.Query("SELECT * FROM `taikhoan` ")
		if err != nil{
			panic(err.Error())
		}
		
		for taikhoans.Next(){
		
		   var tk taikhoan
		  err = taikhoans.Scan(&tk.Idtaikhoan,&tk.Tennguoidung,&tk.Ho,&tk.Matkhau,&tk.Email,&tk.Loaitaikhoan)
	 
		  mangtaikhoan = append(mangtaikhoan, tk)
	   if err != nil {
		   panic(err.Error())
	   }
		}
		taikhoans.Close()
	
   if r.Method=="POST" &&r.FormValue("ten")!=""{
	dangkytaikhoan,err := db1.Query("insert into taikhoan values('','"+r.FormValue("ten")+"','"+r.FormValue("ho")+"','"+r.FormValue("matkhau")+"','"+r.FormValue("email")+"','"+r.FormValue("loaitaikhoan")+"')")
	if err != nil {
		panic(err.Error())
		dangkytaikhoan.Close()
	}
   }
   if r.Method == "GET" {

	var idspd = r.FormValue("haha")
	delete, err := db1.Query("delete from taikhoan where idtaikhoan='" + idspd + "'")
	if err != nil {
		panic(err.Error())
	}
	delete.Close()
}
if r.Method == "POST" && r.FormValue("idtaikhoan") == "1" {

	insert, err := db1.Query("UPDATE `taikhoan` SET tenkhachhang='" +  r.FormValue("ten")  + "', hokhachhang='" +  r.FormValue("ho")  +"',matkhau='"+ r.FormValue("matkhau") +"',email='"+ r.FormValue("email")+"',loaitaikhoan='"+ r.FormValue("loaitaikhoan")+ "' where idtaikhoan='" + r.FormValue("idtaikhoan") + "'")
	if err != nil {
		panic(err.Error())
	}

	insert.Close()

}
	tpl, err := template.ParseFiles("static/quanlyuser.html", "static/quanly.html")
	if err != nil {
		panic(err.Error())
	}
	tpl.ExecuteTemplate(w, "quanlyuser.html", struct {
		Loaitaikhoan []loaitaikhoan
		Taikhoan []taikhoan
	}{mang,mangtaikhoan})
	}else{

		if SS ==0 {

			tpl,err:= template.New("static/page-login.html").Funcs(template.FuncMap{"tinhgiasale":tinhgiasale,"add": add,"floatchuoi":floatchuoi}).ParseFiles("static/page-login.html")
			if err != nil{
				panic(err.Error())
			}
			tpl.ExecuteTemplate(w, "page-login.html", "")
		
		}else{
		tpl, err := template.ParseFiles("static/chuyentrang.html", "static/chuyentrang.html")
		if err != nil {
			panic(err.Error())
		}
		tpl.ExecuteTemplate(w, "chuyentrang.html", "")
	  }
	}
}


type noidungtemp struct{
	Ten string
	Noidung string
	Giatri int
}
func quanlythongke(w http.ResponseWriter, r *http.Request) {

	session,_ := store.Get(r,"session")

	var SS  int
	
	
	if session.Values["iduser"]!=nil{
	var S= session.Values["iduser"]
	SS = S.(int)
	
	
	
	}else{SS=0

	}

	selectloaitaikhoan,err := db1.Query("select loaitaikhoan from taikhoan where idtaikhoan='"+strconv.Itoa(SS)+"'")

	if err!=nil{
	panic(err.Error())
	}
	var loaitaikhoanxacnhan int
	for selectloaitaikhoan.Next(){
	err = selectloaitaikhoan.Scan(&loaitaikhoanxacnhan)
	if err!=nil{
	panic(err.Error())
	}
	}
if loaitaikhoanxacnhan !=3 && SS !=0 {



	var demsotaikhoan int
	var demtongip int
	var demdonhang int
	var tongtien int
	var tongsanpham int
	var tongsanphamdaban int
	var tongsonguoitruycapmoi int
	tongtien=0
	demdonhang=0
		taikhoans,err := db1.Query("SELECT * FROM `taikhoan` where loaitaikhoan=3")
		if err != nil{
			panic(err.Error())
		}
		
		for taikhoans.Next(){
		demsotaikhoan+=1
		
	if err != nil {
		panic(err.Error())
	}
		}
		taikhoans.Close()
		t := time.Now()
		p:=t.Month()
		y:=t.Year()
		

		var mangdonhangtheothang []noidungtemp
		for i := 1; i <= 12; i++ {
		var tongtemp =0
			donhang,err := db1.Query("SELECT madh FROM `donhang` where tinhtrang=1 and MONTH(ngaylap)="+strconv.Itoa(i))
			if err != nil{
				panic(err.Error())
			}


			for donhang.Next(){
				tongtemp+=1
			
		
			}
			donhang.Close()
			var mangtemp noidungtemp
			mangtemp.Ten=strconv.Itoa(i)
			mangtemp.Giatri=tongtemp
			mangdonhangtheothang =append(mangdonhangtheothang,mangtemp)


		}
		var mangdonhangtheonam []noidungtemp

		for i := 10; i >= 1; i-- {
			var tongtemp =0
				donhang,err := db1.Query("SELECT madh FROM `donhang` where tinhtrang=1 and YEAR(ngaylap)="+strconv.Itoa((y+1)-i))
				if err != nil{
					panic(err.Error())
				}
	
	
				for donhang.Next(){
					tongtemp+=1
				
			
				}
				donhang.Close()
				var mangtemp noidungtemp
				mangtemp.Ten=strconv.Itoa((y+1)-i)
				mangtemp.Giatri=tongtemp
				mangdonhangtheonam =append(mangdonhangtheonam,mangtemp)
	
	
			}

		sanpham,err := db1.Query("SELECT * FROM `sanpham`")
		if err != nil{
			panic(err.Error())
		}
		
		for sanpham.Next(){
			tongsanpham+=1
		
	if err != nil {
		panic(err.Error())
	}
		}
		sanpham.Close()
	
		
		
		taikhoanip,err := db1.Query("SELECT MONTH(ngaylap)  FROM `ipnguoidung` ")
		if err != nil{
			panic(err.Error())
		}
		
		for taikhoanip.Next(){

			var ngaylap string
			err = taikhoanip.Scan(&ngaylap)
			if ngaylap == strconv.Itoa(int(p)){
				tongsonguoitruycapmoi+=1
			}
			demtongip+=1
		
	if err != nil {
		panic(err.Error())
	}
		}
		
		taikhoanip.Close()

	
		donhang,err := db1.Query("SELECT madh FROM `donhang` where tinhtrang=1 and MONTH(ngaylap)="+strconv.Itoa(int(p)) )
		if err != nil{
			panic(err.Error())
		}
		
		for donhang.Next(){
			demdonhang+=1

			var madondonhang string
			err = donhang.Scan(&madondonhang)
			hoadon,err := db1.Query("SELECT itemhoadon.tongtien , itemhoadon.soluong FROM itemhoadon,hoadon  WHERE hoadon.mahd=itemhoadon.mahd and hoadon.madh='"+madondonhang+"'")
		
	if err != nil {
		panic(err.Error())
	}
				
	donhang.Close()
		for hoadon.Next(){
			var tien int
			var soluong int
			err = hoadon.Scan(&tien,&soluong)
			tongtien+=tien
			tongsanphamdaban+=soluong

		}
		hoadon.Close()

		}
		donhang.Close()

	

var arrsanpham1xemnhieu [] noidungtemp
		slsanpham1xemnhieu,err := db1.Query("SELECT sanpham.tensanpham ,sanpham1.luotxem FROM sanpham1,sanpham WHERE sanpham1.idsp=sanpham.idsp ORDER by sanpham1.luotxem DESC LIMIT 10")
	if err != nil{
		panic(err.Error())
	}
	for slsanpham1xemnhieu.Next(){
		var sanpham1 noidungtemp
		err = slsanpham1xemnhieu.Scan(&sanpham1.Ten,&sanpham1.Giatri)
		if err != nil {
			panic(err.Error())
		}
		arrsanpham1xemnhieu = append(arrsanpham1xemnhieu,sanpham1)
	}
	slsanpham1xemnhieu.Close()


	var arrsanpham1binhchon [] noidungtemp
		slsanpham1binhchon,err := db1.Query("SELECT sanpham.tensanpham ,sanpham1.luotbinhchon FROM sanpham1,sanpham WHERE sanpham1.idsp=sanpham.idsp ORDER by sanpham1.luotbinhchon DESC LIMIT 10")
	if err != nil{
		panic(err.Error())
	}
	for slsanpham1binhchon.Next(){
		var sanpham1 noidungtemp
		err = slsanpham1binhchon.Scan(&sanpham1.Ten,&sanpham1.Giatri)
		if err != nil {
			panic(err.Error())
		}
		arrsanpham1binhchon = append(arrsanpham1binhchon,sanpham1)
	}
	slsanpham1binhchon.Close()

	var arrluotmua [] noidungtemp
	slsanpham1luotmua,err := db1.Query("SELECT sanpham.tensanpham,itemhoadon.idsanpham FROM `itemhoadon`,sanpham WHERE sanpham.idsp=itemhoadon.idsanpham LIMIT 30")
if err != nil{
	panic(err.Error())
}
var demso int
demso =0
for slsanpham1luotmua.Next(){
	var sanpham1 noidungtemp
	err = slsanpham1luotmua.Scan(&sanpham1.Ten,&sanpham1.Noidung)
	if err != nil {
		panic(err.Error())
	}
var kt int
	for i := 0; i < len(arrluotmua); i++ {
		kt=1
	if strings.Contains(arrluotmua[i].Noidung ,sanpham1.Noidung){
		
		arrluotmua[i].Giatri+=1 
		kt=0
		break
	}
	
	}
	if kt==1{
		
		arrluotmua = append(arrluotmua,sanpham1)
		
	}
	if demso==0{
		arrluotmua = append(arrluotmua,sanpham1)
		demso+=1
	}
	
}

slsanpham1luotmua.Close()



var arrluotbinhluan [] noidungtemp
slsanphambinhluan,err := db1.Query("SELECT sanpham.tensanpham,sanpham.idsp FROM `binhluan`,sanpham WHERE sanpham.idsp=binhluan.idsanpham LIMIT 30")
if err != nil{
panic(err.Error())
}

demso =0
for slsanphambinhluan.Next(){
var sanpham1 noidungtemp
err = slsanphambinhluan.Scan(&sanpham1.Ten,&sanpham1.Noidung)
if err != nil {
	panic(err.Error())
}
slsanphambinhluan.Close()
var kt int
for i := 0; i < len(arrluotbinhluan); i++ {
	kt=1
if strings.Contains(arrluotbinhluan[i].Noidung ,sanpham1.Noidung){
	
	arrluotbinhluan[i].Giatri+=1 
	kt=0
	break
}

}
if kt==1{
	
	arrluotbinhluan = append(arrluotbinhluan,sanpham1)
	
}
if demso==0{
	arrluotbinhluan = append(arrluotbinhluan,sanpham1)
	demso+=1
}

}




arrluotbinhluan=sapxepmang(arrluotbinhluan)

arrluotmua=sapxepmang(arrluotmua)


var mangnguoidungtheothang []noidungtemp

for i := 1; i <= 12; i++ {
taikhoanip,err := db1.Query("SELECT MONTH(ngaylap)   FROM `ipnguoidung` where  MONTH(ngaylap)="+strconv.Itoa(int(i))+"")
		if err != nil{
			panic(err.Error())
		}
		var demtemp int
		for taikhoanip.Next(){

			
			
			demtemp+=1
		
		
		}
		var mm noidungtemp
		mm.Giatri =demtemp 
		mm.Ten=strconv.Itoa(int(i))
		mangnguoidungtheothang = append(mangnguoidungtheothang,mm)
		taikhoanip.Close()

	

	}


	var mangnguoidungtheonam []noidungtemp

	for i := 10; i >= 1; i-- {
		taikhoanip,err := db1.Query("SELECT MONTH(ngaylap)   FROM `ipnguoidung` where  year(ngaylap)="+strconv.Itoa(int(y+1-i))+"")
		if err != nil{
			panic(err.Error())
		}
		var demtemp int
		for taikhoanip.Next(){

			
			
			demtemp+=1
		
		
		}
		var mm noidungtemp
		mm.Giatri =demtemp 
		mm.Ten=strconv.Itoa(int(y+1-i))
		mangnguoidungtheonam = append(mangnguoidungtheonam,mm)
		taikhoanip.Close()
	}

	tpl, err := template.ParseFiles("static/quanlythongke.html", "static/quanly.html")
	if err != nil {
		panic(err.Error())
	}
	tpl.ExecuteTemplate(w, "quanlythongke.html", struct {
	Tongsoluongsanpham int
	Tongsoluongdonhangtrongthang int
	Tongsoluongkhachhang  int
	Tongluongtruycap    int
	Tongtientrongthang int
	Thang   int
	Tongsanphamdabantrongthang int
	Tongsoluongusermoi int
	Mangtheothangdonhang  []noidungtemp
	Mangtheonamdonhang  []noidungtemp
	Sanphamxemnhieunhat []noidungtemp
	Sanphambinhchon []noidungtemp
	Sanphammuanhieunhat []noidungtemp
	Sanphambinhluannhieunhat []noidungtemp

	Nguoidungtheothang []noidungtemp

	Nguoidungtheonam []noidungtemp
	}{tongsanpham,demdonhang,demsotaikhoan,demtongip,tongtien,int(p),tongsanphamdaban,tongsonguoitruycapmoi,mangdonhangtheothang,mangdonhangtheonam,arrsanpham1xemnhieu,arrsanpham1binhchon,arrluotmua,arrluotbinhluan,mangnguoidungtheothang,mangnguoidungtheonam})

	}else{

		if SS ==0 {

			tpl,err:= template.New("static/page-login.html").Funcs(template.FuncMap{"tinhgiasale":tinhgiasale,"add": add,"floatchuoi":floatchuoi}).ParseFiles("static/page-login.html")
			if err != nil{
				panic(err.Error())
			}
			tpl.ExecuteTemplate(w, "page-login.html", "")
		
		}else{
		tpl, err := template.ParseFiles("static/chuyentrang.html", "static/chuyentrang.html")
		if err != nil {
			panic(err.Error())
		}
		tpl.ExecuteTemplate(w, "chuyentrang.html", "")
	}
	}
}
func sapxepmang (a [] noidungtemp)  [] noidungtemp{
	var temp noidungtemp
	for i := 0; i < len(a)-1; i++ {

		for j := 0; j < len(a)-1-i; j++ {
			if a[j].Giatri < a[j+1].Giatri{
				temp	 =a[j]
				a[j]=a[j+1]
				a[j+1]=temp
			}
		}
	}
	return a
}

func quanlybinhluan(w http.ResponseWriter, r *http.Request) {
	session,_ := store.Get(r,"session")
	var SS  int
  
  
	if session.Values["iduser"]!=nil{
	  var S= session.Values["iduser"]
	   SS = S.(int)
	  
	
	   
	}else{SS=0
  
	}
  
	selectloaitaikhoan,err := db1.Query("select loaitaikhoan from taikhoan where idtaikhoan='"+strconv.Itoa(SS)+"'")
  
	if err!=nil{
	  panic(err.Error())
	}
	var loaitaikhoanxacnhan int
	for selectloaitaikhoan.Next(){
	 err = selectloaitaikhoan.Scan(&loaitaikhoanxacnhan)
	 if err!=nil{
	  panic(err.Error())
	}
	}
  if loaitaikhoanxacnhan !=3 && SS !=0 {
  
  
  
	var mang []binhluan
	if r.Method == "POST" && r.FormValue("chinhsua") != "1" {
	
		var a int
		a=0
		
    	if session.Values["username"]!=nil{

	    var S1= session.Values["iduser"]
	    a = S1.(int)
	   }else{
		a=0
	   }
		taikhoans,err := db1.Query("SELECT * FROM `taikhoan` WHERE idtaikhoan='"+strconv.Itoa(a) +"' ")
		if err != nil{
			panic(err.Error())
		}
		
		for taikhoans.Next(){
		
		   var tk taikhoan
		  err = taikhoans.Scan(&tk.Idtaikhoan,&tk.Tennguoidung,&tk.Ho,&tk.Matkhau,&tk.Email,&tk.Loaitaikhoan)
	
		  dt := time.Now()
		  sl,err := db1.Query("insert into itembinhluan values('','"+r.FormValue("noidung")+ "','"+tk.Tennguoidung+"','" +tk.Email+"','0"+"','"+ dt.Format("2006-01-02")+"','"+r.FormValue("Idbinhluan")+"','"+strconv.Itoa(a)+"')")
		  if err != nil{
			  panic(err.Error())
		  }
		  sl.Close()
		
	   if err != nil {
		   panic(err.Error())
	   }
		}
		taikhoans.Close()



	

	}
	if r.Method == "POST" && r.FormValue("chinhsua") == "1" {
		tensanpham := r.FormValue("txttenloaisp")
		mota := r.FormValue("txtmota")
		insert, err := db1.Query("UPDATE `loaisanpham` SET tenloaisanpham='" + mota + "', mota='" + tensanpham + "' where idlsp='" + r.FormValue("idsp") + "'")
		if err != nil {
			panic(err.Error())
		}
	
		insert.Close()

	}
	if r.Method == "GET" {

		var idspd = r.FormValue("haha")
		delete, err := db1.Query("delete from binhluan where mabinhluan='" + idspd + "'")
		if err != nil {
			panic(err.Error())
		}
		delete.Close()
	}




		rows, err := db1.Query("SELECT * FROM `binhluan` WHERE 1 ORDER by ngaycapnhat DESC ")
		for rows.Next() {
			var binhluan binhluan
			err = rows.Scan(&binhluan.Idbinhluan, &binhluan.Noidung, &binhluan.Tenkhachhang,&binhluan.Email,&binhluan.Sdt,&binhluan.Ngaycapnhat,&binhluan.Idsanpham,&binhluan.Iduser)
			if err != nil {
				panic(err.Error())
			}

			mang = append(mang, binhluan)

	}
	rows.Close()
var mang1 []sanpham
	row, err := db1.Query("select * from sanpham  ")
	for row.Next() {
		var sanpham1 sanpham
		err = row.Scan(&sanpham1.Idsp, &sanpham1.Tensanpham, &sanpham1.Loaisanpham, &sanpham1.Gia, &sanpham1.Mota, &sanpham1.Hinhanh1, &sanpham1.Hinhanh2, &sanpham1.Hinhanh3)
		if err != nil {
			panic(err.Error())
		}

		mang1 = append(mang1, sanpham1)

	}

	row.Close()
	tpl, err := template.ParseFiles("static/quanlybinhluan.html", "static/quanly.html")
	if err != nil {
		panic(err.Error())
	}
	tpl.ExecuteTemplate(w, "quanlybinhluan.html", struct {
		Binhluan []binhluan
        Sanpham  []sanpham
	}{mang,mang1})
	}else{
		
	if SS ==0 {

		tpl,err:= template.New("static/page-login.html").Funcs(template.FuncMap{"tinhgiasale":tinhgiasale,"add": add,"floatchuoi":floatchuoi}).ParseFiles("static/page-login.html")
		if err != nil{
			panic(err.Error())
		}
		tpl.ExecuteTemplate(w, "page-login.html", "")
	
	}else{

		tpl, err := template.ParseFiles("static/chuyentrang.html", "static/chuyentrang.html")
		if err != nil {
			panic(err.Error())
		}
		tpl.ExecuteTemplate(w, "chuyentrang.html", "")
	  }
	}
}


func quanlydonhang(w http.ResponseWriter, r *http.Request) {
	session,_ := store.Get(r,"session")
  
	var SS  int
	
	
	if session.Values["iduser"]!=nil{
	  var S= session.Values["iduser"]
	   SS = S.(int)
	  
	
	   
	}else{SS=0
  
	}
  
	selectloaitaikhoan,err := db1.Query("select loaitaikhoan from taikhoan where idtaikhoan='"+strconv.Itoa(SS)+"'")
  
	if err!=nil{
	  panic(err.Error())
	}
	var loaitaikhoanxacnhan int
	for selectloaitaikhoan.Next(){
	 err = selectloaitaikhoan.Scan(&loaitaikhoanxacnhan)
	 if err!=nil{
	  panic(err.Error())
	}
	}
  if loaitaikhoanxacnhan !=3 && SS !=0 {
  
  
  
	var arrhoadon [] structhoadon

	hoadon1,err := db1.Query("SELECT * FROM `hoadon` where 1")
	if err != nil{
		panic(err.Error())
	}
	for hoadon1.Next(){
		var hoadon structhoadon
		err = hoadon1.Scan(&hoadon.Mahd,&hoadon.Madh,&hoadon.Iduser,&hoadon.Ship)
		if err != nil {
			panic(err.Error())
		}
		arrhoadon = append(arrhoadon,hoadon)
	}
	if err != nil {
		panic(err.Error())
	}
	hoadon1.Close()
var arrdonhang[] donhang
	donhang1,err := db1.Query("SELECT * FROM `donhang`")
	if err != nil{
		panic(err.Error())
	}
	for donhang1.Next(){
		var donhang donhang
		err = donhang1.Scan(&donhang.Madh,&donhang.Ten,&donhang.Ho,&donhang.Diachi,&donhang.Email,&donhang.Sdt,&donhang.Huyen,&donhang.TP,&donhang.Tinhtrang,&donhang.Note,&donhang.Ngaylap)
		if err != nil {
			panic(err.Error())
		}
		arrdonhang = append(arrdonhang,donhang)
	}
	donhang1.Close()
	if err != nil {
		panic(err.Error())
	}
	var arritemhoadon       []itemhoadon
	slihd,err := db1.Query("SELECT * FROM `itemhoadon` WHERE 1 ")
	 if err != nil{
		 panic(err.Error())
	 }
	 for slihd.Next(){
		 var itemhoadon itemhoadon
		 err = slihd.Scan(&itemhoadon.Iditemhoadon,&itemhoadon.Idsanpham,&itemhoadon.Soluong,&itemhoadon.Tongtien,&itemhoadon.Mahd)
		 if err != nil {
			 panic(err.Error())
		 }
		 arritemhoadon = append(arritemhoadon,itemhoadon)
	 }
	 slihd.Close()
	 if err != nil {
		panic(err.Error())
	}

    if r.Method == "POST" && r.FormValue("madh") != "" {

		updatedh, err := db1.Query("UPDATE `donhang` SET tenkhachhang='" +  r.FormValue("txttenkhachhang") +"',hokhachhang='"+ r.FormValue("txtho")+"',diachi='"+ r.FormValue("txtdiachi")+"',email='"+ r.FormValue("email")+"',sdt='"+ r.FormValue("sdt")+"',huyen='"+ r.FormValue("huyen")+"',thanhpho='"+ r.FormValue("thanhpho")+"',tinhtrang='"+ r.FormValue("tinhtrang")+"',note='"+ r.FormValue("ghichu")+"',ngaylap='"+ r.FormValue("date")+ "' where madh='" + r.FormValue("madh") + "'")
		if err != nil {
			panic(err.Error())
		}

		updatedh.Close()

	}

	if r.Method == "GET" && r.FormValue("tt") != "" {
		
	    
		updatedh, err := db1.Query("UPDATE `donhang` SET tinhtrang='" +  r.FormValue("tt") + "' where madh='" + r.FormValue("iddh") + "'")
		if err != nil {
			panic(err.Error())
		}

		updatedh.Close()
	}
	if r.Method == "GET" && r.FormValue("haha") != "" {
		
	

		var mahd string
		slhd, err := db1.Query("select mahd FROM hoadon    where hoadon.madh='" + r.FormValue("haha") + "'")
		if err != nil {
			panic(err.Error())
		}

		for slhd.Next(){

			err = slhd.Scan(&mahd)
		}
		slhd.Close()

		deleteihd, err := db1.Query("DELETE FROM itemhoadon where itemhoadon.mahd ='" + mahd + "'")
		if err != nil {
			panic(err.Error())
		}

		deleteihd.Close()


		deletehd, err := db1.Query("DELETE FROM hoadon where hoadon.madh ='" +  r.FormValue("haha")  + "'")
		if err != nil {
			panic(err.Error())
		}

		deletehd.Close()

		
		deletedh, err := db1.Query(" delete from  donhang where madh='" + r.FormValue("haha") + "'")
		if err != nil {
			panic(err.Error())
		}

		deletedh.Close()

		deletegiaohang,err:=db1.Query(" delete from  giaohang where madonhang='" + mahd+ "'")
		deletegiaohang.Close()
	}

var mang[] sanpham

	rows, err := db1.Query("select * from sanpham  ")
		for rows.Next() {
			var sanpham1 sanpham
			err = rows.Scan(&sanpham1.Idsp, &sanpham1.Tensanpham, &sanpham1.Loaisanpham, &sanpham1.Gia, &sanpham1.Mota, &sanpham1.Hinhanh1, &sanpham1.Hinhanh2, &sanpham1.Hinhanh3)
			if err != nil {
				panic(err.Error())
			}

			mang = append(mang, sanpham1)

		}
		rows.Close()

type giaohang struct{
	Madonhang int
	Tinhtrang  int
}

selectgiaohang,err :=db1.Query("select * from giaohang")
if err != nil {
	panic(err.Error())
}
var giaohangtmp []giaohang

for selectgiaohang.Next(){
	var giaohang1 giaohang
err = selectgiaohang.Scan(&giaohang1.Madonhang,&giaohang1.Tinhtrang)
if err != nil {
	panic(err.Error())
}
giaohangtmp = append(giaohangtmp,giaohang1)
}
selectgiaohang.Close()

if r.Method == "GET" && r.FormValue("mavanchuyen") != "" {
		
	    
	updatedh, err := db1.Query("UPDATE `giaohang` SET tinhtrang='" +  r.FormValue("tt1") + "' where madonhang='" + r.FormValue("mavanchuyen") + "'")
	if err != nil {
		panic(err.Error())
	}

	updatedh.Close()
}

		tpl,err:= template.New("static/quanlydonhang.html").Funcs(template.FuncMap{"tinhgiasale":tinhgiasale,"add": add,"floatchuoi":floatchuoi}).ParseFiles("static/quanlydonhang.html")

	if err != nil {
		panic(err.Error())
	}
	tpl.ExecuteTemplate(w, "quanlydonhang.html", struct {
		Hoadon [] structhoadon
		Hoadons  []itemhoadon
		Donhang []donhang
		Sanpham [] sanpham
		Giaohang [] giaohang
	}{arrhoadon,arritemhoadon,arrdonhang,mang,giaohangtmp})

}else{

	if SS ==0 {

		tpl,err:= template.New("static/page-login.html").Funcs(template.FuncMap{"tinhgiasale":tinhgiasale,"add": add,"floatchuoi":floatchuoi}).ParseFiles("static/page-login.html")
		if err != nil{
			panic(err.Error())
		}
		tpl.ExecuteTemplate(w, "page-login.html", "")
	
	}else{
		tpl, err := template.ParseFiles("static/chuyentrang.html", "static/chuyentrang.html")
		if err != nil {
			panic(err.Error())
		}
		tpl.ExecuteTemplate(w, "chuyentrang.html", "")
	  }
	}
}


func themthuoctinhloaisanpham(w http.ResponseWriter, r *http.Request) {

	var mang []thuoctinhloaisanpham
	if r.Method == "POST" && r.FormValue("chinhsua") != "1" {
		tensanpham := r.FormValue("txttensanpham")
		loaisanpham := r.FormValue("hahaa")
		mota := r.FormValue("txtmota")
		f := func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsNumber(c)
		}
		var ints = strings.FieldsFunc(loaisanpham, f)

		for i, num := range ints {
         fmt.Print(i)
			
			insert, err := db1.Query("INSERT INTO `thuoctinhloaisanpham` VALUES ('" + string("") + "','" + string(num) + "','" + tensanpham + "','" + mota + "')")
			if err != nil {
				panic(err.Error())
			}
			insert.Close()

		}

	}
	if r.Method == "POST" && r.FormValue("chinhsua") == "1" {
		tensanpham := r.FormValue("txttensanpham")
		loaisanpham := r.FormValue("hahaa")
		mota := r.FormValue("txtmota")
		f := func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsNumber(c)
		}
		var ints = strings.FieldsFunc(loaisanpham, f)
		for i, num := range ints {
			fmt.Println("index:", i)
			
			insert, err := db1.Query("UPDATE `thuoctinhloaisanpham` SET tenthuoctinh='" + tensanpham + "',idloaisanpham='" + string(num) + "', mota='" + mota + "' where idthuoctinh='" + r.FormValue("idsp") + "'")
			if err != nil {
				panic(err.Error())
			}
		
			insert.Close()
			if err != nil {
				panic(err.Error())
			}
			insert.Close()

		}

	}
	if r.Method == "GET" {

		var idspd = r.FormValue("haha")
		delete, err := db1.Query("delete from thuoctinhloaisanpham where idthuoctinh='" + idspd + "'")
		if err != nil {
			panic(err.Error())
		}
		delete.Close()
	}
	if r.Method == "GET" && r.FormValue("search") != "" {
		sr := r.FormValue("search")
		rows, err := db1.Query("select * from thuoctinhloaisanpham  ")
		for rows.Next() {

			var sanpham1 thuoctinhloaisanpham
			err = rows.Scan(&sanpham1.Idthuoctinh, &sanpham1.Idloaisanpham, &sanpham1.Tenthuoctinh, &sanpham1.Mota)
			if err != nil {
				panic(err.Error())
			}
			if 0 < strings.Count(sanpham1.Tenthuoctinh, sr) {
				mang = append(mang, sanpham1)
			}

		}
		rows.Close()
	} else {

		rows, err := db1.Query("select * from thuoctinhloaisanpham  ")
		for rows.Next() {
			var sanpham1 thuoctinhloaisanpham
			err = rows.Scan(&sanpham1.Idthuoctinh, &sanpham1.Idloaisanpham, &sanpham1.Tenthuoctinh, &sanpham1.Mota)
			if err != nil {
				panic(err.Error())
			}

			mang = append(mang, sanpham1)

		}
		rows.Close()
	}
	var loaisanpham1111 []loaisanpham
	loaisanpham11, err := db1.Query("select * from loaisanpham  ")
	for loaisanpham11.Next() {
		var sanpham1 loaisanpham
		err = loaisanpham11.Scan(&sanpham1.Idloaisp, &sanpham1.Tenloaisanpham, &sanpham1.Mota)
		if err != nil {
			panic(err.Error())
		}
		loaisanpham1111 = append(loaisanpham1111, sanpham1)

	}
	loaisanpham11.Close()
	tpl, err := template.ParseFiles("static/thuoctinhtrongloaisanpham.html", "static/quanly.html")
	if err != nil {
		panic(err.Error())
	}
	tpl.ExecuteTemplate(w, "thuoctinhtrongloaisanpham.html", struct {
		Loaisanpham []loaisanpham
		Thuoctinh   []thuoctinhloaisanpham
	}{loaisanpham1111, mang})

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	session,_ := store.Get(r,"session")
	session.Values["nhasanxuat"]=nil 
    session.Values["nhacungcap"]=nil
	session.Values["giacantim"]=nil
	session.Values["sao"]=nil
     session.Save(r,w) 
    var lsp [] loaisanpham
	var  arrSanpham [] sanpham
	var  arrsanpham1 [] sanpham1
	var   arrsanpham1giamgia []sanpham1
    var   arrsanpham1xemnhieu []sanpham1
	slloaisanpham,err := db1.Query("SELECT * FROM `loaisanpham` WHERE 1 ")
	if err != nil{
		panic(err.Error())
	}
	for slloaisanpham.Next(){
		var loaisanpham loaisanpham
		err = slloaisanpham.Scan(&loaisanpham.Idloaisp,&loaisanpham.Tenloaisanpham,&loaisanpham.Mota)
		if err != nil {
			panic(err.Error())
		}
		lsp = append(lsp,loaisanpham)
	}
	if err != nil {
		panic(err.Error())
	}
	slloaisanpham.Close()
	
	slsanpham,err := db1.Query("SELECT * FROM `sanpham` WHERE 1 ")
	if err != nil{
		panic(err.Error())
	}
	for slsanpham.Next(){
		var sanpham sanpham
		err = slsanpham.Scan(&sanpham.Idsp, &sanpham.Tensanpham, &sanpham.Loaisanpham, &sanpham.Gia, &sanpham.Mota, &sanpham.Hinhanh1, &sanpham.Hinhanh2, &sanpham.Hinhanh3)
		if err != nil {
			panic(err.Error())
		}
		arrSanpham = append(arrSanpham,sanpham)
	}
	slsanpham.Close()

	slsanpham1,err := db1.Query("SELECT * FROM `sanpham1` WHERE 1 ORDER BY ngaycapnhat asc LIMIT 15 ")
	if err != nil{
		panic(err.Error())
	}
	for slsanpham1.Next(){
		var sanpham1 sanpham1

		
		
		err = slsanpham1.Scan(&sanpham1.Giamgia, &sanpham1.Manhasanxuat, &sanpham1.Manhacungcap, &sanpham1.Ngaycapnhat, &sanpham1.Luotxem, &sanpham1.Luotbinhchon, &sanpham1.Soluongton, &sanpham1.Cauhinh,&sanpham1.Idsp1)
		if err != nil {
			panic(err.Error())
		}
	
		arrsanpham1 = append(arrsanpham1,sanpham1)
	}

	slsanpham1.Close()


	slsanpham1giamgia,err := db1.Query("SELECT * FROM `sanpham1` WHERE 1 ORDER BY giamgia asc LIMIT 15 ")
	if err != nil{
		panic(err.Error())
	}
	for slsanpham1giamgia.Next(){
		var sanpham1 sanpham1
		err = slsanpham1giamgia.Scan(&sanpham1.Giamgia, &sanpham1.Manhasanxuat, &sanpham1.Manhacungcap, &sanpham1.Ngaycapnhat, &sanpham1.Luotxem, &sanpham1.Luotbinhchon, &sanpham1.Soluongton, &sanpham1.Cauhinh,&sanpham1.Idsp1)
		if err != nil {
			panic(err.Error())
		}
		arrsanpham1giamgia = append(arrsanpham1giamgia,sanpham1)
	}
	slsanpham1giamgia.Close()

	slsanpham1xemnhieu,err := db1.Query("SELECT * FROM `sanpham1` WHERE 1 ORDER BY luotxem asc LIMIT 15  ")
	if err != nil{
		panic(err.Error())
	}
	for slsanpham1xemnhieu.Next(){
		var sanpham1 sanpham1
		err = slsanpham1xemnhieu.Scan(&sanpham1.Giamgia, &sanpham1.Manhasanxuat, &sanpham1.Manhacungcap, &sanpham1.Ngaycapnhat, &sanpham1.Luotxem, &sanpham1.Luotbinhchon, &sanpham1.Soluongton, &sanpham1.Cauhinh,&sanpham1.Idsp1)
		if err != nil {
			panic(err.Error())
		}
		arrsanpham1xemnhieu = append(arrsanpham1xemnhieu,sanpham1)
	}
	slsanpham1xemnhieu.Close()
	
	if err != nil {
		panic(err.Error())
	}
	var SS string
	if session.Values["username"]!=nil{

	var S= session.Values["username"]
	SS = S.(string)
	}else{
		SS=""
	}
	tpl,err:= template.New("static/index.html").Funcs(template.FuncMap{"tinhgiasale":tinhgiasale,"add": add,"floatchuoi":floatchuoi}).ParseFiles("static/index.html")
	tpl.ExecuteTemplate(w, "index.html", struct {
		Loaisanpham []loaisanpham
		Sanpham   []sanpham
		Sanpham1  []sanpham1
		Sanpham1giamgia  []sanpham1
		Sanpham1xemnhieu  []sanpham1
		Username string
	}{lsp,arrSanpham,arrsanpham1,arrsanpham1giamgia,arrsanpham1xemnhieu,SS})
	
}



func loadheader(w http.ResponseWriter, r *http.Request) {
	session,_ := store.Get(r,"session")
	
    var lsp [] loaisanpham

	slloaisanpham,err := db1.Query("SELECT * FROM `loaisanpham` WHERE 1 ")
	if err != nil{
		panic(err.Error())
	}
	for slloaisanpham.Next(){
		var loaisanpham loaisanpham
		err = slloaisanpham.Scan(&loaisanpham.Idloaisp,&loaisanpham.Tenloaisanpham,&loaisanpham.Mota)
		if err != nil {
			panic(err.Error())
		}
		lsp = append(lsp,loaisanpham)
	}
	if err != nil {
		panic(err.Error())
	}
	slloaisanpham.Close()

	var SS string
	if session.Values["username"]!=nil{

	var S= session.Values["username"]
	SS = S.(string)
	}else{
		SS=""
	}
	tpl,err:= template.New("static/header.html").Funcs(template.FuncMap{"tinhgiasale":tinhgiasale,"add": add,"floatchuoi":floatchuoi}).ParseFiles("static/header.html")
	tpl.ExecuteTemplate(w, "header.html", struct {
		Loaisanpham []loaisanpham
		Username string
	}{lsp,SS})
	
}

func chitietthongtinsanpham(w http.ResponseWriter, r *http.Request){
   var idsanpham = r.FormValue("idsp")
   var  luotxemsanpham int
 

	slluotxem,err := db1.Query("SELECT luotxem FROM `sanpham1` WHERE idsp='"+string(idsanpham)+"'")
	
	for slluotxem.Next(){
		err = slluotxem.Scan(&luotxemsanpham)
		if err != nil {
			panic(err.Error())
		}
		
	}
	luotxemsanpham=luotxemsanpham+1
		updatelx,err :=db1.Query("UPDATE `sanpham1` SET luotxem='"+strconv.Itoa(luotxemsanpham)+"' WHERE idsp='"+string(idsanpham)+"'")
		updatelx.Close()
		slluotxem.Close()
   session,_ := store.Get(r,"session")
   session.Values["nhasanxuat"]=nil 
   session.Values["nhacungcap"]=nil
   session.Values["giacantim"]=nil
	session.Save(r,w)

	
   var arrSanpham[]sanpham
   var arrsanpham1[]sanpham1
   var lsp       []loaisanpham
   var arrsanphamlienquan []sanpham
   slloaisanpham,err := db1.Query("SELECT * FROM `loaisanpham` WHERE 1 ")
	if err != nil{
		panic(err.Error())
	}
	for slloaisanpham.Next(){
		var loaisanpham loaisanpham
		err = slloaisanpham.Scan(&loaisanpham.Idloaisp,&loaisanpham.Tenloaisanpham,&loaisanpham.Mota)
		if err != nil {
			panic(err.Error())
		}
		lsp = append(lsp,loaisanpham)
	}
	if err != nil {
		panic(err.Error())
	}
	slloaisanpham.Close()
   slsanpham,err := db1.Query("SELECT * FROM `sanpham` WHERE idsp='"+string(idsanpham)+"'")
	if err != nil{
		panic(err.Error())
	}
	for slsanpham.Next(){
		var sanpham sanpham
		err = slsanpham.Scan(&sanpham.Idsp, &sanpham.Tensanpham, &sanpham.Loaisanpham, &sanpham.Gia, &sanpham.Mota, &sanpham.Hinhanh1, &sanpham.Hinhanh2, &sanpham.Hinhanh3)
		if err != nil {
			panic(err.Error())
		}
		arrSanpham = append(arrSanpham,sanpham)
	}
	slsanpham.Close()

	slsanpham1,err := db1.Query("SELECT * FROM `sanpham1` WHERE 1")
	if err != nil{
		panic(err.Error())
	}
	for slsanpham1.Next(){
		var sanpham1 sanpham1
		err = slsanpham1.Scan(&sanpham1.Giamgia, &sanpham1.Manhasanxuat, &sanpham1.Manhacungcap, &sanpham1.Ngaycapnhat, &sanpham1.Luotxem, &sanpham1.Luotbinhchon, &sanpham1.Soluongton, &sanpham1.Cauhinh,&sanpham1.Idsp1)
		if err != nil {
			panic(err.Error())
		}
		arrsanpham1 = append(arrsanpham1,sanpham1)
	}
	slsanpham1.Close()
	slsanphamlq,err := db1.Query("SELECT * FROM `sanpham` WHERE 1")
	if err != nil{
		panic(err.Error())
	}
	for slsanphamlq.Next(){
		var sanpham sanpham
		err = slsanphamlq.Scan(&sanpham.Idsp, &sanpham.Tensanpham, &sanpham.Loaisanpham, &sanpham.Gia, &sanpham.Mota, &sanpham.Hinhanh1, &sanpham.Hinhanh2, &sanpham.Hinhanh3)
		if err != nil {
			panic(err.Error())
		}
		arrsanphamlienquan = append(arrsanphamlienquan,sanpham)
	}
	slsanphamlq.Close()
var Nhasanxuat [] nhasanxuat
var Nhacungcap [] nhacungcap
rows, err := db1.Query("select * from nhasanxuat ")
		for rows.Next() {
			var sanpham1 nhasanxuat
			err = rows.Scan(&sanpham1.Idnhasanxuat, &sanpham1.Tennhasanxuat, &sanpham1.Thongtin, &sanpham1.Logo)
			if err != nil {
				panic(err.Error())
			}
		

			Nhasanxuat = append(Nhasanxuat, sanpham1)

		}
		rows.Close()
		row, err := db1.Query("select * from nhacungcap ")
		for row.Next() {
			var sanpham1 nhacungcap
			err = row.Scan(&sanpham1.Idnhacungcap, &sanpham1.TenNhaCungCap,&sanpham1.Diachi, &sanpham1.Email, &sanpham1.Sdt,&sanpham1.Fax)
			if err != nil {
				panic(err.Error())
			}

			Nhacungcap = append(Nhacungcap, sanpham1)
		}
		row.Close()
	
		var tmp string
		var mangdanhgia []danhgia

		var booldanhgia string
		var username int
		
	if session.Values["username"]!= nil{
      var ttusername =session.Values["iduser"]
		username = ttusername.(int)
		selectkt , err := db1.Query("SELECT * FROM `itemhoadon`,`hoadon` WHERE itemhoadon.idsanpham='"+string(idsanpham)+"' and hoadon.mahd=hoadon.mahd and hoadon.iduser='"+strconv.Itoa(username)+"'")
	   var demkt =0
	   if err != nil{
		panic(err.Error())
	}
		for selectkt.Next(){
		 demkt=1
		
	   }
	   selectkt.Close()
		if demkt ==1{
			booldanhgia="1"
		}else{
			booldanhgia="0"
		}

		
	
      tmp ="1"
	}else{
		tmp="0"
	}
	sldanhgia,err := db1.Query("select * from danhgia where idsanpham='"+string(idsanpham)+"'")
		
		for sldanhgia.Next(){
			var danhgia danhgia
			err = sldanhgia.Scan(&danhgia.Iddanhgia,&danhgia.Noidung,&danhgia.Idsanpham,&danhgia.Iduser,&danhgia.Tenkhachhang,&danhgia.Sao,&danhgia.Ngaycapnhat)
			if err != nil{
				panic(err.Error())
			}
			
			
			mangdanhgia =append(mangdanhgia,danhgia)
		}
		sldanhgia.Close()
	
		
	
		var SS string
		if session.Values["username"]!=nil{
	
		var S= session.Values["username"]
		SS = S.(string)
		}else{
			SS=""
		}
     
		tpl,err:= template.New("static/single-product.html").Funcs(template.FuncMap{"tinhgiasale":tinhgiasale,"add": add,"floatchuoi":floatchuoi}).ParseFiles("static/single-product.html")
		if err != nil {
			panic(err.Error())
		}

	tpl.ExecuteTemplate(w, "single-product.html", struct {
		Loaisanpham []loaisanpham
		Sanpham   []sanpham
		Sanpham1  []sanpham1
		Sanphamlienquan   []sanpham
		Nhasanxuat    []nhasanxuat
		Nhacungcap    []nhacungcap
		User  string
		Danhgia []danhgia
		Damuahang    string
		Username  string
	

	}{lsp,arrSanpham,arrsanpham1,arrsanphamlienquan,Nhasanxuat,Nhacungcap,tmp,mangdanhgia,booldanhgia,SS})
}
func binhluansanpham(w http.ResponseWriter, r *http.Request){
	
	var mangbinhluan []binhluan
	var mangbinhluans []binhluans
	slbinhluan,err := db1.Query("select * from binhluan where idsanpham='"+r.FormValue("idsanpham")+"'")
	if err != nil{
		panic(err.Error())
	}
	var kt string
	session,_ := store.Get(r,"session")
     if  session.Values["username"] !=nil{
		kt="1"
	 }else{
		kt="0"
	 }
	
	for slbinhluan.Next(){
		var danhgia binhluan
		err = slbinhluan.Scan(&danhgia.Idbinhluan,&danhgia.Noidung,&danhgia.Tenkhachhang,&danhgia.Email,&danhgia.Sdt,&danhgia.Ngaycapnhat,&danhgia.Idsanpham,&danhgia.Iduser)
		if err != nil{
			panic(err.Error())
		}
		
	
		mangbinhluan =append(mangbinhluan,danhgia)
	}
	slbinhluan.Close()
	slbinhluans,err := db1.Query("select * from itembinhluan where 1")
	if err != nil{
		panic(err.Error())
	}
	for slbinhluans.Next(){
		var danhgia binhluans
		err = slbinhluans.Scan(&danhgia.Idbinhluans,&danhgia.Noidung,&danhgia.Tenkhachhang,&danhgia.Email,&danhgia.Sdt,&danhgia.Ngaycapnhat,&danhgia.Idbinhluan,&danhgia.Iduser)
		if err != nil{
			panic(err.Error())
		}
		
		
		mangbinhluans =append(mangbinhluans,danhgia)
	}
	slbinhluans.Close()
	var  arranhcomment  []anhcomment
	anhcmt, err := db1.Query("select * from anhcomment ")
	for anhcmt.Next() {
		var anh anhcomment
		err = anhcmt.Scan(&anh.Idanh,&anh.Loaianh,&anh.Duongdan)
	
		if err != nil {
			panic(err.Error())
		}

		arranhcomment = append(arranhcomment, anh)
	}
	anhcmt.Close()

	tpl, err := template.ParseFiles("static/binhluan.html")
	if err != nil {
		panic(err.Error())
	}
	tpl.ExecuteTemplate(w, "binhluan.html", struct {
		Binhluan []binhluan
		Binhluans []binhluans
		Kt string
        Anhbinhluan [] anhcomment
	}{mangbinhluan ,mangbinhluans,kt,arranhcomment})
}


func binhluansanphams(w http.ResponseWriter, r *http.Request){
	
	var mangbinhluan []binhluan
	var mangbinhluans []binhluans
	slbinhluan,err := db1.Query("select * from binhluan where idsanpham='"+r.FormValue("idsanpham")+"'")
	if err != nil{
		panic(err.Error())
	}
	var kt string
	session,_ := store.Get(r,"session")
     if  session.Values["username"] !=nil{
		kt="1"
	 }else{
		kt="0"
	 }
	
	for slbinhluan.Next(){
		var danhgia binhluan
		err = slbinhluan.Scan(&danhgia.Idbinhluan,&danhgia.Noidung,&danhgia.Tenkhachhang,&danhgia.Email,&danhgia.Sdt,&danhgia.Ngaycapnhat,&danhgia.Idsanpham,&danhgia.Iduser)
		if err != nil{
			panic(err.Error())
		}
		
		
		mangbinhluan =append(mangbinhluan,danhgia)
	}
	slbinhluan.Close()
	slbinhluans,err := db1.Query("select * from itembinhluan where  idbinhluan='"+r.FormValue("idbinhluan")+"'")
	if err != nil{
		panic(err.Error())
	}
	for slbinhluans.Next(){
		var danhgia binhluans
		err = slbinhluans.Scan(&danhgia.Idbinhluans,&danhgia.Noidung,&danhgia.Tenkhachhang,&danhgia.Email,&danhgia.Sdt,&danhgia.Ngaycapnhat,&danhgia.Idbinhluan,&danhgia.Iduser)
		if err != nil{
			panic(err.Error())
		}
		
	
		mangbinhluans =append(mangbinhluans,danhgia)
	}
	slbinhluans.Close()
  var idbl=r.FormValue("idsanpham")
	tpl, err := template.ParseFiles("static/itembinhluan.html")
	if err != nil {
		panic(err.Error())
	}
	tpl.ExecuteTemplate(w, "itembinhluan.html", struct {
		Binhluan []binhluan
		Binhluans []binhluans
		Kt string
		Idbinhluan string

	}{mangbinhluan ,mangbinhluans,kt,idbl})
}
func thongtinloaisanpham(w http.ResponseWriter, r *http.Request){
	session,_ := store.Get(r,"session")
	var nhasanxuatss string
	if r.FormValue("nhasanxuat") != ""{
	
		session.Values["nhasanxuat"]=r.FormValue("nhasanxuat")
		session.Save(r,w) 

	}
  var sapxeptheo string
  sapxeptheo=""
  
	if r.FormValue("sapxep") != ""{
	
		session.Values["sapxep"]=r.FormValue("sapxep")
		session.Save(r,w) 
		

	}

	if session.Values["sapxep"] != nil{
		if session.Values["sapxep"] == "0"{
			sapxeptheo =""
		}
		if session.Values["sapxep"] == "1"{
			sapxeptheo ="  ORDER BY `sanpham`.`tensanpham` ASC"
		}
		if session.Values["sapxep"] == "2"{
			sapxeptheo ="  ORDER BY `sanpham`.`tensanpham` DESC"
		}
		if session.Values["sapxep"] == "3"{
			sapxeptheo ="  ORDER BY `sanpham`.`gia` ASC"
		}
		if session.Values["sapxep"] == "4"{
			sapxeptheo ="  ORDER BY `sanpham`.`gia` DESC"
		}
	}
     if r.FormValue("reset") != ""{
		session.Values["nhasanxuat"]=nil 
		session.Values["nhacungcap"]=nil
		session.Values["giacantim"]=nil
		session.Values["sao"]=nil
		 session.Save(r,w) 
	 }



	if session.Values["nhasanxuat"] != nil{
		var tempsx string
		var tmp =session.Values["nhasanxuat"]
		tempsx = tmp.(string)
    	f := func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsNumber(c)
		}
		nhasanxuatss=""
		var ints = strings.FieldsFunc(tempsx, f)

		for i, num := range ints {
		  
			
			if i == 0{
				nhasanxuatss+="and  manhasanxuat='"+string(num)+"'"
			}else{
				nhasanxuatss+="or manhasanxuat='"+string(num)+"'"
			}

		}
	}
   
	var giatim string
	if r.FormValue("min")!=""{
		giatim+=" and gia>="+r.FormValue("min")
		session.Values["giacantim"]=giatim
		session.Save(r,w) 
	}else{
		giatim+=" and gia>=0"
	
	}
	if r.FormValue("max")!=""{
		
		giatim+=" and gia<="+r.FormValue("max")
		session.Values["giacantim"]=giatim
	session.Save(r,w) 
	}
   if session.Values["giacantim"]!=nil{
	var temp string
	var tmp = session.Values["giacantim"]
	temp = tmp.(string)
	
	giatim=temp
	
   }
	var nhacungcapss string
	if r.FormValue("nhacungcap") != ""{
	
		session.Values["nhacungcap"]=r.FormValue("nhacungcap")
		session.Save(r,w) 
	}
   
   
	if session.Values["nhacungcap"] != nil{
		var tempsx string
		var tmp =session.Values["nhacungcap"]
		tempsx = tmp.(string)
    	f := func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsNumber(c)
		}
		nhacungcapss=""
		var ints = strings.FieldsFunc(tempsx, f)

		for i, num := range ints {
		  
			
			if i == 0{
				nhasanxuatss+="and  manhacungcap='"+string(num)+"'"
			}else{
				nhasanxuatss+="or manhacungcap='"+string(num)+"'"
			}

		}
	}



	var saoss string
	if r.FormValue("sao") != ""{
	
		session.Values["sao"]=r.FormValue("sao")
		session.Save(r,w) 
	}
   
	saoss=""
	if session.Values["sao"] != nil{
		var tempsx string
		var tmp =session.Values["sao"]
		tempsx = tmp.(string)
    	saoss +="and luotbinhchon="+tempsx
	}
	
	var idloaisanpham = r.FormValue("idloaisanpham")
	if idloaisanpham == "0"{
		idloaisanpham=""
	}

	if idloaisanpham == "NaN"{
		idloaisanpham=""
		
	}
	var page =  r.FormValue("page")
	if page ==""{
		page="1"
		
	}
	var pageso int64
	if i, err := strconv.ParseInt(page, 10, 64); err == nil {
		pageso =i
	}
	pagetp,err :=strconv.ParseFloat(page,64)
	pagetp =  math.Round(pagetp*12 /12)
	
	var arrsanpham1 []sanpham1
	var arrsanphamlienquan []sanpham
	slsanpham1,err := db1.Query("SELECT * FROM `sanpham1` where 1 "+nhacungcapss+nhasanxuatss+saoss)
	
	if err != nil{
		panic(err.Error())
	}
	for slsanpham1.Next(){
		var sanpham1 sanpham1
		err = slsanpham1.Scan(&sanpham1.Giamgia, &sanpham1.Manhasanxuat, &sanpham1.Manhacungcap, &sanpham1.Ngaycapnhat, &sanpham1.Luotxem, &sanpham1.Luotbinhchon, &sanpham1.Soluongton, &sanpham1.Cauhinh,&sanpham1.Idsp1)
		if err != nil {
			panic(err.Error())
		}
		arrsanpham1 = append(arrsanpham1,sanpham1)
	}
	var temp=pagetp
    var bien =0
	if pagetp <1.0 {
		temp =1
	}
	slsanpham1.Close()
	var soluongpage =1.0;

	if r.FormValue("q") ==""{
	slsanphamlq,err := db1.Query("SELECT * FROM `sanpham` WHERE loaisanpham='"+string(idloaisanpham)+"'"+giatim+" "+sapxeptheo)
	if err != nil{
		panic(err.Error())
	}
	
	for slsanphamlq.Next(){
		var sanpham sanpham
        if bien >= (int(temp)-1)*12 &&bien <= int(pageso) *12 {
		err = slsanphamlq.Scan(&sanpham.Idsp, &sanpham.Tensanpham, &sanpham.Loaisanpham, &sanpham.Gia, &sanpham.Mota, &sanpham.Hinhanh1, &sanpham.Hinhanh2, &sanpham.Hinhanh3)
		if err != nil {
			panic(err.Error())
		}
		arrsanphamlienquan = append(arrsanphamlienquan,sanpham)
     	}
	
	
		soluongpage+=1.0
		bien +=1
	}
	slsanphamlq.Close()
	
    }else{
	
	if idloaisanpham == ""{
	slsanphamlq,err := db1.Query("SELECT * FROM `sanpham` where 1"+giatim+" "+sapxeptheo)

		for slsanphamlq.Next(){
		var sanpham sanpham
		
		err = slsanphamlq.Scan(&sanpham.Idsp, &sanpham.Tensanpham, &sanpham.Loaisanpham, &sanpham.Gia, &sanpham.Mota, &sanpham.Hinhanh1, &sanpham.Hinhanh2, &sanpham.Hinhanh3)
		if 0 < strings.Count(sanpham.Tensanpham, r.FormValue("q")) {
		if bien >= (int(temp)-1)*12 &&bien <= int(pageso) *12 {
	
		if err != nil {
			panic(err.Error())
		}
		
		
		
		
		
		arrsanphamlienquan = append(arrsanphamlienquan,sanpham)
		
	   
	
		}
		bien +=1
     	}
	
		 soluongpage+=1.0
		 
	}
	slsanphamlq.Close()
	}else{
		slsanphamlq,err := db1.Query("SELECT * FROM `sanpham` WHERE loaisanpham='"+string(idloaisanpham)+"'"+giatim+" "+sapxeptheo)
		for slsanphamlq.Next(){
			var sanpham sanpham
			err = slsanphamlq.Scan(&sanpham.Idsp, &sanpham.Tensanpham, &sanpham.Loaisanpham, &sanpham.Gia, &sanpham.Mota, &sanpham.Hinhanh1, &sanpham.Hinhanh2, &sanpham.Hinhanh3)
			if 0 < strings.Count(sanpham.Tensanpham, r.FormValue("q")) {
			if bien >= (int(temp)-1)*12 &&bien <= int(pageso) *12 {
			
			if err != nil {
				panic(err.Error())
			}
			
			arrsanphamlienquan = append(arrsanphamlienquan,sanpham)
			
			
		
			}
			bien +=1
			 }
			 soluongpage+=1.0
			 
			
		}
		slsanphamlq.Close()
	}
	
	


}
var Nhasanxuat [] nhasanxuat
var Nhacungcap [] nhacungcap
rows1, err := db1.Query("select * from nhasanxuat ")
		for rows1.Next() {
			var sanpham1 nhasanxuat
			err = rows1.Scan(&sanpham1.Idnhasanxuat, &sanpham1.Tennhasanxuat, &sanpham1.Thongtin, &sanpham1.Logo)
			if err != nil {
				panic(err.Error())
			}
			

			Nhasanxuat = append(Nhasanxuat, sanpham1)

		}
		rows1.Close()
		row1, err := db1.Query("select * from nhacungcap ")
		for row1.Next() {
			var sanpham1 nhacungcap
			err = row1.Scan(&sanpham1.Idnhacungcap, &sanpham1.TenNhaCungCap,&sanpham1.Diachi, &sanpham1.Email, &sanpham1.Sdt,&sanpham1.Fax)
			if err != nil {
				panic(err.Error())
			}

			Nhacungcap = append(Nhacungcap, sanpham1)
		}
		row1.Close()

var teslp =soluongpage
	soluongpage = math.Round(soluongpage/12)
	if soluongpage <1.0{
		soluongpage=1
		
	}
	if soluongpage <2.0 && teslp/12>1.07 {
		soluongpage=2
	
	}

	var SS string
	if session.Values["username"]!=nil{

	var S= session.Values["username"]
	SS = S.(string)
	}else{
		SS=""
	}

	tpl,err:= template.New("static/shop-left-sidebar.html").Funcs(template.FuncMap{"tinhgiasale":tinhgiasale,"add": add,"floatchuoi":floatchuoi}).ParseFiles("static/shop-left-sidebar.html")

	if err != nil {
		panic(err.Error())
	}
	var lsp       []loaisanpham
	slloaisanpham,err := db1.Query("SELECT * FROM `loaisanpham` WHERE 1 ")
	 if err != nil{
		 panic(err.Error())
	 }
	 for slloaisanpham.Next(){
		 var loaisanpham loaisanpham
		 err = slloaisanpham.Scan(&loaisanpham.Idloaisp,&loaisanpham.Tenloaisanpham,&loaisanpham.Mota)
		 if err != nil {
			 panic(err.Error())
		 }
		 lsp = append(lsp,loaisanpham)
	 }
	 slloaisanpham.Close()
	 if err != nil {
		 panic(err.Error())
	 }
	 var mangpage []int
	 for i := 1; i <= int(soluongpage); i++ {
		mangpage = append(mangpage,i)
	}
	var Q =r.FormValue("q")
	tpl.ExecuteTemplate(w, "shop-left-sidebar.html", struct {
		Sanpham   []sanpham
		Sanpham1  []sanpham1
		Loaisanpham []loaisanpham
		Soluongpage []int
		Pagehientai string
		Loaisp string
		Q string
		Nhasanxuat [] nhasanxuat
		Nhacungcap  []nhacungcap
		Username   string
	}{arrsanphamlienquan,arrsanpham1,lsp, mangpage,page,idloaisanpham,Q,Nhasanxuat,Nhacungcap,SS})

}
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

var tpl *template.Template

func init() {
	// tpl = template.Must(template.ParseGlob("static/*.html"))
	// tpl, err := template.ParseGlob("static/*.html")
	// fmt.Println(tpl)
	// fmt.Println(err)
}

// }
func themvaogiohang(w http.ResponseWriter, r *http.Request) {
	session,_ := store.Get(r,"session")
	var a string
	if session.Values["username"] !=nil{
      a="1"
	}else{
		a="0"
	}
    if session.Values["idsanpham"] == nil{
		session.Values["idsanpham"]=""
		session.Values["soluong"]=""
	}

	ktsoluong, err := db1.Query("select soluongton from sanpham1 where idsp='"+r.FormValue("idsanpham")+"'")
		if err != nil {
			panic(err.Error())
			}

			var ktthemvao int
    for ktsoluong.Next(){
        var soluongtmps int
		err = ktsoluong.Scan(&soluongtmps)
		i2, err := strconv.ParseInt(r.FormValue("soluong"), 10, 64)
	if err == nil {
	
	}
		if soluongtmps < int(i2){
			ktthemvao=0
		} else{
			ktthemvao=1
		}
	}
	ktsoluong.Close()
	

	if ktthemvao==1{

	if session.Values["idsanpham"]!=nil && a=="0"{
		
	type giohang struct{
		idsanpham  string
		soluong   string
	}
	var mang1 [] string
	var mang2 [] string
	var mang []giohang
	var idsanpham =session.Values["idsanpham"] 
	var stridsanpham string
	var strsoluong string
	strsoluong =""
	stridsanpham = ""
	if idsanpham != nil{
		 stridsanpham = idsanpham.(string)
	
	}
	
	
	var soluong = session.Values["soluong"]
	if soluong != nil{
		strsoluong = soluong.(string)
		
	}
	if strsoluong != ""{
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	var ints = strings.FieldsFunc(stridsanpham, f)

	for i, num := range ints {
		if i != -1{

		}
		mang1 = append(mang1,num)

	}

	ints = strings.FieldsFunc(strsoluong, f)

	for i, num := range ints {

		if i != -1{

		}
		mang2 = append(mang2,num)

	}

	for i := 0; i < len(mang1); i++ {
		var temp =struct {
		idsanpham  string
		soluong   string
	}{mang1[i],mang2[i]}
       mang = append(mang,temp)
	}

}
	if len(mang1) <1{
		session.Values["idsanpham"] = r.FormValue("idsanpham")+";"
		session.Values["soluong"] = r.FormValue("soluong")+";"
		session.Save(r,w) 
	
	
	}else{
		if  strings.Count( stridsanpham,r.FormValue("idsanpham")) ==0 {
			var c=stridsanpham+r.FormValue("idsanpham")+";"
			var d = strsoluong+r.FormValue("soluong")+";"
			session.Values["idsanpham"] = c
			session.Values["soluong"]  = d
			session.Save(r,w) 
		
		}else{
			var str1,str2 string
			for i := 0; i < len(mang1); i++ {
				if r.FormValue("idsanpham")== mang1[i]{
					str1+=mang1[i]+";"
					str2+=r.FormValue("soluong")+";"
				}else{
				   str1+=mang1[i]+";"
				   str2+=mang2[i]+";"
				  
		
				}
			}
			session.Values["idsanpham"] = str1
			session.Values["soluong"]  = str2
			session.Save(r,w)
		
		}
	   
		if r.FormValue("capnhat")!=""{
			var str1,str2 string
			for i := 0; i < len(mang1); i++ {
				if r.FormValue("idsanpham")== mang1[i]{
					str1+=mang1[i]+";"
					str2+=r.FormValue("soluong")+";"
				}else{
				   str1+=mang1[i]+";"
				   str2+=mang2[i]+";"
				  
		
				}
			}
			session.Values["idsanpham"] = str1
			session.Values["soluong"]  = str2
			session.Save(r,w)
		
		}

	}
	}else {
		
		var iduser =session.Values["iduser"]
		var striduser = iduser.(int)
        var idgiohang int
		row, err := db1.Query("select * from giohang where user='"+strconv.Itoa(striduser)+"'")
		if err != nil {
			panic(err.Error())
			}
		var count int
		count =0
			for row.Next() {
				var giohang giohang
			err = row.Scan(&giohang.Idgiohang,&giohang.Username)
			idgiohang =giohang.Idgiohang
			if err != nil {
				panic(err.Error())
			}
				count =1
			}
			row.Close()
		if count <1{
					insert, err := db1.Query("insert into giohang values('"+string("")+"','"+strconv.Itoa(striduser)+"')")
						if err != nil {
						panic(err.Error())
					}
					insert.Close()

					row, err := db1.Query("select * from giohang where user='"+strconv.Itoa(striduser)+"'")
					if err != nil {
						panic(err.Error())
						}
					for row.Next() {
						var giohang giohang
					err = row.Scan(&giohang.Idgiohang,&giohang.Username)
					idgiohang =giohang.Idgiohang
					if err != nil {
						panic(err.Error())
					}

					}
			

					row.Close()
					demsoluong, err := db1.Query("select * from cartmini where idsanpham='"+r.FormValue("idsanpham")+"'")
					if err != nil {
						panic(err.Error())
						}
					var soluongcuasanpham int 
					soluongcuasanpham =0
					for demsoluong.Next() {
					
						soluongcuasanpham=1
					}
					demsoluong.Close()
            if soluongcuasanpham <1{
					insertcartmini, err := db1.Query("insert into cartmini values('"+string("")+"','"+r.FormValue("idsanpham")+"','"+r.FormValue("soluong")+"','"+strconv.Itoa(idgiohang)+"')")
					if err != nil {
						panic(err.Error())
					}
					insertcartmini.Close()
					}else{
						updatecartmini, err := db1.Query("update cartmini set soluong='"+r.FormValue("soluong")+"' where idgiohang='"+strconv.Itoa(idgiohang)+"' and idsanpham='"+r.FormValue("idsanpham")+"'")
					if err != nil {
						panic(err.Error())
					}
					updatecartmini.Close()
			}

		}else{
					demsoluong, err := db1.Query("select * from cartmini where idsanpham='"+r.FormValue("idsanpham")+"' and idgiohang='"+strconv.Itoa(idgiohang)+"'")
					if err != nil {
						panic(err.Error())
						}
					var soluongcuasanpham int 
					soluongcuasanpham =0
					for demsoluong.Next() {
					
						soluongcuasanpham=1
					}
					demsoluong.Close()
				
            if soluongcuasanpham <1{
					insertcartmini, err := db1.Query("insert into cartmini values('"+string("")+"','"+r.FormValue("idsanpham")+"','"+r.FormValue("soluong")+"','"+strconv.Itoa(idgiohang)+"')")
					if err != nil {
						panic(err.Error())
					}
				
					insertcartmini.Close()
		    }else{
						updatecartmini, err := db1.Query("update cartmini set soluong='"+r.FormValue("soluong")+"' where idgiohang='"+strconv.Itoa(idgiohang)+"' and idsanpham='"+r.FormValue("idsanpham")+"'")
					if err != nil {
						panic(err.Error())
					}
					updatecartmini.Close()
			}

		}
		// for row.Next() {
		// 	var sanpham1 nhacungcap
		// 	err = row.Scan(&sanpham1.Idnhacungcap, &sanpham1.TenNhaCungCap,&sanpham1.Diachi, &sanpham1.Email, &sanpham1.Sdt,&sanpham1.Fax)
		// 	if err != nil {
		// 		panic(err.Error())
		// 	}

		// 	Nhacungcap = append(Nhacungcap, sanpham1)
		// }
	}
	io.WriteString(w,"sản phẩm đã được thêm vào giỏ hàng")
}else{
	io.WriteString(w,"số Lượng Sản Phẩm Không Đủ")
}

}
func danhsachsanphamtronggiohang(w http.ResponseWriter, r *http.Request){
	session,_ := store.Get(r,"session")
	if session.Values["username"] !=nil{
		var iduser =session.Values["iduser"]
		var striduser = iduser.(int)
		var idgiohang int
		type cart struct{
			Idsanpham  int
			Soluong   string
		}
		var mang [] cart
		row, err := db1.Query("select * from giohang where user='"+strconv.Itoa(striduser)+"'")
		if err != nil {
			panic(err.Error())
			}

			for row.Next() {
				var giohang giohang
			err = row.Scan(&giohang.Idgiohang,&giohang.Username)
			idgiohang =giohang.Idgiohang
			if err != nil {
				panic(err.Error())
			}
			
			}
			row.Close()
			if   r.FormValue("idspcanxoa") != "" {
				row, err := db1.Query("delete from cartmini where idgiohang='"+strconv.Itoa(idgiohang)+"' and idsanpham='"+r.FormValue("idspcanxoa")+"'")
				if err != nil {
					panic(err.Error())
					}
		
				io.WriteString(w,"Xóa Thành Công!")
				row.Close()
			}
			cartmini, err := db1.Query("select * from cartmini where idgiohang='"+strconv.Itoa(idgiohang)+"'")
			if err != nil {
				panic(err.Error())
				}
			var soluongtrongio int
			soluongtrongio =0
			for cartmini.Next() {
			 var giohang itemgiohang
			 soluongtrongio+=1
			 err = cartmini.Scan(&giohang.Iditem,&giohang.Idsanpham,&giohang.Soluong,&giohang.Idgiohang)
			 var temp = struct{
				 Idsanpham int
				 Soluong   string
			 }{giohang.Idsanpham,strconv.Itoa(giohang.Soluong)}
			 mang = append(mang,temp)
			if err != nil {
				panic(err.Error())
			}
			}
			cartmini.Close()
			var arrSanpham []sanpham
			slsanpham,err := db1.Query("SELECT * FROM `sanpham` WHERE 1 ")
			if err != nil{
				panic(err.Error())
			}
			for slsanpham.Next(){
				var sanpham sanpham
				err = slsanpham.Scan(&sanpham.Idsp, &sanpham.Tensanpham, &sanpham.Loaisanpham, &sanpham.Gia, &sanpham.Mota, &sanpham.Hinhanh1, &sanpham.Hinhanh2, &sanpham.Hinhanh3)
				if err != nil {
					panic(err.Error())
				}
				arrSanpham = append(arrSanpham,sanpham)
			}
						slsanpham1,err := db1.Query("SELECT * FROM `sanpham1` WHERE 1  ")
						if err != nil{
							panic(err.Error())
						}

						var arrsanpham1 []sanpham1
						for slsanpham1.Next(){
							var sanpham1 sanpham1

							
							
							err = slsanpham1.Scan(&sanpham1.Giamgia, &sanpham1.Manhasanxuat, &sanpham1.Manhacungcap, &sanpham1.Ngaycapnhat, &sanpham1.Luotxem, &sanpham1.Luotbinhchon, &sanpham1.Soluongton, &sanpham1.Cauhinh,&sanpham1.Idsp1)
							if err != nil {
								panic(err.Error())
							}
						
							arrsanpham1 = append(arrsanpham1,sanpham1)
						}
						slsanpham.Close()
			
			tpl,err:= template.New("static/danhsachsanphamtronggiohang.html").Funcs(template.FuncMap{"tinhgiasale":tinhgiasale,"add": add,"floatchuoi":floatchuoi}).ParseFiles("static/danhsachsanphamtronggiohang.html")
			if err != nil {
				panic(err.Error())
			}
		 
				tpl.ExecuteTemplate(w, "danhsachsanphamtronggiohang.html",struct{
					Giohang  [] cart
					Sanpham   [] sanpham
					Soluong   int
					Sanpham1  []sanpham1
				}{mang,arrSanpham,soluongtrongio,arrsanpham1})


	}else{
		type giohang struct{
			Idsanpham  int
			Soluong   string
		}
		var soluongtrongio int
		var mang1 [] int
		var mang2 [] string
		var mang []giohang
		var idsanpham =session.Values["idsanpham"] 
		var stridsanpham string
		var strsoluong string
		strsoluong =""
		stridsanpham = ""
		if idsanpham != nil{
			stridsanpham = idsanpham.(string)
		
		}
		
		
		var soluong = session.Values["soluong"]
		if soluong != nil{
			strsoluong = soluong.(string)
		
		}
		if strsoluong != ""{
		f := func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsNumber(c)
		}
		var ints = strings.FieldsFunc(stridsanpham, f)

		var stridsp string

		var strsl string
		stridsp=""
		strsl =""
		if stridsanpham != ""{
		for i, num := range ints {
			if i != -1{

			}

			
			a,err := strconv.ParseInt(num,10,64)
			if err !=nil{

			
			panic(err.Error())
			}
			mang1 = append(mang1,int(a))
			

		}

		ints = strings.FieldsFunc(strsoluong, f)

		for i, num := range ints {

			if i != -1{

			}
			mang2 = append(mang2,num)

		}

		for i := 0; i < len(mang1); i++ {
			if   r.FormValue("idspcanxoa") != "" && r.FormValue("idspcanxoa") ==  strconv.Itoa(mang1[i] ){
			
				
			}else{
				
				stridsp += strconv.Itoa(mang1[i])+";"
				strsl +=mang2[i]+";"
			
			}
			
			var temp =struct {
			Idsanpham  int
			Soluong   string
		}{mang1[i],mang2[i]}
		mang = append(mang,temp)
		}

		if r.FormValue("idspcanxoa") != ""{
			session.Values["idsanpham"] = stridsp
			session.Values["soluong"]  = strsl
			session.Save(r,w) 
			
			
		}
		soluongtrongio=len(mang1)
var arrSanpham []sanpham
slsanpham,err := db1.Query("SELECT * FROM `sanpham` WHERE 1 ")
if err != nil{
	panic(err.Error())
}
for slsanpham.Next(){
	var sanpham sanpham
	err = slsanpham.Scan(&sanpham.Idsp, &sanpham.Tensanpham, &sanpham.Loaisanpham, &sanpham.Gia, &sanpham.Mota, &sanpham.Hinhanh1, &sanpham.Hinhanh2, &sanpham.Hinhanh3)
	if err != nil {
		panic(err.Error())
	}
	arrSanpham = append(arrSanpham,sanpham)
}
slsanpham.Close()

slsanpham1,err := db1.Query("SELECT * FROM `sanpham1` WHERE 1  ")
if err != nil{
	panic(err.Error())
}

var arrsanpham1 []sanpham1
for slsanpham1.Next(){
	var sanpham1 sanpham1

	
	
	err = slsanpham1.Scan(&sanpham1.Giamgia, &sanpham1.Manhasanxuat, &sanpham1.Manhacungcap, &sanpham1.Ngaycapnhat, &sanpham1.Luotxem, &sanpham1.Luotbinhchon, &sanpham1.Soluongton, &sanpham1.Cauhinh,&sanpham1.Idsp1)
	if err != nil {
		panic(err.Error())
	}

	arrsanpham1 = append(arrsanpham1,sanpham1)
}

slsanpham1.Close()

tpl,err:= template.New("static/danhsachsanphamtronggiohang.html").Funcs(template.FuncMap{"tinhgiasale":tinhgiasale,"add": add,"floatchuoi":floatchuoi}).ParseFiles("static/danhsachsanphamtronggiohang.html")
	if err != nil {
		panic(err.Error())
	}
 
		tpl.ExecuteTemplate(w, "danhsachsanphamtronggiohang.html",struct{
			Giohang  [] giohang
			Sanpham   [] sanpham
			Soluong   int
			Sanpham1  []sanpham1
		}{mang,arrSanpham,soluongtrongio,arrsanpham1})
	}
	}

	}

}
func dangnhapvadangkytaikhoan(w http.ResponseWriter, r *http.Request){
	session,_ := store.Get(r,"session")
    
	if r.FormValue("dangnhap")!=""{
		var a int
		a=0
		email := r.FormValue("email")
		password := r.FormValue("password")
		if strings.Count(password,"'") <=0{
		taikhoans,err := db1.Query("SELECT * FROM `taikhoan` WHERE email='"+email+"' and matkhau='"+password+"'")
	 if err != nil{
		 panic(err.Error())
	 }
	
	 for taikhoans.Next(){
		a+=1
		var tk taikhoan
	   err = taikhoans.Scan(&tk.Idtaikhoan,&tk.Tennguoidung,&tk.Ho,&tk.Matkhau,&tk.Email,&tk.Loaitaikhoan)
	   session.Values["iduser"]=tk.Idtaikhoan
	
		 session.Values["username"]=tk.Ho+" "+tk.Tennguoidung
		 session.Save(r,w) 
		if err != nil {
				panic(err.Error())
			}
	 }

	 taikhoans.Close()
	}
	
	 if a ==1{
		 io.WriteString(w,"Đăng Nhập Thành Công")

		 
		 
	 } else{
		io.WriteString(w,"Đăng Nhập Thất bại")
	 }
	 
	}
	if r.FormValue("dangxuat") != ""{

		session.Values["iduser"]=nil
		  session.Values["username"]=nil
		  session.Save(r,w) 
	}


	if r.FormValue("dangky")!=""{
          if r.FormValue("maxacnhan") == session.Values["maxacnhanlai"]{
			email := r.FormValue("email")
			password := r.FormValue("password")
			ho := r.FormValue("ho")
			ten := r.FormValue("ten")
			taikhoans,err := db1.Query("SELECT * FROM `taikhoan` WHERE email='"+email+"'")
		if err != nil{
			panic(err.Error())
			}
		var a int
		a=0
		for taikhoans.Next(){
			a+=1
			
		}
		taikhoans.Close()
		if a!=1 {
			dangkytaikhoan,err := db1.Query("insert into taikhoan values('','"+ten+"','"+ho+"','"+password+"','"+email+"','3"+"')")
			dangkytaikhoan.Close()
			taikhoans,err := db1.Query("SELECT * FROM `taikhoan` WHERE email='"+email+"' and matkhau='"+password+"'")
			if err != nil{
				panic(err.Error())
			}
			
			var a int
			a=0
						for taikhoans.Next(){
						a+=1
						var tk taikhoan
						err = taikhoans.Scan(&tk.Idtaikhoan,&tk.Tennguoidung,&tk.Ho,&tk.Matkhau,&tk.Email,&tk.Loaitaikhoan)
						session.Values["iduser"]=tk.Idtaikhoan
						
							session.Values["username"]=tk.Ho+" "+tk.Tennguoidung
							session.Save(r,w) 
								if err != nil {
									panic(err.Error())
								}
						}
						taikhoans.Close()
			io.WriteString(w,"Đăng Ký Thành Công ^^")
		 }else{
			io.WriteString(w,"Email Đã Tồn Tại ^^")
		 }
	 }else{
		io.WriteString(w,"Mã Xác Nhận Không Chính Xác ^^")
}
	}
				
}
func dangnhaptaikhoan(w http.ResponseWriter, r *http.Request){
	session,_ := store.Get(r,"session")
	var SS  string
	
	
	if session.Values["username"]!=nil{
		var S= session.Values["username"]
		 SS = S.(string)
		
	
		 
	}else{SS="0"

	}
	var lsp [] loaisanpham

	slloaisanpham,err := db1.Query("SELECT * FROM `loaisanpham` WHERE 1 ")
	if err != nil{
		panic(err.Error())
	}
	for slloaisanpham.Next(){
		var loaisanpham loaisanpham
		err = slloaisanpham.Scan(&loaisanpham.Idloaisp,&loaisanpham.Tenloaisanpham,&loaisanpham.Mota)
		if err != nil {
			panic(err.Error())
		}
		lsp = append(lsp,loaisanpham)
	}

	slloaisanpham.Close()
	if err != nil {
		panic(err.Error())
	}
	

	var SS1 string
	if session.Values["username"]!=nil{

	var S1= session.Values["username"]
	SS1 = S1.(string)
	}else{
		SS1=""
	}
		
	    tpl, err := template.ParseFiles("static/login-register.html")
		tpl.ExecuteTemplate(w, "login-register.html",struct{
		S string
		Username string
		Loaisanpham  []loaisanpham
		}{SS,SS1,lsp})
		if err != nil {
			panic(err.Error())
		}
}
func viewthongtinsanpham(w http.ResponseWriter, r *http.Request){
	
	var lsp [] loaisanpham
	var  arrSanpham [] sanpham

	slloaisanpham,err := db1.Query("SELECT * FROM `loaisanpham` WHERE 1 ")
	if err != nil{
		panic(err.Error())
	}
	for slloaisanpham.Next(){
		var loaisanpham loaisanpham
		err = slloaisanpham.Scan(&loaisanpham.Idloaisp,&loaisanpham.Tenloaisanpham,&loaisanpham.Mota)
		if err != nil {
			panic(err.Error())
		}
		lsp = append(lsp,loaisanpham)
	}
	if err != nil {
		panic(err.Error())
	}
	slloaisanpham.Close()
	
	slsanpham,err := db1.Query("SELECT * FROM `sanpham` WHERE idsp='"+r.FormValue("idsanpham")+"'")
	if err != nil{
		panic(err.Error())
	}
	for slsanpham.Next(){
		var sanpham sanpham
		err = slsanpham.Scan(&sanpham.Idsp, &sanpham.Tensanpham, &sanpham.Loaisanpham, &sanpham.Gia, &sanpham.Mota, &sanpham.Hinhanh1, &sanpham.Hinhanh2, &sanpham.Hinhanh3)
		if err != nil {
			panic(err.Error())
		}
		arrSanpham = append(arrSanpham,sanpham)
	}

	slsanpham.Close()
var arrsanpham1 []sanpham1
	slsanpham1,err := db1.Query("SELECT * FROM `sanpham1` WHERE 1")
	if err != nil{
		panic(err.Error())
	}
	for slsanpham1.Next(){
		var sanpham1 sanpham1
		err = slsanpham1.Scan(&sanpham1.Giamgia, &sanpham1.Manhasanxuat, &sanpham1.Manhacungcap, &sanpham1.Ngaycapnhat, &sanpham1.Luotxem, &sanpham1.Luotbinhchon, &sanpham1.Soluongton, &sanpham1.Cauhinh,&sanpham1.Idsp1)
		if err != nil {
			panic(err.Error())
		}
		arrsanpham1 = append(arrsanpham1,sanpham1)
	}	
	slsanpham1.Close()
	tpl,err:= template.New("static/viewchititetsanpham.html").Funcs(template.FuncMap{"tinhgiasale":tinhgiasale,"add": add,"floatchuoi":floatchuoi}).ParseFiles("static/viewchititetsanpham.html")
	
	tpl.ExecuteTemplate(w, "viewchititetsanpham.html", struct {
		Loaisanpham []loaisanpham
		Sanpham   []sanpham
		Sanpham1   []sanpham1
	}{lsp,arrSanpham,arrsanpham1})
	if err != nil {
		panic(err.Error())
	}
}

func viewcart(w http.ResponseWriter, r *http.Request){

	session,_ := store.Get(r,"session")
	type tempsanpham struct{
		Idsanpham int
		Soluong int
	}
	var mnagtempsanpham [] tempsanpham
	if session.Values["username"] !=nil{
		var iduser =session.Values["iduser"]
		var striduser = iduser.(int)
		var idgiohang int
		type cart struct{
			Idsanpham  int
			Soluong   string
		}
		var mang [] cart
		row, err := db1.Query("select * from giohang where user='"+strconv.Itoa(striduser)+"'")
		if err != nil {
			panic(err.Error())
			}

			for row.Next() {
				var giohang giohang
			err = row.Scan(&giohang.Idgiohang,&giohang.Username)
			idgiohang =giohang.Idgiohang
			if err != nil {
				panic(err.Error())
			}
			
			}
			row.Close()
			if   r.FormValue("idspcanxoa") != "" {
				row, err := db1.Query("delete from cartmini where idgiohang='"+strconv.Itoa(idgiohang)+"' and idsanpham='"+r.FormValue("idspcanxoa")+"'")
				if err != nil {
					panic(err.Error())
					}
		
				io.WriteString(w,"Xóa Thành Công!")
				row.Close()
			}
			cartmini, err := db1.Query("select * from cartmini where idgiohang='"+strconv.Itoa(idgiohang)+"'")
			if err != nil {
				panic(err.Error())
				}
			var soluongtrongio int
			soluongtrongio =0
			for cartmini.Next() {
			 var giohang itemgiohang
			 soluongtrongio+=1
			 err = cartmini.Scan(&giohang.Iditem,&giohang.Idsanpham,&giohang.Soluong,&giohang.Idgiohang)

			 ktsoluongton, err:=db1.Query("select  soluongton from sanpham1 where idsp='"+strconv.Itoa(giohang.Idsanpham)+"'")
			 var soluongtontemp int
			 for ktsoluongton.Next(){
				err = ktsoluongton.Scan(&soluongtontemp)
			 }
			 if soluongtontemp < giohang.Soluong{
               
               if soluongtontemp>0 {

				  updateslton, err := db1.Query("update cartmini set soluong="+strconv.Itoa(soluongtontemp)+" where idsanpham='"+strconv.Itoa(giohang.Idsanpham)+"'")
				  if err != nil {
					panic(err.Error())
				}


				  updateslton.Close()
				  var temmm tempsanpham
				  temmm.Idsanpham =giohang.Idsanpham
				  temmm.Soluong = giohang.Soluong
				  mnagtempsanpham =append(mnagtempsanpham,temmm)
			}else{
				updateslton, err := db1.Query("delete from cartmini where idsanpham='"+strconv.Itoa(giohang.Idsanpham)+"'")
				if err != nil {
				  panic(err.Error())
			  }


				updateslton.Close()
				var temmm tempsanpham
				temmm.Idsanpham =giohang.Idsanpham
				temmm.Soluong = giohang.Soluong
				mnagtempsanpham =append(mnagtempsanpham,temmm)
			}
			 }else{
			 var temp = struct{
				 Idsanpham int
				 Soluong   string
			 }{giohang.Idsanpham,strconv.Itoa(giohang.Soluong)}
			 mang = append(mang,temp)
			if err != nil {
				panic(err.Error())
			}
			 }
			 ktsoluongton.Close()
			}
			cartmini.Close()
			var arrSanpham []sanpham
			slsanpham,err := db1.Query("SELECT * FROM `sanpham` WHERE 1 ")
			if err != nil{
				panic(err.Error())
			}
			for slsanpham.Next(){
				var sanpham sanpham
				err = slsanpham.Scan(&sanpham.Idsp, &sanpham.Tensanpham, &sanpham.Loaisanpham, &sanpham.Gia, &sanpham.Mota, &sanpham.Hinhanh1, &sanpham.Hinhanh2, &sanpham.Hinhanh3)
				if err != nil {
					panic(err.Error())
				}
				arrSanpham = append(arrSanpham,sanpham)
			}
			slsanpham.Close()
	slsanpham1,err := db1.Query("SELECT * FROM `sanpham1` WHERE 1  ")
	if err != nil{
		panic(err.Error())
	}

	var arrsanpham1 []sanpham1
	for slsanpham1.Next(){
		var sanpham1 sanpham1

		
		
		err = slsanpham1.Scan(&sanpham1.Giamgia, &sanpham1.Manhasanxuat, &sanpham1.Manhacungcap, &sanpham1.Ngaycapnhat, &sanpham1.Luotxem, &sanpham1.Luotbinhchon, &sanpham1.Soluongton, &sanpham1.Cauhinh,&sanpham1.Idsp1)
		if err != nil {
			panic(err.Error())
		}
	
		arrsanpham1 = append(arrsanpham1,sanpham1)
	}
	slsanpham1.Close()

			
			tpl,err:= template.New("static/shopping-cartmini.html").Funcs(template.FuncMap{"add": add,"floatchuoi":floatchuoi}).ParseFiles("static/shopping-cartmini.html")
			if err != nil {
				panic(err.Error())
			}
		 
				tpl.ExecuteTemplate(w, "shopping-cartmini.html",struct{
					Giohang  [] cart
					Sanpham   [] sanpham
					Soluong   int
					Sanpham1   []sanpham1
					Thongbao   []tempsanpham
				}{mang,arrSanpham,soluongtrongio,arrsanpham1,mnagtempsanpham})


	}else{
		type giohang struct{
			Idsanpham  int
			Soluong   string
		}
		var soluongtrongio int
		var mang1 [] int
		var mang2 [] string
		var mang []giohang
		var idsanpham =session.Values["idsanpham"] 
		var stridsanpham string
		var strsoluong string
		strsoluong =""
		stridsanpham = ""
		if idsanpham != nil{
			stridsanpham = idsanpham.(string)
		
		}
		
		
		var soluong = session.Values["soluong"]
		if soluong != nil{
			strsoluong = soluong.(string)
		
		}
		if strsoluong != ""{
		f := func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsNumber(c)
		}
		var ints = strings.FieldsFunc(stridsanpham, f)

		var stridsp string

		var strsl string
		stridsp=""
		strsl =""
		if stridsanpham != ""{
		for i, num := range ints {
			if i != -1{

			}

			
			a,err := strconv.ParseInt(num,10,64)
			if err !=nil{

			
		
			}
			mang1 = append(mang1,int(a))
			

		}

		ints = strings.FieldsFunc(strsoluong, f)

		for i, num := range ints {

			if i != -1{

			}
			mang2 = append(mang2,num)

		}

		for i := 0; i < len(mang1); i++ {
			if   r.FormValue("idspcanxoa") != "" && r.FormValue("idspcanxoa") ==  strconv.Itoa(mang1[i] ){
			
				
			}else{
				
				stridsp += strconv.Itoa(mang1[i])+";"
				strsl +=mang2[i]+";"
			
			}
			
			var temp =struct {
			Idsanpham  int
			Soluong   string
		}{mang1[i],mang2[i]}
		mang = append(mang,temp)
		}

		var capnhatlaisoluong int
		capnhatlaisoluong=0

		var stridsp11 string
		stridsp11=""
		var strsl11 string
		strsl11=""
		for i := 0; i < len(mang1); i++ {
		ktsoluongton, err:=db1.Query("select  soluongton from sanpham1 where idsp='"+strconv.Itoa(mang1[i])+"'")
		var soluongtontemp int
		for ktsoluongton.Next(){
		   err = ktsoluongton.Scan(&soluongtontemp)
		}
		itemp,err :=strconv.ParseInt(mang2[i],10,64)
		if err!=nil{
			panic(err.Error())
		}
	
		
		if soluongtontemp <int(itemp){
			stridsp11 += strconv.Itoa(mang1[i])+";"
			strsl11 +=strconv.Itoa(soluongtontemp)+";"
		
				capnhatlaisoluong=1
		var temmm tempsanpham
		temmm.Idsanpham =mang1[i]
		temmm.Soluong = int(itemp)
		mnagtempsanpham =append(mnagtempsanpham,temmm)
	    }else{
			if soluongtontemp!=0{
			stridsp11 += strconv.Itoa(mang1[i])+";"
			strsl11 +=mang2[i]+";"
			}
		}
          ktsoluongton.Close()

			
		}
if capnhatlaisoluong==1{
	session.Values["idsanpham"] = stridsp11
	session.Values["soluong"]  = strsl11
	session.Save(r,w) 
}



		if r.FormValue("idspcanxoa") != ""{
			session.Values["idsanpham"] = stridsp
			session.Values["soluong"]  = strsl
			session.Save(r,w) 
			
			
		}
		soluongtrongio=len(mang1)
var arrSanpham []sanpham
slsanpham,err := db1.Query("SELECT * FROM `sanpham` WHERE 1 ")
if err != nil{
	panic(err.Error())
}
for slsanpham.Next(){
	var sanpham sanpham
	err = slsanpham.Scan(&sanpham.Idsp, &sanpham.Tensanpham, &sanpham.Loaisanpham, &sanpham.Gia, &sanpham.Mota, &sanpham.Hinhanh1, &sanpham.Hinhanh2, &sanpham.Hinhanh3)
	if err != nil {
		panic(err.Error())
	}
	arrSanpham = append(arrSanpham,sanpham)
}
slsanpham.Close()
slsanpham1,err := db1.Query("SELECT * FROM `sanpham1` WHERE 1  ")
if err != nil{
	panic(err.Error())
}

var arrsanpham1 []sanpham1
for slsanpham1.Next(){
	var sanpham1 sanpham1

	
	
	err = slsanpham1.Scan(&sanpham1.Giamgia, &sanpham1.Manhasanxuat, &sanpham1.Manhacungcap, &sanpham1.Ngaycapnhat, &sanpham1.Luotxem, &sanpham1.Luotbinhchon, &sanpham1.Soluongton, &sanpham1.Cauhinh,&sanpham1.Idsp1)
	if err != nil {
		panic(err.Error())
	}

	arrsanpham1 = append(arrsanpham1,sanpham1)
}
slsanpham1.Close()
tpl,err:= template.New("static/shopping-cartmini.html").Funcs(template.FuncMap{"add": add,"floatchuoi":floatchuoi}).ParseFiles("static/shopping-cartmini.html")
	if err != nil {
		panic(err.Error())
	}
 
		tpl.ExecuteTemplate(w, "shopping-cartmini.html",struct{
			Giohang  [] giohang
			Sanpham   [] sanpham
			Soluong   int
			Sanpham1   []sanpham1
			Thongbao   []tempsanpham
		}{mang,arrSanpham,soluongtrongio,arrsanpham1,mnagtempsanpham})
	}
	}

	}

}

func cart(w http.ResponseWriter, r *http.Request){
	session,_ := store.Get(r,"session")
	
    var lsp [] loaisanpham

	slloaisanpham,err := db1.Query("SELECT * FROM `loaisanpham` WHERE 1 ")
	if err != nil{
		panic(err.Error())
	}
	for slloaisanpham.Next(){
		var loaisanpham loaisanpham
		err = slloaisanpham.Scan(&loaisanpham.Idloaisp,&loaisanpham.Tenloaisanpham,&loaisanpham.Mota)
		if err != nil {
			panic(err.Error())
		}
		lsp = append(lsp,loaisanpham)
	}
	if err != nil {
		panic(err.Error())
	}
	
	slloaisanpham.Close()
	var SS string
	if session.Values["username"]!=nil{

	var S= session.Values["username"]
	SS = S.(string)
	}else{
		SS=""
	}
	tpl, err := template.ParseFiles("static/shopping-cart.html")
	if err != nil {
		panic(err.Error())
	}
	tpl.ExecuteTemplate(w, "shopping-cart.html", struct{
		Loaisanpham []loaisanpham
		Username  string
	}{lsp,SS})
}
type Person struct {
	Frist string
	Last  string
	Age   int
}
func add(x string, y string,z string) string {
	if z == "*"{
	var a1,b1 int64
	
	if a, err := strconv.ParseInt(x, 10, 64); err == nil {
		a1=a
	
	}
	if b, err := strconv.ParseInt(y, 10, 64); err == nil {
		b1=b
	}
	 var temp string
	 var res string
	 res = ""
	 temp =Reverse(strconv.Itoa(int(b1) * int(a1) ))
	 for i := 0; i < len(temp); i++ {
        if i %3==0 && i!=0{
			res+=","
		}
		res+= temp[i:i+1]
      
    }
	 temp = Reverse(res)

	return temp+"VNĐ"
  }else{

  var a1,b1 int64
	
	if a, err := strconv.ParseInt(x, 10, 64); err == nil {
		a1=a
	
	}
	if b, err := strconv.ParseInt(y, 10, 64); err == nil {
		b1=b
	}
	return strconv.Itoa( int(b1) + int(a1)) +"VNĐ"
	
}
}
func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
func cong(x, y int) int {
    return x + y
}
func tinhgiasale(x string, y int) string {
	var a int
	if a, err := strconv.ParseInt(x, 10, 64); err == nil {
		fmt.Printf("i=%d, type: %T\n", a, a)
	}
	var b float64
	b= math.Log((float64(y)/float64(a))*100)
    fmt.Print(b)
   var c string
   c= ""
    return c
}
func floatchuoi(x int) string{
	return   strconv.Itoa(x)
}
var db1 *sql.DB
func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
			log.Panicf("%s environment variable not set.", k)
	}
	return v
}
func moketnoi() *sql.DB {

	db, err := sql.Open("mysql", "g6cuvyj767pq9rrs:xbss6qkudakpmfux@tcp(wvulqmhjj9tbtc1w.cbetxkdyhwsb.us-east-1.rds.amazonaws.com:3306)/kd48zxahcswhpezv")
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	// db, err := sql.Open("mysql", "luandeptrai98:root123@cloudsql(34.68.0.92)/qlbh")
	// 	if err != nil {
	// 	panic(err.Error())
	// }
	return db

}

type temp1 struct {
	bien string
}

func hoadon(w http.ResponseWriter, r *http.Request){

	session,_ := store.Get(r,"session")
	var tongtien int
	if session.Values["username"] !=nil{
		var iduser =session.Values["iduser"]
		var striduser = iduser.(int)
		var idgiohang int
		type cart struct{
			Idsanpham  int
			Soluong   string
		}
		var mang [] cart
		row, err := db1.Query("select * from giohang where user='"+strconv.Itoa(striduser)+"'")
		if err != nil {
			panic(err.Error())
			}

			for row.Next() {
				var giohang giohang
			err = row.Scan(&giohang.Idgiohang,&giohang.Username)
			idgiohang =giohang.Idgiohang
			if err != nil {
				panic(err.Error())
			}
			
			}
			row.Close()
			if   r.FormValue("idspcanxoa") != "" {
				row, err := db1.Query("delete from cartmini where idgiohang='"+strconv.Itoa(idgiohang)+"' and idsanpham='"+r.FormValue("idspcanxoa")+"'")
				if err != nil {
					panic(err.Error())
					}
		
				io.WriteString(w,"Xóa Thành Công!")
				row.Close()
			}
			cartmini, err := db1.Query("select * from cartmini where idgiohang='"+strconv.Itoa(idgiohang)+"'")
			if err != nil {
				panic(err.Error())
				}
			var soluongtrongio int
			soluongtrongio =0
			for cartmini.Next() {
			 var giohang itemgiohang
			 soluongtrongio+=1
			 err = cartmini.Scan(&giohang.Iditem,&giohang.Idsanpham,&giohang.Soluong,&giohang.Idgiohang)
			 var temp = struct{
				 Idsanpham int
				 Soluong   string
			 }{giohang.Idsanpham,strconv.Itoa(giohang.Soluong)}
		
			 mang = append(mang,temp)
			if err != nil {
				panic(err.Error())
			}
			}
			cartmini.Close()
			var arrSanpham []sanpham
			slsanpham,err := db1.Query("SELECT * FROM `sanpham` WHERE 1 ")
			if err != nil{
				panic(err.Error())
			}
			
			for slsanpham.Next(){
				var sanpham sanpham
				err = slsanpham.Scan(&sanpham.Idsp, &sanpham.Tensanpham, &sanpham.Loaisanpham, &sanpham.Gia, &sanpham.Mota, &sanpham.Hinhanh1, &sanpham.Hinhanh2, &sanpham.Hinhanh3)
				if err != nil {
					panic(err.Error())
				}
				cartmini, err := db1.Query("select * from cartmini where idgiohang='"+strconv.Itoa(idgiohang)+"'")
				for cartmini.Next(){
					var giohang itemgiohang
					err = cartmini.Scan(&giohang.Iditem,&giohang.Idsanpham,&giohang.Soluong,&giohang.Idgiohang)
					if giohang.Idsanpham == sanpham.Idsp{
						var a int64
						a,err =strconv.ParseInt(sanpham.Gia,10,64 )
						if err != nil {
							panic(err.Error())
						}

													
								slsanpham1,err := db1.Query("SELECT * FROM `sanpham1` WHERE 1  ")
								if err != nil{
									panic(err.Error())
								}
								for slsanpham1.Next(){
									var sanpham1 sanpham1

									
									
									err = slsanpham1.Scan(&sanpham1.Giamgia, &sanpham1.Manhasanxuat, &sanpham1.Manhacungcap, &sanpham1.Ngaycapnhat, &sanpham1.Luotxem, &sanpham1.Luotbinhchon, &sanpham1.Soluongton, &sanpham1.Cauhinh,&sanpham1.Idsp1)
									if err != nil {
										panic(err.Error())
									}
								 if	sanpham1.Idsp1 == sanpham.Idsp{
                                     if sanpham1.Giamgia ==0{
										tongtien += int(a) * int(giohang.Soluong)
									 }else
									 {	tongtien += sanpham1.Giamgia * int(giohang.Soluong)

									 }
								 }
									
								}


						
					
					}
				
				}
				cartmini.Close()
			
				arrSanpham = append(arrSanpham,sanpham)
			}
			slsanpham.Close()
			if r.FormValue("thanhpho") != "1"{
				if tongtien <=2000000{
					tongtien += 25000
					}
			}
			slsanpham1,err := db1.Query("SELECT * FROM `sanpham1` WHERE 1 ORDER BY ngaycapnhat LIMIT 15 ")
			if err != nil{
				panic(err.Error())
			}
			var arrsanpham1  [] sanpham1 
			for slsanpham1.Next(){
				var sanpham1 sanpham1
		
				
				
				err = slsanpham1.Scan(&sanpham1.Giamgia, &sanpham1.Manhasanxuat, &sanpham1.Manhacungcap, &sanpham1.Ngaycapnhat, &sanpham1.Luotxem, &sanpham1.Luotbinhchon, &sanpham1.Soluongton, &sanpham1.Cauhinh,&sanpham1.Idsp1)
				if err != nil {
					panic(err.Error())
				}
			
				arrsanpham1 = append(arrsanpham1,sanpham1)
			}
			slsanpham1.Close()
			tpl,err:= template.New("static/donhang.html").Funcs(template.FuncMap{"tinhgiasale":tinhgiasale,"add": add,"floatchuoi":floatchuoi}).ParseFiles("static/donhang.html")

			if err != nil {
				panic(err.Error())
			}
		 
				tpl.ExecuteTemplate(w, "donhang.html",struct{
					Giohang  [] cart
					Sanpham   [] sanpham
					Soluong   int
					Ship      string
					Tongtien  string
					Sanpham1   []sanpham1
				}{mang,arrSanpham,soluongtrongio,r.FormValue("thanhpho"),strconv.Itoa(tongtien),arrsanpham1})


	}else{
		type giohang struct{
			Idsanpham  int
			Soluong   string
		}
		var soluongtrongio int
		var mang1 [] int
		var mang2 [] string
		var mang []giohang
		var idsanpham =session.Values["idsanpham"] 
		var stridsanpham string
		var strsoluong string
		strsoluong =""
		stridsanpham = ""
		if idsanpham != nil{
			stridsanpham = idsanpham.(string)
		
		}
		
		
		var soluong = session.Values["soluong"]
		if soluong != nil{
			strsoluong = soluong.(string)
		
		}
		if strsoluong != ""{
		f := func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsNumber(c)
		}
		var ints = strings.FieldsFunc(stridsanpham, f)

		var stridsp string

		var strsl string
		stridsp=""
		strsl =""
		if stridsanpham != ""{
		for i, num := range ints {
			if i != -1{

			}

			
			a,err := strconv.ParseInt(num,10,64)
			if err !=nil{

			
			panic(err.Error())
			}
			mang1 = append(mang1,int(a))
			

		}

		ints = strings.FieldsFunc(strsoluong, f)

		for i, num := range ints {

			if i != -1{

			}
			mang2 = append(mang2,num)

		}

		for i := 0; i < len(mang1); i++ {
			if   r.FormValue("idspcanxoa") != "" && r.FormValue("idspcanxoa") ==  strconv.Itoa(mang1[i] ){
			
				
			}else{
				
				stridsp += strconv.Itoa(mang1[i])+";"
				strsl +=mang2[i]+";"
			
			}
			
			var temp =struct {
			Idsanpham  int
			Soluong   string
		}{mang1[i],mang2[i]}
		mang = append(mang,temp)
		}

		if r.FormValue("idspcanxoa") != ""{
			session.Values["idsanpham"] = stridsp
			session.Values["soluong"]  = strsl
			session.Save(r,w) 
			
			
		}
		soluongtrongio=len(mang1)
var arrSanpham []sanpham
slsanpham,err := db1.Query("SELECT * FROM `sanpham` WHERE 1 ")
if err != nil{
	panic(err.Error())
}
for slsanpham.Next(){
	var sanpham sanpham
	err = slsanpham.Scan(&sanpham.Idsp, &sanpham.Tensanpham, &sanpham.Loaisanpham, &sanpham.Gia, &sanpham.Mota, &sanpham.Hinhanh1, &sanpham.Hinhanh2, &sanpham.Hinhanh3)
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < len(mang1); i++ {
		if mang1[i] == sanpham.Idsp{
			var a,giasp int64
			a,err = strconv.ParseInt(mang2[i],10,64)
			giasp,err = strconv.ParseInt(sanpham.Gia,10,64)
			slsanpham1,err := db1.Query("SELECT * FROM `sanpham1`")
			if err != nil{
				panic(err.Error())
			}
			for slsanpham1.Next(){
				var sanpham1 sanpham1

				
				
				err = slsanpham1.Scan(&sanpham1.Giamgia, &sanpham1.Manhasanxuat, &sanpham1.Manhacungcap, &sanpham1.Ngaycapnhat, &sanpham1.Luotxem, &sanpham1.Luotbinhchon, &sanpham1.Soluongton, &sanpham1.Cauhinh,&sanpham1.Idsp1)
				if err != nil {
					panic(err.Error())
				}
			 if	sanpham1.Idsp1 == sanpham.Idsp{
				 if sanpham1.Giamgia ==0{
					tongtien += int(a) * int(giasp)
				 }else
				 {	tongtien += sanpham1.Giamgia * int(a)

				 }
			 }
				
			}
			slsanpham1.Close()

		}
	}
	arrSanpham = append(arrSanpham,sanpham)
}
if r.FormValue("thanhpho") != "1"{
	if tongtien <=2000000{
	tongtien += 25000
	}
}
slsanpham1,err := db1.Query("SELECT * FROM `sanpham1` WHERE 1 ORDER BY ngaycapnhat LIMIT 15 ")
if err != nil{
	panic(err.Error())
}
var arrsanpham1  [] sanpham1 
for slsanpham1.Next(){
	var sanpham1 sanpham1

	
	
	err = slsanpham1.Scan(&sanpham1.Giamgia, &sanpham1.Manhasanxuat, &sanpham1.Manhacungcap, &sanpham1.Ngaycapnhat, &sanpham1.Luotxem, &sanpham1.Luotbinhchon, &sanpham1.Soluongton, &sanpham1.Cauhinh,&sanpham1.Idsp1)
	if err != nil {
		panic(err.Error())
	}

	arrsanpham1 = append(arrsanpham1,sanpham1)
}
slsanpham1.Close()

tpl,err:= template.New("static/donhang.html").Funcs(template.FuncMap{"tinhgiasale":tinhgiasale,"add": add,"floatchuoi":floatchuoi}).ParseFiles("static/donhang.html")
	if err != nil {
		panic(err.Error())
	}
 
		tpl.ExecuteTemplate(w, "donhang.html",struct{
			Giohang  [] giohang
			Sanpham   [] sanpham
			Soluong   int
			Ship      string
			Tongtien  string
			Sanpham1 []sanpham1
		}{mang,arrSanpham,soluongtrongio,r.FormValue("thanhpho"),strconv.Itoa(tongtien),arrsanpham1})
	}
	}

	}


}
func laphoadon(w http.ResponseWriter, r *http.Request,idhoadon string) string{
	
	session,_ := store.Get(r,"session")
	var tongtien int
	var mang11 []string
    var mang22 []string
   var mang33 []string
	var mangsanpham []string
	var mangsoluong []string
	var thanhtien  []string
	
	if session.Values["username"] !=nil{
		var iduser =session.Values["iduser"]
		var striduser = iduser.(int)
		var idgiohang int
		type cart struct{
			Idsanpham  int
			Soluong   string
		}
		
		row, err := db1.Query("select * from giohang where user='"+strconv.Itoa(striduser)+"'")
		if err != nil {
			panic(err.Error())
			}

			for row.Next() {
				var giohang giohang
			err = row.Scan(&giohang.Idgiohang,&giohang.Username)
			idgiohang =giohang.Idgiohang
			if err != nil {
				panic(err.Error())
			}
			
			}
		
			row.Close()

		
			slsanpham,err := db1.Query("SELECT * FROM `sanpham` WHERE 1 ")
			if err != nil{
				panic(err.Error())
			}
			for slsanpham.Next(){
				var sanpham sanpham
				err = slsanpham.Scan(&sanpham.Idsp, &sanpham.Tensanpham, &sanpham.Loaisanpham, &sanpham.Gia, &sanpham.Mota, &sanpham.Hinhanh1, &sanpham.Hinhanh2, &sanpham.Hinhanh3)
				if err != nil {
					panic(err.Error())
				}
				cartmini, err := db1.Query("select * from cartmini where idgiohang='"+strconv.Itoa(idgiohang)+"'")
				for cartmini.Next(){
					var giohang itemgiohang
					err = cartmini.Scan(&giohang.Iditem,&giohang.Idsanpham,&giohang.Soluong,&giohang.Idgiohang)
					if giohang.Idsanpham == sanpham.Idsp{
						var a int64
						a,err =strconv.ParseInt(sanpham.Gia,10,64 )
						if err != nil {
							panic(err.Error())
						}
						
						
						slsanpham1,err := db1.Query("SELECT * FROM `sanpham1` WHERE 1")
						if err != nil{
							panic(err.Error())
						}
						
						for slsanpham1.Next(){
							var sanpham1 sanpham1
						
							
							
							err = slsanpham1.Scan(&sanpham1.Giamgia, &sanpham1.Manhasanxuat, &sanpham1.Manhacungcap, &sanpham1.Ngaycapnhat, &sanpham1.Luotxem, &sanpham1.Luotbinhchon, &sanpham1.Soluongton, &sanpham1.Cauhinh,&sanpham1.Idsp1)
							if err != nil {
								panic(err.Error())
							}
						
							if sanpham1.Idsp1 == sanpham.Idsp{
								if sanpham1.Giamgia ==0{
									tongtien += int(a) * int(giohang.Soluong)
									thanhtien = append(thanhtien,strconv.Itoa(int(a) * int(giohang.Soluong)))
								}else{
									tongtien += sanpham1.Giamgia * int(giohang.Soluong)
									thanhtien = append(thanhtien,strconv.Itoa(sanpham1.Giamgia * int(giohang.Soluong)))
								}
							}
						}
						slsanpham1.Close()
						mangsanpham  = append(mangsanpham,strconv.Itoa(giohang.Idsanpham))
						mangsoluong = append(mangsoluong,strconv.Itoa(giohang.Soluong))
						
						row, err := db1.Query("delete from cartmini where idgiohang='"+strconv.Itoa(idgiohang)+"' and idsanpham='"+strconv.Itoa(giohang.Idsanpham)+"'")
						
						if err != nil {
						panic(err.Error())

							}
							row.Close()
					}
				
				}
				cartmini.Close()
				
			}
			slsanpham.Close()
			for i := 0; i < len(mangsanpham); i++ {
				selectsp,err :=db1.Query("select soluongton from sanpham1 where idsp='"+mangsanpham[i]+"'")
				if err != nil {
					panic(err.Error())
						}
				var soluongtonsp int
				
				for selectsp.Next(){
					err = selectsp.Scan(&soluongtonsp)
				}
				selectsp.Close()
				isoluong,err :=strconv.ParseInt(mangsoluong[i],10,64)
				if err != nil {
					panic(err.Error())
						}
				if soluongtonsp <int(isoluong){
			
		
				}else{
					hoadon ,err := db1.Query("insert into itemhoadon values('','"+mangsanpham[i]+"','"+mangsoluong[i]+"','"+thanhtien[i]+"','"+idhoadon+"')")
					if err != nil {
						panic(err.Error())
							}
					hoadon.Close()
					i2,err :=strconv.ParseInt(mangsoluong[i],10,64)
					if err != nil {
						panic(err.Error())
							}
				updatesp,err:=db1.Query("update sanpham1 set soluongton='"+strconv.Itoa(soluongtonsp-int(i2))+"' where idsp='"+mangsanpham[i]+"'")
				if err != nil {
					panic(err.Error())
						}
				updatesp.Close()
				}


				if err != nil {
					panic(err.Error())
						}
						
			}
	
		if tongtien >=2000000{
			update,err :=db1.Query("update hoadon set phiship=0 where mahd='"+idhoadon+"'")
			if err != nil {
				panic(err.Error())
					}
					update.Close()
		}
				
		
			
			


	}else{
		type giohang struct{
			Idsanpham  int
			Soluong   string
		}
		var soluongtrongio int
		fmt.Print(soluongtrongio)
		var mang1 [] int
		var mang2 [] string
		var mang []giohang
		var idsanpham =session.Values["idsanpham"] 
		var stridsanpham string
		var strsoluong string
		strsoluong =""
		stridsanpham = ""
		if idsanpham != nil{
			stridsanpham = idsanpham.(string)
		
		}
		
		
		var soluong = session.Values["soluong"]
		if soluong != nil{
			strsoluong = soluong.(string)
		
		}
		if strsoluong != ""{
		f := func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsNumber(c)
		}
		var ints = strings.FieldsFunc(stridsanpham, f)

		var stridsp string

		var strsl string
		stridsp=""
		strsl =""
		if stridsanpham != ""{
		for i, num := range ints {
			if i != -1{

			}

			
			a,err := strconv.ParseInt(num,10,64)
			if err !=nil{

			
		panic(err.Error())
			}
		
			mang1 = append(mang1,int(a))
			

		}

		ints = strings.FieldsFunc(strsoluong, f)

		for i, num := range ints {

			if i != -1{

			}
			mang2 = append(mang2,num)

		}

		for i := 0; i < len(mang1); i++ {
			if   r.FormValue("idspcanxoa") != "" && r.FormValue("idspcanxoa") ==  strconv.Itoa(mang1[i] ){
			
				
			}else{
				
				stridsp += strconv.Itoa(mang1[i])+";"
				strsl +=mang2[i]+";"
			
			}
			
			var temp =struct {
			Idsanpham  int
			Soluong   string
		}{mang1[i],mang2[i]}
		mang = append(mang,temp)
		}

		if r.FormValue("idspcanxoa") != ""{
			session.Values["idsanpham"] = stridsp
			session.Values["soluong"]  = strsl
			session.Save(r,w) 
			
			
		}
		soluongtrongio=len(mang1)

slsanpham,err := db1.Query("SELECT * FROM `sanpham` WHERE 1 ")
if err != nil{
	panic(err.Error())
}
var tien string
tien="0"


for slsanpham.Next(){
	var sanpham sanpham
	err = slsanpham.Scan(&sanpham.Idsp, &sanpham.Tensanpham, &sanpham.Loaisanpham, &sanpham.Gia, &sanpham.Mota, &sanpham.Hinhanh1, &sanpham.Hinhanh2, &sanpham.Hinhanh3)
	if err != nil {
		panic(err.Error())
	}
	for i := 0; i < len(mang1); i++ {

		
		if mang1[i] == sanpham.Idsp{
			var a int64
			
			
			a,err = strconv.ParseInt(mang2[i],10,64)
			var b int64
			b,err = strconv.ParseInt(sanpham.Gia,10,64)
		    tongtien+=	int(a)* int(b)
		


		slsanpham1,err := db1.Query("SELECT * FROM `sanpham1` WHERE 1")
		if err != nil{
			panic(err.Error())
		}
		
		for slsanpham1.Next(){
			var sanpham1 sanpham1
		
			
			
			err = slsanpham1.Scan(&sanpham1.Giamgia, &sanpham1.Manhasanxuat, &sanpham1.Manhacungcap, &sanpham1.Ngaycapnhat, &sanpham1.Luotxem, &sanpham1.Luotbinhchon, &sanpham1.Soluongton, &sanpham1.Cauhinh,&sanpham1.Idsp1)
			if err != nil {
				panic(err.Error())
			}
		
			if sanpham1.Idsp1 == sanpham.Idsp{
				if sanpham1.Giamgia ==0{
					tien =strconv.Itoa(int(a)* int(b))
					mang33 =append(mang33,tien)
				}else{
					tien =strconv.Itoa(int(a)*sanpham1.Giamgia )
					mang33 =append(mang33,tien)
				}
			}
		}
		slsanpham1.Close()
		
		
		mang11 =append(mang11,strconv.Itoa(mang1[i]))
		mang22 =append(mang22,mang2[i])
		
		}
	}

}
slsanpham.Close()

var tientmp int64
tientmp,err = strconv.ParseInt(tien,10,64)
if err!=nil{
	panic(err.Error())
}
if int(tientmp) >=2000000{
	update,err :=db1.Query("update hoadon set phiship=0 where mahd='"+idhoadon+"'")
	if err != nil {
		panic(err.Error())
			}
			update.Close()
}
session.Values["idsanpham"] = " "
session.Values["soluong"]  = " "
session.Save(r,w) 


for i := 0; i < len(mang11); i++ {
	selectsp,err :=db1.Query("select soluongton from sanpham1 where idsp='"+mang11[i]+"'")

	var soluongtonsp int
				for selectsp.Next(){
					err = selectsp.Scan(&soluongtonsp)
				}
				selectsp.Close()
				isl,err :=strconv.ParseInt(mang22[i],10,64)
				if err != nil {
					panic(err.Error())
						}
				if err != nil {
					panic(err.Error())
						}
				if soluongtonsp <int(isl){
			
		
				}else{
					hoadon ,err := db1.Query("insert into itemhoadon values('','"+mang11[i]+"','"+mang22[i]+"','"+mang33[i]+"','"+idhoadon+"')")
					if err != nil {
						panic(err.Error())
							}
					hoadon.Close()
					i2,err :=strconv.ParseInt(mang22[i],10,64)
					if err != nil {
						panic(err.Error())
							}
				updatesp,err:=db1.Query("update sanpham1 set soluongton='"+strconv.Itoa(soluongtonsp-int(i2))+"' where idsp='"+mang11[i]+"'")
				if err != nil {
					panic(err.Error())
						}
				updatesp.Close()
				if err != nil {
					panic(err.Error())
						}
				}
			if err != nil {
				panic(err.Error())
					}

		
}		        
	
	
	
	}
	}

	}
return ""
}
func xoasanpham(w http.ResponseWriter, r *http.Request) {

	p := temp1{bien: "haha"}
	tpl.ExecuteTemplate(w, "quanlysanpham.html", p)
}
func thanhtoan(w http.ResponseWriter, r *http.Request) {
	session,_ := store.Get(r,"session")
	
	var a string
	if r.FormValue("ten")==""|| r.FormValue("ho")==""||r.FormValue("sdt") =="" || r.FormValue("huyen")==""||r.FormValue("diachi") =="" || r.FormValue("note")==""{
		// a="Vui Lòng Điền Đầy Đủ Thông Tin"
	}else{
		dt := time.Now()
		var kt int
		kt =0

		for kt == 0 {
			var mdh =String1(9)
			var dem int
			dem=0
			sl,err := db1.Query("SELECT * FROM `donhang` WHERE madh='"+mdh+"'")
			if err != nil {
				panic(err.Error())
			}
			for sl.Next(){
				dem=1
			}
			sl.Close()
			if dem ==0{
				
			insert,err := db1.Query("insert into donhang values('"+mdh+"','"+r.FormValue("ten")+"','"+ r.FormValue("ho")+"','"+r.FormValue("diachi")+"','"+r.FormValue("email")+"','"+r.FormValue("sdt")+"','"+r.FormValue("huyen")+"','"+r.FormValue("thanhpho")+"','"+string(0)+"','"+r.FormValue("note")+"','"+dt.Format("2006-01-02")+"')")
						if err != nil {
							panic(err.Error())
						}
						insert.Close()

			insertgiaohang,err := db1.Query("insert into giaohang values('"+mdh+"','0')")
			if err != nil {
				panic(err.Error())
			}
			insertgiaohang.Close()
					  var phiship int
						if session.Values["username"] !=nil{
							if r.FormValue("thanhpho") != "1"{
								phiship =25000
							}else
							{
								phiship =0
							}
							  
							hoadon ,err := db1.Query("insert into hoadon values('','"+mdh+"','"+ strconv.Itoa( session.Values["iduser"].(int))+"','"+strconv.Itoa(phiship)+"')")
							slhd,err := db1.Query("SELECT mahd FROM `hoadon` WHERE madh='"+mdh+"'")
							if err != nil {
								panic(err.Error())
							}
							var id int
							for slhd.Next(){
								var id1 int
								err = slhd.Scan(&id1)
								hoadon.Close()
								id=id1
							}
							laphoadon(w,r,strconv.Itoa(id))

						}else{
							if r.FormValue("thanhpho") != "1"{
								phiship =25000
							}else
							{
								phiship =0
							}
							  
							hoadon ,err := db1.Query("insert into hoadon values('','"+mdh+"','"+ "1"+"','"+strconv.Itoa(phiship)+"')")
							slhd,err := db1.Query("SELECT mahd FROM `hoadon` WHERE madh='"+mdh+"'")
							if err != nil {
								panic(err.Error())
							}
							var id int
							for slhd.Next(){
								var id1 int
								err = slhd.Scan(&id1)
						   id=id1
					
							 }
							 slhd.Close()
						hoadon.Close()
						fmt.Print("day la id")
						fmt.Print(id)
						fmt.Print("HET")
						laphoadon(w,r,strconv.Itoa(id))
						
					}
						kt=1
			}else{
			kt=0
			}
		}

	}

		
    var lsp [] loaisanpham

	slloaisanpham,err := db1.Query("SELECT * FROM `loaisanpham` WHERE 1 ")
	if err != nil{
		panic(err.Error())
	}
	for slloaisanpham.Next(){
		var loaisanpham loaisanpham
		err = slloaisanpham.Scan(&loaisanpham.Idloaisp,&loaisanpham.Tenloaisanpham,&loaisanpham.Mota)
		if err != nil {
			panic(err.Error())
		}
		lsp = append(lsp,loaisanpham)
	}
	if err != nil {
		panic(err.Error())
	}
	slloaisanpham.Close()

	var SS string
	if session.Values["username"]!=nil{

	var S= session.Values["username"]
	SS = S.(string)
	}else{
		SS=""
	}

	tpl,err:= template.New("static/checkout.html").Funcs(template.FuncMap{"cong":cong,"add": add}).ParseFiles("static/checkout.html")
	if err != nil {
		panic(err.Error())
	}
 
		tpl.ExecuteTemplate(w, "checkout.html",struct{
		Loaisanpham []loaisanpham
		Username string
		A string
		}{lsp,SS,a})
}

func sendmail(w http.ResponseWriter, r *http.Request){
	session,_ := store.Get(r,"session")
	hostURL := "smtp.gmail.com"
	hostPort := "587"
	emailSender := "myphamonline2019@gmail.com"
	password :="Myphamonlien2019"
	emailRCeveiver:= r.FormValue("email")
	emailAuth := smtp.PlainAuth(
		 "",
		 emailSender,
		 password,
		 hostURL,
	)
	var maxacnhan = String(6)
	session.Values["maxacnhan"]=maxacnhan
	session.Save(r,w) 


//    msg := []byte("To:"+emailRCeveiver+"\r\n"+"Subject:"+"Mã Xác Nhận"+"\r\n"+"<p>how are you doing<p>")
   mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n";
    subject := "Subject: Xác Nhận Email!\n"
   msg := []byte(subject + mime  +templatesendmain("")+"<a  target='_blank'>Đây Là Mã Xác Nhận Của Bạn <a></br><div target='_blank' style='font-size:30px;color:seashell'>"+maxacnhan+"</div>"+"</td> </tr> </table> </td> </tr> <tr class='footer'> <td style='padding: 40px;'> <a target='_blank'></a> </td> </tr> </table> </body> </html>")

   err := smtp.SendMail(
	   hostURL +":"+hostPort,
	   emailAuth,
	   emailSender,
	   []string{emailRCeveiver},
	   msg)
	   if err != nil {
		panic(err.Error())
	}

}
func kiemtrathongtin(w http.ResponseWriter, r *http.Request){
	
	
		email := r.FormValue("email")
		taikhoans,err := db1.Query("SELECT * FROM `taikhoan` WHERE email='"+email+"'")
	   if err != nil{
		 panic(err.Error())
	    }
	 var a int
	 a=0
	 for taikhoans.Next(){
		a+=1
		
	 }
	 if a==1 {
	 io.WriteString(w,"Email đã tồn tại")
	 
	 }
	 taikhoans.Close()
	
}
func  xacnhanmacode(w http.ResponseWriter, r *http.Request){
	session,_ := store.Get(r,"session")

 if r.FormValue("maxacnhan") == session.Values["maxacnhan"].(string){
	session.Values["maxacnhanlai"]=session.Values["maxacnhan"]
	session.Values["maxacnhan"]=nil
	session.Save(r,w) 
  io.WriteString(w,"xác nhận thành công")
 
 }else{
	io.WriteString(w,"Mã Không Chính Xác")
	
 }
}
func  funcdanhgia(w http.ResponseWriter, r *http.Request){
	session,_ := store.Get(r,"session")
	var b int
	
var 	a = session.Values["iduser"]
	b = a.(int)
	sl,err := db1.Query("select madanhgia from danhgia where iduser='"+strconv.Itoa(b)+"' and idsanpham='"+r.FormValue("idsanpham")+"'")
	if err != nil{
		panic(err.Error())
	}
	var dem int 
	dem =0
	var iddanhgia int
	for sl.Next(){
		
		err = sl.Scan(&iddanhgia)
		if err != nil{
			panic(err.Error)
		}
		sl.Close()
  dem =1
	}
	if dem ==0{
		var bienusername string
		var aaaa =session.Values["username"]
		bienusername=aaaa.(string)
		dt :=time.Now()
		insert,err := db1.Query("insert into  danhgia values('','"+r.FormValue("noidung")+"','"+r.FormValue("idsanpham")+"','"+strconv.Itoa(b)+"','"+bienusername+"','"+r.FormValue("sao")+"','"+dt.Format("2006-01-02")+"')")
		if err != nil{
			panic(err.Error())
		}
		insert.Close()
	
	}else{
		
   update, err := db1.Query("update danhgia set noidung='"+r.FormValue("noidung")+ "',sao='"+r.FormValue("sao")+ "' where madanhgia='"+  strconv.Itoa(iddanhgia)+"' and idsanpham='"+r.FormValue("idsanpham")+"'")
   if err != nil{
	panic(err.Error())
}
update.Close()
	}
	slsao,err := db1.Query("select sao from danhgia where idsanpham='" +r.FormValue("idsanpham")+"'")
	if err != nil{
		panic(err.Error())
	}
	var demsao int
	var tongbinhluan int
	for slsao.Next(){
		var biensao int
		err = slsao.Scan(&biensao)
		demsao+=biensao
		tongbinhluan+=1
	}
	slsao.Close()
	var tempsao int
	tempsao = int(math.Ceil(float64(demsao/tongbinhluan)))

	
	upslsao,err := db1.Query("UPDATE `sanpham1` SET luotbinhchon="+strconv.Itoa(tempsao)+" where idsp='" +r.FormValue("idsanpham")+"'")
	
	if err != nil{
		panic(err.Error)
	}
	upslsao.Close()

}

func  khungdangnhap(w http.ResponseWriter, r *http.Request){


	tpl,err:= template.New("static/dangnhap.html").Funcs(template.FuncMap{"cong":cong,"add": add}).ParseFiles("static/dangnhap.html")
	if err != nil {
		panic(err.Error())
	}
 
		tpl.ExecuteTemplate(w, "dangnhap.html","")
}


func  funcbinhluan(w http.ResponseWriter, r *http.Request){
	session,_ := store.Get(r,"session")
	var b string
	var bid int

	if r.FormValue("binhluan")!=""{

		
	
	if session.Values["iduser"]!=nil{
		dt := time.Now()
		var a = session.Values["username"]
		b = a.(string)

		var aid = session.Values["iduser"]
		bid = aid.(int)
	sl,err := db1.Query("insert into binhluan values('','"+r.FormValue("noidung")+ "','"+b+"','" +r.FormValue("email")+"','"+r.FormValue("sdt")+"','"+ dt.Format("2006-01-02")+"','"+r.FormValue("idsanpham")+"','"+ strconv.Itoa( bid)+"')  ")
	if err != nil{
		panic(err.Error())
	}
	sl.Close()
}else{
	dt := time.Now()
	
	sl,err := db1.Query("insert into binhluan values('','"+r.FormValue("noidung")+ "','"+r.FormValue("ten")+"','" +r.FormValue("email")+"','"+r.FormValue("sdt")+"','"+ dt.Format("2006-01-02")+"','"+r.FormValue("idsanpham")+"','"+"1"+"')")
	if err != nil{
		panic(err.Error())
	}
		sl.Close()
}}else{
	
	var aid = session.Values["iduser"]
	bid = aid.(int)
	var a = session.Values["username"]
		b = a.(string)
	if session.Values["iduser"]!=nil{
		dt := time.Now()
		
	sl,err := db1.Query("insert into itembinhluan values('','"+r.FormValue("noidung")+ "','"+b+"','" +r.FormValue("email")+"','"+r.FormValue("sdt")+"','"+ dt.Format("2006-01-02")+"','"+r.FormValue("idbinhluan")+"','"+strconv.Itoa(bid)+"') ")
	if err != nil{
		panic(err.Error())
	}
	sl.Close()
    }else{
	dt := time.Now()
		
	sl,err := db1.Query("insert into itembinhluan values('','"+r.FormValue("noidung")+ "','"+r.FormValue("ten")+"','" +r.FormValue("email")+"','"+r.FormValue("sdt")+"','"+ dt.Format("2006-01-02")+"','"+r.FormValue("idbinhluan")+"','"+"1')")
	if err != nil{
		panic(err.Error())
	}
	sl.Close()


}
}

}




func  quanlytaikhoan(w http.ResponseWriter, r *http.Request){
	session,_ := store.Get(r,"session")
	
	var SS  int
	
	
	if session.Values["iduser"]!=nil{
		var S= session.Values["iduser"]
		 SS = S.(int)
		
	
		 
	}else{SS=0

	}
	var SS2 string
	if session.Values["username"]!=nil{

	var S2= session.Values["username"]
	SS2 = S2.(string)
	}else{
		SS2=""
	}
	var lsp       []loaisanpham
	sllsp,err := db1.Query("SELECT * FROM `loaisanpham` WHERE 1 ")
	 if err != nil{
		 panic(err.Error())
	 }
	 for sllsp.Next(){
		 var loaisanpham loaisanpham
		 err = sllsp.Scan(&loaisanpham.Idloaisp,&loaisanpham.Tenloaisanpham,&loaisanpham.Mota)
		 if err != nil {
			 panic(err.Error())
		 }
		 lsp = append(lsp,loaisanpham)
	 }
	 sllsp.Close()

    if SS!=0{

				var arrtaikhoan [] taikhoan

				slloaisanpham,err := db1.Query("SELECT idtaikhoan,tenkhachhang,hokhachhang,email FROM `taikhoan` WHERE 1 and idtaikhoan='"+strconv.Itoa(SS)+"'")
				if err != nil{
					panic(err.Error())
				}
				for slloaisanpham.Next(){
					var taikhoan taikhoan
					err = slloaisanpham.Scan(&taikhoan.Idtaikhoan,&taikhoan.Tennguoidung,&taikhoan.Ho,&taikhoan.Email)
					if err != nil {
						panic(err.Error())
					}
					arrtaikhoan = append(arrtaikhoan,taikhoan)
				}
				slloaisanpham.Close()
				if err != nil {
					panic(err.Error())
				}

				var arrhoadon [] structhoadon

				hoadon1,err := db1.Query("SELECT * FROM `hoadon` where iduser='"+strconv.Itoa(SS)+"'")
				if err != nil{
					panic(err.Error())
				}
				for hoadon1.Next(){
					var hoadon structhoadon
					err = hoadon1.Scan(&hoadon.Mahd,&hoadon.Madh,&hoadon.Iduser,&hoadon.Ship)
					if err != nil {
						panic(err.Error())
					}
					arrhoadon = append(arrhoadon,hoadon)
				}
				hoadon1.Close()
				if err != nil {
					panic(err.Error())
				}
			var arrdonhang[] donhang
				donhang1,err := db1.Query("SELECT * FROM `donhang`")
				if err != nil{
					panic(err.Error())
				}
				for donhang1.Next(){
					var donhang donhang
					err = donhang1.Scan(&donhang.Madh,&donhang.Ten,&donhang.Ho,&donhang.Diachi,&donhang.Email,&donhang.Sdt,&donhang.Huyen,&donhang.TP,&donhang.Tinhtrang,&donhang.Note,&donhang.Ngaylap)
					if err != nil {
						panic(err.Error())
					}
					arrdonhang = append(arrdonhang,donhang)
				}
				donhang1.Close()
				if err != nil {
					panic(err.Error())
				}
				

				var arritemhoadon       []itemhoadon
				slihd,err := db1.Query("SELECT * FROM `itemhoadon` WHERE 1 ")
				if err != nil{
					panic(err.Error())
				}
				for slihd.Next(){
					var itemhoadon itemhoadon
					err = slihd.Scan(&itemhoadon.Iditemhoadon,&itemhoadon.Idsanpham,&itemhoadon.Soluong,&itemhoadon.Tongtien,&itemhoadon.Mahd)
					if err != nil {
						panic(err.Error())
					}
					arritemhoadon = append(arritemhoadon,itemhoadon)
				}
				slihd.Close()
				
				if err != nil {
					panic(err.Error())
				}


			
			if r.FormValue("capnhatdiachigiaohang")!=""{
				selectdiachi,err := db1.Query("select * from diachigiaohang where idkhachhang='"+strconv.Itoa(SS)+"'")
				if err != nil {
					panic(err.Error())
				}
				var bien int
				for selectdiachi.Next(){
				bien+=1
				}
				if bien >0{
				updatediachi,err := db1.Query("update diachigiaohang set tenkhachhang='"+ r.FormValue("ten")+"',hokhachhang='"+r.FormValue("ho")+"',thanhpho='"+r.FormValue("thanhpho")+"',huyen='"+ r.FormValue("huyen")+"',diachigiaohang='"+r.FormValue("diachi")+"',sodienthoai='"+r.FormValue("sdt")+"'  where idkhachhang='"+strconv.Itoa(SS)+"'" )
				if err != nil {
					panic(err.Error())
				}
				updatediachi.Close()

				}else{
					insertdiachi,err :=db1.Query("insert into  diachigiaohang values('','"+ r.FormValue("ten")+"','"+ r.FormValue("ho")+"','"+ r.FormValue("thanhpho")+"','"+ r.FormValue("huyen")+"','"+r.FormValue("diachi")+"','"+r.FormValue("sdt")+"','"+strconv.Itoa(SS)+"')")
					if err != nil {
						panic(err.Error())
					}
					insertdiachi.Close()
				
					}
			}




			var diachigiaohang1 []diachigiaohang



			selectdc,err := db1.Query("select * from diachigiaohang")
			if err != nil {
				panic(err.Error())
			}
			for selectdc.Next(){
				var diachigiaohang diachigiaohang
				err = selectdc.Scan(&diachigiaohang.Madiachi,&diachigiaohang.Tenkhachhang,&diachigiaohang.Hokhachhang,&diachigiaohang.Thanhpho,&diachigiaohang.Huyen,&diachigiaohang.Diachi,&diachigiaohang.Sodienthoai,&diachigiaohang.Idkhachhang)
				
				if err != nil {
					panic(err.Error())
				}
				diachigiaohang1 = append(diachigiaohang1,diachigiaohang)
			}
			selectdc.Close()
				
				if err != nil {
					panic(err.Error())
				}
				var arrSanpham [] sanpham
				slsanpham,err := db1.Query("SELECT * FROM `sanpham` WHERE 1 ")
				if err != nil{
					panic(err.Error())
				}
				for slsanpham.Next(){
					var sanpham sanpham
					err = slsanpham.Scan(&sanpham.Idsp, &sanpham.Tensanpham, &sanpham.Loaisanpham, &sanpham.Gia, &sanpham.Mota, &sanpham.Hinhanh1, &sanpham.Hinhanh2, &sanpham.Hinhanh3)
					if err != nil {
						panic(err.Error())
					}
					arrSanpham = append(arrSanpham,sanpham)
				}
				slsanpham.Close()

				
				tpl,err:= template.New("static/thongtintaikhoan.html").Funcs(template.FuncMap{"tinhgiasale":tinhgiasale,"add": add,"floatchuoi":floatchuoi}).ParseFiles("static/thongtintaikhoan.html")
				
					tpl.ExecuteTemplate(w, "thongtintaikhoan.html",struct{
					Taikhoan []taikhoan
					Username string
					Hoadon []structhoadon
					Loaisanpham []loaisanpham
					Itemhoadon    []itemhoadon
					Iduser   int
					Donhang    []donhang
					Diachigiaohang []diachigiaohang
					Sanpham []sanpham
					}{arrtaikhoan,SS2,arrhoadon,lsp,arritemhoadon,SS,arrdonhang,diachigiaohang1,arrSanpham})
					if err != nil {
						panic(err.Error())
					}

	}else{

		tpl, err := template.ParseFiles("static/vuilongdangnhap.html")
		tpl.ExecuteTemplate(w, "vuilongdangnhap.html",struct{

		Username string

		Loaisanpham []loaisanpham
		}{SS2,lsp})
		if err != nil {
			panic(err.Error())
		}

	}
}

func thaydoimatkhau(w http.ResponseWriter, r *http.Request){

	session,_ := store.Get(r,"session")
	
	var SS  int
	
	
	if session.Values["iduser"]!=nil{
		var S= session.Values["iduser"]
		 SS = S.(int)
		
	
		 
	}else{SS=0

	}
	if r.FormValue("thaydoimatkhau")!=""{
		matkhaucu := r.FormValue("matkhaucu")
		matkhaumoi := r.FormValue("matkhaumoi")
		if strings.Count(matkhaucu,"'")==0{
		
	
			taikhoans,err := db1.Query("SELECT * FROM `taikhoan` WHERE idtaikhoan='"+strconv.Itoa(SS)+"' and matkhau='"+matkhaucu+"'")
			if err != nil{
				panic(err.Error())
			}
			var dem int
			dem=0
			for taikhoans.Next(){
			  dem=1
			}
		   
			if dem==1{
				
			 updatetaikhoan,err := db1.Query("update taikhoan set matkhau='"+matkhaumoi+"' where idtaikhoan='"+strconv.Itoa(SS)+"'")
			 if err != nil{
				panic(err.Error())
			}
			 updatetaikhoan.Close()
			io.WriteString(w,"Thay Đổi Mật Khẩu Thành Công")
			
			}else{
			
				io.WriteString(w,"Mật Khẩu Cũ Không Chính Xác")
			}
			taikhoans.Close()
		}else{
			io.WriteString(w,"Hệ Thống Bó Tay ^^")
		}
	  }
}

func lienhechungtoi(w http.ResponseWriter, r *http.Request){


	if r.Method=="POST"{
		
	hostURL := "smtp.gmail.com"
	hostPort := "587"
	emailSender := "myphamonline2019@gmail.com"
	password :="Myphamonlien2019"
	emailRCeveiver:= "kingminhluan9@gmail.com"
	emailAuth := smtp.PlainAuth(
		 "",
		 emailSender,
		 password,
		 hostURL,
	)




   mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n";
    subject := "Subject: !"+r.FormValue("contactSubject")+"\n"
   msg := []byte(subject + mime  +"Khách Hàng "+r.FormValue("customerName")+" Đóng Góp Ý kiến :<p>"+r.FormValue("contactMessage")+"</p> <p> liên hệ trả lời qua email:   "+r.FormValue("customerEmail")+"</p>")

   err := smtp.SendMail(
	   hostURL +":"+hostPort,
	   emailAuth,
	   emailSender,
	   []string{emailRCeveiver},
	   msg)
	   if err != nil {
		panic(err.Error())
	}
	}
	  
	session,_ := store.Get(r,"session")
  var lsp [] loaisanpham

	slloaisanpham,err := db1.Query("SELECT * FROM `loaisanpham` WHERE 1 ")
	if err != nil{
		panic(err.Error())
	}
	for slloaisanpham.Next(){
		var loaisanpham loaisanpham
		err = slloaisanpham.Scan(&loaisanpham.Idloaisp,&loaisanpham.Tenloaisanpham,&loaisanpham.Mota)
		if err != nil {
			panic(err.Error())
		}
		lsp = append(lsp,loaisanpham)
	}
	if err != nil {
		panic(err.Error())
	}
	slloaisanpham.Close()

	var SS string
	if session.Values["username"]!=nil{

	var S= session.Values["username"]
	SS = S.(string)
	}else{
		SS=""
	}
	tpl,err:= template.New("static/contact.html").Funcs(template.FuncMap{"cong":cong,"add": add}).ParseFiles("static/contact.html")
	if err != nil {
		panic(err.Error())
	}
 
		tpl.ExecuteTemplate(w, "contact.html",struct{
			Loaisanpham []loaisanpham
			Username string
		}{lsp,SS})

}

func thongtinvechungtoi(w http.ResponseWriter, r *http.Request){

	var soluongsanpham int 
	slsp,err := db1.Query("select count(*) from sanpham")
	for slsp.Next(){
		err=slsp.Scan(&soluongsanpham)
	}
	var soluongtruycap int 
	sltc,err := db1.Query("select count(*) from ipnguoidung")
	for sltc.Next(){
		err=sltc.Scan(&soluongtruycap)
	}

	var soluonguser int 
	slus,err := db1.Query("select count(*) from taikhoan where loaitaikhoan=3")
	for slus.Next(){
		err=slus.Scan(&soluonguser)
	}
	var soluongdanhgia int
	sldg,err := db1.Query("select count(*) from danhgia ")
	for sldg.Next(){
		err=sldg.Scan(&soluongdanhgia)
	}
	  
	session,_ := store.Get(r,"session")
  var lsp [] loaisanpham

	slloaisanpham,err := db1.Query("SELECT * FROM `loaisanpham` WHERE 1 ")
	if err != nil{
		panic(err.Error())
	}
	for slloaisanpham.Next(){
		var loaisanpham loaisanpham
		err = slloaisanpham.Scan(&loaisanpham.Idloaisp,&loaisanpham.Tenloaisanpham,&loaisanpham.Mota)
		if err != nil {
			panic(err.Error())
		}
		lsp = append(lsp,loaisanpham)
	}
	if err != nil {
		panic(err.Error())
	}
	slloaisanpham.Close()

	var SS string
	if session.Values["username"]!=nil{

	var S= session.Values["username"]
	SS = S.(string)
	}else{
		SS=""
	}

	tpl,err:= template.New("static/about-us.html").Funcs(template.FuncMap{"cong":cong,"add": add}).ParseFiles("static/about-us.html")
	if err != nil {
		panic(err.Error())
	}
 
		tpl.ExecuteTemplate(w, "about-us.html",struct{
			Soluongsanpham int
			Soluottruycap int
			Soluongkhachhang int
			Soluotdanhgia int
			Loaisanpham []loaisanpham
			Username string
		}{soluongsanpham,soluongtruycap,soluonguser,soluongdanhgia,lsp,SS})
}


func laylaimatkhau(w http.ResponseWriter, r *http.Request){


	selectmk,err := db1.Query("select matkhau from taikhoan where email='"+r.FormValue("email")+"'")
	if err!=nil{
		panic(err.Error())
	}
	var matkhau string
	matkhau=""
	for selectmk.Next(){
		err =selectmk.Scan(&matkhau)
	}
	if matkhau!=""{
	if r.Method=="POST"{
		
	hostURL := "smtp.gmail.com"
	hostPort := "587"
	emailSender := "myphamonline2019@gmail.com"
	password :="Myphamonlien2019"
	emailRCeveiver:= r.FormValue("email")
	emailAuth := smtp.PlainAuth(
		 "",
		 emailSender,
		 password,
		 hostURL,
	)




   mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n";
    subject := "Subject: ! Lấy Lại Mật Khẩu"+"\n"
   msg := []byte(subject + mime  +"<p> Xin Chào Đây Là Mật Khẩu Của Bạn:  "+"<div > Mật Khẩu: <input type='text' style='width=100px;height:30px' value='"+matkhau+"' id='myInput'>" +"</div>" +"<div></div>  <p>Nếu Bạn Không Có Yêu Cầu Lấy Lại Mật Khẩu Hãy Liên Hệ Với Chúng tôi qua Email: yenngan1990@gmail.com</p> </p> ")

   err := smtp.SendMail(
	   hostURL +":"+hostPort,
	   emailAuth,
	   emailSender,
	   []string{emailRCeveiver},
	   msg)
	   if err != nil {
		panic(err.Error())
	}
	}
	
	}
}

func htmllaylaimatkhau(w http.ResponseWriter, r *http.Request){
	tpl,err:= template.New("static/laylaimatkhau.html").Funcs(template.FuncMap{"tinhgiasale":tinhgiasale,"add": add,"floatchuoi":floatchuoi}).ParseFiles("static/laylaimatkhau.html")
	if err != nil{
		panic(err.Error())
	}
	tpl.ExecuteTemplate(w, "laylaimatkhau.html", "")
}

func printbill(w http.ResponseWriter, r *http.Request){

	session,_ := store.Get(r,"session")
	
	var SS  int
	

	if session.Values["iduser"]!=nil{
		var S= session.Values["iduser"]
		 SS = S.(int)
		
	
		 
	}else{SS=0

	}

var kiemtra int 
soluongkt,err := db1.Query("SELECT * FROM `hoadon`,donhang WHERE hoadon.madh=donhang.madh and hoadon.iduser='"+strconv.Itoa(SS)+"'")
if err != nil{
	panic(err.Error())
}
for soluongkt.Next(){
	kiemtra=1
}
if kiemtra==1{
						var arrsanpham1 [] sanpham1
						slsanpham1,err := db1.Query("SELECT * FROM `sanpham1` WHERE 1 ORDER BY ngaycapnhat desc LIMIT 15 ")
						if err != nil{
							panic(err.Error())
						}
						for slsanpham1.Next(){
							var sanpham1 sanpham1

							
							
							err = slsanpham1.Scan(&sanpham1.Giamgia, &sanpham1.Manhasanxuat, &sanpham1.Manhacungcap, &sanpham1.Ngaycapnhat, &sanpham1.Luotxem, &sanpham1.Luotbinhchon, &sanpham1.Soluongton, &sanpham1.Cauhinh,&sanpham1.Idsp1)
							if err != nil {
								panic(err.Error())
							}
						
							arrsanpham1 = append(arrsanpham1,sanpham1)
						}

						slsanpham1.Close()

					var arrSanpham [] sanpham
						slsanpham,err := db1.Query("SELECT * FROM `sanpham` WHERE 1 ")
						if err != nil{
							panic(err.Error())
						}
						for slsanpham.Next(){
							var sanpham sanpham
							err = slsanpham.Scan(&sanpham.Idsp, &sanpham.Tensanpham, &sanpham.Loaisanpham, &sanpham.Gia, &sanpham.Mota, &sanpham.Hinhanh1, &sanpham.Hinhanh2, &sanpham.Hinhanh3)
							if err != nil {
								panic(err.Error())
							}
							arrSanpham = append(arrSanpham,sanpham)
						}
						slsanpham.Close()
						var mahoadon string
						mahoadon =r.FormValue("mahoadon")
						s,err:=strconv.ParseInt(mahoadon,10,64)
						


						var arrdonhang[] donhang
						donhang1,err := db1.Query("SELECT * FROM `donhang` where madh='"+strconv.Itoa(int(s))+"'")
						if err != nil{
							panic(err.Error())
						}
						for donhang1.Next(){
							var donhang donhang
							err = donhang1.Scan(&donhang.Madh,&donhang.Ten,&donhang.Ho,&donhang.Diachi,&donhang.Email,&donhang.Sdt,&donhang.Huyen,&donhang.TP,&donhang.Tinhtrang,&donhang.Note,&donhang.Ngaylap)
							if err != nil {
								panic(err.Error())
							}
							arrdonhang = append(arrdonhang,donhang)
						}
						donhang1.Close()
						if err != nil {
							panic(err.Error())
						}


						if err!=nil{
							panic(err.Error())
						}
						selecthd,err := db1.Query("select mahd from hoadon where madh='"+strconv.Itoa(int(s))+"' and iduser='"+strconv.Itoa(SS)+"'") 
						if err!=nil{
							panic(err.Error())
						}
						var mahd int
						for selecthd.Next(){
							err=selecthd.Scan(&mahd)
						}

					selecthd.Close()

						var arritemhoadon   []itemhoadon
						slihd,err := db1.Query("SELECT * FROM `itemhoadon` WHERE mahd='"+strconv.Itoa(mahd)+"'")
						if err != nil{
							panic(err.Error())
						}
						for slihd.Next(){
							var itemhoadon itemhoadon
							err = slihd.Scan(&itemhoadon.Iditemhoadon,&itemhoadon.Idsanpham,&itemhoadon.Soluong,&itemhoadon.Tongtien,&itemhoadon.Mahd)
							if err != nil {
								panic(err.Error())
							}
							arritemhoadon = append(arritemhoadon,itemhoadon)
						}
						slihd.Close()
						tpl,err:= template.New("static/printbill.html").Funcs(template.FuncMap{"tinhgiasale":tinhgiasale,"add": add,"floatchuoi":floatchuoi}).ParseFiles("static/printbill.html")
						tpl.ExecuteTemplate(w, "printbill.html", struct{
							Itemhoadon [] itemhoadon
							Donhang [] donhang
							Madonhang string
						Sanpham []sanpham
						Sanpham1 []sanpham1
							}{arritemhoadon,arrdonhang,mahoadon,arrSanpham,arrsanpham1})
					}else{
						tpl,err:= template.New("static/404.html").Funcs(template.FuncMap{"tinhgiasale":tinhgiasale,"add": add,"floatchuoi":floatchuoi}).ParseFiles("static/404.html")
						if err != nil{
							panic(err.Error())
						}
						tpl.ExecuteTemplate(w, "404.html", "")
					}

						

	  
}

const charset = "0123456789abcdefghijklmnopqrstuvwxyz" +
  "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const char ="0123456789"
var seededRand *rand.Rand = rand.New(
  rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
  b := make([]byte, length)
  for i := range b {
    b[i] = charset[seededRand.Intn(len(charset))]
  }
  return string(b)
}

func String(length int) string {
  return StringWithCharset(length, charset)
}
func String1(length int) string {
	return StringWithCharset(length, char)
  }


  func quanlyip(w http.ResponseWriter, r *http.Request) {
	var dem int
	sllsp,err := db1.Query("SELECT * FROM `ipnguoidung` WHERE diachi='"+r.FormValue("ip")+"'")
	 if err != nil{
		 panic(err.Error())
	 }
     if r.FormValue("de")=="10"{
		deleteip,err := db1.Query("delete  FROM `ipnguoidung` WHERE 1 ")
		if err != nil{
			panic(err.Error())
		}
		deleteip.Close()
	 }

	 for sllsp.Next(){
		dem+=1
		 if err != nil {
			 panic(err.Error())
		 }
		 
	 }
	 dt := time.Now()
	
if dem==0{
	sllsp,err := db1.Query("insert into ipnguoidung values('"+r.FormValue("ip")+"','"+dt.Format("2006-01-02")+"')")
	if err != nil{
		panic(err.Error())
	}
	sllsp.Close()
}



  }



  func huydonhang(w http.ResponseWriter, r *http.Request) {

	if r.FormValue("madonhang") !=""{
		selecta, err := db1.Query("select tinhtrang from giaohang where madonhang='"+r.FormValue("madonhang")+"'")
		if err!=nil{
			panic(err.Error())
		}
		var tinhtrang int
		for selecta.Next(){
		
			err = selecta.Scan(&tinhtrang)
			if err!=nil{
				panic(err.Error())
			}
		}
		selecta.Close()
		if tinhtrang !=1{
			
			deletegiaohang,err:=db1.Query("delete from giaohang where madonhang='"+r.FormValue("madonhang")+"'")
			if err!=nil{
				panic(err.Error())
			}
			deletegiaohang.Close()
			

			
			
			deletedonhang,err := db1.Query("delete from donhang where madh='"+r.FormValue("madonhang")+"'")
			if err!=nil{
				panic(err.Error())
			}
			deletedonhang.Close()


			selectmahoadon,err := db1.Query("select * from hoadon where madh='"+r.FormValue("madonhang")+"'")
			if err!=nil{
				panic(err.Error())
			}
			var mahoadon string
            for selectmahoadon.Next(){
				err = selecta.Scan(&mahoadon)
			}

			selectmahoadon.Close()


			deletehoadon,err := db1.Query("delete from hoadon where mahd='"+mahoadon+"'")
			if err!=nil{
				panic(err.Error())
			}
			deletehoadon.Close()

			deletehoadonitem,err := db1.Query("delete from itemhoadon where mahd='"+mahoadon+"'")
			if err!=nil{
				panic(err.Error())
			}
			deletehoadonitem.Close()

			io.WriteString(w,"Đơn Hàng Đã Được Hủy ^^")
		
		}else{
         
			io.WriteString(w,"Đơn Hàng Đã Được Giao Cho Nhà Vận Chuyển")
		}
	}
  }

  
  func dangnhapquanly(w http.ResponseWriter, r *http.Request) {
	session,_ := store.Get(r,"session")
	
	var SS  int
	
	
	if session.Values["iduser"]!=nil{
		var S= session.Values["iduser"]
		 SS = S.(int)
		
	
		 
	}else{SS=0

	}

	selectloaitaikhoan,err := db1.Query("select loaitaikhoan from taikhoan where idtaikhoan='"+strconv.Itoa(SS)+"'")

	if err!=nil{
		panic(err.Error())
	}
	var loaitaikhoanxacnhan int
	for selectloaitaikhoan.Next(){
	 err = selectloaitaikhoan.Scan(&loaitaikhoanxacnhan)
	 if err!=nil{
		panic(err.Error())
	}
	}
if loaitaikhoanxacnhan !=3  && SS !=0 {

	tpl,err:= template.New("static/page-login.html").Funcs(template.FuncMap{"tinhgiasale":tinhgiasale,"add": add,"floatchuoi":floatchuoi}).ParseFiles("static/page-login.html")
	if err != nil{
		panic(err.Error())
	}
	tpl.ExecuteTemplate(w, "page-login.html", "")

}else{
	tpl,err:= template.New("static/chuyentrang.html").Funcs(template.FuncMap{"tinhgiasale":tinhgiasale,"add": add,"floatchuoi":floatchuoi}).ParseFiles("static/chuyentrang.html")
	if err != nil{
		panic(err.Error())
	}
	tpl.ExecuteTemplate(w, "chuyentrang.html", "")

}
  }


  func thaydoimatkhauquanly(w http.ResponseWriter, r *http.Request) {
	session,_ := store.Get(r,"session")
	
	var SS  int
	
	
	if session.Values["iduser"]!=nil{
		var S= session.Values["iduser"]
		 SS = S.(int)
		
	
		 
	}else{SS=0

	}

	selectloaitaikhoan,err := db1.Query("select loaitaikhoan from taikhoan where idtaikhoan='"+strconv.Itoa(SS)+"'")

	if err!=nil{
		panic(err.Error())
	}
	var loaitaikhoanxacnhan int
	for selectloaitaikhoan.Next(){
	 err = selectloaitaikhoan.Scan(&loaitaikhoanxacnhan)
	 if err!=nil{
		panic(err.Error())
	}
	}
if loaitaikhoanxacnhan !=3  && SS !=0 {
	tpl,err:= template.New("static/thaydoimatkhau.html").Funcs(template.FuncMap{"tinhgiasale":tinhgiasale,"add": add,"floatchuoi":floatchuoi}).ParseFiles("static/thaydoimatkhau.html")
	if err != nil{
		panic(err.Error())
	}
	tpl.ExecuteTemplate(w, "thaydoimatkhau.html", "")
	}else{
		tpl,err:= template.New("static/chuyentrang.html").Funcs(template.FuncMap{"tinhgiasale":tinhgiasale,"add": add,"floatchuoi":floatchuoi}).ParseFiles("static/chuyentrang.html")
		if err != nil{
			panic(err.Error())
		}
		tpl.ExecuteTemplate(w, "chuyentrang.html", "")
	
	}
  }
func main() {


    
	db1 = moketnoi()
	statement, _ := db1.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
    statement.Exec()
	fmt.Println("Kết nối cơ sở dữ liệu thành công ^_^")
	// session,_ := store.Get(r,"session")
	http.HandleFunc("/hoadoncuaban",printbill)
	http.HandleFunc("/xacnhancode",xacnhanmacode)
	http.HandleFunc("/trangchu",indexHandler)
	http.HandleFunc("/quanlynhasanxuat", themnhasanxuat)
	http.HandleFunc("/loaisanpham",thongtinloaisanpham)
	http.HandleFunc("/chitietsanpham", themchitietsanpham)
	http.HandleFunc("/thongtinsanpham", chitietthongtinsanpham)
	http.HandleFunc("/quanlynhacungcap", themnhacungcap)
	http.HandleFunc("/xoasanpham", xoasanpham)
	http.HandleFunc("/quanlysanpham", themsanpham)
	http.HandleFunc("/quanlyloaisanpham", themloaisanpham)
	http.HandleFunc("/themthuoctinnh", themthuoctinhloaisanpham)
	http.HandleFunc("/themvaogiohang",themvaogiohang)
	http.HandleFunc("/giohang",danhsachsanphamtronggiohang)
	http.HandleFunc("/dangnhap",dangnhaptaikhoan)
	http.HandleFunc("/xacnhantaikhoan",dangnhapvadangkytaikhoan)
	http.HandleFunc("/viewthongtinsanpham",viewthongtinsanpham)
	http.HandleFunc("/sendmail",sendmail)
	http.HandleFunc("/kiemtrathongtin",kiemtrathongtin)
	http.HandleFunc("/cart",viewcart)
	http.HandleFunc("/giohangcuatoi",cart)
	http.HandleFunc("/hoadon",hoadon)
	http.HandleFunc("/thanhtoan",thanhtoan)
	http.HandleFunc("/danhgia",funcdanhgia)
	http.HandleFunc("/binhluan",funcbinhluan)
	http.HandleFunc("/loadcomment",binhluansanpham)
	http.HandleFunc("/loadcomments",binhluansanphams)
	http.HandleFunc("/khungdangnhap",khungdangnhap)
	http.HandleFunc("/loadheader",loadheader)
	http.HandleFunc("/quanlytaikhoan",quanlytaikhoan)
	http.HandleFunc("/quanlybinhluan",quanlybinhluan)
	http.HandleFunc("/quanlydonhang",quanlydonhang)
	http.HandleFunc("/quanlyuser",quanlyuser)
	http.HandleFunc("/xulyip",quanlyip)
	http.HandleFunc("/quanly",quanlythongke)
	http.HandleFunc("/lienhechungtoi",lienhechungtoi)
	http.HandleFunc("/thongtinvechungtoi",thongtinvechungtoi)
	http.HandleFunc("/thaydoimatkhau",thaydoimatkhau)
	http.HandleFunc("/thaydoimatkhauquanly",thaydoimatkhauquanly)
	http.HandleFunc("/huydonhang",huydonhang)
	http.HandleFunc("/dangnhapquanly",dangnhapquanly)
	http.HandleFunc("/laylaimatkhau",laylaimatkhau)
	http.HandleFunc("/quenmatkhau",htmllaylaimatkhau)
	
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("static"))))
	
	port := os.Getenv("PORT")
	if port ==""{

	 port ="80"
	}
	http.ListenAndServe(":"+port, nil)

}

func templatesendmain(string string) string{
	string ="<!DOCTYPE html PUBLIC '-//W3C//DTD XHTML 1.0 Transitional//EN' 'http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd'><html xmlns='http://www.w3.org/1999/xhtml'> <head><meta name='viewport' content='width=device-width'/><meta http-equiv='Content-Type' content='text/html; charset=UTF-8' /><title></title><style type='text/css'> body{ margin: 0 auto; padding: 0; min-width: 100%; font-family: sans-serif; } table{ margin: 50px 0 50px 0; } .header{ height: 40px; text-align: center; text-transform: uppercase; font-size: 24px; font-weight: bold; } .content{ height: 100px; font-size: 18px; line-height: 30px; } .subscribe{ height: 70px; text-align: center; } .button{ text-align: center; font-size: 18px; font-family: sans-serif; font-weight: bold; padding: 0 30px 0 30px; } .button a{ color: #FFFFFF; text-decoration: none; } .buttonwrapper{ margin: 0 auto; } .footer{ text-transform: uppercase; text-align: center; height: 40px; font-size: 14px; font-style: italic; } .footer a{ color: #000000; text-decoration: none; font-style: normal; } </style> </head> <body bgcolor='#009587'> <table bgcolor='#FFFFFF' width='100%' border='0' cellspacing='0' cellpadding='0'><tr class='header'> <td style='padding: 40px;'> Mã Có Hiệu Lực Trong 60s ! </td> </tr> <tr class='subscribe'> <td style='padding: 20px 0 0 0;'> <table bgcolor='#009587' border='0' cellspacing='0' cellpadding='0' class='buttonwrapper'> <tr> <td class='button' height='105'>  "
	return string
}