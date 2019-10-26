import React from 'react';
import { Link } from "react-router-dom";

function Header() {
    return (
        <header style={headerStyle}>
            <img src="/klogo.png" style={logoStyle} />
            <h1>TodoList</h1>
            <Link style={linkStyle} to="/">Home</Link> | <Link style={linkStyle} to="/about">About</Link>
        </header>
    );
}

const logoStyle = {
    height: '80px',
    width: '80px'
};

const linkStyle = {
    color: '#fff',
    textDecoration: 'none'
};

const headerStyle = {
    background: 'black',
    color: 'white',
    textAlign: 'center',
    padding: '10px'
};

export default Header;