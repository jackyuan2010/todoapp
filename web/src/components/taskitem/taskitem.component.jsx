import React from "react";
import {TaskItemContainer, TaskItemDetail, RemoveTaskItem, TaskItemDone} from "./taskitem.style"

const TaskItem = ({taskItem}) => {
    const { Task_Description, Is_Finished } = taskItem;
    return (
        <TaskItemContainer>
            <TaskItemDetail>
                <span>{ Task_Description }</span>
                <RemoveTaskItem>X</RemoveTaskItem>
            </TaskItemDetail>
            <TaskItemDone>{Is_Finished === true ? "Done" : "Not Done"}</TaskItemDone>
        </TaskItemContainer>
    );
};

export default TaskItem;