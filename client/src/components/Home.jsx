import React from 'react';
import { Router, Link } from 'react-router';

import MinesStore from '../stores/minestore.js';

class Home extends React.Component {
  handleChange(e) {
    MinesStore.addUsername(e.target.value);
  }
  validatePlayerName(){
    //show some custom message if name is empty
    var name = MinesStore.getUsername();
    if(name == ""){
      alert("Enter nickname to play");
      return false;
    }
    return true;
  }
  render() {
    return (
      <div className="home">
        <h1 className="home__title">Entrapped v0.0</h1>
        <blockquote>Tiny little minefield of our dreams.</blockquote>
        <div className="nickname">
          <input className="nickname__input" placeholder="nickname" onChange={this.handleChange}/>
        </div>
        <Link to="/minefield" onClick={this.validatePlayerName}>
        Play
        </Link>
      </div>
    )
  }
};

export default Home;
