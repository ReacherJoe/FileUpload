import  Todoitem  from "./Todoitem"

 function TodoList ({todos , toggleTodo ,deleteTodo} ) {
  return ( 
  <ul className="list">
    {todos.length === 0 && "No Todos"}
    {todos.map((todo, key) => {
      return <Todoitem {...todo}  key={key} toggleTodo={toggleTodo}
      deleteTodo={deleteTodo}/>
    })}
    
  </ul>
  )
}
export default  TodoList;