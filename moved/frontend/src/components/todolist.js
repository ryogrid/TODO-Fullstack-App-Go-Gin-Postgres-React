import React from "react";
import "./styles/Todolist.css";

class Todolist extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      items: [],
      randToReload: props.val
    };
  }

  removeItem(id) {
    fetch(`https://golang-todo-with-samehada.herokuapp.com/item/delete/${id}`).then(
      this.setState({
        items: this.state.items.filter(item => item.id !== id),
      })
    );
  }

  toggleDone(id) {
    let items = [...this.state.items];
    let item = items.find(item => item.id === id);
    item.done = !item.done;

    fetch(`https://golang-todo-with-samehada.herokuapp.com/item/update/${id}/${item.done}`).then(
      this.setState({ items })
    );
  }

  isDone(done) {
    if (done) {
      return "Done";
    } else {
      return "Not Done";
    }
  }

  createItem(item) {
    return (
      <div className="ListItem" key={item.id} id={item.id}>
        <div className="Title">
          <div className="RemoveItem" onClick={() => this.removeItem(item.id)}>
            X
          </div>
          {item.item}
        </div>
        <div className="Status" onClick={() => this.toggleDone(item.id)}>
          {this.isDone(item.done)}
        </div>
      </div>
    );
  }

  componentDidMount() {
    fetch("https://golang-todo-with-samehada.herokuapp.com/items")
      .then(res => res.json())
      .then(json => this.setState({ items: json.items }));
  }

  render() {
    var items = this.state.items;
    return (
      <div className="TodoList">
        <div className="List">{items.map(item => this.createItem(item))}</div>
      </div>
    );
  }
}

export default Todolist;
