import React, { Component } from 'react';
import PropTypes from 'prop-types';

class TodoItem extends Component{
    getStyle = () => {
        return {
            textDecoration: this.props.todo.completed
                ? 'line-through'
                : 'none',
            background: 'lightGray',
            padding: '10px',
            borderBottom: '1px #ccc dotted'
        }
    }

    render() {
        const { id, title } = this.props.todo;

        return (
            <div style={this.getStyle()}>
                <p>
                    <input type="checkbox" onChange={this.props.markComplete.bind(this, id)} /> {' '}
                    {title}
                    <button onClick={this.props.delTodo.bind(this, id)} style={btnDeleteStyle}>x</button>
                </p>
            </div>
        );
    }
}

TodoItem.propTypes = {
    todo: PropTypes.object.isRequired,
    markComplete: PropTypes.func.isRequired,
    delTodo: PropTypes.func.isRequired
}

const btnDeleteStyle = {
    background: 'red',
    color: 'white',
    border: 'none',
    padding: '5px 9px',
    borderRadius: '50%',
    cursor: 'pointer',
    float: 'right'
};

export default TodoItem;
