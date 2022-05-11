package main
import (
	"fmt"
	"os"
	"strings"
	"math/rand"
	"encoding/json"
	enk "encoding/base64"
	"time"
)
var RootDir = ".Cf"
var Home string
func version(){
	fmt.Println("v1.0")
}
func main(){
	home,_ := os.UserHomeDir()
	if _,err := os.Stat(home+"/"+RootDir);err != nil {
		os.Mkdir(home+"/"+RootDir,0700)
		Home = home+"/"+RootDir
	}
	Home = home +"/"+RootDir
	stdout := ""
	out := &stdout
	//i := true
	if len(os.Args) == 1 {
		return
	}
	switch os.Args[1]{
		case "shell" :
			shell()
		case "version" :
			version()
		case "login" :
			Login()
		case "me" :
			fmt.Println("<",me())
		case "sign":
			fmt.Print("username : ")
			username := ""
			fmt.Scanf("%s",&username)
			pass := uuid()
			Sign(username,&pass)
			if pass == "" {
				return
			}
			fmt.Println(username,"uuid is",pass)
		case "uuid":
			fmt.Print("username : ")
			username := ""
			fmt.Scanf("%s",&username)
			if getuuid(username) == "" {
				fmt.Println("uuid not found")
				return
			}
			fmt.Println("uuid is",getuuid(username))
		case "help":
			help := []string{
				"help -- display this message",
				"uuid -- get uuid for specific user",
				"sign -- register new user to tuxy",
				"me -- display who use tuxy now",
				"login -- for user login",
				"shell -- open tuxy shell",
				"version -- display tuxy version",
			}
			fmt.Println("usage command : ")
			for _,v := range help {
				fmt.Println("\t"+v)
			}
		default :
			*out = "no command"
			fmt.Println(stdout)
	}
}

