// src/hooks/useTodos.js
import { useQuery, useMutation, useQueryClient } from "react-query";
import axiosInstance from "../utils/axiosInstance";

// Fetch all todos
export const useFetchTodos = () => {
  return useQuery("todos", async () => {
    const { data } = await axiosInstance.get("/todos");
    return data;
  });
};

// Create a new todo
export const useCreateTodo = () => {
  const queryClient = useQueryClient();
  return useMutation(
    async (newTodo) => {
      const { data } = await axiosInstance.post("/todos", newTodo);
      return data;
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries("todos"); // Refresh todos
      },
    }
  );
};

// Update a todo (mark as completed)
export const useUpdateTodo = () => {
  const queryClient = useQueryClient();
  return useMutation(
    async (id) => {
      const { data } = await axiosInstance.patch(`/todos/${id}`);
      return data;
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries("todos"); // Refresh todos
      },
    }
  );
};

// Delete a todo
export const useDeleteTodo = () => {
  const queryClient = useQueryClient();
  return useMutation(
    async (id) => {
      const { data } = await axiosInstance.delete(`/todos/${id}`);
      return data;
    },
    {
      onSuccess: () => {
        queryClient.invalidateQueries("todos"); // Refresh todos
      },
    }
  );
};
