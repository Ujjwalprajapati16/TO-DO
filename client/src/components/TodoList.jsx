import { useState } from "react";
import {
  useFetchTodos,
  useCreateTodo,
  useUpdateTodo,
  useDeleteTodo,
} from "../hooks/useTodos";
import { FaTrash, FaCheck } from "react-icons/fa";

const TodoList = () => {
  const { data: todos, isLoading } = useFetchTodos();
  const createTodo = useCreateTodo();
  const updateTodo = useUpdateTodo();
  const deleteTodo = useDeleteTodo();
  const [newTodo, setNewTodo] = useState("");

  const handleAddTodo = () => {
    if (newTodo.trim()) {
      createTodo.mutate({ body: newTodo, completed: false });
      setNewTodo("");
    }
  };

  if (isLoading) return <p>Loading todos...</p>;

  return (
    <div className="p-4 max-w-md mx-auto bg-quinary dark:bg-gray-800 rounded-lg shadow-md">
      <h1 className="text-xl font-bold text-center mb-4">Todo List</h1>
      <div className="flex mb-4">
        <input
          type="text"
          className="flex-1 p-2 border rounded-l-lg focus:outline-none"
          placeholder="Add a new todo..."
          value={newTodo}
          onChange={(e) => setNewTodo(e.target.value)}
        />
        <button
          className="p-2 bg-blue-500 text-white rounded-r-lg"
          onClick={handleAddTodo}
        >
          Add
        </button>
      </div>
      <ul>
        {todos.map((todo) => (
          <li
            key={todo._id}
            className={`flex justify-between items-center p-2 border-b ${
              todo.completed ? "line-through text-gray-500" : ""
            }`}
          >
            <span>{todo.body}</span>
            <div className="flex space-x-2">
              {!todo.completed && (
                <button
                  className="text-green-500"
                  onClick={() => updateTodo.mutate(todo._id)}
                >
                  <FaCheck />
                </button>
              )}
              <button
                className="text-red-500"
                onClick={() => deleteTodo.mutate(todo._id)}
              >
                <FaTrash />
              </button>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default TodoList;
