import React, {Component} from 'react'

export default class TicketKey extends Component {
  render() {
    return <div className="row key-row">
      <div className="col-md-6">
        <span className="text-muted">Key:</span>
      </div>
      <div className="col-md-6">
        <span>{this.props.ticketKey}</span>
      </div>
    </div> 
  }
}


