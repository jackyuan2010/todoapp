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

// .List {
//     width: 80%;
//     padding: 1% 2%;
//     display: flex;
//     flex-direction: column;
//     justify-content: center;
//     align-items: center;
// }

// .ListItem {
//     width: 100%;
//     padding: 1%;
//     margin: 0.5%;
//     display: flex;
//     flex-direction: row;
//     justify-content: space-between;
//     align-items: center;
//     background-color: #fafafa;
//     border-radius: 4px;
//     transition: all 0.3s cubic-bezier(0.25, 0.8, 0.25, 1);
// }

// .ListItem:hover {
//     box-shadow: 0 3px 6px rgba(0, 0, 0, 0.16), 0 3px 6px rgba(0, 0, 0, 0.23);
// }

// .Status {
//     transition: all 0.3s cubic-bezier(0.25, 0.8, 0.25, 1);
// }

// .Status:hover {
//     cursor: pointer;
//     transform: scale(1.05);
// }

// .Title {
//     width: 80%;
//     display: flex;
//     flex-direction: row;
//     justify-content: flex-start;
//     align-items: center;
// }

// .RemoveItem {
//     color: #878787;
//     margin-right: 1%;
//     font-weight: bold;
//     transform: scale(0.9);
//     transition: all 0.3s cubic-bezier(0.25, 0.8, 0.25, 1);
// }

// .RemoveItem:hover {
//     color: #e81744;
//     cursor: pointer;
// }