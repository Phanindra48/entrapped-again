import React from 'react/addons';

import Api from '../api/api.js';

class Block extends React.Component {
  constructor(props) {
    super(props);
    this.handleClick = this.handleClick.bind(this);
  }

  render() {
    var cx = React.addons.classSet;
    var classes = cx({
      'block': this.props.of !== "player",
      'block__enemy': this.props.of === "player",
      'block__visited': this.props.status === 0 || this.props.status === 9,
      'fa fa-2x fa-bomb block__dead': this.props.status === -2
    });

    return(
      <td className={classes} onClick={this.handleClick}></td>
    )
  }

  handleClick() {
    if (this.props.of === "player") {
      return;
    }

    var msg = "data:open:[idx=" + this.props.index + "]:[help=9]";
    Api.send(msg);

    /* beep is a global function intended for fun */
    beep();
  }
};
Block.propTypes = { status: React.PropTypes.number, index: React.PropTypes.number, of: React.PropTypes.string };
Block.defaultProps = { status: 0, index: 0, of: null };

class Blocks extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    var nodes = [];
    var size = this.props.size;

    /* construct blocks */
    for (var i = 0; i < size; i++) {
      var _blocks = [];

      for (var j = 0; j < size; j++) {
        var index = i * size + j;
        _blocks.push(<Block key={'block-' + index} status={this.props.mines[index]} index={index} of={this.props.of} />);
      }

      nodes.push((<tr key={'blocks-' + i}>{_blocks}</tr>));
    }

    return (
      <table className="blocks">
        <tbody>
          {nodes}
        </tbody>
      </table>
    );
  }
};
Blocks.propTypes = { size: React.PropTypes.number, mines: React.PropTypes.array, of: React.PropTypes.string };
Blocks.defaultProps = { size: 7, mines: [], of: null }

export default Blocks;
