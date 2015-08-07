import React from 'react';
import Timer from './Timer.jsx';
import Field from './Field.jsx';
import MinesStore from '../stores/minestore.js';
import Api from '../api/api.js';
var _ = require('../../src/utility/utils.js');
class Minefield extends React.Component {

  exitGame(){
    if(confirm('Are you sure you want to exit the game?'))
      return true;
  }

  constructor() {
    super();

    this.state = {
      player: MinesStore.getUser(),
      enemy: MinesStore.getEnemy()
    };

    // connect
    Api.connect(MinesStore.getUsername());
    MinesStore.addChangeListener(this._onChange.bind(this));
  }

  componentWillUnmount() {
    MinesStore.removeChangeListener(this._onChange);
  }
  
  _onChange() {
    this.setState({
      player: MinesStore.getUser(),
      enemy: MinesStore.getEnemy()
    });
  }

  render() {
    var content = (
      <div className="minefield__wrapper">
        <Timer />
        <Field of="enemy"  player={this.state.player} />
        <Field of="player" player={this.state.enemy} />
        <br/>
        <button id='exitGame' className='btn btn-sm btn-danger' onClick={this.exitGame}> Exit </button>
      </div>
    );

    if (!this.state.player.mines || !this.state.player.mines.length) {
      content = (
        <div className="minefield__waiting">
          <h1 className="minefield__text--waiting">Waiting for an opponent...</h1>
        </div>
      );
    }

    if (!this.state.player.life || !this.state.enemy.life) {
      var winner = this.state.player.life ? this.state.player.username : this.state.enemy.username;
      content = (
        <div className="minefield__waiting">
          <h1 className="minefield__text--waiting">Game over. {winner} won.</h1>
        </div>
      )
    }

    if (!this.state.player.life || !this.state.enemy.life) {
      var gameEndedBy = this.state.player.life ? this.state.player.username : this.state.enemy.username;
      content = (
        <div className="minefield__waiting">
          <h1 className="minefield__text--waiting">Game ended by {winner}.</h1>
        </div>
      )
    }

    return (
      <div>
        {content}
      </div>
    )
  }
};

export default Minefield;
