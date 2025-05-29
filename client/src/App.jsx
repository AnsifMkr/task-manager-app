import React, { useState, useEffect } from "react";
import axios from "axios";
import "./App.css";

const API_URL = "http://localhost:8080";

function App() {
  const [todos, setTodos] = useState([]);
  const [newTodo, setNewTodo] = useState("");

  useEffect(() => {
    fetchTodos();
  }, []);

  const fetchTodos = async () => {
    const res = await axios.get(`${API_URL}/api/todos`);
    setTodos(res.data);
  };

  const addTodo = async () => {
    await axios.post(`${API_URL}/api/todos`, { title: newTodo, done: false });
    setNewTodo("");
    fetchTodos();
  };

  const toggleTodo = async (id, done) => {
    const todo = todos.find(t => t.id === id);
    await axios.put(`${API_URL}/api/todos/${id}`, { ...todo, done: !done });
    fetchTodos();
  };

  const deleteTodo = async (id) => {
    await axios.delete(`${API_URL}/api/todos/${id}`);
    fetchTodos();
  };

  return (
    <div className="App">
      <h1>Task Manager</h1>
      <div>
          <input value={newTodo} onChange={e => setNewTodo(e.target.value)} />
          <button onClick={addTodo}>+</button>
      </div>
      <ul className="lists">
        {todos.map(todo => (
        <li key={todo.id}>
          <input
            type="checkbox"
            checked={todo.done}
            onChange={() => toggleTodo(todo.id, todo.done)}
          />
          <span className={todo.done ? "done" : ""}>{todo.title}</span>
          <button onClick={() => deleteTodo(todo.id)}>Delete</button>
        </li>
        ))}
      </ul>


    </div>
  );
}

export default App;
