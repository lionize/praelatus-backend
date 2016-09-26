import React, {Component} from 'react'

export default class TicketStatus extends Component {
	render() {
		return <div className="row status-row">
            <div className="col-md-6">
              <span className="text-muted">Status:</span>
            </div>
            <div className="col-md-6">
              <span>{this.props.status.name}</span>
            </div>
          </div> 
	}
}
