import { useState } from "react";
import TodoList from "../components/TodoList";

const Home = () => {
  const [todos, setTodos] = useState([]);

  const deleteTodo = (id) => {
    setTodos(todos.filter((todo) => todo.id !== id));
  };

  const toggleTodo = (id) => {
    setTodos(
      todos.map((todo) =>
        todo.id === id ? { ...todo, completed: !todo.completed } : todo
      )
    );
  };

  return (
    <div className="p-4 max-w-2xl mx-auto bg-transparent transition-colors">
      {/* Todo List Section */}
      <TodoList todos={todos} onDelete={deleteTodo} onToggle={toggleTodo} />
    </div>
  );
};

export default Home;
