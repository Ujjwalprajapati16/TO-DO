import { motion } from "framer-motion";

const TodoItem = ({ todo, onDelete, onToggle }) => (
  <motion.div
    className={`p-4 flex flex-col sm:flex-row justify-between items-center border rounded-lg transition-colors ${
      todo.completed
        ? "bg-green-100 dark:bg-primary"
        : "bg-white dark:bg-tertiary"
    }`}
    whileHover={{ scale: 1.05 }}
  >
    <div
      onClick={() => onToggle(todo.id)}
      className="cursor-pointer text-black dark:text-white"
    >
      <h2
        className={`text-lg ${todo.completed ? "line-through" : ""} ${
          todo.completed ? "text-gray-600 dark:text-gray-300" : ""
        }`}
      >
        {todo.title}
      </h2>
    </div>

    <button
      onClick={() => onDelete(todo.id)}
      className="mt-2 sm:mt-0 bg-red-500 text-white px-4 py-2 rounded-lg transition-all hover:bg-red-600"
    >
      Delete
    </button>
  </motion.div>
);

export default TodoItem;
