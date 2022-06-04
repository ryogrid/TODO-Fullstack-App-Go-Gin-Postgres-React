import React from "react";
import "./styles/Addbar.css";

class AddBar extends React.Component {
  addItem = event => {
    if (event.key === "Enter") {
      fetch(`https://golang-todo-with-samehada.herokuapp.com/item/create/${event.target.value}`).then(
        alert("Bug: Reload The Page To View Changes")
      );
    }
  };

  render() {
    return (
      <div className="AddBar">
        <input
          className="AddBar-Text"
          type="text"
          placeholder="Enter TODO Item"
          onKeyDown={this.addItem}
        />
      </div>
    );
  }
}

export default AddBar;
