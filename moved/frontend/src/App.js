import React from "react";

import Header from "./components/header";
import AddBar from "./components/addbar";
import TodoList from "./components/todolist";

import "./App.css";

class App extends React.Component {  

  render() {
    const [val, setVal] = React.useState();

    return (
      <div className="App">
        <Header />
        <AddBar emitVal={setVal} />
        <TodoList val={val} />
      </div>
    );
  }
}

export default App;
