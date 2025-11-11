const express = require('express');
const cors = require('cors');
const bodyParser = require('body-parser');

const authRoutes = require('./routes/auth');
const coffeeRoutes = require('./routes/coffees');

const app = express();
const PORT = 3000;

app.use(cors()); 
app.use(bodyParser.json()); 

app.use('/api/auth', authRoutes);
app.use('/api/coffees', coffeeRoutes);

app.get('/', (req, res) => {
  res.json({ message: 'Servidor funcionando correctamente' });
});

app.listen(PORT, () => {
  console.log(`🚀 Servidor corriendo en http://localhost:${PORT}`);
});