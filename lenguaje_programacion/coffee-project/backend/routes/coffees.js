const express = require('express');
const router = express.Router();
const fs = require('fs');
const path = require('path');

const coffeesFile = path.join(__dirname, '../data/coffees.json');

router.get('/', (req, res) => {
  try {
    const data = fs.readFileSync(coffeesFile, 'utf8');
    const coffees = JSON.parse(data);
    res.json({ success: true, coffees });
  } catch (error) {
    res.status(500).json({ 
      success: false, 
      message: 'Error al obtener cafés' 
    });
  }
});

module.exports = router;