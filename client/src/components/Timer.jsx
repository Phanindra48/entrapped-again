import React from 'react/addons';

var _ = require('../../src/utility/utils.js');

class Timer extends React.Component {
  componentDidMount(){
    var fiveMinutes = 60 * 5,
        display = document.querySelector('.timer');
    _.startTimer(fiveMinutes, display);
  }  

  render() {
    return(
      <div >Countdown Timer :<span className='timer'></span></div>
    )
  }

};
export default Timer;