const express = require('express');
const router = express.Router();
const fs = require('fs');
const path = require('path');

const usersFile = path.join(__dirname, '../data/users.json');

const getUsers = () => {
  try {
    const data = fs.readFileSync(usersFile, 'utf8');
    return JSON.parse(data);
  } catch (error) {
    return [];
  }
};

const saveUsers = (users) => {
  fs.writeFileSync(usersFile, JSON.stringify(users, null, 2));
};

router.post('/register', (req, res) => {
  const { username, email, password } = req.body;

  if (!username || !email || !password) {
    return res.status(400).json({ 
      success: false, 
      message: 'Todos los campos son obligatorios' 
    });
  }

  const users = getUsers();

  const userExists = users.find(user => user.email === email);
  if (userExists) {
    return res.status(400).json({ 
      success: false, 
      message: 'El email ya está registrado' 
    });
  }

  const newUser = {
    id: Date.now(),
    username,
    email,
    password 
  };

  users.push(newUser);
  saveUsers(users);

  res.status(201).json({ 
    success: true, 
    message: 'Usuario registrado exitosamente',
    user: { id: newUser.id, username: newUser.username, email: newUser.email }
  });
});

router.post('/login', (req, res) => {
  const { email, password } = req.body;

  if (!email || !password) {
    return res.status(400).json({ 
      success: false, 
      message: 'Email y contraseña son obligatorios' 
    });
  }

  const users = getUsers();

  const user = users.find(u => u.email === email && u.password === password);

  if (!user) {
    return res.status(401).json({ 
      success: false, 
      message: 'Credenciales incorrectas' 
    });
  }

  res.json({ 
    success: true, 
    message: 'Inicio de sesión exitoso',
    user: { id: user.id, username: user.username, email: user.email }
  });
});

module.exports = router;