const express = require('express');
const app = express();
const port = 3000;

app.get('/', (req, res) => {
  res.send(`
    <div style="text-align: center; margin-top: 50px; font-family: sans-serif;">
      <h1 style="color: #4CAF50;">Auth Service</h1>
      <p>Status: 🟢 <strong>Online</strong></p>
      <p>Version: 1.0</p>
      <p>Technlogy: Node.js</p>
    </div>
  `);
});

app.listen(port, () => {
  console.log(`Auth service listening at http://localhost:${port}`);
});
