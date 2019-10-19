import React, { Component } from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import Todos from "./components/todos/Todos";
import './App.css';
import Header from "./components/layouts/Header";
import AddTodo from "./components/todos/AddTodo";
import About from "./components/pages/About";
import axios from "axios";

class App extends Component{
    state = {
        todos: []
    };

    componentDidMount() {
        axios.get('https://jsonplaceholder.typicode.com/todos?_limit=10')
            .then(response =>
                this.setState({
                    todos: response.data
                }))
    }

    markComplete = (id) => {
        this.setState({
            todos: this.state.todos.map((todo) => {
                if (todo.id === id) {
                    todo.completed = !todo.completed
                }
                return todo;
            })
        });
    };

    delTodo = (id) => {
        axios.delete(`https://jsonplaceholder.typicode.com/todos/${id}`)
            .then(() => this.filterIdOut(id));
    };

    filterIdOut = (id) => {
        this.setState({
            todos: [
                ...this.state.todos.filter((todo) => todo.id !== id)
            ]
        });
    };

    addTodo = (title) => {
        const newTodo = {
            title: title,
            completed: false
        };

        // can use .catch to deal with exceptions after .then
        axios.post('https://jsonplaceholder.typicode.com/todos', newTodo)
            .then(response =>
                this.setState({
                    todos: [
                        ...this.state.todos, response.data
                    ]
                })
            );
    };

    addTodoWithoutAPI = (title) => {
        const newTodo = {
            id: this.generateID(),
            title: title,
            completed: false
        };

        this.setState({
            todos: [
                ...this.state.todos,
                newTodo
            ]
        })
    };

    generateID() {
        let idArray = [];

        this.state.todos.map(todo => idArray.push(todo.id));
        idArray.sort(function(a, b) {return a-b});

        return this.getLastItem(idArray) + 1;
    }

    getLastItem(array) {
        return array.slice(-1)[0];
    }

    render() {
        return (
            <Router>
                <div className="App">
                    <div className="container">
                        <Header />

                        <Route exact path="/" render={props => (
                            <React.Fragment>
                                <AddTodo addTodo={this.addTodo}/>
                                <Todos todos={this.state.todos}
                                       markComplete={this.markComplete}
                                       delTodo={this.delTodo}/>
                            </React.Fragment>
                        )}/>

                        <Route path="/about" component={About}/>

                    </div>
                </div>
            </Router>
        );
    }
}

export default App;
