import React, {Component} from 'react'

export default class IssueType extends Component {
  render() {
    return <div className="row issue-type-row">
      <div className="col-md-6">
        <span className="text-muted">Issue Type:</span>
      </div>
      <div className="col-md-6">
        <span>{this.props.ticketType}</span>
      </div>
    </div> 
  }
}
