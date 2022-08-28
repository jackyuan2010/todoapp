import React from 'react';
import './App.css';
import ToDoHeader from './components/todoheader/todoheader.component';
import ToDoList from './components/todolist/todolist.component';

const App = () => {
  return (
    <div className="App">
      <ToDoHeader/>
      <ToDoList/>
    </div>
  ); 
};

export default App;
