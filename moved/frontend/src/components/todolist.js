import React from "react";
import "./styles/Todolist.css";

class Todolist extends React.Component {
  bkend_url_base = null

  constructor(props) {
    super(props);
    this.state = {
      items: [],
      //randToReload: props.val
    };

    if(window.location.hostname == "localhost"){
      this.bkend_url_base = "http://localhost:8080";
    }else{
      this.bkend_url_base = "https://" + window.location.hostname;
    }
  }

  

  removeItem(id) {
    fetch(this.bkend_url_base + `/item/delete/${id}`).then(
      this.setState({
        items: this.state.items.filter(item => item.id !== id),
      })
    );
  }

  toggleDone(id) {
    let items = [...this.state.items];
    let item = items.find(item => item.id === id);
    item.done = !item.done;

    fetch(this.bkend_url_base + `/item/update/${id}/${item.done}`).then(
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
    fetch(this.bkend_url_base + "/items")
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
