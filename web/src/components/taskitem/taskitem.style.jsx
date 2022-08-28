import styled  from "styled-components";

export const TaskItemContainer = styled.div`
    width: 100%;
    padding: 1%;
    margin: 0.5%;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    background-color: #fafafa;
    border-radius: 4px;
    transition: all 0.3s cubic-bezier(0.25, 0.8, 0.25, 1); 
    &:hover {
        box-shadow: 0 3px 6px rgba(0, 0, 0, 0.16), 0 3px 6px rgba(0, 0, 0, 0.23);
    }  
`;

export const TaskItemDetail = styled.div`
    width: 80%;
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    align-items: center;       
`;

export const RemoveTaskItem = styled.div`
    color: #878787;
    margin-right: 10%;
    font-weight: bold;
    transform: scale(0.9);
    transition: all 0.3s cubic-bezier(0.25, 0.8, 0.25, 1);  
    
    &:hover {
        color: #e81744;
        cursor: pointer;
    }
`;

export const TaskItemDone = styled.div`
    transition: all 0.3s cubic-bezier(0.25, 0.8, 0.25, 1); 
    &:hover {
        cursor: pointer;
        transform: scale(1.05);
    }     
`;