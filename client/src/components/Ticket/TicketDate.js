import React, {Component} from 'react'

export default class TicketDate extends Component {
	render() {
		return (
			<div className="row key-row">
      			<div className="col-md-6">
      			  <span className="text-muted">{this.props.title}</span>
      			</div>
      			<div className="col-md-6">
      			  <span>{new Date(this.props.date).toLocaleString()}</span>
      			</div>
    		</div> 
		)
	}
}
