import React, {Component} from 'react';
import PropTypes from "prop-types";

class AddTodo extends Component {
    state = {
        title: ''
    };

    updateState = (event) =>
        this.setState({
            // event.target.name will get whatever 'name' tag is equal to
            [event.target.name]: event.target.value
        });

    onSubmit = (event) => {
        // prevent submission to the file
        event.preventDefault();

        this.props.addTodo(this.state.title);

        // reset state
        this.setState({title: ''});
    };

    render() {
        return (
            <form onSubmit={this.onSubmit} style={{display: 'flex'}}>
                <input type="text"
                       name="title"
                       style={{flex: '10', padding: '5px'}}
                       placeholder="Add todo ..."
                       value={this.state.title}
                       onChange={this.updateState}/>

                <input type="submit"
                       value="Submit"
                       className="btn"
                       style={{flex: '1'}}/>
            </form>
        );
    }
}

AddTodo.propTypes = {
    addTodo: PropTypes.func.isRequired
};

export default AddTodo;