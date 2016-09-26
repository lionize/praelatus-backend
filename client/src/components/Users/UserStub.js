import React, {Component} from 'react'

export default class UserStub extends Component {
	render() {
		return (
			<div>
				<img src={this.props.profilePicture} />
				<span>{this.props.fullName}</span>
			</div>
		)
	}
}