func shell(){
	fmt.Println("Welcome to TUXY")
	i := true
	for i == true {
		Command := ""
		fmt.Print("> ")
		fmt.Scanf("%s",&Command)
		Handlers(Command,&i)
	}
}
func Handlers(command string,i *bool){
	//fmt.Println(">_",command)
	cmd := strings.Split(command,".")
	if len(cmd) == 1 {
		return
	}
	if cmd[0] == "tuxy" {
		tuxy(cmd[1],i)
	}else if cmd[0] == "self" {
		self(cmd[1])
	}else if cmd[0] == "task" {
		task(cmd)
	}
}
func tuxy(cmd string,i *bool){
	if cmd == "exit()" {
		confirm := ""
		fmt.Print("< are you sure to exit ? (y/n) ")
		fmt.Scanf("%s",&confirm)
		if confirm == "n" {
			*i = true
			return
		}
		*i = false
	} else if cmd == "version()" {
		fmt.Println("< v1.2")
	} else if cmd == "help()" {
		fmt.Println("< parent.function()\n\nfor parent :\n")
		parent := []string{
			"tuxy -- this parrent usage for shell command",
			"task -- this parrent for manage task",
			"self -- this parrent for agrv handle",
			//"bash -- this parrent to exec bash command",
		}
		TuxyFunc :=[]string{
			".exit() -- this function for exit from tuxy shell",
			".version() -- this function to diplays version of shell",
			".help() -- this function for display this message",
			".init() -- this function for create dir for task usage",
		}
		//TaskFunc :=[]string{}
		//SelfFunc :=[]string{}
		//BashFunc :=[]string{}
		for _,v := range parent {
			fmt.Println("\t"+v)
		}
		fmt.Println("\n< for function :\n")
		for _,v := range TuxyFunc {
			fmt.Println("\ttuxy - "+v)
		}
		fmt.Println("< use parrent.help() to display available command on other parent")
	}else if cmd == "init()"{
		fmt.Println("< create root dir for task,history and status")
		initDir()
		fmt.Println("< initilized...")
		fmt.Println("< done")
	}else {
		fmt.Println("< not tuxy command",cmd,"or you not call it ?")
		fmt.Println("< type tuxy.help() for command information")
	}
}
func self(cmd string){
	if cmd == "help()" {
		fmt.Println("< Command is ...")
		command := []string{
			//"self - .shell() -- for entering tuxy shell",
			//"self - .kill() -- for close current terminal",
			//"self - .history() -- for display history for current user",
			"self - .login() -- for login with your username and uuid",
			"self - .uuid() -- for get uuid for a username",
			"self - .sign() -- for add a username",
			"self - .me() -- print the user use now",
		}
		for _,v := range command {
			fmt.Println("\t",v)
		}
	}else if cmd == "sign()"{
		fmt.Print("< username To Sign : ")
		username := ""
		fmt.Scanf("%s",&username)
		Id := uuid()
		Sign(username,&Id)
		if Id == "" {
			return
		}
		fmt.Println("< uuid created",Id)
	}else if cmd == "me()"{
		fmt.Println("<",me())
	}else if cmd == "login()"{
		Login()
	}else if cmd == "uuid()"{
		username := ""
		fmt.Print("< username : ")
		fmt.Scanf("%s",&username)
		if getuuid(username) != ""{
			fmt.Println("< uuid for",username,"is",getuuid(username))
		}else {
			fmt.Println("< uuid for",username,"not found")
		}
	}
}
func uuid() string {
	rand.Seed(time.Now().UnixNano())
	id := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	S := make([]rune,12)
	for i := range S {
		S[i] = id[rand.Intn(len(id))]
	}
	return string(S)
}
func Sign(username string,id *string) interface{} {
	var db []interface{}
	if _,err := os.Stat(Home+"/all-user");err == nil {
		var user []interface{}
		Users,_ := os.ReadFile(Home+"/all-user")
		Str,_ := enk.StdEncoding.DecodeString(string(Users))
		Users = []byte(Str)
		json.Unmarshal(Users,&user)
		//fmt.Println(err)
		//fmt.Println(string(Users),user)
		db = user
		//db = db.([]interface{})
		for _,i := range db {
			if i.(map[string]interface{})["User"] == username {
				fmt.Println(username,"is Exists")
				username = ""
				*id = ""
				return nil
			}
		}
	}
	if username != "" {
		var user map[string]string
		user = map[string]string{"User":username,"Id":*id}
		db = append(db,user)
		JSONSTR,_ := json.Marshal(db)
		JSONSTRs := enk.StdEncoding.EncodeToString([]byte(JSONSTR))
		//fmt.Println(Home)
		os.WriteFile(Home+"/all-user",[]byte(JSONSTRs),0700)
		//fmt.Println("> user",username,"succes added")
	}
	return db
}
func initDir() interface{} {
	x,_ := os.UserHomeDir()
	y := ".tuxy"
	//fmt.Println(Home)
	if _,err := os.Stat(Home+"/user");err != nil {
		//fmt.Println(x)
		//fmt.Println("> no user found\n> using tuxy default")
		os.Mkdir(x+"/"+y,0700)
		return nil
	}else{
		c,_ := os.ReadFile(Home+"/user")
		c,_ = enk.StdEncoding.DecodeString(string(c))
		//fmt.Println(string(c))
		//fmt.Println("> using",strings.Replace(string(c),"\n","",1))
		user := strings.Replace(string(c),"\n","",1)
		os.Mkdir(x+"/.tuxy_user",0700)
		path := ".tuxy_user/"+user+".tuxy"
		os.Mkdir(x+"/"+path,0700)
		return user
	}
}
func me() string {
	if user := initDir();user != nil{
		//fmt.Print(user)
		//users,err:= enk.StdEncoding.DecodeString(user.(string))
		//if err != nil {
			//fmt.Println(err)
		//return
		//}
		//fmt.Println("<",user)
		return user.(string)
	}
	//fmt.Println("< tuxy")
	return "tuxy"
}
func Login(){
	var user string
	var pass string
	fmt.Print("< Username : ")
	fmt.Scanf("%s",&user)
	fmt.Print("< Uuid : ")
	fmt.Scanf("%s",&pass)
	//fmt.Print(user)
	Ids := ""
	users := Sign("",&Ids)
	if users == nil {
		fmt.Println("< file all-user not found")
		return
	}
	//fmt.Println(users)
	//verify
	Data := users.([]interface{})
	for _,i := range Data {
		if i.(map[string]interface{})["User"] == user && i.(map[string]interface{})["Id"] == pass {
			os.WriteFile(Home+"/user",[]byte(enk.StdEncoding.EncodeToString([]byte(user))),0700)
			fmt.Println("< user",user,"login success")
			return
		}
		//fmt.Println("> user or uuid wrong")
	}
	fmt.Println("< user or uuid wrong")
}
func getuuid(username string) string {
	id := ""
	allUser := Sign("",&id)
	allUsers := allUser.([]interface{})
	for _,v := range allUsers {
		if v.(map[string]interface{})["User"] == username {
			return v.(map[string]interface{})["Id"].(string)
		}
	}
	return ""
}
func task(cmd []string){
	command := cmd
	if command[1] == "Create()" {
		DailyTask()
	}else if command[1] == "List()" {
		listDailyTask("")
	}else if command[1] == "Percent()"{
		percent()
	}else if command[1] == "Help()"{
		helper()
	}
	if len(command) <= 2 {
		//fmt.Print("hai")
		return
	}
	if command[2] == "Detail()" {
		Detail(command[1])
	}
	if command[2] == "Solve()"{
		solve(command[1])
	}
	if command[2] == "Resolve()"{
		resolve(command[1])
	}
}
func DailyTask(){
	//"Daily/daily-Rabu20092004.json"
	Button := true
	type Task struct {
		Name string
		Uuid string
		In string
	}
	var AllTask []interface{}
	name := ""
	for Button {
		fmt.Print("< task name : ")
		fmt.Scanf("%s",&name)
		NewTask := Task{Name:name,Uuid:uuid(),In:"no"}
		AllTask = append(AllTask,NewTask)
		Confirm := ""
		fmt.Print("< add new task again ? (y/n)")
		fmt.Scanf("%s",&Confirm)
		if Confirm == "n" {
			Button = false
		}
	}
	uuid := string(time.Now().Format("01-02-2006"))
	Filename := "Daily-"+uuid+".json"
	//fmt.Println(Filename)
	User := me()
	Homes,_ := os.UserHomeDir()
	if User == "tuxy" {
		os.Mkdir(Homes+"/.tuxy/Daily",0700)
		if _,err:= os.Stat(Homes+"/.tuxy/Daily/"+Filename);err != nil {
			JSONSTR,_ := json.Marshal(AllTask)
			JSONSTRS := enk.StdEncoding.EncodeToString([]byte(JSONSTR))
			os.WriteFile(Homes+"/.tuxy/Daily/"+Filename,[]byte(JSONSTRS),0700)
			return
		}
		JSONSTRS,_ := os.ReadFile(Homes+"/.tuxy/Daily/"+Filename)
		JSONSTR,_ := enk.StdEncoding.DecodeString(string(JSONSTRS))
		var JSONPARSE []interface{}
		json.Unmarshal(JSONSTR,&JSONPARSE)
		for _,v := range JSONPARSE {
			AllTask = append(AllTask,v.(map[string]interface{}))
		}
		JSONStR,_ := json.Marshal(AllTask)
		JSONStRs := enk.StdEncoding.EncodeToString([]byte(JSONStR))
		os.WriteFile(Homes+"/.tuxy/Daily/"+Filename,[]byte(JSONStRs),0700)
		return
	}
	os.Mkdir(Homes+"/.tuxy_user/"+User+".tuxy/Daily",0700)
	if _,err := os.Stat(Filename);err != nil {
		JSONSTR,_ := json.Marshal(AllTask)
		JSONSTRS := enk.StdEncoding.EncodeToString([]byte(JSONSTR))
		os.WriteFile(Homes+"/.tuxy_user/"+User+".tuxy/Daily/"+Filename,[]byte(JSONSTRS),0700)
		return
	}
	JSONSTRS,_ := os.ReadFile(Homes+"/.tuxy_user/"+User+".tuxy/Daily/"+Filename)
	JSONSTR,_ := enk.StdEncoding.DecodeString(string(JSONSTRS))
	var JSONPARSE []interface{}
	json.Unmarshal(JSONSTR,&JSONPARSE)
	for _,v := range JSONPARSE {
		AllTask = append(AllTask,v.(map[string]interface{}))
	}
	JSONStR,_ := json.Marshal(AllTask)
	JSONStRs := enk.StdEncoding.EncodeToString([]byte(JSONStR))
	os.WriteFile(Homes+"/.tuxy_user/"+User+".tuxy/Daily/"+Filename,[]byte(JSONStRs),0700)
	return
}
func listDailyTask(is interface{}) (int,int) {
	Homes, _ := os.UserHomeDir()
	Filename := "Daily-"+time.Now().Format("01-02-2006")+".json"
	User := me()
	if me() == "tuxy" {
		JSONSTRS,_ := os.ReadFile(Homes+"/.tuxy/Daily/"+Filename)
		JSONSTR,_ := enk.StdEncoding.DecodeString(string(JSONSTRS))
		var JSONPARSE []interface{}
		json.Unmarshal(JSONSTR,&JSONPARSE)
		var yes int
		var no int
		for i,v := range JSONPARSE {
			if v.(map[string]interface{})["In"] == "no" {
				if is != "percent" {
					fmt.Println(" [",i,"] Not do >",v.(map[string]interface{})["Name"])
				}
				no++
			}else if v.(map[string]interface{})["In"] == "yes" {
				if is != "percent" {
					fmt.Println(" [",i,"] Yes do >",v.(map[string]interface{})["Name"])
				}
				yes++
			}
		}
		return yes,no
	}
	JSONSTRS,_ := os.ReadFile(Homes+"/.tuxy_user/"+User+".tuxy/Daily/"+Filename)
	JSONSTR,_ := enk.StdEncoding.DecodeString(string(JSONSTRS))
	var JSONPARSE []interface{}
	json.Unmarshal(JSONSTR,&JSONPARSE)
	yes,no := 0,0
	//fmt.Println()
	for i,v := range JSONPARSE {
		if v.(map[string]interface{})["In"] == "no" {
			if is != "percent" {
				fmt.Println(" [",i,"] Not do >",v.(map[string]interface{})["Name"])
			}
			no++
		}else if v.(map[string]interface{})["In"] == "yes" {
			if is != "percent" {
				fmt.Println(" [",i,"] Yes do >",v.(map[string]interface{})["Name"])
			}
			yes++
		}
	}
	return yes,no
}
func percent(){
	yes,no := listDailyTask("percent")
	allTask := yes+no
	//fmt.Print(yes/allTask)
	yesP := (yes * 100)/allTask 
	noP := (no * 100)/allTask 
	//yesnoP := yesno/allTask * 100
	fmt.Printf("\tNot do Percentage %v%% \n",noP)
	fmt.Printf("\tYes do Percentage %v%% \n",yesP)
	//fmt.Printf("\tYes but Late percentage %v%% \n",yesnoP)
}
func Detail(name string) {
	Homes, _ := os.UserHomeDir()
	Filename := "Daily-"+time.Now().Format("01-02-2006")+".json"
	User := me()
	if me() == "tuxy" {
		JSONSTRS,_ := os.ReadFile(Homes+"/.tuxy/Daily/"+Filename)
		JSONSTR,_ := enk.StdEncoding.DecodeString(string(JSONSTRS))
		var JSONPARSE []interface{}
		json.Unmarshal(JSONSTR,&JSONPARSE)
		for _,v := range JSONPARSE {
			if v.(map[string]interface{})["Name"] == name {
				fmt.Println("Name   :",v.(map[string]interface{})["Name"])
				fmt.Println("Uuid   :",v.(map[string]interface{})["Uuid"])
				fmt.Println("status :",v.(map[string]interface{})["In"])
			}
		}
		return 
	}
	JSONSTRS,_ := os.ReadFile(Homes+"/.tuxy_user/"+User+".tuxy/Daily/"+Filename)
	JSONSTR,_ := enk.StdEncoding.DecodeString(string(JSONSTRS))
	var JSONPARSE []interface{}
	json.Unmarshal(JSONSTR,&JSONPARSE)
	//yes,no,yesno := 0,0,0
	//fmt.Println()
	for _,v := range JSONPARSE {
		if v.(map[string]interface{})["Name"]== name {
			fmt.Println("Name   :",v.(map[string]interface{})["Name"])
			fmt.Println("Uuid   :",v.(map[string]interface{})["Uuid"])
			fmt.Println("status :",v.(map[string]interface{})["In"])
		}
	}
	return 
}
func solve(name string){
	uuid := string(time.Now().Format("01-02-2006"))
	Filename := "Daily-"+uuid+".json"
	//fmt.Println(Filename)
	User := me()
	Homes,_ := os.UserHomeDir()
	//fmt.Print("ac")
	if User == "tuxy" {
		//os.Mkdir(Homes+"/.tuxy/Daily",0700)
		if _,err:= os.Stat(Homes+"/.tuxy/Daily/"+Filename);err != nil {
			fmt.Println("< task file not found")
			return
		}
		JSONSTRS,_ := os.ReadFile(Homes+"/.tuxy/Daily/"+Filename)
		JSONSTR,_ := enk.StdEncoding.DecodeString(string(JSONSTRS))
		var JSONPARSE []interface{}
		json.Unmarshal(JSONSTR,&JSONPARSE)
		for _,v := range JSONPARSE {
			if name == v.(map[string]interface{})["Name"] {
				v.(map[string]interface{})["In"] = "yes"
			}
		}
		JSONStR,_ := json.Marshal(JSONPARSE)
		JSONStRs := enk.StdEncoding.EncodeToString([]byte(JSONStR))
		os.WriteFile(Homes+"/.tuxy/Daily/"+Filename,[]byte(JSONStRs),0700)
		return
	}
	//os.Mkdir(Homes+"/.tuxy_user/"+User+".tuxy/Daily",0700)
	if _,err := os.Stat(Homes+"/.tuxy_user/"+User+".tuxy/Daily/"+Filename);err != nil {
		//fmt.Print(err)
		return
	}
	JSONSTRS,_ := os.ReadFile(Homes+"/.tuxy_user/"+User+".tuxy/Daily/"+Filename)
	JSONSTR,_ := enk.StdEncoding.DecodeString(string(JSONSTRS))
	var JSONPARSE []interface{}
	json.Unmarshal(JSONSTR,&JSONPARSE)
	for _,v := range JSONPARSE {
		if name == v.(map[string]interface{})["Name"] {
			v.(map[string]interface{})["In"] = "yes"
		}
	}
	fmt.Println("< done change",name,"to yes")
	JSONStR,_ := json.Marshal(JSONPARSE)
	JSONStRs := enk.StdEncoding.EncodeToString([]byte(JSONStR))
	os.WriteFile(Homes+"/.tuxy_user/"+User+".tuxy/Daily/"+Filename,[]byte(JSONStRs),0700)
	return
}
func resolve(name string){
	uuid := string(time.Now().Format("01-02-2006"))
	Filename := "Daily-"+uuid+".json"
	//fmt.Println(Filename)
	User := me()
	Homes,_ := os.UserHomeDir()
	//fmt.Print("ac")
	if User == "tuxy" {
		//os.Mkdir(Homes+"/.tuxy/Daily",0700)
		if _,err:= os.Stat(Homes+"/.tuxy/Daily/"+Filename);err != nil {
			fmt.Println("< task file not found")
			return
		}
		JSONSTRS,_ := os.ReadFile(Homes+"/.tuxy/Daily/"+Filename)
		JSONSTR,_ := enk.StdEncoding.DecodeString(string(JSONSTRS))
		var JSONPARSE []interface{}
		json.Unmarshal(JSONSTR,&JSONPARSE)
		for _,v := range JSONPARSE {
			if name == v.(map[string]interface{})["Name"] {
				v.(map[string]interface{})["In"] = "no"
			}
		}
		JSONStR,_ := json.Marshal(JSONPARSE)
		JSONStRs := enk.StdEncoding.EncodeToString([]byte(JSONStR))
		os.WriteFile(Homes+"/.tuxy/Daily/"+Filename,[]byte(JSONStRs),0700)
		return
	}
	//os.Mkdir(Homes+"/.tuxy_user/"+User+".tuxy/Daily",0700)
	if _,err := os.Stat(Homes+"/.tuxy_user/"+User+".tuxy/Daily/"+Filename);err != nil {
		//fmt.Print(err)
		return
	}
	JSONSTRS,_ := os.ReadFile(Homes+"/.tuxy_user/"+User+".tuxy/Daily/"+Filename)
	JSONSTR,_ := enk.StdEncoding.DecodeString(string(JSONSTRS))
	var JSONPARSE []interface{}
	json.Unmarshal(JSONSTR,&JSONPARSE)
	for _,v := range JSONPARSE {
		if name == v.(map[string]interface{})["Name"] {
			v.(map[string]interface{})["In"] = "no"
		}
	}
	fmt.Println("< done change",name,"to no :(")
	JSONStR,_ := json.Marshal(JSONPARSE)
	JSONStRs := enk.StdEncoding.EncodeToString([]byte(JSONStR))
	os.WriteFile(Homes+"/.tuxy_user/"+User+".tuxy/Daily/"+Filename,[]byte(JSONStRs),0700)
	return
}
func helper(){
	command := []string{
		"task - .Create() -- create new task",
		"task - TASKNAME.Solve() -- solved a task by name",
		"task - TASKNAME.Resolve() -- re-solved a task by name",
		"task - TASKNAME.Detail() -- display detail from a task by name",
		"task - .Percent() -- display your progress to day",
		"task - .Help() -- display this message",
		"task - .List() -- list alltask what you has been created",
	}
	fmt.Println("< task parrent command :")
	for _,v := range command {
		fmt.Println("\t",v)
	}
}
