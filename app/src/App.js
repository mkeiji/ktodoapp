import React, { Component } from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import Todos from "./components/todos/Todos";
import './App.css';
import Header from "./components/layouts/Header";
import AddTodo from "./components/todos/AddTodo";
import About from "./components/pages/About";
import axios from "axios";

const appUrlRoot = 'http://ktodoapp.info.tm';

class App extends Component{
    placeholderUrls = {
        getMany: 'https://jsonplaceholder.typicode.com/todos?_limit=10',
        delete: 'https://jsonplaceholder.typicode.com/todos/',
        post: 'https://jsonplaceholder.typicode.com/todos'
    };

    kserver = {
        getMany: `${appUrlRoot}/server/todos`,
        delete: `${appUrlRoot}/server/todoItem/`,
        post: `${appUrlRoot}/server/todoItem`
    };

    state = {
        todos: []
    };

    componentDidMount() {
        axios.get(this.kserver.getMany)
            .then(response =>
                this.setState({
                    todos: response.data
                }))
    }

    markComplete = (id) => {
        this.setState({
            todos: this.state.todos.map((todo) => {
                if (todo.Id === id) {
                    todo.Completed = !todo.Completed
                }
                return todo;
            })
        });
    };

    delTodo = (id) => {
        axios.delete(this.kserver.delete + id)
            .then(() => this.filterIdOut(id));
    };

    filterIdOut = (id) => {
        this.setState({
            todos: [
                ...this.state.todos.filter((todo) => todo.Id !== id)
            ]
        });
    };

    addTodo = (title) => {
        const newTodo = {
            Title: title,
            Completed: false
        };

        // can use .catch to deal with exceptions after .then
        axios.post(this.kserver.post, newTodo)
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
            Id: this.generateID(),
            Title: title,
            Completed: false
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

        this.state.todos.map(todo => idArray.push(todo.Id));
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
