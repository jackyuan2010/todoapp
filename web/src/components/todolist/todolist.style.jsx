import styled from "styled-components"

export const TodoListContainer =styled.div`
    width: 100%;
    /* height: 80%; */
    display: block;
    overflow: auto;
    padding: 2% 0;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
`

export const TaskItemsContainer =styled.div`
    width: 80%;
    padding: 1% 2%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
`
export const EmptyMessage = styled.span`
  font-size: 18px;
  margin: 10px auto;
`;