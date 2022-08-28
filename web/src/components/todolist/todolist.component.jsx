import React from "react";
import {TodoListContainer, TaskItemsContainer, EmptyMessage} from "./todolist.style"
import TaskItem from "../taskitem/taskitem.component";

const ToDoList  = () => {
    const taskItems = [{Id: "1", Task_Description: "To do task 1", Is_Finished: true}, 
        {Id: "2", Task_Description: "To do task 2", Is_Finished: false}];
    return (
        <TodoListContainer>
          <TaskItemsContainer>
            {taskItems.length ? (taskItems.map((task) => <TaskItem key={task.Id} taskItem={task} />)) 
                : (<EmptyMessage>Your To Do List is empty</EmptyMessage>)}
          </TaskItemsContainer>
        </TodoListContainer>
      );
};

export default ToDoList;