# Gogen

//gogen domain todocore
//gogen entity Todo
//gogen error MessageMustNotEmpty
//gogen usecase RunTodoCreate
//gogen usecase GetAllTodo
//gogen usecase RunTodoCheck
//gogen repository SaveTodo Todo RunTodoCreate
//gogen repository FindAllTodo Todo GetAllTodo
//gogen repository FindOneTodoByID Todo RunTodoCheck
//gogen error TodoHasBeenChecked
//gogen gateway withgorm
//gogen controller resapi
//gogen application todoapp
//go run main.go todoapp
