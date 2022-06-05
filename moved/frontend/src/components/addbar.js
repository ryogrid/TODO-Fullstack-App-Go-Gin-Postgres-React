import React from "react";
import "./styles/Addbar.css";

class AddBar extends React.Component {
  bkend_url_base = null

  constructor(props) {
    super(props);
    if(window.location.hostname == "localhost"){
      this.bkend_url_base = "http://localhost:8080";
    }else{
      this.bkend_url_base = "https://" + window.location.hostname;
    }

    //this.props = props;
  }

  //props = null;

  addItem = event => {
    if (event.key === "Enter") {
      fetch(this.bkend_url_base + `/item/create/${event.target.value}`).then(
        //alert("Bug: Reload The Page To View Changes")
        //this.props.emitVal(Math.random())
        window.location.reload()
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
