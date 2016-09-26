import React, {Component} from 'react'

export default class TicketDescription extends Component {
  render() {
    return <div className="description">
      {this.props.description}
    </div>
  }
}

