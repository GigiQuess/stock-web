import React, { useState, useEffect } from 'react';
import './App.css';

interface User {
  id: number;
  name: string;
  email: string;
}

function App() {
  const [message, setMessage] = useState<string>('');
  const [users, setUsers] = useState<User[]>([]);
  const [name, setName] = useState<string>('');
  const [email, setEmail] = useState<string>('');

  useEffect(() => {
    fetch('/api/ping')
      .then(res => res.json())
      .then(data => setMessage(data.message));
  }, []);

  const createUser = (e: React.FormEvent) => {
    e.preventDefault();
    fetch('/api/users', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ name, email }),
    })
      .then(res => res.json())
      .then((newUser: User) => {
        setUsers([...users, newUser]);
        setName('');
        setEmail('');
      });
  };

  return (
    <div className="App">
      <header className="App-header">
        <p>
          Message from backend: {message}
        </p>
        <form onSubmit={createUser}>
          <input
            type="text"
            placeholder="Name"
            value={name}
            onChange={(e: React.ChangeEvent<HTMLInputElement>) => setName(e.target.value)}
          />
          <input
            type="email"
            placeholder="Email"
            value={email}
            onChange={(e: React.ChangeEvent<HTMLInputElement>) => setEmail(e.target.value)}
          />
          <button type="submit">Create User</button>
        </form>
        <h2>Users</h2>
        <ul>
          {users.map(user => (
            <li key={user.id}>{user.name} - {user.email}</li>
          ))}
        </ul>
      </header>
    </div>
  );
}

export default App;
